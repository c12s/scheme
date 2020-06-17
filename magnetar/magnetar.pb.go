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
	Metricttl            string   `protobuf:"bytes,1,opt,name=metricttl,proto3" json:"metricttl,omitempty"`
	Ids                  []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
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

func (m *ReserveMsg) GetMetricttl() string {
	if m != nil {
		return m.Metricttl
	}
	return ""
}

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

type EventMsg struct {
	Data                 map[string]*DataMsg `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *EventMsg) Reset()         { *m = EventMsg{} }
func (m *EventMsg) String() string { return proto.CompactTextString(m) }
func (*EventMsg) ProtoMessage()    {}
func (*EventMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_d955f06648f68850, []int{4}
}

func (m *EventMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventMsg.Unmarshal(m, b)
}
func (m *EventMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventMsg.Marshal(b, m, deterministic)
}
func (m *EventMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMsg.Merge(m, src)
}
func (m *EventMsg) XXX_Size() int {
	return xxx_messageInfo_EventMsg.Size(m)
}
func (m *EventMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMsg.DiscardUnknown(m)
}

var xxx_messageInfo_EventMsg proto.InternalMessageInfo

func (m *EventMsg) GetData() map[string]*DataMsg {
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
	proto.RegisterType((*EventMsg)(nil), "magnetar.EventMsg")
	proto.RegisterMapType((map[string]*DataMsg)(nil), "magnetar.EventMsg.DataEntry")
}

func init() {
	proto.RegisterFile("magnetar/magnetar.proto", fileDescriptor_d955f06648f68850)
}

var fileDescriptor_d955f06648f68850 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xc1, 0x6e, 0xda, 0x40,
	0x10, 0x65, 0x01, 0x17, 0x3c, 0x1c, 0x5a, 0x56, 0x48, 0xb5, 0x28, 0x52, 0x5d, 0x5f, 0xea, 0x93,
	0x5d, 0xdc, 0x4a, 0x54, 0x55, 0x95, 0x43, 0x14, 0x72, 0x88, 0xc2, 0x21, 0x9b, 0x5b, 0x6e, 0x8b,
	0x3d, 0x32, 0x16, 0xd8, 0xa0, 0xdd, 0xb5, 0x15, 0x7e, 0x20, 0xca, 0x87, 0xe6, 0x43, 0x22, 0x83,
	0xb1, 0x89, 0x1c, 0x22, 0xe5, 0x90, 0xdb, 0xcc, 0xdb, 0xf7, 0xfc, 0xde, 0x68, 0x3c, 0xf0, 0x35,
	0xe6, 0x61, 0x82, 0x8a, 0x0b, 0xf7, 0x50, 0x38, 0x1b, 0xb1, 0x56, 0x6b, 0xda, 0x3d, 0xf4, 0xd6,
	0x7f, 0x00, 0x86, 0x12, 0x45, 0x86, 0x33, 0x19, 0xd2, 0x11, 0xe8, 0x31, 0x2a, 0x11, 0xf9, 0x4a,
	0xad, 0x0c, 0x62, 0x12, 0x5b, 0x67, 0x15, 0x40, 0xbf, 0x40, 0x2b, 0x0a, 0xa4, 0xd1, 0x34, 0x5b,
	0xb6, 0xce, 0xf2, 0xd2, 0x3a, 0x2b, 0xd5, 0x4c, 0x6e, 0xe8, 0x00, 0x34, 0xc5, 0x97, 0x98, 0xec,
	0x94, 0x1a, 0xdb, 0x37, 0x74, 0x08, 0x5d, 0xbc, 0xf7, 0x57, 0x69, 0x80, 0x41, 0x21, 0x2d, 0x7b,
	0xeb, 0x81, 0x40, 0xe7, 0x3a, 0x92, 0x2a, 0x57, 0xbb, 0xd0, 0x0e, 0xb8, 0xe2, 0x06, 0x31, 0x5b,
	0x76, 0xcf, 0xfb, 0xe6, 0x94, 0x91, 0x0b, 0x82, 0x73, 0xc1, 0x15, 0x9f, 0x26, 0x4a, 0x6c, 0xd9,
	0x8e, 0x38, 0xbc, 0x02, 0xbd, 0x84, 0xf2, 0x6c, 0x4b, 0xdc, 0x16, 0x99, 0xf3, 0x92, 0xfe, 0x04,
	0x2d, 0xe3, 0xab, 0x14, 0x8d, 0xa6, 0x49, 0xec, 0x9e, 0xd7, 0xaf, 0x3e, 0x98, 0xab, 0x66, 0x32,
	0x64, 0xfb, 0xf7, 0x7f, 0xcd, 0xbf, 0xc4, 0x92, 0xd0, 0x29, 0xd0, 0xd3, 0x39, 0x0a, 0x42, 0x2d,
	0xc7, 0xe4, 0xed, 0x1c, 0x83, 0xe3, 0x1c, 0xfa, 0xb1, 0xe9, 0x23, 0x81, 0xee, 0x34, 0xc3, 0x44,
	0xe5, 0xb6, 0xbf, 0x5e, 0xd8, 0x8e, 0x2a, 0xdb, 0x03, 0xe3, 0x23, 0xe7, 0xf7, 0x9e, 0x08, 0x7c,
	0x9e, 0x15, 0xef, 0xb7, 0x28, 0xb2, 0xc8, 0x47, 0x3a, 0x81, 0x4e, 0xb1, 0x5c, 0x3a, 0xa8, 0xc4,
	0xd5, 0xdf, 0x32, 0xac, 0xa3, 0x4c, 0x6e, 0xac, 0x06, 0x1d, 0x43, 0x3b, 0xdf, 0xd9, 0x09, 0x55,
	0xbf, 0xb6, 0x59, 0xab, 0x41, 0xff, 0x40, 0xfb, 0x52, 0xe0, 0x7b, 0x8d, 0x5c, 0xd0, 0x6e, 0x52,
	0x14, 0x5b, 0x5a, 0x1f, 0xee, 0x55, 0x9b, 0xf3, 0x1f, 0x77, 0xdf, 0xc3, 0x48, 0x2d, 0xd2, 0xb9,
	0xe3, 0xaf, 0x63, 0xd7, 0x1f, 0x7b, 0xd2, 0x95, 0xfe, 0x02, 0x63, 0x2c, 0x0f, 0x64, 0xfe, 0x69,
	0x77, 0x21, 0xbf, 0x9f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x81, 0x2d, 0x35, 0x0e, 0x3c, 0x03, 0x00,
	0x00,
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
	Query(ctx context.Context, in *DataMsg, opts ...grpc.CallOption) (*ListRsp, error)
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

func (c *magnetarServiceClient) Query(ctx context.Context, in *DataMsg, opts ...grpc.CallOption) (*ListRsp, error) {
	out := new(ListRsp)
	err := c.cc.Invoke(ctx, "/magnetar.MagnetarService/Query", in, out, opts...)
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
	Query(context.Context, *DataMsg) (*ListRsp, error)
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
func (*UnimplementedMagnetarServiceServer) Query(ctx context.Context, req *DataMsg) (*ListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
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

func _MagnetarService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MagnetarServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magnetar.MagnetarService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MagnetarServiceServer).Query(ctx, req.(*DataMsg))
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
		{
			MethodName: "Query",
			Handler:    _MagnetarService_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "magnetar/magnetar.proto",
}
