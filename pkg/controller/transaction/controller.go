// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transaction

import (
	"context"
	"time"

	"github.com/onosproject/onos-config/pkg/utils/tree"

	"github.com/onosproject/onos-config/pkg/utils"

	"github.com/onosproject/onos-config/pkg/pluginregistry"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	configapi "github.com/onosproject/onos-api/go/onos/config/v2"
	"github.com/onosproject/onos-config/pkg/store/configuration"
	"github.com/onosproject/onos-config/pkg/store/topo"
	"github.com/onosproject/onos-config/pkg/store/transaction"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

const defaultTimeout = 30 * time.Second

var log = logging.GetLogger("controller", "transaction")

// NewController returns a new control relation  controller
func NewController(topo topo.Store, transactions transaction.Store, configurations configuration.Store, pluginRegistry *pluginregistry.PluginRegistry) *controller.Controller {
	c := controller.NewController("transaction")

	c.Watch(&Watcher{
		transactions: transactions,
	})

	c.Watch(&ConfigurationWatcher{
		transactions:   transactions,
		configurations: configurations,
	})

	c.Reconcile(&Reconciler{
		transactions:   transactions,
		topo:           topo,
		configurations: configurations,
		pluginRegistry: pluginRegistry,
	})

	return c
}

// Reconciler reconciles transactions
type Reconciler struct {
	topo           topo.Store
	transactions   transaction.Store
	configurations configuration.Store
	pluginRegistry *pluginregistry.PluginRegistry
}

// Reconcile reconciles transactions
func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	index := id.Value.(configapi.Index)
	transaction, err := r.transactions.GetByIndex(ctx, index)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Failed to reconcile transaction %d, %s", index, err)
			return controller.Result{}, err
		}
		log.Debugf("Transaction %d not found", index)
		return controller.Result{}, nil
	}

	log.Infof("Reconciling transaction %v", transaction)

	// If the transaction is in a completed state, queue the next transaction to be reconciled
	if transaction.Status.State == configapi.TransactionState_TRANSACTION_COMPLETE ||
		transaction.Status.State == configapi.TransactionState_TRANSACTION_FAILED {
		log.Debugf("Queueing next transaction %d", index+1)
		return controller.Result{Requeue: controller.NewID(index + 1)}, nil
	}

	if ok, err := r.reconcileTransaction(ctx, transaction); err != nil {
		log.Warnf("Failed to reconcile transaction: %d, %s", index, err)
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileTransaction(ctx context.Context, transaction *configapi.Transaction) (bool, error) {
	switch t := transaction.Transaction.(type) {
	case *configapi.Transaction_Change:
		return r.reconcileTransactionChange(ctx, transaction, t.Change)
	case *configapi.Transaction_Rollback:
		return r.reconcileTransactionRollback(ctx, transaction, t.Rollback)
	}
	return false, nil
}

func (r *Reconciler) reconcileTransactionChange(ctx context.Context, transaction *configapi.Transaction, change *configapi.TransactionChange) (bool, error) {
	log.Debugf("Reconciling transaction change %s in %s state", transaction.ID, transaction.Status.State)
	switch transaction.Status.State {
	case configapi.TransactionState_TRANSACTION_PENDING:
		// If the transaction is Pending, begin validation if the prior transaction
		// has already been applied. This simplifies concurrency control in the controller
		// and guarantees transactions are applied to the configurations in sequential order.
		return r.reconcileTransactionPending(ctx, transaction)
		// if the transaction is in the Validating state, Validate the changes in a transaction and
		// change the state to Applying if validation passed otherwise fail the transaction
	case configapi.TransactionState_TRANSACTION_VALIDATING:
		return r.reconcileTransactionChangeValidating(ctx, transaction, change)
		// If the transaction is in the Applying state, update the Configuration for each
		// target and Complete the transaction.
	case configapi.TransactionState_TRANSACTION_APPLYING:
		return r.reconcileTransactionChangeApplying(ctx, transaction, change)
	}
	return false, nil
}

func (r *Reconciler) reconcileTransactionPending(ctx context.Context, transaction *configapi.Transaction) (bool, error) {
	prevTransaction, err := r.transactions.GetByIndex(ctx, transaction.Index-1)
	if err != nil && !errors.IsNotFound(err) {
		return false, err
	}
	if errors.IsNotFound(err) ||
		prevTransaction.Status.State == configapi.TransactionState_TRANSACTION_COMPLETE ||
		prevTransaction.Status.State == configapi.TransactionState_TRANSACTION_FAILED {
		transaction.Status.State = configapi.TransactionState_TRANSACTION_VALIDATING
		err = r.transactions.Update(ctx, transaction)
		if err != nil {
			if !errors.IsNotFound(err) && !errors.IsConflict(err) {
				return false, err
			}
			return false, nil
		}
		return true, nil
	}
	return false, nil
}

func (r *Reconciler) reconcileTransactionChangeValidating(ctx context.Context, transaction *configapi.Transaction, change *configapi.TransactionChange) (bool, error) {
	// Look through the change targets and validate changes for each target
	for targetID, change := range change.Changes {
		modelName := utils.ToModelNameV2(change.TargetType, change.TargetVersion)
		modelPlugin, ok := r.pluginRegistry.GetPlugin(modelName)
		if !ok {
			return false, errors.NewNotFound("model plugin not found")
		}

		pathValues := make([]*configapi.PathValue, 0, len(change.Values))
		for path, changeValue := range change.Values {
			pathValue := &configapi.PathValue{
				Path:    path,
				Value:   changeValue.Value,
				Deleted: changeValue.Delete,
			}
			pathValues = append(pathValues, pathValue)
		}

		jsonTree, err := tree.BuildTree(pathValues, true)
		if err != nil {
			return false, err
		}
		// If validation fails any target, mark the transaction Failed.
		// If validation is successful, proceed to Applying.
		err = modelPlugin.Validate(ctx, jsonTree)
		if err != nil {
			transaction.Status.State = configapi.TransactionState_TRANSACTION_FAILED
			err = r.transactions.Update(ctx, transaction)
			if err != nil {
				if !errors.IsNotFound(err) && !errors.IsConflict(err) {
					return false, err
				}
				return false, nil
			}
			return true, nil
		}

		// Get the target configuration and record the source values in the transaction status
		var configValues map[string]*configapi.PathValue
		configID := configuration.NewID(targetID, change.TargetType, change.TargetVersion)
		if config, err := r.configurations.Get(ctx, configID); err != nil {
			if !errors.IsNotFound(err) {
				return false, err
			}
			configValues = make(map[string]*configapi.PathValue)
		} else if config.Values != nil {
			configValues = config.Values
		} else {
			configValues = make(map[string]*configapi.PathValue)
		}

		source := configapi.Source{
			TargetType:    change.TargetType,
			TargetVersion: change.TargetVersion,
			Values:        make(map[string]configapi.PathValue),
		}
		for path := range change.Values {
			pathValue, ok := configValues[path]
			if ok {
				source.Values[path] = *pathValue
			} else {
				source.Values[path] = configapi.PathValue{
					Path:    path,
					Deleted: true,
				}
			}
		}
		transaction.Status.Sources[targetID] = source
	}

	// Store configuration sources and move the transaction to the APPLYING state
	transaction.Status.State = configapi.TransactionState_TRANSACTION_APPLYING
	err := r.transactions.Update(ctx, transaction)
	if err != nil {
		if !errors.IsNotFound(err) && !errors.IsConflict(err) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (r *Reconciler) reconcileTransactionChangeApplying(ctx context.Context, transaction *configapi.Transaction, change *configapi.TransactionChange) (bool, error) {
	// Once the source configurations have been stored we can update the target configurations
	for targetID, change := range change.Changes {
		configID := configuration.NewID(targetID, change.TargetType, change.TargetVersion)
		config, err := r.configurations.Get(ctx, configID)
		if err != nil {
			if !errors.IsNotFound(err) {
				return false, err
			}

			config = &configapi.Configuration{
				ID:            configID,
				TargetID:      targetID,
				TargetType:    change.TargetType,
				TargetVersion: change.TargetVersion,
				Values:        make(map[string]*configapi.PathValue),
			}
			for path, changeValue := range change.Values {
				config.Values[path] = &configapi.PathValue{
					Value:   changeValue.Value,
					Deleted: changeValue.Delete,
					Index:   transaction.Index,
				}
			}

			config.Status.State = configapi.ConfigurationState_CONFIGURATION_PENDING
			config.Status.TransactionIndex = transaction.Index
			err = r.configurations.Create(ctx, config)
			if err != nil {
				if !errors.IsAlreadyExists(err) {
					return false, err
				}
				return false, nil
			}
		} else {
			if config.Values == nil {
				config.Values = make(map[string]*configapi.PathValue)
			}
			for path, changeValue := range change.Values {
				config.Values[path] = &configapi.PathValue{
					Value:   changeValue.Value,
					Deleted: changeValue.Delete,
					Index:   transaction.Index,
				}
			}

			config.Status.State = configapi.ConfigurationState_CONFIGURATION_PENDING
			config.Status.TransactionIndex = transaction.Index
			err = r.configurations.Update(ctx, config)
			if err != nil {
				if !errors.IsConflict(err) && !errors.IsNotFound(err) {
					return false, err
				}
				return false, nil
			}
		}
	}

	// Complete the transaction once the target configurations have been updated
	transaction.Status.State = configapi.TransactionState_TRANSACTION_COMPLETE
	err := r.transactions.Update(ctx, transaction)
	if err != nil {
		if !errors.IsNotFound(err) && !errors.IsConflict(err) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (r *Reconciler) reconcileTransactionRollback(ctx context.Context, transaction *configapi.Transaction, rollback *configapi.TransactionRollback) (bool, error) {
	log.Debugf("Reconciling transaction rollback %s in %s state", transaction.ID, transaction.Status.State)
	switch transaction.Status.State {
	case configapi.TransactionState_TRANSACTION_PENDING:
		// If the transaction is Pending, begin validation if the prior transaction
		// has already been applied. This simplifies concurrency control in the controller
		// and guarantees transactions are applied to the configurations in sequential order.
		return r.reconcileTransactionPending(ctx, transaction)
		// if the transaction is in the Validating state, Validate the rollback in a transaction and
		// change the state to Applying if validation passed otherwise fail the transaction
	case configapi.TransactionState_TRANSACTION_VALIDATING:
		return r.reconcileTransactionRollbackValidating(ctx, transaction, rollback)
		// If the transaction is in the Applying state, rollback the Configuration for each
		// target and Complete the transaction.
	case configapi.TransactionState_TRANSACTION_APPLYING:
		return r.reconcileTransactionRollbackApplying(ctx, transaction, rollback)
	}
	return false, nil
}

func (r *Reconciler) reconcileTransactionRollbackValidating(ctx context.Context, transaction *configapi.Transaction, rollback *configapi.TransactionRollback) (bool, error) {
	// Get the transaction being rolled back and apply its sources to this transaction
	// The source transaction's sources are stored in the rollback transaction to ensure
	// the rollback can be applied once it's in the APPLYING state even if the source
	// transaction is deleted from the log during compaction.
	targetTransaction, err := r.transactions.GetByIndex(ctx, rollback.Index)
	if err != nil {
		if !errors.IsNotFound(err) {
			return false, err
		}
		transaction.Status.State = configapi.TransactionState_TRANSACTION_FAILED
	} else {
		switch targetTransaction.Transaction.(type) {
		case *configapi.Transaction_Change:
			transaction.Status.State = configapi.TransactionState_TRANSACTION_APPLYING
			transaction.Status.Sources = targetTransaction.Status.Sources
		case *configapi.Transaction_Rollback:
			transaction.Status.State = configapi.TransactionState_TRANSACTION_FAILED
		}
	}

	err = r.transactions.Update(ctx, transaction)
	if err != nil {
		if !errors.IsNotFound(err) && !errors.IsConflict(err) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (r *Reconciler) reconcileTransactionRollbackApplying(ctx context.Context, transaction *configapi.Transaction, rollback *configapi.TransactionRollback) (bool, error) {
	// Once the source configurations have been stored we can update the target configurations
	for targetID, source := range transaction.Status.Sources {
		config, err := r.configurations.Get(ctx, configuration.NewID(targetID, source.TargetType, source.TargetVersion))
		if err != nil {
			if !errors.IsNotFound(err) {
				return false, err
			}
			return false, nil
		}

		// Update the configuration's values with the transaction index
		for path, pathValue := range source.Values {
			if config.Values == nil {
				config.Values = make(map[string]*configapi.PathValue)
			}
			config.Values[path] = &configapi.PathValue{
				Path:    path,
				Value:   pathValue.Value,
				Deleted: pathValue.Deleted,
				Index:   transaction.Index,
			}
		}

		config.Status.State = configapi.ConfigurationState_CONFIGURATION_PENDING
		config.Status.TransactionIndex = transaction.Index
		err = r.configurations.Update(ctx, config)
		if err != nil {
			if !errors.IsConflict(err) && !errors.IsNotFound(err) {
				return false, err
			}
			return false, nil
		}
	}

	// Complete the transaction once the target configurations have been updated
	transaction.Status.State = configapi.TransactionState_TRANSACTION_COMPLETE
	err := r.transactions.Update(ctx, transaction)
	if err != nil {
		if !errors.IsNotFound(err) && !errors.IsConflict(err) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}
