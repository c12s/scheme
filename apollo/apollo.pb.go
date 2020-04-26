// Code generated by protoc-gen-go. DO NOT EDIT.
// source: apollo/apollo.proto

package apollo

import (
	context "context"
	fmt "fmt"
	celestial "github.com/c12s/scheme/celestial"
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

type OptExtras struct {
	Data                 []string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OptExtras) Reset()         { *m = OptExtras{} }
func (m *OptExtras) String() string { return proto.CompactTextString(m) }
func (*OptExtras) ProtoMessage()    {}
func (*OptExtras) Descriptor() ([]byte, []int) {
	return fileDescriptor_556057f10c4e752b, []int{0}
}

func (m *OptExtras) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OptExtras.Unmarshal(m, b)
}
func (m *OptExtras) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OptExtras.Marshal(b, m, deterministic)
}
func (m *OptExtras) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OptExtras.Merge(m, src)
}
func (m *OptExtras) XXX_Size() int {
	return xxx_messageInfo_OptExtras.Size(m)
}
func (m *OptExtras) XXX_DiscardUnknown() {
	xxx_messageInfo_OptExtras.DiscardUnknown(m)
}

var xxx_messageInfo_OptExtras proto.InternalMessageInfo

func (m *OptExtras) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

type AuthOpt struct {
	Data                 map[string]string     `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Extras               map[string]*OptExtras `protobuf:"bytes,2,rep,name=extras,proto3" json:"extras,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AuthOpt) Reset()         { *m = AuthOpt{} }
func (m *AuthOpt) String() string { return proto.CompactTextString(m) }
func (*AuthOpt) ProtoMessage()    {}
func (*AuthOpt) Descriptor() ([]byte, []int) {
	return fileDescriptor_556057f10c4e752b, []int{1}
}

func (m *AuthOpt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthOpt.Unmarshal(m, b)
}
func (m *AuthOpt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthOpt.Marshal(b, m, deterministic)
}
func (m *AuthOpt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthOpt.Merge(m, src)
}
func (m *AuthOpt) XXX_Size() int {
	return xxx_messageInfo_AuthOpt.Size(m)
}
func (m *AuthOpt) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthOpt.DiscardUnknown(m)
}

var xxx_messageInfo_AuthOpt proto.InternalMessageInfo

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
	Value                bool              `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Data                 map[string]string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AuthResp) Reset()         { *m = AuthResp{} }
func (m *AuthResp) String() string { return proto.CompactTextString(m) }
func (*AuthResp) ProtoMessage()    {}
func (*AuthResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_556057f10c4e752b, []int{2}
}

func (m *AuthResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResp.Unmarshal(m, b)
}
func (m *AuthResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResp.Marshal(b, m, deterministic)
}
func (m *AuthResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResp.Merge(m, src)
}
func (m *AuthResp) XXX_Size() int {
	return xxx_messageInfo_AuthResp.Size(m)
}
func (m *AuthResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResp.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResp proto.InternalMessageInfo

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
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReq) Reset()         { *m = GetReq{} }
func (m *GetReq) String() string { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()    {}
func (*GetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_556057f10c4e752b, []int{3}
}

func (m *GetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReq.Unmarshal(m, b)
}
func (m *GetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReq.Marshal(b, m, deterministic)
}
func (m *GetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReq.Merge(m, src)
}
func (m *GetReq) XXX_Size() int {
	return xxx_messageInfo_GetReq.Size(m)
}
func (m *GetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetReq proto.InternalMessageInfo

func (m *GetReq) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type GetResp struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResp) Reset()         { *m = GetResp{} }
func (m *GetResp) String() string { return proto.CompactTextString(m) }
func (*GetResp) ProtoMessage()    {}
func (*GetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_556057f10c4e752b, []int{4}
}

func (m *GetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResp.Unmarshal(m, b)
}
func (m *GetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResp.Marshal(b, m, deterministic)
}
func (m *GetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResp.Merge(m, src)
}
func (m *GetResp) XXX_Size() int {
	return xxx_messageInfo_GetResp.Size(m)
}
func (m *GetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetResp proto.InternalMessageInfo

func (m *GetResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ACL struct {
	Token                []string `protobuf:"bytes,1,rep,name=token,proto3" json:"token,omitempty"`
	Created              int64    `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ACL) Reset()         { *m = ACL{} }
