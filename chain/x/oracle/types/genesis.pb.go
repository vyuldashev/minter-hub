// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/v1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Params struct {
	SignedClaimsWindow            uint64                                 `protobuf:"varint,1,opt,name=signed_claims_window,json=signedClaimsWindow,proto3" json:"signed_claims_window,omitempty"`
	SlashFractionClaim            github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=slash_fraction_claim,json=slashFractionClaim,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_fraction_claim"`
	SlashFractionConflictingClaim github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=slash_fraction_conflicting_claim,json=slashFractionConflictingClaim,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_fraction_conflicting_claim"`
	Coins                         []*Coin                                `protobuf:"bytes,4,rep,name=coins,proto3" json:"coins,omitempty"`
	MinSingleWithdrawGas          uint64                                 `protobuf:"varint,5,opt,name=min_single_withdraw_gas,json=minSingleWithdrawGas,proto3" json:"min_single_withdraw_gas,omitempty"`
	MinBatchGas                   uint64                                 `protobuf:"varint,6,opt,name=min_batch_gas,json=minBatchGas,proto3" json:"min_batch_gas,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_14b982a0a6345d1d, []int{0}
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

func (m *Params) GetSignedClaimsWindow() uint64 {
	if m != nil {
		return m.SignedClaimsWindow
	}
	return 0
}

func (m *Params) GetCoins() []*Coin {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *Params) GetMinSingleWithdrawGas() uint64 {
	if m != nil {
		return m.MinSingleWithdrawGas
	}
	return 0
}

func (m *Params) GetMinBatchGas() uint64 {
	if m != nil {
		return m.MinBatchGas
	}
	return 0
}

// GenesisState struct
type GenesisState struct {
	Params       *Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	Attestations []Attestation `protobuf:"bytes,2,rep,name=attestations,proto3" json:"attestations"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_14b982a0a6345d1d, []int{1}
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

func (m *GenesisState) GetAttestations() []Attestation {
	if m != nil {
		return m.Attestations
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "oracle.v1.Params")
	proto.RegisterType((*GenesisState)(nil), "oracle.v1.GenesisState")
}

func init() { proto.RegisterFile("oracle/v1/genesis.proto", fileDescriptor_14b982a0a6345d1d) }

