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

// PlacementTicketPayload indicates data of bet placement ticket.
type PlacementTicketPayload struct {
	// selected_odds is the user-selected odds to place bet.
	SelectedOdds *BetOdds `protobuf:"bytes,1,opt,name=selected_odds,json=selectedOdds,proto3" json:"selected_odds,omitempty"`
	// kyc_data contains the details of user kyc.
	KycData types.KycDataPayload `protobuf:"bytes,2,opt,name=kyc_data,json=kycData,proto3" json:"kyc_data"`
	// odds_type is the type of odds that are going to be placed
	// such as decimal, fraction, moneyline.
	OddsType OddsType `protobuf:"varint,3,opt,name=odds_type,json=oddsType,proto3,enum=sgenetwork.sge.bet.OddsType" json:"odds_type,omitempty"`
}

func (m *PlacementTicketPayload) Reset()         { *m = PlacementTicketPayload{} }
func (m *PlacementTicketPayload) String() string { return proto.CompactTextString(m) }
func (*PlacementTicketPayload) ProtoMessage()    {}
func (*PlacementTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6959e7db451613, []int{0}
}
func (m *PlacementTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlacementTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlacementTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlacementTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlacementTicketPayload.Merge(m, src)
}
func (m *PlacementTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *PlacementTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_PlacementTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_PlacementTicketPayload proto.InternalMessageInfo

func (m *PlacementTicketPayload) GetSelectedOdds() *BetOdds {
	if m != nil {
		return m.SelectedOdds
	}
	return nil
}

func (m *PlacementTicketPayload) GetKycData() types.KycDataPayload {
	if m != nil {
		return m.KycData
	}
	return types.KycDataPayload{}
}

func (m *PlacementTicketPayload) GetOddsType() OddsType {
	if m != nil {
		return m.OddsType
	}
	return OddsType_ODDS_TYPE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*PlacementTicketPayload)(nil), "sgenetwork.sge.bet.PlacementTicketPayload")
}

func init() { proto.RegisterFile("sge/bet/ticket.proto", fileDescriptor_cf6959e7db451613) }

var fileDescriptor_cf6959e7db451613 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0x27, 0xdf, 0x27, 0x5a, 0xc7, 0x3f, 0x8b, 0xa1, 0xd4, 0x52, 0x25, 0x16, 0x05, 0xe9,
	0xc6, 0x04, 0xea, 0xca, 0x5d, 0x29, 0xdd, 0xb9, 0xb0, 0x94, 0xae, 0xdc, 0x94, 0x4c, 0x72, 0x89,
	0x65, 0xda, 0x66, 0x68, 0xae, 0x68, 0xde, 0xc2, 0xc7, 0xea, 0xb2, 0x4b, 0x41, 0x10, 0x69, 0x5f,
	0x44, 0x92, 0x99, 0x71, 0xa1, 0xee, 0xce, 0x9c, 0x7b, 0xcf, 0xb9, 0x3f, 0x26, 0x71, 0xdd, 0x6a,
	0xe0, 0x29, 0x20, 0xc7, 0xa9, 0xcc, 0x00, 0x59, 0xbe, 0x34, 0x68, 0x92, 0xc4, 0x6a, 0x58, 0x00,
	0x3e, 0x9b, 0x65, 0xc6, 0xac, 0x06, 0x96, 0x02, 0xb6, 0xea, 0xda, 0x68, 0x13, 0xc6, 0xdc, 0xab,
	0x62, 0xb3, 0xe5, 0x37, 0x39, 0xba, 0x1c, 0x78, 0xe6, 0x64, 0xe9, 0x35, 0xaa, 0xce, 0x14, 0x70,
	0x62, 0x94, 0xb2, 0xa5, 0x7f, 0x52, 0xf9, 0xde, 0x9b, 0xf8, 0x50, 0x31, 0xb8, 0x78, 0x27, 0x71,
	0x63, 0x38, 0x13, 0x12, 0xe6, 0xb0, 0xc0, 0x71, 0x00, 0x19, 0x0a, 0x37, 0x33, 0x42, 0x25, 0xbd,
	0xf8, 0xc8, 0xc2, 0x0c, 0x24, 0x82, 0x0a, 0x55, 0x4d, 0xd2, 0x26, 0x9d, 0x83, 0xee, 0x29, 0xfb,
	0x4d, 0xc8, 0xfa, 0x80, 0xf7, 0x4a, 0xd9, 0xd1, 0x61, 0x95, 0xf0, 0x5f, 0xc9, 0x20, 0xae, 0x65,
	0x4e, 0x4e, 0x94, 0x40, 0xd1, 0xfc, 0x17, 0xc2, 0x97, 0x3f, 0xc3, 0x01, 0xe5, 0xce, 0xc9, 0x81,
	0x40, 0x51, 0x1e, 0xee, 0xef, 0xac, 0x3e, 0xce, 0xa3, 0xd1, 0x5e, 0x56, 0xb8, 0xc9, 0x6d, 0xbc,
	0xff, 0x4d, 0xdd, 0xfc, 0xdf, 0x26, 0x9d, 0xe3, 0xee, 0xd9, 0x5f, 0x0c, 0xfe, 0xe4, 0xd8, 0xe5,
	0x30, 0xaa, 0x99, 0x52, 0xf5, 0x7b, 0xab, 0x0d, 0x25, 0xeb, 0x0d, 0x25, 0x9f, 0x1b, 0x4a, 0x5e,
	0xb7, 0x34, 0x5a, 0x6f, 0x69, 0xf4, 0xb6, 0xa5, 0xd1, 0xc3, 0x95, 0x9e, 0xe2, 0xe3, 0x53, 0xca,
	0xa4, 0x99, 0x73, 0xab, 0xe1, 0xba, 0x2c, 0xf3, 0x9a, 0xbf, 0x14, 0xaf, 0xe2, 0x72, 0xb0, 0xe9,
	0x6e, 0xf8, 0x4d, 0x37, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x0f, 0x0c, 0x5b, 0xad, 0x01,
	0x00, 0x00,
}

func (m *PlacementTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlacementTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlacementTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *PlacementTicketPayload) Size() (n int) {
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
func (m *PlacementTicketPayload) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PlacementTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlacementTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
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
