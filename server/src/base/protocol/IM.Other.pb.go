// Code generated by protoc-gen-go. DO NOT EDIT.
// source: IM.Other.proto

package IM_Other

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type IMHeartBeat struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMHeartBeat) Reset()         { *m = IMHeartBeat{} }
func (m *IMHeartBeat) String() string { return proto.CompactTextString(m) }
func (*IMHeartBeat) ProtoMessage()    {}
func (*IMHeartBeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bd32544ae6b3c91, []int{0}
}

func (m *IMHeartBeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMHeartBeat.Unmarshal(m, b)
}
func (m *IMHeartBeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMHeartBeat.Marshal(b, m, deterministic)
}
func (m *IMHeartBeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMHeartBeat.Merge(m, src)
}
func (m *IMHeartBeat) XXX_Size() int {
	return xxx_messageInfo_IMHeartBeat.Size(m)
}
func (m *IMHeartBeat) XXX_DiscardUnknown() {
	xxx_messageInfo_IMHeartBeat.DiscardUnknown(m)
}

var xxx_messageInfo_IMHeartBeat proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IMHeartBeat)(nil), "IM.Other.IMHeartBeat")
}

func init() { proto.RegisterFile("IM.Other.proto", fileDescriptor_6bd32544ae6b3c91) }

var fileDescriptor_6bd32544ae6b3c91 = []byte{
	// 87 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xf3, 0xf4, 0xd5, 0xf3,
	0x2f, 0xc9, 0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x78,
	0xb9, 0xb8, 0x3d, 0x7d, 0x3d, 0x52, 0x13, 0x8b, 0x4a, 0x9c, 0x52, 0x13, 0x4b, 0x9c, 0x24, 0xb9,
	0xc4, 0x93, 0xf3, 0x73, 0xf5, 0x72, 0xf3, 0xd3, 0x4b, 0xb3, 0x32, 0x53, 0xf5, 0x4a, 0x4a, 0x20,
	0x1a, 0x92, 0x4a, 0xd3, 0x92, 0xd8, 0xc0, 0x2c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95,
	0x69, 0xa2, 0x88, 0x4c, 0x00, 0x00, 0x00,
}
