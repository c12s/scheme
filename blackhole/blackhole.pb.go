// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blackhole/blackhole.proto

/*
Package blackhole is a generated protocol buffer package.

It is generated from these files:
	blackhole/blackhole.proto

It has these top-level messages:
	Strategy
	KV
	Selector
	Payload
	Metadata
	PutTask
	PutReq
	Resp
*/
package blackhole

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

// Enums
type TaskKind int32

const (
	TaskKind_CONFIGS    TaskKind = 0
	TaskKind_ACTIONS    TaskKind = 1
	TaskKind_SECRETS    TaskKind = 2
	TaskKind_NAMESPACES TaskKind = 3
)

var TaskKind_name = map[int32]string{
	0: "CONFIGS",
	1: "ACTIONS",
	2: "SECRETS",
	3: "NAMESPACES",
}
var TaskKind_value = map[string]int32{
	"CONFIGS":    0,
	"ACTIONS":    1,
	"SECRETS":    2,
	"NAMESPACES": 3,
}

func (x TaskKind) String() string {
	return proto.EnumName(TaskKind_name, int32(x))
}
func (TaskKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PayloadKind int32

const (
	PayloadKind_FILE   PayloadKind = 0
	PayloadKind_ENV    PayloadKind = 1
	PayloadKind_ACTION PayloadKind = 2
)

var PayloadKind_name = map[int32]string{
	0: "FILE",
	1: "ENV",
	2: "ACTION",
}
var PayloadKind_value = map[string]int32{
	"FILE":   0,
	"ENV":    1,
	"ACTION": 2,
}

func (x PayloadKind) String() string {
	return proto.EnumName(PayloadKind_name, int32(x))
}
func (PayloadKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type StrategyKind int32

const (
	StrategyKind_AT_ONCE        StrategyKind = 0
	StrategyKind_ROLLING_UPDATE StrategyKind = 1
)

var StrategyKind_name = map[int32]string{
	0: "AT_ONCE",
	1: "ROLLING_UPDATE",
}
var StrategyKind_value = map[string]int32{
	"AT_ONCE":        0,
	"ROLLING_UPDATE": 1,
}

func (x StrategyKind) String() string {
	return proto.EnumName(StrategyKind_name, int32(x))
}
func (StrategyKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type CompareKind int32

const (
	CompareKind_ALL CompareKind = 0
	CompareKind_ANY CompareKind = 1
)

var CompareKind_name = map[int32]string{
	0: "ALL",
	1: "ANY",
}
var CompareKind_value = map[string]int32{
	"ALL": 0,
	"ANY": 1,
}

func (x CompareKind) String() string {
	return proto.EnumName(CompareKind_name, int32(x))
}
func (CompareKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Request helper messages
type Strategy struct {
	Type string       `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Kind StrategyKind `protobuf:"varint,2,opt,name=kind,enum=blackhole.StrategyKind" json:"kind,omitempty"`
}

func (m *Strategy) Reset()                    { *m = Strategy{} }
func (m *Strategy) String() string            { return proto.CompactTextString(m) }
func (*Strategy) ProtoMessage()               {}
func (*Strategy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Strategy) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Strategy) GetKind() StrategyKind {
	if m != nil {
		return m.Kind
	}
	return StrategyKind_AT_ONCE
}

type KV struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *KV) Reset()                    { *m = KV{} }
func (m *KV) String() string            { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()               {}
func (*KV) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *KV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KV) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Selector struct {
	Kind   CompareKind `protobuf:"varint,1,opt,name=kind,enum=blackhole.CompareKind" json:"kind,omitempty"`
	Labels []*KV       `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty"`
}

func (m *Selector) Reset()                    { *m = Selector{} }
func (m *Selector) String() string            { return proto.CompactTextString(m) }
func (*Selector) ProtoMessage()               {}
func (*Selector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Selector) GetKind() CompareKind {
	if m != nil {
		return m.Kind
	}
	return CompareKind_ALL
}

func (m *Selector) GetLabels() []*KV {
	if m != nil {
		return m.Labels
	}
	return nil
}

type Payload struct {
	Kind  PayloadKind `protobuf:"varint,1,opt,name=kind,enum=blackhole.PayloadKind" json:"kind,omitempty"`
	Value []*KV       `protobuf:"bytes,2,rep,name=value" json:"value,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Payload) GetKind() PayloadKind {
	if m != nil {
		return m.Kind
	}
	return PayloadKind_FILE
}

func (m *Payload) GetValue() []*KV {
	if m != nil {
		return m.Value
	}
	return nil
}

type Metadata struct {
	TaskName            string `protobuf:"bytes,1,opt,name=taskName" json:"taskName,omitempty"`
	Timestamp           int64  `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Namespace           string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	ForceNamespaceQueue bool   `protobuf:"varint,4,opt,name=forceNamespaceQueue" json:"forceNamespaceQueue,omitempty"`
	Queue               string `protobuf:"bytes,5,opt,name=queue" json:"queue,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Metadata) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *Metadata) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Metadata) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Metadata) GetForceNamespaceQueue() bool {
	if m != nil {
		return m.ForceNamespaceQueue
	}
	return false
}

func (m *Metadata) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

type PutTask struct {
	RegionId  string     `protobuf:"bytes,1,opt,name=region_id,json=regionId" json:"region_id,omitempty"`
	ClusterId string     `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId" json:"cluster_id,omitempty"`
	Strategy  *Strategy  `protobuf:"bytes,3,opt,name=strategy" json:"strategy,omitempty"`
	Selector  *Selector  `protobuf:"bytes,4,opt,name=selector" json:"selector,omitempty"`
	Payload   []*Payload `protobuf:"bytes,5,rep,name=payload" json:"payload,omitempty"`
}

func (m *PutTask) Reset()                    { *m = PutTask{} }
func (m *PutTask) String() string            { return proto.CompactTextString(m) }
func (*PutTask) ProtoMessage()               {}
func (*PutTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PutTask) GetRegionId() string {
	if m != nil {
		return m.RegionId
	}
	return ""
}

func (m *PutTask) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *PutTask) GetStrategy() *Strategy {
	if m != nil {
		return m.Strategy
	}
	return nil
}

func (m *PutTask) GetSelector() *Selector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *PutTask) GetPayload() []*Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Reques messages
type PutReq struct {
	Version string     `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	UserId  string     `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Kind    TaskKind   `protobuf:"varint,3,opt,name=kind,enum=blackhole.TaskKind" json:"kind,omitempty"`
	Mtdata  *Metadata  `protobuf:"bytes,4,opt,name=mtdata" json:"mtdata,omitempty"`
	Tasks   []*PutTask `protobuf:"bytes,5,rep,name=tasks" json:"tasks,omitempty"`
	Extras  []*KV      `protobuf:"bytes,6,rep,name=extras" json:"extras,omitempty"`
}

func (m *PutReq) Reset()                    { *m = PutReq{} }
func (m *PutReq) String() string            { return proto.CompactTextString(m) }
func (*PutReq) ProtoMessage()               {}
func (*PutReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PutReq) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *PutReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *PutReq) GetKind() TaskKind {
	if m != nil {
		return m.Kind
	}
	return TaskKind_CONFIGS
}

func (m *PutReq) GetMtdata() *Metadata {
	if m != nil {
		return m.Mtdata
	}
	return nil
}

func (m *PutReq) GetTasks() []*PutTask {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *PutReq) GetExtras() []*KV {
	if m != nil {
		return m.Extras
	}
	return nil
}

// Response message
type Resp struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Resp) Reset()                    { *m = Resp{} }
func (m *Resp) String() string            { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()               {}
func (*Resp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Resp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Strategy)(nil), "blackhole.Strategy")
	proto.RegisterType((*KV)(nil), "blackhole.KV")
	proto.RegisterType((*Selector)(nil), "blackhole.Selector")
	proto.RegisterType((*Payload)(nil), "blackhole.Payload")
	proto.RegisterType((*Metadata)(nil), "blackhole.Metadata")
	proto.RegisterType((*PutTask)(nil), "blackhole.PutTask")
	proto.RegisterType((*PutReq)(nil), "blackhole.PutReq")
	proto.RegisterType((*Resp)(nil), "blackhole.Resp")
	proto.RegisterEnum("blackhole.TaskKind", TaskKind_name, TaskKind_value)
	proto.RegisterEnum("blackhole.PayloadKind", PayloadKind_name, PayloadKind_value)
	proto.RegisterEnum("blackhole.StrategyKind", StrategyKind_name, StrategyKind_value)
	proto.RegisterEnum("blackhole.CompareKind", CompareKind_name, CompareKind_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BlackHoleService service

type BlackHoleServiceClient interface {
	Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*Resp, error)
}

type blackHoleServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlackHoleServiceClient(cc *grpc.ClientConn) BlackHoleServiceClient {
	return &blackHoleServiceClient{cc}
}

func (c *blackHoleServiceClient) Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/blackhole.BlackHoleService/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BlackHoleService service

type BlackHoleServiceServer interface {
	Put(context.Context, *PutReq) (*Resp, error)
}

func RegisterBlackHoleServiceServer(s *grpc.Server, srv BlackHoleServiceServer) {
	s.RegisterService(&_BlackHoleService_serviceDesc, srv)
}

func _BlackHoleService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlackHoleServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blackhole.BlackHoleService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlackHoleServiceServer).Put(ctx, req.(*PutReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlackHoleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blackhole.BlackHoleService",
	HandlerType: (*BlackHoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _BlackHoleService_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blackhole/blackhole.proto",
}

func init() { proto.RegisterFile("blackhole/blackhole.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x7d, 0x6e, 0xd3, 0x4a,
	0x10, 0xaf, 0xe3, 0x7c, 0x38, 0x93, 0xf7, 0xfa, 0xfc, 0xb6, 0x88, 0x9a, 0x02, 0x22, 0x32, 0x42,
	0x44, 0x69, 0xd5, 0x40, 0x38, 0x00, 0x4a, 0x43, 0x5a, 0xa2, 0xa4, 0x4e, 0x58, 0x87, 0x4a, 0x54,
	0x42, 0xd5, 0xc6, 0x5e, 0x52, 0x2b, 0x76, 0xec, 0x7a, 0xd7, 0x15, 0xb9, 0x12, 0xf7, 0xe1, 0x02,
	0x9c, 0x04, 0xed, 0xda, 0x4e, 0x0c, 0x44, 0xfc, 0xb7, 0x33, 0xf3, 0x9b, 0xdf, 0x7c, 0x2f, 0x3c,
	0x9a, 0xfb, 0xc4, 0x59, 0xde, 0x86, 0x3e, 0xed, 0x6c, 0x5e, 0xa7, 0x51, 0x1c, 0xf2, 0x10, 0xd5,
	0x37, 0x0a, 0x73, 0x04, 0x9a, 0xcd, 0x63, 0xc2, 0xe9, 0x62, 0x8d, 0x10, 0x94, 0xf9, 0x3a, 0xa2,
	0x86, 0xd2, 0x54, 0x5a, 0x75, 0x2c, 0xdf, 0xe8, 0x18, 0xca, 0x4b, 0x6f, 0xe5, 0x1a, 0xa5, 0xa6,
	0xd2, 0xda, 0xef, 0x1e, 0x9e, 0x6e, 0xa9, 0x72, 0xb7, 0x91, 0xb7, 0x72, 0xb1, 0x04, 0x99, 0x27,
	0x50, 0x1a, 0x5d, 0x21, 0x1d, 0xd4, 0x25, 0x5d, 0x67, 0x2c, 0xe2, 0x89, 0x1e, 0x40, 0xe5, 0x9e,
	0xf8, 0x09, 0x95, 0x2c, 0x75, 0x9c, 0x0a, 0xe6, 0x67, 0xd0, 0x6c, 0xea, 0x53, 0x87, 0x87, 0x31,
	0x6a, 0x67, 0x61, 0x14, 0x19, 0xe6, 0x61, 0x21, 0x4c, 0x3f, 0x0c, 0x22, 0x12, 0xd3, 0x6d, 0x14,
	0xf4, 0x02, 0xaa, 0x3e, 0x99, 0x53, 0x9f, 0x19, 0xa5, 0xa6, 0xda, 0x6a, 0x74, 0xff, 0x2d, 0xa0,
	0x47, 0x57, 0x38, 0x33, 0x9a, 0xd7, 0x50, 0x9b, 0x92, 0xb5, 0x1f, 0x12, 0xf7, 0x2f, 0xec, 0x19,
	0xa2, 0xc0, 0xfe, 0x7c, 0x9b, 0xeb, 0x0e, 0xf2, 0x2c, 0xf5, 0x6f, 0x0a, 0x68, 0x97, 0x94, 0x13,
	0x97, 0x70, 0x82, 0x8e, 0x40, 0xe3, 0x84, 0x2d, 0x2d, 0x12, 0xe4, 0xad, 0xdb, 0xc8, 0xe8, 0x09,
	0xd4, 0xb9, 0x17, 0x50, 0xc6, 0x49, 0x10, 0xc9, 0xea, 0x55, 0xbc, 0x55, 0x08, 0xeb, 0x8a, 0x04,
	0x94, 0x45, 0xc4, 0xa1, 0x86, 0x2a, 0x5d, 0xb7, 0x0a, 0xf4, 0x0a, 0x0e, 0xbe, 0x84, 0xb1, 0x43,
	0xad, 0x5c, 0xf3, 0x21, 0xa1, 0x09, 0x35, 0xca, 0x4d, 0xa5, 0xa5, 0xe1, 0x5d, 0x26, 0xd1, 0xe7,
	0x3b, 0x89, 0xa9, 0xa4, 0x7d, 0x96, 0x82, 0xf9, 0x5d, 0x81, 0xda, 0x34, 0xe1, 0x33, 0xc2, 0x96,
	0xe8, 0x31, 0xd4, 0x63, 0xba, 0xf0, 0xc2, 0xd5, 0x8d, 0xe7, 0xe6, 0xc9, 0xa6, 0x8a, 0xa1, 0x8b,
	0x9e, 0x02, 0x38, 0x7e, 0xc2, 0x38, 0x8d, 0x85, 0x35, 0x9d, 0x55, 0x3d, 0xd3, 0x0c, 0x5d, 0xd4,
	0x01, 0x8d, 0x65, 0x33, 0x97, 0xc9, 0x36, 0xba, 0x07, 0x3b, 0xd6, 0x01, 0x6f, 0x40, 0xd2, 0x21,
	0x1b, 0xb0, 0xcc, 0xfa, 0x37, 0x87, 0xcc, 0x84, 0x37, 0x20, 0x74, 0x02, 0xb5, 0x28, 0x1d, 0x88,
	0x51, 0x91, 0xdd, 0x47, 0x7f, 0x8e, 0x0a, 0xe7, 0x10, 0xf3, 0x87, 0x02, 0xd5, 0x69, 0xc2, 0x31,
	0xbd, 0x43, 0x06, 0xd4, 0xee, 0x69, 0xcc, 0xbc, 0x70, 0x95, 0x15, 0x95, 0x8b, 0xe8, 0x10, 0x6a,
	0x09, 0x2b, 0x16, 0x54, 0x15, 0xe2, 0xd0, 0x45, 0x2f, 0xb3, 0x9d, 0x50, 0xe5, 0x4e, 0x14, 0x13,
	0x13, 0x8d, 0x2a, 0x2c, 0xc4, 0x31, 0x54, 0x03, 0x2e, 0x06, 0xbd, 0xa3, 0x86, 0x7c, 0x07, 0x70,
	0x06, 0x41, 0x2d, 0xa8, 0x88, 0xd9, 0xb3, 0x5d, 0xf9, 0xa7, 0x23, 0xc0, 0x29, 0x40, 0x6c, 0x31,
	0xfd, 0xca, 0x63, 0xc2, 0x8c, 0xea, 0xce, 0x2d, 0x4e, 0x8d, 0xa6, 0x01, 0x65, 0x4c, 0x59, 0x24,
	0x8e, 0x2a, 0x60, 0x8b, 0xfc, 0xa8, 0x02, 0xb6, 0x68, 0xf7, 0x40, 0xcb, 0x33, 0x45, 0x0d, 0xa8,
	0xf5, 0x27, 0xd6, 0xf9, 0xf0, 0xc2, 0xd6, 0xf7, 0x84, 0xd0, 0xeb, 0xcf, 0x86, 0x13, 0xcb, 0xd6,
	0x15, 0x21, 0xd8, 0x83, 0x3e, 0x1e, 0xcc, 0x6c, 0xbd, 0x84, 0xf6, 0x01, 0xac, 0xde, 0xe5, 0xc0,
	0x9e, 0xf6, 0xfa, 0x03, 0x5b, 0x57, 0xdb, 0x27, 0xd0, 0x28, 0x1c, 0x00, 0xd2, 0xa0, 0x7c, 0x3e,
	0x1c, 0x0f, 0xf4, 0x3d, 0x54, 0x03, 0x75, 0x60, 0x5d, 0xe9, 0x0a, 0x02, 0xa8, 0xa6, 0x5c, 0x7a,
	0xa9, 0xdd, 0x81, 0x7f, 0x8a, 0x37, 0x2f, 0xe3, 0xcc, 0x6e, 0x26, 0x56, 0x5f, 0x78, 0x20, 0xd8,
	0xc7, 0x93, 0xf1, 0x78, 0x68, 0x5d, 0xdc, 0x7c, 0x9c, 0xbe, 0xeb, 0xcd, 0x06, 0xba, 0xd2, 0x7e,
	0x06, 0x8d, 0xc2, 0xf5, 0x0a, 0xd2, 0xde, 0x78, 0x9c, 0xb2, 0xf7, 0xac, 0x4f, 0xba, 0xd2, 0x7d,
	0x0b, 0xfa, 0x99, 0x28, 0xfa, 0x7d, 0xe8, 0x53, 0x9b, 0xc6, 0xf7, 0x9e, 0x23, 0x3e, 0x1c, 0x75,
	0x9a, 0x70, 0xf4, 0xff, 0xaf, 0x9d, 0xc3, 0xf4, 0xee, 0xe8, 0xbf, 0x82, 0x4a, 0xf4, 0xc4, 0xdc,
	0x3b, 0x33, 0xaf, 0x9b, 0x0b, 0x8f, 0xdf, 0x26, 0xf3, 0x53, 0x27, 0x0c, 0x3a, 0xce, 0xeb, 0x2e,
	0xeb, 0x30, 0xe7, 0x96, 0x06, 0x85, 0x2f, 0x6f, 0x5e, 0x95, 0x7f, 0xde, 0x9b, 0x9f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x90, 0x37, 0x86, 0x84, 0x10, 0x05, 0x00, 0x00,
}
