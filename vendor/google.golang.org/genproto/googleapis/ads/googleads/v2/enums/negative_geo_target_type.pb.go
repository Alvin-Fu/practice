// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/negative_geo_target_type.proto

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

// The possible negative geo target types.
type NegativeGeoTargetTypeEnum_NegativeGeoTargetType int32

const (
	// Not specified.
	NegativeGeoTargetTypeEnum_UNSPECIFIED NegativeGeoTargetTypeEnum_NegativeGeoTargetType = 0
	// The value is unknown in this version.
	NegativeGeoTargetTypeEnum_UNKNOWN NegativeGeoTargetTypeEnum_NegativeGeoTargetType = 1
	// Specifies that a user is excluded from seeing the ad if they
	// are in, or show interest in, advertiser's excluded locations.
	NegativeGeoTargetTypeEnum_PRESENCE_OR_INTEREST NegativeGeoTargetTypeEnum_NegativeGeoTargetType = 4
	// Specifies that a user is excluded from seeing the ad if they
	// are in advertiser's excluded locations.
	NegativeGeoTargetTypeEnum_PRESENCE NegativeGeoTargetTypeEnum_NegativeGeoTargetType = 5
)

var NegativeGeoTargetTypeEnum_NegativeGeoTargetType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	4: "PRESENCE_OR_INTEREST",
	5: "PRESENCE",
}

var NegativeGeoTargetTypeEnum_NegativeGeoTargetType_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"PRESENCE_OR_INTEREST": 4,
	"PRESENCE":             5,
}

func (x NegativeGeoTargetTypeEnum_NegativeGeoTargetType) String() string {
	return proto.EnumName(NegativeGeoTargetTypeEnum_NegativeGeoTargetType_name, int32(x))
}

func (NegativeGeoTargetTypeEnum_NegativeGeoTargetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c8d8058d6d7174fa, []int{0, 0}
}

// Container for enum describing possible negative geo target types.
type NegativeGeoTargetTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NegativeGeoTargetTypeEnum) Reset()         { *m = NegativeGeoTargetTypeEnum{} }
func (m *NegativeGeoTargetTypeEnum) String() string { return proto.CompactTextString(m) }
func (*NegativeGeoTargetTypeEnum) ProtoMessage()    {}
func (*NegativeGeoTargetTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8d8058d6d7174fa, []int{0}
}

func (m *NegativeGeoTargetTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NegativeGeoTargetTypeEnum.Unmarshal(m, b)
}
func (m *NegativeGeoTargetTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NegativeGeoTargetTypeEnum.Marshal(b, m, deterministic)
}
func (m *NegativeGeoTargetTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NegativeGeoTargetTypeEnum.Merge(m, src)
}
func (m *NegativeGeoTargetTypeEnum) XXX_Size() int {
	return xxx_messageInfo_NegativeGeoTargetTypeEnum.Size(m)
}
func (m *NegativeGeoTargetTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_NegativeGeoTargetTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_NegativeGeoTargetTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.NegativeGeoTargetTypeEnum_NegativeGeoTargetType", NegativeGeoTargetTypeEnum_NegativeGeoTargetType_name, NegativeGeoTargetTypeEnum_NegativeGeoTargetType_value)
	proto.RegisterType((*NegativeGeoTargetTypeEnum)(nil), "google.ads.googleads.v2.enums.NegativeGeoTargetTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/negative_geo_target_type.proto", fileDescriptor_c8d8058d6d7174fa)
}

