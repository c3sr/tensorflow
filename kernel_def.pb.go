// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kernel_def.proto

package tensorflow

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KernelDef struct {
	// Must match the name of an Op.
	Op string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	// Type of device this kernel runs on.
	DeviceType string                      `protobuf:"bytes,2,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	Constraint []*KernelDef_AttrConstraint `protobuf:"bytes,3,rep,name=constraint,proto3" json:"constraint,omitempty"`
	// Names of the Op's input_/output_args that reside in host memory
	// instead of device memory.
	HostMemoryArg []string `protobuf:"bytes,4,rep,name=host_memory_arg,json=hostMemoryArg,proto3" json:"host_memory_arg,omitempty"`
	// This allows experimental kernels to be registered for an op that
	// won't be used unless the user specifies a "_kernel" attr with
	// value matching this.
	Label                string   `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KernelDef) Reset()         { *m = KernelDef{} }
func (m *KernelDef) String() string { return proto.CompactTextString(m) }
func (*KernelDef) ProtoMessage()    {}
func (*KernelDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8551f1c1caf11ba, []int{0}
}

func (m *KernelDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelDef.Unmarshal(m, b)
}
func (m *KernelDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelDef.Marshal(b, m, deterministic)
}
func (m *KernelDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelDef.Merge(m, src)
}
func (m *KernelDef) XXX_Size() int {
	return xxx_messageInfo_KernelDef.Size(m)
}
func (m *KernelDef) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelDef.DiscardUnknown(m)
}

var xxx_messageInfo_KernelDef proto.InternalMessageInfo

func (m *KernelDef) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *KernelDef) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *KernelDef) GetConstraint() []*KernelDef_AttrConstraint {
	if m != nil {
		return m.Constraint
	}
	return nil
}

func (m *KernelDef) GetHostMemoryArg() []string {
	if m != nil {
		return m.HostMemoryArg
	}
	return nil
}

func (m *KernelDef) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type KernelDef_AttrConstraint struct {
	// Name of an attr from the Op.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of values that this kernel supports for this attr.
	// Like OpDef.AttrDef.allowed_values, except for kernels instead of Ops.
	AllowedValues        *AttrValue `protobuf:"bytes,2,opt,name=allowed_values,json=allowedValues,proto3" json:"allowed_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *KernelDef_AttrConstraint) Reset()         { *m = KernelDef_AttrConstraint{} }
func (m *KernelDef_AttrConstraint) String() string { return proto.CompactTextString(m) }
func (*KernelDef_AttrConstraint) ProtoMessage()    {}
func (*KernelDef_AttrConstraint) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8551f1c1caf11ba, []int{0, 0}
}

func (m *KernelDef_AttrConstraint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelDef_AttrConstraint.Unmarshal(m, b)
}
func (m *KernelDef_AttrConstraint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelDef_AttrConstraint.Marshal(b, m, deterministic)
}
func (m *KernelDef_AttrConstraint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelDef_AttrConstraint.Merge(m, src)
}
func (m *KernelDef_AttrConstraint) XXX_Size() int {
	return xxx_messageInfo_KernelDef_AttrConstraint.Size(m)
}
func (m *KernelDef_AttrConstraint) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelDef_AttrConstraint.DiscardUnknown(m)
}

var xxx_messageInfo_KernelDef_AttrConstraint proto.InternalMessageInfo

func (m *KernelDef_AttrConstraint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KernelDef_AttrConstraint) GetAllowedValues() *AttrValue {
	if m != nil {
		return m.AllowedValues
	}
	return nil
}

// A collection of KernelDefs
type KernelList struct {
	Kernel               []*KernelDef `protobuf:"bytes,1,rep,name=kernel,proto3" json:"kernel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *KernelList) Reset()         { *m = KernelList{} }
func (m *KernelList) String() string { return proto.CompactTextString(m) }
func (*KernelList) ProtoMessage()    {}
func (*KernelList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8551f1c1caf11ba, []int{1}
}

func (m *KernelList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelList.Unmarshal(m, b)
}
func (m *KernelList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelList.Marshal(b, m, deterministic)
}
func (m *KernelList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelList.Merge(m, src)
}
func (m *KernelList) XXX_Size() int {
	return xxx_messageInfo_KernelList.Size(m)
}
func (m *KernelList) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelList.DiscardUnknown(m)
}

var xxx_messageInfo_KernelList proto.InternalMessageInfo

func (m *KernelList) GetKernel() []*KernelDef {
	if m != nil {
		return m.Kernel
	}
	return nil
}

func init() {
	proto.RegisterType((*KernelDef)(nil), "tensorflow.KernelDef")
	proto.RegisterType((*KernelDef_AttrConstraint)(nil), "tensorflow.KernelDef.AttrConstraint")
	proto.RegisterType((*KernelList)(nil), "tensorflow.KernelList")
}

