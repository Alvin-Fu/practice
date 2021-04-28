// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/interaction_type.proto

package enums

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

// Enum describing possible interaction types.
type InteractionTypeEnum_InteractionType int32

const (
	// Not specified.
	InteractionTypeEnum_UNSPECIFIED InteractionTypeEnum_InteractionType = 0
	// Used for return value only. Represents value unknown in this version.
	InteractionTypeEnum_UNKNOWN InteractionTypeEnum_InteractionType = 1
	// Calls.
	InteractionTypeEnum_CALLS InteractionTypeEnum_InteractionType = 8000
)

var InteractionTypeEnum_InteractionType_name = map[int32]string{
	0:    "UNSPECIFIED",
	1:    "UNKNOWN",
	8000: "CALLS",
}

var InteractionTypeEnum_InteractionType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"CALLS":       8000,
}

func (x InteractionTypeEnum_InteractionType) String() string {
	return proto.EnumName(InteractionTypeEnum_InteractionType_name, int32(x))
}

func (InteractionTypeEnum_InteractionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_01158f56321f0e95, []int{0, 0}
}

// Container for enum describing possible interaction types.
type InteractionTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InteractionTypeEnum) Reset()         { *m = InteractionTypeEnum{} }
func (m *InteractionTypeEnum) String() string { return proto.CompactTextString(m) }
func (*InteractionTypeEnum) ProtoMessage()    {}
func (*InteractionTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_01158f56321f0e95, []int{0}
}

func (m *InteractionTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InteractionTypeEnum.Unmarshal(m, b)
}
func (m *InteractionTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InteractionTypeEnum.Marshal(b, m, deterministic)
}
func (m *InteractionTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InteractionTypeEnum.Merge(m, src)
}
func (m *InteractionTypeEnum) XXX_Size() int {
	return xxx_messageInfo_InteractionTypeEnum.Size(m)
}
func (m *InteractionTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_InteractionTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_InteractionTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.InteractionTypeEnum_InteractionType", InteractionTypeEnum_InteractionType_name, InteractionTypeEnum_InteractionType_value)
	proto.RegisterType((*InteractionTypeEnum)(nil), "google.ads.googleads.v2.enums.InteractionTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/interaction_type.proto", fileDescriptor_01158f56321f0e95)
}

var fileDescriptor_01158f56321f0e95 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x23, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xcc, 0xbc, 0x92, 0xd4, 0xa2, 0xc4, 0xe4, 0x92, 0xcc, 0xfc, 0xbc, 0xf8,
	0x92, 0xca, 0x82, 0x54, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x52, 0xbd, 0xc4,
	0x94, 0x62, 0x3d, 0xb8, 0x2e, 0xbd, 0x32, 0x23, 0x3d, 0xb0, 0x2e, 0x29, 0x19, 0x98, 0xa1, 0x05,
	0x99, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x20, 0x03, 0x8a, 0x21, 0x9a, 0x95, 0x82, 0xb8,
	0x84, 0x3d, 0x11, 0xc6, 0x86, 0x54, 0x16, 0xa4, 0xba, 0xe6, 0x95, 0xe6, 0x2a, 0x59, 0x73, 0xf1,
	0xa3, 0x09, 0x0b, 0xf1, 0x73, 0x71, 0x87, 0xfa, 0x05, 0x07, 0xb8, 0x3a, 0x7b, 0xba, 0x79, 0xba,
	0xba, 0x08, 0x30, 0x08, 0x71, 0x73, 0xb1, 0x87, 0xfa, 0x79, 0xfb, 0xf9, 0x87, 0xfb, 0x09, 0x30,
	0x0a, 0x71, 0x71, 0xb1, 0x3a, 0x3b, 0xfa, 0xf8, 0x04, 0x0b, 0x1c, 0xb0, 0x73, 0x7a, 0xc9, 0xc8,
	0xa5, 0x98, 0x9c, 0x9f, 0xab, 0x87, 0xd7, 0x5d, 0x4e, 0x22, 0x68, 0x16, 0x04, 0x80, 0xdc, 0x13,
	0xc0, 0x18, 0xe5, 0x04, 0xd5, 0x96, 0x9e, 0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f, 0x94, 0xae,
	0x9f, 0x9e, 0x9a, 0x07, 0x76, 0x2d, 0x2c, 0x50, 0x0a, 0x32, 0x8b, 0x71, 0x84, 0x91, 0x35, 0x98,
	0x5c, 0xc4, 0xc4, 0xec, 0xee, 0xe8, 0xb8, 0x8a, 0x49, 0xd6, 0x1d, 0x62, 0x94, 0x63, 0x4a, 0xb1,
	0x1e, 0x84, 0x09, 0x62, 0x85, 0x19, 0xe9, 0x81, 0xfc, 0x58, 0x7c, 0x0a, 0x26, 0x1f, 0xe3, 0x98,
	0x52, 0x1c, 0x03, 0x97, 0x8f, 0x09, 0x33, 0x8a, 0x01, 0xcb, 0xbf, 0x62, 0x52, 0x84, 0x08, 0x5a,
	0x59, 0x39, 0xa6, 0x14, 0x5b, 0x59, 0xc1, 0x55, 0x58, 0x59, 0x85, 0x19, 0x59, 0x59, 0x81, 0xd5,
	0x24, 0xb1, 0x81, 0x1d, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x38, 0x83, 0x9e, 0xf8, 0xbb,
	0x01, 0x00, 0x00,
}
