// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/session/v1/proof.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/duration"
	types "github.com/sentinel-official/hub/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Proof struct {
	Channel      uint64          `protobuf:"varint,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Subscription uint64          `protobuf:"varint,2,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Node         string          `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
	Duration     time.Duration   `protobuf:"bytes,4,opt,name=duration,proto3,stdduration" json:"duration"`
	Bandwidth    types.Bandwidth `protobuf:"bytes,5,opt,name=bandwidth,proto3" json:"bandwidth"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c7e6824be930eaf, []int{0}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return m.Size()
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Proof)(nil), "sentinel.session.v1.Proof")
}

func init() { proto.RegisterFile("sentinel/session/v1/proof.proto", fileDescriptor_5c7e6824be930eaf) }

var fileDescriptor_5c7e6824be930eaf = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x3f, 0x4e, 0xc3, 0x30,
	0x14, 0xc6, 0x63, 0x48, 0xa1, 0x35, 0x4c, 0x86, 0x21, 0x54, 0xc8, 0x2d, 0x9d, 0xba, 0x60, 0x53,
	0x38, 0x00, 0x28, 0xe2, 0x00, 0x28, 0x23, 0x5b, 0xfe, 0x38, 0x89, 0xa5, 0xe0, 0x17, 0xc5, 0x49,
	0x81, 0x5b, 0x30, 0x72, 0x04, 0x8e, 0xd2, 0xb1, 0x23, 0x03, 0xe2, 0x4f, 0x7a, 0x11, 0x14, 0xa7,
	0x09, 0x62, 0x7b, 0xcf, 0xef, 0xf7, 0xbe, 0xef, 0xb3, 0x8d, 0x27, 0x5a, 0xa8, 0x52, 0x2a, 0x91,
	0x71, 0x2d, 0xb4, 0x96, 0xa0, 0xf8, 0x72, 0xc1, 0xf3, 0x02, 0x20, 0x66, 0x79, 0x01, 0x25, 0x90,
	0xa3, 0x0e, 0x60, 0x5b, 0x80, 0x2d, 0x17, 0xe3, 0xe3, 0x04, 0x12, 0x30, 0x73, 0xde, 0x54, 0x2d,
	0x3a, 0xa6, 0x09, 0x40, 0x92, 0x09, 0x6e, 0xba, 0xa0, 0x8a, 0x79, 0x54, 0x15, 0x7e, 0xd9, 0xac,
	0xb4, 0xf3, 0xb3, 0xde, 0xab, 0x7c, 0xce, 0x85, 0x6e, 0x9c, 0x02, 0x5f, 0x45, 0x8f, 0x32, 0x2a,
	0xd3, 0x16, 0x99, 0x7d, 0x20, 0x3c, 0xb8, 0x6b, 0xdc, 0x89, 0x83, 0xf7, 0xc3, 0xd4, 0x57, 0x4a,
	0x64, 0x0e, 0x9a, 0xa2, 0xb9, 0xed, 0x75, 0x2d, 0x99, 0xe1, 0x43, 0x5d, 0x05, 0x3a, 0x2c, 0x64,
	0xde, 0x88, 0x3b, 0x3b, 0x66, 0xfc, 0xef, 0x8c, 0x10, 0x6c, 0x2b, 0x88, 0x84, 0xb3, 0x3b, 0x45,
	0xf3, 0x91, 0x67, 0x6a, 0x72, 0x8d, 0x87, 0x5d, 0x20, 0xc7, 0x9e, 0xa2, 0xf9, 0xc1, 0xe5, 0x09,
	0x6b, 0x13, 0xb3, 0x2e, 0x31, 0xbb, 0xdd, 0x02, 0xee, 0x70, 0xf5, 0x39, 0xb1, 0x5e, 0xbf, 0x26,
	0xc8, 0xeb, 0x97, 0xc8, 0x0d, 0x1e, 0xf5, 0x79, 0x9d, 0x81, 0x51, 0x38, 0x65, 0xfd, 0xf3, 0x98,
	0x3b, 0xb1, 0xe5, 0x82, 0xb9, 0x1d, 0xe3, 0xda, 0x8d, 0x88, 0xf7, 0xb7, 0xe4, 0x7a, 0xab, 0x1f,
	0x6a, 0xbd, 0xd5, 0xd4, 0x5a, 0xd5, 0x14, 0xad, 0x6b, 0x8a, 0xbe, 0x6b, 0x8a, 0x5e, 0x36, 0xd4,
	0x5a, 0x6f, 0xa8, 0xf5, 0xbe, 0xa1, 0xd6, 0xfd, 0x45, 0x22, 0xcb, 0xb4, 0x0a, 0x58, 0x08, 0x0f,
	0xbc, 0x93, 0x3e, 0x87, 0x38, 0x96, 0xa1, 0xf4, 0x33, 0x9e, 0x56, 0x01, 0x7f, 0xea, 0x7f, 0xca,
	0x38, 0x06, 0x7b, 0x26, 0xfc, 0xd5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0xa8, 0xbc, 0x92,
	0xca, 0x01, 0x00, 0x00,
}

func (m *Proof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Bandwidth.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProof(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintProof(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	if len(m.Node) > 0 {
		i -= len(m.Node)
		copy(dAtA[i:], m.Node)
		i = encodeVarintProof(dAtA, i, uint64(len(m.Node)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Subscription != 0 {
		i = encodeVarintProof(dAtA, i, uint64(m.Subscription))
		i--
		dAtA[i] = 0x10
	}
	if m.Channel != 0 {
		i = encodeVarintProof(dAtA, i, uint64(m.Channel))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProof(dAtA []byte, offset int, v uint64) int {
	offset -= sovProof(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Proof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Channel != 0 {
		n += 1 + sovProof(uint64(m.Channel))
	}
	if m.Subscription != 0 {
		n += 1 + sovProof(uint64(m.Subscription))
	}
	l = len(m.Node)
	if l > 0 {
		n += 1 + l + sovProof(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovProof(uint64(l))
	l = m.Bandwidth.Size()
	n += 1 + l + sovProof(uint64(l))
	return n
}

func sovProof(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProof(x uint64) (n int) {
	return sovProof(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Proof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProof
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
			return fmt.Errorf("proto: Proof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			m.Channel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Channel |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscription", wireType)
			}
			m.Subscription = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Subscription |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bandwidth", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bandwidth.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProof
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
func skipProof(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
				return 0, ErrInvalidLengthProof
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProof
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProof
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProof        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProof          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProof = fmt.Errorf("proto: unexpected end of group")
)
