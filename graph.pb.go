// Code generated by protoc-gen-go. DO NOT EDIT.
// source: graph.proto

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

// Represents the graph of operations
type GraphDef struct {
	Node []*NodeDef `protobuf:"bytes,1,rep,name=node,proto3" json:"node,omitempty"`
	// Compatibility versions of the graph.  See core/public/version.h for version
	// history.  The GraphDef version is distinct from the TensorFlow version, and
	// each release of TensorFlow will support a range of GraphDef versions.
	Versions *VersionDef `protobuf:"bytes,4,opt,name=versions,proto3" json:"versions,omitempty"`
	// Deprecated single version field; use versions above instead.  Since all
	// GraphDef changes before "versions" was introduced were forward
	// compatible, this field is entirely ignored.
	Version int32 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"` // Deprecated: Do not use.
	// EXPERIMENTAL. DO NOT USE OR DEPEND ON THIS YET.
	//
	// "library" provides user-defined functions.
	//
	// Naming:
	//   * library.function.name are in a flat namespace.
	//     NOTE: We may need to change it to be hierarchical to support
	//     different orgs. E.g.,
	//     { "/google/nn", { ... }},
	//     { "/google/vision", { ... }}
	//     { "/org_foo/module_bar", { ... }}
	//     map<string, FunctionDefLib> named_lib;
	//   * If node[i].op is the name of one function in "library",
	//     node[i] is deemed as a function call. Otherwise, node[i].op
	//     must be a primitive operation supported by the runtime.
	//
	//
	// Function call semantics:
	//
	//   * The callee may start execution as soon as some of its inputs
	//     are ready. The caller may want to use Tuple() mechanism to
	//     ensure all inputs are ready in the same time.
	//
	//   * The consumer of return values may start executing as soon as
	//     the return values the consumer depends on are ready.  The
	//     consumer may want to use Tuple() mechanism to ensure the
	//     consumer does not start until all return values of the callee
	//     function are ready.
	Library              *FunctionDefLibrary `protobuf:"bytes,2,opt,name=library,proto3" json:"library,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GraphDef) Reset()         { *m = GraphDef{} }
func (m *GraphDef) String() string { return proto.CompactTextString(m) }
func (*GraphDef) ProtoMessage()    {}
func (*GraphDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e4c656902fc0e6b, []int{0}
}

func (m *GraphDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphDef.Unmarshal(m, b)
}
func (m *GraphDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphDef.Marshal(b, m, deterministic)
}
func (m *GraphDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphDef.Merge(m, src)
}
func (m *GraphDef) XXX_Size() int {
	return xxx_messageInfo_GraphDef.Size(m)
}
func (m *GraphDef) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphDef.DiscardUnknown(m)
}

var xxx_messageInfo_GraphDef proto.InternalMessageInfo

func (m *GraphDef) GetNode() []*NodeDef {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *GraphDef) GetVersions() *VersionDef {
	if m != nil {
		return m.Versions
	}
	return nil
}

// Deprecated: Do not use.
func (m *GraphDef) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *GraphDef) GetLibrary() *FunctionDefLibrary {
	if m != nil {
		return m.Library
	}
	return nil
}

func init() {
	proto.RegisterType((*GraphDef)(nil), "tensorflow.GraphDef")
}

func init() { proto.RegisterFile("graph.proto", fileDescriptor_3e4c656902fc0e6b) }

var fileDescriptor_3e4c656902fc0e6b = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0xd9, 0xb6, 0xda, 0xb2, 0x81, 0x1e, 0x56, 0x90, 0xa5, 0x88, 0x44, 0x2f, 0xe6, 0x62,
	0x0a, 0xf1, 0xe2, 0x39, 0x04, 0xbd, 0x88, 0x96, 0x1c, 0x3c, 0x78, 0x91, 0x7c, 0xcc, 0xa6, 0xab,
	0xed, 0x4e, 0x98, 0x6c, 0x2d, 0xfe, 0x3a, 0xff, 0x96, 0x47, 0x49, 0xd2, 0xb5, 0xb9, 0xed, 0xec,
	0x3c, 0xef, 0x03, 0xef, 0x70, 0xaf, 0xa2, 0xac, 0x5e, 0x87, 0x35, 0xa1, 0x45, 0xc1, 0x2d, 0x98,
	0x06, 0x49, 0x6d, 0x70, 0xbf, 0x98, 0x1b, 0x2c, 0xe1, 0xbd, 0x04, 0xd5, 0xef, 0x16, 0x73, 0xb5,
	0x33, 0x85, 0xd5, 0x68, 0xdc, 0xfc, 0x05, 0xd4, 0x68, 0x34, 0x4d, 0x3f, 0x5f, 0xff, 0x30, 0x3e,
	0x7b, 0x6c, 0x5d, 0x09, 0x28, 0x71, 0xc3, 0x27, 0x6d, 0x5c, 0x32, 0x7f, 0x1c, 0x78, 0xd1, 0x59,
	0x78, 0xf4, 0x86, 0xcf, 0x58, 0x42, 0x02, 0x2a, 0xed, 0x00, 0x11, 0xf1, 0x99, 0xf3, 0xc8, 0x89,
	0xcf, 0x02, 0x2f, 0x3a, 0x1f, 0xc2, 0xaf, 0xfd, 0xae, 0xe5, 0xff, 0x39, 0x71, 0xc1, 0xa7, 0x87,
	0xb7, 0x1c, 0xfb, 0x2c, 0x38, 0x89, 0x47, 0x92, 0xa5, 0xee, 0x4b, 0xdc, 0xf3, 0xe9, 0x46, 0xe7,
	0x94, 0xd1, 0xb7, 0x1c, 0x75, 0xc2, 0xcb, 0xa1, 0xf0, 0xe1, 0x50, 0x22, 0x01, 0xf5, 0xd4, 0x53,
	0xa9, 0xc3, 0xe3, 0x17, 0x2e, 0x91, 0xaa, 0x21, 0xad, 0x28, 0xdb, 0xc2, 0x1e, 0xe9, 0x33, 0xf6,
	0xba, 0x6a, 0xab, 0xb6, 0x69, 0xb3, 0x62, 0x6f, 0x57, 0x95, 0xb6, 0xeb, 0x5d, 0x1e, 0x16, 0xb8,
	0x5d, 0x52, 0xa6, 0x6f, 0x6b, 0xc2, 0x0f, 0x28, 0xec, 0xf2, 0x98, 0xfd, 0x65, 0x2c, 0x3f, 0xed,
	0x2e, 0x73, 0xf7, 0x17, 0x00, 0x00, 0xff, 0xff, 0x85, 0x33, 0xee, 0x44, 0x64, 0x01, 0x00, 0x00,
}
