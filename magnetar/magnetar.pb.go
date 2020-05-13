// Code generated by protoc-gen-go. DO NOT EDIT.
// source: magnetar/magnetar.proto

package magnetar

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

type ReserveMsg struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReserveMsg) Reset()         { *m = ReserveMsg{} }
func (m *ReserveMsg) String() string { return proto.CompactTextString(m) }
func (*ReserveMsg) ProtoMessage()    {}
func (*ReserveMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_d955f06648f68850, []int{0}
}

func (m *ReserveMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReserveMsg.Unmarshal(m, b)
}
func (m *ReserveMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReserveMsg.Marshal(b, m, deterministic)
}
func (m *ReserveMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReserveMsg.Merge(m, src)
}
func (m *ReserveMsg) XXX_Size() int {
	return xxx_messageInfo_ReserveMsg.Size(m)
}
func (m *ReserveMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ReserveMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ReserveMsg proto.InternalMessageInfo

func (m *ReserveMsg) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type ReserveRsp struct {
	Taken                int32    `protobuf:"varint,1,opt,name=taken,proto3" json:"taken,omitempty"`
	Excluded             []string `protobuf:"bytes,2,rep,name=excluded,proto3" json:"excluded,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReserveRsp) Reset()         { *m = ReserveRsp{} }
func (m *ReserveRsp) String() string { return proto.CompactTextString(m) }
func (*ReserveRsp) ProtoMessage()    {}
func (*ReserveRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d955f06648f68850, []int{1}
}

func (m *ReserveRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReserveRsp.Unmarshal(m, b)
}
func (m *ReserveRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReserveRsp.Marshal(b, m, deterministic)
}
func (m *ReserveRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReserveRsp.Merge(m, src)
}
func (m *ReserveRsp) XXX_Size() int {
	return xxx_messageInfo_ReserveRsp.Size(m)
}
func (m *ReserveRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ReserveRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ReserveRsp proto.InternalMessageInfo

func (m *ReserveRsp) GetTaken() int32 {
	if m != nil {
		return m.Taken
	}
	return 0
}

func (m *ReserveRsp) GetExcluded() []string {
	if m != nil {
		return m.Excluded
	}
	return nil
}

type ListRsp struct {
	Data                 map[string]*DataMsg `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ListRsp) Reset()         { *m = ListRsp{} }
func (m *ListRsp) String() string { return proto.CompactTextString(m) }
func (*ListRsp) ProtoMessage()    {}
func (*ListRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d955f06648f68850, []int{2}
}

func (m *ListRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRsp.Unmarshal(m, b)
}
func (m *ListRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRsp.Marshal(b, m, deterministic)
}
func (m *ListRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRsp.Merge(m, src)
}
func (m *ListRsp) XXX_Size() int {
	return xxx_messageInfo_ListRsp.Size(m)
}
func (m *ListRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ListRsp proto.InternalMessageInfo

func (m *ListRsp) GetData() map[string]*DataMsg {
	if m != nil {
		return m.Data
	}
	return nil
}

type DataMsg struct {
	Data                 map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DataMsg) Reset()         { *m = DataMsg{} }
func (m *DataMsg) String() string { return proto.CompactTextString(m) }
func (*DataMsg) ProtoMessage()    {}
func (*DataMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_d955f06648f68850, []int{3}
}

func (m *DataMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataMsg.Unmarshal(m, b)
}
func (m *DataMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataMsg.Marshal(b, m, deterministic)
}
func (m *DataMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataMsg.Merge(m, src)
}
func (m *DataMsg) XXX_Size() int {
	return xxx_messageInfo_DataMsg.Size(m)
}
func (m *DataMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_DataMsg.DiscardUnknown(m)
}

var xxx_messageInfo_DataMsg proto.InternalMessageInfo

func (m *DataMsg) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ReserveMsg)(nil), "magnetar.ReserveMsg")
	proto.RegisterType((*ReserveRsp)(nil), "magnetar.ReserveRsp")
	proto.RegisterType((*ListRsp)(nil), "magnetar.ListRsp")
	proto.RegisterMapType((map[string]*DataMsg)(nil), "magnetar.ListRsp.DataEntry")
	proto.RegisterType((*DataMsg)(nil), "magnetar.DataMsg")
	proto.RegisterMapType((map[string]string)(nil), "magnetar.DataMsg.DataEntry")
}

