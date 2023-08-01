// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/type/kyc.proto

package types

import (
	fmt "fmt"
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

// KycDataPayload is the KYC info.
type KycDataPayload struct {
	// ignore is true if KYC validation is not required.
	Ignore bool `protobuf:"varint,1,opt,name=ignore,proto3" json:"ignore,omitempty"`
	// approved represent if KYC validation is approved or not.
	Approved bool `protobuf:"varint,2,opt,name=approved,proto3" json:"approved,omitempty"`
	// id is the id of the KYC user.
	ID string `protobuf:"bytes,3,opt,name=id,proto3" json:"id"`
}

func (m *KycDataPayload) Reset()         { *m = KycDataPayload{} }
func (m *KycDataPayload) String() string { return proto.CompactTextString(m) }
func (*KycDataPayload) ProtoMessage()    {}
func (*KycDataPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_aefa821fa9aaec33, []int{0}
}
func (m *KycDataPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KycDataPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KycDataPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KycDataPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KycDataPayload.Merge(m, src)
}
func (m *KycDataPayload) XXX_Size() int {
	return m.Size()
}
func (m *KycDataPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_KycDataPayload.DiscardUnknown(m)
}

var xxx_messageInfo_KycDataPayload proto.InternalMessageInfo

func (m *KycDataPayload) GetIgnore() bool {
	if m != nil {
		return m.Ignore
	}
	return false
}

func (m *KycDataPayload) GetApproved() bool {
	if m != nil {
		return m.Approved
	}
	return false
}

func (m *KycDataPayload) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*KycDataPayload)(nil), "sgenetwork.sge.type.KycDataPayload")
}

func init() { proto.RegisterFile("sge/type/kyc.proto", fileDescriptor_aefa821fa9aaec33) }

var fileDescriptor_aefa821fa9aaec33 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4e, 0x4f, 0xd5,
	0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0xae, 0x4c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x2e, 0x4e, 0x4f, 0xcd, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2b, 0x4e, 0x4f, 0xd5, 0x03,
	0x49, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0xe5, 0xf5, 0x41, 0x2c, 0x88, 0x52, 0xa5, 0x24,
	0x2e, 0x3e, 0xef, 0xca, 0x64, 0x97, 0xc4, 0x92, 0xc4, 0x80, 0xc4, 0xca, 0x9c, 0xfc, 0xc4, 0x14,
	0x21, 0x31, 0x2e, 0xb6, 0xcc, 0xf4, 0xbc, 0xfc, 0xa2, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x8e,
	0x20, 0x28, 0x4f, 0x48, 0x8a, 0x8b, 0x23, 0xb1, 0xa0, 0xa0, 0x28, 0xbf, 0x2c, 0x35, 0x45, 0x82,
	0x09, 0x2c, 0x03, 0xe7, 0x0b, 0xc9, 0x70, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2b, 0x30, 0x6a, 0x70,
	0x3a, 0xf1, 0x3c, 0xba, 0x27, 0xcf, 0xe4, 0xe9, 0xf2, 0xea, 0x9e, 0x3c, 0x53, 0x66, 0x4a, 0x10,
	0x53, 0x66, 0x8a, 0x93, 0xd5, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x29,
	0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x17, 0xa7, 0xa7, 0xea, 0x42,
	0x1d, 0xad, 0x0f, 0xf3, 0x53, 0x71, 0x12, 0x1b, 0xd8, 0x99, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x34, 0x51, 0xd7, 0x69, 0xe7, 0x00, 0x00, 0x00,
}

func (m *KycDataPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KycDataPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KycDataPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintKyc(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Approved {
		i--
		if m.Approved {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Ignore {
		i--
		if m.Ignore {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintKyc(dAtA []byte, offset int, v uint64) int {
	offset -= sovKyc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KycDataPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ignore {
		n += 2
	}
	if m.Approved {
		n += 2
	}
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovKyc(uint64(l))
	}
	return n
}

func sovKyc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKyc(x uint64) (n int) {
	return sovKyc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KycDataPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKyc
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
			return fmt.Errorf("proto: KycDataPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KycDataPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ignore", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKyc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Ignore = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approved", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKyc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Approved = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKyc
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
				return ErrInvalidLengthKyc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKyc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKyc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKyc
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
func skipKyc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKyc
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
					return 0, ErrIntOverflowKyc
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
					return 0, ErrIntOverflowKyc
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
				return 0, ErrInvalidLengthKyc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKyc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKyc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKyc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKyc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKyc = fmt.Errorf("proto: unexpected end of group")
)
