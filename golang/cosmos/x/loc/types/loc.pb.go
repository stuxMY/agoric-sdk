// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: agoric/loc/loc.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// Params defines the set of parameters for the line-of-credit module.
type Params struct {
	// Gives the loan that can be obtained for the collatoral of one staking token.
	// The staking token is the one defined by the x/staking module.
	// If absent, standalone lines of credit may not be initiated.
	StandaloneRatio *types.DecCoin `protobuf:"bytes,1,opt,name=standalone_ratio,json=standaloneRatio,proto3" json:"standalone_ratio,omitempty" yaml:"standalone_ratio"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_9610db7b09723a56, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetStandaloneRatio() *types.DecCoin {
	if m != nil {
		return m.StandaloneRatio
	}
	return nil
}

// Loc records an active line of credit for an implicit account.
// An Loc consisting of collateral and a loan amount.
type Loc struct {
	// The collateral is liened (preventing transfer) until the loan is repaid.
	Collateral *types.Coin `protobuf:"bytes,1,opt,name=collateral,proto3" json:"collateral,omitempty"`
	// The loan outstanding.
	// It must be repaid to lift the lien on the collateral.
	Loan *types.Coin `protobuf:"bytes,2,opt,name=loan,proto3" json:"loan,omitempty"`
}

func (m *Loc) Reset()         { *m = Loc{} }
func (m *Loc) String() string { return proto.CompactTextString(m) }
func (*Loc) ProtoMessage()    {}
func (*Loc) Descriptor() ([]byte, []int) {
	return fileDescriptor_9610db7b09723a56, []int{1}
}
func (m *Loc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Loc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Loc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Loc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Loc.Merge(m, src)
}
func (m *Loc) XXX_Size() int {
	return m.Size()
}
func (m *Loc) XXX_DiscardUnknown() {
	xxx_messageInfo_Loc.DiscardUnknown(m)
}

var xxx_messageInfo_Loc proto.InternalMessageInfo

func (m *Loc) GetCollateral() *types.Coin {
	if m != nil {
		return m.Collateral
	}
	return nil
}

func (m *Loc) GetLoan() *types.Coin {
	if m != nil {
		return m.Loan
	}
	return nil
}

// Account LoC binds an LoC with an address.
// It is used to export and import the total LoC state.
type AccountLoc struct {
	// address is the address of the LoC borrower.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// loc is the single LoC held by this borrower.
	Loc *Loc `protobuf:"bytes,2,opt,name=loc,proto3" json:"loc,omitempty"`
}

func (m *AccountLoc) Reset()         { *m = AccountLoc{} }
func (m *AccountLoc) String() string { return proto.CompactTextString(m) }
func (*AccountLoc) ProtoMessage()    {}
func (*AccountLoc) Descriptor() ([]byte, []int) {
	return fileDescriptor_9610db7b09723a56, []int{2}
}
func (m *AccountLoc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountLoc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountLoc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountLoc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountLoc.Merge(m, src)
}
func (m *AccountLoc) XXX_Size() int {
	return m.Size()
}
func (m *AccountLoc) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountLoc.DiscardUnknown(m)
}

var xxx_messageInfo_AccountLoc proto.InternalMessageInfo

func (m *AccountLoc) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AccountLoc) GetLoc() *Loc {
	if m != nil {
		return m.Loc
	}
	return nil
}

// GenesisState defines the loc module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	// account_loc defines the active LoCs.
	AccountLoc []*AccountLoc `protobuf:"bytes,2,rep,name=account_loc,json=accountLoc,proto3" json:"account_loc,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_9610db7b09723a56, []int{3}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *GenesisState) GetAccountLoc() []*AccountLoc {
	if m != nil {
		return m.AccountLoc
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "agoric.loc.Params")
	proto.RegisterType((*Loc)(nil), "agoric.loc.Loc")
	proto.RegisterType((*AccountLoc)(nil), "agoric.loc.AccountLoc")
	proto.RegisterType((*GenesisState)(nil), "agoric.loc.GenesisState")
}

func init() { proto.RegisterFile("agoric/loc/loc.proto", fileDescriptor_9610db7b09723a56) }

