// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: critical_section.proto

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

// Protocol buffer representing a CriticalSection.
type CriticalSectionDef struct {
	// Name of the critical section handle.
	CriticalSectionName string `protobuf:"bytes,1,opt,name=critical_section_name,json=criticalSectionName,proto3" json:"critical_section_name,omitempty"`
}

func (m *CriticalSectionDef) Reset()         { *m = CriticalSectionDef{} }
func (m *CriticalSectionDef) String() string { return proto.CompactTextString(m) }
func (*CriticalSectionDef) ProtoMessage()    {}
func (*CriticalSectionDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_496f0ce04a3aeb97, []int{0}
}
func (m *CriticalSectionDef) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CriticalSectionDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CriticalSectionDef.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CriticalSectionDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriticalSectionDef.Merge(m, src)
}
func (m *CriticalSectionDef) XXX_Size() int {
	return m.Size()
}
func (m *CriticalSectionDef) XXX_DiscardUnknown() {
	xxx_messageInfo_CriticalSectionDef.DiscardUnknown(m)
}

var xxx_messageInfo_CriticalSectionDef proto.InternalMessageInfo

func (m *CriticalSectionDef) GetCriticalSectionName() string {
	if m != nil {
		return m.CriticalSectionName
	}
	return ""
}

// Protocol buffer representing a CriticalSection execution.
type CriticalSectionExecutionDef struct {
	// Name of the critical section handle.
	ExecuteInCriticalSectionName string `protobuf:"bytes,1,opt,name=execute_in_critical_section_name,json=executeInCriticalSectionName,proto3" json:"execute_in_critical_section_name,omitempty"`
	// Whether this operation requires exclusive access to its resources,
	// (i.e., no other CriticalSections may request the same resources).
	ExclusiveResourceAccess bool `protobuf:"varint,2,opt,name=exclusive_resource_access,json=exclusiveResourceAccess,proto3" json:"exclusive_resource_access,omitempty"`
}

func (m *CriticalSectionExecutionDef) Reset()         { *m = CriticalSectionExecutionDef{} }
func (m *CriticalSectionExecutionDef) String() string { return proto.CompactTextString(m) }
func (*CriticalSectionExecutionDef) ProtoMessage()    {}
func (*CriticalSectionExecutionDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_496f0ce04a3aeb97, []int{1}
}
func (m *CriticalSectionExecutionDef) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CriticalSectionExecutionDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CriticalSectionExecutionDef.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CriticalSectionExecutionDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriticalSectionExecutionDef.Merge(m, src)
}
func (m *CriticalSectionExecutionDef) XXX_Size() int {
	return m.Size()
}
func (m *CriticalSectionExecutionDef) XXX_DiscardUnknown() {
	xxx_messageInfo_CriticalSectionExecutionDef.DiscardUnknown(m)
}

var xxx_messageInfo_CriticalSectionExecutionDef proto.InternalMessageInfo

func (m *CriticalSectionExecutionDef) GetExecuteInCriticalSectionName() string {
	if m != nil {
		return m.ExecuteInCriticalSectionName
	}
	return ""
}

func (m *CriticalSectionExecutionDef) GetExclusiveResourceAccess() bool {
	if m != nil {
		return m.ExclusiveResourceAccess
	}
	return false
}

func init() {
	proto.RegisterType((*CriticalSectionDef)(nil), "tensorflow.CriticalSectionDef")
	proto.RegisterType((*CriticalSectionExecutionDef)(nil), "tensorflow.CriticalSectionExecutionDef")
}

func init() { proto.RegisterFile("critical_section.proto", fileDescriptor_496f0ce04a3aeb97) }

