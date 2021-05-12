// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/education_placeholder_field.proto

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

// Possible values for Education placeholder fields.
type EducationPlaceholderFieldEnum_EducationPlaceholderField int32

const (
	// Not specified.
	EducationPlaceholderFieldEnum_UNSPECIFIED EducationPlaceholderFieldEnum_EducationPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	EducationPlaceholderFieldEnum_UNKNOWN EducationPlaceholderFieldEnum_EducationPlaceholderField = 1
	// Data Type: STRING. Required. Combination of PROGRAM ID and LOCATION ID
	// must be unique per offer.
	EducationPlaceholderFieldEnum_PROGRAM_ID EducationPlaceholderFieldEnum_EducationPlaceholderField = 2
	// Data Type: STRING. Combination of PROGRAM ID and LOCATION ID must be
	// unique per offer.
	EducationPlaceholderFieldEnum_LOCATION_ID EducationPlaceholderFieldEnum_EducationPlaceholderField = 3
	// Data Type: STRING. Required. Main headline with program name to be shown
	// in dynamic ad.
	EducationPlaceholderFieldEnum_PROGRAM_NAME EducationPlaceholderFieldEnum_EducationPlaceholderField = 4
	// Data Type: STRING. Area of study that can be shown in dynamic ad.
	EducationPlaceholderFieldEnum_AREA_OF_STUDY EducationPlaceholderFieldEnum_EducationPlaceholderField = 5
	// Data Type: STRING. Description of program that can be shown in dynamic
	// ad.
	EducationPlaceholderFieldEnum_PROGRAM_DESCRIPTION EducationPlaceholderFieldEnum_EducationPlaceholderField = 6
	// Data Type: STRING. Name of school that can be shown in dynamic ad.
	EducationPlaceholderFieldEnum_SCHOOL_NAME EducationPlaceholderFieldEnum_EducationPlaceholderField = 7
	// Data Type: STRING. Complete school address, including postal code.
	EducationPlaceholderFieldEnum_ADDRESS EducationPlaceholderFieldEnum_EducationPlaceholderField = 8
	// Data Type: URL. Image to be displayed in ads.
	EducationPlaceholderFieldEnum_THUMBNAIL_IMAGE_URL EducationPlaceholderFieldEnum_EducationPlaceholderField = 9
	// Data Type: URL. Alternative hosted file of image to be used in the ad.
	EducationPlaceholderFieldEnum_ALTERNATIVE_THUMBNAIL_IMAGE_URL EducationPlaceholderFieldEnum_EducationPlaceholderField = 10
	// Data Type: URL_LIST. Required. Final URLs to be used in ad when using
	// Upgraded URLs; the more specific the better (e.g. the individual URL of a
	// specific program and its location).
	EducationPlaceholderFieldEnum_FINAL_URLS EducationPlaceholderFieldEnum_EducationPlaceholderField = 11
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	EducationPlaceholderFieldEnum_FINAL_MOBILE_URLS EducationPlaceholderFieldEnum_EducationPlaceholderField = 12
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	EducationPlaceholderFieldEnum_TRACKING_URL EducationPlaceholderFieldEnum_EducationPlaceholderField = 13
	// Data Type: STRING_LIST. Keywords used for product retrieval.
	EducationPlaceholderFieldEnum_CONTEXTUAL_KEYWORDS EducationPlaceholderFieldEnum_EducationPlaceholderField = 14
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//   scheme.
	// host_path: identifies the specific content within your application.
	EducationPlaceholderFieldEnum_ANDROID_APP_LINK EducationPlaceholderFieldEnum_EducationPlaceholderField = 15
	// Data Type: STRING_LIST. List of recommended program IDs to show together
	// with this item.
	EducationPlaceholderFieldEnum_SIMILAR_PROGRAM_IDS EducationPlaceholderFieldEnum_EducationPlaceholderField = 16
	// Data Type: STRING. iOS app link.
	EducationPlaceholderFieldEnum_IOS_APP_LINK EducationPlaceholderFieldEnum_EducationPlaceholderField = 17
	// Data Type: INT64. iOS app store ID.
	EducationPlaceholderFieldEnum_IOS_APP_STORE_ID EducationPlaceholderFieldEnum_EducationPlaceholderField = 18
)

var EducationPlaceholderFieldEnum_EducationPlaceholderField_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "PROGRAM_ID",
	3:  "LOCATION_ID",
	4:  "PROGRAM_NAME",
	5:  "AREA_OF_STUDY",
	6:  "PROGRAM_DESCRIPTION",
	7:  "SCHOOL_NAME",
	8:  "ADDRESS",
	9:  "THUMBNAIL_IMAGE_URL",
	10: "ALTERNATIVE_THUMBNAIL_IMAGE_URL",
	11: "FINAL_URLS",
	12: "FINAL_MOBILE_URLS",
	13: "TRACKING_URL",
	14: "CONTEXTUAL_KEYWORDS",
	15: "ANDROID_APP_LINK",
	16: "SIMILAR_PROGRAM_IDS",
	17: "IOS_APP_LINK",
	18: "IOS_APP_STORE_ID",
}

