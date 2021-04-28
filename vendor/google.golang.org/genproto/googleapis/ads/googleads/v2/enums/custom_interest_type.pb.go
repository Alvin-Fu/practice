// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/custom_interest_type.proto

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

// Enum containing possible custom interest types.
type CustomInterestTypeEnum_CustomInterestType int32

const (
	// Not specified.
	CustomInterestTypeEnum_UNSPECIFIED CustomInterestTypeEnum_CustomInterestType = 0
	// Used for return value only. Represents value unknown in this version.
	CustomInterestTypeEnum_UNKNOWN CustomInterestTypeEnum_CustomInterestType = 1
	// Allows brand advertisers to define custom affinity audience lists.
	CustomInterestTypeEnum_CUSTOM_AFFINITY CustomInterestTypeEnum_CustomInterestType = 2
	// Allows advertisers to define custom intent audience lists.
	CustomInterestTypeEnum_CUSTOM_INTENT CustomInterestTypeEnum_CustomInterestType = 3
)

var CustomInterestTypeEnum_CustomInterestType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CUSTOM_AFFINITY",
	3: "CUSTOM_INTENT",
}

var CustomInterestTypeEnum_CustomInterestType_value = map[string]int32{
	"UNSPECIFIED":     0,
	"UNKNOWN":         1,
	"CUSTOM_AFFINITY": 2,
	"CUSTOM_INTENT":   3,
}

func (x CustomInterestTypeEnum_CustomInterestType) String() string {
	return proto.EnumName(CustomInterestTypeEnum_CustomInterestType_name, int32(x))
}

func (CustomInterestTypeEnum_CustomInterestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2915d04628c70c4a, []int{0, 0}
}

// The types of custom interest.
type CustomInterestTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomInterestTypeEnum) Reset()         { *m = CustomInterestTypeEnum{} }
func (m *CustomInterestTypeEnum) String() string { return proto.CompactTextString(m) }
func (*CustomInterestTypeEnum) ProtoMessage()    {}
func (*CustomInterestTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_2915d04628c70c4a, []int{0}
}

func (m *CustomInterestTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomInterestTypeEnum.Unmarshal(m, b)
}
func (m *CustomInterestTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomInterestTypeEnum.Marshal(b, m, deterministic)
}
func (m *CustomInterestTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomInterestTypeEnum.Merge(m, src)
}
func (m *CustomInterestTypeEnum) XXX_Size() int {
	return xxx_messageInfo_CustomInterestTypeEnum.Size(m)
}
func (m *CustomInterestTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomInterestTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CustomInterestTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.CustomInterestTypeEnum_CustomInterestType", CustomInterestTypeEnum_CustomInterestType_name, CustomInterestTypeEnum_CustomInterestType_value)
	proto.RegisterType((*CustomInterestTypeEnum)(nil), "google.ads.googleads.v2.enums.CustomInterestTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/custom_interest_type.proto", fileDescriptor_2915d04628c70c4a)
}

var fileDescriptor_2915d04628c70c4a = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdf, 0x6a, 0xf2, 0x30,
	0x1c, 0xfd, 0xac, 0xf0, 0x0d, 0x22, 0xc3, 0xae, 0x83, 0x0d, 0xc6, 0xbc, 0xd0, 0x07, 0x48, 0xa0,
	0xbb, 0x19, 0xd9, 0x55, 0x74, 0x55, 0xc2, 0x58, 0x14, 0xac, 0x8e, 0x49, 0x41, 0x3a, 0x1b, 0x42,
	0xc1, 0x26, 0xc5, 0x44, 0xc1, 0xd7, 0xd9, 0xe5, 0x1e, 0x65, 0x8f, 0x32, 0xf6, 0x10, 0xa3, 0x89,
	0xf6, 0x46, 0xb6, 0x9b, 0x72, 0xf8, 0x9d, 0x3f, 0x3d, 0x39, 0xe0, 0x5e, 0x28, 0x25, 0xd6, 0x1c,
	0xa5, 0x99, 0x46, 0x0e, 0x56, 0x68, 0x17, 0x22, 0x2e, 0xb7, 0x85, 0x46, 0xab, 0xad, 0x36, 0xaa,
	0x58, 0xe6, 0xd2, 0xf0, 0x0d, 0xd7, 0x66, 0x69, 0xf6, 0x25, 0x87, 0xe5, 0x46, 0x19, 0x15, 0x74,
	0x9c, 0x1c, 0xa6, 0x99, 0x86, 0xb5, 0x13, 0xee, 0x42, 0x68, 0x9d, 0x37, 0xb7, 0xc7, 0xe0, 0x32,
	0x47, 0xa9, 0x94, 0xca, 0xa4, 0x26, 0x57, 0x52, 0x3b, 0x73, 0xcf, 0x80, 0xab, 0x81, 0x8d, 0xa6,
	0x87, 0xe4, 0x78, 0x5f, 0xf2, 0x48, 0x6e, 0x8b, 0xde, 0x02, 0x04, 0xa7, 0x4c, 0xd0, 0x06, 0xad,
	0x19, 0x9b, 0x4e, 0xa2, 0x01, 0x1d, 0xd2, 0xe8, 0xd1, 0xff, 0x17, 0xb4, 0xc0, 0xd9, 0x8c, 0x3d,
	0xb1, 0xf1, 0x0b, 0xf3, 0x1b, 0xc1, 0x25, 0x68, 0x0f, 0x66, 0xd3, 0x78, 0xfc, 0xbc, 0x24, 0xc3,
	0x21, 0x65, 0x34, 0x7e, 0xf5, 0xbd, 0xe0, 0x02, 0x9c, 0x1f, 0x8e, 0x94, 0xc5, 0x11, 0x8b, 0xfd,
	0x66, 0xff, 0xbb, 0x01, 0xba, 0x2b, 0x55, 0xc0, 0x3f, 0x9b, 0xf7, 0xaf, 0x4f, 0xff, 0x3f, 0xa9,
	0x4a, 0x4f, 0x1a, 0x8b, 0xfe, 0xc1, 0x29, 0xd4, 0x3a, 0x95, 0x02, 0xaa, 0x8d, 0x40, 0x82, 0x4b,
	0xfb, 0xa4, 0xe3, 0x7a, 0x65, 0xae, 0x7f, 0x19, 0xf3, 0xc1, 0x7e, 0xdf, 0xbd, 0xe6, 0x88, 0x90,
	0x0f, 0xaf, 0x33, 0x72, 0x51, 0x24, 0xd3, 0xd0, 0xc1, 0x0a, 0xcd, 0x43, 0x58, 0xad, 0xa0, 0x3f,
	0x8f, 0x7c, 0x42, 0x32, 0x9d, 0xd4, 0x7c, 0x32, 0x0f, 0x13, 0xcb, 0x7f, 0x79, 0x5d, 0x77, 0xc4,
	0x98, 0x64, 0x1a, 0xe3, 0x5a, 0x81, 0xf1, 0x3c, 0xc4, 0xd8, 0x6a, 0xde, 0xfe, 0xdb, 0x62, 0x77,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x13, 0x07, 0xd8, 0x06, 0xe4, 0x01, 0x00, 0x00,
}
