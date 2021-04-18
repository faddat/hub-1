// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/node/v1/msg.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	types1 "github.com/sentinel-official/hub/types"
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

type MsgRegisterRequest struct {
	From      string                                   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Provider  string                                   `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Price     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=price,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"price"`
	RemoteURL string                                   `protobuf:"bytes,4,opt,name=remote_url,json=remoteUrl,proto3" json:"remote_url,omitempty"`
}

func (m *MsgRegisterRequest) Reset()         { *m = MsgRegisterRequest{} }
func (m *MsgRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterRequest) ProtoMessage()    {}
func (*MsgRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f2b348927ca90b6, []int{0}
}
func (m *MsgRegisterRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterRequest.Merge(m, src)
}
func (m *MsgRegisterRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterRequest proto.InternalMessageInfo

type MsgUpdateRequest struct {
	From      string                                   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Provider  string                                   `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Price     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=price,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"price"`
	RemoteURL string                                   `protobuf:"bytes,4,opt,name=remote_url,json=remoteUrl,proto3" json:"remote_url,omitempty"`
}

func (m *MsgUpdateRequest) Reset()         { *m = MsgUpdateRequest{} }
func (m *MsgUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateRequest) ProtoMessage()    {}
func (*MsgUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f2b348927ca90b6, []int{1}
}
func (m *MsgUpdateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateRequest.Merge(m, src)
}
func (m *MsgUpdateRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateRequest proto.InternalMessageInfo

type MsgSetStatusRequest struct {
	From   string        `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Status types1.Status `protobuf:"varint,2,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty"`
}

func (m *MsgSetStatusRequest) Reset()         { *m = MsgSetStatusRequest{} }
func (m *MsgSetStatusRequest) String() string { return proto.CompactTextString(m) }
func (*MsgSetStatusRequest) ProtoMessage()    {}
func (*MsgSetStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f2b348927ca90b6, []int{2}
}
func (m *MsgSetStatusRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetStatusRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetStatusRequest.Merge(m, src)
}
func (m *MsgSetStatusRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetStatusRequest proto.InternalMessageInfo

type MsgRegisterResponse struct {
}

func (m *MsgRegisterResponse) Reset()         { *m = MsgRegisterResponse{} }
func (m *MsgRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterResponse) ProtoMessage()    {}
func (*MsgRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f2b348927ca90b6, []int{3}
}
func (m *MsgRegisterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterResponse.Merge(m, src)
}
func (m *MsgRegisterResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterResponse proto.InternalMessageInfo

type MsgUpdateResponse struct {
}

func (m *MsgUpdateResponse) Reset()         { *m = MsgUpdateResponse{} }
func (m *MsgUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateResponse) ProtoMessage()    {}
func (*MsgUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f2b348927ca90b6, []int{4}
}
func (m *MsgUpdateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateResponse.Merge(m, src)
}
func (m *MsgUpdateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateResponse proto.InternalMessageInfo

type MsgSetStatusResponse struct {
}

func (m *MsgSetStatusResponse) Reset()         { *m = MsgSetStatusResponse{} }
func (m *MsgSetStatusResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetStatusResponse) ProtoMessage()    {}
func (*MsgSetStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f2b348927ca90b6, []int{5}
}
func (m *MsgSetStatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetStatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetStatusResponse.Merge(m, src)
}
func (m *MsgSetStatusResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetStatusResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegisterRequest)(nil), "sentinel.node.v1.MsgRegisterRequest")
	proto.RegisterType((*MsgUpdateRequest)(nil), "sentinel.node.v1.MsgUpdateRequest")
	proto.RegisterType((*MsgSetStatusRequest)(nil), "sentinel.node.v1.MsgSetStatusRequest")
	proto.RegisterType((*MsgRegisterResponse)(nil), "sentinel.node.v1.MsgRegisterResponse")
	proto.RegisterType((*MsgUpdateResponse)(nil), "sentinel.node.v1.MsgUpdateResponse")
	proto.RegisterType((*MsgSetStatusResponse)(nil), "sentinel.node.v1.MsgSetStatusResponse")
}

func init() { proto.RegisterFile("sentinel/node/v1/msg.proto", fileDescriptor_6f2b348927ca90b6) }

var fileDescriptor_6f2b348927ca90b6 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x54, 0xbf, 0x73, 0xd3, 0x30,
	0x18, 0xb5, 0xdb, 0xd2, 0x23, 0x6a, 0xe1, 0x8a, 0x5a, 0xb8, 0xd4, 0x83, 0xdd, 0x0b, 0x3f, 0x2e,
	0x03, 0x91, 0x70, 0xd8, 0x18, 0xc3, 0xc0, 0x42, 0x16, 0x97, 0x2c, 0xbd, 0xe3, 0x38, 0xdb, 0x51,
	0x5c, 0x41, 0x6c, 0x19, 0x7d, 0xb2, 0x0f, 0xfe, 0x03, 0x46, 0x46, 0x16, 0xee, 0x32, 0xf3, 0x97,
	0x74, 0xec, 0xc8, 0x54, 0x20, 0x59, 0xf8, 0x33, 0x38, 0x4b, 0x4e, 0x48, 0x43, 0xaf, 0xd9, 0x99,
	0x2c, 0x7f, 0xdf, 0xd3, 0xfb, 0xde, 0xb3, 0x9e, 0x8c, 0x1c, 0x60, 0x99, 0xe2, 0x19, 0x1b, 0xd3,
	0x4c, 0x0c, 0x19, 0x2d, 0x7d, 0x9a, 0x42, 0x42, 0x72, 0x29, 0x94, 0xc0, 0x7b, 0xf3, 0x1e, 0xa9,
	0x7a, 0xa4, 0xf4, 0x9d, 0x83, 0x44, 0x24, 0x42, 0x37, 0x69, 0xb5, 0x32, 0x38, 0xc7, 0x8d, 0x05,
	0xa4, 0x02, 0x68, 0x14, 0x42, 0xc5, 0x10, 0x31, 0x15, 0xfa, 0x34, 0x16, 0x3c, 0x9b, 0xf7, 0x17,
	0x33, 0xd4, 0xc7, 0x9c, 0x41, 0x35, 0x04, 0x54, 0xa8, 0x0a, 0x30, 0xfd, 0xd6, 0xcc, 0x46, 0xb8,
	0x0f, 0x49, 0xc0, 0x12, 0x0e, 0x8a, 0xc9, 0x80, 0xbd, 0x2f, 0x18, 0x28, 0x8c, 0xd1, 0xd6, 0x48,
	0x8a, 0xb4, 0x69, 0x1f, 0xd9, 0xed, 0x46, 0xa0, 0xd7, 0xd8, 0x41, 0x37, 0x73, 0x29, 0x4a, 0x3e,
	0x64, 0xb2, 0xb9, 0xa1, 0xeb, 0x8b, 0x77, 0x1c, 0xa2, 0x1b, 0xb9, 0xe4, 0x31, 0x6b, 0x6e, 0x1e,
	0x6d, 0xb6, 0x77, 0xba, 0x87, 0xc4, 0xc8, 0x22, 0x95, 0x2c, 0x52, 0xcb, 0x22, 0xcf, 0x05, 0xcf,
	0x7a, 0x4f, 0xce, 0x2e, 0x3c, 0xeb, 0xdb, 0x0f, 0xaf, 0x9d, 0x70, 0x75, 0x5a, 0x44, 0x24, 0x16,
	0x29, 0xad, 0x3d, 0x98, 0x47, 0x07, 0x86, 0xef, 0x8c, 0x58, 0xbd, 0x01, 0x02, 0xc3, 0x8c, 0x1f,
	0x23, 0x24, 0x59, 0x2a, 0x14, 0x7b, 0x53, 0xc8, 0x71, 0x73, 0xab, 0x12, 0xd0, 0xbb, 0x35, 0xbd,
	0xf0, 0x1a, 0x81, 0xae, 0x0e, 0x82, 0x97, 0x41, 0xc3, 0x00, 0x06, 0x72, 0xfc, 0x6c, 0xf7, 0xd3,
	0xc4, 0xb3, 0xbe, 0x4c, 0x3c, 0xfb, 0xf7, 0xc4, 0xb3, 0x5a, 0xbf, 0x6c, 0xb4, 0xd7, 0x87, 0x64,
	0x90, 0x0f, 0x43, 0xc5, 0xfe, 0x53, 0x8f, 0x6f, 0xd1, 0x7e, 0x1f, 0x92, 0x63, 0xa6, 0x8e, 0xf5,
	0xf9, 0x5e, 0xe7, 0xd2, 0x47, 0xdb, 0x26, 0x04, 0xda, 0xe3, 0xed, 0xee, 0x21, 0x59, 0xa4, 0xcd,
	0x88, 0x2a, 0x7d, 0x52, 0xb3, 0xd4, 0xc0, 0x95, 0x59, 0x77, 0xf5, 0xac, 0xbf, 0xa1, 0x81, 0x5c,
	0x64, 0xc0, 0x5a, 0xfb, 0xe8, 0xce, 0xd2, 0x57, 0xae, 0x8b, 0xf7, 0xd0, 0xc1, 0x65, 0x5d, 0xa6,
	0xde, 0xfd, 0xba, 0x81, 0x90, 0x6e, 0xc8, 0xb2, 0xb2, 0x7e, 0x82, 0x76, 0x96, 0x28, 0xf1, 0x03,
	0xb2, 0x7a, 0x01, 0xc8, 0xbf, 0x31, 0x75, 0x1e, 0xae, 0x41, 0x99, 0x51, 0xf8, 0x15, 0x6a, 0x2c,
	0x74, 0xe1, 0xd6, 0x95, 0x7b, 0x2e, 0x45, 0xc3, 0xb9, 0x7f, 0x2d, 0xa6, 0x66, 0x7d, 0x8d, 0x76,
	0x97, 0x8d, 0xe1, 0xab, 0xc5, 0xac, 0x1e, 0x88, 0xf3, 0x68, 0x1d, 0xcc, 0xd0, 0xf7, 0x5e, 0x9c,
	0x4d, 0x5d, 0xfb, 0x7c, 0xea, 0xda, 0x3f, 0xa7, 0xae, 0xfd, 0x79, 0xe6, 0x5a, 0xe7, 0x33, 0xd7,
	0xfa, 0x3e, 0x73, 0xad, 0x93, 0xce, 0x52, 0xac, 0xe6, 0x5c, 0x1d, 0x31, 0x1a, 0xf1, 0x98, 0x87,
	0x63, 0x7a, 0x5a, 0x44, 0xf4, 0x83, 0xf9, 0xa3, 0xe8, 0xc3, 0x8c, 0xb6, 0xf5, 0x4d, 0x7f, 0xfa,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x53, 0x6c, 0x3c, 0xc2, 0x6f, 0x04, 0x00, 0x00,
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
	MsgRegister(ctx context.Context, in *MsgRegisterRequest, opts ...grpc.CallOption) (*MsgRegisterResponse, error)
	MsgUpdate(ctx context.Context, in *MsgUpdateRequest, opts ...grpc.CallOption) (*MsgUpdateResponse, error)
	MsgSetStatus(ctx context.Context, in *MsgSetStatusRequest, opts ...grpc.CallOption) (*MsgSetStatusResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) MsgRegister(ctx context.Context, in *MsgRegisterRequest, opts ...grpc.CallOption) (*MsgRegisterResponse, error) {
	out := new(MsgRegisterResponse)
	err := c.cc.Invoke(ctx, "/sentinel.node.v1.MsgService/MsgRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) MsgUpdate(ctx context.Context, in *MsgUpdateRequest, opts ...grpc.CallOption) (*MsgUpdateResponse, error) {
	out := new(MsgUpdateResponse)
	err := c.cc.Invoke(ctx, "/sentinel.node.v1.MsgService/MsgUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) MsgSetStatus(ctx context.Context, in *MsgSetStatusRequest, opts ...grpc.CallOption) (*MsgSetStatusResponse, error) {
	out := new(MsgSetStatusResponse)
	err := c.cc.Invoke(ctx, "/sentinel.node.v1.MsgService/MsgSetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	MsgRegister(context.Context, *MsgRegisterRequest) (*MsgRegisterResponse, error)
	MsgUpdate(context.Context, *MsgUpdateRequest) (*MsgUpdateResponse, error)
	MsgSetStatus(context.Context, *MsgSetStatusRequest) (*MsgSetStatusResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) MsgRegister(ctx context.Context, req *MsgRegisterRequest) (*MsgRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgRegister not implemented")
}
func (*UnimplementedMsgServiceServer) MsgUpdate(ctx context.Context, req *MsgUpdateRequest) (*MsgUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgUpdate not implemented")
}
func (*UnimplementedMsgServiceServer) MsgSetStatus(ctx context.Context, req *MsgSetStatusRequest) (*MsgSetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgSetStatus not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_MsgRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).MsgRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.node.v1.MsgService/MsgRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).MsgRegister(ctx, req.(*MsgRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_MsgUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).MsgUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.node.v1.MsgService/MsgUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).MsgUpdate(ctx, req.(*MsgUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_MsgSetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).MsgSetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.node.v1.MsgService/MsgSetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).MsgSetStatus(ctx, req.(*MsgSetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sentinel.node.v1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MsgRegister",
			Handler:    _MsgService_MsgRegister_Handler,
		},
		{
			MethodName: "MsgUpdate",
			Handler:    _MsgService_MsgUpdate_Handler,
		},
		{
			MethodName: "MsgSetStatus",
			Handler:    _MsgService_MsgSetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sentinel/node/v1/msg.proto",
}

func (m *MsgRegisterRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RemoteURL) > 0 {
		i -= len(m.RemoteURL)
		copy(dAtA[i:], m.RemoteURL)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.RemoteURL)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Price) > 0 {
		for iNdEx := len(m.Price) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Price[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMsg(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Provider)))
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

func (m *MsgUpdateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RemoteURL) > 0 {
		i -= len(m.RemoteURL)
		copy(dAtA[i:], m.RemoteURL)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.RemoteURL)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Price) > 0 {
		for iNdEx := len(m.Price) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Price[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMsg(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Provider)))
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

func (m *MsgSetStatusRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetStatusRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetStatusRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
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

func (m *MsgRegisterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSetStatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetStatusResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetStatusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgRegisterRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if len(m.Price) > 0 {
		for _, e := range m.Price {
			l = e.Size()
			n += 1 + l + sovMsg(uint64(l))
		}
	}
	l = len(m.RemoteURL)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *MsgUpdateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if len(m.Price) > 0 {
		for _, e := range m.Price {
			l = e.Size()
			n += 1 + l + sovMsg(uint64(l))
		}
	}
	l = len(m.RemoteURL)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *MsgSetStatusRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovMsg(uint64(m.Status))
	}
	return n
}

func (m *MsgRegisterResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSetStatusResponse) Size() (n int) {
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
func (m *MsgRegisterRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRegisterRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
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
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = append(m.Price, types.Coin{})
			if err := m.Price[len(m.Price)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteURL", wireType)
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
			m.RemoteURL = string(dAtA[iNdEx:postIndex])
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
func (m *MsgUpdateRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
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
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = append(m.Price, types.Coin{})
			if err := m.Price[len(m.Price)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteURL", wireType)
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
			m.RemoteURL = string(dAtA[iNdEx:postIndex])
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
func (m *MsgSetStatusRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetStatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetStatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types1.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *MsgRegisterResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRegisterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgUpdateResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgSetStatusResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
