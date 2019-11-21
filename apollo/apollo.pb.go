// Code generated by protoc-gen-go. DO NOT EDIT.
// source: apollo/apollo.proto

/*
Package apollo is a generated protocol buffer package.

It is generated from these files:
	apollo/apollo.proto

It has these top-level messages:
	AuthOpt
	AuthResp
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

type AuthOpt struct {
	Data map[string]string `protobuf:"bytes,1,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AuthOpt) Reset()                    { *m = AuthOpt{} }
func (m *AuthOpt) String() string            { return proto.CompactTextString(m) }
func (*AuthOpt) ProtoMessage()               {}
func (*AuthOpt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthOpt) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type AuthResp struct {
	Value bool              `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Data  map[string]string `protobuf:"bytes,2,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AuthResp) Reset()                    { *m = AuthResp{} }
func (m *AuthResp) String() string            { return proto.CompactTextString(m) }
func (*AuthResp) ProtoMessage()               {}
func (*AuthResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthResp) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func (m *AuthResp) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetReq struct {
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *GetReq) Reset()                    { *m = GetReq{} }
func (m *GetReq) String() string            { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()               {}
func (*GetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
func (*GetResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthOpt)(nil), "apollo.AuthOpt")
	proto.RegisterType((*AuthResp)(nil), "apollo.AuthResp")
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
	Auth(ctx context.Context, in *AuthOpt, opts ...grpc.CallOption) (*AuthResp, error)
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

func (c *apolloServiceClient) Auth(ctx context.Context, in *AuthOpt, opts ...grpc.CallOption) (*AuthResp, error) {
	out := new(AuthResp)
	err := grpc.Invoke(ctx, "/apollo.ApolloService/Auth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApolloService service

type ApolloServiceServer interface {
	GetToken(context.Context, *GetReq) (*GetResp, error)
	Auth(context.Context, *AuthOpt) (*AuthResp, error)
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

func _ApolloService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthOpt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApolloServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apollo.ApolloService/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApolloServiceServer).Auth(ctx, req.(*AuthOpt))
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
		{
			MethodName: "Auth",
			Handler:    _ApolloService_Auth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apollo/apollo.proto",
}

func init() { proto.RegisterFile("apollo/apollo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xcd, 0x4e, 0x83, 0x50,
	0x10, 0x85, 0x0b, 0x45, 0x4a, 0xc7, 0x68, 0x9b, 0xd1, 0x05, 0x12, 0x4d, 0x1b, 0x56, 0x4d, 0x4c,
	0x21, 0xe2, 0x42, 0xe3, 0xae, 0x46, 0xd3, 0xa5, 0x09, 0xba, 0x72, 0x77, 0x8b, 0x13, 0x69, 0xa0,
	0xbd, 0x14, 0x2e, 0x4d, 0xfa, 0x08, 0xbe, 0xb5, 0xb9, 0x3f, 0xfe, 0x75, 0xed, 0x8a, 0x39, 0x73,
	0xce, 0xc0, 0x37, 0x0c, 0x9c, 0xb0, 0x8a, 0x97, 0x25, 0x8f, 0xf5, 0x23, 0xaa, 0x6a, 0x2e, 0x38,
	0xba, 0x5a, 0x85, 0x1b, 0xe8, 0xcd, 0x5a, 0x91, 0x3f, 0x55, 0x02, 0xa7, 0xe0, 0xbc, 0x31, 0xc1,
	0x7c, 0x6b, 0xdc, 0x9d, 0x1c, 0x26, 0x67, 0x91, 0xc9, 0x1b, 0x3b, 0x7a, 0x60, 0x82, 0x3d, 0xae,
	0x45, 0xbd, 0x4b, 0x55, 0x2c, 0xb8, 0x81, 0xfe, 0x77, 0x0b, 0x87, 0xd0, 0x2d, 0x68, 0xe7, 0x5b,
	0x63, 0x6b, 0xd2, 0x4f, 0x65, 0x89, 0xa7, 0x70, 0xb0, 0x65, 0x65, 0x4b, 0xbe, 0xad, 0x7a, 0x5a,
	0xdc, 0xd9, 0xb7, 0x56, 0xf8, 0x61, 0x81, 0x27, 0x5f, 0x9a, 0x52, 0x53, 0xfd, 0xc4, 0xe4, 0xa8,
	0x67, 0x62, 0x18, 0x19, 0x14, 0x5b, 0xa1, 0x04, 0xbf, 0x51, 0xe4, 0xd4, 0xff, 0xb1, 0x9c, 0x83,
	0x3b, 0x27, 0x91, 0xd2, 0x06, 0x11, 0x9c, 0xb6, 0xa1, 0xda, 0x8c, 0xa9, 0x3a, 0x1c, 0x41, 0x4f,
	0xb9, 0x9a, 0x53, 0xf0, 0x82, 0xd6, 0xc6, 0xd7, 0x22, 0x29, 0xe0, 0x68, 0xa6, 0xd0, 0x9e, 0xa9,
	0xde, 0x2e, 0x33, 0xc2, 0x29, 0x78, 0x73, 0x12, 0x2f, 0xd2, 0xc4, 0xe3, 0x2f, 0x6c, 0xfd, 0x85,
	0x60, 0xf0, 0x47, 0x37, 0x55, 0xd8, 0xc1, 0x4b, 0x70, 0xe4, 0x4e, 0x38, 0xd8, 0xfb, 0xd9, 0xc1,
	0x70, 0x7f, 0xe5, 0xb0, 0x73, 0x3f, 0x7a, 0xbd, 0x78, 0x5f, 0x8a, 0xbc, 0x5d, 0x44, 0x19, 0x5f,
	0xc5, 0xd9, 0x55, 0xd2, 0xc4, 0x4d, 0x96, 0xd3, 0x8a, 0xcc, 0x65, 0x17, 0xae, 0x3a, 0xed, 0xf5,
	0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0xdc, 0x4c, 0xf8, 0xf1, 0x01, 0x00, 0x00,
}
