// Code generated by protoc-gen-go. DO NOT EDIT.
// source: TvAuthenticationService.proto

package TvAuthenticationService

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

type TargetTv struct {
	Emac                 string   `protobuf:"bytes,1,opt,name=emac,proto3" json:"emac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TargetTv) Reset()         { *m = TargetTv{} }
func (m *TargetTv) String() string { return proto.CompactTextString(m) }
func (*TargetTv) ProtoMessage()    {}
func (*TargetTv) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9ed2356ff4a0833, []int{0}
}

func (m *TargetTv) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetTv.Unmarshal(m, b)
}
func (m *TargetTv) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetTv.Marshal(b, m, deterministic)
}
func (m *TargetTv) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetTv.Merge(m, src)
}
func (m *TargetTv) XXX_Size() int {
	return xxx_messageInfo_TargetTv.Size(m)
}
func (m *TargetTv) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetTv.DiscardUnknown(m)
}

var xxx_messageInfo_TargetTv proto.InternalMessageInfo

func (m *TargetTv) GetEmac() string {
	if m != nil {
		return m.Emac
	}
	return ""
}

type IsAuthorised struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAuthorised) Reset()         { *m = IsAuthorised{} }
func (m *IsAuthorised) String() string { return proto.CompactTextString(m) }
func (*IsAuthorised) ProtoMessage()    {}
func (*IsAuthorised) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9ed2356ff4a0833, []int{1}
}

func (m *IsAuthorised) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorised.Unmarshal(m, b)
}
func (m *IsAuthorised) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorised.Marshal(b, m, deterministic)
}
func (m *IsAuthorised) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorised.Merge(m, src)
}
func (m *IsAuthorised) XXX_Size() int {
	return xxx_messageInfo_IsAuthorised.Size(m)
}
func (m *IsAuthorised) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorised.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorised proto.InternalMessageInfo

func (m *IsAuthorised) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*TargetTv)(nil), "TvAuthenticationService.TargetTv")
	proto.RegisterType((*IsAuthorised)(nil), "TvAuthenticationService.IsAuthorised")
}

func init() { proto.RegisterFile("TvAuthenticationService.proto", fileDescriptor_c9ed2356ff4a0833) }

var fileDescriptor_c9ed2356ff4a0833 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x0d, 0x29, 0x73, 0x2c,
	0x2d, 0xc9, 0x48, 0xcd, 0x2b, 0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x0b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc7, 0x21, 0xad, 0x24, 0xc7,
	0xc5, 0x11, 0x92, 0x58, 0x94, 0x9e, 0x5a, 0x12, 0x52, 0x26, 0x24, 0xc4, 0xc5, 0x92, 0x9a, 0x9b,
	0x98, 0x2c, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x2b, 0xa9, 0x71, 0xf1, 0x78, 0x16,
	0x83, 0xb4, 0xe6, 0x17, 0x65, 0x16, 0xa7, 0xa6, 0x08, 0x89, 0x71, 0xb1, 0x15, 0xa5, 0x16, 0x97,
	0xe6, 0x94, 0x80, 0x55, 0x71, 0x04, 0x41, 0x79, 0x46, 0x55, 0x5c, 0xb8, 0xac, 0x10, 0x8a, 0xe7,
	0x12, 0x76, 0xce, 0xc9, 0x2f, 0x4d, 0x29, 0x4f, 0xcc, 0xc9, 0x4e, 0x2d, 0x02, 0xa9, 0x09, 0x4f,
	0xcc, 0xc9, 0x11, 0x52, 0xd4, 0xc3, 0xe5, 0x64, 0x98, 0x83, 0xa4, 0x54, 0x71, 0x2a, 0x41, 0x76,
	0x53, 0x12, 0x1b, 0xd8, 0x8f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x56, 0xd0, 0xc5,
	0x04, 0x01, 0x00, 0x00,
}