var EducationPlaceholderFieldEnum_EducationPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":                     0,
	"UNKNOWN":                         1,
	"PROGRAM_ID":                      2,
	"LOCATION_ID":                     3,
	"PROGRAM_NAME":                    4,
	"AREA_OF_STUDY":                   5,
	"PROGRAM_DESCRIPTION":             6,
	"SCHOOL_NAME":                     7,
	"ADDRESS":                         8,
	"THUMBNAIL_IMAGE_URL":             9,
	"ALTERNATIVE_THUMBNAIL_IMAGE_URL": 10,
	"FINAL_URLS":                      11,
	"FINAL_MOBILE_URLS":               12,
	"TRACKING_URL":                    13,
	"CONTEXTUAL_KEYWORDS":             14,
	"ANDROID_APP_LINK":                15,
	"SIMILAR_PROGRAM_IDS":             16,
	"IOS_APP_LINK":                    17,
	"IOS_APP_STORE_ID":                18,
}

func (x EducationPlaceholderFieldEnum_EducationPlaceholderField) String() string {
	return proto.EnumName(EducationPlaceholderFieldEnum_EducationPlaceholderField_name, int32(x))
}

func (EducationPlaceholderFieldEnum_EducationPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8fd4c47241194ad2, []int{0, 0}
}

// Values for Education placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type EducationPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EducationPlaceholderFieldEnum) Reset()         { *m = EducationPlaceholderFieldEnum{} }
func (m *EducationPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*EducationPlaceholderFieldEnum) ProtoMessage()    {}
func (*EducationPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fd4c47241194ad2, []int{0}
}