var fileDescriptor_c8d8058d6d7174fa = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdd, 0x4a, 0xc3, 0x30,
	0x1c, 0xc5, 0xdd, 0xfc, 0x24, 0x13, 0x2c, 0x45, 0x41, 0x87, 0xbb, 0xd8, 0x1e, 0x20, 0x81, 0x7a,
	0x17, 0xbd, 0xe9, 0x66, 0x1c, 0x43, 0xc8, 0xca, 0xd6, 0x4d, 0x90, 0x4a, 0x89, 0x36, 0x84, 0xc2,
	0x96, 0x94, 0x25, 0x1b, 0xcc, 0xc7, 0xf1, 0xd2, 0x47, 0xf1, 0x51, 0xbc, 0xf1, 0x15, 0xa4, 0xc9,
	0xba, 0xab, 0xe9, 0x4d, 0x39, 0xf4, 0xfc, 0x7f, 0x87, 0x93, 0x03, 0xee, 0x84, 0x52, 0x62, 0xc6,
	0x11, 0xcb, 0x34, 0x72, 0xb2, 0x54, 0xab, 0x00, 0x71, 0xb9, 0x9c, 0x6b, 0x24, 0xb9, 0x60, 0x26,
	0x5f, 0xf1, 0x54, 0x70, 0x95, 0x1a, 0xb6, 0x10, 0xdc, 0xa4, 0x66, 0x5d, 0x70, 0x58, 0x2c, 0x94,
	0x51, 0x7e, 0xcb, 0x21, 0x90, 0x65, 0x1a, 0x6e, 0x69, 0xb8, 0x0a, 0xa0, 0xa5, 0x9b, 0xd7, 0x55,
	0x78, 0x91, 0x23, 0x26, 0xa5, 0x32, 0xcc, 0xe4, 0x4a, 0x6a, 0x07, 0x77, 0xde, 0xc1, 0x15, 0xdd,
	0xc4, 0xf7, 0xb9, 0x8a, 0x6d, 0x78, 0xbc, 0x2e, 0x38, 0x91, 0xcb, 0x79, 0xe7, 0x05, 0x5c, 0xec,
	0x34, 0xfd, 0x33, 0xd0, 0x98, 0xd0, 0x71, 0x44, 0x7a, 0x83, 0x87, 0x01, 0xb9, 0xf7, 0xf6, 0xfc,
	0x06, 0x38, 0x9e, 0xd0, 0x47, 0x3a, 0x7c, 0xa2, 0x5e, 0xcd, 0xbf, 0x04, 0xe7, 0xd1, 0x88, 0x8c,
	0x09, 0xed, 0x91, 0x74, 0x38, 0x4a, 0x07, 0x34, 0x26, 0x23, 0x32, 0x8e, 0xbd, 0x03, 0xff, 0x14,
	0x9c, 0x54, 0x8e, 0x77, 0xd8, 0xfd, 0xa9, 0x81, 0xf6, 0x9b, 0x9a, 0xc3, 0x7f, 0xfb, 0x77, 0x9b,
	0x3b, 0x2b, 0x44, 0x65, 0xfb, 0xa8, 0xf6, 0xdc, 0xdd, 0xc0, 0x42, 0xcd, 0x98, 0x14, 0x50, 0x2d,
	0x04, 0x12, 0x5c, 0xda, 0xb7, 0x55, 0x53, 0x16, 0xb9, 0xfe, 0x63, 0xd9, 0x5b, 0xfb, 0xfd, 0xa8,
	0xef, 0xf7, 0xc3, 0xf0, 0xb3, 0xde, 0xea, 0xbb, 0xa8, 0x30, 0xd3, 0xd0, 0xc9, 0x52, 0x4d, 0x03,
	0x58, 0x6e, 0xa1, 0xbf, 0x2a, 0x3f, 0x09, 0x33, 0x9d, 0x6c, 0xfd, 0x64, 0x1a, 0x24, 0xd6, 0xff,
	0xae, 0xb7, 0xdd, 0x4f, 0x8c, 0xc3, 0x4c, 0x63, 0xbc, 0xbd, 0xc0, 0x78, 0x1a, 0x60, 0x6c, 0x6f,
	0x5e, 0x8f, 0x6c, 0xb1, 0x9b, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x0f, 0x86, 0x8a, 0xf1,
	0x01, 0x00, 0x00,
}
