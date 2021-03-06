// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensor_bundle.proto

package tensorflow

import (
	encoding_binary "encoding/binary"
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

// An enum indicating the endianness of the platform that produced this
// bundle.  A bundle can only be read by a platform with matching endianness.
// Defaults to LITTLE, as most modern platforms are little-endian.
//
// Affects the binary tensor data bytes only, not the metadata in protobufs.
type BundleHeaderProto_Endianness int32

const (
	BundleHeaderProto_LITTLE BundleHeaderProto_Endianness = 0
	BundleHeaderProto_BIG    BundleHeaderProto_Endianness = 1
)

var BundleHeaderProto_Endianness_name = map[int32]string{
	0: "LITTLE",
	1: "BIG",
}

var BundleHeaderProto_Endianness_value = map[string]int32{
	"LITTLE": 0,
	"BIG":    1,
}

func (x BundleHeaderProto_Endianness) String() string {
	return proto.EnumName(BundleHeaderProto_Endianness_name, int32(x))
}

func (BundleHeaderProto_Endianness) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2bf7a626254a7cdc, []int{0, 0}
}

// Special header that is associated with a bundle.
//
// TODO(zongheng,zhifengc): maybe in the future, we can add information about
// which binary produced this checkpoint, timestamp, etc. Sometime, these can be
// valuable debugging information. And if needed, these can be used as defensive
// information ensuring reader (binary version) of the checkpoint and the writer
// (binary version) must match within certain range, etc.
type BundleHeaderProto struct {
	// Number of data files in the bundle.
	NumShards  int32                        `protobuf:"varint,1,opt,name=num_shards,json=numShards,proto3" json:"num_shards,omitempty"`
	Endianness BundleHeaderProto_Endianness `protobuf:"varint,2,opt,name=endianness,proto3,enum=tensorflow.BundleHeaderProto_Endianness" json:"endianness,omitempty"`
	// Versioning of the tensor bundle format.
	Version *VersionDef `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *BundleHeaderProto) Reset()         { *m = BundleHeaderProto{} }
func (m *BundleHeaderProto) String() string { return proto.CompactTextString(m) }
func (*BundleHeaderProto) ProtoMessage()    {}
func (*BundleHeaderProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bf7a626254a7cdc, []int{0}
}
func (m *BundleHeaderProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BundleHeaderProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BundleHeaderProto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BundleHeaderProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BundleHeaderProto.Merge(m, src)
}
func (m *BundleHeaderProto) XXX_Size() int {
	return m.Size()
}
func (m *BundleHeaderProto) XXX_DiscardUnknown() {
	xxx_messageInfo_BundleHeaderProto.DiscardUnknown(m)
}

var xxx_messageInfo_BundleHeaderProto proto.InternalMessageInfo

func (m *BundleHeaderProto) GetNumShards() int32 {
	if m != nil {
		return m.NumShards
	}
	return 0
}

func (m *BundleHeaderProto) GetEndianness() BundleHeaderProto_Endianness {
	if m != nil {
		return m.Endianness
	}
	return BundleHeaderProto_LITTLE
}

func (m *BundleHeaderProto) GetVersion() *VersionDef {
	if m != nil {
		return m.Version
	}
	return nil
}

// Describes the metadata related to a checkpointed tensor.
type BundleEntryProto struct {
	// The tensor dtype and shape.
	Dtype DataType          `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape *TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	// The binary content of the tensor lies in:
	//   File "shard_id": bytes [offset, offset + size).
	ShardId int32 `protobuf:"varint,3,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	Offset  int64 `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Size_   int64 `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	// The CRC32C checksum of the tensor bytes.
	Crc32C uint32 `protobuf:"fixed32,6,opt,name=crc32c,proto3" json:"crc32c,omitempty"`
	// Iff present, this entry represents a partitioned tensor.  The previous
	// fields are interpreted as follows:
	//
	//   "dtype", "shape": describe the full tensor.
	//   "shard_id", "offset", "size", "crc32c": all IGNORED.
	//      These information for each slice can be looked up in their own
	//      BundleEntryProto, keyed by each "slice_name".
	Slices []*TensorSliceProto `protobuf:"bytes,7,rep,name=slices,proto3" json:"slices,omitempty"`
}

func (m *BundleEntryProto) Reset()         { *m = BundleEntryProto{} }
func (m *BundleEntryProto) String() string { return proto.CompactTextString(m) }
func (*BundleEntryProto) ProtoMessage()    {}
func (*BundleEntryProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bf7a626254a7cdc, []int{1}
}
func (m *BundleEntryProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BundleEntryProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BundleEntryProto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BundleEntryProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BundleEntryProto.Merge(m, src)
}
func (m *BundleEntryProto) XXX_Size() int {
	return m.Size()
}
func (m *BundleEntryProto) XXX_DiscardUnknown() {
	xxx_messageInfo_BundleEntryProto.DiscardUnknown(m)
}

var xxx_messageInfo_BundleEntryProto proto.InternalMessageInfo

func (m *BundleEntryProto) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func (m *BundleEntryProto) GetShape() *TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *BundleEntryProto) GetShardId() int32 {
	if m != nil {
		return m.ShardId
	}
	return 0
}

func (m *BundleEntryProto) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *BundleEntryProto) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *BundleEntryProto) GetCrc32C() uint32 {
	if m != nil {
		return m.Crc32C
	}
	return 0
}

func (m *BundleEntryProto) GetSlices() []*TensorSliceProto {
	if m != nil {
		return m.Slices
	}
	return nil
}

func init() {
	proto.RegisterEnum("tensorflow.BundleHeaderProto_Endianness", BundleHeaderProto_Endianness_name, BundleHeaderProto_Endianness_value)
	proto.RegisterType((*BundleHeaderProto)(nil), "tensorflow.BundleHeaderProto")
	proto.RegisterType((*BundleEntryProto)(nil), "tensorflow.BundleEntryProto")
}

func init() { proto.RegisterFile("tensor_bundle.proto", fileDescriptor_2bf7a626254a7cdc) }

var fileDescriptor_2bf7a626254a7cdc = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xcd, 0x36, 0xb5, 0x0d, 0x13, 0x29, 0x0a, 0x5b, 0x54, 0x19, 0x04, 0x96, 0x9b, 0x93, 0x85,
	0x84, 0x8b, 0x5c, 0xbe, 0x20, 0x6a, 0x44, 0x23, 0xf5, 0x50, 0xb9, 0x11, 0x07, 0x2e, 0x95, 0x63,
	0x6f, 0xda, 0x45, 0xc9, 0xae, 0xb5, 0xbb, 0x06, 0x85, 0x1f, 0xe0, 0xca, 0x67, 0x71, 0x2c, 0x37,
	0x8e, 0x28, 0xf9, 0x09, 0x8e, 0xc8, 0xb3, 0x5b, 0x62, 0x09, 0x71, 0xdb, 0xf7, 0xe6, 0xcd, 0x9b,
	0x37, 0x63, 0xc3, 0x91, 0x61, 0x42, 0x4b, 0x75, 0xb3, 0x68, 0x44, 0xb5, 0x62, 0x69, 0xad, 0xa4,
	0x91, 0x14, 0x2c, 0xb9, 0x5c, 0xc9, 0xcf, 0xcf, 0xa9, 0x13, 0xe8, 0xbb, 0xa2, 0x76, 0xf5, 0x3d,
	0xb7, 0xe2, 0xe5, 0x03, 0x37, 0x30, 0x9b, 0x9a, 0x69, 0x07, 0x86, 0x9f, 0x98, 0xd2, 0x5c, 0x0a,
	0x87, 0xc7, 0x3f, 0x08, 0x3c, 0x99, 0xe0, 0x84, 0x0b, 0x56, 0x54, 0x4c, 0x5d, 0xe1, 0x98, 0x97,
	0x00, 0xa2, 0x59, 0xb7, 0xce, 0xaa, 0xd2, 0x21, 0x89, 0x49, 0xe2, 0xe5, 0x8f, 0x45, 0xb3, 0xbe,
	0x46, 0x82, 0x5e, 0x00, 0x30, 0x51, 0xf1, 0x42, 0x08, 0xa6, 0x75, 0x78, 0x10, 0x93, 0x64, 0x98,
	0x25, 0xe9, 0x3e, 0x5a, 0xfa, 0x8f, 0x63, 0x3a, 0xfd, 0xab, 0xcf, 0x3b, 0xbd, 0xf4, 0x0d, 0x04,
	0x2e, 0x50, 0xd8, 0x8f, 0x49, 0x32, 0xc8, 0x8e, 0xbb, 0x36, 0xef, 0x6d, 0xe9, 0x9c, 0x2d, 0xf3,
	0x07, 0xd9, 0xf8, 0x04, 0x60, 0xef, 0x45, 0x01, 0xfc, 0xcb, 0xd9, 0x7c, 0x7e, 0x39, 0x1d, 0xf5,
	0x68, 0x00, 0xfd, 0xc9, 0xec, 0xdd, 0x88, 0x8c, 0xbf, 0x1e, 0xc0, 0xc8, 0x26, 0x98, 0x0a, 0xa3,
	0x36, 0x76, 0xa5, 0x57, 0xe0, 0x55, 0xed, 0x21, 0x70, 0x9b, 0x61, 0xf6, 0xb4, 0x3b, 0xe7, 0xbc,
	0x30, 0xc5, 0x7c, 0x53, 0xb3, 0xdc, 0x4a, 0x68, 0x06, 0x1e, 0x1e, 0x15, 0x57, 0x1b, 0x64, 0x2f,
	0xba, 0xda, 0x39, 0x3e, 0xaf, 0xdb, 0x32, 0x1a, 0xe7, 0x56, 0x4a, 0x9f, 0xc1, 0x23, 0x3c, 0xd7,
	0x0d, 0xaf, 0x70, 0x15, 0x2f, 0x0f, 0x10, 0xcf, 0x2a, 0x7a, 0x0c, 0xbe, 0x5c, 0x2e, 0x35, 0x33,
	0xe1, 0x61, 0x4c, 0x92, 0x7e, 0xee, 0x10, 0xa5, 0x70, 0xa8, 0xf9, 0x17, 0x16, 0x7a, 0xc8, 0xe2,
	0xbb, 0xd5, 0x96, 0xaa, 0x3c, 0xcb, 0xca, 0xd0, 0x8f, 0x49, 0x12, 0xe4, 0x0e, 0xd1, 0xb7, 0xe0,
	0xe3, 0x37, 0xd5, 0x61, 0x10, 0xf7, 0xff, 0x93, 0xa9, 0xad, 0xdb, 0x4c, 0x4e, 0x3b, 0xe1, 0xdf,
	0xb7, 0x11, 0xb9, 0xdf, 0x46, 0xe4, 0xd7, 0x36, 0x22, 0xdf, 0x76, 0x51, 0xef, 0x7e, 0x17, 0xf5,
	0x7e, 0xee, 0xa2, 0x1e, 0x1c, 0x49, 0x75, 0xdb, 0xb5, 0x68, 0x0c, 0x5f, 0x4d, 0xa8, 0x35, 0xb2,
	0xb7, 0x43, 0x27, 0x7d, 0x45, 0x3e, 0x9c, 0xdc, 0x72, 0x73, 0xd7, 0x2c, 0xd2, 0x52, 0xae, 0x4f,
	0x55, 0xc1, 0x5f, 0xd7, 0x4a, 0x7e, 0x64, 0xa5, 0x39, 0xdd, 0x77, 0xff, 0x26, 0x64, 0xe1, 0xe3,
	0xff, 0x74, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x69, 0x34, 0x5b, 0xb7, 0x02, 0x00, 0x00,
}

func (m *BundleHeaderProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BundleHeaderProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NumShards != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTensorBundle(dAtA, i, uint64(m.NumShards))
	}
	if m.Endianness != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTensorBundle(dAtA, i, uint64(m.Endianness))
	}
	if m.Version != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTensorBundle(dAtA, i, uint64(m.Version.Size()))
		n1, err1 := m.Version.MarshalTo(dAtA[i:])
		if err1 != nil {
			return 0, err1
		}
		i += n1
	}
	return i, nil
}

func (m *BundleEntryProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BundleEntryProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Dtype != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTensorBundle(dAtA, i, uint64(m.Dtype))
	}
	if m.Shape != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTensorBundle(dAtA, i, uint64(m.Shape.Size()))
		n2, err2 := m.Shape.MarshalTo(dAtA[i:])
		if err2 != nil {
			return 0, err2
		}
		i += n2
	}
	if m.ShardId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintTensorBundle(dAtA, i, uint64(m.ShardId))
	}
	if m.Offset != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintTensorBundle(dAtA, i, uint64(m.Offset))
	}
	if m.Size_ != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintTensorBundle(dAtA, i, uint64(m.Size_))
	}
	if m.Crc32C != 0 {
		dAtA[i] = 0x35
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.Crc32C))
		i += 4
	}
	if len(m.Slices) > 0 {
		for _, msg := range m.Slices {
			dAtA[i] = 0x3a
			i++
			i = encodeVarintTensorBundle(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintTensorBundle(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BundleHeaderProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumShards != 0 {
		n += 1 + sovTensorBundle(uint64(m.NumShards))
	}
	if m.Endianness != 0 {
		n += 1 + sovTensorBundle(uint64(m.Endianness))
	}
	if m.Version != nil {
		l = m.Version.Size()
		n += 1 + l + sovTensorBundle(uint64(l))
	}
	return n
}

func (m *BundleEntryProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Dtype != 0 {
		n += 1 + sovTensorBundle(uint64(m.Dtype))
	}
	if m.Shape != nil {
		l = m.Shape.Size()
		n += 1 + l + sovTensorBundle(uint64(l))
	}
	if m.ShardId != 0 {
		n += 1 + sovTensorBundle(uint64(m.ShardId))
	}
	if m.Offset != 0 {
		n += 1 + sovTensorBundle(uint64(m.Offset))
	}
	if m.Size_ != 0 {
		n += 1 + sovTensorBundle(uint64(m.Size_))
	}
	if m.Crc32C != 0 {
		n += 5
	}
	if len(m.Slices) > 0 {
		for _, e := range m.Slices {
			l = e.Size()
			n += 1 + l + sovTensorBundle(uint64(l))
		}
	}
	return n
}

func sovTensorBundle(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTensorBundle(x uint64) (n int) {
	return sovTensorBundle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BundleHeaderProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTensorBundle
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
			return fmt.Errorf("proto: BundleHeaderProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BundleHeaderProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumShards", wireType)
			}
			m.NumShards = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorBundle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumShards |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endianness", wireType)
			}
			m.Endianness = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorBundle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Endianness |= BundleHeaderProto_Endianness(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorBundle
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
				return ErrInvalidLengthTensorBundle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTensorBundle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Version == nil {
				m.Version = &VersionDef{}
			}
			if err := m.Version.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTensorBundle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTensorBundle
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTensorBundle
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
func (m *BundleEntryProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTensorBundle
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
			return fmt.Errorf("proto: BundleEntryProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BundleEntryProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dtype", wireType)
			}
			m.Dtype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorBundle
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
					return ErrIntOverflowTensorBundle
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
				return ErrInvalidLengthTensorBundle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTensorBundle
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardId", wireType)
			}
			m.ShardId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorBundle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShardId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorBundle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorBundle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Crc32C", wireType)
			}
			m.Crc32C = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.Crc32C = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorBundle
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
				return ErrInvalidLengthTensorBundle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTensorBundle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slices = append(m.Slices, &TensorSliceProto{})
			if err := m.Slices[len(m.Slices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTensorBundle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTensorBundle
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTensorBundle
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
func skipTensorBundle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTensorBundle
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
					return 0, ErrIntOverflowTensorBundle
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
					return 0, ErrIntOverflowTensorBundle
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
				return 0, ErrInvalidLengthTensorBundle
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTensorBundle
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTensorBundle
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
				next, err := skipTensorBundle(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTensorBundle
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
	ErrInvalidLengthTensorBundle = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTensorBundle   = fmt.Errorf("proto: integer overflow")
)
