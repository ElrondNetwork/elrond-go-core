// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scheduled.proto

package scheduled

import (
	bytes "bytes"
	fmt "fmt"
	smartContractResult "github.com/ElrondNetwork/elrond-go-core/data/smartContractResult"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SmartContractResults struct {
	TxHandlers []smartContractResult.SmartContractResult `protobuf:"bytes,1,rep,name=TxHandlers,proto3" json:"TxHandlers"`
}

func (m *SmartContractResults) Reset()      { *m = SmartContractResults{} }
func (*SmartContractResults) ProtoMessage() {}
func (*SmartContractResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80076f37bd30c16, []int{0}
}
func (m *SmartContractResults) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SmartContractResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SmartContractResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartContractResults.Merge(m, src)
}
func (m *SmartContractResults) XXX_Size() int {
	return m.Size()
}
func (m *SmartContractResults) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartContractResults.DiscardUnknown(m)
}

var xxx_messageInfo_SmartContractResults proto.InternalMessageInfo

func (m *SmartContractResults) GetTxHandlers() []smartContractResult.SmartContractResult {
	if m != nil {
		return m.TxHandlers
	}
	return nil
}

type ScheduledSCRs struct {
	RootHash [][]byte                        `protobuf:"bytes,1,rep,name=rootHash,proto3" json:"rootHash,omitempty"`
	Scrs     map[int32]*SmartContractResults `protobuf:"bytes,2,rep,name=scrs,proto3" json:"scrs,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ScheduledSCRs) Reset()      { *m = ScheduledSCRs{} }
func (*ScheduledSCRs) ProtoMessage() {}
func (*ScheduledSCRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80076f37bd30c16, []int{1}
}
func (m *ScheduledSCRs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduledSCRs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ScheduledSCRs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduledSCRs.Merge(m, src)
}
func (m *ScheduledSCRs) XXX_Size() int {
	return m.Size()
}
func (m *ScheduledSCRs) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduledSCRs.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduledSCRs proto.InternalMessageInfo

func (m *ScheduledSCRs) GetRootHash() [][]byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

func (m *ScheduledSCRs) GetScrs() map[int32]*SmartContractResults {
	if m != nil {
		return m.Scrs
	}
	return nil
}

func init() {
	proto.RegisterType((*SmartContractResults)(nil), "proto.SmartContractResults")
	proto.RegisterType((*ScheduledSCRs)(nil), "proto.ScheduledSCRs")
	proto.RegisterMapType((map[int32]*SmartContractResults)(nil), "proto.ScheduledSCRs.ScrsEntry")
}

func init() { proto.RegisterFile("scheduled.proto", fileDescriptor_f80076f37bd30c16) }

var fileDescriptor_f80076f37bd30c16 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0x7b, 0xf9, 0xf3, 0xe5, 0x63, 0xd0, 0x68, 0x1a, 0x17, 0x4d, 0x4d, 0xae, 0x84, 0x15,
	0x1b, 0xda, 0x88, 0x1b, 0xe3, 0xca, 0x40, 0x48, 0x58, 0xb9, 0x68, 0x59, 0x18, 0x76, 0xa5, 0x1d,
	0x8b, 0xa1, 0x74, 0xcc, 0xcc, 0x54, 0x65, 0xe7, 0x23, 0xf8, 0x18, 0xbe, 0x80, 0xef, 0xc0, 0x92,
	0x25, 0x2b, 0x23, 0xc3, 0xc6, 0x25, 0x8f, 0x60, 0x9c, 0x06, 0x42, 0x22, 0xae, 0x7a, 0xcf, 0xe9,
	0xf9, 0xdd, 0x33, 0x93, 0x21, 0x47, 0x22, 0x1c, 0xd1, 0x28, 0x4b, 0x68, 0xe4, 0x3c, 0x70, 0x26,
	0x99, 0x59, 0xd6, 0x1f, 0xbb, 0x19, 0xdf, 0xcb, 0x51, 0x36, 0x74, 0x42, 0x36, 0x71, 0x63, 0x16,
	0x33, 0x57, 0xdb, 0xc3, 0xec, 0x4e, 0x2b, 0x2d, 0xf4, 0x94, 0x53, 0xf6, 0x60, 0x27, 0xde, 0x4d,
	0x38, 0x4b, 0xa3, 0x1b, 0x2a, 0x9f, 0x18, 0x1f, 0xbb, 0x54, 0xab, 0x66, 0xcc, 0x9a, 0x21, 0xe3,
	0xd4, 0x8d, 0x02, 0x19, 0xb8, 0x62, 0x12, 0x70, 0xd9, 0x61, 0xa9, 0xe4, 0x41, 0x28, 0x3d, 0x2a,
	0xb2, 0x44, 0xee, 0xf3, 0xf2, 0xdd, 0xf5, 0x5b, 0x72, 0xe2, 0xff, 0xfe, 0x29, 0xcc, 0x6b, 0x42,
	0xfa, 0xcf, 0xbd, 0x20, 0x8d, 0x12, 0xca, 0x85, 0x05, 0xb5, 0x62, 0xa3, 0xda, 0xb2, 0x73, 0xc6,
	0xd9, 0x03, 0xb4, 0x4b, 0xb3, 0x8f, 0x33, 0xc3, 0xdb, 0x61, 0xea, 0xef, 0x40, 0x0e, 0xfd, 0xcd,
	0xfd, 0xfd, 0x8e, 0x27, 0x4c, 0x9b, 0xfc, 0xe7, 0x8c, 0xc9, 0x5e, 0x20, 0x46, 0x7a, 0xe3, 0x81,
	0xb7, 0xd5, 0x66, 0x8b, 0x94, 0x44, 0xc8, 0x85, 0x55, 0xd0, 0x4d, 0xb8, 0x69, 0xda, 0xe5, 0x1d,
	0x3f, 0xe4, 0xa2, 0x9b, 0x4a, 0x3e, 0xf5, 0x74, 0xd6, 0xee, 0x93, 0xca, 0xd6, 0x32, 0x8f, 0x49,
	0x71, 0x4c, 0xa7, 0x16, 0xd4, 0xa0, 0x51, 0xf6, 0x7e, 0x46, 0xf3, 0x9c, 0x94, 0x1f, 0x83, 0x24,
	0xa3, 0x56, 0xa1, 0x06, 0x8d, 0x6a, 0xeb, 0xf4, 0xef, 0xd3, 0x0b, 0x2f, 0x4f, 0x5e, 0x15, 0x2e,
	0xa1, 0xdd, 0x99, 0x2f, 0xd1, 0x58, 0x2c, 0xd1, 0x58, 0x2f, 0x11, 0x5e, 0x14, 0xc2, 0x9b, 0x42,
	0x98, 0x29, 0x84, 0xb9, 0x42, 0x58, 0x28, 0x84, 0x4f, 0x85, 0xf0, 0xa5, 0xd0, 0x58, 0x2b, 0x84,
	0xd7, 0x15, 0x1a, 0xf3, 0x15, 0x1a, 0x8b, 0x15, 0x1a, 0x83, 0xca, 0xf6, 0xb9, 0x87, 0xff, 0x74,
	0xd7, 0xc5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x85, 0x34, 0x09, 0x23, 0x02, 0x02, 0x00, 0x00,
}

func (this *SmartContractResults) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SmartContractResults)
	if !ok {
		that2, ok := that.(SmartContractResults)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.TxHandlers) != len(that1.TxHandlers) {
		return false
	}
	for i := range this.TxHandlers {
		if !this.TxHandlers[i].Equal(&that1.TxHandlers[i]) {
			return false
		}
	}
	return true
}
func (this *ScheduledSCRs) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ScheduledSCRs)
	if !ok {
		that2, ok := that.(ScheduledSCRs)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.RootHash) != len(that1.RootHash) {
		return false
	}
	for i := range this.RootHash {
		if !bytes.Equal(this.RootHash[i], that1.RootHash[i]) {
			return false
		}
	}
	if len(this.Scrs) != len(that1.Scrs) {
		return false
	}
	for i := range this.Scrs {
		if !this.Scrs[i].Equal(that1.Scrs[i]) {
			return false
		}
	}
	return true
}
func (this *SmartContractResults) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&scheduled.SmartContractResults{")
	if this.TxHandlers != nil {
		vs := make([]smartContractResult.SmartContractResult, len(this.TxHandlers))
		for i := range vs {
			vs[i] = this.TxHandlers[i]
		}
		s = append(s, "TxHandlers: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ScheduledSCRs) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&scheduled.ScheduledSCRs{")
	s = append(s, "RootHash: "+fmt.Sprintf("%#v", this.RootHash)+",\n")
	keysForScrs := make([]int32, 0, len(this.Scrs))
	for k, _ := range this.Scrs {
		keysForScrs = append(keysForScrs, k)
	}
	github_com_gogo_protobuf_sortkeys.Int32s(keysForScrs)
	mapStringForScrs := "map[int32]*SmartContractResults{"
	for _, k := range keysForScrs {
		mapStringForScrs += fmt.Sprintf("%#v: %#v,", k, this.Scrs[k])
	}
	mapStringForScrs += "}"
	if this.Scrs != nil {
		s = append(s, "Scrs: "+mapStringForScrs+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringScheduled(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *SmartContractResults) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SmartContractResults) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SmartContractResults) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxHandlers) > 0 {
		for iNdEx := len(m.TxHandlers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TxHandlers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintScheduled(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ScheduledSCRs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduledSCRs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduledSCRs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Scrs) > 0 {
		keysForScrs := make([]int32, 0, len(m.Scrs))
		for k := range m.Scrs {
			keysForScrs = append(keysForScrs, int32(k))
		}
		github_com_gogo_protobuf_sortkeys.Int32s(keysForScrs)
		for iNdEx := len(keysForScrs) - 1; iNdEx >= 0; iNdEx-- {
			v := m.Scrs[int32(keysForScrs[iNdEx])]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintScheduled(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i = encodeVarintScheduled(dAtA, i, uint64(keysForScrs[iNdEx]))
			i--
			dAtA[i] = 0x8
			i = encodeVarintScheduled(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.RootHash) > 0 {
		for iNdEx := len(m.RootHash) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.RootHash[iNdEx])
			copy(dAtA[i:], m.RootHash[iNdEx])
			i = encodeVarintScheduled(dAtA, i, uint64(len(m.RootHash[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintScheduled(dAtA []byte, offset int, v uint64) int {
	offset -= sovScheduled(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SmartContractResults) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TxHandlers) > 0 {
		for _, e := range m.TxHandlers {
			l = e.Size()
			n += 1 + l + sovScheduled(uint64(l))
		}
	}
	return n
}

func (m *ScheduledSCRs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RootHash) > 0 {
		for _, b := range m.RootHash {
			l = len(b)
			n += 1 + l + sovScheduled(uint64(l))
		}
	}
	if len(m.Scrs) > 0 {
		for k, v := range m.Scrs {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovScheduled(uint64(l))
			}
			mapEntrySize := 1 + sovScheduled(uint64(k)) + l
			n += mapEntrySize + 1 + sovScheduled(uint64(mapEntrySize))
		}
	}
	return n
}

func sovScheduled(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozScheduled(x uint64) (n int) {
	return sovScheduled(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SmartContractResults) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForTxHandlers := "[]SmartContractResult{"
	for _, f := range this.TxHandlers {
		repeatedStringForTxHandlers += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForTxHandlers += "}"
	s := strings.Join([]string{`&SmartContractResults{`,
		`TxHandlers:` + repeatedStringForTxHandlers + `,`,
		`}`,
	}, "")
	return s
}
func (this *ScheduledSCRs) String() string {
	if this == nil {
		return "nil"
	}
	keysForScrs := make([]int32, 0, len(this.Scrs))
	for k, _ := range this.Scrs {
		keysForScrs = append(keysForScrs, k)
	}
	github_com_gogo_protobuf_sortkeys.Int32s(keysForScrs)
	mapStringForScrs := "map[int32]*SmartContractResults{"
	for _, k := range keysForScrs {
		mapStringForScrs += fmt.Sprintf("%v: %v,", k, this.Scrs[k])
	}
	mapStringForScrs += "}"
	s := strings.Join([]string{`&ScheduledSCRs{`,
		`RootHash:` + fmt.Sprintf("%v", this.RootHash) + `,`,
		`Scrs:` + mapStringForScrs + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringScheduled(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SmartContractResults) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScheduled
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
			return fmt.Errorf("proto: SmartContractResults: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SmartContractResults: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHandlers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduled
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
				return ErrInvalidLengthScheduled
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthScheduled
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHandlers = append(m.TxHandlers, smartContractResult.SmartContractResult{})
			if err := m.TxHandlers[len(m.TxHandlers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipScheduled(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthScheduled
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthScheduled
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
func (m *ScheduledSCRs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScheduled
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
			return fmt.Errorf("proto: ScheduledSCRs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduledSCRs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduled
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthScheduled
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthScheduled
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RootHash = append(m.RootHash, make([]byte, postIndex-iNdEx))
			copy(m.RootHash[len(m.RootHash)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scrs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduled
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
				return ErrInvalidLengthScheduled
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthScheduled
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Scrs == nil {
				m.Scrs = make(map[int32]*SmartContractResults)
			}
			var mapkey int32
			var mapvalue *SmartContractResults
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowScheduled
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowScheduled
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowScheduled
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthScheduled
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthScheduled
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &SmartContractResults{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipScheduled(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthScheduled
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Scrs[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipScheduled(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthScheduled
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthScheduled
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
func skipScheduled(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowScheduled
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
					return 0, ErrIntOverflowScheduled
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowScheduled
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
				return 0, ErrInvalidLengthScheduled
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupScheduled
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthScheduled
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthScheduled        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowScheduled          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupScheduled = fmt.Errorf("proto: unexpected end of group")
)
