// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/hotel_placeholder_field.proto

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

// Possible values for Hotel placeholder fields.
type HotelPlaceholderFieldEnum_HotelPlaceholderField int32

const (
	// Not specified.
	HotelPlaceholderFieldEnum_UNSPECIFIED HotelPlaceholderFieldEnum_HotelPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	HotelPlaceholderFieldEnum_UNKNOWN HotelPlaceholderFieldEnum_HotelPlaceholderField = 1
	// Data Type: STRING. Required. Unique ID.
	HotelPlaceholderFieldEnum_PROPERTY_ID HotelPlaceholderFieldEnum_HotelPlaceholderField = 2
	// Data Type: STRING. Required. Main headline with property name to be shown
	// in dynamic ad.
	HotelPlaceholderFieldEnum_PROPERTY_NAME HotelPlaceholderFieldEnum_HotelPlaceholderField = 3
	// Data Type: STRING. Name of destination to be shown in dynamic ad.
	HotelPlaceholderFieldEnum_DESTINATION_NAME HotelPlaceholderFieldEnum_HotelPlaceholderField = 4
	// Data Type: STRING. Description of destination to be shown in dynamic ad.
	HotelPlaceholderFieldEnum_DESCRIPTION HotelPlaceholderFieldEnum_HotelPlaceholderField = 5
	// Data Type: STRING. Complete property address, including postal code.
	HotelPlaceholderFieldEnum_ADDRESS HotelPlaceholderFieldEnum_HotelPlaceholderField = 6
	// Data Type: STRING. Price to be shown in the ad.
	// Example: "100.00 USD"
	HotelPlaceholderFieldEnum_PRICE HotelPlaceholderFieldEnum_HotelPlaceholderField = 7
	// Data Type: STRING. Formatted price to be shown in the ad.
	// Example: "Starting at $100.00 USD", "$80 - $100"
	HotelPlaceholderFieldEnum_FORMATTED_PRICE HotelPlaceholderFieldEnum_HotelPlaceholderField = 8
	// Data Type: STRING. Sale price to be shown in the ad.
	// Example: "80.00 USD"
	HotelPlaceholderFieldEnum_SALE_PRICE HotelPlaceholderFieldEnum_HotelPlaceholderField = 9
	// Data Type: STRING. Formatted sale price to be shown in the ad.
	// Example: "On sale for $80.00", "$60 - $80"
	HotelPlaceholderFieldEnum_FORMATTED_SALE_PRICE HotelPlaceholderFieldEnum_HotelPlaceholderField = 10
	// Data Type: URL. Image to be displayed in the ad.
	HotelPlaceholderFieldEnum_IMAGE_URL HotelPlaceholderFieldEnum_HotelPlaceholderField = 11
	// Data Type: STRING. Category of property used to group like items together
	// for recommendation engine.
	HotelPlaceholderFieldEnum_CATEGORY HotelPlaceholderFieldEnum_HotelPlaceholderField = 12
	// Data Type: INT64. Star rating (1 to 5) used to group like items
	// together for recommendation engine.
	HotelPlaceholderFieldEnum_STAR_RATING HotelPlaceholderFieldEnum_HotelPlaceholderField = 13
	// Data Type: STRING_LIST. Keywords used for product retrieval.
	HotelPlaceholderFieldEnum_CONTEXTUAL_KEYWORDS HotelPlaceholderFieldEnum_HotelPlaceholderField = 14
	// Data Type: URL_LIST. Required. Final URLs for the ad when using Upgraded
	// URLs. User will be redirected to these URLs when they click on an ad, or
	// when they click on a specific flight for ads that show multiple
	// flights.
	HotelPlaceholderFieldEnum_FINAL_URLS HotelPlaceholderFieldEnum_HotelPlaceholderField = 15
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	HotelPlaceholderFieldEnum_FINAL_MOBILE_URLS HotelPlaceholderFieldEnum_HotelPlaceholderField = 16
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	HotelPlaceholderFieldEnum_TRACKING_URL HotelPlaceholderFieldEnum_HotelPlaceholderField = 17
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//   scheme.
	// host_path: identifies the specific content within your application.
	HotelPlaceholderFieldEnum_ANDROID_APP_LINK HotelPlaceholderFieldEnum_HotelPlaceholderField = 18
	// Data Type: STRING_LIST. List of recommended property IDs to show together
	// with this item.
	HotelPlaceholderFieldEnum_SIMILAR_PROPERTY_IDS HotelPlaceholderFieldEnum_HotelPlaceholderField = 19
	// Data Type: STRING. iOS app link.
	HotelPlaceholderFieldEnum_IOS_APP_LINK HotelPlaceholderFieldEnum_HotelPlaceholderField = 20
	// Data Type: INT64. iOS app store ID.
	HotelPlaceholderFieldEnum_IOS_APP_STORE_ID HotelPlaceholderFieldEnum_HotelPlaceholderField = 21
)

