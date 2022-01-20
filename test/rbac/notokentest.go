// Copyright 2019-present Open Networking Foundation.
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

package rbac

import (
	"context"
	gpb "github.com/openconfig/gnmi/proto/gnmi"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/onosproject/onos-config/test/utils/gnmi"
	"github.com/onosproject/onos-config/test/utils/proto"
)

// TestNoToken tests access to a protected API with no access token supplied
func (s *TestSuite) TestNoToken(t *testing.T) {
	const (
		tzValue = "Europe/Dublin"
		tzPath  = "/system/clock/config/timezone-name"
	)
	// Create a simulated device
	simulator := gnmi.CreateSimulator(t)
	defer gnmi.DeleteSimulator(t, simulator)

	// Make a GNMI client to use for requests
	gnmiClient := gnmi.GetGNMIClientOrFail(t)

	// Try to fetch a value from the GNMI client
	devicePath := gnmi.GetDevicePathWithValue(simulator.Name(), tzPath, tzValue, proto.StringVal)
	_, _, err := gnmi.GetGNMIValue(context.Background(), gnmiClient, devicePath, gpb.Encoding_PROTO)

	// An error indicating an unauthenticated request is expected
	assert.Error(t, err)
	if err != nil {
		assert.Contains(t, err.Error(), "Request unauthenticated with bearer")
	}
}