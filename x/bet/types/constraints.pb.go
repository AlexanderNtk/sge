// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/bet/constraints.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Constraints is the bet constrains type for the bets
type Constraints struct {
	// min_amount is the minimum allowed bet amount.
	MinAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=min_amount,json=minAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_amount"`
	// fee is the fee that the bettor needs to pay to bet.
	Fee github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=fee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"fee"`
}

func (m *Constraints) Reset()         { *m = Constraints{} }
func (m *Constraints) String() string { return proto.CompactTextString(m) }
func (*Constraints) ProtoMessage()    {}
func (*Constraints) Descriptor() ([]byte, []int) {
	return fileDescriptor_902ec3e683a9aac6, []int{0}
}
func (m *Constraints) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Constraints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Constraints.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Constraints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Constraints.Merge(m, src)
}
func (m *Constraints) XXX_Size() int {
	return m.Size()
}
func (m *Constraints) XXX_DiscardUnknown() {
	xxx_messageInfo_Constraints.DiscardUnknown(m)
}

var xxx_messageInfo_Constraints proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Constraints)(nil), "sgenetwork.sge.bet.Constraints")
}

func init() { proto.RegisterFile("sge/bet/constraints.proto", fileDescriptor_902ec3e683a9aac6) }

var fileDescriptor_902ec3e683a9aac6 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x4e, 0x4f, 0xd5,
	0x4f, 0x4a, 0x2d, 0xd1, 0x4f, 0xce, 0xcf, 0x2b, 0x2e, 0x29, 0x4a, 0xcc, 0xcc, 0x2b, 0x29, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2a, 0x4e, 0x4f, 0xcd, 0x4b, 0x2d, 0x29, 0xcf, 0x2f,
	0xca, 0xd6, 0x2b, 0x4e, 0x4f, 0xd5, 0x4b, 0x4a, 0x2d, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07,
	0x4b, 0xeb, 0x83, 0x58, 0x10, 0x95, 0x4a, 0xf3, 0x18, 0xb9, 0xb8, 0x9d, 0x11, 0xfa, 0x85, 0x7c,
	0xb9, 0xb8, 0x72, 0x33, 0xf3, 0xe2, 0x13, 0x73, 0xf3, 0x4b, 0xf3, 0x4a, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x9d, 0xf4, 0x4e, 0xdc, 0x93, 0x67, 0xb8, 0x75, 0x4f, 0x5e, 0x2d, 0x3d, 0xb3, 0x24,
	0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x39, 0xbf, 0x38, 0x37, 0xbf, 0x18, 0x4a, 0xe9,
	0x16, 0xa7, 0x64, 0xeb, 0x97, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0x79, 0xe6, 0x95, 0x04, 0x71, 0xe6,
	0x66, 0xe6, 0x39, 0x82, 0x0d, 0x10, 0x72, 0xe0, 0x62, 0x4e, 0x4b, 0x4d, 0x95, 0x60, 0x22, 0xcb,
	0x1c, 0x90, 0x56, 0x27, 0x87, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x42,
	0x36, 0xa6, 0x38, 0x3d, 0x55, 0x17, 0xea, 0x61, 0x10, 0x5b, 0xbf, 0x02, 0x1c, 0x30, 0x60, 0xa3,
	0x92, 0xd8, 0xc0, 0x3e, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x30, 0xdb, 0xa2, 0x4c, 0x30,
	0x01, 0x00, 0x00,
}

func (m *Constraints) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Constraints) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Constraints) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Fee.Size()
		i -= size
		if _, err := m.Fee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintConstraints(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MinAmount.Size()
		i -= size
		if _, err := m.MinAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintConstraints(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintConstraints(dAtA []byte, offset int, v uint64) int {
	offset -= sovConstraints(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Constraints) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinAmount.Size()
	n += 1 + l + sovConstraints(uint64(l))
	l = m.Fee.Size()
	n += 1 + l + sovConstraints(uint64(l))
	return n
}

func sovConstraints(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConstraints(x uint64) (n int) {
	return sovConstraints(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Constraints) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConstraints
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
			return fmt.Errorf("proto: Constraints: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Constraints: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConstraints
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
				return ErrInvalidLengthConstraints
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConstraints
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConstraints
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
				return ErrInvalidLengthConstraints
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConstraints
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConstraints(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConstraints
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
func skipConstraints(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConstraints
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
					return 0, ErrIntOverflowConstraints
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
					return 0, ErrIntOverflowConstraints
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
				return 0, ErrInvalidLengthConstraints
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConstraints
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConstraints
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConstraints        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConstraints          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConstraints = fmt.Errorf("proto: unexpected end of group")
)
