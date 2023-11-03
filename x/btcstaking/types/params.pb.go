// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/btcstaking/v1/params.proto

package types

import (
	fmt "fmt"
	github_com_babylonchain_babylon_types "github.com/babylonchain/babylon/types"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the module.
type Params struct {
	// jury_pk is the public key of jury
	// the PK follows encoding in BIP-340 spec on Bitcoin
	JuryPk *github_com_babylonchain_babylon_types.BIP340PubKey `protobuf:"bytes,1,opt,name=jury_pk,json=juryPk,proto3,customtype=github.com/babylonchain/babylon/types.BIP340PubKey" json:"jury_pk,omitempty"`
	// slashing address is the address that the slashed BTC goes to
	// the address is in string on Bitcoin
	SlashingAddress string `protobuf:"bytes,2,opt,name=slashing_address,json=slashingAddress,proto3" json:"slashing_address,omitempty"`
	// min_slashing_tx_fee_sat is the minimum amount of tx fee (quantified
	// in Satoshi) needed for the pre-signed slashing tx
	// TODO: change to satoshi per byte?
	MinSlashingTxFeeSat int64 `protobuf:"varint,3,opt,name=min_slashing_tx_fee_sat,json=minSlashingTxFeeSat,proto3" json:"min_slashing_tx_fee_sat,omitempty"`
	// min_commission_rate is the chain-wide minimum commission rate that a validator can charge their delegators
	MinCommissionRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=min_commission_rate,json=minCommissionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_commission_rate"`
	// slashing_rate determines the portion of the staked amount to be slashed,
	// expressed as a decimal (e.g., 0.5 for 50%).
	SlashingRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=slashing_rate,json=slashingRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slashing_rate"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d1392776a3e15b9, []int{0}
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

func (m *Params) GetSlashingAddress() string {
	if m != nil {
		return m.SlashingAddress
	}
	return ""
}

func (m *Params) GetMinSlashingTxFeeSat() int64 {
	if m != nil {
		return m.MinSlashingTxFeeSat
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "babylon.btcstaking.v1.Params")
}

func init() {
	proto.RegisterFile("babylon/btcstaking/v1/params.proto", fileDescriptor_8d1392776a3e15b9)
}

var fileDescriptor_8d1392776a3e15b9 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0x4a, 0x4c, 0xaa,
	0xcc, 0xc9, 0xcf, 0xd3, 0x4f, 0x2a, 0x49, 0x2e, 0x2e, 0x49, 0xcc, 0xce, 0xcc, 0x4b, 0xd7, 0x2f,
	0x33, 0xd4, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x85, 0xaa, 0xd1, 0x43, 0xa8, 0xd1, 0x2b, 0x33, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xab,
	0xd0, 0x07, 0xb1, 0x20, 0x8a, 0xa5, 0x24, 0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xe3, 0x21, 0x12,
	0x10, 0x0e, 0x44, 0x4a, 0xa9, 0x99, 0x99, 0x8b, 0x2d, 0x00, 0x6c, 0xb0, 0x90, 0x3f, 0x17, 0x7b,
	0x56, 0x69, 0x51, 0x65, 0x7c, 0x41, 0xb6, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x8f, 0x93, 0xd9, 0xad,
	0x7b, 0xf2, 0x46, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x50, 0x2b,
	0x93, 0x33, 0x12, 0x33, 0xf3, 0x60, 0x1c, 0xfd, 0x92, 0xca, 0x82, 0xd4, 0x62, 0x3d, 0x27, 0xcf,
	0x00, 0x63, 0x13, 0x83, 0x80, 0xd2, 0x24, 0xef, 0xd4, 0xca, 0x20, 0x36, 0x90, 0x31, 0x01, 0xd9,
	0x42, 0x9a, 0x5c, 0x02, 0xc5, 0x39, 0x89, 0xc5, 0x19, 0x99, 0x79, 0xe9, 0xf1, 0x89, 0x29, 0x29,
	0x45, 0xa9, 0xc5, 0xc5, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xfc, 0x30, 0x71, 0x47, 0x88,
	0xb0, 0x90, 0x09, 0x97, 0x78, 0x6e, 0x66, 0x5e, 0x3c, 0x5c, 0x79, 0x49, 0x45, 0x7c, 0x5a, 0x6a,
	0x6a, 0x7c, 0x71, 0x62, 0x89, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x73, 0x90, 0x70, 0x6e, 0x66, 0x5e,
	0x30, 0x54, 0x36, 0xa4, 0xc2, 0x2d, 0x35, 0x35, 0x38, 0xb1, 0x44, 0x28, 0x8e, 0x0b, 0x24, 0x1c,
	0x9f, 0x9c, 0x9f, 0x9b, 0x9b, 0x59, 0x5c, 0x9c, 0x99, 0x9f, 0x17, 0x5f, 0x94, 0x58, 0x92, 0x2a,
	0xc1, 0x02, 0xb2, 0xc3, 0x49, 0xef, 0xc4, 0x3d, 0x79, 0x86, 0x5b, 0xf7, 0xe4, 0xd5, 0x90, 0x7c,
	0x00, 0xf1, 0x3a, 0x94, 0xd2, 0x2d, 0x4e, 0xc9, 0x86, 0x3a, 0xdf, 0x25, 0x35, 0x39, 0x48, 0x30,
	0x37, 0x33, 0xcf, 0x19, 0x6e, 0x52, 0x50, 0x62, 0x49, 0xaa, 0x50, 0x22, 0x17, 0x2f, 0xdc, 0x45,
	0x60, 0x93, 0x59, 0xc1, 0x26, 0xdb, 0x90, 0x66, 0xf2, 0xa5, 0x2d, 0xba, 0x5c, 0xd0, 0x30, 0x07,
	0xd9, 0xc3, 0x03, 0x33, 0x12, 0x64, 0x85, 0x15, 0xcb, 0x8c, 0x05, 0xf2, 0x0c, 0x4e, 0x3e, 0x27,
	0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c,
	0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x45, 0x30, 0xfc, 0x2b, 0x90, 0x53, 0x09,
	0xd8, 0xce, 0x24, 0x36, 0x70, 0xd4, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x85, 0xbe,
	0x67, 0x48, 0x02, 0x00, 0x00,
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
	{
		size := m.SlashingRate.Size()
		i -= size
		if _, err := m.SlashingRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.MinCommissionRate.Size()
		i -= size
		if _, err := m.MinCommissionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.MinSlashingTxFeeSat != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinSlashingTxFeeSat))
		i--
		dAtA[i] = 0x18
	}
	if len(m.SlashingAddress) > 0 {
		i -= len(m.SlashingAddress)
		copy(dAtA[i:], m.SlashingAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.SlashingAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.JuryPk != nil {
		{
			size := m.JuryPk.Size()
			i -= size
			if _, err := m.JuryPk.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
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
	if m.JuryPk != nil {
		l = m.JuryPk.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.SlashingAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.MinSlashingTxFeeSat != 0 {
		n += 1 + sovParams(uint64(m.MinSlashingTxFeeSat))
	}
	l = m.MinCommissionRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SlashingRate.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
				return fmt.Errorf("proto: wrong wireType = %d for field JuryPk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.BIP340PubKey
			m.JuryPk = &v
			if err := m.JuryPk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashingAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SlashingAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSlashingTxFeeSat", wireType)
			}
			m.MinSlashingTxFeeSat = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinSlashingTxFeeSat |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCommissionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinCommissionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashingRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashingRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
