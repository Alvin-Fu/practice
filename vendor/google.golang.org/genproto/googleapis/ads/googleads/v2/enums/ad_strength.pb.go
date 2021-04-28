// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/ad_strength.proto

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

// Enum listing the possible ad strengths.
type AdStrengthEnum_AdStrength int32

const (
	// Not specified.
	AdStrengthEnum_UNSPECIFIED AdStrengthEnum_AdStrength = 0
	// Used for return value only. Represents value unknown in this version.
	AdStrengthEnum_UNKNOWN AdStrengthEnum_AdStrength = 1
	// The ad strength is currently pending.
	AdStrengthEnum_PENDING AdStrengthEnum_AdStrength = 2
	// No ads could be generated.
	AdStrengthEnum_NO_ADS AdStrengthEnum_AdStrength = 3
	// Poor strength.
	AdStrengthEnum_POOR AdStrengthEnum_AdStrength = 4
	// Average strength.
	AdStrengthEnum_AVERAGE AdStrengthEnum_AdStrength = 5
	// Good strength.
	AdStrengthEnum_GOOD AdStrengthEnum_AdStrength = 6
	// Excellent strength.
	AdStrengthEnum_EXCELLENT AdStrengthEnum_AdStrength = 7
)

var AdStrengthEnum_AdStrength_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PENDING",
	3: "NO_ADS",
	4: "POOR",
	5: "AVERAGE",
	6: "GOOD",
	7: "EXCELLENT",
}

var AdStrengthEnum_AdStrength_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"PENDING":     2,
	"NO_ADS":      3,
	"POOR":        4,
	"AVERAGE":     5,
	"GOOD":        6,
	"EXCELLENT":   7,
}

func (x AdStrengthEnum_AdStrength) String() string {
	return proto.EnumName(AdStrengthEnum_AdStrength_name, int32(x))
}

func (AdStrengthEnum_AdStrength) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_967c381b85f11def, []int{0, 0}
}

// Container for enum describing possible ad strengths.
type AdStrengthEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdStrengthEnum) Reset()         { *m = AdStrengthEnum{} }
func (m *AdStrengthEnum) String() string { return proto.CompactTextString(m) }
func (*AdStrengthEnum) ProtoMessage()    {}
func (*AdStrengthEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_967c381b85f11def, []int{0}
}

func (m *AdStrengthEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdStrengthEnum.Unmarshal(m, b)
}
func (m *AdStrengthEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdStrengthEnum.Marshal(b, m, deterministic)
}
func (m *AdStrengthEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdStrengthEnum.Merge(m, src)
}
func (m *AdStrengthEnum) XXX_Size() int {
	return xxx_messageInfo_AdStrengthEnum.Size(m)
}
func (m *AdStrengthEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdStrengthEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdStrengthEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.AdStrengthEnum_AdStrength", AdStrengthEnum_AdStrength_name, AdStrengthEnum_AdStrength_value)
	proto.RegisterType((*AdStrengthEnum)(nil), "google.ads.googleads.v2.enums.AdStrengthEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/ad_strength.proto", fileDescriptor_967c381b85f11def)
}

var fileDescriptor_967c381b85f11def = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdd, 0x4e, 0xc2, 0x30,
	0x18, 0x75, 0x03, 0x41, 0x4b, 0x94, 0x66, 0x97, 0x46, 0x2e, 0xe0, 0x01, 0xba, 0x64, 0xde, 0xd5,
	0xab, 0xc2, 0xea, 0x42, 0x24, 0xdd, 0x02, 0x32, 0x8d, 0x59, 0x42, 0xa6, 0x5d, 0x2a, 0x09, 0xb4,
	0x84, 0x0e, 0xde, 0xc0, 0x17, 0xf1, 0xd2, 0x47, 0xf1, 0x39, 0xbc, 0xf2, 0x29, 0x4c, 0x3b, 0x7e,
	0xae, 0xf4, 0xa6, 0x39, 0xfd, 0xbe, 0x73, 0x4e, 0xce, 0x77, 0x80, 0x2f, 0x94, 0x12, 0x8b, 0xc2,
	0xcf, 0xb9, 0xde, 0x41, 0x83, 0xb6, 0x81, 0x5f, 0xc8, 0xcd, 0x52, 0xfb, 0x39, 0x9f, 0xe9, 0x72,
	0x5d, 0x48, 0x51, 0xbe, 0xa1, 0xd5, 0x5a, 0x95, 0xca, 0xeb, 0x54, 0x2c, 0x94, 0x73, 0x8d, 0x0e,
	0x02, 0xb4, 0x0d, 0x90, 0x15, 0x5c, 0x5d, 0xef, 0xfd, 0x56, 0x73, 0x3f, 0x97, 0x52, 0x95, 0x79,
	0x39, 0x57, 0x52, 0x57, 0xe2, 0xde, 0xbb, 0x03, 0x2e, 0x09, 0x9f, 0xec, 0x1c, 0xa9, 0xdc, 0x2c,
	0x7b, 0x1a, 0x80, 0xe3, 0xc4, 0x6b, 0x83, 0xd6, 0x94, 0x4d, 0x12, 0x3a, 0x18, 0xde, 0x0d, 0x69,
	0x08, 0x4f, 0xbc, 0x16, 0x68, 0x4e, 0xd9, 0x3d, 0x8b, 0x1f, 0x19, 0x74, 0xcc, 0x27, 0xa1, 0x2c,
	0x1c, 0xb2, 0x08, 0xba, 0x1e, 0x00, 0x0d, 0x16, 0xcf, 0x48, 0x38, 0x81, 0x35, 0xef, 0x0c, 0xd4,
	0x93, 0x38, 0x1e, 0xc3, 0xba, 0xa1, 0x90, 0x94, 0x8e, 0x49, 0x44, 0xe1, 0xa9, 0x19, 0x47, 0x71,
	0x1c, 0xc2, 0x86, 0x77, 0x01, 0xce, 0xe9, 0xd3, 0x80, 0x8e, 0x46, 0x94, 0x3d, 0xc0, 0x66, 0xff,
	0xdb, 0x01, 0xdd, 0x57, 0xb5, 0x44, 0xff, 0xde, 0xd2, 0x6f, 0x1f, 0x83, 0x25, 0x26, 0x7e, 0xe2,
	0x3c, 0xf7, 0x77, 0x0a, 0xa1, 0x16, 0xb9, 0x14, 0x48, 0xad, 0x85, 0x2f, 0x0a, 0x69, 0x8f, 0xdb,
	0xd7, 0xb7, 0x9a, 0xeb, 0x3f, 0xda, 0xbc, 0xb5, 0xef, 0x87, 0x5b, 0x8b, 0x08, 0xf9, 0x74, 0x3b,
	0x51, 0x65, 0x45, 0xb8, 0x46, 0x15, 0x34, 0x28, 0x0d, 0x90, 0xa9, 0x45, 0x7f, 0xed, 0xf7, 0x19,
	0xe1, 0x3a, 0x3b, 0xec, 0xb3, 0x34, 0xc8, 0xec, 0xfe, 0xc7, 0xed, 0x56, 0x43, 0x8c, 0x09, 0xd7,
	0x18, 0x1f, 0x18, 0x18, 0xa7, 0x01, 0xc6, 0x96, 0xf3, 0xd2, 0xb0, 0xc1, 0x6e, 0x7e, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xa4, 0x2e, 0x5b, 0xbf, 0xe5, 0x01, 0x00, 0x00,
}
