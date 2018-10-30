// Code generated by protoc-gen-go. DO NOT EDIT.
// source: apollo/apollo.proto

/*
Package apollo is a generated protocol buffer package.

It is generated from these files:
	apollo/apollo.proto

It has these top-level messages:
	GetReq
	GetResp
*/
package apollo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetReq struct {
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *GetReq) Reset()                    { *m = GetReq{} }
func (m *GetReq) String() string            { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()               {}
func (*GetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetReq) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type GetResp struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *GetResp) Reset()                    { *m = GetResp{} }
func (m *GetResp) String() string            { return proto.CompactTextString(m) }
func (*GetResp) ProtoMessage()               {}
func (*GetResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*GetReq)(nil), "apollo.GetReq")
	proto.RegisterType((*GetResp)(nil), "apollo.GetResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ApolloService service

type ApolloServiceClient interface {
	GetToken(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
}

type apolloServiceClient struct {
	cc *grpc.ClientConn
}

func NewApolloServiceClient(cc *grpc.ClientConn) ApolloServiceClient {
	return &apolloServiceClient{cc}
}

func (c *apolloServiceClient) GetToken(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := grpc.Invoke(ctx, "/apollo.ApolloService/GetToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApolloService service

type ApolloServiceServer interface {
	GetToken(context.Context, *GetReq) (*GetResp, error)
}

func RegisterApolloServiceServer(s *grpc.Server, srv ApolloServiceServer) {
	s.RegisterService(&_ApolloService_serviceDesc, srv)
}

func _ApolloService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApolloServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apollo.ApolloService/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApolloServiceServer).GetToken(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApolloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apollo.ApolloService",
	HandlerType: (*ApolloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToken",
			Handler:    _ApolloService_GetToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apollo/apollo.proto",
}

func init() { proto.RegisterFile("apollo/apollo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2c, 0xc8, 0xcf,
	0xc9, 0xc9, 0xd7, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x92,
	0x0c, 0x17, 0x9b, 0x7b, 0x6a, 0x49, 0x50, 0x6a, 0xa1, 0x90, 0x10, 0x17, 0x4b, 0x69, 0x71, 0x6a,
	0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xad, 0x24, 0xcf, 0xc5, 0x0e, 0x96, 0x2d,
	0x2e, 0x10, 0x12, 0xe1, 0x62, 0x2d, 0xc9, 0xcf, 0x4e, 0xcd, 0x83, 0xca, 0x43, 0x38, 0x46, 0x76,
	0x5c, 0xbc, 0x8e, 0x60, 0x83, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x74, 0xb9, 0x38,
	0xdc, 0x53, 0x4b, 0x42, 0x40, 0x92, 0x42, 0x7c, 0x7a, 0x50, 0x2b, 0x21, 0x36, 0x48, 0xf1, 0xa3,
	0xf0, 0x8b, 0x0b, 0x94, 0x18, 0x9c, 0xe4, 0xa3, 0x64, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4,
	0x92, 0xf3, 0x73, 0xf5, 0x93, 0x0d, 0x8d, 0x8a, 0xf5, 0x8b, 0x93, 0x33, 0x52, 0x73, 0x53, 0xa1,
	0xae, 0x4d, 0x62, 0x03, 0x3b, 0xd7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xb9, 0x2f, 0x54,
	0xc5, 0x00, 0x00, 0x00,
}