func (m *ACL) String() string { return proto.CompactTextString(m) }
func (*ACL) ProtoMessage()    {}
func (*ACL) Descriptor() ([]byte, []int) {
	return fileDescriptor_556057f10c4e752b, []int{5}
}

func (m *ACL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ACL.Unmarshal(m, b)
}
func (m *ACL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ACL.Marshal(b, m, deterministic)
}
func (m *ACL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACL.Merge(m, src)
}
func (m *ACL) XXX_Size() int {
	return xxx_messageInfo_ACL.Size(m)
}
func (m *ACL) XXX_DiscardUnknown() {
	xxx_messageInfo_ACL.DiscardUnknown(m)
}

var xxx_messageInfo_ACL proto.InternalMessageInfo

func (m *ACL) GetToken() []string {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *ACL) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

type User struct {
	Firtname             string            `protobuf:"bytes,1,opt,name=firtname,proto3" json:"firtname,omitempty"`
	Lastname             string            `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Organisation         string            `protobuf:"bytes,3,opt,name=organisation,proto3" json:"organisation,omitempty"`
	Role                 string            `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	Username             string            `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Password             string            `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Labels               map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_556057f10c4e752b, []int{6}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetFirtname() string {
	if m != nil {
		return m.Firtname
	}
	return ""
}

func (m *User) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *User) GetOrganisation() string {
	if m != nil {
		return m.Organisation
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*OptExtras)(nil), "apollo.OptExtras")
	proto.RegisterType((*AuthOpt)(nil), "apollo.AuthOpt")
	proto.RegisterMapType((map[string]string)(nil), "apollo.AuthOpt.DataEntry")
	proto.RegisterMapType((map[string]*OptExtras)(nil), "apollo.AuthOpt.ExtrasEntry")
	proto.RegisterType((*AuthResp)(nil), "apollo.AuthResp")
	proto.RegisterMapType((map[string]string)(nil), "apollo.AuthResp.DataEntry")
	proto.RegisterType((*GetReq)(nil), "apollo.GetReq")
	proto.RegisterType((*GetResp)(nil), "apollo.GetResp")
	proto.RegisterType((*ACL)(nil), "apollo.ACL")
	proto.RegisterType((*User)(nil), "apollo.User")
	proto.RegisterMapType((map[string]string)(nil), "apollo.User.LabelsEntry")
}

func init() {
	proto.RegisterFile("apollo/apollo.proto", fileDescriptor_556057f10c4e752b)
}

