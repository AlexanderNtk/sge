// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/bet/ticket.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/sge-network/sge/types"
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

// WagerTicketPayload indicates data of bet placement ticket.
type WagerTicketPayload struct {
	// selected_odds is the user-selected odds to place bet.
	SelectedOdds *BetOdds `protobuf:"bytes,1,opt,name=selected_odds,json=selectedOdds,proto3" json:"selected_odds,omitempty"`
	// kyc_data contains the details of user kyc.
	KycData types.KycDataPayload `protobuf:"bytes,2,opt,name=kyc_data,json=kycData,proto3" json:"kyc_data"`
	// odds_type is the type of odds that are going to be placed
	// such as decimal, fraction, moneyline.
	OddsType OddsType `protobuf:"varint,3,opt,name=odds_type,json=oddsType,proto3,enum=sgenetwork.sge.bet.OddsType" json:"odds_type,omitempty"`
}

func (m *WagerTicketPayload) Reset()         { *m = WagerTicketPayload{} }
func (m *WagerTicketPayload) String() string { return proto.CompactTextString(m) }
func (*WagerTicketPayload) ProtoMessage()    {}
func (*WagerTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6959e7db451613, []int{0}
}
func (m *WagerTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WagerTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WagerTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WagerTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WagerTicketPayload.Merge(m, src)
}
func (m *WagerTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *WagerTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_WagerTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_WagerTicketPayload proto.InternalMessageInfo

func (m *WagerTicketPayload) GetSelectedOdds() *BetOdds {
	if m != nil {
		return m.SelectedOdds
	}
	return nil
}

func (m *WagerTicketPayload) GetKycData() types.KycDataPayload {
	if m != nil {
		return m.KycData
	}
	return types.KycDataPayload{}
}

func (m *WagerTicketPayload) GetOddsType() OddsType {
	if m != nil {
		return m.OddsType
	}
	return OddsType_ODDS_TYPE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*WagerTicketPayload)(nil), "sgenetwork.sge.bet.WagerTicketPayload")
}

func init() { proto.RegisterFile("sge/bet/ticket.proto", fileDescriptor_cf6959e7db451613) }

var fileDescriptor_cf6959e7db451613 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x37, 0x2a, 0x5a, 0xd7, 0x3f, 0x87, 0x50, 0xb4, 0x54, 0x89, 0x45, 0x41, 0x7a, 0x31,
	0x81, 0x7a, 0xf2, 0x56, 0x4a, 0x6f, 0x1e, 0x94, 0x52, 0x10, 0xbc, 0x94, 0x6c, 0x32, 0xc4, 0xb2,
	0xd5, 0x2c, 0xcd, 0x88, 0xe6, 0x5b, 0xf8, 0xb1, 0x7a, 0xec, 0x4d, 0x4f, 0x22, 0xed, 0x17, 0x91,
	0x64, 0xb7, 0x1e, 0xd4, 0xdb, 0xdb, 0x37, 0xf3, 0xde, 0xfc, 0xd8, 0xa4, 0x75, 0x67, 0x40, 0x64,
	0x80, 0x02, 0xc7, 0x2a, 0x07, 0xe4, 0xc5, 0xd4, 0xa2, 0xa5, 0xd4, 0x19, 0x78, 0x02, 0x7c, 0xb1,
	0xd3, 0x9c, 0x3b, 0x03, 0x3c, 0x03, 0x6c, 0xd6, 0x8d, 0x35, 0x36, 0x8e, 0x45, 0x50, 0xe5, 0x66,
	0x33, 0x6c, 0x0a, 0xf4, 0x05, 0x88, 0xdc, 0xab, 0xca, 0x3b, 0x58, 0x75, 0x66, 0x80, 0x23, 0xab,
	0xb5, 0xab, 0xfc, 0xc3, 0x95, 0x1f, 0xbc, 0x51, 0x08, 0x95, 0x83, 0xd3, 0x77, 0x92, 0xd2, 0x3b,
	0x69, 0x60, 0x3a, 0x8c, 0x10, 0xb7, 0xd2, 0x4f, 0xac, 0xd4, 0xb4, 0x9b, 0xee, 0x39, 0x98, 0x80,
	0x42, 0xd0, 0xb1, 0xa6, 0x41, 0x5a, 0xa4, 0xbd, 0xd3, 0x39, 0xe2, 0x7f, 0xe9, 0x78, 0x0f, 0xf0,
	0x46, 0x6b, 0x37, 0xd8, 0x5d, 0x25, 0xc2, 0x17, 0xed, 0xa7, 0xb5, 0xdc, 0xab, 0x91, 0x96, 0x28,
	0x1b, 0x6b, 0x31, 0x7c, 0xf6, 0x3b, 0x1c, 0x31, 0xae, 0xbd, 0xea, 0x4b, 0x94, 0xd5, 0xe1, 0xde,
	0xc6, 0xec, 0xf3, 0x24, 0x19, 0x6c, 0xe5, 0xa5, 0x4b, 0xaf, 0xd2, 0xed, 0x1f, 0xe2, 0xc6, 0x7a,
	0x8b, 0xb4, 0xf7, 0x3b, 0xc7, 0xff, 0x31, 0x84, 0x93, 0x43, 0x5f, 0xc0, 0xa0, 0x66, 0x2b, 0xd5,
	0xeb, 0xce, 0x16, 0x8c, 0xcc, 0x17, 0x8c, 0x7c, 0x2d, 0x18, 0x79, 0x5b, 0xb2, 0x64, 0xbe, 0x64,
	0xc9, 0xc7, 0x92, 0x25, 0xf7, 0xe7, 0x66, 0x8c, 0x0f, 0xcf, 0x19, 0x57, 0xf6, 0x51, 0x38, 0x03,
	0x17, 0x55, 0x59, 0xd0, 0xe2, 0xb5, 0x7c, 0x11, 0x5f, 0x80, 0xcb, 0x36, 0xe3, 0x2f, 0xba, 0xfc,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x55, 0x90, 0x94, 0x38, 0xa9, 0x01, 0x00, 0x00,
}

func (m *WagerTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WagerTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WagerTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OddsType != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.OddsType))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.KycData.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.SelectedOdds != nil {
		{
			size, err := m.SelectedOdds.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTicket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTicket(dAtA []byte, offset int, v uint64) int {
	offset -= sovTicket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WagerTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SelectedOdds != nil {
		l = m.SelectedOdds.Size()
		n += 1 + l + sovTicket(uint64(l))
	}
	l = m.KycData.Size()
	n += 1 + l + sovTicket(uint64(l))
	if m.OddsType != 0 {
		n += 1 + sovTicket(uint64(m.OddsType))
	}
	return n
}

func sovTicket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTicket(x uint64) (n int) {
	return sovTicket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WagerTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: WagerTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WagerTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelectedOdds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SelectedOdds == nil {
				m.SelectedOdds = &BetOdds{}
			}
			if err := m.SelectedOdds.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KycData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.KycData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OddsType", wireType)
			}
			m.OddsType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OddsType |= OddsType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func skipTicket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
				return 0, ErrInvalidLengthTicket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTicket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTicket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTicket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTicket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTicket = fmt.Errorf("proto: unexpected end of group")
)
