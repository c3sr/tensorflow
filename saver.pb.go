// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: saver.proto

package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A version number that identifies a different on-disk checkpoint format.
// Usually, each subclass of BaseSaverBuilder works with a particular
// version/format.  However, it is possible that the same builder may be
// upgraded to support a newer checkpoint format in the future.
type SaverDef_CheckpointFormatVersion int32

const (
	// Internal legacy format.
	SaverDef_LEGACY SaverDef_CheckpointFormatVersion = 0
	// Current format: tf.Saver() which works with tensorflow::table::Table.
	SaverDef_V1 SaverDef_CheckpointFormatVersion = 1
	// Experimental format under development.
	SaverDef_V2 SaverDef_CheckpointFormatVersion = 2
)

var SaverDef_CheckpointFormatVersion_name = map[int32]string{
	0: "LEGACY",
	1: "V1",
	2: "V2",
}
var SaverDef_CheckpointFormatVersion_value = map[string]int32{
	"LEGACY": 0,
	"V1":     1,
	"V2":     2,
}

func (x SaverDef_CheckpointFormatVersion) String() string {
	return proto.EnumName(SaverDef_CheckpointFormatVersion_name, int32(x))
}
func (SaverDef_CheckpointFormatVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorSaver, []int{0, 0}
}

// Protocol buffer representing the configuration of a Saver.
type SaverDef struct {
	// The name of the tensor in which to specify the filename when saving or
	// restoring a model checkpoint.
	FilenameTensorName string `protobuf:"bytes,1,opt,name=filename_tensor_name,json=filenameTensorName,proto3" json:"filename_tensor_name,omitempty"`
	// The operation to run when saving a model checkpoint.
	SaveTensorName string `protobuf:"bytes,2,opt,name=save_tensor_name,json=saveTensorName,proto3" json:"save_tensor_name,omitempty"`
	// The operation to run when restoring a model checkpoint.
	RestoreOpName string `protobuf:"bytes,3,opt,name=restore_op_name,json=restoreOpName,proto3" json:"restore_op_name,omitempty"`
	// Maximum number of checkpoints to keep.  If 0, no checkpoints are deleted.
	MaxToKeep int32 `protobuf:"varint,4,opt,name=max_to_keep,json=maxToKeep,proto3" json:"max_to_keep,omitempty"`
	// Shard the save files, one per device that has Variable nodes.
	Sharded bool `protobuf:"varint,5,opt,name=sharded,proto3" json:"sharded,omitempty"`
	// How often to keep an additional checkpoint. If not specified, only the last
	// "max_to_keep" checkpoints are kept; if specified, in addition to keeping
	// the last "max_to_keep" checkpoints, an additional checkpoint will be kept
	// for every n hours of training.
	KeepCheckpointEveryNHours float32                          `protobuf:"fixed32,6,opt,name=keep_checkpoint_every_n_hours,json=keepCheckpointEveryNHours,proto3" json:"keep_checkpoint_every_n_hours,omitempty"`
	Version                   SaverDef_CheckpointFormatVersion `protobuf:"varint,7,opt,name=version,proto3,enum=tensorflow.SaverDef_CheckpointFormatVersion" json:"version,omitempty"`
}

func (m *SaverDef) Reset()                    { *m = SaverDef{} }
func (m *SaverDef) String() string            { return proto.CompactTextString(m) }
func (*SaverDef) ProtoMessage()               {}
func (*SaverDef) Descriptor() ([]byte, []int) { return fileDescriptorSaver, []int{0} }

func (m *SaverDef) GetFilenameTensorName() string {
	if m != nil {
		return m.FilenameTensorName
	}
	return ""
}

func (m *SaverDef) GetSaveTensorName() string {
	if m != nil {
		return m.SaveTensorName
	}
	return ""
}

func (m *SaverDef) GetRestoreOpName() string {
	if m != nil {
		return m.RestoreOpName
	}
	return ""
}

func (m *SaverDef) GetMaxToKeep() int32 {
	if m != nil {
		return m.MaxToKeep
	}
	return 0
}

func (m *SaverDef) GetSharded() bool {
	if m != nil {
		return m.Sharded
	}
	return false
}