func (m *EducationPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EducationPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *EducationPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EducationPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (m *EducationPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EducationPlaceholderFieldEnum.Merge(m, src)
}
func (m *EducationPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_EducationPlaceholderFieldEnum.Size(m)
}
func (m *EducationPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_EducationPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_EducationPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.EducationPlaceholderFieldEnum_EducationPlaceholderField", EducationPlaceholderFieldEnum_EducationPlaceholderField_name, EducationPlaceholderFieldEnum_EducationPlaceholderField_value)
	proto.RegisterType((*EducationPlaceholderFieldEnum)(nil), "google.ads.googleads.v1.enums.EducationPlaceholderFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/education_placeholder_field.proto", fileDescriptor_8fd4c47241194ad2)
}

var fileDescriptor_8fd4c47241194ad2 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0xd9, 0x16, 0x76, 0xc1, 0xdd, 0x3f, 0x6e, 0x00, 0x21, 0x10, 0x05, 0xed, 0x72, 0x4f,
	0x54, 0x71, 0x0b, 0x07, 0x34, 0x4d, 0xdc, 0xae, 0xd5, 0xc4, 0x8e, 0xec, 0xa4, 0xcb, 0xa2, 0x4a,
	0x56, 0x68, 0x42, 0xa8, 0xd4, 0x26, 0x55, 0xd3, 0xee, 0xdb, 0x70, 0xe1, 0xc8, 0x4b, 0x70, 0xe7,
	0x51, 0xb8, 0xf2, 0x02, 0xc8, 0x09, 0x69, 0x2f, 0x94, 0x4b, 0x35, 0x9e, 0xf9, 0xe6, 0xf7, 0xb9,
	0xce, 0x87, 0xde, 0x67, 0x45, 0x91, 0x2d, 0x52, 0x2b, 0x4e, 0x4a, 0xab, 0x2e, 0x75, 0x75, 0xd7,
	0xb7, 0xd2, 0x7c, 0xbb, 0x2c, 0xad, 0x34, 0xd9, 0xce, 0xe2, 0xcd, 0xbc, 0xc8, 0xd5, 0x6a, 0x11,
	0xcf, 0xd2, 0x2f, 0xc5, 0x22, 0x49, 0xd7, 0xea, 0xf3, 0x3c, 0x5d, 0x24, 0xe6, 0x6a, 0x5d, 0x6c,
	0x0a, 0xa3, 0x57, 0x6f, 0x99, 0x71, 0x52, 0x9a, 0x3b, 0x80, 0x79, 0xd7, 0x37, 0x2b, 0xc0, 0x8b,
	0x97, 0x0d, 0x7f, 0x35, 0xb7, 0xe2, 0x3c, 0x2f, 0x36, 0x15, 0xad, 0xac, 0x97, 0xaf, 0x7e, 0xb4,
	0x51, 0x8f, 0x34, 0x16, 0xc1, 0xde, 0x61, 0xa8, 0x0d, 0x48, 0xbe, 0x5d, 0x5e, 0x7d, 0x6d, 0xa3,
	0xe7, 0x07, 0x15, 0xc6, 0x05, 0xea, 0x44, 0x4c, 0x06, 0xc4, 0xa1, 0x43, 0x4a, 0x5c, 0x7c, 0xcf,
	0xe8, 0xa0, 0x93, 0x88, 0x8d, 0x19, 0xbf, 0x61, 0xf8, 0xc8, 0x38, 0x47, 0x28, 0x10, 0x7c, 0x24,
	0xc0, 0x57, 0xd4, 0xc5, 0x2d, 0xad, 0xf6, 0xb8, 0x03, 0x21, 0xe5, 0x4c, 0x37, 0xda, 0x06, 0x46,
	0xa7, 0x8d, 0x80, 0x81, 0x4f, 0xf0, 0x7d, 0xa3, 0x8b, 0xce, 0x40, 0x10, 0x50, 0x7c, 0xa8, 0x64,
	0x18, 0xb9, 0xb7, 0xf8, 0x81, 0xf1, 0x0c, 0x3d, 0x6e, 0x44, 0x2e, 0x91, 0x8e, 0xa0, 0x81, 0x06,
	0xe0, 0x63, 0x8d, 0x93, 0xce, 0x35, 0xe7, 0x5e, 0xbd, 0x7c, 0xa2, 0xcd, 0xc1, 0x75, 0x05, 0x91,
	0x12, 0x3f, 0xd4, 0x6b, 0xe1, 0x75, 0xe4, 0x0f, 0x18, 0x50, 0x4f, 0x51, 0x1f, 0x46, 0x44, 0x45,
	0xc2, 0xc3, 0x8f, 0x8c, 0x37, 0xe8, 0x35, 0x78, 0x21, 0x11, 0x0c, 0x42, 0x3a, 0x21, 0xea, 0x5f,
	0x22, 0xa4, 0xaf, 0x3e, 0xa4, 0x0c, 0x3c, 0x7d, 0x94, 0xb8, 0x63, 0x3c, 0x45, 0xdd, 0xfa, 0xec,
	0xf3, 0x01, 0xf5, 0x48, 0xdd, 0x3e, 0xd5, 0x7f, 0x20, 0x14, 0xe0, 0x8c, 0x29, 0x1b, 0x55, 0x8b,
	0x67, 0xda, 0xd6, 0xe1, 0x2c, 0x24, 0x1f, 0xc2, 0x08, 0x3c, 0x35, 0x26, 0xb7, 0x37, 0x5c, 0xb8,
	0x12, 0x9f, 0x1b, 0x4f, 0x10, 0x06, 0xe6, 0x0a, 0x4e, 0x5d, 0x05, 0x41, 0xa0, 0x3c, 0xca, 0xc6,
	0xf8, 0x42, 0xcb, 0x25, 0xf5, 0xa9, 0x07, 0x42, 0xed, 0x9f, 0x4a, 0x62, 0xac, 0xc9, 0x94, 0xcb,
	0xbd, 0xb4, 0xab, 0x01, 0x4d, 0x47, 0x86, 0x5c, 0x10, 0xfd, 0x84, 0xc6, 0xe0, 0xf7, 0x11, 0xba,
	0x9c, 0x15, 0x4b, 0xf3, 0xbf, 0x29, 0x18, 0xbc, 0x3a, 0xf8, 0x09, 0x03, 0x9d, 0x83, 0xe0, 0xe8,
	0xe3, 0xe0, 0x2f, 0x20, 0x2b, 0x16, 0x71, 0x9e, 0x99, 0xc5, 0x3a, 0xb3, 0xb2, 0x34, 0xaf, 0x52,
	0xd2, 0xe4, 0x72, 0x35, 0x2f, 0x0f, 0xc4, 0xf4, 0x5d, 0xf5, 0xfb, 0xad, 0xd5, 0x1e, 0x01, 0x7c,
	0x6f, 0xf5, 0x46, 0x35, 0x0a, 0x92, 0xd2, 0xac, 0x4b, 0x5d, 0x4d, 0xfa, 0xa6, 0x0e, 0x54, 0xf9,
	0xb3, 0x99, 0x4f, 0x21, 0x29, 0xa7, 0xbb, 0xf9, 0x74, 0xd2, 0x9f, 0x56, 0xf3, 0x5f, 0xad, 0xcb,
	0xba, 0x69, 0xdb, 0x90, 0x94, 0xb6, 0xbd, 0x53, 0xd8, 0xf6, 0xa4, 0x6f, 0xdb, 0x95, 0xe6, 0xd3,
	0x71, 0x75, 0xb1, 0xb7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x24, 0xab, 0x32, 0x3e, 0x03,
	0x00, 0x00,
}
