// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc2.proto

package rpc2

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The response message
type Rpc2Reply struct {
	Rpc2Msg              string   `protobuf:"bytes,1,opt,name=rpc2_msg,json=rpc2Msg,proto3" json:"rpc2_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rpc2Reply) Reset()         { *m = Rpc2Reply{} }
func (m *Rpc2Reply) String() string { return proto.CompactTextString(m) }
func (*Rpc2Reply) ProtoMessage()    {}
func (*Rpc2Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_af0916bb5e6806d0, []int{0}
}

func (m *Rpc2Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rpc2Reply.Unmarshal(m, b)
}
func (m *Rpc2Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rpc2Reply.Marshal(b, m, deterministic)
}
func (m *Rpc2Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rpc2Reply.Merge(m, src)
}
func (m *Rpc2Reply) XXX_Size() int {
	return xxx_messageInfo_Rpc2Reply.Size(m)
}
func (m *Rpc2Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Rpc2Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Rpc2Reply proto.InternalMessageInfo

func (m *Rpc2Reply) GetRpc2Msg() string {
	if m != nil {
		return m.Rpc2Msg
	}
	return ""
}

//Empty message for request header since currently there are no request params
type RequestParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestParams) Reset()         { *m = RequestParams{} }
func (m *RequestParams) String() string { return proto.CompactTextString(m) }
func (*RequestParams) ProtoMessage()    {}
func (*RequestParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_af0916bb5e6806d0, []int{1}
}

func (m *RequestParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestParams.Unmarshal(m, b)
}
func (m *RequestParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestParams.Marshal(b, m, deterministic)
}
func (m *RequestParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestParams.Merge(m, src)
}
func (m *RequestParams) XXX_Size() int {
	return xxx_messageInfo_RequestParams.Size(m)
}
func (m *RequestParams) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestParams.DiscardUnknown(m)
}

var xxx_messageInfo_RequestParams proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Rpc2Reply)(nil), "rpc2.Rpc2Reply")
	proto.RegisterType((*RequestParams)(nil), "rpc2.RequestParams")
}

func init() { proto.RegisterFile("rpc2.proto", fileDescriptor_af0916bb5e6806d0) }

var fileDescriptor_af0916bb5e6806d0 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2a, 0x48, 0x36,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xd4, 0xb8, 0x38, 0x83, 0x0a,
	0x92, 0x8d, 0x82, 0x52, 0x0b, 0x72, 0x2a, 0x85, 0x24, 0xb9, 0x38, 0x40, 0x82, 0xf1, 0xb9, 0xc5,
	0xe9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xec, 0x20, 0xbe, 0x6f, 0x71, 0xba, 0x12, 0x3f,
	0x17, 0x6f, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x49, 0x40, 0x62, 0x51, 0x62, 0x6e, 0xb1, 0x91,
	0x07, 0x97, 0x40, 0x50, 0x41, 0xb2, 0x73, 0x7e, 0x6e, 0x6e, 0x69, 0x5e, 0x66, 0x72, 0x62, 0x49,
	0x7e, 0x91, 0x91, 0x90, 0x09, 0x17, 0x57, 0x7a, 0x6a, 0x49, 0x10, 0x44, 0x8b, 0x90, 0xb0, 0x1e,
	0xd8, 0x36, 0x14, 0x6d, 0x52, 0xfc, 0x50, 0x41, 0x98, 0x9d, 0x4a, 0x0c, 0x4e, 0x6c, 0x51, 0x60,
	0xa7, 0x24, 0xb1, 0x81, 0xdd, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x70, 0x3b, 0x2d,
	0xa5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RpcCommunicator2Client is the client API for RpcCommunicator2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RpcCommunicator2Client interface {
	GetRpc2Msg(ctx context.Context, in *RequestParams, opts ...grpc.CallOption) (*Rpc2Reply, error)
}

type rpcCommunicator2Client struct {
	cc *grpc.ClientConn
}

func NewRpcCommunicator2Client(cc *grpc.ClientConn) RpcCommunicator2Client {
	return &rpcCommunicator2Client{cc}
}

func (c *rpcCommunicator2Client) GetRpc2Msg(ctx context.Context, in *RequestParams, opts ...grpc.CallOption) (*Rpc2Reply, error) {
	out := new(Rpc2Reply)
	err := c.cc.Invoke(ctx, "/rpc2.RpcCommunicator2/getRpc2Msg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcCommunicator2Server is the server API for RpcCommunicator2 service.
type RpcCommunicator2Server interface {
	GetRpc2Msg(context.Context, *RequestParams) (*Rpc2Reply, error)
}

// UnimplementedRpcCommunicator2Server can be embedded to have forward compatible implementations.
type UnimplementedRpcCommunicator2Server struct {
}

func (*UnimplementedRpcCommunicator2Server) GetRpc2Msg(ctx context.Context, req *RequestParams) (*Rpc2Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRpc2Msg not implemented")
}

func RegisterRpcCommunicator2Server(s *grpc.Server, srv RpcCommunicator2Server) {
	s.RegisterService(&_RpcCommunicator2_serviceDesc, srv)
}

func _RpcCommunicator2_GetRpc2Msg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcCommunicator2Server).GetRpc2Msg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc2.RpcCommunicator2/GetRpc2Msg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcCommunicator2Server).GetRpc2Msg(ctx, req.(*RequestParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _RpcCommunicator2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc2.RpcCommunicator2",
	HandlerType: (*RpcCommunicator2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getRpc2Msg",
			Handler:    _RpcCommunicator2_GetRpc2Msg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc2.proto",
}
