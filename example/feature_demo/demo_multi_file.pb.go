// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/feature_demo/demo_multi_file.proto

package example // import "github.com/YReshetko/protoc-gen-gorm/example/feature_demo"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/YReshetko/protoc-gen-gorm/options"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ExternalChild struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExternalChild) Reset()         { *m = ExternalChild{} }
func (m *ExternalChild) String() string { return proto.CompactTextString(m) }
func (*ExternalChild) ProtoMessage()    {}
func (*ExternalChild) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_multi_file_ff24e9aae73eb453, []int{0}
}
func (m *ExternalChild) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalChild.Unmarshal(m, b)
}
func (m *ExternalChild) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalChild.Marshal(b, m, deterministic)
}
func (dst *ExternalChild) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalChild.Merge(dst, src)
}
func (m *ExternalChild) XXX_Size() int {
	return xxx_messageInfo_ExternalChild.Size(m)
}
func (m *ExternalChild) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalChild.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalChild proto.InternalMessageInfo

func (m *ExternalChild) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type BlogPost struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author               string   `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlogPost) Reset()         { *m = BlogPost{} }
func (m *BlogPost) String() string { return proto.CompactTextString(m) }
func (*BlogPost) ProtoMessage()    {}
func (*BlogPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_multi_file_ff24e9aae73eb453, []int{1}
}
func (m *BlogPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogPost.Unmarshal(m, b)
}
func (m *BlogPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogPost.Marshal(b, m, deterministic)
}
func (dst *BlogPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogPost.Merge(dst, src)
}
func (m *BlogPost) XXX_Size() int {
	return xxx_messageInfo_BlogPost.Size(m)
}
func (m *BlogPost) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogPost.DiscardUnknown(m)
}

var xxx_messageInfo_BlogPost proto.InternalMessageInfo

func (m *BlogPost) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BlogPost) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *BlogPost) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func init() {
	proto.RegisterType((*ExternalChild)(nil), "example.ExternalChild")
	proto.RegisterType((*BlogPost)(nil), "example.BlogPost")
}

func init() {
	proto.RegisterFile("example/feature_demo/demo_multi_file.proto", fileDescriptor_demo_multi_file_ff24e9aae73eb453)
}

var fileDescriptor_demo_multi_file_ff24e9aae73eb453 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4a, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x4f, 0x4b, 0x4d, 0x2c, 0x29, 0x2d, 0x4a, 0x8d, 0x4f, 0x49, 0xcd, 0xcd,
	0xd7, 0x07, 0x11, 0xf1, 0xb9, 0xa5, 0x39, 0x25, 0x99, 0xf1, 0x69, 0x99, 0x39, 0xa9, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xb5, 0x52, 0xe6, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49,
	0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x91, 0x41, 0xa9, 0xc5, 0x19, 0xa9, 0x25, 0xd9, 0xf9, 0xfa, 0x60,
	0x45, 0xc9, 0xba, 0xe9, 0xa9, 0x79, 0xba, 0xe9, 0xf9, 0x45, 0xb9, 0xfa, 0xf9, 0x05, 0x25, 0x99,
	0xf9, 0x79, 0xc5, 0xfa, 0x20, 0x0e, 0xc4, 0x04, 0x25, 0x75, 0x2e, 0x5e, 0xd7, 0x8a, 0x92, 0xd4,
	0xa2, 0xbc, 0xc4, 0x1c, 0xe7, 0x8c, 0xcc, 0x9c, 0x14, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xa6, 0xcc, 0x14, 0x2b, 0xb6, 0x5d, 0x3b, 0x25, 0x99, 0x38,
	0x18, 0x95, 0x02, 0xb8, 0x38, 0x9c, 0x72, 0xf2, 0xd3, 0x03, 0xf2, 0x8b, 0x4b, 0x90, 0xd4, 0xb0,
	0x80, 0xd4, 0x08, 0x89, 0x70, 0xb1, 0x96, 0x64, 0x96, 0xe4, 0xa4, 0x4a, 0x30, 0x81, 0xb5, 0x41,
	0x38, 0x42, 0x62, 0x5c, 0x6c, 0x89, 0xa5, 0x25, 0x19, 0xf9, 0x45, 0x12, 0xcc, 0x60, 0x61, 0x28,
	0x0f, 0x66, 0xa2, 0x93, 0x73, 0x94, 0x23, 0x51, 0xae, 0xc6, 0x16, 0x1e, 0xd6, 0x50, 0xc1, 0x24,
	0x36, 0xb0, 0x6a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x19, 0x44, 0x3d, 0x36, 0x01,
	0x00, 0x00,
}
