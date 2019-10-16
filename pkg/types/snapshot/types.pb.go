// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/types/snapshot/types.proto

package snapshot

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Phase is a snapshot phase
type Phase int32

const (
	// MARK is the first phase in which changes are marked for deletion
	Phase_MARK Phase = 0
	// DELETE is the second phase in which changes are deleted from stores
	Phase_DELETE Phase = 1
)

var Phase_name = map[int32]string{
	0: "MARK",
	1: "DELETE",
}

var Phase_value = map[string]int32{
	"MARK":   0,
	"DELETE": 1,
}

func (x Phase) String() string {
	return proto.EnumName(Phase_name, int32(x))
}

func (Phase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c57a42bda5c26e41, []int{0}
}

// State is the state of a snapshot within a phase
type State int32

const (
	// PENDING indicates the snapshot is pending
	State_PENDING State = 0
	// RUNNING indicates the snapshot is in progress
	State_RUNNING State = 1
	// COMPLETE indicates the snapshot is complete
	State_COMPLETE State = 2
)

var State_name = map[int32]string{
	0: "PENDING",
	1: "RUNNING",
	2: "COMPLETE",
}

var State_value = map[string]int32{
	"PENDING":  0,
	"RUNNING":  1,
	"COMPLETE": 2,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c57a42bda5c26e41, []int{1}
}