var fileDescriptor_496f0ce04a3aeb97 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x63, 0x06, 0x04, 0x1e, 0x83, 0x0a, 0x41, 0x20, 0x2b, 0x74, 0xea, 0x42, 0x2a, 0xc1,
	0xc6, 0x46, 0x0b, 0x08, 0x16, 0x54, 0x85, 0x8d, 0xc5, 0x72, 0x9f, 0x5e, 0x8a, 0x21, 0x89, 0xcb,
	0xb3, 0x43, 0xfb, 0x19, 0xfc, 0x02, 0x7f, 0xc3, 0xd8, 0x91, 0x11, 0x25, 0x3f, 0xc1, 0x88, 0x08,
	0x81, 0x8a, 0x20, 0x75, 0xb4, 0xee, 0xb9, 0xc7, 0x4f, 0x97, 0x6f, 0x03, 0x69, 0xa7, 0x41, 0xa5,
	0xd2, 0x22, 0x38, 0x6d, 0xf2, 0x68, 0x4a, 0xc6, 0x19, 0x9f, 0x3b, 0xcc, 0xad, 0xa1, 0x24, 0x35,
	0xb3, 0xee, 0x25, 0xf7, 0x87, 0x0d, 0x75, 0xf3, 0x0d, 0x9d, 0x61, 0xe2, 0x1f, 0xf1, 0x4e, 0xbb,
	0x2b, 0x73, 0x95, 0x61, 0xc0, 0x42, 0xd6, 0xdb, 0x8c, 0xb7, 0xe0, 0x6f, 0xe5, 0x5a, 0x65, 0xd8,
	0x7d, 0x61, 0x7c, 0xaf, 0xa5, 0x3a, 0x9f, 0x23, 0x14, 0x3f, 0xce, 0x0b, 0x1e, 0x62, 0xfd, 0x46,
	0xa9, 0x73, 0xb9, 0x4a, 0xbf, 0xdf, 0x70, 0x57, 0xf9, 0xf0, 0xff, 0x3f, 0xfe, 0x09, 0xdf, 0xc5,
	0x39, 0xa4, 0x85, 0xd5, 0x4f, 0x28, 0x09, 0xad, 0x29, 0x08, 0x50, 0x2a, 0x00, 0xb4, 0x36, 0x58,
	0x0b, 0x59, 0x6f, 0x23, 0xde, 0xf9, 0x05, 0xe2, 0x26, 0x3f, 0xad, 0xe3, 0xc1, 0xe3, 0x6b, 0x29,
	0xd8, 0xa2, 0x14, 0xec, 0xbd, 0x14, 0xec, 0xb9, 0x12, 0xde, 0xa2, 0x12, 0xde, 0x5b, 0x25, 0x3c,
	0x1e, 0x18, 0x9a, 0x44, 0xcb, 0x5d, 0xa2, 0x84, 0x54, 0x86, 0x33, 0x43, 0x0f, 0x83, 0x4e, 0xeb,
	0x88, 0xd1, 0xd7, 0x86, 0x76, 0xc4, 0x6e, 0x0f, 0x26, 0xda, 0xdd, 0x15, 0xe3, 0x08, 0x4c, 0xd6,
	0x27, 0xa5, 0x0f, 0xa7, 0x64, 0xee, 0x11, 0x5c, 0x7f, 0x69, 0xf9, 0x60, 0x6c, 0xbc, 0x5e, 0x6f,
	0x7e, 0xfc, 0x19, 0x00, 0x00, 0xff, 0xff, 0x83, 0x2f, 0xce, 0x4b, 0x8d, 0x01, 0x00, 0x00,
}

func (m *CriticalSectionDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CriticalSectionDef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.CriticalSectionName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCriticalSection(dAtA, i, uint64(len(m.CriticalSectionName)))
		i += copy(dAtA[i:], m.CriticalSectionName)
	}
	return i, nil
}

func (m *CriticalSectionExecutionDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CriticalSectionExecutionDef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ExecuteInCriticalSectionName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCriticalSection(dAtA, i, uint64(len(m.ExecuteInCriticalSectionName)))
		i += copy(dAtA[i:], m.ExecuteInCriticalSectionName)
	}
	if m.ExclusiveResourceAccess {
		dAtA[i] = 0x10
		i++
		if m.ExclusiveResourceAccess {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintCriticalSection(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CriticalSectionDef) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CriticalSectionName)
	if l > 0 {
		n += 1 + l + sovCriticalSection(uint64(l))
	}
	return n
}

func (m *CriticalSectionExecutionDef) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ExecuteInCriticalSectionName)
	if l > 0 {
		n += 1 + l + sovCriticalSection(uint64(l))
	}
	if m.ExclusiveResourceAccess {
		n += 2
	}
	return n
}

func sovCriticalSection(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCriticalSection(x uint64) (n int) {
	return sovCriticalSection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CriticalSectionDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCriticalSection
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
			return fmt.Errorf("proto: CriticalSectionDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CriticalSectionDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CriticalSectionName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCriticalSection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCriticalSection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCriticalSection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CriticalSectionName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCriticalSection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCriticalSection
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCriticalSection
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
func (m *CriticalSectionExecutionDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCriticalSection
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
			return fmt.Errorf("proto: CriticalSectionExecutionDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CriticalSectionExecutionDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecuteInCriticalSectionName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCriticalSection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCriticalSection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCriticalSection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecuteInCriticalSectionName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExclusiveResourceAccess", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCriticalSection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ExclusiveResourceAccess = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCriticalSection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCriticalSection
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCriticalSection
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
func skipCriticalSection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCriticalSection
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
					return 0, ErrIntOverflowCriticalSection
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
					return 0, ErrIntOverflowCriticalSection
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
				return 0, ErrInvalidLengthCriticalSection
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCriticalSection
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCriticalSection
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
				next, err := skipCriticalSection(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCriticalSection
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
	ErrInvalidLengthCriticalSection = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCriticalSection   = fmt.Errorf("proto: integer overflow")
)