func init() {
	proto.RegisterFile("magnetar/magnetar.proto", fileDescriptor_d955f06648f68850)
}

var fileDescriptor_d955f06648f68850 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0xbb, 0xfd, 0xb0, 0xcd, 0xf4, 0xa0, 0x5d, 0x0a, 0x86, 0x0a, 0x5a, 0xf7, 0x62, 0x4f,
	0x09, 0x8d, 0x42, 0xc5, 0x83, 0x07, 0x51, 0x0f, 0x62, 0x2f, 0xeb, 0xcd, 0xdb, 0x36, 0x19, 0xd2,
	0xd0, 0xaf, 0xb0, 0xbb, 0x09, 0xf6, 0x05, 0x7c, 0x21, 0x5f, 0x50, 0x36, 0x4d, 0x93, 0x42, 0xa8,
	0xe0, 0x6d, 0x66, 0x76, 0xfe, 0xf3, 0xff, 0x31, 0x3b, 0x70, 0xbe, 0x12, 0xe1, 0x1a, 0xb5, 0x90,
	0xee, 0x3e, 0x70, 0x62, 0xb9, 0xd1, 0x1b, 0xda, 0xd9, 0xe7, 0xec, 0x12, 0x80, 0xa3, 0x42, 0x99,
	0xe2, 0x54, 0x85, 0xf4, 0x0c, 0x1a, 0x51, 0xa0, 0x6c, 0x32, 0x6c, 0x8c, 0x2c, 0x6e, 0x42, 0xf6,
	0x58, 0xbc, 0x73, 0x15, 0xd3, 0x3e, 0xb4, 0xb4, 0x58, 0xe0, 0xda, 0x26, 0x43, 0x32, 0x6a, 0xf1,
	0x5d, 0x42, 0x07, 0xd0, 0xc1, 0x2f, 0x7f, 0x99, 0x04, 0x18, 0xd8, 0xf5, 0x4c, 0x5a, 0xe4, 0xec,
	0x9b, 0x40, 0xfb, 0x3d, 0x52, 0xda, 0xa8, 0x5d, 0x68, 0x06, 0x42, 0x8b, 0x6c, 0x7c, 0xd7, 0xbb,
	0x70, 0x0a, 0xa8, 0xbc, 0xc1, 0x79, 0x16, 0x5a, 0xbc, 0xac, 0xb5, 0xdc, 0xf2, 0xac, 0x71, 0xf0,
	0x06, 0x56, 0x51, 0x32, 0x6c, 0x0b, 0xdc, 0x66, 0xce, 0x16, 0x37, 0x21, 0xbd, 0x81, 0x56, 0x2a,
	0x96, 0x09, 0xda, 0xf5, 0x21, 0x19, 0x75, 0xbd, 0x5e, 0x39, 0xd0, 0xa8, 0xa6, 0x2a, 0xe4, 0xbb,
	0xf7, 0x87, 0xfa, 0x3d, 0x61, 0x0a, 0xda, 0x79, 0xf5, 0x38, 0x47, 0xde, 0x50, 0xe1, 0x98, 0xfc,
	0xcd, 0xd1, 0x3f, 0xe4, 0xb0, 0x0e, 0x4c, 0xbd, 0x1f, 0x02, 0xa7, 0xd3, 0x7c, 0xfa, 0x07, 0xca,
	0x34, 0xf2, 0x91, 0x4e, 0xa0, 0x9d, 0x6f, 0x94, 0xf6, 0x4b, 0xeb, 0xf2, 0x13, 0x06, 0xd5, 0x2a,
	0x57, 0x31, 0xab, 0xd1, 0x31, 0x34, 0xcd, 0xa2, 0x8e, 0xa8, 0x7a, 0x95, 0x75, 0xb2, 0x1a, 0xbd,
	0x83, 0xe6, 0xab, 0xc4, 0x7f, 0x1a, 0x3d, 0x5d, 0x7f, 0x5e, 0x85, 0x91, 0x9e, 0x27, 0x33, 0xc7,
	0xdf, 0xac, 0x5c, 0x7f, 0xec, 0x29, 0x57, 0xf9, 0x73, 0x5c, 0x61, 0x71, 0x46, 0xb3, 0x93, 0xec,
	0x8e, 0x6e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x44, 0x9a, 0x8e, 0xb7, 0x62, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MagnetarServiceClient is the client API for MagnetarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MagnetarServiceClient interface {
	Reserve(ctx context.Context, in *ReserveMsg, opts ...grpc.CallOption) (*ReserveRsp, error)
	List(ctx context.Context, in *ReserveMsg, opts ...grpc.CallOption) (*ListRsp, error)
	Free(ctx context.Context, in *ReserveMsg, opts ...grpc.CallOption) (*ReserveRsp, error)
}

type magnetarServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMagnetarServiceClient(cc grpc.ClientConnInterface) MagnetarServiceClient {
	return &magnetarServiceClient{cc}
}

func (c *magnetarServiceClient) Reserve(ctx context.Context, in *ReserveMsg, opts ...grpc.CallOption) (*ReserveRsp, error) {
	out := new(ReserveRsp)
	err := c.cc.Invoke(ctx, "/magnetar.MagnetarService/Reserve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *magnetarServiceClient) List(ctx context.Context, in *ReserveMsg, opts ...grpc.CallOption) (*ListRsp, error) {
	out := new(ListRsp)
	err := c.cc.Invoke(ctx, "/magnetar.MagnetarService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *magnetarServiceClient) Free(ctx context.Context, in *ReserveMsg, opts ...grpc.CallOption) (*ReserveRsp, error) {
	out := new(ReserveRsp)
	err := c.cc.Invoke(ctx, "/magnetar.MagnetarService/Free", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MagnetarServiceServer is the server API for MagnetarService service.
type MagnetarServiceServer interface {
	Reserve(context.Context, *ReserveMsg) (*ReserveRsp, error)
	List(context.Context, *ReserveMsg) (*ListRsp, error)
	Free(context.Context, *ReserveMsg) (*ReserveRsp, error)
}

// UnimplementedMagnetarServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMagnetarServiceServer struct {
}

func (*UnimplementedMagnetarServiceServer) Reserve(ctx context.Context, req *ReserveMsg) (*ReserveRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reserve not implemented")
}
func (*UnimplementedMagnetarServiceServer) List(ctx context.Context, req *ReserveMsg) (*ListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedMagnetarServiceServer) Free(ctx context.Context, req *ReserveMsg) (*ReserveRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Free not implemented")
}

func RegisterMagnetarServiceServer(s *grpc.Server, srv MagnetarServiceServer) {
	s.RegisterService(&_MagnetarService_serviceDesc, srv)
}

func _MagnetarService_Reserve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MagnetarServiceServer).Reserve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magnetar.MagnetarService/Reserve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MagnetarServiceServer).Reserve(ctx, req.(*ReserveMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _MagnetarService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MagnetarServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magnetar.MagnetarService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MagnetarServiceServer).List(ctx, req.(*ReserveMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _MagnetarService_Free_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MagnetarServiceServer).Free(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magnetar.MagnetarService/Free",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MagnetarServiceServer).Free(ctx, req.(*ReserveMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _MagnetarService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magnetar.MagnetarService",
	HandlerType: (*MagnetarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reserve",
			Handler:    _MagnetarService_Reserve_Handler,
		},
		{
			MethodName: "List",
			Handler:    _MagnetarService_List_Handler,
		},
		{
			MethodName: "Free",
			Handler:    _MagnetarService_Free_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "magnetar/magnetar.proto",
}