// Status is the status of a snapshot
type Status struct {
	// 'phase' is the snapshot phase
	Phase Phase `protobuf:"varint,1,opt,name=phase,proto3,enum=onos.config.snapshot.Phase" json:"phase,omitempty"`
	// 'state' is the state of a snapshot
	State State `protobuf:"varint,2,opt,name=state,proto3,enum=onos.config.snapshot.State" json:"state,omitempty"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_c57a42bda5c26e41, []int{0}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Status.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return m.Size()
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetPhase() Phase {
	if m != nil {
		return m.Phase
	}
	return Phase_MARK
}

func (m *Status) GetState() State {
	if m != nil {
		return m.State
	}
	return State_PENDING
}

// RetentionOptions specifies the retention policy for a change log
type RetentionOptions struct {
	// 'retain_window' is the duration for which to retain network changes
	RetainWindow *time.Duration `protobuf:"bytes,1,opt,name=retain_window,json=retainWindow,proto3,stdduration" json:"retain_window,omitempty"`
}

func (m *RetentionOptions) Reset()         { *m = RetentionOptions{} }
func (m *RetentionOptions) String() string { return proto.CompactTextString(m) }
func (*RetentionOptions) ProtoMessage()    {}
func (*RetentionOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c57a42bda5c26e41, []int{1}
}
func (m *RetentionOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RetentionOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RetentionOptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RetentionOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetentionOptions.Merge(m, src)
}
func (m *RetentionOptions) XXX_Size() int {
	return m.Size()
}
func (m *RetentionOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_RetentionOptions.DiscardUnknown(m)
}

var xxx_messageInfo_RetentionOptions proto.InternalMessageInfo

func (m *RetentionOptions) GetRetainWindow() *time.Duration {
	if m != nil {
		return m.RetainWindow
	}
	return nil
}

func init() {
	proto.RegisterEnum("onos.config.snapshot.Phase", Phase_name, Phase_value)
	proto.RegisterEnum("onos.config.snapshot.State", State_name, State_value)
	proto.RegisterType((*Status)(nil), "onos.config.snapshot.Status")
	proto.RegisterType((*RetentionOptions)(nil), "onos.config.snapshot.RetentionOptions")
}

func init() { proto.RegisterFile("pkg/types/snapshot/types.proto", fileDescriptor_c57a42bda5c26e41) }

var fileDescriptor_c57a42bda5c26e41 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xc1, 0x4e, 0x2a, 0x31,
	0x14, 0x86, 0xa7, 0x04, 0xb8, 0xe4, 0xc0, 0x35, 0x93, 0x86, 0x05, 0x62, 0xac, 0x86, 0x95, 0x61,
	0xd1, 0x89, 0xf8, 0x04, 0xe2, 0x4c, 0x8c, 0x51, 0x06, 0x32, 0x6a, 0x74, 0x67, 0x06, 0x29, 0xc3,
	0x44, 0xd3, 0x33, 0xa1, 0x25, 0xc4, 0xb7, 0x70, 0xe9, 0x23, 0xb9, 0x64, 0xe9, 0x4e, 0x03, 0x2f,
	0x62, 0xda, 0xca, 0xce, 0xb8, 0x69, 0xce, 0x69, 0xfe, 0xaf, 0xfd, 0x7e, 0x60, 0xc5, 0x53, 0x16,
	0xe8, 0x97, 0x42, 0xa8, 0x40, 0xc9, 0xb4, 0x50, 0x33, 0xd4, 0x6e, 0xe5, 0xc5, 0x1c, 0x35, 0xd2,
	0x26, 0x4a, 0x54, 0xfc, 0x11, 0xe5, 0x34, 0xcf, 0xf8, 0x36, 0xd1, 0x66, 0x19, 0x62, 0xf6, 0x2c,
	0x02, 0x9b, 0x19, 0x2f, 0xa6, 0xc1, 0x64, 0x31, 0x4f, 0x75, 0x8e, 0xd2, 0x51, 0xed, 0x66, 0x86,
	0x19, 0xda, 0x31, 0x30, 0x93, 0xbb, 0xed, 0x48, 0xa8, 0x5e, 0xeb, 0x54, 0x2f, 0x14, 0x3d, 0x86,
	0x4a, 0x31, 0x4b, 0x95, 0x68, 0x91, 0x43, 0x72, 0xb4, 0xd3, 0xdb, 0xe3, 0xbf, 0xfd, 0xc2, 0x47,
	0x26, 0x92, 0xb8, 0xa4, 0x41, 0x94, 0x4e, 0xb5, 0x68, 0x95, 0xfe, 0x42, 0xcc, 0xfb, 0x22, 0x71,
	0xc9, 0xce, 0x3d, 0xf8, 0x89, 0xd0, 0x42, 0x1a, 0xb1, 0x61, 0x61, 0x4e, 0x45, 0x43, 0xf8, 0x3f,
	0x17, 0x3a, 0xcd, 0xe5, 0xc3, 0x32, 0x97, 0x13, 0x5c, 0x5a, 0x83, 0x7a, 0x6f, 0x97, 0xbb, 0x46,
	0x7c, 0xdb, 0x88, 0x87, 0x3f, 0x8d, 0xfa, 0xe5, 0xb7, 0xcf, 0x03, 0x92, 0x34, 0x1c, 0x75, 0x67,
	0xa1, 0xee, 0x3e, 0x54, 0xac, 0x1c, 0xad, 0x41, 0x79, 0x70, 0x9a, 0x5c, 0xfa, 0x1e, 0x05, 0xa8,
	0x86, 0xd1, 0x55, 0x74, 0x13, 0xf9, 0xa4, 0x1b, 0x40, 0xc5, 0x8a, 0xd0, 0x3a, 0xfc, 0x1b, 0x45,
	0x71, 0x78, 0x11, 0x9f, 0xfb, 0x9e, 0x59, 0x92, 0xdb, 0x38, 0x36, 0x0b, 0xa1, 0x0d, 0xa8, 0x9d,
	0x0d, 0x07, 0x23, 0x0b, 0x94, 0xfa, 0xad, 0xf7, 0x35, 0x23, 0xab, 0x35, 0x23, 0x5f, 0x6b, 0x46,
	0x5e, 0x37, 0xcc, 0x5b, 0x6d, 0x98, 0xf7, 0xb1, 0x61, 0xde, 0xb8, 0x6a, 0x85, 0x4e, 0xbe, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x8a, 0x3d, 0x34, 0xc8, 0xa8, 0x01, 0x00, 0x00,
}

func (m *Status) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Status) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Status) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.State != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	if m.Phase != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Phase))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RetentionOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RetentionOptions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RetentionOptions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RetainWindow != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.RetainWindow, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(*m.RetainWindow):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintTypes(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Status) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Phase != 0 {
		n += 1 + sovTypes(uint64(m.Phase))
	}
	if m.State != 0 {
		n += 1 + sovTypes(uint64(m.State))
	}
	return n
}

func (m *RetentionOptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RetainWindow != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.RetainWindow)
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Status) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Status: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Status: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phase", wireType)
			}
			m.Phase = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Phase |= Phase(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RetentionOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RetentionOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RetentionOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetainWindow", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RetainWindow == nil {
				m.RetainWindow = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.RetainWindow, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTypes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTypes(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTypes
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTypes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes   = fmt.Errorf("proto: integer overflow")
)