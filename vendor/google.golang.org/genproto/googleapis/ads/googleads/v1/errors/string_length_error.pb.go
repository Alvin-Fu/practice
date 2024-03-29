// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/string_length_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible string length errors.
type StringLengthErrorEnum_StringLengthError int32

const (
	// Enum unspecified.
	StringLengthErrorEnum_UNSPECIFIED StringLengthErrorEnum_StringLengthError = 0
	// The received error code is not known in this version.
	StringLengthErrorEnum_UNKNOWN StringLengthErrorEnum_StringLengthError = 1
	// Too short.
	StringLengthErrorEnum_TOO_SHORT StringLengthErrorEnum_StringLengthError = 2
	// Too long.
	StringLengthErrorEnum_TOO_LONG StringLengthErrorEnum_StringLengthError = 3
)

var StringLengthErrorEnum_StringLengthError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "TOO_SHORT",
	3: "TOO_LONG",
}

var StringLengthErrorEnum_StringLengthError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"TOO_SHORT":   2,
	"TOO_LONG":    3,
}

func (x StringLengthErrorEnum_StringLengthError) String() string {
	return proto.EnumName(StringLengthErrorEnum_StringLengthError_name, int32(x))
}

func (StringLengthErrorEnum_StringLengthError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cc7d1cf13cf6443b, []int{0, 0}
}

// Container for enum describing possible string length errors.
type StringLengthErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringLengthErrorEnum) Reset()         { *m = StringLengthErrorEnum{} }
func (m *StringLengthErrorEnum) String() string { return proto.CompactTextString(m) }
func (*StringLengthErrorEnum) ProtoMessage()    {}
func (*StringLengthErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc7d1cf13cf6443b, []int{0}
}

func (m *StringLengthErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringLengthErrorEnum.Unmarshal(m, b)
}
func (m *StringLengthErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringLengthErrorEnum.Marshal(b, m, deterministic)
}
func (m *StringLengthErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringLengthErrorEnum.Merge(m, src)
}
func (m *StringLengthErrorEnum) XXX_Size() int {
	return xxx_messageInfo_StringLengthErrorEnum.Size(m)
}
func (m *StringLengthErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_StringLengthErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_StringLengthErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.StringLengthErrorEnum_StringLengthError", StringLengthErrorEnum_StringLengthError_name, StringLengthErrorEnum_StringLengthError_value)
	proto.RegisterType((*StringLengthErrorEnum)(nil), "google.ads.googleads.v1.errors.StringLengthErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/string_length_error.proto", fileDescriptor_cc7d1cf13cf6443b)
}

var fileDescriptor_cc7d1cf13cf6443b = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4e, 0x84, 0x30,
	0x14, 0xc6, 0x85, 0x49, 0xfc, 0xd3, 0xd1, 0x88, 0x24, 0xba, 0x30, 0x66, 0x16, 0x1c, 0xa0, 0x0d,
	0x71, 0x63, 0xea, 0x8a, 0x71, 0x10, 0x27, 0x4e, 0x0a, 0x91, 0x19, 0x4c, 0x0c, 0x09, 0x41, 0x21,
	0x95, 0x84, 0x69, 0x49, 0x8b, 0x73, 0x20, 0x97, 0x1e, 0xc5, 0xa3, 0xb8, 0xf3, 0x06, 0x86, 0x56,
	0xd8, 0x4c, 0x74, 0xd5, 0xaf, 0xaf, 0xbf, 0xef, 0x7b, 0xaf, 0x0f, 0x5c, 0x51, 0xce, 0x69, 0x5d,
	0xa2, 0xbc, 0x90, 0x48, 0xcb, 0x4e, 0x6d, 0x5c, 0x54, 0x0a, 0xc1, 0x85, 0x44, 0xb2, 0x15, 0x15,
	0xa3, 0x59, 0x5d, 0x32, 0xda, 0xbe, 0x66, 0xaa, 0x08, 0x1b, 0xc1, 0x5b, 0x6e, 0x4f, 0x34, 0x0e,
	0xf3, 0x42, 0xc2, 0xc1, 0x09, 0x37, 0x2e, 0xd4, 0xce, 0xf3, 0x8b, 0x3e, 0xb9, 0xa9, 0x50, 0xce,
	0x18, 0x6f, 0xf3, 0xb6, 0xe2, 0x4c, 0x6a, 0xb7, 0x43, 0xc1, 0x69, 0xac, 0xa2, 0x17, 0x2a, 0xd9,
	0xef, 0x3c, 0x3e, 0x7b, 0x5b, 0x3b, 0x04, 0x9c, 0x6c, 0x3d, 0xd8, 0xc7, 0x60, 0xbc, 0x22, 0x71,
	0xe4, 0xdf, 0xcc, 0x6f, 0xe7, 0xfe, 0xcc, 0xda, 0xb1, 0xc7, 0x60, 0x6f, 0x45, 0xee, 0x49, 0xf8,
	0x48, 0x2c, 0xc3, 0x3e, 0x02, 0x07, 0xcb, 0x30, 0xcc, 0xe2, 0xbb, 0xf0, 0x61, 0x69, 0x99, 0xf6,
	0x21, 0xd8, 0xef, 0xae, 0x8b, 0x90, 0x04, 0xd6, 0x68, 0xfa, 0x6d, 0x00, 0xe7, 0x85, 0xaf, 0xe1,
	0xff, 0xd3, 0x4e, 0xcf, 0xb6, 0x9a, 0x46, 0xdd, 0x9c, 0x91, 0xf1, 0x34, 0xfb, 0x75, 0x52, 0x5e,
	0xe7, 0x8c, 0x42, 0x2e, 0x28, 0xa2, 0x25, 0x53, 0xbf, 0xe8, 0x37, 0xd6, 0x54, 0xf2, 0xaf, 0x05,
	0x5e, 0xeb, 0xe3, 0xdd, 0x1c, 0x05, 0x9e, 0xf7, 0x61, 0x4e, 0x02, 0x1d, 0xe6, 0x15, 0x12, 0x6a,
	0xd9, 0xa9, 0xc4, 0x85, 0xaa, 0xa5, 0xfc, 0xec, 0x81, 0xd4, 0x2b, 0x64, 0x3a, 0x00, 0x69, 0xe2,
	0xa6, 0x1a, 0xf8, 0x32, 0x1d, 0x5d, 0xc5, 0xd8, 0x2b, 0x24, 0xc6, 0x03, 0x82, 0x71, 0xe2, 0x62,
	0xac, 0xa1, 0xe7, 0x5d, 0x35, 0xdd, 0xe5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x64, 0x6d, 0x29,
	0x4e, 0xdd, 0x01, 0x00, 0x00,
}
