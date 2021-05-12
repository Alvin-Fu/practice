// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/affiliate_location_placeholder_field.proto

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

// Possible values for Affiliate Location placeholder fields.
type AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField int32

const (
	// Not specified.
	AffiliateLocationPlaceholderFieldEnum_UNSPECIFIED AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	AffiliateLocationPlaceholderFieldEnum_UNKNOWN AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField = 1
	// Data Type: STRING. The name of the business.
	AffiliateLocationPlaceholderFieldEnum_BUSINESS_NAME AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField = 2
	// Data Type: STRING. Line 1 of the business address.
	AffiliateLocationPlaceholderFieldEnum_ADDRESS_LINE_1 AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField = 3
	// Data Type: STRING. Line 2 of the business address.
	AffiliateLocationPlaceholderFieldEnum_ADDRESS_LINE_2 AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField = 4
	// Data Type: STRING. City of the business address.
	AffiliateLocationPlaceholderFieldEnum_CITY AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField = 5
	// Data Type: STRING. Province of the business address.
	AffiliateLocationPlaceholderFieldEnum_PROVINCE AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField = 6
	// Data Type: STRING. Postal code of the business address.
	AffiliateLocationPlaceholderFieldEnum_POSTAL_CODE AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField = 7
	// Data Type: STRING. Country code of the business address.
	AffiliateLocationPlaceholderFieldEnum_COUNTRY_CODE AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField = 8
	// Data Type: STRING. Phone number of the business.
	AffiliateLocationPlaceholderFieldEnum_PHONE_NUMBER AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField = 9
	// Data Type: STRING. Language code of the business.
	AffiliateLocationPlaceholderFieldEnum_LANGUAGE_CODE AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField = 10
	// Data Type: INT64. ID of the chain.
	AffiliateLocationPlaceholderFieldEnum_CHAIN_ID AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField = 11
	// Data Type: STRING. Name of the chain.
	AffiliateLocationPlaceholderFieldEnum_CHAIN_NAME AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField = 12
)

var AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "BUSINESS_NAME",
	3:  "ADDRESS_LINE_1",
	4:  "ADDRESS_LINE_2",
	5:  "CITY",
	6:  "PROVINCE",
	7:  "POSTAL_CODE",
	8:  "COUNTRY_CODE",
	9:  "PHONE_NUMBER",
	10: "LANGUAGE_CODE",
	11: "CHAIN_ID",
	12: "CHAIN_NAME",
}

var AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"BUSINESS_NAME":  2,
	"ADDRESS_LINE_1": 3,
	"ADDRESS_LINE_2": 4,
	"CITY":           5,
	"PROVINCE":       6,
	"POSTAL_CODE":    7,
	"COUNTRY_CODE":   8,
	"PHONE_NUMBER":   9,
	"LANGUAGE_CODE":  10,
	"CHAIN_ID":       11,
	"CHAIN_NAME":     12,
}

func (x AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField) String() string {
	return proto.EnumName(AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField_name, int32(x))
}

func (AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00ea8dfbc7e37169, []int{0, 0}
}

// Values for Affiliate Location placeholder fields.
type AffiliateLocationPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AffiliateLocationPlaceholderFieldEnum) Reset()         { *m = AffiliateLocationPlaceholderFieldEnum{} }
func (m *AffiliateLocationPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*AffiliateLocationPlaceholderFieldEnum) ProtoMessage()    {}
func (*AffiliateLocationPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_00ea8dfbc7e37169, []int{0}
}