func (m *SaverDef) GetKeepCheckpointEveryNHours() float32 {
	if m != nil {
		return m.KeepCheckpointEveryNHours
	}
	return 0
}

func (m *SaverDef) GetVersion() SaverDef_CheckpointFormatVersion {
	if m != nil {
		return m.Version
	}
	return SaverDef_LEGACY
}

func init() {
	proto.RegisterType((*SaverDef)(nil), "tensorflow.SaverDef")
	proto.RegisterEnum("tensorflow.SaverDef_CheckpointFormatVersion", SaverDef_CheckpointFormatVersion_name, SaverDef_CheckpointFormatVersion_value)
}
func (m *SaverDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SaverDef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FilenameTensorName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSaver(dAtA, i, uint64(len(m.FilenameTensorName)))
		i += copy(dAtA[i:], m.FilenameTensorName)
	}
	if len(m.SaveTensorName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSaver(dAtA, i, uint64(len(m.SaveTensorName)))
		i += copy(dAtA[i:], m.SaveTensorName)
	}
	if len(m.RestoreOpName) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSaver(dAtA, i, uint64(len(m.RestoreOpName)))
		i += copy(dAtA[i:], m.RestoreOpName)
	}
	if m.MaxToKeep != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintSaver(dAtA, i, uint64(m.MaxToKeep))
	}
	if m.Sharded {
		dAtA[i] = 0x28
		i++
		if m.Sharded {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.KeepCheckpointEveryNHours != 0 {
		dAtA[i] = 0x35
		i++
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.KeepCheckpointEveryNHours))))
		i += 4
	}
	if m.Version != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintSaver(dAtA, i, uint64(m.Version))
	}
	return i, nil
}

func encodeVarintSaver(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SaverDef) Size() (n int) {
	var l int
	_ = l
	l = len(m.FilenameTensorName)
	if l > 0 {
		n += 1 + l + sovSaver(uint64(l))
	}
	l = len(m.SaveTensorName)
	if l > 0 {
		n += 1 + l + sovSaver(uint64(l))
	}
	l = len(m.RestoreOpName)
	if l > 0 {
		n += 1 + l + sovSaver(uint64(l))
	}
	if m.MaxToKeep != 0 {
		n += 1 + sovSaver(uint64(m.MaxToKeep))
	}
	if m.Sharded {
		n += 2
	}
	if m.KeepCheckpointEveryNHours != 0 {
		n += 5
	}
	if m.Version != 0 {
		n += 1 + sovSaver(uint64(m.Version))
	}
	return n
}

