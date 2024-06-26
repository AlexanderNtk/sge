// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sgenetwork/sge/house/withdraw.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
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

// WithdrawalMode is the enum type for the withdrawal mode.
type WithdrawalMode int32

const (
	// invalid
	WithdrawalMode_WITHDRAWAL_MODE_UNSPECIFIED WithdrawalMode = 0
	// full
	WithdrawalMode_WITHDRAWAL_MODE_FULL WithdrawalMode = 1
	// partial
	WithdrawalMode_WITHDRAWAL_MODE_PARTIAL WithdrawalMode = 2
)

var WithdrawalMode_name = map[int32]string{
	0: "WITHDRAWAL_MODE_UNSPECIFIED",
	1: "WITHDRAWAL_MODE_FULL",
	2: "WITHDRAWAL_MODE_PARTIAL",
}

var WithdrawalMode_value = map[string]int32{
	"WITHDRAWAL_MODE_UNSPECIFIED": 0,
	"WITHDRAWAL_MODE_FULL":        1,
	"WITHDRAWAL_MODE_PARTIAL":     2,
}

func (x WithdrawalMode) String() string {
	return proto.EnumName(WithdrawalMode_name, int32(x))
}

func (WithdrawalMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bc9c262f4256300e, []int{0}
}

// Withdrawal represents the withdrawal against a deposit.
type Withdrawal struct {
	// creator is the bech32-encoded address of the depositor.
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
	// withdrawal is the withdrawal attempt id.
	ID uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id" yaml:"id"`
	// address is the bech32-encoded address of the depositor.
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// market_uid is the uid of market against which the deposit is
	// being made.
	MarketUID string `protobuf:"bytes,4,opt,name=market_uid,proto3" json:"market_uid"`
	// participation_index is the id corresponding to the book participation
	ParticipationIndex uint64 `protobuf:"varint,5,opt,name=participation_index,json=participationIndex,proto3" json:"participation_index,omitempty" yaml:"participation_index"`
	// mode is the withdrawal mode enum value
	Mode WithdrawalMode `protobuf:"varint,6,opt,name=mode,proto3,enum=sgenetwork.sge.house.WithdrawalMode" json:"mode,omitempty" yaml:"mode"`
	// amount is the amount being withdrawn.
	Amount cosmossdk_io_math.Int `protobuf:"bytes,7,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount" yaml:"amount"`
}

func (m *Withdrawal) Reset()      { *m = Withdrawal{} }
func (*Withdrawal) ProtoMessage() {}
func (*Withdrawal) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc9c262f4256300e, []int{0}
}
func (m *Withdrawal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Withdrawal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Withdrawal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Withdrawal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Withdrawal.Merge(m, src)
}
func (m *Withdrawal) XXX_Size() int {
	return m.Size()
}
func (m *Withdrawal) XXX_DiscardUnknown() {
	xxx_messageInfo_Withdrawal.DiscardUnknown(m)
}

var xxx_messageInfo_Withdrawal proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("sgenetwork.sge.house.WithdrawalMode", WithdrawalMode_name, WithdrawalMode_value)
	proto.RegisterType((*Withdrawal)(nil), "sgenetwork.sge.house.Withdrawal")
}

func init() {
	proto.RegisterFile("sgenetwork/sge/house/withdraw.proto", fileDescriptor_bc9c262f4256300e)
}

