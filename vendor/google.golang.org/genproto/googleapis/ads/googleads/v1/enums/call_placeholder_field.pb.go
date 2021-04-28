// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/call_placeholder_field.proto

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

// Possible values for Call placeholder fields.
type CallPlaceholderFieldEnum_CallPlaceholderField int32

const (
	// Not specified.
	CallPlaceholderFieldEnum_UNSPECIFIED CallPlaceholderFieldEnum_CallPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	CallPlaceholderFieldEnum_UNKNOWN CallPlaceholderFieldEnum_CallPlaceholderField = 1
	// Data Type: STRING. The advertiser's phone number to append to the ad.
	CallPlaceholderFieldEnum_PHONE_NUMBER CallPlaceholderFieldEnum_CallPlaceholderField = 2
	// Data Type: STRING. Uppercase two-letter country code of the advertiser's
	// phone number.
	CallPlaceholderFieldEnum_COUNTRY_CODE CallPlaceholderFieldEnum_CallPlaceholderField = 3
	// Data Type: BOOLEAN. Indicates whether call tracking is enabled. Default:
	// true.
	CallPlaceholderFieldEnum_TRACKED CallPlaceholderFieldEnum_CallPlaceholderField = 4
	// Data Type: INT64. The ID of an AdCallMetricsConversion object. This
	// object contains the phoneCallDurationfield which is the minimum duration
	// (in seconds) of a call to be considered a conversion.
	CallPlaceholderFieldEnum_CONVERSION_TYPE_ID CallPlaceholderFieldEnum_CallPlaceholderField = 5
	// Data Type: STRING. Indicates whether this call extension uses its own
	// call conversion setting or follows the account level setting.
	// Valid values are: USE_ACCOUNT_LEVEL_CALL_CONVERSION_ACTION and
	// USE_RESOURCE_LEVEL_CALL_CONVERSION_ACTION.
	CallPlaceholderFieldEnum_CONVERSION_REPORTING_STATE CallPlaceholderFieldEnum_CallPlaceholderField = 6
)

var CallPlaceholderFieldEnum_CallPlaceholderField_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PHONE_NUMBER",
	3: "COUNTRY_CODE",
	4: "TRACKED",
	5: "CONVERSION_TYPE_ID",
	6: "CONVERSION_REPORTING_STATE",
}

var CallPlaceholderFieldEnum_CallPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":                0,
	"UNKNOWN":                    1,
	"PHONE_NUMBER":               2,
	"COUNTRY_CODE":               3,
	"TRACKED":                    4,
	"CONVERSION_TYPE_ID":         5,
	"CONVERSION_REPORTING_STATE": 6,
}

func (x CallPlaceholderFieldEnum_CallPlaceholderField) String() string {
	return proto.EnumName(CallPlaceholderFieldEnum_CallPlaceholderField_name, int32(x))
}

func (CallPlaceholderFieldEnum_CallPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_08ddc446f9f82f36, []int{0, 0}
}

// Values for Call placeholder fields.
type CallPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallPlaceholderFieldEnum) Reset()         { *m = CallPlaceholderFieldEnum{} }
func (m *CallPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*CallPlaceholderFieldEnum) ProtoMessage()    {}
func (*CallPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_08ddc446f9f82f36, []int{0}
}

