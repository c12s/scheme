// Code generated by protoc-gen-go. DO NOT EDIT.
// source: apollo/apollo.proto

/*
Package apollo is a generated protocol buffer package.

It is generated from these files:
	apollo/apollo.proto

It has these top-level messages:
	OptExtras
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

type OptExtras struct {
	Data []string `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
}

func (m *OptExtras) Reset()                    { *m = OptExtras{} }
func (m *OptExtras) String() string            { return proto.CompactTextString(m) }
func (*OptExtras) ProtoMessage()               {}
func (*OptExtras) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OptExtras) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

type AuthOpt struct {
	Data   map[string]string     `protobuf:"bytes,1,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Extras map[string]*OptExtras `protobuf:"bytes,2,rep,name=extras" json:"extras,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AuthOpt) Reset()                    { *m = AuthOpt{} }
func (m *AuthOpt) String() string            { return proto.CompactTextString(m) }
func (*AuthOpt) ProtoMessage()               {}
func (*AuthOpt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthOpt) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *AuthOpt) GetExtras() map[string]*OptExtras {
	if m != nil {
		return m.Extras
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
func (*AuthResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
func (*GetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
func (*GetResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*OptExtras)(nil), "apollo.OptExtras")
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
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0xa5, 0xc0, 0x57, 0xe8, 0x25, 0x9f, 0xe0, 0xe8, 0xa2, 0x56, 0x0d, 0x64, 0x36, 0x92, 0x18,
	0x4a, 0x2c, 0x0b, 0x8d, 0x3b, 0x8c, 0x84, 0x8d, 0x09, 0x49, 0x75, 0xe5, 0x6e, 0xa8, 0x37, 0x42,
	0xf8, 0x69, 0xed, 0x4c, 0x89, 0x3c, 0x82, 0xef, 0xeb, 0x03, 0x98, 0xf9, 0x01, 0x6b, 0xe3, 0xca,
	0x55, 0xef, 0xed, 0x99, 0x73, 0xee, 0x39, 0x33, 0x17, 0x8e, 0x58, 0x12, 0x2f, 0x97, 0x71, 0x5f,
	0x7f, 0xfc, 0x24, 0x8d, 0x45, 0x4c, 0x6c, 0xdd, 0xd1, 0x36, 0x38, 0x93, 0x44, 0x8c, 0xde, 0x45,
	0xca, 0x38, 0x21, 0x50, 0x7d, 0x61, 0x82, 0xb9, 0x56, 0xa7, 0xd2, 0x75, 0x42, 0x55, 0xd3, 0x4f,
	0x0b, 0x6a, 0xc3, 0x4c, 0xcc, 0x26, 0x89, 0x20, 0xbd, 0x1c, 0xde, 0x08, 0x4e, 0x7c, 0xa3, 0x68,
	0x60, 0xff, 0x9e, 0x09, 0x36, 0x5a, 0x8b, 0x74, 0xab, 0xa9, 0x64, 0x00, 0x36, 0x2a, 0x61, 0xb7,
	0xac, 0x08, 0xa7, 0x45, 0x82, 0x1e, 0xab, 0x29, 0xe6, 0xa8, 0x77, 0x0d, 0xce, 0x5e, 0x87, 0xb4,
	0xa0, 0xb2, 0xc0, 0xad, 0x6b, 0x75, 0xac, 0xae, 0x13, 0xca, 0x92, 0x1c, 0xc3, 0xbf, 0x0d, 0x5b,
	0x66, 0xe8, 0x96, 0xd5, 0x3f, 0xdd, 0xdc, 0x96, 0x6f, 0x2c, 0xef, 0x01, 0x1a, 0x39, 0xbd, 0x5f,
	0xa8, 0x17, 0x79, 0x6a, 0x23, 0x38, 0xdc, 0xb9, 0xd9, 0xe7, 0xcf, 0xa9, 0xd1, 0x0f, 0x0b, 0xea,
	0xd2, 0x66, 0x88, 0x3c, 0xf9, 0x1e, 0x2a, 0xd5, 0xea, 0xe6, 0x18, 0xf1, 0xcd, 0x6d, 0xe8, 0x70,
	0x5e, 0x3e, 0x9c, 0x64, 0x15, 0xaf, 0xe3, 0xcf, 0xc9, 0xe8, 0x19, 0xd8, 0x63, 0x14, 0x21, 0xbe,
	0xc9, 0x07, 0xca, 0x38, 0xa6, 0x86, 0xa6, 0x6a, 0xda, 0x86, 0x9a, 0x42, 0xb5, 0x4f, 0x11, 0x2f,
	0x70, 0x6d, 0x70, 0xdd, 0x04, 0x0b, 0xf8, 0x3f, 0x54, 0xd6, 0x1e, 0x31, 0xdd, 0xcc, 0x23, 0x24,
	0x3d, 0xa8, 0x8f, 0x51, 0x3c, 0x49, 0x90, 0x1c, 0xec, 0x6c, 0xeb, 0x09, 0x5e, 0xf3, 0x47, 0xcf,
	0x13, 0x5a, 0x22, 0x97, 0x50, 0x95, 0x99, 0x48, 0xb3, 0xf0, 0x7c, 0x5e, 0xab, 0x18, 0x99, 0x96,
	0xee, 0xda, 0xcf, 0xe7, 0xaf, 0x73, 0x31, 0xcb, 0xa6, 0x7e, 0x14, 0xaf, 0xfa, 0xd1, 0x55, 0xc0,
	0xfb, 0x3c, 0x9a, 0xe1, 0x0a, 0xcd, 0xfa, 0x4d, 0x6d, 0xb5, 0x7f, 0x83, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xaf, 0x7e, 0x4a, 0x86, 0x96, 0x02, 0x00, 0x00,
}