var HotelPlaceholderFieldEnum_HotelPlaceholderField_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "PROPERTY_ID",
	3:  "PROPERTY_NAME",
	4:  "DESTINATION_NAME",
	5:  "DESCRIPTION",
	6:  "ADDRESS",
	7:  "PRICE",
	8:  "FORMATTED_PRICE",
	9:  "SALE_PRICE",
	10: "FORMATTED_SALE_PRICE",
	11: "IMAGE_URL",
	12: "CATEGORY",
	13: "STAR_RATING",
	14: "CONTEXTUAL_KEYWORDS",
	15: "FINAL_URLS",
	16: "FINAL_MOBILE_URLS",
	17: "TRACKING_URL",
	18: "ANDROID_APP_LINK",
	19: "SIMILAR_PROPERTY_IDS",
	20: "IOS_APP_LINK",
	21: "IOS_APP_STORE_ID",
}

var HotelPlaceholderFieldEnum_HotelPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"PROPERTY_ID":          2,
	"PROPERTY_NAME":        3,
	"DESTINATION_NAME":     4,
	"DESCRIPTION":          5,
	"ADDRESS":              6,
	"PRICE":                7,
	"FORMATTED_PRICE":      8,
	"SALE_PRICE":           9,
	"FORMATTED_SALE_PRICE": 10,
	"IMAGE_URL":            11,
	"CATEGORY":             12,
	"STAR_RATING":          13,
	"CONTEXTUAL_KEYWORDS":  14,
	"FINAL_URLS":           15,
	"FINAL_MOBILE_URLS":    16,
	"TRACKING_URL":         17,
	"ANDROID_APP_LINK":     18,
	"SIMILAR_PROPERTY_IDS": 19,
	"IOS_APP_LINK":         20,
	"IOS_APP_STORE_ID":     21,
}

func (x HotelPlaceholderFieldEnum_HotelPlaceholderField) String() string {
	return proto.EnumName(HotelPlaceholderFieldEnum_HotelPlaceholderField_name, int32(x))
}

func (HotelPlaceholderFieldEnum_HotelPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bf09812e91f0c054, []int{0, 0}
}

// Values for Hotel placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type HotelPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotelPlaceholderFieldEnum) Reset()         { *m = HotelPlaceholderFieldEnum{} }
func (m *HotelPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*HotelPlaceholderFieldEnum) ProtoMessage()    {}
func (*HotelPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf09812e91f0c054, []int{0}
}

