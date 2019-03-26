// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types.proto

package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// LINT.IfChange
type DataType int32

const (
	// Not a legal value for DataType.  Used to indicate a DataType field
	// has not been set.
	DataType_DT_INVALID DataType = 0
	// Data types that all computation devices are expected to be
	// capable to support.
	DataType_DT_FLOAT      DataType = 1
	DataType_DT_DOUBLE     DataType = 2
	DataType_DT_INT32      DataType = 3
	DataType_DT_UINT8      DataType = 4
	DataType_DT_INT16      DataType = 5
	DataType_DT_INT8       DataType = 6
	DataType_DT_STRING     DataType = 7
	DataType_DT_COMPLEX64  DataType = 8
	DataType_DT_INT64      DataType = 9
	DataType_DT_BOOL       DataType = 10
	DataType_DT_QINT8      DataType = 11
	DataType_DT_QUINT8     DataType = 12
	DataType_DT_QINT32     DataType = 13
	DataType_DT_BFLOAT16   DataType = 14
	DataType_DT_QINT16     DataType = 15
	DataType_DT_QUINT16    DataType = 16
	DataType_DT_UINT16     DataType = 17
	DataType_DT_COMPLEX128 DataType = 18
	DataType_DT_HALF       DataType = 19
	DataType_DT_RESOURCE   DataType = 20
	// Do not use!  These are only for parameters.  Every enum above
	// should have a corresponding value below (verified by types_test).
	DataType_DT_FLOAT_REF      DataType = 101
	DataType_DT_DOUBLE_REF     DataType = 102
	DataType_DT_INT32_REF      DataType = 103
	DataType_DT_UINT8_REF      DataType = 104
	DataType_DT_INT16_REF      DataType = 105
	DataType_DT_INT8_REF       DataType = 106
	DataType_DT_STRING_REF     DataType = 107
	DataType_DT_COMPLEX64_REF  DataType = 108
	DataType_DT_INT64_REF      DataType = 109
	DataType_DT_BOOL_REF       DataType = 110
	DataType_DT_QINT8_REF      DataType = 111
	DataType_DT_QUINT8_REF     DataType = 112
	DataType_DT_QINT32_REF     DataType = 113
	DataType_DT_BFLOAT16_REF   DataType = 114
	DataType_DT_QINT16_REF     DataType = 115
	DataType_DT_QUINT16_REF    DataType = 116
	DataType_DT_UINT16_REF     DataType = 117
	DataType_DT_COMPLEX128_REF DataType = 118
	DataType_DT_HALF_REF       DataType = 119
	DataType_DT_RESOURCE_REF   DataType = 120
)

