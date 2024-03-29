// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/campaign_feed_error.proto

package errors

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

// Enum describing possible campaign feed errors.
type CampaignFeedErrorEnum_CampaignFeedError int32

const (
	// Enum unspecified.
	CampaignFeedErrorEnum_UNSPECIFIED CampaignFeedErrorEnum_CampaignFeedError = 0
	// The received error code is not known in this version.
	CampaignFeedErrorEnum_UNKNOWN CampaignFeedErrorEnum_CampaignFeedError = 1
	// An active feed already exists for this campaign and placeholder type.
	CampaignFeedErrorEnum_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE CampaignFeedErrorEnum_CampaignFeedError = 2
	// The specified feed is removed.
	CampaignFeedErrorEnum_CANNOT_CREATE_FOR_REMOVED_FEED CampaignFeedErrorEnum_CampaignFeedError = 4
	// The CampaignFeed already exists. UPDATE should be used to modify the
	// existing CampaignFeed.
	CampaignFeedErrorEnum_CANNOT_CREATE_ALREADY_EXISTING_CAMPAIGN_FEED CampaignFeedErrorEnum_CampaignFeedError = 5
	// Cannot update removed campaign feed.
	CampaignFeedErrorEnum_CANNOT_MODIFY_REMOVED_CAMPAIGN_FEED CampaignFeedErrorEnum_CampaignFeedError = 6
	// Invalid placeholder type.
	CampaignFeedErrorEnum_INVALID_PLACEHOLDER_TYPE CampaignFeedErrorEnum_CampaignFeedError = 7
	// Feed mapping for this placeholder type does not exist.
	CampaignFeedErrorEnum_MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE CampaignFeedErrorEnum_CampaignFeedError = 8
)

var CampaignFeedErrorEnum_CampaignFeedError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE",
	4: "CANNOT_CREATE_FOR_REMOVED_FEED",
	5: "CANNOT_CREATE_ALREADY_EXISTING_CAMPAIGN_FEED",
	6: "CANNOT_MODIFY_REMOVED_CAMPAIGN_FEED",
	7: "INVALID_PLACEHOLDER_TYPE",
	8: "MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE",
}

var CampaignFeedErrorEnum_CampaignFeedError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE":     2,
	"CANNOT_CREATE_FOR_REMOVED_FEED":               4,
	"CANNOT_CREATE_ALREADY_EXISTING_CAMPAIGN_FEED": 5,
	"CANNOT_MODIFY_REMOVED_CAMPAIGN_FEED":          6,
	"INVALID_PLACEHOLDER_TYPE":                     7,
	"MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE":     8,
}

func (x CampaignFeedErrorEnum_CampaignFeedError) String() string {
	return proto.EnumName(CampaignFeedErrorEnum_CampaignFeedError_name, int32(x))
}

func (CampaignFeedErrorEnum_CampaignFeedError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_75d91332b6012d3f, []int{0, 0}
}

// Container for enum describing possible campaign feed errors.
type CampaignFeedErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignFeedErrorEnum) Reset()         { *m = CampaignFeedErrorEnum{} }
func (m *CampaignFeedErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignFeedErrorEnum) ProtoMessage()    {}
func (*CampaignFeedErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_75d91332b6012d3f, []int{0}
}

func (m *CampaignFeedErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignFeedErrorEnum.Unmarshal(m, b)
}
func (m *CampaignFeedErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignFeedErrorEnum.Marshal(b, m, deterministic)
}
func (m *CampaignFeedErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignFeedErrorEnum.Merge(m, src)
}
func (m *CampaignFeedErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignFeedErrorEnum.Size(m)
}
func (m *CampaignFeedErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignFeedErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignFeedErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.CampaignFeedErrorEnum_CampaignFeedError", CampaignFeedErrorEnum_CampaignFeedError_name, CampaignFeedErrorEnum_CampaignFeedError_value)
	proto.RegisterType((*CampaignFeedErrorEnum)(nil), "google.ads.googleads.v1.errors.CampaignFeedErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/campaign_feed_error.proto", fileDescriptor_75d91332b6012d3f)
}