var fileDescriptor_9610db7b09723a56 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4f, 0x8b, 0xd3, 0x40,
	0x14, 0x4f, 0xec, 0x52, 0xf1, 0x55, 0x58, 0x19, 0x16, 0xad, 0xab, 0xcc, 0x6a, 0x4e, 0x8b, 0xb0,
	0x33, 0x6c, 0x05, 0xc5, 0xde, 0x5a, 0x05, 0x11, 0x7a, 0x90, 0xe8, 0xc9, 0x4b, 0x79, 0x99, 0x0c,
	0x31, 0x38, 0x99, 0x57, 0x32, 0x53, 0xb1, 0xdf, 0xc2, 0xa3, 0x47, 0x8f, 0x7e, 0x14, 0x8f, 0x3d,
	0x7a, 0x12, 0x49, 0xbf, 0x81, 0x9f, 0x40, 0x92, 0x89, 0xb6, 0x88, 0xe0, 0x21, 0x30, 0xf9, 0xcd,
	0xef, 0xcf, 0x9b, 0xc7, 0x0f, 0x4e, 0xb0, 0xa0, 0xba, 0x54, 0xd2, 0x50, 0xf7, 0x89, 0x55, 0x4d,
	0x9e, 0x18, 0x04, 0x54, 0x18, 0x52, 0xa7, 0x27, 0x05, 0x15, 0xd4, 0xc1, 0xb2, 0x3d, 0x05, 0xc6,
	0x29, 0x57, 0xe4, 0x2a, 0x72, 0x32, 0x43, 0xa7, 0xe5, 0xfb, 0xcb, 0x4c, 0x7b, 0xbc, 0x94, 0x8a,
	0x4a, 0x1b, 0xee, 0x13, 0x0f, 0xc3, 0x97, 0x58, 0x63, 0xe5, 0x58, 0x0e, 0x37, 0x9c, 0x47, 0x9b,
	0xa3, 0x21, 0xab, 0x97, 0x35, 0xfa, 0x92, 0xc6, 0xf1, 0xbd, 0xf8, 0x7c, 0x34, 0xb9, 0x2b, 0x82,
	0x89, 0x68, 0x4d, 0x44, 0x6f, 0x22, 0x9e, 0x69, 0xf5, 0x94, 0x4a, 0x3b, 0xbf, 0xf3, 0xf3, 0xfb,
	0xd9, 0xad, 0x0d, 0x56, 0x66, 0x9a, 0xfc, 0xad, 0x4f, 0xd2, 0xe3, 0x3d, 0x94, 0xb6, 0xc8, 0xf4,
	0xe8, 0xd3, 0xe7, 0xb3, 0x28, 0x21, 0x18, 0x2c, 0x48, 0xb1, 0x27, 0x00, 0x8a, 0x8c, 0x41, 0xaf,
	0x6b, 0x34, 0x7d, 0xd8, 0xed, 0x7f, 0x86, 0xb5, 0x49, 0xe9, 0x01, 0x99, 0x5d, 0xc0, 0x91, 0x21,
	0xb4, 0xe3, 0x2b, 0xff, 0x13, 0x75, 0xb4, 0xe4, 0x05, 0xc0, 0x4c, 0x29, 0x5a, 0x5b, 0xdf, 0xe6,
	0x8e, 0xe1, 0x2a, 0xe6, 0x79, 0xad, 0x9d, 0xeb, 0x42, 0xaf, 0xa5, 0xbf, 0x7f, 0xd9, 0x7d, 0x18,
	0x18, 0x52, 0xbd, 0xeb, 0xb1, 0xd8, 0xaf, 0x57, 0x2c, 0x48, 0xa5, 0xed, 0x5d, 0xe2, 0xe0, 0xfa,
	0x73, 0x6d, 0xb5, 0x2b, 0xdd, 0x2b, 0x8f, 0x5e, 0xb3, 0x07, 0x30, 0x5c, 0x75, 0x1b, 0xec, 0x1f,
	0xc0, 0x0e, 0x55, 0x61, 0xb7, 0x69, 0xcf, 0x60, 0x8f, 0x61, 0x84, 0x61, 0x8c, 0x65, 0x88, 0x19,
	0x9c, 0x8f, 0x26, 0x37, 0x0f, 0x05, 0xfb, 0x29, 0x53, 0xc0, 0x3f, 0xe7, 0xf9, 0xeb, 0x2f, 0x0d,
	0x8f, 0xbf, 0x36, 0x3c, 0xde, 0x36, 0x3c, 0xfe, 0xd1, 0xf0, 0xf8, 0xe3, 0x8e, 0x47, 0xdb, 0x1d,
	0x8f, 0xbe, 0xed, 0x78, 0xf4, 0xe6, 0x51, 0x51, 0xfa, 0xb7, 0xeb, 0x4c, 0x28, 0xaa, 0xe4, 0x2c,
	0xf4, 0x24, 0x58, 0x5e, 0xb8, 0xfc, 0x9d, 0x2c, 0xc8, 0xa0, 0x2d, 0x64, 0x5f, 0x84, 0x0f, 0x5d,
	0x85, 0xfc, 0x66, 0xa5, 0x5d, 0x36, 0xec, 0x3a, 0xf0, 0xf0, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xa1, 0xb4, 0x98, 0x92, 0x5d, 0x02, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if !this.StandaloneRatio.Equal(that1.StandaloneRatio) {
		return false
	}
	return true
}
func (this *Loc) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Loc)
	if !ok {
		that2, ok := that.(Loc)
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
	if !this.Collateral.Equal(that1.Collateral) {
		return false
	}
	if !this.Loan.Equal(that1.Loan) {
		return false
	}
	return true
}
func (this *AccountLoc) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccountLoc)
	if !ok {
		that2, ok := that.(AccountLoc)
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
	if this.Address != that1.Address {
		return false
	}
	if !this.Loc.Equal(that1.Loc) {
		return false
	}
	return true
}
func (this *GenesisState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenesisState)
	if !ok {
		that2, ok := that.(GenesisState)
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
	if !this.Params.Equal(that1.Params) {
		return false
	}
	if len(this.AccountLoc) != len(that1.AccountLoc) {
		return false
	}
	for i := range this.AccountLoc {
		if !this.AccountLoc[i].Equal(that1.AccountLoc[i]) {
			return false
		}
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StandaloneRatio != nil {
		{
			size, err := m.StandaloneRatio.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLoc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Loc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Loc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Loc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Loan != nil {
		{
			size, err := m.Loan.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLoc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Collateral != nil {
		{
			size, err := m.Collateral.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLoc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AccountLoc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountLoc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountLoc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Loc != nil {
		{
			size, err := m.Loc.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLoc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintLoc(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AccountLoc) > 0 {
		for iNdEx := len(m.AccountLoc) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccountLoc[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLoc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLoc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLoc(dAtA []byte, offset int, v uint64) int {
	offset -= sovLoc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StandaloneRatio != nil {
		l = m.StandaloneRatio.Size()
		n += 1 + l + sovLoc(uint64(l))
	}
	return n
}

func (m *Loc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Collateral != nil {
		l = m.Collateral.Size()
		n += 1 + l + sovLoc(uint64(l))
	}
	if m.Loan != nil {
		l = m.Loan.Size()
		n += 1 + l + sovLoc(uint64(l))
	}
	return n
}

func (m *AccountLoc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovLoc(uint64(l))
	}
	if m.Loc != nil {
		l = m.Loc.Size()
		n += 1 + l + sovLoc(uint64(l))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovLoc(uint64(l))
	}
	if len(m.AccountLoc) > 0 {
		for _, e := range m.AccountLoc {
			l = e.Size()
			n += 1 + l + sovLoc(uint64(l))
		}
	}
	return n
}

func sovLoc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLoc(x uint64) (n int) {
	return sovLoc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoc
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StandaloneRatio", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoc
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
				return ErrInvalidLengthLoc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLoc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StandaloneRatio == nil {
				m.StandaloneRatio = &types.DecCoin{}
			}
			if err := m.StandaloneRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLoc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLoc
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
func (m *Loc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoc
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
			return fmt.Errorf("proto: Loc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Loc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoc
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
				return ErrInvalidLengthLoc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLoc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Collateral == nil {
				m.Collateral = &types.Coin{}
			}
			if err := m.Collateral.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Loan", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoc
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
				return ErrInvalidLengthLoc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLoc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Loan == nil {
				m.Loan = &types.Coin{}
			}
			if err := m.Loan.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLoc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLoc
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
func (m *AccountLoc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoc
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
			return fmt.Errorf("proto: AccountLoc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountLoc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLoc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Loc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoc
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
				return ErrInvalidLengthLoc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLoc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Loc == nil {
				m.Loc = &Loc{}
			}
			if err := m.Loc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLoc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLoc
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoc
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoc
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
				return ErrInvalidLengthLoc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLoc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountLoc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoc
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
				return ErrInvalidLengthLoc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLoc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountLoc = append(m.AccountLoc, &AccountLoc{})
			if err := m.AccountLoc[len(m.AccountLoc)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLoc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLoc
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
func skipLoc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLoc
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
					return 0, ErrIntOverflowLoc
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
					return 0, ErrIntOverflowLoc
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
				return 0, ErrInvalidLengthLoc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLoc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLoc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLoc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLoc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLoc = fmt.Errorf("proto: unexpected end of group")
)