var DataType_name = map[int32]string{
	0:   "DT_INVALID",
	1:   "DT_FLOAT",
	2:   "DT_DOUBLE",
	3:   "DT_INT32",
	4:   "DT_UINT8",
	5:   "DT_INT16",
	6:   "DT_INT8",
	7:   "DT_STRING",
	8:   "DT_COMPLEX64",
	9:   "DT_INT64",
	10:  "DT_BOOL",
	11:  "DT_QINT8",
	12:  "DT_QUINT8",
	13:  "DT_QINT32",
	14:  "DT_BFLOAT16",
	15:  "DT_QINT16",
	16:  "DT_QUINT16",
	17:  "DT_UINT16",
	18:  "DT_COMPLEX128",
	19:  "DT_HALF",
	20:  "DT_RESOURCE",
	101: "DT_FLOAT_REF",
	102: "DT_DOUBLE_REF",
	103: "DT_INT32_REF",
	104: "DT_UINT8_REF",
	105: "DT_INT16_REF",
	106: "DT_INT8_REF",
	107: "DT_STRING_REF",
	108: "DT_COMPLEX64_REF",
	109: "DT_INT64_REF",
	110: "DT_BOOL_REF",
	111: "DT_QINT8_REF",
	112: "DT_QUINT8_REF",
	113: "DT_QINT32_REF",
	114: "DT_BFLOAT16_REF",
	115: "DT_QINT16_REF",
	116: "DT_QUINT16_REF",
	117: "DT_UINT16_REF",
	118: "DT_COMPLEX128_REF",
	119: "DT_HALF_REF",
	120: "DT_RESOURCE_REF",
}
var DataType_value = map[string]int32{
	"DT_INVALID":        0,
	"DT_FLOAT":          1,
	"DT_DOUBLE":         2,
	"DT_INT32":          3,
	"DT_UINT8":          4,
	"DT_INT16":          5,
	"DT_INT8":           6,
	"DT_STRING":         7,
	"DT_COMPLEX64":      8,
	"DT_INT64":          9,
	"DT_BOOL":           10,
	"DT_QINT8":          11,
	"DT_QUINT8":         12,
	"DT_QINT32":         13,
	"DT_BFLOAT16":       14,
	"DT_QINT16":         15,
	"DT_QUINT16":        16,
	"DT_UINT16":         17,
	"DT_COMPLEX128":     18,
	"DT_HALF":           19,
	"DT_RESOURCE":       20,
	"DT_FLOAT_REF":      101,
	"DT_DOUBLE_REF":     102,
	"DT_INT32_REF":      103,
	"DT_UINT8_REF":      104,
	"DT_INT16_REF":      105,
	"DT_INT8_REF":       106,
	"DT_STRING_REF":     107,
	"DT_COMPLEX64_REF":  108,
	"DT_INT64_REF":      109,
	"DT_BOOL_REF":       110,
	"DT_QINT8_REF":      111,
	"DT_QUINT8_REF":     112,
	"DT_QINT32_REF":     113,
	"DT_BFLOAT16_REF":   114,
	"DT_QINT16_REF":     115,
	"DT_QUINT16_REF":    116,
	"DT_UINT16_REF":     117,
	"DT_COMPLEX128_REF": 118,
	"DT_HALF_REF":       119,
	"DT_RESOURCE_REF":   120,
}

func (x DataType) String() string {
	return proto.EnumName(DataType_name, int32(x))
}
func (DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_types_95c72c0adfd0761c, []int{0}
}

func init() {
	proto.RegisterEnum("tensorflow.DataType", DataType_name, DataType_value)
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_types_95c72c0adfd0761c) }