func sovSaver(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSaver(x uint64) (n int) {
	return sovSaver(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SaverDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSaver
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
			return fmt.Errorf("proto: SaverDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SaverDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilenameTensorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSaver
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
				return ErrInvalidLengthSaver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FilenameTensorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SaveTensorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSaver
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
				return ErrInvalidLengthSaver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SaveTensorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RestoreOpName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSaver
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
				return ErrInvalidLengthSaver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RestoreOpName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxToKeep", wireType)
			}
			m.MaxToKeep = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSaver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxToKeep |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sharded", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSaver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Sharded = bool(v != 0)
		case 6:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeepCheckpointEveryNHours", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.KeepCheckpointEveryNHours = float32(math.Float32frombits(v))
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSaver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (SaverDef_CheckpointFormatVersion(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSaver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSaver
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
func skipSaver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSaver
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
					return 0, ErrIntOverflowSaver
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
					return 0, ErrIntOverflowSaver
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
				return 0, ErrInvalidLengthSaver
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSaver
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
				next, err := skipSaver(dAtA[start:])
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
	ErrInvalidLengthSaver = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSaver   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("saver.proto", fileDescriptorSaver) }

var fileDescriptorSaver = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd1, 0x41, 0x4f, 0xea, 0x40,
	0x10, 0x07, 0x70, 0xb6, 0x3c, 0x0a, 0x0c, 0x79, 0xbc, 0x66, 0xdf, 0x4b, 0x5e, 0x3d, 0xd8, 0x34,
	0x1c, 0x4c, 0x0f, 0xa6, 0x51, 0x8c, 0xf1, 0xaa, 0x20, 0x68, 0xa2, 0x41, 0x52, 0x09, 0x89, 0xa7,
	0x4d, 0x85, 0x41, 0x1a, 0x68, 0xb7, 0xd9, 0x16, 0xc4, 0x8f, 0xe0, 0xcd, 0x8f, 0xe5, 0xd1, 0x8f,
	0x60, 0xf0, 0x4b, 0x78, 0x34, 0xbb, 0x58, 0xc1, 0x83, 0xa7, 0xb6, 0x33, 0xbf, 0x7f, 0x33, 0x99,
	0x81, 0x4a, 0xe2, 0xcf, 0x51, 0xb8, 0xb1, 0xe0, 0x29, 0xa7, 0x90, 0x62, 0x94, 0x70, 0x31, 0x9a,
	0xf2, 0xfb, 0xda, 0x63, 0x1e, 0x4a, 0xd7, 0xb2, 0x77, 0x8a, 0x23, 0xba, 0x07, 0xff, 0x46, 0xc1,
	0x14, 0x23, 0x3f, 0x44, 0xb6, 0x32, 0x4c, 0xbe, 0x9b, 0xc4, 0x26, 0x4e, 0xd9, 0xa3, 0x59, 0xaf,
	0xa7, 0x5a, 0x1d, 0x3f, 0x44, 0xea, 0x80, 0x21, 0xff, 0xfc, 0x4d, 0x6b, 0x4a, 0x57, 0x65, 0x7d,
	0x43, 0xee, 0xc0, 0x1f, 0x81, 0x49, 0xca, 0x05, 0x32, 0x1e, 0xaf, 0x60, 0x5e, 0xc1, 0xdf, 0x9f,
	0xe5, 0xab, 0x58, 0x39, 0x0b, 0x2a, 0xa1, 0xbf, 0x60, 0x29, 0x67, 0x13, 0xc4, 0xd8, 0xfc, 0x65,
	0x13, 0xa7, 0xe0, 0x95, 0x43, 0x7f, 0xd1, 0xe3, 0x17, 0x88, 0x31, 0x35, 0xa1, 0x98, 0x8c, 0x7d,
	0x31, 0xc4, 0xa1, 0x59, 0xb0, 0x89, 0x53, 0xf2, 0xb2, 0x4f, 0x7a, 0x0c, 0xdb, 0x32, 0xc2, 0x06,
	0x63, 0x1c, 0x4c, 0x62, 0x1e, 0x44, 0x29, 0xc3, 0x39, 0x8a, 0x07, 0x16, 0xb1, 0x31, 0x9f, 0x89,
	0xc4, 0xd4, 0x6d, 0xe2, 0x68, 0xde, 0x96, 0x44, 0xcd, 0x2f, 0xd3, 0x92, 0xa4, 0x73, 0x2e, 0x01,
	0x6d, 0x43, 0x71, 0x8e, 0x22, 0x09, 0x78, 0x64, 0x16, 0x6d, 0xe2, 0x54, 0xeb, 0xbb, 0xee, 0x7a,
	0x55, 0x6e, 0xb6, 0x26, 0x77, 0x1d, 0x6e, 0x73, 0x11, 0xfa, 0x69, 0x7f, 0x95, 0xf1, 0xb2, 0x70,
	0xed, 0x10, 0xfe, 0xff, 0x60, 0x28, 0x80, 0x7e, 0xd9, 0x3a, 0x3b, 0x69, 0xde, 0x18, 0x39, 0xaa,
	0x83, 0xd6, 0xdf, 0x37, 0x88, 0x7a, 0xd6, 0x0d, 0xad, 0x71, 0xf4, 0xbc, 0xb4, 0xc8, 0xcb, 0xd2,
	0x22, 0xaf, 0x4b, 0x8b, 0x3c, 0xbd, 0x59, 0x39, 0xf8, 0xcb, 0xc5, 0xdd, 0xe6, 0x08, 0xb3, 0x34,
	0x98, 0x36, 0x2a, 0x6a, 0x90, 0xae, 0x3c, 0x65, 0xd2, 0x25, 0xef, 0x84, 0xdc, 0xea, 0xea, 0xae,
	0x07, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x51, 0xc2, 0x2d, 0xe6, 0x01, 0x00, 0x00,
}
