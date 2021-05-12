// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/change_status_resource_type.proto

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

// Enum listing the resource types support by the ChangeStatus resource.
type ChangeStatusResourceTypeEnum_ChangeStatusResourceType int32

const (
	// No value has been specified.
	ChangeStatusResourceTypeEnum_UNSPECIFIED ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 0
	// Used for return value only. Represents an unclassified resource unknown
	// in this version.
	ChangeStatusResourceTypeEnum_UNKNOWN ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 1
	// An AdGroup resource change.
	ChangeStatusResourceTypeEnum_AD_GROUP ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 3
	// An AdGroupAd resource change.
	ChangeStatusResourceTypeEnum_AD_GROUP_AD ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 4
	// An AdGroupCriterion resource change.
	ChangeStatusResourceTypeEnum_AD_GROUP_CRITERION ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 5
	// A Campaign resource change.
	ChangeStatusResourceTypeEnum_CAMPAIGN ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 6
	// A CampaignCriterion resource change.
	ChangeStatusResourceTypeEnum_CAMPAIGN_CRITERION ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 7
	// A Feed resource change.
	ChangeStatusResourceTypeEnum_FEED ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 9
	// A FeedItem resource change.
	ChangeStatusResourceTypeEnum_FEED_ITEM ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 10
	// An AdGroupFeed resource change.
	ChangeStatusResourceTypeEnum_AD_GROUP_FEED ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 11
	// A CampaignFeed resource change.
	ChangeStatusResourceTypeEnum_CAMPAIGN_FEED ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 12
	// An AdGroupBidModifier resource change.
	ChangeStatusResourceTypeEnum_AD_GROUP_BID_MODIFIER ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 13
)

var ChangeStatusResourceTypeEnum_ChangeStatusResourceType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	3:  "AD_GROUP",
	4:  "AD_GROUP_AD",
	5:  "AD_GROUP_CRITERION",
	6:  "CAMPAIGN",
	7:  "CAMPAIGN_CRITERION",
	9:  "FEED",
	10: "FEED_ITEM",
	11: "AD_GROUP_FEED",
	12: "CAMPAIGN_FEED",
	13: "AD_GROUP_BID_MODIFIER",
}

var ChangeStatusResourceTypeEnum_ChangeStatusResourceType_value = map[string]int32{
	"UNSPECIFIED":           0,
	"UNKNOWN":               1,
	"AD_GROUP":              3,
	"AD_GROUP_AD":           4,
	"AD_GROUP_CRITERION":    5,
	"CAMPAIGN":              6,
	"CAMPAIGN_CRITERION":    7,
	"FEED":                  9,
	"FEED_ITEM":             10,
	"AD_GROUP_FEED":         11,
	"CAMPAIGN_FEED":         12,
	"AD_GROUP_BID_MODIFIER": 13,
}

func (x ChangeStatusResourceTypeEnum_ChangeStatusResourceType) String() string {
	return proto.EnumName(ChangeStatusResourceTypeEnum_ChangeStatusResourceType_name, int32(x))
}

func (ChangeStatusResourceTypeEnum_ChangeStatusResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_be6c457197dd2372, []int{0, 0}
}

// Container for enum describing supported resource types for the ChangeStatus
// resource.
type ChangeStatusResourceTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeStatusResourceTypeEnum) Reset()         { *m = ChangeStatusResourceTypeEnum{} }
func (m *ChangeStatusResourceTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ChangeStatusResourceTypeEnum) ProtoMessage()    {}
func (*ChangeStatusResourceTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_be6c457197dd2372, []int{0}
}

func (m *ChangeStatusResourceTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeStatusResourceTypeEnum.Unmarshal(m, b)
}
func (m *ChangeStatusResourceTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeStatusResourceTypeEnum.Marshal(b, m, deterministic)
}
func (m *ChangeStatusResourceTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeStatusResourceTypeEnum.Merge(m, src)
}
func (m *ChangeStatusResourceTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ChangeStatusResourceTypeEnum.Size(m)
}
func (m *ChangeStatusResourceTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeStatusResourceTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeStatusResourceTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType", ChangeStatusResourceTypeEnum_ChangeStatusResourceType_name, ChangeStatusResourceTypeEnum_ChangeStatusResourceType_value)
	proto.RegisterType((*ChangeStatusResourceTypeEnum)(nil), "google.ads.googleads.v2.enums.ChangeStatusResourceTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/change_status_resource_type.proto", fileDescriptor_be6c457197dd2372)
}