var fileDescriptor_bc9c262f4256300e = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x6f, 0x12, 0x41,
	0x14, 0xc7, 0x77, 0xb7, 0x48, 0xc3, 0xa8, 0x48, 0x46, 0x8c, 0x6b, 0x49, 0x76, 0xc8, 0xea, 0x01,
	0x8d, 0x2e, 0x89, 0xde, 0xea, 0x09, 0x04, 0xe2, 0x26, 0x60, 0x9b, 0xb5, 0x84, 0xc4, 0x0b, 0x99,
	0x32, 0x93, 0x65, 0x42, 0x97, 0x21, 0x3b, 0x43, 0x68, 0xbf, 0x81, 0x47, 0x8f, 0x1e, 0xf9, 0x2c,
	0x9e, 0x7a, 0xec, 0xd1, 0x78, 0x98, 0x18, 0xb8, 0x18, 0x8f, 0xfb, 0x09, 0x9a, 0x1d, 0xb6, 0x29,
	0x25, 0xbd, 0xbd, 0xbc, 0xf7, 0xfb, 0xbf, 0xff, 0xcc, 0xcb, 0x1f, 0xbc, 0x14, 0x21, 0x9d, 0x52,
	0xb9, 0xe0, 0xf1, 0xa4, 0x2e, 0x42, 0x5a, 0x1f, 0xf3, 0xb9, 0xa0, 0xf5, 0x05, 0x93, 0x63, 0x12,
	0xe3, 0x85, 0x37, 0x8b, 0xb9, 0xe4, 0xb0, 0x7c, 0x0b, 0x79, 0x22, 0xa4, 0x9e, 0x86, 0x0e, 0xca,
	0x21, 0x0f, 0xb9, 0x06, 0xea, 0x69, 0xb5, 0x61, 0xdd, 0x5f, 0x7b, 0x00, 0x0c, 0x32, 0x39, 0x3e,
	0x83, 0x6f, 0xc1, 0xfe, 0x28, 0xa6, 0x58, 0xf2, 0xd8, 0x36, 0xab, 0x66, 0xad, 0xd0, 0x84, 0x89,
	0x42, 0xc5, 0x0b, 0x1c, 0x9d, 0x1d, 0xba, 0xd9, 0xc0, 0x0d, 0x6e, 0x10, 0xf8, 0x1a, 0x58, 0x8c,
	0xd8, 0x56, 0xd5, 0xac, 0xe5, 0x9a, 0x2f, 0x56, 0x0a, 0x59, 0x7e, 0xeb, 0xbf, 0x42, 0x16, 0x23,
	0x89, 0x42, 0x85, 0x8d, 0x88, 0x11, 0x37, 0xb0, 0x18, 0x49, 0x17, 0x63, 0x42, 0x62, 0x2a, 0x84,
	0xbd, 0xb7, 0xbb, 0x38, 0x1b, 0xb8, 0xc1, 0x0d, 0x02, 0x3f, 0x02, 0x10, 0xe1, 0x78, 0x42, 0xe5,
	0x70, 0xce, 0x88, 0x9d, 0xd3, 0x82, 0xca, 0x4a, 0xa1, 0x42, 0x4f, 0x77, 0xfb, 0xda, 0x67, 0x0b,
	0x09, 0xb6, 0x6a, 0x78, 0x04, 0x9e, 0xce, 0x70, 0x2c, 0xd9, 0x88, 0xcd, 0xb0, 0x64, 0x7c, 0x3a,
	0x64, 0x53, 0x42, 0xcf, 0xed, 0x07, 0xfa, 0x99, 0x4e, 0xa2, 0xd0, 0xc1, 0xc6, 0xf6, 0x1e, 0xc8,
	0x0d, 0xe0, 0x9d, 0xae, 0x9f, 0x36, 0xa1, 0x0f, 0x72, 0x11, 0x27, 0xd4, 0xce, 0x57, 0xcd, 0x5a,
	0xf1, 0xfd, 0x2b, 0xef, 0xbe, 0xf3, 0x7a, 0xb7, 0x47, 0xec, 0x71, 0x42, 0x9b, 0x4f, 0x12, 0x85,
	0x1e, 0x6e, 0x7c, 0x52, 0xad, 0x1b, 0xe8, 0x15, 0xb0, 0x03, 0xf2, 0x38, 0xe2, 0xf3, 0xa9, 0xb4,
	0xf7, 0xf5, 0xa7, 0xbc, 0x4b, 0x85, 0x8c, 0x3f, 0x0a, 0x3d, 0x1b, 0x71, 0x11, 0x71, 0x21, 0xc8,
	0xc4, 0x63, 0xbc, 0x1e, 0x61, 0x39, 0xf6, 0xfc, 0xa9, 0x4c, 0x14, 0x7a, 0x9c, 0x9d, 0x48, 0x8b,
	0xdc, 0x20, 0x53, 0x1f, 0x3e, 0xfa, 0xbe, 0x44, 0xc6, 0xcf, 0x25, 0x32, 0xfe, 0x2d, 0x91, 0xf1,
	0x66, 0x0c, 0x8a, 0x77, 0xed, 0x21, 0x02, 0x95, 0x81, 0x7f, 0xf2, 0xb9, 0x15, 0x34, 0x06, 0x8d,
	0xee, 0xb0, 0x77, 0xd4, 0x6a, 0x0f, 0xfb, 0x5f, 0xbe, 0x1e, 0xb7, 0x3f, 0xf9, 0x1d, 0xbf, 0xdd,
	0x2a, 0x19, 0xd0, 0x06, 0xe5, 0x5d, 0xa0, 0xd3, 0xef, 0x76, 0x4b, 0x26, 0xac, 0x80, 0xe7, 0xbb,
	0x93, 0xe3, 0x46, 0x70, 0xe2, 0x37, 0xba, 0x25, 0xab, 0xd9, 0xbc, 0x5c, 0x39, 0xe6, 0xd5, 0xca,
	0x31, 0xff, 0xae, 0x1c, 0xf3, 0xc7, 0xda, 0x31, 0xae, 0xd6, 0x8e, 0xf1, 0x7b, 0xed, 0x18, 0xdf,
	0x6a, 0x21, 0x93, 0xe3, 0xf9, 0xa9, 0x37, 0xe2, 0x51, 0x9a, 0xcc, 0x77, 0xdb, 0x29, 0x3d, 0xcf,
	0x72, 0x2a, 0x2f, 0x66, 0x54, 0x9c, 0xe6, 0x75, 0xf2, 0x3e, 0x5c, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xc7, 0xb9, 0xe4, 0x97, 0xcc, 0x02, 0x00, 0x00,
}

