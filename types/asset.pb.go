// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: asset.proto

/*
	Package types is a generated protocol buffer package.

	It is generated from these files:
		asset.proto
		authentication.proto
		check.proto
		entity.proto
		environment.proto
		event.proto
		metrics.proto

	It has these top-level messages:
		Asset
		Tokens
		CheckRequest
		CheckConfig
		Check
		CheckHistory
		Entity
		System
		Network
		NetworkInterface
		Deregistration
		Environment
		Event
		Metrics
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Asset defines an asset agents install as a dependency for a check.
type Asset struct {
	// Name is the unique identifier for an asset
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// URL is the location of the asset
	URL string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// Sha512 is the SHA-512 checksum of the asset
	Sha512 string `protobuf:"bytes,3,opt,name=sha512,proto3" json:"sha512,omitempty"`
	// Metadata is a set of key value pair associated with the asset
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata" json:"metadata" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Filters are a collection of sensu queries, used by the system to determine
	// if the asset should be installed. If more than one filter is present the
	// queries are joined by the "AND" operator.
	Filters []string `protobuf:"bytes,5,rep,name=filters" json:"filters"`
	// Organization indicates to which org an asset belongs to
	Organization string `protobuf:"bytes,6,opt,name=organization,proto3" json:"organization,omitempty"`
}

func (m *Asset) Reset()                    { *m = Asset{} }
func (m *Asset) String() string            { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()               {}
func (*Asset) Descriptor() ([]byte, []int) { return fileDescriptorAsset, []int{0} }

func (m *Asset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Asset) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *Asset) GetSha512() string {
	if m != nil {
		return m.Sha512
	}
	return ""
}

func (m *Asset) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Asset) GetFilters() []string {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *Asset) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

func init() {
	proto.RegisterType((*Asset)(nil), "sensu.types.Asset")
}
func (this *Asset) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Asset)
	if !ok {
		that2, ok := that.(Asset)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.URL != that1.URL {
		return false
	}
	if this.Sha512 != that1.Sha512 {
		return false
	}
	if len(this.Metadata) != len(that1.Metadata) {
		return false
	}
	for i := range this.Metadata {
		if this.Metadata[i] != that1.Metadata[i] {
			return false
		}
	}
	if len(this.Filters) != len(that1.Filters) {
		return false
	}
	for i := range this.Filters {
		if this.Filters[i] != that1.Filters[i] {
			return false
		}
	}
	if this.Organization != that1.Organization {
		return false
	}
	return true
}
func (m *Asset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Asset) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.URL) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAsset(dAtA, i, uint64(len(m.URL)))
		i += copy(dAtA[i:], m.URL)
	}
	if len(m.Sha512) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Sha512)))
		i += copy(dAtA[i:], m.Sha512)
	}
	if len(m.Metadata) > 0 {
		for k, _ := range m.Metadata {
			dAtA[i] = 0x22
			i++
			v := m.Metadata[k]
			mapSize := 1 + len(k) + sovAsset(uint64(len(k))) + 1 + len(v) + sovAsset(uint64(len(v)))
			i = encodeVarintAsset(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintAsset(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintAsset(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.Filters) > 0 {
		for _, s := range m.Filters {
			dAtA[i] = 0x2a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Organization) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Organization)))
		i += copy(dAtA[i:], m.Organization)
	}
	return i, nil
}

func encodeVarintAsset(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedAsset(r randyAsset, easy bool) *Asset {
	this := &Asset{}
	this.Name = string(randStringAsset(r))
	this.URL = string(randStringAsset(r))
	this.Sha512 = string(randStringAsset(r))
	if r.Intn(10) != 0 {
		v1 := r.Intn(10)
		this.Metadata = make(map[string]string)
		for i := 0; i < v1; i++ {
			this.Metadata[randStringAsset(r)] = randStringAsset(r)
		}
	}
	v2 := r.Intn(10)
	this.Filters = make([]string, v2)
	for i := 0; i < v2; i++ {
		this.Filters[i] = string(randStringAsset(r))
	}
	this.Organization = string(randStringAsset(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyAsset interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneAsset(r randyAsset) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringAsset(r randyAsset) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneAsset(r)
	}
	return string(tmps)
}
func randUnrecognizedAsset(r randyAsset, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldAsset(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldAsset(dAtA []byte, r randyAsset, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateAsset(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateAsset(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateAsset(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateAsset(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateAsset(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateAsset(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateAsset(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Asset) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	l = len(m.URL)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	l = len(m.Sha512)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	if len(m.Metadata) > 0 {
		for k, v := range m.Metadata {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovAsset(uint64(len(k))) + 1 + len(v) + sovAsset(uint64(len(v)))
			n += mapEntrySize + 1 + sovAsset(uint64(mapEntrySize))
		}
	}
	if len(m.Filters) > 0 {
		for _, s := range m.Filters {
			l = len(s)
			n += 1 + l + sovAsset(uint64(l))
		}
	}
	l = len(m.Organization)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	return n
}

func sovAsset(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAsset(x uint64) (n int) {
	return sovAsset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Asset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAsset
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
			return fmt.Errorf("proto: Asset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Asset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sha512", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sha512 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAsset
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
							return ErrIntOverflowAsset
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
						return ErrInvalidLengthAsset
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
							return ErrIntOverflowAsset
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
						return ErrInvalidLengthAsset
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipAsset(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthAsset
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Metadata[mapkey] = mapvalue
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filters = append(m.Filters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Organization", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Organization = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAsset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAsset
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
func skipAsset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAsset
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
					return 0, ErrIntOverflowAsset
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
					return 0, ErrIntOverflowAsset
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
				return 0, ErrInvalidLengthAsset
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAsset
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
				next, err := skipAsset(dAtA[start:])
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
	ErrInvalidLengthAsset = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAsset   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("asset.proto", fileDescriptorAsset) }

var fileDescriptorAsset = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2c, 0x2e, 0x4e,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2e, 0x4e, 0xcd, 0x2b, 0x2e, 0xd5, 0x2b,
	0xa9, 0x2c, 0x48, 0x2d, 0x96, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x49, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30,
	0x0b, 0xa2, 0x57, 0x69, 0x16, 0x13, 0x17, 0xab, 0x23, 0xc8, 0x2c, 0x21, 0x21, 0x2e, 0x96, 0xbc,
	0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x92, 0x8b, 0xb9,
	0xb4, 0x28, 0x47, 0x82, 0x09, 0x24, 0xe4, 0xc4, 0xfe, 0xe8, 0x9e, 0x3c, 0x73, 0x68, 0x90, 0x4f,
	0x10, 0x48, 0x4c, 0x48, 0x8c, 0x8b, 0xad, 0x38, 0x23, 0xd1, 0xd4, 0xd0, 0x48, 0x82, 0x19, 0xac,
	0x01, 0xca, 0x13, 0x72, 0xe2, 0xe2, 0xc8, 0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0x94, 0x60,
	0x51, 0x60, 0xd6, 0xe0, 0x36, 0x52, 0xd0, 0x43, 0x72, 0x9f, 0x1e, 0xd8, 0x32, 0x3d, 0x5f, 0xa8,
	0x12, 0xd7, 0xbc, 0x92, 0xa2, 0x4a, 0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0xe0, 0xfa, 0x84,
	0x54, 0xb9, 0xd8, 0xd3, 0x32, 0x73, 0x4a, 0x52, 0x8b, 0x8a, 0x25, 0x58, 0x15, 0x98, 0x35, 0x38,
	0x9d, 0xb8, 0x5f, 0xdd, 0x93, 0x87, 0x09, 0x05, 0xc1, 0x18, 0x42, 0x4a, 0x5c, 0x3c, 0xf9, 0x45,
	0xe9, 0x89, 0x79, 0x99, 0x55, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x12, 0x6c, 0x60, 0x87, 0xa0, 0x88,
	0x49, 0x59, 0x73, 0xf1, 0xa2, 0xd8, 0x25, 0x24, 0xc0, 0xc5, 0x9c, 0x9d, 0x5a, 0x09, 0xf5, 0x25,
	0x88, 0x29, 0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x0a, 0xf1, 0x66, 0x10, 0x84, 0x63,
	0xc5, 0x64, 0xc1, 0xe8, 0xa4, 0xfc, 0xe3, 0xa1, 0x1c, 0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x3b, 0x1e,
	0xc9, 0x31, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33,
	0x1e, 0xcb, 0x31, 0x44, 0xb1, 0x82, 0x3d, 0x94, 0xc4, 0x06, 0x0e, 0x48, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x54, 0x0d, 0xc6, 0xf2, 0x93, 0x01, 0x00, 0x00,
}