func (m *CallPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *CallPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (m *CallPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallPlaceholderFieldEnum.Merge(m, src)
}
func (m *CallPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_CallPlaceholderFieldEnum.Size(m)
}
func (m *CallPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CallPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CallPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.CallPlaceholderFieldEnum_CallPlaceholderField", CallPlaceholderFieldEnum_CallPlaceholderField_name, CallPlaceholderFieldEnum_CallPlaceholderField_value)
	proto.RegisterType((*CallPlaceholderFieldEnum)(nil), "google.ads.googleads.v1.enums.CallPlaceholderFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/call_placeholder_field.proto", fileDescriptor_08ddc446f9f82f36)
}

var fileDescriptor_08ddc446f9f82f36 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x6a, 0xdb, 0x40,
	0x14, 0xc6, 0x2b, 0xb9, 0x75, 0x61, 0x5c, 0xe8, 0x20, 0x4a, 0x69, 0x4d, 0x5d, 0xb0, 0x0f, 0x30,
	0x42, 0x74, 0x37, 0x5d, 0xe9, 0xcf, 0xd8, 0x15, 0xa6, 0x23, 0x21, 0x4b, 0x0a, 0x0e, 0x02, 0x31,
	0xb1, 0x14, 0x45, 0x30, 0xd6, 0x08, 0x8f, 0xed, 0xdb, 0x64, 0x93, 0x65, 0x8e, 0x90, 0x23, 0xe4,
	0x28, 0x59, 0xe4, 0x0c, 0x41, 0x52, 0xec, 0x64, 0xe1, 0x64, 0x33, 0x7c, 0xbc, 0xef, 0xfd, 0x3e,
	0xe6, 0xbd, 0x07, 0x70, 0x21, 0x44, 0xc1, 0x73, 0x9d, 0x65, 0x52, 0xef, 0x64, 0xa3, 0xf6, 0x86,
	0x9e, 0x57, 0xbb, 0xb5, 0xd4, 0x57, 0x8c, 0xf3, 0xb4, 0xe6, 0x6c, 0x95, 0x5f, 0x09, 0x9e, 0xe5,
	0x9b, 0xf4, 0xb2, 0xcc, 0x79, 0x86, 0xea, 0x8d, 0xd8, 0x0a, 0x6d, 0xd4, 0x01, 0x88, 0x65, 0x12,
	0x1d, 0x59, 0xb4, 0x37, 0x50, 0xcb, 0x0e, 0x7f, 0x1d, 0xa2, 0xeb, 0x52, 0x67, 0x55, 0x25, 0xb6,
	0x6c, 0x5b, 0x8a, 0x4a, 0x76, 0xf0, 0xe4, 0x4e, 0x01, 0x3f, 0x6c, 0xc6, 0xb9, 0xff, 0x12, 0x3e,
	0x6d, 0xb2, 0x49, 0xb5, 0x5b, 0x4f, 0xae, 0x15, 0xf0, 0xed, 0x94, 0xa9, 0x7d, 0x05, 0x83, 0x88,
	0x2e, 0x7c, 0x62, 0xbb, 0x53, 0x97, 0x38, 0xf0, 0x83, 0x36, 0x00, 0x9f, 0x23, 0x3a, 0xa7, 0xde,
	0x19, 0x85, 0x8a, 0x06, 0xc1, 0x17, 0xff, 0x9f, 0x47, 0x49, 0x4a, 0xa3, 0xff, 0x16, 0x09, 0xa0,
	0xda, 0x54, 0x6c, 0x2f, 0xa2, 0x61, 0xb0, 0x4c, 0x6d, 0xcf, 0x21, 0xb0, 0xd7, 0x00, 0x61, 0x60,
	0xda, 0x73, 0xe2, 0xc0, 0x8f, 0xda, 0x77, 0xa0, 0xd9, 0x1e, 0x8d, 0x49, 0xb0, 0x70, 0x3d, 0x9a,
	0x86, 0x4b, 0x9f, 0xa4, 0xae, 0x03, 0x3f, 0x69, 0xbf, 0xc1, 0xf0, 0x55, 0x3d, 0x20, 0xbe, 0x17,
	0x84, 0x2e, 0x9d, 0xa5, 0x8b, 0xd0, 0x0c, 0x09, 0xec, 0x5b, 0x8f, 0x0a, 0x18, 0xaf, 0xc4, 0x1a,
	0xbd, 0xbb, 0x00, 0xeb, 0xe7, 0xa9, 0x11, 0xfc, 0x66, 0x7a, 0x5f, 0x39, 0xb7, 0x9e, 0xd9, 0x42,
	0x70, 0x56, 0x15, 0x48, 0x6c, 0x0a, 0xbd, 0xc8, 0xab, 0x76, 0x37, 0x87, 0x43, 0xd4, 0xa5, 0x7c,
	0xe3, 0x2e, 0x7f, 0xdb, 0xf7, 0x46, 0xed, 0xcd, 0x4c, 0xf3, 0x56, 0x1d, 0xcd, 0xba, 0x28, 0x33,
	0x93, 0xa8, 0x93, 0x8d, 0x8a, 0x0d, 0xd4, 0xec, 0x52, 0xde, 0x1f, 0xfc, 0xc4, 0xcc, 0x64, 0x72,
	0xf4, 0x93, 0xd8, 0x48, 0x5a, 0xff, 0x41, 0x1d, 0x77, 0x45, 0x8c, 0xcd, 0x4c, 0x62, 0x7c, 0xec,
	0xc0, 0x38, 0x36, 0x30, 0x6e, 0x7b, 0x2e, 0xfa, 0xed, 0xc7, 0xfe, 0x3c, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x44, 0xc8, 0x4a, 0x76, 0x2f, 0x02, 0x00, 0x00,
}
