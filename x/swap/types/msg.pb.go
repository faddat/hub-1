// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/swap/v1/msg.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgSwapRequest struct {
	From     string                                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	TxHash   []byte                                 `protobuf:"bytes,2,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	Receiver string                                 `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Amount   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
}

func (m *MsgSwapRequest) Reset()         { *m = MsgSwapRequest{} }
func (m *MsgSwapRequest) String() string { return proto.CompactTextString(m) }
func (*MsgSwapRequest) ProtoMessage()    {}
func (*MsgSwapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2007fffb242cd53c, []int{0}
}
func (m *MsgSwapRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapRequest.Merge(m, src)
}
func (m *MsgSwapRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapRequest proto.InternalMessageInfo

type MsgSwapResponse struct {
}

func (m *MsgSwapResponse) Reset()         { *m = MsgSwapResponse{} }
func (m *MsgSwapResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSwapResponse) ProtoMessage()    {}
func (*MsgSwapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2007fffb242cd53c, []int{1}
}
func (m *MsgSwapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapResponse.Merge(m, src)
}
func (m *MsgSwapResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSwapRequest)(nil), "sentinel.swap.v1.MsgSwapRequest")
	proto.RegisterType((*MsgSwapResponse)(nil), "sentinel.swap.v1.MsgSwapResponse")
}

func init() { proto.RegisterFile("sentinel/swap/v1/msg.proto", fileDescriptor_2007fffb242cd53c) }

var fileDescriptor_2007fffb242cd53c = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xbf, 0x4e, 0xeb, 0x30,
	0x14, 0x87, 0xe3, 0x7b, 0xab, 0x16, 0x2c, 0xc4, 0x1f, 0x0b, 0x89, 0x28, 0x83, 0x5b, 0x3a, 0xa0,
	0x2e, 0xb5, 0x55, 0x78, 0x83, 0x0e, 0x08, 0x06, 0x40, 0x0a, 0x1b, 0x42, 0x42, 0x6e, 0x70, 0x13,
	0x8b, 0x26, 0x0e, 0x39, 0x4e, 0x5a, 0xde, 0x82, 0x57, 0x60, 0xe3, 0x51, 0x3a, 0x76, 0x44, 0x0c,
	0x15, 0xa4, 0x2f, 0x82, 0x92, 0xa6, 0x15, 0x30, 0x30, 0xf9, 0xd8, 0xdf, 0xef, 0x1c, 0xf9, 0xb3,
	0xb1, 0x03, 0x32, 0x32, 0x2a, 0x92, 0x23, 0x0e, 0x63, 0x11, 0xf3, 0xac, 0xc7, 0x43, 0xf0, 0x59,
	0x9c, 0x68, 0xa3, 0xc9, 0xee, 0x8a, 0xb1, 0x82, 0xb1, 0xac, 0xe7, 0xec, 0xfb, 0xda, 0xd7, 0x25,
	0xe4, 0x45, 0xb5, 0xcc, 0xb5, 0x5f, 0x10, 0xde, 0xbe, 0x00, 0xff, 0x7a, 0x2c, 0x62, 0x57, 0x3e,
	0xa6, 0x12, 0x0c, 0x21, 0xb8, 0x36, 0x4c, 0x74, 0x68, 0xa3, 0x16, 0xea, 0x6c, 0xba, 0x65, 0x4d,
	0x0e, 0x70, 0xc3, 0x4c, 0xee, 0x02, 0x01, 0x81, 0xfd, 0xaf, 0x85, 0x3a, 0x5b, 0x6e, 0xdd, 0x4c,
	0xce, 0x04, 0x04, 0xc4, 0xc1, 0x1b, 0x89, 0xf4, 0xa4, 0xca, 0x64, 0x62, 0xff, 0x2f, 0x1b, 0xd6,
	0x7b, 0x72, 0x8a, 0xeb, 0x22, 0xd4, 0x69, 0x64, 0xec, 0x5a, 0x41, 0xfa, 0x6c, 0x3a, 0x6f, 0x5a,
	0xef, 0xf3, 0xe6, 0x91, 0xaf, 0x4c, 0x90, 0x0e, 0x98, 0xa7, 0x43, 0xee, 0x69, 0x08, 0x35, 0x54,
	0x4b, 0x17, 0xee, 0x1f, 0xb8, 0x79, 0x8a, 0x25, 0xb0, 0xf3, 0xc8, 0xb8, 0x55, 0x77, 0x7b, 0x0f,
	0xef, 0xac, 0xaf, 0x08, 0xb1, 0x8e, 0x40, 0x1e, 0xdf, 0x62, 0x5c, 0x1c, 0xc9, 0x24, 0x53, 0x9e,
	0x24, 0x97, 0xb8, 0x51, 0x05, 0x48, 0x8b, 0xfd, 0x16, 0x67, 0x3f, 0xf5, 0x9c, 0xc3, 0x3f, 0x12,
	0xcb, 0xe9, 0xfd, 0xab, 0xe9, 0x27, 0xb5, 0x5e, 0x73, 0x6a, 0x4d, 0x73, 0x8a, 0x66, 0x39, 0x45,
	0x1f, 0x39, 0x45, 0xcf, 0x0b, 0x6a, 0xcd, 0x16, 0xd4, 0x7a, 0x5b, 0x50, 0xeb, 0xa6, 0xfb, 0x4d,
	0x61, 0x35, 0xae, 0xab, 0x87, 0x43, 0xe5, 0x29, 0x31, 0xe2, 0x41, 0x3a, 0xe0, 0x93, 0xe5, 0xa7,
	0x94, 0x36, 0x83, 0x7a, 0xf9, 0xd8, 0x27, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x97, 0xb9, 0x8b,
	0xc8, 0xb2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	MsgSwap(ctx context.Context, in *MsgSwapRequest, opts ...grpc.CallOption) (*MsgSwapResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) MsgSwap(ctx context.Context, in *MsgSwapRequest, opts ...grpc.CallOption) (*MsgSwapResponse, error) {
	out := new(MsgSwapResponse)
	err := c.cc.Invoke(ctx, "/sentinel.swap.v1.MsgService/MsgSwap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	MsgSwap(context.Context, *MsgSwapRequest) (*MsgSwapResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) MsgSwap(ctx context.Context, req *MsgSwapRequest) (*MsgSwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgSwap not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_MsgSwap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).MsgSwap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.swap.v1.MsgService/MsgSwap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).MsgSwap(ctx, req.(*MsgSwapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sentinel.swap.v1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MsgSwap",
			Handler:    _MsgService_MsgSwap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sentinel/swap/v1/msg.proto",
}

func (m *MsgSwapRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSwapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSwapRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovMsg(uint64(l))
	return n
}

func (m *MsgSwapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSwapRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgSwapRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
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
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgSwapResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgSwapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)
