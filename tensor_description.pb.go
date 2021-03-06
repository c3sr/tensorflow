// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensor_description.proto

package tensorflow

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TensorDescription struct {
	// Data type of tensor elements
	Dtype DataType `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	// Shape of the tensor.
	Shape *TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	// Information about the size and allocator used for the data
	AllocationDescription *AllocationDescription `protobuf:"bytes,4,opt,name=allocation_description,json=allocationDescription,proto3" json:"allocation_description,omitempty"`
}

func (m *TensorDescription) Reset()         { *m = TensorDescription{} }
func (m *TensorDescription) String() string { return proto.CompactTextString(m) }
func (*TensorDescription) ProtoMessage()    {}
func (*TensorDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e7257165d0caf56, []int{0}
}
func (m *TensorDescription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TensorDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TensorDescription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TensorDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TensorDescription.Merge(m, src)
}
func (m *TensorDescription) XXX_Size() int {
	return m.Size()
}
func (m *TensorDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_TensorDescription.DiscardUnknown(m)
}

var xxx_messageInfo_TensorDescription proto.InternalMessageInfo

func (m *TensorDescription) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func (m *TensorDescription) GetShape() *TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *TensorDescription) GetAllocationDescription() *AllocationDescription {
	if m != nil {
		return m.AllocationDescription
	}
	return nil
}

func init() {
	proto.RegisterType((*TensorDescription)(nil), "tensorflow.TensorDescription")
}

func init() { proto.RegisterFile("tensor_description.proto", fileDescriptor_3e7257165d0caf56) }

var fileDescriptor_3e7257165d0caf56 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x49, 0xcd, 0x2b,
	0xce, 0x2f, 0x8a, 0x4f, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0xc8, 0xa4, 0xe5, 0xe4, 0x97, 0x4b, 0x71, 0x97, 0x54,
	0x16, 0xa4, 0x16, 0x43, 0x24, 0xa4, 0x84, 0xa0, 0x5a, 0x8a, 0x33, 0x12, 0x0b, 0x52, 0xa1, 0x62,
	0x32, 0x89, 0x39, 0x39, 0xf9, 0xc9, 0x89, 0x20, 0xed, 0x98, 0x46, 0x29, 0x9d, 0x65, 0xe4, 0x12,
	0x0c, 0x01, 0x6b, 0x72, 0x41, 0xc8, 0x09, 0x69, 0x71, 0xb1, 0xa6, 0x80, 0xcc, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x33, 0x12, 0xd1, 0x43, 0x58, 0xa8, 0xe7, 0x92, 0x58, 0x92, 0x18, 0x52, 0x59,
	0x90, 0x1a, 0x04, 0x51, 0x22, 0x64, 0xc4, 0xc5, 0x0a, 0xb6, 0x4e, 0x82, 0x49, 0x81, 0x51, 0x83,
	0xdb, 0x48, 0x06, 0x59, 0x2d, 0xc4, 0xe4, 0x60, 0x90, 0x74, 0x00, 0xc8, 0xba, 0x20, 0x88, 0x52,
	0xa1, 0x08, 0x2e, 0x31, 0xec, 0xae, 0x92, 0x60, 0x01, 0x1b, 0xa2, 0x88, 0x6c, 0x88, 0x23, 0x5c,
	0x25, 0x92, 0x13, 0x83, 0x44, 0x13, 0xb1, 0x09, 0x3b, 0x15, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70,
	0xe3, 0xb1, 0x1c, 0x03, 0x97, 0x44, 0x7e, 0x51, 0x3a, 0xb2, 0xb1, 0x69, 0x45, 0x89, 0xb9, 0xa9,
	0xe5, 0xf9, 0x45, 0xd9, 0x4e, 0xe2, 0x18, 0x01, 0x00, 0x76, 0x6c, 0x71, 0x00, 0x63, 0x94, 0x62,
	0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x51, 0x62, 0xa6, 0x6e, 0x41,
	0x51, 0x7e, 0x56, 0x6a, 0x72, 0x89, 0x3e, 0xc2, 0x9c, 0x1f, 0x8c, 0x8c, 0x49, 0x6c, 0xe0, 0xb0,
	0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xca, 0x7e, 0x01, 0xb2, 0x01, 0x00, 0x00,
}

func (m *TensorDescription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TensorDescription) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Dtype != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTensorDescription(dAtA, i, uint64(m.Dtype))
	}
	if m.Shape != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTensorDescription(dAtA, i, uint64(m.Shape.Size()))
		n1, err1 := m.Shape.MarshalTo(dAtA[i:])
		if err1 != nil {
			return 0, err1
		}
		i += n1
	}
	if m.AllocationDescription != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTensorDescription(dAtA, i, uint64(m.AllocationDescription.Size()))
		n2, err2 := m.AllocationDescription.MarshalTo(dAtA[i:])
		if err2 != nil {
			return 0, err2
		}
		i += n2
	}
	return i, nil
}

func encodeVarintTensorDescription(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TensorDescription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Dtype != 0 {
		n += 1 + sovTensorDescription(uint64(m.Dtype))
	}
	if m.Shape != nil {
		l = m.Shape.Size()
		n += 1 + l + sovTensorDescription(uint64(l))
	}
	if m.AllocationDescription != nil {
		l = m.AllocationDescription.Size()
		n += 1 + l + sovTensorDescription(uint64(l))
	}
	return n
}

func sovTensorDescription(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTensorDescription(x uint64) (n int) {
	return sovTensorDescription(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TensorDescription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTensorDescription
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
			return fmt.Errorf("proto: TensorDescription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TensorDescription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dtype", wireType)
			}
			m.Dtype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorDescription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Dtype |= DataType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shape", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorDescription
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
				return ErrInvalidLengthTensorDescription
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTensorDescription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Shape == nil {
				m.Shape = &TensorShapeProto{}
			}
			if err := m.Shape.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocationDescription", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorDescription
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
				return ErrInvalidLengthTensorDescription
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTensorDescription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AllocationDescription == nil {
				m.AllocationDescription = &AllocationDescription{}
			}
			if err := m.AllocationDescription.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTensorDescription(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTensorDescription
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTensorDescription
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
func skipTensorDescription(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTensorDescription
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
					return 0, ErrIntOverflowTensorDescription
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
					return 0, ErrIntOverflowTensorDescription
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
				return 0, ErrInvalidLengthTensorDescription
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTensorDescription
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTensorDescription
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
				next, err := skipTensorDescription(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTensorDescription
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
	ErrInvalidLengthTensorDescription = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTensorDescription   = fmt.Errorf("proto: integer overflow")
)
