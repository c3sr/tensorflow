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
func (DataType) EnumDescriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func init() {
	proto.RegisterEnum("tensorflow.DataType", DataType_name, DataType_value)
}

func init() { proto.RegisterFile("types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x93, 0xcd, 0x6e, 0xd3, 0x50,
	0x10, 0x85, 0x6b, 0xa0, 0x6d, 0x3a, 0xce, 0xcf, 0xe4, 0xb6, 0x48, 0xac, 0xf2, 0x00, 0x2c, 0x22,
	0xb9, 0xad, 0xac, 0x2c, 0xd8, 0xc4, 0xb5, 0x03, 0x96, 0x8c, 0x6f, 0xec, 0xde, 0x20, 0x76, 0x56,
	0x90, 0x9c, 0x02, 0xfd, 0xb1, 0xb1, 0x0d, 0xa1, 0x6f, 0xc1, 0x33, 0xb1, 0x62, 0xc9, 0x23, 0xa0,
	0xf0, 0x12, 0x2c, 0x91, 0xc7, 0xe3, 0x5e, 0x77, 0x39, 0xdf, 0xcc, 0x9c, 0x39, 0xf7, 0xc8, 0x06,
	0xb3, 0xba, 0xcf, 0xd3, 0x72, 0x9a, 0x17, 0x59, 0x95, 0x09, 0xa8, 0xd2, 0xbb, 0x32, 0x2b, 0x36,
	0x37, 0xd9, 0xf6, 0xe5, 0xcf, 0x7d, 0xe8, 0xb9, 0xeb, 0x6a, 0xad, 0xee, 0xf3, 0x54, 0x0c, 0x01,
	0x5c, 0x95, 0xf8, 0xe1, 0xbb, 0x79, 0xe0, 0xbb, 0xb8, 0x27, 0xfa, 0xd0, 0x73, 0x55, 0xb2, 0x08,
	0xe4, 0x5c, 0xa1, 0x21, 0x06, 0x70, 0xe4, 0xaa, 0xc4, 0x95, 0x2b, 0x27, 0xf0, 0xf0, 0x09, 0x37,
	0xfd, 0x50, 0x9d, 0x9d, 0xe2, 0x53, 0xae, 0x56, 0x7e, 0xa8, 0x66, 0xf8, 0x4c, 0xf7, 0x2c, 0x1b,
	0xf7, 0x85, 0x09, 0x87, 0x4d, 0x35, 0xc3, 0x03, 0x56, 0xb9, 0x54, 0xb1, 0x1f, 0xbe, 0xc6, 0x43,
	0x81, 0xd0, 0x77, 0x55, 0x72, 0x21, 0xdf, 0x2e, 0x03, 0xef, 0xbd, 0x7d, 0x8e, 0x3d, 0xbd, 0x6b,
	0x9f, 0xe3, 0x11, 0xef, 0x3a, 0x52, 0x06, 0x08, 0xdc, 0x8a, 0x48, 0xc9, 0x64, 0xa5, 0xa8, 0xb9,
	0xd9, 0x6f, 0xcb, 0xc6, 0xd0, 0x40, 0x8c, 0xc0, 0xac, 0x17, 0xc9, 0xbc, 0x65, 0xe3, 0xb0, 0xd3,
	0xb7, 0x6c, 0x1c, 0xf1, 0x5b, 0x69, 0xdb, 0xb2, 0x11, 0xb9, 0xcd, 0xe5, 0x58, 0x8c, 0x61, 0xa0,
	0x7d, 0x59, 0xa7, 0x33, 0x14, 0x6c, 0xe5, 0xcd, 0x3c, 0x58, 0xe0, 0x31, 0xcb, 0xc7, 0xde, 0xa5,
	0x5c, 0xc5, 0x17, 0x1e, 0x9e, 0xf0, 0x43, 0xe8, 0x5c, 0x12, 0x7b, 0x0b, 0x4c, 0x59, 0xa2, 0xc9,
	0x8b, 0xd0, 0x86, 0x87, 0xc8, 0x22, 0x91, 0x2b, 0x26, 0xf4, 0x06, 0x22, 0x1f, 0xf5, 0x8c, 0x65,
	0x13, 0xf9, 0xc4, 0xb7, 0x1e, 0x46, 0x3e, 0xb3, 0x72, 0x93, 0x21, 0xa1, 0x6b, 0x71, 0x02, 0xd8,
	0xcd, 0x91, 0xe8, 0x8d, 0xd6, 0x62, 0x72, 0xdb, 0xc6, 0x22, 0x65, 0x40, 0xe0, 0x8e, 0x47, 0xa2,
	0x07, 0xf5, 0x8c, 0xd5, 0x23, 0xed, 0x29, 0x6f, 0x91, 0x36, 0xfe, 0x45, 0x1c, 0xc3, 0xa8, 0x93,
	0x2f, 0xc1, 0xa2, 0x33, 0xc7, 0xa8, 0x14, 0x02, 0x86, 0x3a, 0x67, 0x62, 0x15, 0x8f, 0x75, 0xd0,
	0x57, 0xf1, 0x1c, 0xc6, 0x8f, 0xf2, 0x26, 0xfc, 0x8d, 0xed, 0xd6, 0x99, 0x13, 0xd8, 0xf2, 0xd9,
	0x36, 0x77, 0x82, 0xdf, 0x9d, 0x57, 0xbf, 0x76, 0x13, 0xe3, 0xf7, 0x6e, 0x62, 0xfc, 0xd9, 0x4d,
	0x8c, 0x1f, 0x7f, 0x27, 0x7b, 0xf0, 0x22, 0x2b, 0xae, 0xa6, 0xfa, 0x33, 0x9f, 0x6e, 0x8a, 0xf5,
	0x6d, 0xba, 0xcd, 0x8a, 0x6b, 0xc7, 0xac, 0xbf, 0xf4, 0x72, 0x59, 0xff, 0x08, 0xe5, 0xd2, 0xf8,
	0x67, 0x18, 0x1f, 0x0e, 0xe8, 0xaf, 0x38, 0xfb, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x01, 0x84, 0xfe,
	0xff, 0x24, 0x03, 0x00, 0x00,
}