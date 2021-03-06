// Code generated by protoc-gen-go.
// source: foo.proto
// DO NOT EDIT!

/*
Package foo is a generated protocol buffer package.

It is generated from these files:
	foo.proto

It has these top-level messages:
	Message
	Header
	Body
	SimpleResponse
	Rc
*/
package foo

import proto "gitlab.ucloudadmin.com/udb/uframework/message/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Message struct {
	Header           *Header `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Body             *Body   `protobuf:"bytes,2,req,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Message) GetBody() *Body {
	if m != nil {
		return m.Body
	}
	return nil
}

type Header struct {
	Type             *int32 `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Header) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

type Body struct {
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Body) Reset()                    { *m = Body{} }
func (m *Body) String() string            { return proto.CompactTextString(m) }
func (*Body) ProtoMessage()               {}
func (*Body) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

var extRange_Body = []proto.ExtensionRange{
	{1, 536870911},
}

func (*Body) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Body
}
func (m *Body) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

type SimpleResponse struct {
	Rc               *Rc    `protobuf:"bytes,1,req,name=rc" json:"rc,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SimpleResponse) Reset()                    { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string            { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()               {}
func (*SimpleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SimpleResponse) GetRc() *Rc {
	if m != nil {
		return m.Rc
	}
	return nil
}

type Rc struct {
	RetCode          *int32  `protobuf:"varint,1,req,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	RetMessage       *string `protobuf:"bytes,2,req,name=ret_message,json=retMessage" json:"ret_message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Rc) Reset()                    { *m = Rc{} }
func (m *Rc) String() string            { return proto.CompactTextString(m) }
func (*Rc) ProtoMessage()               {}
func (*Rc) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Rc) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *Rc) GetRetMessage() string {
	if m != nil && m.RetMessage != nil {
		return *m.RetMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "foo.Message")
	proto.RegisterType((*Header)(nil), "foo.Header")
	proto.RegisterType((*Body)(nil), "foo.Body")
	proto.RegisterType((*SimpleResponse)(nil), "foo.SimpleResponse")
	proto.RegisterType((*Rc)(nil), "foo.Rc")
}

var fileDescriptor0 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xb1, 0x4e, 0xc4, 0x30,
	0x10, 0x44, 0x75, 0x26, 0x5c, 0x92, 0x8d, 0x84, 0x4e, 0xdb, 0x70, 0x48, 0x20, 0x90, 0x69, 0x80,
	0xe2, 0x0a, 0xbe, 0x00, 0x41, 0x43, 0x73, 0xcd, 0xf2, 0x01, 0x28, 0xd8, 0x1b, 0x40, 0x22, 0xac,
	0x65, 0xbb, 0x49, 0xe7, 0x4f, 0x47, 0xd9, 0x84, 0x6e, 0x35, 0xf3, 0x76, 0x34, 0x03, 0xed, 0x20,
	0x72, 0x08, 0x51, 0xb2, 0xe0, 0xc9, 0x20, 0x62, 0x8f, 0x50, 0x1f, 0x39, 0xa5, 0xfe, 0x93, 0xf1,
	0x16, 0xb6, 0x5f, 0xdc, 0x7b, 0x8e, 0xfb, 0xcd, 0x8d, 0xb9, 0xeb, 0x1e, 0xbb, 0xc3, 0xcc, 0xbe,
	0xaa, 0x44, 0xab, 0x85, 0x57, 0x50, 0x7d, 0x88, 0x9f, 0xf6, 0x46, 0x91, 0x56, 0x91, 0x67, 0xf1,
	0x13, 0xa9, 0x6c, 0x2f, 0x61, 0xbb, 0x3c, 0x20, 0x42, 0x95, 0xa7, 0xc0, 0x9a, 0x75, 0x4a, 0x7a,
	0xdb, 0x1d, 0x54, 0x33, 0xfb, 0xd0, 0x34, 0x9b, 0x5d, 0x29, 0xa5, 0x18, 0x7b, 0x0f, 0x67, 0x6f,
	0xdf, 0x63, 0xf8, 0x61, 0xe2, 0x14, 0xe4, 0x37, 0x31, 0x9e, 0x83, 0x89, 0x6e, 0x6d, 0x50, 0x6b,
	0x3c, 0x39, 0x32, 0xd1, 0xd9, 0x27, 0x30, 0xe4, 0xf0, 0x02, 0x9a, 0xc8, 0xf9, 0xdd, 0x89, 0xff,
	0x8f, 0xae, 0x23, 0xe7, 0x17, 0xf1, 0x8c, 0xd7, 0xd0, 0xcd, 0xd6, 0xb8, 0xcc, 0xd1, 0x86, 0x2d,
	0x41, 0xe4, 0xbc, 0x0e, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x17, 0xca, 0x9b, 0xaa, 0xfc, 0x00,
	0x00, 0x00,
}