var fileDescriptor_556057f10c4e752b = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xae, 0xe3, 0xd4, 0x49, 0x26, 0xbf, 0x1f, 0x2d, 0xdb, 0x22, 0x6d, 0x0d, 0x28, 0xd5, 0x5e,
	0xa8, 0x84, 0xea, 0xd0, 0x54, 0xa8, 0xc0, 0x2d, 0x40, 0xd5, 0x4b, 0x50, 0xa5, 0x00, 0x17, 0x6e,
	0x1b, 0x67, 0x68, 0xac, 0x3a, 0x59, 0xb3, 0xbb, 0x2e, 0xf4, 0x11, 0x78, 0x06, 0x1e, 0x8c, 0x17,
	0xe1, 0x01, 0xd0, 0xfe, 0xb1, 0xeb, 0x86, 0x5e, 0xe0, 0x94, 0xf9, 0xe6, 0x9b, 0xef, 0x9b, 0x9d,
	0x9d, 0x8d, 0x61, 0x87, 0x17, 0x22, 0xcf, 0xc5, 0xd0, 0xfd, 0x24, 0x85, 0x14, 0x5a, 0x90, 0xc8,
	0xa1, 0x78, 0x2f, 0xc5, 0x1c, 0x95, 0xce, 0x78, 0x3e, 0xac, 0x23, 0x57, 0xc2, 0x06, 0xd0, 0x3b,
	0x2f, 0xf4, 0xe9, 0x37, 0x2d, 0xb9, 0x22, 0x04, 0xda, 0x73, 0xae, 0x39, 0x0d, 0xf6, 0xc3, 0x83,
	0xde, 0xd4, 0xc6, 0xec, 0x57, 0x00, 0x9d, 0x71, 0xa9, 0x17, 0xe7, 0x85, 0x26, 0x87, 0x0d, 0xbe,
	0x3f, 0xda, 0x4b, 0x7c, 0x33, 0x4f, 0x27, 0x6f, 0xb9, 0xe6, 0xa7, 0x2b, 0x2d, 0xaf, 0x9d, 0x94,
	0x1c, 0x43, 0x84, 0xd6, 0x98, 0xb6, 0xac, 0xe0, 0xe1, 0xba, 0xc0, 0xb5, 0x75, 0x12, 0x5f, 0x1a,
	0x9f, 0x40, 0xaf, 0xf6, 0x21, 0xdb, 0x10, 0x5e, 0xe2, 0x35, 0x0d, 0xf6, 0x83, 0x83, 0xde, 0xd4,
	0x84, 0x64, 0x17, 0x36, 0xaf, 0x78, 0x5e, 0x22, 0x6d, 0xd9, 0x9c, 0x03, 0xaf, 0x5a, 0x2f, 0x82,
	0x78, 0x02, 0xfd, 0x86, 0xdf, 0x1d, 0xd2, 0x27, 0x4d, 0x69, 0x7f, 0x74, 0xbf, 0x3a, 0x4d, 0x3d,
	0x7f, 0xc3, 0x8d, 0x7d, 0x0f, 0xa0, 0x6b, 0x8e, 0x39, 0x45, 0x55, 0xdc, 0x34, 0x35, 0x6e, 0x5d,
	0x5f, 0x46, 0x12, 0x7f, 0x1b, 0x6e, 0xb8, 0xb8, 0x39, 0x9c, 0x51, 0xad, 0x5f, 0xc7, 0x3f, 0x4f,
	0xc6, 0x1e, 0x41, 0x74, 0x86, 0x7a, 0x8a, 0x5f, 0xcc, 0x82, 0x4a, 0x85, 0xd2, 0xcb, 0x6c, 0xcc,
	0x06, 0xd0, 0xb1, 0xac, 0x3b, 0xa7, 0x16, 0x97, 0xb8, 0xf2, 0xbc, 0x03, 0xec, 0x39, 0x84, 0xe3,
	0x37, 0x93, 0x26, 0x19, 0xd6, 0x24, 0xa1, 0xd0, 0x49, 0x25, 0x72, 0x8d, 0x73, 0xdb, 0x37, 0x9c,
	0x56, 0x90, 0xfd, 0x68, 0x41, 0xfb, 0xa3, 0x42, 0x49, 0x62, 0xe8, 0x7e, 0xce, 0xa4, 0x5e, 0xf1,
	0x25, 0x7a, 0xe3, 0x1a, 0x1b, 0x2e, 0xe7, 0xca, 0x71, 0xee, 0xdc, 0x35, 0x26, 0x0c, 0xfe, 0x13,
	0xf2, 0x82, 0xaf, 0x32, 0xc5, 0x75, 0x26, 0x56, 0x34, 0xb4, 0xfc, 0xad, 0x9c, 0x19, 0x48, 0x8a,
	0x1c, 0x69, 0xdb, 0x0d, 0x64, 0x62, 0xe3, 0x69, 0x06, 0xb3, 0x9e, 0x9b, 0xce, 0xb3, 0xc2, 0x86,
	0x2b, 0xb8, 0x52, 0x5f, 0x85, 0x9c, 0xd3, 0xc8, 0x71, 0x15, 0x26, 0xcf, 0x20, 0xca, 0xf9, 0x0c,
	0x73, 0x45, 0x3b, 0x76, 0x23, 0xb4, 0xda, 0x88, 0x99, 0x22, 0x99, 0x58, 0xca, 0xbf, 0x35, 0x57,
	0x17, 0xbf, 0x84, 0x7e, 0x23, 0xfd, 0x37, 0x3b, 0x19, 0xfd, 0x0c, 0xe0, 0xff, 0xb1, 0xb5, 0x7f,
	0x8f, 0xf2, 0x2a, 0x4b, 0x91, 0x1c, 0x42, 0xf7, 0x0c, 0xf5, 0x07, 0x7b, 0xab, 0xf7, 0xaa, 0xd6,
	0x6e, 0x6f, 0xf1, 0xd6, 0x2d, 0xac, 0x0a, 0xb6, 0x41, 0x9e, 0x42, 0xdb, 0xbc, 0x14, 0xb2, 0xb5,
	0xf6, 0xa7, 0x88, 0xb7, 0xd7, 0x1f, 0x12, 0xdb, 0x20, 0x47, 0xd0, 0x9e, 0x64, 0x4a, 0x13, 0x92,
	0xdc, 0xfc, 0x7f, 0x4d, 0xc2, 0x78, 0xef, 0xfc, 0x91, 0xb3, 0x92, 0x13, 0x88, 0xde, 0x95, 0x9a,
	0x6b, 0x24, 0xbb, 0x8d, 0x02, 0x97, 0x32, 0xb2, 0x07, 0x77, 0x64, 0x8d, 0xf0, 0xf5, 0xe0, 0xd3,
	0xe3, 0x8b, 0x4c, 0x2f, 0xca, 0x59, 0x92, 0x8a, 0xe5, 0x30, 0x3d, 0x1a, 0xa9, 0xa1, 0x4a, 0x17,
	0xb8, 0x44, 0xff, 0x6d, 0x99, 0x45, 0xf6, 0xcb, 0x71, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x99,
	0x8f, 0xe5, 0x12, 0x73, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ApolloServiceClient is the client API for ApolloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApolloServiceClient interface {
	GetToken(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
	Auth(ctx context.Context, in *AuthOpt, opts ...grpc.CallOption) (*AuthResp, error)
	List(ctx context.Context, in *celestial.ListReq, opts ...grpc.CallOption) (*celestial.ListResp, error)
	Mutate(ctx context.Context, in *celestial.MutateReq, opts ...grpc.CallOption) (*celestial.MutateResp, error)
}

type apolloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApolloServiceClient(cc grpc.ClientConnInterface) ApolloServiceClient {
	return &apolloServiceClient{cc}
}

func (c *apolloServiceClient) GetToken(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := c.cc.Invoke(ctx, "/apollo.ApolloService/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apolloServiceClient) Auth(ctx context.Context, in *AuthOpt, opts ...grpc.CallOption) (*AuthResp, error) {
	out := new(AuthResp)
	err := c.cc.Invoke(ctx, "/apollo.ApolloService/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apolloServiceClient) List(ctx context.Context, in *celestial.ListReq, opts ...grpc.CallOption) (*celestial.ListResp, error) {
	out := new(celestial.ListResp)
	err := c.cc.Invoke(ctx, "/apollo.ApolloService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apolloServiceClient) Mutate(ctx context.Context, in *celestial.MutateReq, opts ...grpc.CallOption) (*celestial.MutateResp, error) {
	out := new(celestial.MutateResp)
	err := c.cc.Invoke(ctx, "/apollo.ApolloService/Mutate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApolloServiceServer is the server API for ApolloService service.
type ApolloServiceServer interface {
	GetToken(context.Context, *GetReq) (*GetResp, error)
	Auth(context.Context, *AuthOpt) (*AuthResp, error)
	List(context.Context, *celestial.ListReq) (*celestial.ListResp, error)
	Mutate(context.Context, *celestial.MutateReq) (*celestial.MutateResp, error)
}

// UnimplementedApolloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedApolloServiceServer struct {
}

func (*UnimplementedApolloServiceServer) GetToken(ctx context.Context, req *GetReq) (*GetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (*UnimplementedApolloServiceServer) Auth(ctx context.Context, req *AuthOpt) (*AuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (*UnimplementedApolloServiceServer) List(ctx context.Context, req *celestial.ListReq) (*celestial.ListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedApolloServiceServer) Mutate(ctx context.Context, req *celestial.MutateReq) (*celestial.MutateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mutate not implemented")
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

func _ApolloService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(celestial.ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApolloServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apollo.ApolloService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApolloServiceServer).List(ctx, req.(*celestial.ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApolloService_Mutate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(celestial.MutateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApolloServiceServer).Mutate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apollo.ApolloService/Mutate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApolloServiceServer).Mutate(ctx, req.(*celestial.MutateReq))
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
		{
			MethodName: "List",
			Handler:    _ApolloService_List_Handler,
		},
		{
			MethodName: "Mutate",
			Handler:    _ApolloService_Mutate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apollo/apollo.proto",
}