var fileDescriptor_14b982a0a6345d1d = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x93, 0xee, 0x76, 0xc1, 0xd9, 0x15, 0x71, 0x58, 0x6d, 0xa8, 0x98, 0x86, 0x05, 0x65,
	0x3d, 0x98, 0xd8, 0x8a, 0x17, 0x4f, 0xba, 0x15, 0x17, 0x0f, 0x82, 0xa4, 0x42, 0xc1, 0x4b, 0x9c,
	0x9d, 0x9d, 0x26, 0x0f, 0x33, 0x33, 0x4b, 0xde, 0xb4, 0x51, 0xbc, 0xf8, 0x11, 0xfc, 0x02, 0x7e,
	0x9f, 0x1e, 0x7b, 0x14, 0x91, 0x22, 0xbb, 0x5f, 0x44, 0x32, 0x19, 0xdc, 0xad, 0xc7, 0x9e, 0x12,
	0xe6, 0xf7, 0xfe, 0xef, 0x3f, 0xf3, 0x7f, 0x8f, 0xec, 0xe8, 0x8a, 0xf1, 0x52, 0x24, 0x67, 0xfb,
	0x49, 0x2e, 0x94, 0x40, 0xc0, 0x78, 0x51, 0x69, 0xa3, 0xe9, 0x8d, 0x16, 0xc4, 0x67, 0xfb, 0xbb,
	0xc3, 0x5c, 0xe7, 0xda, 0x9e, 0x26, 0xcd, 0x5f, 0x5b, 0xb0, 0x7b, 0x6f, 0xad, 0x64, 0xc6, 0x08,
	0x34, 0xcc, 0x80, 0x56, 0x0e, 0xde, 0x59, 0x43, 0xf3, 0x65, 0x21, 0x5c, 0xd3, 0xd1, 0x8f, 0x0e,
	0xe9, 0xbd, 0x63, 0x15, 0x93, 0x48, 0x9f, 0x90, 0x21, 0x42, 0xae, 0xc4, 0x3c, 0xe3, 0x25, 0x03,
	0x89, 0x59, 0x0d, 0x6a, 0xae, 0xeb, 0xc0, 0x8f, 0xfc, 0x71, 0x37, 0xa5, 0x2d, 0x3b, 0xb4, 0xe8,
	0xd8, 0x12, 0xfa, 0x91, 0x0c, 0xb1, 0x64, 0x58, 0x64, 0x27, 0x15, 0xe3, 0x8d, 0x57, 0xab, 0x0c,
	0xb6, 0x22, 0x7f, 0x3c, 0x98, 0xc4, 0xe7, 0x97, 0x7b, 0xde, 0xaf, 0xcb, 0xbd, 0x87, 0x39, 0x98,
	0xe2, 0x74, 0x16, 0x73, 0x2d, 0x13, 0xae, 0x51, 0x6a, 0x74, 0x9f, 0xc7, 0x38, 0xff, 0xe4, 0x2e,
	0xf3, 0x4a, 0xf0, 0x94, 0xda, 0x5e, 0xaf, 0x5d, 0x2b, 0x6b, 0x44, 0x6b, 0x12, 0xfd, 0xef, 0xa0,
	0xd5, 0x49, 0x09, 0xdc, 0x80, 0xca, 0x9d, 0x5b, 0xe7, 0x5a, 0x6e, 0xf7, 0xaf, 0xba, 0xad, 0xbb,
	0xb6, 0xc6, 0x0f, 0xc8, 0x36, 0xd7, 0xa0, 0x30, 0xe8, 0x46, 0x9d, 0x71, 0xff, 0xe0, 0x56, 0xfc,
	0x2f, 0xfc, 0xf8, 0x50, 0x83, 0x4a, 0x5b, 0x4a, 0x9f, 0x91, 0x1d, 0x09, 0x2a, 0x43, 0x50, 0x79,
	0x29, 0xb2, 0x1a, 0x4c, 0x31, 0xaf, 0x58, 0x9d, 0xe5, 0x0c, 0x83, 0x6d, 0x1b, 0xdb, 0x50, 0x82,
	0x3a, 0xb2, 0xf4, 0xd8, 0xc1, 0x29, 0x43, 0x3a, 0x22, 0x37, 0x1b, 0xd9, 0x8c, 0x19, 0x5e, 0xd8,
	0xe2, 0x9e, 0x2d, 0xee, 0x4b, 0x50, 0x93, 0xe6, 0x6c, 0xca, 0xf0, 0x79, 0xf7, 0xdb, 0xef, 0xc8,
	0x1b, 0x7d, 0x25, 0x83, 0x69, 0xbb, 0x05, 0x47, 0x86, 0x19, 0x41, 0x1f, 0x91, 0xde, 0xc2, 0x8e,
	0xcb, 0x8e, 0xa5, 0x7f, 0x70, 0x7b, 0xe3, 0x62, 0xed, 0x1c, 0x53, 0x57, 0x40, 0x5f, 0x90, 0xc1,
	0xc6, 0x1a, 0x60, 0xb0, 0x65, 0x5f, 0x72, 0x77, 0x43, 0xf0, 0x72, 0x8d, 0x27, 0xdd, 0x26, 0xbf,
	0xf4, 0x8a, 0x62, 0xf2, 0xe6, 0x7c, 0x19, 0xfa, 0x17, 0xcb, 0xd0, 0xff, 0xb3, 0x0c, 0xfd, 0xef,
	0xab, 0xd0, 0xbb, 0x58, 0x85, 0xde, 0xcf, 0x55, 0xe8, 0x7d, 0x48, 0x36, 0x52, 0x7e, 0x0b, 0xca,
	0x88, 0xea, 0xbd, 0x60, 0x32, 0x91, 0xc5, 0xe9, 0x2c, 0xe1, 0x05, 0x03, 0x95, 0x7c, 0x4e, 0xdc,
	0xc2, 0xd9, 0xc8, 0x67, 0x3d, 0xbb, 0x6e, 0x4f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xd5,
	0x17, 0x54, 0xde, 0x02, 0x00, 0x00,
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
	if m.MinBatchGas != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MinBatchGas))
		i--
		dAtA[i] = 0x30
	}
	if m.MinSingleWithdrawGas != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MinSingleWithdrawGas))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size := m.SlashFractionConflictingClaim.Size()
		i -= size
		if _, err := m.SlashFractionConflictingClaim.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.SlashFractionClaim.Size()
		i -= size
		if _, err := m.SlashFractionClaim.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.SignedClaimsWindow != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SignedClaimsWindow))
		i--
		dAtA[i] = 0x8
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
	if len(m.Attestations) > 0 {
		for iNdEx := len(m.Attestations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attestations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
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
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
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
	if m.SignedClaimsWindow != 0 {
		n += 1 + sovGenesis(uint64(m.SignedClaimsWindow))
	}
	l = m.SlashFractionClaim.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.SlashFractionConflictingClaim.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.MinSingleWithdrawGas != 0 {
		n += 1 + sovGenesis(uint64(m.MinSingleWithdrawGas))
	}
	if m.MinBatchGas != 0 {
		n += 1 + sovGenesis(uint64(m.MinBatchGas))
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
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Attestations) > 0 {
		for _, e := range m.Attestations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedClaimsWindow", wireType)
			}
			m.SignedClaimsWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignedClaimsWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionClaim", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFractionClaim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionConflictingClaim", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFractionConflictingClaim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, &Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSingleWithdrawGas", wireType)
			}
			m.MinSingleWithdrawGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinSingleWithdrawGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBatchGas", wireType)
			}
			m.MinBatchGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinBatchGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
				return ErrIntOverflowGenesis
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
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
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
				return fmt.Errorf("proto: wrong wireType = %d for field Attestations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attestations = append(m.Attestations, Attestation{})
			if err := m.Attestations[len(m.Attestations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