func (m *HotelPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HotelPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *HotelPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HotelPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (m *HotelPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotelPlaceholderFieldEnum.Merge(m, src)
}
func (m *HotelPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_HotelPlaceholderFieldEnum.Size(m)
}
func (m *HotelPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_HotelPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_HotelPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.HotelPlaceholderFieldEnum_HotelPlaceholderField", HotelPlaceholderFieldEnum_HotelPlaceholderField_name, HotelPlaceholderFieldEnum_HotelPlaceholderField_value)
	proto.RegisterType((*HotelPlaceholderFieldEnum)(nil), "google.ads.googleads.v2.enums.HotelPlaceholderFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/hotel_placeholder_field.proto", fileDescriptor_bf09812e91f0c054)
}

var fileDescriptor_bf09812e91f0c054 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xdd, 0x6e, 0xda, 0x30,
	0x18, 0x5d, 0x61, 0xfd, 0x73, 0xa1, 0x18, 0x03, 0xda, 0x6f, 0x2f, 0xda, 0x07, 0x48, 0x24, 0x76,
	0x97, 0x5e, 0x99, 0xc4, 0x64, 0x16, 0xc1, 0x89, 0x6c, 0x43, 0xc7, 0x84, 0x14, 0x65, 0x4d, 0x96,
	0x22, 0x85, 0x18, 0x11, 0xda, 0x07, 0xda, 0xe5, 0x6e, 0xf6, 0x1e, 0x93, 0xf6, 0x22, 0xbb, 0xda,
	0x23, 0x4c, 0x4e, 0x0a, 0x54, 0xd3, 0xb6, 0x9b, 0xe8, 0xf3, 0x39, 0xdf, 0x39, 0xfe, 0xe2, 0xef,
	0x80, 0xeb, 0x54, 0xa9, 0x34, 0x4b, 0xcc, 0x28, 0x2e, 0xcc, 0xaa, 0xd4, 0xd5, 0x43, 0xdf, 0x4c,
	0xf2, 0xfb, 0x65, 0x61, 0xde, 0xa9, 0x4d, 0x92, 0x85, 0xab, 0x2c, 0xba, 0x4d, 0xee, 0x54, 0x16,
	0x27, 0xeb, 0xf0, 0xf3, 0x22, 0xc9, 0x62, 0x63, 0xb5, 0x56, 0x1b, 0x85, 0x2e, 0x2a, 0x85, 0x11,
	0xc5, 0x85, 0xb1, 0x13, 0x1b, 0x0f, 0x7d, 0xa3, 0x14, 0xbf, 0x7e, 0xbb, 0xf5, 0x5e, 0x2d, 0xcc,
	0x28, 0xcf, 0xd5, 0x26, 0xda, 0x2c, 0x54, 0x5e, 0x54, 0xe2, 0xab, 0x1f, 0x75, 0xf0, 0xea, 0xbd,
	0xb6, 0x0f, 0xf6, 0xee, 0x43, 0x6d, 0x4e, 0xf2, 0xfb, 0xe5, 0xd5, 0xb7, 0x3a, 0xe8, 0xfd, 0x95,
	0x45, 0x2d, 0x70, 0x36, 0x61, 0x22, 0x20, 0x36, 0x1d, 0x52, 0xe2, 0xc0, 0x67, 0xe8, 0x0c, 0x1c,
	0x4f, 0xd8, 0x88, 0xf9, 0x37, 0x0c, 0x1e, 0x68, 0x36, 0xe0, 0x7e, 0x40, 0xb8, 0x9c, 0x85, 0xd4,
	0x81, 0x35, 0xd4, 0x06, 0xcd, 0x1d, 0xc0, 0xf0, 0x98, 0xc0, 0x3a, 0xea, 0x02, 0xe8, 0x10, 0x21,
	0x29, 0xc3, 0x92, 0xfa, 0xac, 0x42, 0x9f, 0x6b, 0xa5, 0x43, 0x84, 0xcd, 0x69, 0xa0, 0x51, 0x78,
	0xa8, 0x7d, 0xb1, 0xe3, 0x70, 0x22, 0x04, 0x3c, 0x42, 0xa7, 0xe0, 0x30, 0xe0, 0xd4, 0x26, 0xf0,
	0x18, 0x75, 0x40, 0x6b, 0xe8, 0xf3, 0x31, 0x96, 0x92, 0x38, 0x61, 0x05, 0x9e, 0xa0, 0x73, 0x00,
	0x04, 0xf6, 0xc8, 0xe3, 0xf9, 0x14, 0xbd, 0x04, 0xdd, 0x7d, 0xd3, 0x13, 0x06, 0xa0, 0x26, 0x38,
	0xa5, 0x63, 0xec, 0x92, 0x70, 0xc2, 0x3d, 0x78, 0x86, 0x1a, 0xe0, 0xc4, 0xc6, 0x92, 0xb8, 0x3e,
	0x9f, 0xc1, 0x86, 0x1e, 0x42, 0x48, 0xcc, 0x43, 0x8e, 0x25, 0x65, 0x2e, 0x6c, 0xa2, 0x17, 0xa0,
	0x63, 0xfb, 0x4c, 0x92, 0x0f, 0x72, 0x82, 0xbd, 0x70, 0x44, 0x66, 0x37, 0x3e, 0x77, 0x04, 0x3c,
	0xd7, 0x17, 0x0e, 0x29, 0xc3, 0x9e, 0xb6, 0x11, 0xb0, 0x85, 0x7a, 0xa0, 0x5d, 0x9d, 0xc7, 0xfe,
	0x80, 0x7a, 0xa4, 0x82, 0x21, 0x82, 0xa0, 0x21, 0x39, 0xb6, 0x47, 0x94, 0xb9, 0xe5, 0x85, 0x6d,
	0xfd, 0xf7, 0x98, 0x39, 0xdc, 0xa7, 0x4e, 0x88, 0x83, 0x20, 0xf4, 0x28, 0x1b, 0x41, 0xa4, 0xe7,
	0x15, 0x74, 0x4c, 0x3d, 0xcc, 0xc3, 0x27, 0xef, 0x27, 0x60, 0x47, 0x3b, 0x50, 0x5f, 0xec, 0x7b,
	0xbb, 0xda, 0x61, 0x8b, 0x08, 0xe9, 0x73, 0xa2, 0x1f, 0xba, 0x37, 0xf8, 0x75, 0x00, 0x2e, 0x6f,
	0xd5, 0xd2, 0xf8, 0x6f, 0x26, 0x06, 0x6f, 0xca, 0xa5, 0x16, 0x7f, 0x6e, 0x35, 0xd0, 0x91, 0x08,
	0x0e, 0x3e, 0x0e, 0x1e, 0xd5, 0xa9, 0xca, 0xa2, 0x3c, 0x35, 0xd4, 0x3a, 0x35, 0xd3, 0x24, 0x2f,
	0x03, 0xb3, 0x8d, 0xe7, 0x6a, 0x51, 0xfc, 0x23, 0xad, 0xd7, 0xe5, 0xf7, 0x4b, 0xad, 0xee, 0x62,
	0xfc, 0xb5, 0x76, 0xe1, 0x56, 0x56, 0x38, 0x2e, 0x8c, 0xaa, 0xd4, 0xd5, 0xb4, 0x6f, 0xe8, 0x7c,
	0x15, 0xdf, 0xb7, 0xfc, 0x1c, 0xc7, 0xc5, 0x7c, 0xc7, 0xcf, 0xa7, 0xfd, 0x79, 0xc9, 0xff, 0xac,
	0x5d, 0x56, 0xa0, 0x65, 0xe1, 0xb8, 0xb0, 0xac, 0x5d, 0x87, 0x65, 0x4d, 0xfb, 0x96, 0x55, 0xf6,
	0x7c, 0x3a, 0x2a, 0x07, 0x7b, 0xf7, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x02, 0xac, 0x61, 0x97, 0x45,
	0x03, 0x00, 0x00,
}