var fileDescriptor_75d91332b6012d3f = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x8e, 0xd3, 0x30,
	0x14, 0x86, 0x69, 0x80, 0x19, 0xe4, 0x59, 0x50, 0x2c, 0x81, 0x10, 0x1a, 0x75, 0x11, 0x16, 0xb0,
	0x18, 0x39, 0x44, 0x6c, 0x90, 0x59, 0x79, 0x62, 0xa7, 0x58, 0x34, 0x8e, 0x95, 0x74, 0x02, 0x45,
	0x95, 0xac, 0x30, 0x09, 0x51, 0xa5, 0x69, 0x5c, 0xc5, 0x65, 0x0e, 0xc4, 0xb2, 0xe7, 0x60, 0xc5,
	0x51, 0xd8, 0x71, 0x03, 0xe4, 0xb8, 0xad, 0x54, 0x15, 0x66, 0x95, 0x5f, 0xce, 0xf7, 0xff, 0x7e,
	0xcf, 0xef, 0x81, 0x77, 0x8d, 0xd6, 0xcd, 0x4d, 0x1d, 0x94, 0x95, 0x09, 0x9c, 0xb4, 0xea, 0x36,
	0x0c, 0xea, 0xae, 0xd3, 0x9d, 0x09, 0xae, 0xcb, 0xe5, 0xaa, 0x5c, 0x34, 0xad, 0xfa, 0x56, 0xd7,
	0x95, 0xea, 0x0f, 0xd1, 0xaa, 0xd3, 0x6b, 0x0d, 0x47, 0x0e, 0x47, 0x65, 0x65, 0xd0, 0xde, 0x89,
	0x6e, 0x43, 0xe4, 0x9c, 0x2f, 0xce, 0x77, 0xc9, 0xab, 0x45, 0x50, 0xb6, 0xad, 0x5e, 0x97, 0xeb,
	0x85, 0x6e, 0x8d, 0x73, 0xfb, 0x3f, 0x3d, 0xf0, 0x34, 0xda, 0x66, 0xc7, 0x75, 0x5d, 0x31, 0x6b,
	0x62, 0xed, 0xf7, 0xa5, 0xbf, 0xf1, 0xc0, 0x93, 0xa3, 0x3f, 0xf0, 0x31, 0x38, 0xbb, 0x12, 0xb9,
	0x64, 0x11, 0x8f, 0x39, 0xa3, 0xc3, 0x7b, 0xf0, 0x0c, 0x9c, 0x5e, 0x89, 0x8f, 0x22, 0xfd, 0x24,
	0x86, 0x03, 0x78, 0x01, 0x5e, 0xc7, 0x8c, 0x51, 0x45, 0x26, 0x19, 0x23, 0x74, 0xa6, 0xd8, 0x67,
	0x9e, 0x4f, 0x73, 0x15, 0xa7, 0x99, 0x92, 0x13, 0x12, 0xb1, 0x0f, 0xe9, 0x84, 0xb2, 0x4c, 0x4d,
	0x67, 0x92, 0x0d, 0x3d, 0xe8, 0x83, 0x51, 0x44, 0x84, 0x48, 0xa7, 0x2a, 0xca, 0x18, 0x99, 0xb2,
	0x9e, 0xcb, 0x58, 0x92, 0x16, 0x8c, 0x2a, 0x9b, 0x33, 0x7c, 0x00, 0xdf, 0x80, 0x8b, 0x43, 0xe6,
	0x20, 0x9a, 0x8b, 0xb1, 0x8a, 0x48, 0x22, 0x09, 0x1f, 0x0b, 0xe7, 0x78, 0x08, 0x5f, 0x81, 0x97,
	0x5b, 0x47, 0x92, 0x52, 0x1e, 0xcf, 0xf6, 0x89, 0x87, 0xe0, 0x09, 0x3c, 0x07, 0xcf, 0xb9, 0x28,
	0xc8, 0x84, 0xd3, 0xe3, 0xe2, 0x4e, 0x6d, 0x2b, 0x09, 0xcf, 0x73, 0x7b, 0x83, 0xe5, 0x13, 0x22,
	0x65, 0xaf, 0xff, 0xd5, 0xca, 0xa3, 0xcb, 0x3f, 0x03, 0xe0, 0x5f, 0xeb, 0x25, 0xba, 0x7b, 0x16,
	0x97, 0xcf, 0x8e, 0x1e, 0x54, 0xda, 0x29, 0xc8, 0xc1, 0x17, 0xba, 0x75, 0x36, 0xfa, 0xa6, 0x6c,
	0x1b, 0xa4, 0xbb, 0x26, 0x68, 0xea, 0xb6, 0x9f, 0xd1, 0x6e, 0x1f, 0x56, 0x0b, 0xf3, 0xbf, 0xf5,
	0x78, 0xef, 0x3e, 0x3f, 0xbc, 0xfb, 0x63, 0x42, 0x36, 0xde, 0x68, 0xec, 0xc2, 0x48, 0x65, 0x90,
	0x93, 0x56, 0x15, 0x21, 0xea, 0xaf, 0x34, 0xbf, 0x76, 0xc0, 0x9c, 0x54, 0x66, 0xbe, 0x07, 0xe6,
	0x45, 0x38, 0x77, 0xc0, 0x6f, 0xcf, 0x77, 0xa7, 0x18, 0x93, 0xca, 0x60, 0xbc, 0x47, 0x30, 0x2e,
	0x42, 0x8c, 0x1d, 0xf4, 0xf5, 0xa4, 0xaf, 0xee, 0xed, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2,
	0x70, 0xf2, 0xc1, 0xbb, 0x02, 0x00, 0x00,
}
