// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: resource_handle.proto

package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type ResourceHandle struct {
	// Unique name for the device containing the resource.
	Device string `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	// Container in which this resource is placed.
	Container string `protobuf:"bytes,2,opt,name=container,proto3" json:"container,omitempty"`
	// Unique name of this resource.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Hash code for the type of the resource. Is only valid in the same device
	// and in the same execution.
	HashCode uint64 `protobuf:"varint,4,opt,name=hash_code,json=hashCode,proto3" json:"hash_code,omitempty"`
	// For debug-only, the name of the type pointed to by this handle, if
	// available.
	MaybeTypeName string `protobuf:"bytes,5,opt,name=maybe_type_name,json=maybeTypeName,proto3" json:"maybe_type_name,omitempty"`
}

func (m *ResourceHandle) Reset()         { *m = ResourceHandle{} }
func (m *ResourceHandle) String() string { return proto.CompactTextString(m) }
func (*ResourceHandle) ProtoMessage()    {}
func (*ResourceHandle) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_handle_fd663e99a69cddeb, []int{0}
}
func (m *ResourceHandle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceHandle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceHandle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ResourceHandle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceHandle.Merge(dst, src)
}
func (m *ResourceHandle) XXX_Size() int {
	return m.Size()
}
func (m *ResourceHandle) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceHandle.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceHandle proto.InternalMessageInfo

func (m *ResourceHandle) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *ResourceHandle) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *ResourceHandle) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceHandle) GetHashCode() uint64 {
	if m != nil {
		return m.HashCode
	}
	return 0
}

func (m *ResourceHandle) GetMaybeTypeName() string {
	if m != nil {
		return m.MaybeTypeName
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceHandle)(nil), "tensorflow.ResourceHandle")
}
func (m *ResourceHandle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceHandle) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Device) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(len(m.Device)))
		i += copy(dAtA[i:], m.Device)
	}
	if len(m.Container) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(len(m.Container)))
		i += copy(dAtA[i:], m.Container)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.HashCode != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(m.HashCode))
	}
	if len(m.MaybeTypeName) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(len(m.MaybeTypeName)))
		i += copy(dAtA[i:], m.MaybeTypeName)
	}
	return i, nil
}

func encodeVarintResourceHandle(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ResourceHandle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Device)
	if l > 0 {
		n += 1 + l + sovResourceHandle(uint64(l))
	}
	l = len(m.Container)
	if l > 0 {
		n += 1 + l + sovResourceHandle(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovResourceHandle(uint64(l))
	}
	if m.HashCode != 0 {
		n += 1 + sovResourceHandle(uint64(m.HashCode))
	}
	l = len(m.MaybeTypeName)
	if l > 0 {
		n += 1 + l + sovResourceHandle(uint64(l))
	}
	return n
}

func sovResourceHandle(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozResourceHandle(x uint64) (n int) {
	return sovResourceHandle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResourceHandle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourceHandle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResourceHandle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceHandle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Device", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResourceHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Device = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Container", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResourceHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Container = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResourceHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashCode", wireType)
			}
			m.HashCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HashCode |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaybeTypeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResourceHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaybeTypeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResourceHandle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResourceHandle
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
func skipResourceHandle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResourceHandle
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
					return 0, ErrIntOverflowResourceHandle
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
					return 0, ErrIntOverflowResourceHandle
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthResourceHandle
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowResourceHandle
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
				next, err := skipResourceHandle(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthResourceHandle = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResourceHandle   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("resource_handle.proto", fileDescriptor_resource_handle_fd663e99a69cddeb)
}

var fileDescriptor_resource_handle_fd663e99a69cddeb = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4a, 0x2d, 0xce,
	0x2f, 0x2d, 0x4a, 0x4e, 0x8d, 0xcf, 0x48, 0xcc, 0x4b, 0xc9, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x2a, 0x49, 0xcd, 0x2b, 0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0x57, 0x9a, 0xcf,
	0xc8, 0xc5, 0x17, 0x04, 0x55, 0xe5, 0x01, 0x56, 0x24, 0x24, 0xc6, 0xc5, 0x96, 0x92, 0x5a, 0x96,
	0x99, 0x9c, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe5, 0x09, 0xc9, 0x70, 0x71, 0x26,
	0xe7, 0xe7, 0x95, 0x24, 0x66, 0xe6, 0xa5, 0x16, 0x49, 0x30, 0x81, 0xa5, 0x10, 0x02, 0x42, 0x42,
	0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0xcc, 0x60, 0x09, 0x30, 0x5b, 0x48, 0x9a, 0x8b, 0x33,
	0x23, 0xb1, 0x38, 0x23, 0x3e, 0x39, 0x3f, 0x25, 0x55, 0x82, 0x45, 0x81, 0x51, 0x83, 0x25, 0x88,
	0x03, 0x24, 0xe0, 0x9c, 0x9f, 0x92, 0x2a, 0xa4, 0xc6, 0xc5, 0x9f, 0x9b, 0x58, 0x99, 0x94, 0x1a,
	0x5f, 0x52, 0x59, 0x90, 0x1a, 0x0f, 0xd6, 0xcb, 0x0a, 0xd6, 0xcb, 0x0b, 0x16, 0x0e, 0xa9, 0x2c,
	0x48, 0xf5, 0x4b, 0xcc, 0x4d, 0x75, 0xf2, 0x39, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x06, 0x2e, 0x89, 0xfc, 0xa2, 0x74, 0x3d, 0x84, 0x5f, 0xf4, 0xd2, 0x8a, 0x12, 0x73, 0x53, 0xcb,
	0xf3, 0x8b, 0xb2, 0x9d, 0x84, 0x51, 0xbd, 0x14, 0x00, 0xf2, 0x76, 0x00, 0xe3, 0x0f, 0x46, 0xc6,
	0x24, 0x36, 0x70, 0x10, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xf3, 0x9d, 0x9b, 0x1b,
	0x01, 0x00, 0x00,
}
