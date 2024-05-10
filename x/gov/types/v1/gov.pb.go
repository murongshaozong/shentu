// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shentu/gov/v1/gov.proto

package v1

import (
	fmt "fmt"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
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

type CustomParams struct {
	CertifierUpdateSecurityVoteTally *v1.TallyParams `protobuf:"bytes,1,opt,name=certifier_update_security_vote_tally,json=certifierUpdateSecurityVoteTally,proto3" json:"certifier_update_security_vote_tally,omitempty"`
	CertifierUpdateStakeVoteTally    *v1.TallyParams `protobuf:"bytes,2,opt,name=certifier_update_stake_vote_tally,json=certifierUpdateStakeVoteTally,proto3" json:"certifier_update_stake_vote_tally,omitempty"`
}

func (m *CustomParams) Reset()         { *m = CustomParams{} }
func (m *CustomParams) String() string { return proto.CompactTextString(m) }
func (*CustomParams) ProtoMessage()    {}
func (*CustomParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_30426ec55f769237, []int{0}
}
func (m *CustomParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CustomParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CustomParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CustomParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomParams.Merge(m, src)
}
func (m *CustomParams) XXX_Size() int {
	return m.Size()
}
func (m *CustomParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomParams.DiscardUnknown(m)
}

var xxx_messageInfo_CustomParams proto.InternalMessageInfo

func (m *CustomParams) GetCertifierUpdateSecurityVoteTally() *v1.TallyParams {
	if m != nil {
		return m.CertifierUpdateSecurityVoteTally
	}
	return nil
}

func (m *CustomParams) GetCertifierUpdateStakeVoteTally() *v1.TallyParams {
	if m != nil {
		return m.CertifierUpdateStakeVoteTally
	}
	return nil
}

func init() {
	proto.RegisterType((*CustomParams)(nil), "shentu.gov.v1.CustomParams")
}

func init() { proto.RegisterFile("shentu/gov/v1/gov.proto", fileDescriptor_30426ec55f769237) }

var fileDescriptor_30426ec55f769237 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x6b, 0x06, 0x86, 0x00, 0x4b, 0x84, 0x54, 0x14, 0x09, 0xab, 0x20, 0x06, 0xa6, 0x58,
	0x81, 0x37, 0x80, 0x91, 0x05, 0xf1, 0x6f, 0x60, 0x89, 0xdc, 0xc4, 0x4d, 0x0d, 0x4d, 0x2e, 0xb2,
	0xcf, 0x16, 0x79, 0x0b, 0x1e, 0x8b, 0xb1, 0x23, 0x63, 0x95, 0xbc, 0x08, 0xb2, 0x1d, 0x55, 0x88,
	0x0e, 0x4c, 0xb6, 0xee, 0xbb, 0xfb, 0xfd, 0xa4, 0x2f, 0x9a, 0xea, 0xa5, 0x68, 0xd0, 0xb0, 0x0a,
	0x2c, 0xb3, 0x99, 0x7b, 0xd2, 0x56, 0x01, 0x42, 0x7c, 0x14, 0x82, 0xd4, 0x4d, 0x6c, 0x96, 0x1c,
	0x57, 0x50, 0x81, 0x4f, 0x98, 0xfb, 0x85, 0xa5, 0x64, 0x5a, 0x80, 0xae, 0x41, 0xef, 0x5c, 0x9f,
	0x6f, 0x48, 0x74, 0x78, 0x6b, 0x34, 0x42, 0x7d, 0xcf, 0x15, 0xaf, 0x75, 0xfc, 0x16, 0x5d, 0x14,
	0x42, 0xa1, 0x5c, 0x48, 0xa1, 0x72, 0xd3, 0x96, 0x1c, 0x45, 0xae, 0x45, 0x61, 0x94, 0xc4, 0x2e,
	0xb7, 0x80, 0x22, 0x47, 0xbe, 0x5a, 0x75, 0x27, 0x64, 0x46, 0x2e, 0x0f, 0xae, 0x92, 0x34, 0x80,
	0x47, 0x7b, 0xfa, 0xe4, 0xb2, 0x40, 0x7a, 0x98, 0x6d, 0x39, 0xcf, 0x1e, 0xf3, 0x38, 0x52, 0x5e,
	0x00, 0x85, 0xdf, 0x8b, 0xcb, 0xe8, 0x6c, 0xd7, 0x85, 0xfc, 0x5d, 0xfc, 0x16, 0xed, 0xfd, 0x2b,
	0x3a, 0xfd, 0x2b, 0x72, 0x88, 0xad, 0xe5, 0xe6, 0xee, 0xab, 0xa7, 0x64, 0xdd, 0x53, 0xb2, 0xe9,
	0x29, 0xf9, 0x1c, 0xe8, 0x64, 0x3d, 0xd0, 0xc9, 0xf7, 0x40, 0x27, 0xaf, 0x59, 0x25, 0x71, 0x69,
	0xe6, 0x69, 0x01, 0x35, 0x0b, 0x2d, 0x2e, 0xc0, 0x34, 0x25, 0x47, 0x09, 0xcd, 0x38, 0x60, 0x1f,
	0xbe, 0x33, 0xec, 0x5a, 0xa1, 0x99, 0xcd, 0xe6, 0xfb, 0xbe, 0xb6, 0xeb, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x21, 0x4b, 0x85, 0x0e, 0x8f, 0x01, 0x00, 0x00,
}

func (m *CustomParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CustomParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CustomParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CertifierUpdateStakeVoteTally != nil {
		{
			size, err := m.CertifierUpdateStakeVoteTally.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGov(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.CertifierUpdateSecurityVoteTally != nil {
		{
			size, err := m.CertifierUpdateSecurityVoteTally.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGov(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CustomParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CertifierUpdateSecurityVoteTally != nil {
		l = m.CertifierUpdateSecurityVoteTally.Size()
		n += 1 + l + sovGov(uint64(l))
	}
	if m.CertifierUpdateStakeVoteTally != nil {
		l = m.CertifierUpdateStakeVoteTally.Size()
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CustomParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: CustomParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CustomParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertifierUpdateSecurityVoteTally", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CertifierUpdateSecurityVoteTally == nil {
				m.CertifierUpdateSecurityVoteTally = &v1.TallyParams{}
			}
			if err := m.CertifierUpdateSecurityVoteTally.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertifierUpdateStakeVoteTally", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CertifierUpdateStakeVoteTally == nil {
				m.CertifierUpdateStakeVoteTally = &v1.TallyParams{}
			}
			if err := m.CertifierUpdateStakeVoteTally.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
