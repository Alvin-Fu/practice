// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/mobile_device_type.proto

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

// The type of mobile device.
type MobileDeviceTypeEnum_MobileDeviceType int32

const (
	// Not specified.
	MobileDeviceTypeEnum_UNSPECIFIED MobileDeviceTypeEnum_MobileDeviceType = 0
	// Used for return value only. Represents value unknown in this version.
	MobileDeviceTypeEnum_UNKNOWN MobileDeviceTypeEnum_MobileDeviceType = 1
	// Mobile phones.
	MobileDeviceTypeEnum_MOBILE MobileDeviceTypeEnum_MobileDeviceType = 2
	// Tablets.
	MobileDeviceTypeEnum_TABLET MobileDeviceTypeEnum_MobileDeviceType = 3
)

var MobileDeviceTypeEnum_MobileDeviceType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "MOBILE",
	3: "TABLET",
}

var MobileDeviceTypeEnum_MobileDeviceType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"MOBILE":      2,
	"TABLET":      3,
}

func (x MobileDeviceTypeEnum_MobileDeviceType) String() string {
	return proto.EnumName(MobileDeviceTypeEnum_MobileDeviceType_name, int32(x))
}

func (MobileDeviceTypeEnum_MobileDeviceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5e403a42cd2383c3, []int{0, 0}
}

// Container for enum describing the types of mobile device.
type MobileDeviceTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MobileDeviceTypeEnum) Reset()         { *m = MobileDeviceTypeEnum{} }
func (m *MobileDeviceTypeEnum) String() string { return proto.CompactTextString(m) }
func (*MobileDeviceTypeEnum) ProtoMessage()    {}
func (*MobileDeviceTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e403a42cd2383c3, []int{0}
}

func (m *MobileDeviceTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileDeviceTypeEnum.Unmarshal(m, b)
}
func (m *MobileDeviceTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileDeviceTypeEnum.Marshal(b, m, deterministic)
}
func (m *MobileDeviceTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileDeviceTypeEnum.Merge(m, src)
}
func (m *MobileDeviceTypeEnum) XXX_Size() int {
	return xxx_messageInfo_MobileDeviceTypeEnum.Size(m)
}
func (m *MobileDeviceTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileDeviceTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MobileDeviceTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.MobileDeviceTypeEnum_MobileDeviceType", MobileDeviceTypeEnum_MobileDeviceType_name, MobileDeviceTypeEnum_MobileDeviceType_value)
	proto.RegisterType((*MobileDeviceTypeEnum)(nil), "google.ads.googleads.v2.enums.MobileDeviceTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/mobile_device_type.proto", fileDescriptor_5e403a42cd2383c3)
}

var fileDescriptor_5e403a42cd2383c3 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xfe, 0xad, 0x83, 0xfd, 0x20, 0x3b, 0x58, 0x8a, 0x5e, 0xc4, 0x1d, 0xb6, 0x0f, 0x90, 0x40,
	0x05, 0x0f, 0xf1, 0x94, 0xba, 0x3a, 0x87, 0x5b, 0x57, 0xb0, 0xab, 0x20, 0x85, 0xd9, 0x2d, 0x21,
	0x14, 0xda, 0xa4, 0x2c, 0x5d, 0x61, 0x5f, 0xc7, 0xa3, 0x1f, 0xc5, 0x8f, 0xb2, 0x4f, 0x21, 0x4d,
	0x6c, 0x0f, 0x03, 0xbd, 0x84, 0x87, 0xf7, 0xf9, 0x93, 0xe7, 0x7d, 0xc1, 0x1d, 0x97, 0x92, 0xe7,
	0x0c, 0xa5, 0x54, 0x21, 0x03, 0x1b, 0x54, 0xbb, 0x88, 0x89, 0x43, 0xa1, 0x50, 0x21, 0xb7, 0x59,
	0xce, 0x36, 0x94, 0xd5, 0xd9, 0x8e, 0x6d, 0xaa, 0x63, 0xc9, 0x60, 0xb9, 0x97, 0x95, 0x74, 0x46,
	0x46, 0x0c, 0x53, 0xaa, 0x60, 0xe7, 0x83, 0xb5, 0x0b, 0xb5, 0xef, 0xfa, 0xa6, 0x8d, 0x2d, 0x33,
	0x94, 0x0a, 0x21, 0xab, 0xb4, 0xca, 0xa4, 0x50, 0xc6, 0x3c, 0x79, 0x07, 0x97, 0x4b, 0x1d, 0x3c,
	0xd5, 0xb9, 0xd1, 0xb1, 0x64, 0xbe, 0x38, 0x14, 0x93, 0x27, 0x60, 0x9f, 0xcf, 0x9d, 0x0b, 0x30,
	0x5c, 0x07, 0x2f, 0xa1, 0xff, 0x30, 0x7f, 0x9c, 0xfb, 0x53, 0xfb, 0x9f, 0x33, 0x04, 0xff, 0xd7,
	0xc1, 0x73, 0xb0, 0x7a, 0x0d, 0xec, 0x9e, 0x03, 0xc0, 0x60, 0xb9, 0xf2, 0xe6, 0x0b, 0xdf, 0xb6,
	0x1a, 0x1c, 0x11, 0x6f, 0xe1, 0x47, 0x76, 0xdf, 0x3b, 0xf5, 0xc0, 0x78, 0x27, 0x0b, 0xf8, 0x67,
	0x4b, 0xef, 0xea, 0xfc, 0xb7, 0xb0, 0xa9, 0x17, 0xf6, 0xde, 0xbc, 0x1f, 0x1f, 0x97, 0x79, 0x2a,
	0x38, 0x94, 0x7b, 0x8e, 0x38, 0x13, 0xba, 0x7c, 0x7b, 0xa5, 0x32, 0x53, 0xbf, 0x1c, 0xed, 0x5e,
	0xbf, 0x1f, 0x56, 0x7f, 0x46, 0xc8, 0xa7, 0x35, 0x9a, 0x99, 0x28, 0x42, 0x15, 0x34, 0xb0, 0x41,
	0xb1, 0x0b, 0x9b, 0x8d, 0xd5, 0x57, 0xcb, 0x27, 0x84, 0xaa, 0xa4, 0xe3, 0x93, 0xd8, 0x4d, 0x34,
	0x7f, 0xb2, 0xc6, 0x66, 0x88, 0x31, 0xa1, 0x0a, 0xe3, 0x4e, 0x81, 0x71, 0xec, 0x62, 0xac, 0x35,
	0xdb, 0x81, 0x2e, 0x76, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x39, 0xe7, 0x2f, 0x43, 0xcc, 0x01,
	0x00, 0x00,
}