func (m *AffiliateLocationPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AffiliateLocationPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *AffiliateLocationPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AffiliateLocationPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (m *AffiliateLocationPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AffiliateLocationPlaceholderFieldEnum.Merge(m, src)
}
func (m *AffiliateLocationPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_AffiliateLocationPlaceholderFieldEnum.Size(m)
}
func (m *AffiliateLocationPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AffiliateLocationPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AffiliateLocationPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField", AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField_name, AffiliateLocationPlaceholderFieldEnum_AffiliateLocationPlaceholderField_value)
	proto.RegisterType((*AffiliateLocationPlaceholderFieldEnum)(nil), "google.ads.googleads.v1.enums.AffiliateLocationPlaceholderFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/affiliate_location_placeholder_field.proto", fileDescriptor_00ea8dfbc7e37169)
}

var fileDescriptor_00ea8dfbc7e37169 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0xa5, 0xd9, 0x65, 0xb7, 0xb8, 0x65, 0x31, 0x3e, 0x22, 0xf6, 0xd0, 0x95, 0xe0, 0xe8, 0x28,
	0x70, 0x0b, 0x27, 0x27, 0xf1, 0xb6, 0x11, 0x5d, 0x27, 0x6a, 0x9a, 0xa2, 0x45, 0x95, 0x2c, 0xd3,
	0xa4, 0x21, 0x52, 0x1a, 0x57, 0x75, 0x77, 0x3f, 0x08, 0x71, 0xe2, 0x53, 0xf8, 0x14, 0x3e, 0x01,
	0x2e, 0xc8, 0xf6, 0xa6, 0x1c, 0x10, 0xf4, 0x12, 0xbd, 0xbc, 0x79, 0x7e, 0x6f, 0x34, 0x33, 0x60,
	0x52, 0x49, 0x59, 0x35, 0xa5, 0x2b, 0x0a, 0xe5, 0x5a, 0xa8, 0xd1, 0xbd, 0xe7, 0x96, 0xed, 0xdd,
	0x46, 0xb9, 0x62, 0xbd, 0xae, 0x9b, 0x5a, 0xec, 0x4b, 0xde, 0xc8, 0x95, 0xd8, 0xd7, 0xb2, 0xe5,
	0xdb, 0x46, 0xac, 0xca, 0xcf, 0xb2, 0x29, 0xca, 0x1d, 0x5f, 0xd7, 0x65, 0x53, 0xe0, 0xed, 0x4e,
	0xee, 0x25, 0xba, 0xb4, 0xcf, 0xb1, 0x28, 0x14, 0x3e, 0x38, 0xe1, 0x7b, 0x0f, 0x1b, 0xa7, 0x17,
	0x2f, 0xbb, 0xa0, 0x6d, 0xed, 0x8a, 0xb6, 0x95, 0x7b, 0xe3, 0xa6, 0xec, 0xe3, 0xab, 0xaf, 0x0e,
	0x78, 0x45, 0xba, 0xac, 0xe9, 0x43, 0x54, 0xfa, 0x27, 0xe9, 0x5a, 0x07, 0xd1, 0xf6, 0x6e, 0x73,
	0xf5, 0xb3, 0x07, 0x46, 0x47, 0x95, 0xe8, 0x19, 0x18, 0xe4, 0x2c, 0x4b, 0x69, 0x18, 0x5f, 0xc7,
	0x34, 0x82, 0x8f, 0xd0, 0x00, 0x9c, 0xe7, 0xec, 0x3d, 0x4b, 0x3e, 0x30, 0xd8, 0x43, 0xcf, 0xc1,
	0xd3, 0x20, 0xcf, 0x62, 0x46, 0xb3, 0x8c, 0x33, 0x72, 0x43, 0xa1, 0x83, 0x10, 0xb8, 0x20, 0x51,
	0x34, 0xd3, 0xcc, 0x34, 0x66, 0x94, 0x7b, 0xf0, 0xe4, 0x2f, 0xee, 0x0d, 0x3c, 0x45, 0x7d, 0x70,
	0x1a, 0xc6, 0xf3, 0x5b, 0xf8, 0x18, 0x0d, 0x41, 0x3f, 0x9d, 0x25, 0x8b, 0x98, 0x85, 0x14, 0x9e,
	0xe9, 0xc0, 0x34, 0xc9, 0xe6, 0x64, 0xca, 0xc3, 0x24, 0xa2, 0xf0, 0x1c, 0x41, 0x30, 0x0c, 0x93,
	0x9c, 0xcd, 0x67, 0xb7, 0x96, 0xe9, 0x6b, 0x26, 0x9d, 0x24, 0x8c, 0x72, 0x96, 0xdf, 0x04, 0x74,
	0x06, 0x9f, 0xe8, 0x3e, 0xa6, 0x84, 0x8d, 0x73, 0x32, 0xa6, 0x56, 0x04, 0xb4, 0x6b, 0x38, 0x21,
	0x31, 0xe3, 0x71, 0x04, 0x07, 0xe8, 0x02, 0x00, 0xfb, 0x67, 0xba, 0x1c, 0x06, 0xbf, 0x7a, 0x60,
	0xb4, 0x92, 0x1b, 0xfc, 0xdf, 0x51, 0x07, 0xaf, 0x8f, 0xce, 0x27, 0xd5, 0x43, 0x4f, 0x7b, 0x1f,
	0x83, 0x07, 0xa3, 0x4a, 0x36, 0xa2, 0xad, 0xb0, 0xdc, 0x55, 0x6e, 0x55, 0xb6, 0x66, 0x25, 0xdd,
	0x35, 0x6c, 0x6b, 0xf5, 0x8f, 0xe3, 0x78, 0x67, 0xbe, 0x5f, 0x9c, 0x93, 0x31, 0x21, 0xdf, 0x9c,
	0xcb, 0xb1, 0xb5, 0x22, 0x85, 0xc2, 0x16, 0x6a, 0xb4, 0xf0, 0xb0, 0xde, 0x9a, 0xfa, 0xde, 0xd5,
	0x97, 0xa4, 0x50, 0xcb, 0x43, 0x7d, 0xb9, 0xf0, 0x96, 0xa6, 0xfe, 0xc3, 0x19, 0x59, 0xd2, 0xf7,
	0x49, 0xa1, 0x7c, 0xff, 0xa0, 0xf0, 0xfd, 0x85, 0xe7, 0xfb, 0x46, 0xf3, 0xe9, 0xcc, 0x34, 0xf6,
	0xf6, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x18, 0x25, 0xe9, 0x48, 0xb4, 0x02, 0x00, 0x00,
}