func (m *Withdrawal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Withdrawal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Withdrawal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintWithdraw(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.Mode != 0 {
		i = encodeVarintWithdraw(dAtA, i, uint64(m.Mode))
		i--
		dAtA[i] = 0x30
	}
	if m.ParticipationIndex != 0 {
		i = encodeVarintWithdraw(dAtA, i, uint64(m.ParticipationIndex))
		i--
		dAtA[i] = 0x28
	}
	if len(m.MarketUID) > 0 {
		i -= len(m.MarketUID)
		copy(dAtA[i:], m.MarketUID)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.MarketUID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ID != 0 {
		i = encodeVarintWithdraw(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWithdraw(dAtA []byte, offset int, v uint64) int {
	offset -= sovWithdraw(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Withdrawal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovWithdraw(uint64(m.ID))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	l = len(m.MarketUID)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	if m.ParticipationIndex != 0 {
		n += 1 + sovWithdraw(uint64(m.ParticipationIndex))
	}
	if m.Mode != 0 {
		n += 1 + sovWithdraw(uint64(m.Mode))
	}
	l = m.Amount.Size()
	n += 1 + l + sovWithdraw(uint64(l))
	return n
}

func sovWithdraw(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWithdraw(x uint64) (n int) {
	return sovWithdraw(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Withdrawal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWithdraw
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
			return fmt.Errorf("proto: Withdrawal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Withdrawal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationIndex", wireType)
			}
			m.ParticipationIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParticipationIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
			}
			m.Mode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mode |= WithdrawalMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWithdraw(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWithdraw
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
func skipWithdraw(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWithdraw
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
					return 0, ErrIntOverflowWithdraw
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
					return 0, ErrIntOverflowWithdraw
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
				return 0, ErrInvalidLengthWithdraw
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWithdraw
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWithdraw
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWithdraw        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWithdraw          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWithdraw = fmt.Errorf("proto: unexpected end of group")
)
