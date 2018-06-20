// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: graph.proto

package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Represents the graph of operations
type GraphDef struct {
	Node []*NodeDef `protobuf:"bytes,1,rep,name=node" json:"node,omitempty"`
	// Compatibility versions of the graph.  See core/public/version.h for version
	// history.  The GraphDef version is distinct from the TensorFlow version, and
	// each release of TensorFlow will support a range of GraphDef versions.
	Versions *VersionDef `protobuf:"bytes,4,opt,name=versions" json:"versions,omitempty"`
	// Deprecated single version field; use versions above instead.  Since all
	// GraphDef changes before "versions" was introduced were forward
	// compatible, this field is entirely ignored.
	Version int32 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
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
	Library *FunctionDefLibrary `protobuf:"bytes,2,opt,name=library" json:"library,omitempty"`
}

func (m *GraphDef) Reset()                    { *m = GraphDef{} }
func (m *GraphDef) String() string            { return proto.CompactTextString(m) }
func (*GraphDef) ProtoMessage()               {}
func (*GraphDef) Descriptor() ([]byte, []int) { return fileDescriptorGraph, []int{0} }

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
func (m *GraphDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GraphDef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Node) > 0 {
		for _, msg := range m.Node {
			dAtA[i] = 0xa
			i++
			i = encodeVarintGraph(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Library != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGraph(dAtA, i, uint64(m.Library.Size()))
		n1, err := m.Library.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Version != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGraph(dAtA, i, uint64(m.Version))
	}
	if m.Versions != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintGraph(dAtA, i, uint64(m.Versions.Size()))
		n2, err := m.Versions.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeVarintGraph(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GraphDef) Size() (n int) {
	var l int
	_ = l
	if len(m.Node) > 0 {
		for _, e := range m.Node {
			l = e.Size()
			n += 1 + l + sovGraph(uint64(l))
		}
	}
	if m.Library != nil {
		l = m.Library.Size()
		n += 1 + l + sovGraph(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovGraph(uint64(m.Version))
	}
	if m.Versions != nil {
		l = m.Versions.Size()
		n += 1 + l + sovGraph(uint64(l))
	}
	return n
}

func sovGraph(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGraph(x uint64) (n int) {
	return sovGraph(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GraphDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGraph
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
			return fmt.Errorf("proto: GraphDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GraphDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGraph
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
				return ErrInvalidLengthGraph
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = append(m.Node, &NodeDef{})
			if err := m.Node[len(m.Node)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Library", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGraph
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
				return ErrInvalidLengthGraph
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Library == nil {
				m.Library = &FunctionDefLibrary{}
			}
			if err := m.Library.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGraph
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Versions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGraph
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
				return ErrInvalidLengthGraph
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Versions == nil {
				m.Versions = &VersionDef{}
			}
			if err := m.Versions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGraph(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGraph
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
func skipGraph(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGraph
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
					return 0, ErrIntOverflowGraph
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
					return 0, ErrIntOverflowGraph
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
				return 0, ErrInvalidLengthGraph
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGraph
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
				next, err := skipGraph(dAtA[start:])
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
	ErrInvalidLengthGraph = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGraph   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("graph.proto", fileDescriptorGraph) }

var fileDescriptorGraph = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2f, 0x4a, 0x2c,
	0xc8, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2a, 0x49, 0xcd, 0x2b, 0xce, 0x2f, 0x4a,
	0xcb, 0xc9, 0x2f, 0x97, 0xe2, 0xcb, 0xcb, 0x4f, 0x49, 0x8d, 0x4f, 0x49, 0x4d, 0x83, 0xc8, 0x49,
	0xf1, 0xa5, 0x95, 0xe6, 0x25, 0x97, 0x64, 0xe6, 0xe7, 0xc1, 0xf8, 0x65, 0xa9, 0x45, 0xc5, 0x99,
	0xf9, 0x79, 0xc5, 0x10, 0xbe, 0xd2, 0x7e, 0x46, 0x2e, 0x0e, 0x77, 0x90, 0x59, 0x2e, 0xa9, 0x69,
	0x42, 0xea, 0x5c, 0x2c, 0x20, 0xed, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xc2, 0x7a, 0x08,
	0x73, 0xf5, 0xfc, 0xf2, 0x53, 0x52, 0x5d, 0x52, 0xd3, 0x82, 0xc0, 0x0a, 0x84, 0x2c, 0xb8, 0xd8,
	0x73, 0x32, 0x93, 0x8a, 0x12, 0x8b, 0x2a, 0x25, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0xe4, 0x90,
	0xd5, 0xba, 0x41, 0xad, 0x74, 0x49, 0x4d, 0xf3, 0x81, 0xa8, 0x0a, 0x82, 0x29, 0x17, 0x92, 0xe1,
	0x62, 0x87, 0xba, 0x40, 0x82, 0x59, 0x81, 0x51, 0x83, 0xd5, 0x89, 0x49, 0x82, 0x31, 0x08, 0x26,
	0x24, 0x64, 0xc4, 0xc5, 0x01, 0x73, 0x9f, 0x04, 0x0b, 0xd8, 0x60, 0x31, 0x64, 0x83, 0xc3, 0x20,
	0x72, 0x20, 0x77, 0xc0, 0xd5, 0x39, 0xd9, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3,
	0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x70, 0x49, 0xe4, 0x17, 0xa5, 0x23, 0x6b, 0x4b,
	0x2b, 0x4a, 0xcc, 0x4d, 0x2d, 0xcf, 0x2f, 0xca, 0x76, 0xe2, 0x06, 0x7b, 0x35, 0x00, 0xe4, 0xf3,
	0xe2, 0x00, 0xc6, 0x1f, 0x8c, 0x8c, 0x49, 0x6c, 0xe0, 0x60, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x14, 0x9a, 0xe7, 0xc6, 0x51, 0x01, 0x00, 0x00,
}
