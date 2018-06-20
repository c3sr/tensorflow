// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: device_properties.proto

package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DeviceProperties struct {
	// Device type (CPU, GPU, ...)
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Vendor (Intel, nvidia, ...)
	Vendor string `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// Model (Haswell, K40, ...)
	Model string `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	// Core Frequency in Mhz
	Frequency int64 `protobuf:"varint,4,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// Number of cores
	NumCores int64 `protobuf:"varint,5,opt,name=num_cores,json=numCores,proto3" json:"num_cores,omitempty"`
	// Version of the tools and libraries used with this device (e.g. gcc 4.9,
	// cudnn 5.1)
	Environment map[string]string `protobuf:"bytes,6,rep,name=environment" json:"environment,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Number of registers per core.
	NumRegisters int64 `protobuf:"varint,7,opt,name=num_registers,json=numRegisters,proto3" json:"num_registers,omitempty"`
	// L1 cache size in bytes
	L1CacheSize int64 `protobuf:"varint,8,opt,name=l1_cache_size,json=l1CacheSize,proto3" json:"l1_cache_size,omitempty"`
	// L2 cache size in bytes
	L2CacheSize int64 `protobuf:"varint,9,opt,name=l2_cache_size,json=l2CacheSize,proto3" json:"l2_cache_size,omitempty"`
	// L3 cache size in bytes
	L3CacheSize int64 `protobuf:"varint,10,opt,name=l3_cache_size,json=l3CacheSize,proto3" json:"l3_cache_size,omitempty"`
	// Shared memory size per multiprocessor in bytes. This field is
	// applicable to GPUs only.
	SharedMemorySizePerMultiprocessor int64 `protobuf:"varint,11,opt,name=shared_memory_size_per_multiprocessor,json=sharedMemorySizePerMultiprocessor,proto3" json:"shared_memory_size_per_multiprocessor,omitempty"`
	// Memory size in bytes
	MemorySize int64 `protobuf:"varint,12,opt,name=memory_size,json=memorySize,proto3" json:"memory_size,omitempty"`
	// Memory bandwidth in KB/s
	Bandwidth int64 `protobuf:"varint,13,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
}

func (m *DeviceProperties) Reset()                    { *m = DeviceProperties{} }
func (m *DeviceProperties) String() string            { return proto.CompactTextString(m) }
func (*DeviceProperties) ProtoMessage()               {}
func (*DeviceProperties) Descriptor() ([]byte, []int) { return fileDescriptorDeviceProperties, []int{0} }

func (m *DeviceProperties) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DeviceProperties) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *DeviceProperties) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *DeviceProperties) GetFrequency() int64 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *DeviceProperties) GetNumCores() int64 {
	if m != nil {
		return m.NumCores
	}
	return 0
}

func (m *DeviceProperties) GetEnvironment() map[string]string {
	if m != nil {
		return m.Environment
	}
	return nil
}

func (m *DeviceProperties) GetNumRegisters() int64 {
	if m != nil {
		return m.NumRegisters
	}
	return 0
}

func (m *DeviceProperties) GetL1CacheSize() int64 {
	if m != nil {
		return m.L1CacheSize
	}
	return 0
}

func (m *DeviceProperties) GetL2CacheSize() int64 {
	if m != nil {
		return m.L2CacheSize
	}
	return 0
}

func (m *DeviceProperties) GetL3CacheSize() int64 {
	if m != nil {
		return m.L3CacheSize
	}
	return 0
}

func (m *DeviceProperties) GetSharedMemorySizePerMultiprocessor() int64 {
	if m != nil {
		return m.SharedMemorySizePerMultiprocessor
	}
	return 0
}

func (m *DeviceProperties) GetMemorySize() int64 {
	if m != nil {
		return m.MemorySize
	}
	return 0
}

func (m *DeviceProperties) GetBandwidth() int64 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func init() {
	proto.RegisterType((*DeviceProperties)(nil), "tensorflow.DeviceProperties")
}
func (m *DeviceProperties) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceProperties) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDeviceProperties(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if len(m.Vendor) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDeviceProperties(dAtA, i, uint64(len(m.Vendor)))
		i += copy(dAtA[i:], m.Vendor)
	}
	if len(m.Model) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDeviceProperties(dAtA, i, uint64(len(m.Model)))
		i += copy(dAtA[i:], m.Model)
	}
	if m.Frequency != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.Frequency))
	}
	if m.NumCores != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.NumCores))
	}
	if len(m.Environment) > 0 {
		for k, _ := range m.Environment {
			dAtA[i] = 0x32
			i++
			v := m.Environment[k]
			mapSize := 1 + len(k) + sovDeviceProperties(uint64(len(k))) + 1 + len(v) + sovDeviceProperties(uint64(len(v)))
			i = encodeVarintDeviceProperties(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintDeviceProperties(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintDeviceProperties(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.NumRegisters != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.NumRegisters))
	}
	if m.L1CacheSize != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.L1CacheSize))
	}
	if m.L2CacheSize != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.L2CacheSize))
	}
	if m.L3CacheSize != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.L3CacheSize))
	}
	if m.SharedMemorySizePerMultiprocessor != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.SharedMemorySizePerMultiprocessor))
	}
	if m.MemorySize != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.MemorySize))
	}
	if m.Bandwidth != 0 {
		dAtA[i] = 0x68
		i++
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.Bandwidth))
	}
	return i, nil
}

func encodeVarintDeviceProperties(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DeviceProperties) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovDeviceProperties(uint64(l))
	}
	l = len(m.Vendor)
	if l > 0 {
		n += 1 + l + sovDeviceProperties(uint64(l))
	}
	l = len(m.Model)
	if l > 0 {
		n += 1 + l + sovDeviceProperties(uint64(l))
	}
	if m.Frequency != 0 {
		n += 1 + sovDeviceProperties(uint64(m.Frequency))
	}
	if m.NumCores != 0 {
		n += 1 + sovDeviceProperties(uint64(m.NumCores))
	}
	if len(m.Environment) > 0 {
		for k, v := range m.Environment {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovDeviceProperties(uint64(len(k))) + 1 + len(v) + sovDeviceProperties(uint64(len(v)))
			n += mapEntrySize + 1 + sovDeviceProperties(uint64(mapEntrySize))
		}
	}
	if m.NumRegisters != 0 {
		n += 1 + sovDeviceProperties(uint64(m.NumRegisters))
	}
	if m.L1CacheSize != 0 {
		n += 1 + sovDeviceProperties(uint64(m.L1CacheSize))
	}
	if m.L2CacheSize != 0 {
		n += 1 + sovDeviceProperties(uint64(m.L2CacheSize))
	}
	if m.L3CacheSize != 0 {
		n += 1 + sovDeviceProperties(uint64(m.L3CacheSize))
	}
	if m.SharedMemorySizePerMultiprocessor != 0 {
		n += 1 + sovDeviceProperties(uint64(m.SharedMemorySizePerMultiprocessor))
	}
	if m.MemorySize != 0 {
		n += 1 + sovDeviceProperties(uint64(m.MemorySize))
	}
	if m.Bandwidth != 0 {
		n += 1 + sovDeviceProperties(uint64(m.Bandwidth))
	}
	return n
}

func sovDeviceProperties(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDeviceProperties(x uint64) (n int) {
	return sovDeviceProperties(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeviceProperties) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceProperties
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
			return fmt.Errorf("proto: DeviceProperties: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceProperties: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
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
				return ErrInvalidLengthDeviceProperties
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vendor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
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
				return ErrInvalidLengthDeviceProperties
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vendor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Model", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
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
				return ErrInvalidLengthDeviceProperties
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Model = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frequency", wireType)
			}
			m.Frequency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Frequency |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumCores", wireType)
			}
			m.NumCores = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumCores |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Environment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
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
				return ErrInvalidLengthDeviceProperties
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Environment == nil {
				m.Environment = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDeviceProperties
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDeviceProperties
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthDeviceProperties
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDeviceProperties
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthDeviceProperties
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipDeviceProperties(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthDeviceProperties
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Environment[mapkey] = mapvalue
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumRegisters", wireType)
			}
			m.NumRegisters = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumRegisters |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field L1CacheSize", wireType)
			}
			m.L1CacheSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.L1CacheSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field L2CacheSize", wireType)
			}
			m.L2CacheSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.L2CacheSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field L3CacheSize", wireType)
			}
			m.L3CacheSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.L3CacheSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharedMemorySizePerMultiprocessor", wireType)
			}
			m.SharedMemorySizePerMultiprocessor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SharedMemorySizePerMultiprocessor |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemorySize", wireType)
			}
			m.MemorySize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemorySize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bandwidth", wireType)
			}
			m.Bandwidth = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Bandwidth |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceProperties(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceProperties
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
func skipDeviceProperties(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeviceProperties
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
					return 0, ErrIntOverflowDeviceProperties
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
					return 0, ErrIntOverflowDeviceProperties
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
				return 0, ErrInvalidLengthDeviceProperties
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDeviceProperties
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
				next, err := skipDeviceProperties(dAtA[start:])
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
	ErrInvalidLengthDeviceProperties = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeviceProperties   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("device_properties.proto", fileDescriptorDeviceProperties) }

var fileDescriptorDeviceProperties = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0xf1, 0xda, 0x95, 0xf5, 0x64, 0x95, 0x2a, 0x0b, 0x0d, 0x0b, 0x50, 0x29, 0x43, 0x48,
	0x95, 0x10, 0x95, 0xd6, 0xde, 0x20, 0x2e, 0xb8, 0xd8, 0xd8, 0xe5, 0x44, 0x55, 0x1e, 0x20, 0xea,
	0x92, 0x33, 0x6a, 0x11, 0xdb, 0xe1, 0xd8, 0xe9, 0x94, 0x3d, 0x05, 0xcf, 0xc4, 0x15, 0x97, 0x3c,
	0x02, 0x2a, 0x2f, 0xc1, 0x25, 0xb2, 0x93, 0x35, 0x59, 0xef, 0xfc, 0xff, 0xf9, 0xce, 0xb1, 0x95,
	0xff, 0x87, 0xa7, 0x29, 0x6e, 0x64, 0x82, 0x71, 0x4e, 0x26, 0x47, 0x72, 0x12, 0xed, 0x34, 0x27,
	0xe3, 0x0c, 0x07, 0x87, 0xda, 0x1a, 0xba, 0xc9, 0xcc, 0xed, 0xe9, 0xcf, 0x2e, 0x0c, 0x3f, 0x05,
	0x6e, 0xb1, 0xc3, 0x38, 0x87, 0xae, 0x2b, 0x73, 0x14, 0x6c, 0xcc, 0x26, 0xfd, 0x65, 0x38, 0xf3,
	0x13, 0xe8, 0x6d, 0x50, 0xa7, 0x86, 0xc4, 0x41, 0x70, 0x6b, 0xc5, 0x9f, 0xc0, 0xa1, 0x32, 0x29,
	0x66, 0xa2, 0x13, 0xec, 0x4a, 0xf0, 0x17, 0xd0, 0xbf, 0x21, 0xfc, 0x5e, 0xa0, 0x4e, 0x4a, 0xd1,
	0x1d, 0xb3, 0x49, 0x67, 0xd9, 0x18, 0xfc, 0x39, 0xf4, 0x75, 0xa1, 0xe2, 0xc4, 0x10, 0x5a, 0x71,
	0x18, 0xbe, 0x1e, 0xe9, 0x42, 0x5d, 0x78, 0xcd, 0x3f, 0x43, 0x84, 0x7a, 0x23, 0xc9, 0x68, 0x85,
	0xda, 0x89, 0xde, 0xb8, 0x33, 0x89, 0x66, 0xef, 0xa6, 0xcd, 0x9b, 0xa7, 0xfb, 0xef, 0x9d, 0x5e,
	0x36, 0xfc, 0xa5, 0x76, 0x54, 0x2e, 0xdb, 0x1b, 0xf8, 0x6b, 0x18, 0xf8, 0xdb, 0x08, 0xbf, 0x4a,
	0xeb, 0x90, 0xac, 0x78, 0x1c, 0x6e, 0x3c, 0xd6, 0x85, 0x5a, 0xde, 0x7b, 0xfc, 0x14, 0x06, 0xd9,
	0x59, 0x9c, 0xac, 0x92, 0x35, 0xc6, 0x56, 0xde, 0xa1, 0x38, 0x0a, 0x50, 0x94, 0x9d, 0x5d, 0x78,
	0xef, 0x8b, 0xbc, 0xc3, 0xc0, 0xcc, 0xda, 0x4c, 0xbf, 0x66, 0x66, 0x0f, 0x99, 0x79, 0x9b, 0x81,
	0x9a, 0x99, 0x37, 0xcc, 0x02, 0xde, 0xd8, 0xf5, 0x8a, 0x30, 0x8d, 0x15, 0x2a, 0x43, 0x65, 0x00,
	0xe3, 0x1c, 0x29, 0x56, 0x45, 0xe6, 0x64, 0x4e, 0x26, 0x41, 0x6b, 0x0d, 0x89, 0x28, 0xcc, 0xbe,
	0xaa, 0xe0, 0xab, 0xc0, 0xfa, 0x05, 0x0b, 0xa4, 0xab, 0x07, 0x20, 0x7f, 0x09, 0x51, 0x6b, 0x95,
	0x38, 0x0e, 0x73, 0xa0, 0x76, 0x13, 0x3e, 0x8f, 0xeb, 0x95, 0x4e, 0x6f, 0x65, 0xea, 0xd6, 0x62,
	0x50, 0xe5, 0xb1, 0x33, 0x9e, 0x7d, 0x84, 0xe1, 0xfe, 0x2f, 0xe4, 0x43, 0xe8, 0x7c, 0xc3, 0xb2,
	0xae, 0x80, 0x3f, 0xfa, 0xa4, 0x37, 0xab, 0xac, 0xc0, 0xba, 0x00, 0x95, 0xf8, 0x70, 0xf0, 0x9e,
	0x9d, 0xbf, 0xfd, 0xb5, 0x1d, 0xb1, 0xdf, 0xdb, 0x11, 0xfb, 0xb3, 0x1d, 0xb1, 0x1f, 0x7f, 0x47,
	0x8f, 0xce, 0x4f, 0xf6, 0x33, 0x5a, 0xf8, 0xe6, 0xd9, 0x7f, 0x8c, 0x5d, 0xf7, 0x42, 0x09, 0xe7,
	0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xe1, 0xa7, 0x4d, 0x9f, 0x02, 0x00, 0x00,
}