var fileDescriptor_types_95c72c0adfd0761c = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x63, 0xa0, 0x6d, 0x3a, 0xce, 0x9f, 0xc9, 0xb6, 0x48, 0x9c, 0xfc, 0x00, 0x1c, 0x22,
	0xb9, 0xad, 0xac, 0x5c, 0x93, 0xae, 0x03, 0x96, 0x4c, 0x36, 0x49, 0x37, 0x88, 0x9b, 0x15, 0x24,
	0xa7, 0x40, 0xff, 0x38, 0xd8, 0x86, 0xd0, 0xb7, 0xe0, 0x99, 0x38, 0x71, 0xec, 0x91, 0x23, 0x4a,
	0x5e, 0x82, 0x23, 0xf2, 0x78, 0x9c, 0x75, 0x8f, 0xf3, 0x9b, 0x99, 0x6f, 0xbe, 0xfd, 0x64, 0x83,
	0x9d, 0x3f, 0xac, 0xe3, 0xac, 0xbf, 0x4e, 0x93, 0x3c, 0x11, 0x90, 0xc7, 0xf7, 0x59, 0x92, 0xae,
	0x6e, 0x93, 0xcd, 0xeb, 0x5f, 0x07, 0xd0, 0x94, 0xcb, 0x7c, 0xa9, 0x1f, 0xd6, 0xb1, 0xe8, 0x00,
	0x48, 0x1d, 0x05, 0x93, 0xf7, 0xc3, 0x30, 0x90, 0xd8, 0x10, 0x2d, 0x68, 0x4a, 0x1d, 0x8d, 0x43,
	0x35, 0xd4, 0x68, 0x89, 0x36, 0x1c, 0x4b, 0x1d, 0x49, 0xb5, 0x18, 0x85, 0x3e, 0x3e, 0xe3, 0x66,
	0x30, 0xd1, 0xe7, 0x67, 0xf8, 0x9c, 0xab, 0x45, 0x30, 0xd1, 0x03, 0x7c, 0x61, 0x7a, 0xae, 0x87,
	0x07, 0xc2, 0x86, 0xa3, 0xb2, 0x1a, 0xe0, 0x21, 0xab, 0x5c, 0xe9, 0x79, 0x30, 0x79, 0x83, 0x47,
	0x02, 0xa1, 0x25, 0x75, 0x74, 0xa9, 0xde, 0x4d, 0x43, 0xff, 0x83, 0x77, 0x81, 0x4d, 0xb3, 0xeb,
	0x5d, 0xe0, 0x31, 0xef, 0x8e, 0x94, 0x0a, 0x11, 0xb8, 0x35, 0x23, 0x25, 0x9b, 0x95, 0x66, 0xe5,
	0xcd, 0x56, 0x55, 0x96, 0x86, 0xda, 0xa2, 0x0b, 0x76, 0xb1, 0x48, 0xe6, 0x5d, 0x0f, 0x3b, 0xb5,
	0xbe, 0xeb, 0x61, 0x97, 0xdf, 0x4a, 0xdb, 0xae, 0x87, 0xc8, 0x6d, 0x2e, 0x7b, 0xa2, 0x07, 0x6d,
	0xe3, 0xcb, 0x3d, 0x1b, 0xa0, 0x60, 0x2b, 0x6f, 0x87, 0xe1, 0x18, 0x4f, 0x58, 0x7e, 0xee, 0x5f,
	0xa9, 0xc5, 0xfc, 0xd2, 0xc7, 0x53, 0x7e, 0x08, 0x9d, 0x8b, 0xe6, 0xfe, 0x18, 0x63, 0x96, 0x28,
	0xf3, 0x22, 0xb4, 0xe2, 0x21, 0xb2, 0x48, 0xe4, 0x9a, 0x09, 0xbd, 0x81, 0xc8, 0x27, 0x33, 0xe3,
	0x7a, 0x44, 0x3e, 0xf3, 0xad, 0xfd, 0xc8, 0x17, 0x56, 0x2e, 0x33, 0x24, 0x74, 0x23, 0x4e, 0x01,
	0xeb, 0x39, 0x12, 0xbd, 0x35, 0x5a, 0x4c, 0xee, 0xaa, 0x58, 0x94, 0x0a, 0x09, 0xdc, 0xf3, 0xc8,
	0x6c, 0xaf, 0x9e, 0xb0, 0xfa, 0xcc, 0x78, 0x5a, 0x57, 0xc8, 0x18, 0xff, 0x2a, 0x4e, 0xa0, 0x5b,
	0xcb, 0x97, 0x60, 0x5a, 0x9b, 0x63, 0x94, 0x09, 0x01, 0x1d, 0x93, 0x33, 0xb1, 0x9c, 0xc7, 0x6a,
	0xe8, 0x9b, 0x78, 0x09, 0xbd, 0x27, 0x79, 0x13, 0xfe, 0xce, 0x76, 0x8b, 0xcc, 0x09, 0x6c, 0xf8,
	0x6c, 0x95, 0x3b, 0xc1, 0x1f, 0x23, 0xf9, 0x7b, 0xeb, 0x58, 0x8f, 0x5b, 0xc7, 0xfa, 0xbb, 0x75,
	0xac, 0x9f, 0x3b, 0xa7, 0xf1, 0xb8, 0x73, 0x1a, 0x7f, 0x76, 0x4e, 0x03, 0x5e, 0x25, 0xe9, 0x75,
	0xdf, 0x7c, 0xee, 0xfd, 0x55, 0xba, 0xbc, 0x8b, 0x37, 0x49, 0x7a, 0x33, 0xb2, 0x8b, 0x2f, 0x3e,
	0x9b, 0x16, 0x3f, 0x44, 0x36, 0xb5, 0xfe, 0x59, 0xd6, 0xc7, 0x43, 0xfa, 0x3b, 0xce, 0xff, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x22, 0x73, 0xbf, 0xb4, 0x2c, 0x03, 0x00, 0x00,
}