var fileDescriptor_be6c457197dd2372 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0x8a, 0xdb, 0x30,
	0x10, 0xc7, 0x6b, 0xef, 0x76, 0x3f, 0x94, 0x0d, 0x55, 0x05, 0x2d, 0x6d, 0xd9, 0x1c, 0x76, 0x1f,
	0x40, 0x06, 0xf7, 0xa6, 0x1e, 0x8a, 0xfc, 0xb1, 0xc6, 0x94, 0xd8, 0xc6, 0x9b, 0xa4, 0x50, 0x0c,
	0xc6, 0x8d, 0x85, 0x1b, 0x48, 0x24, 0x63, 0xd9, 0x81, 0xbc, 0x45, 0x9f, 0xa1, 0xc7, 0x3e, 0x4a,
	0x1f, 0xa5, 0x97, 0x42, 0x9f, 0xa0, 0x48, 0x8e, 0x4d, 0x2f, 0xd9, 0x8b, 0xf9, 0xcf, 0xcc, 0x6f,
	0xfe, 0x63, 0xcd, 0x80, 0x8f, 0x95, 0x10, 0xd5, 0x96, 0x59, 0x45, 0x29, 0xad, 0x5e, 0x2a, 0xb5,
	0xb7, 0x2d, 0xc6, 0xbb, 0x9d, 0xb4, 0xd6, 0xdf, 0x0a, 0x5e, 0xb1, 0x5c, 0xb6, 0x45, 0xdb, 0xc9,
	0xbc, 0x61, 0x52, 0x74, 0xcd, 0x9a, 0xe5, 0xed, 0xa1, 0x66, 0xb8, 0x6e, 0x44, 0x2b, 0xd0, 0xac,
	0xef, 0xc2, 0x45, 0x29, 0xf1, 0x68, 0x80, 0xf7, 0x36, 0xd6, 0x06, 0xef, 0x6e, 0x07, 0xff, 0x7a,
	0x63, 0x15, 0x9c, 0x8b, 0xb6, 0x68, 0x37, 0x82, 0xcb, 0xbe, 0xf9, 0xfe, 0xbb, 0x09, 0x6e, 0x5d,
	0x3d, 0xe2, 0x51, 0x4f, 0x48, 0x8f, 0x03, 0x16, 0x87, 0x9a, 0xf9, 0xbc, 0xdb, 0xdd, 0xff, 0x31,
	0xc0, 0x9b, 0x53, 0x00, 0x7a, 0x01, 0x26, 0xcb, 0xe8, 0x31, 0xf1, 0xdd, 0xf0, 0x21, 0xf4, 0x3d,
	0xf8, 0x0c, 0x4d, 0xc0, 0xe5, 0x32, 0xfa, 0x14, 0xc5, 0x9f, 0x23, 0x68, 0xa0, 0x1b, 0x70, 0x45,
	0xbd, 0x3c, 0x48, 0xe3, 0x65, 0x02, 0xcf, 0x14, 0x3b, 0x44, 0x39, 0xf5, 0xe0, 0x39, 0x7a, 0x0d,
	0xd0, 0x98, 0x70, 0xd3, 0x70, 0xe1, 0xa7, 0x61, 0x1c, 0xc1, 0xe7, 0xaa, 0xcd, 0xa5, 0xf3, 0x84,
	0x86, 0x41, 0x04, 0x2f, 0x14, 0x35, 0x44, 0xff, 0x51, 0x97, 0xe8, 0x0a, 0x9c, 0x3f, 0xf8, 0xbe,
	0x07, 0xaf, 0xd1, 0x14, 0x5c, 0x2b, 0x95, 0x87, 0x0b, 0x7f, 0x0e, 0x01, 0x7a, 0x09, 0xa6, 0xa3,
	0xad, 0x26, 0x26, 0x2a, 0x35, 0x7a, 0xe8, 0xd4, 0x0d, 0x7a, 0x0b, 0x5e, 0x8d, 0x94, 0x13, 0x7a,
	0xf9, 0x3c, 0xf6, 0xd4, 0x1b, 0x52, 0x38, 0x75, 0xfe, 0x1a, 0xe0, 0x6e, 0x2d, 0x76, 0xf8, 0xc9,
	0xb5, 0x3a, 0xb3, 0x53, 0x4b, 0x49, 0xd4, 0x5e, 0x13, 0xe3, 0x8b, 0x73, 0xec, 0xaf, 0xc4, 0xb6,
	0xe0, 0x15, 0x16, 0x4d, 0x65, 0x55, 0x8c, 0xeb, 0xad, 0x0f, 0x77, 0xae, 0x37, 0xf2, 0xc4, 0xd9,
	0x3f, 0xe8, 0xef, 0x0f, 0xf3, 0x2c, 0xa0, 0xf4, 0xa7, 0x39, 0x0b, 0x7a, 0x2b, 0x5a, 0x4a, 0xdc,
	0x4b, 0xa5, 0x56, 0x36, 0x56, 0x17, 0x92, 0xbf, 0x86, 0x7a, 0x46, 0x4b, 0x99, 0x8d, 0xf5, 0x6c,
	0x65, 0x67, 0xba, 0xfe, 0xdb, 0xbc, 0xeb, 0x93, 0x84, 0xd0, 0x52, 0x12, 0x32, 0x12, 0x84, 0xac,
	0x6c, 0x42, 0x34, 0xf3, 0xf5, 0x42, 0xff, 0xd8, 0xfb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb4,
	0xcb, 0xd2, 0x65, 0x8e, 0x02, 0x00, 0x00,
}
