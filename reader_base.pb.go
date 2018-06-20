// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reader_base.proto

package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// For serializing and restoring the state of ReaderBase, see
// reader_base.h for details.
type ReaderBaseState struct {
	WorkStarted        int64  `protobuf:"varint,1,opt,name=work_started,json=workStarted,proto3" json:"work_started,omitempty"`
	WorkFinished       int64  `protobuf:"varint,2,opt,name=work_finished,json=workFinished,proto3" json:"work_finished,omitempty"`
	NumRecordsProduced int64  `protobuf:"varint,3,opt,name=num_records_produced,json=numRecordsProduced,proto3" json:"num_records_produced,omitempty"`
	CurrentWork        []byte `protobuf:"bytes,4,opt,name=current_work,json=currentWork,proto3" json:"current_work,omitempty"`
}

func (m *ReaderBaseState) Reset()                    { *m = ReaderBaseState{} }
func (m *ReaderBaseState) String() string            { return proto.CompactTextString(m) }
func (*ReaderBaseState) ProtoMessage()               {}
func (*ReaderBaseState) Descriptor() ([]byte, []int) { return fileDescriptorReaderBase, []int{0} }

func (m *ReaderBaseState) GetWorkStarted() int64 {
	if m != nil {
		return m.WorkStarted
	}
	return 0
}

func (m *ReaderBaseState) GetWorkFinished() int64 {
	if m != nil {
		return m.WorkFinished
	}
	return 0
}

func (m *ReaderBaseState) GetNumRecordsProduced() int64 {
	if m != nil {
		return m.NumRecordsProduced
	}
	return 0
}

func (m *ReaderBaseState) GetCurrentWork() []byte {
	if m != nil {
		return m.CurrentWork
	}
	return nil
}

func init() {
	proto.RegisterType((*ReaderBaseState)(nil), "tensorflow.ReaderBaseState")
}
func (m *ReaderBaseState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReaderBaseState) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.WorkStarted != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintReaderBase(dAtA, i, uint64(m.WorkStarted))
	}
	if m.WorkFinished != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintReaderBase(dAtA, i, uint64(m.WorkFinished))
	}
	if m.NumRecordsProduced != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintReaderBase(dAtA, i, uint64(m.NumRecordsProduced))
	}
	if len(m.CurrentWork) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintReaderBase(dAtA, i, uint64(len(m.CurrentWork)))
		i += copy(dAtA[i:], m.CurrentWork)
	}
	return i, nil
}

func encodeVarintReaderBase(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ReaderBaseState) Size() (n int) {
	var l int
	_ = l
	if m.WorkStarted != 0 {
		n += 1 + sovReaderBase(uint64(m.WorkStarted))
	}
	if m.WorkFinished != 0 {
		n += 1 + sovReaderBase(uint64(m.WorkFinished))
	}
	if m.NumRecordsProduced != 0 {
		n += 1 + sovReaderBase(uint64(m.NumRecordsProduced))
	}
	l = len(m.CurrentWork)
	if l > 0 {
		n += 1 + l + sovReaderBase(uint64(l))
	}
	return n
}

func sovReaderBase(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozReaderBase(x uint64) (n int) {
	return sovReaderBase(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReaderBaseState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReaderBase
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
			return fmt.Errorf("proto: ReaderBaseState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReaderBaseState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkStarted", wireType)
			}
			m.WorkStarted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReaderBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WorkStarted |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkFinished", wireType)
			}
			m.WorkFinished = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReaderBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WorkFinished |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumRecordsProduced", wireType)
			}
			m.NumRecordsProduced = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReaderBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumRecordsProduced |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentWork", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReaderBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthReaderBase
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentWork = append(m.CurrentWork[:0], dAtA[iNdEx:postIndex]...)
			if m.CurrentWork == nil {
				m.CurrentWork = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReaderBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReaderBase
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
func skipReaderBase(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReaderBase
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
					return 0, ErrIntOverflowReaderBase
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
					return 0, ErrIntOverflowReaderBase
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
				return 0, ErrInvalidLengthReaderBase
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowReaderBase
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
				next, err := skipReaderBase(dAtA[start:])
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
	ErrInvalidLengthReaderBase = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReaderBase   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("reader_base.proto", fileDescriptorReaderBase) }

var fileDescriptorReaderBase = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x79, 0x14, 0x31, 0xb8, 0x45, 0x80, 0xc5, 0xe0, 0x29, 0x2a, 0xb0, 0x74, 0x8a, 0x90,
	0x38, 0x01, 0x19, 0x98, 0xa3, 0x74, 0x60, 0xb4, 0xdc, 0xf8, 0x05, 0xaa, 0x12, 0xbf, 0xea, 0xd9,
	0x56, 0xaf, 0xc1, 0x49, 0x38, 0x07, 0x23, 0x47, 0x40, 0xe1, 0x12, 0x8c, 0xc8, 0x6e, 0xa4, 0xac,
	0xdf, 0xff, 0xfd, 0xb6, 0xfe, 0x27, 0xae, 0x19, 0x8d, 0x45, 0xd6, 0x1b, 0xe3, 0xb1, 0xdc, 0x33,
	0x05, 0x92, 0x22, 0xa0, 0xf3, 0xc4, 0xdd, 0x3b, 0x1d, 0xee, 0x3e, 0x41, 0x5c, 0x36, 0xd9, 0xa8,
	0x8c, 0xc7, 0x75, 0x30, 0x01, 0xe5, 0xad, 0x58, 0x1c, 0x88, 0x77, 0xda, 0x07, 0xc3, 0x01, 0xad,
	0x82, 0x25, 0xac, 0x66, 0xcd, 0x3c, 0xb1, 0xf5, 0x11, 0xc9, 0x7b, 0x71, 0x91, 0x95, 0x6e, 0xeb,
	0xb6, 0xfe, 0x0d, 0xad, 0x3a, 0xcd, 0x4e, 0xee, 0x3d, 0x8f, 0x4c, 0x3e, 0x88, 0x1b, 0x17, 0x7b,
	0xcd, 0xd8, 0x12, 0x5b, 0xaf, 0xf7, 0x4c, 0x36, 0xb6, 0x68, 0xd5, 0x2c, 0xbb, 0xd2, 0xc5, 0xbe,
	0x39, 0x46, 0xf5, 0x98, 0xa4, 0x9f, 0xdb, 0xc8, 0x8c, 0x2e, 0xe8, 0xf4, 0x92, 0x3a, 0x5b, 0xc2,
	0x6a, 0xd1, 0xcc, 0x47, 0xf6, 0x42, 0xbc, 0xab, 0x9e, 0xbe, 0x86, 0x02, 0xbe, 0x87, 0x02, 0x7e,
	0x86, 0x02, 0x3e, 0x7e, 0x8b, 0x13, 0xa1, 0x88, 0x5f, 0xcb, 0x69, 0x52, 0xd9, 0xb1, 0xe9, 0x31,
	0xd5, 0xab, 0xab, 0x69, 0x59, 0x9d, 0x96, 0xfb, 0x1a, 0xfe, 0x00, 0x36, 0xe7, 0xf9, 0x0c, 0x8f,
	0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x20, 0xae, 0xfb, 0x11, 0x1b, 0x01, 0x00, 0x00,
}
