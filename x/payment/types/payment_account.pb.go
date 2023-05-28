// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/payment/payment_account.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// PaymentAccount defines a payment account
type PaymentAccount struct {
	// the address of the payment account
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	// the owner address of the payment account
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// whether the payment account is refundable
	Refundable bool `protobuf:"varint,3,opt,name=refundable,proto3" json:"refundable,omitempty"`
}

func (m *PaymentAccount) Reset()         { *m = PaymentAccount{} }
func (m *PaymentAccount) String() string { return proto.CompactTextString(m) }
func (*PaymentAccount) ProtoMessage()    {}
func (*PaymentAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b1cfac7f45dc467, []int{0}
}
func (m *PaymentAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PaymentAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PaymentAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PaymentAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentAccount.Merge(m, src)
}
func (m *PaymentAccount) XXX_Size() int {
	return m.Size()
}
func (m *PaymentAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentAccount.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentAccount proto.InternalMessageInfo

func (m *PaymentAccount) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *PaymentAccount) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *PaymentAccount) GetRefundable() bool {
	if m != nil {
		return m.Refundable
	}
	return false
}

func init() {
	proto.RegisterType((*PaymentAccount)(nil), "greenfield.payment.PaymentAccount")
}

func init() {
	proto.RegisterFile("greenfield/payment/payment_account.proto", fileDescriptor_9b1cfac7f45dc467)
}

var fileDescriptor_9b1cfac7f45dc467 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0x2f, 0x4a, 0x4d,
	0xcd, 0x4b, 0xcb, 0x4c, 0xcd, 0x49, 0xd1, 0x2f, 0x48, 0xac, 0xcc, 0x4d, 0xcd, 0x2b, 0x81, 0xd1,
	0xf1, 0x89, 0xc9, 0xc9, 0xf9, 0xa5, 0x79, 0x25, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x42,
	0x08, 0x95, 0x7a, 0x50, 0x15, 0x52, 0x92, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0xf1, 0x60, 0x15,
	0xfa, 0x10, 0x0e, 0x44, 0xb9, 0x52, 0x1f, 0x23, 0x17, 0x5f, 0x00, 0x44, 0x99, 0x23, 0xc4, 0x1c,
	0x21, 0x1d, 0x2e, 0x96, 0xc4, 0x94, 0x94, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x4e, 0x27, 0x89,
	0x4b, 0x5b, 0x74, 0x45, 0xa0, 0x5a, 0x1c, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x83, 0x4b, 0x8a,
	0x32, 0xf3, 0xd2, 0x83, 0xc0, 0xaa, 0x84, 0xf4, 0xb8, 0x58, 0xf3, 0xcb, 0xf3, 0x52, 0x8b, 0x24,
	0x98, 0x08, 0x28, 0x87, 0x28, 0x13, 0x92, 0xe3, 0xe2, 0x2a, 0x4a, 0x4d, 0x2b, 0xcd, 0x4b, 0x49,
	0x4c, 0xca, 0x49, 0x95, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x08, 0x42, 0x12, 0x71, 0xf2, 0x3c, 0xf1,
	0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8,
	0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xfd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24,
	0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xa4, 0xbc, 0x24, 0xdd, 0xe4, 0x8c, 0xc4, 0xcc, 0x3c, 0x7d, 0xa4,
	0x80, 0xa9, 0x80, 0x07, 0x4d, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x8b, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x66, 0x74, 0xc1, 0x3d, 0x01, 0x00, 0x00,
}

func (m *PaymentAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PaymentAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PaymentAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Refundable {
		i--
		if m.Refundable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintPaymentAccount(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintPaymentAccount(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPaymentAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovPaymentAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PaymentAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovPaymentAccount(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovPaymentAccount(uint64(l))
	}
	if m.Refundable {
		n += 2
	}
	return n
}

func sovPaymentAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPaymentAccount(x uint64) (n int) {
	return sovPaymentAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PaymentAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPaymentAccount
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
			return fmt.Errorf("proto: PaymentAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PaymentAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaymentAccount
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
				return ErrInvalidLengthPaymentAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPaymentAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaymentAccount
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
				return ErrInvalidLengthPaymentAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPaymentAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Refundable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaymentAccount
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
			m.Refundable = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPaymentAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPaymentAccount
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
func skipPaymentAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPaymentAccount
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
					return 0, ErrIntOverflowPaymentAccount
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
					return 0, ErrIntOverflowPaymentAccount
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
				return 0, ErrInvalidLengthPaymentAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPaymentAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPaymentAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPaymentAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPaymentAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPaymentAccount = fmt.Errorf("proto: unexpected end of group")
)