// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: named_tensor.proto

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

// A pair of tensor name and tensor values.
type NamedTensorProto struct {
	// Name of the tensor.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The client can populate a TensorProto using a tensorflow::Tensor`, or
	// directly using the protobuf field accessors.
	//
	// The client specifies whether the returned tensor values should be
	// filled tensor fields (float_val, int_val, etc.) or encoded in a
	// compact form in tensor.tensor_content.
	Tensor *TensorProto `protobuf:"bytes,2,opt,name=tensor" json:"tensor,omitempty"`
}

func (m *NamedTensorProto) Reset()         { *m = NamedTensorProto{} }
func (m *NamedTensorProto) String() string { return proto.CompactTextString(m) }
func (*NamedTensorProto) ProtoMessage()    {}
func (*NamedTensorProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_named_tensor_89ffee8af9c4eff0, []int{0}
}
func (m *NamedTensorProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NamedTensorProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NamedTensorProto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *NamedTensorProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamedTensorProto.Merge(dst, src)
}
func (m *NamedTensorProto) XXX_Size() int {
	return m.Size()
}
func (m *NamedTensorProto) XXX_DiscardUnknown() {
	xxx_messageInfo_NamedTensorProto.DiscardUnknown(m)
}

var xxx_messageInfo_NamedTensorProto proto.InternalMessageInfo

func (m *NamedTensorProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamedTensorProto) GetTensor() *TensorProto {
	if m != nil {
		return m.Tensor
	}
	return nil
}

func init() {
	proto.RegisterType((*NamedTensorProto)(nil), "tensorflow.NamedTensorProto")
}
func (m *NamedTensorProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamedTensorProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNamedTensor(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Tensor != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNamedTensor(dAtA, i, uint64(m.Tensor.Size()))
		n1, err := m.Tensor.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintNamedTensor(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *NamedTensorProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNamedTensor(uint64(l))
	}
	if m.Tensor != nil {
		l = m.Tensor.Size()
		n += 1 + l + sovNamedTensor(uint64(l))
	}
	return n
}

func sovNamedTensor(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNamedTensor(x uint64) (n int) {
	return sovNamedTensor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NamedTensorProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamedTensor
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
			return fmt.Errorf("proto: NamedTensorProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamedTensorProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamedTensor
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
				return ErrInvalidLengthNamedTensor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tensor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamedTensor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNamedTensor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tensor == nil {
				m.Tensor = &TensorProto{}
			}
			if err := m.Tensor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamedTensor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamedTensor
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
func skipNamedTensor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNamedTensor
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
					return 0, ErrIntOverflowNamedTensor
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
					return 0, ErrIntOverflowNamedTensor
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
				return 0, ErrInvalidLengthNamedTensor
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNamedTensor
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
				next, err := skipNamedTensor(dAtA[start:])
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
	ErrInvalidLengthNamedTensor = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNamedTensor   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("named_tensor.proto", fileDescriptor_named_tensor_89ffee8af9c4eff0) }

var fileDescriptor_named_tensor_89ffee8af9c4eff0 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0x4b, 0xcc, 0x4d,
	0x4d, 0x89, 0x2f, 0x49, 0xcd, 0x2b, 0xce, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x82, 0xf0, 0xd2, 0x72, 0xf2, 0xcb, 0xa5, 0x78, 0x90, 0x65, 0x94, 0xc2, 0xb9, 0x04, 0xfc, 0x40,
	0xea, 0x43, 0xc0, 0x82, 0x01, 0x60, 0xd5, 0x42, 0x5c, 0x2c, 0x20, 0x33, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x7d, 0x2e, 0x36, 0x88, 0x3e, 0x09, 0x26, 0x05, 0x46, 0x0d,
	0x6e, 0x23, 0x71, 0x3d, 0x84, 0x91, 0x7a, 0x48, 0x9a, 0x83, 0xa0, 0xca, 0x9c, 0xbc, 0x4e, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x81, 0x4b, 0x22, 0xbf, 0x28, 0x1d, 0x59, 0x77, 0x5a,
	0x51, 0x62, 0x6e, 0x6a, 0x79, 0x7e, 0x51, 0xb6, 0x93, 0x20, 0xba, 0x53, 0x8a, 0x03, 0x18, 0x7f,
	0x30, 0x32, 0x26, 0xb1, 0x81, 0xdd, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x48, 0x35, 0xf7,
	0x4f, 0xdb, 0x00, 0x00, 0x00,
}
