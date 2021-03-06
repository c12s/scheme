// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flusher/flusher.proto

/*
Package flusher is a generated protocol buffer package.

It is generated from these files:
	flusher/flusher.proto

It has these top-level messages:
	Event
	Update
	Error
*/
package flusher

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import blackhole "github.com/c12s/scheme/blackhole"
import stellar "github.com/c12s/scheme/stellar"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Event struct {
	TaskKey     string               `protobuf:"bytes,1,opt,name=taskKey" json:"taskKey,omitempty"`
	Kind        string               `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
	Payload     []*blackhole.Payload `protobuf:"bytes,3,rep,name=payload" json:"payload,omitempty"`
	SpanContext *stellar.SpanContext `protobuf:"bytes,4,opt,name=spanContext" json:"spanContext,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetTaskKey() string {
	if m != nil {
		return m.TaskKey
	}
	return ""
}

func (m *Event) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Event) GetPayload() []*blackhole.Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Event) GetSpanContext() *stellar.SpanContext {
	if m != nil {
		return m.SpanContext
	}
	return nil
}

type Update struct {
	TaskKey     string               `protobuf:"bytes,1,opt,name=taskKey" json:"taskKey,omitempty"`
	Kind        string               `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
	Node        string               `protobuf:"bytes,3,opt,name=node" json:"node,omitempty"`
	SpanContext *stellar.SpanContext `protobuf:"bytes,4,opt,name=spanContext" json:"spanContext,omitempty"`
}

func (m *Update) Reset()                    { *m = Update{} }
func (m *Update) String() string            { return proto.CompactTextString(m) }
func (*Update) ProtoMessage()               {}
func (*Update) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Update) GetTaskKey() string {
	if m != nil {
		return m.TaskKey
	}
	return ""
}

func (m *Update) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Update) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *Update) GetSpanContext() *stellar.SpanContext {
	if m != nil {
		return m.SpanContext
	}
	return nil
}

type Error struct {
	TaskKey   string `protobuf:"bytes,1,opt,name=taskKey" json:"taskKey,omitempty"`
	Kind      string `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
	Node      string `protobuf:"bytes,3,opt,name=node" json:"node,omitempty"`
	Message   string `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	Timestamp int64  `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
	Level     string `protobuf:"bytes,6,opt,name=level" json:"level,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Error) GetTaskKey() string {
	if m != nil {
		return m.TaskKey
	}
	return ""
}

func (m *Error) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Error) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Error) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "flusher.Event")
	proto.RegisterType((*Update)(nil), "flusher.Update")
	proto.RegisterType((*Error)(nil), "flusher.Error")
}

func init() { proto.RegisterFile("flusher/flusher.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xcd, 0x4a, 0xf4, 0x30,
	0x14, 0x25, 0x5f, 0xff, 0xe8, 0xed, 0x2e, 0xcc, 0x40, 0xbe, 0x41, 0xa4, 0xcc, 0xaa, 0x0b, 0x69,
	0xb1, 0x82, 0x0f, 0xa0, 0xb8, 0x72, 0x23, 0x15, 0x37, 0xee, 0xd2, 0xf6, 0x3a, 0x2d, 0x4d, 0x9b,
	0x92, 0x64, 0x06, 0x67, 0xeb, 0x4b, 0xb8, 0xf0, 0x65, 0x65, 0xfa, 0x63, 0xdd, 0x2a, 0xae, 0xee,
	0x39, 0xe7, 0x9e, 0x70, 0x0e, 0x49, 0x60, 0xfd, 0x22, 0xf6, 0xba, 0x42, 0x95, 0x4c, 0x33, 0xee,
	0x95, 0x34, 0x92, 0x7a, 0x13, 0xdd, 0xfc, 0xcf, 0x05, 0x2f, 0x9a, 0x4a, 0x0a, 0x4c, 0xbe, 0xd0,
	0xe8, 0xd9, 0xac, 0xb5, 0x41, 0x21, 0xb8, 0x4a, 0xa6, 0x39, 0xca, 0xdb, 0x0f, 0x02, 0xce, 0xdd,
	0x01, 0x3b, 0x43, 0x19, 0x78, 0x86, 0xeb, 0xe6, 0x1e, 0x8f, 0x8c, 0x84, 0x24, 0xf2, 0xb3, 0x99,
	0x52, 0x0a, 0x76, 0x53, 0x77, 0x25, 0xfb, 0x37, 0xc8, 0x03, 0xa6, 0x17, 0xe0, 0xf5, 0xfc, 0x28,
	0x24, 0x2f, 0x99, 0x15, 0x5a, 0x51, 0x90, 0xd2, 0x78, 0x49, 0x7c, 0x18, 0x37, 0xd9, 0x6c, 0xa1,
	0xd7, 0x10, 0xe8, 0x9e, 0x77, 0xb7, 0xb2, 0x33, 0xf8, 0x6a, 0x98, 0x1d, 0x92, 0x28, 0x48, 0x57,
	0xf1, 0x5c, 0xe5, 0x71, 0xd9, 0x65, 0xdf, 0x8d, 0xdb, 0x37, 0x02, 0xee, 0x53, 0x5f, 0x72, 0x83,
	0x3f, 0xac, 0x47, 0xc1, 0xee, 0x64, 0x89, 0xcc, 0x1a, 0xb5, 0x13, 0xfe, 0x75, 0x89, 0xf7, 0xd3,
	0x15, 0x29, 0x25, 0xd5, 0x1f, 0x74, 0x60, 0xe0, 0xb5, 0xa8, 0x35, 0xdf, 0xe1, 0x90, 0xef, 0x67,
	0x33, 0xa5, 0x67, 0xe0, 0x9b, 0xba, 0x45, 0x6d, 0x78, 0xdb, 0x33, 0x27, 0x24, 0x91, 0x95, 0x2d,
	0x02, 0x5d, 0x81, 0x23, 0xf0, 0x80, 0x82, 0xb9, 0xc3, 0xa9, 0x91, 0xdc, 0x84, 0xcf, 0xe7, 0xbb,
	0xda, 0x54, 0xfb, 0x3c, 0x2e, 0x64, 0x9b, 0x14, 0x97, 0xa9, 0x4e, 0x74, 0x51, 0x61, 0x8b, 0xf3,
	0xff, 0xc8, 0xdd, 0xe1, 0x95, 0xaf, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x88, 0xd2, 0x5c,
	0x39, 0x02, 0x00, 0x00,
}