func init() { proto.RegisterFile("kernel_def.proto", fileDescriptor_c8551f1c1caf11ba) }

var fileDescriptor_c8551f1c1caf11ba = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x6b, 0xc2, 0x40,
	0x10, 0xc5, 0x49, 0xfc, 0x03, 0x8e, 0xf8, 0x87, 0xa5, 0x42, 0xf0, 0x52, 0x2b, 0xa5, 0x78, 0x31,
	0x82, 0x3d, 0xb6, 0x17, 0xad, 0xb7, 0xb6, 0x20, 0x69, 0xe9, 0xa1, 0x97, 0xb0, 0x89, 0x93, 0x98,
	0x9a, 0x64, 0xc2, 0x64, 0x55, 0xfc, 0x3e, 0xfd, 0x90, 0x3d, 0x96, 0x6c, 0x44, 0x2d, 0xf4, 0x96,
	0x99, 0x79, 0xef, 0x65, 0x7e, 0x3b, 0xd0, 0xdd, 0x20, 0xa7, 0x18, 0xbb, 0x2b, 0x0c, 0xec, 0x8c,
	0x49, 0x91, 0x00, 0x85, 0x69, 0x4e, 0x1c, 0xc4, 0xb4, 0xef, 0x77, 0xa5, 0x52, 0xec, 0xee, 0x64,
	0xbc, 0xc5, 0x72, 0x3a, 0xfc, 0x36, 0xa1, 0xf1, 0xac, 0x2d, 0x0b, 0x0c, 0x44, 0x1b, 0x4c, 0xca,
	0x2c, 0x63, 0x60, 0x8c, 0x1a, 0x8e, 0x49, 0x99, 0xb8, 0x86, 0xe6, 0x0a, 0x77, 0x91, 0x8f, 0xae,
	0x3a, 0x64, 0x68, 0x99, 0x7a, 0x00, 0x65, 0xeb, 0xfd, 0x90, 0xa1, 0x58, 0x00, 0xf8, 0x94, 0xe6,
	0x8a, 0x65, 0x94, 0x2a, 0xab, 0x32, 0xa8, 0x8c, 0x9a, 0xd3, 0x5b, 0xfb, 0xfc, 0x47, 0xfb, 0x94,
	0x6d, 0xcf, 0x94, 0xe2, 0xa7, 0x93, 0xd6, 0xb9, 0xf0, 0x89, 0x3b, 0xe8, 0xac, 0x29, 0x57, 0x6e,
	0x82, 0x09, 0xf1, 0xc1, 0x95, 0x1c, 0x5a, 0xd5, 0x41, 0x65, 0xd4, 0x70, 0x5a, 0x45, 0xfb, 0x55,
	0x77, 0x67, 0x1c, 0x8a, 0x2b, 0xa8, 0xc5, 0xd2, 0xc3, 0xd8, 0xaa, 0xe9, 0x45, 0xca, 0xa2, 0xef,
	0x41, 0xfb, 0x6f, 0xb6, 0x10, 0x50, 0x4d, 0x65, 0x82, 0x47, 0x10, 0xfd, 0x2d, 0x1e, 0xa1, 0x2d,
	0xe3, 0x98, 0xf6, 0xb8, 0x2a, 0xf9, 0x73, 0x4d, 0xd3, 0x9c, 0xf6, 0x2e, 0xb7, 0x2d, 0x72, 0x3e,
	0x8a, 0xa9, 0xd3, 0x3a, 0x8a, 0x75, 0x95, 0x0f, 0x1f, 0x00, 0x4a, 0x92, 0x97, 0x28, 0x57, 0x62,
	0x0c, 0xf5, 0xf2, 0x99, 0x2d, 0x43, 0x13, 0xf7, 0xfe, 0x25, 0x76, 0x8e, 0xa2, 0xf9, 0x1b, 0x58,
	0xc4, 0xe1, 0xa5, 0x26, 0x60, 0x99, 0xe0, 0x9e, 0x78, 0x33, 0xef, 0x9c, 0xe4, 0xcb, 0xe2, 0x1e,
	0xf9, 0xd2, 0xf8, 0xbc, 0x09, 0x23, 0xb5, 0xde, 0x7a, 0xb6, 0x4f, 0xc9, 0x84, 0x65, 0x34, 0xce,
	0x98, 0xbe, 0xd0, 0x57, 0x93, 0xb3, 0xff, 0xc7, 0x30, 0xbc, 0xba, 0xbe, 0xdf, 0xfd, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xeb, 0x99, 0x2c, 0x0d, 0xf1, 0x01, 0x00, 0x00,
}
