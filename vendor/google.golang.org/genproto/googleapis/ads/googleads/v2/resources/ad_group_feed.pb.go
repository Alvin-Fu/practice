// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/ad_group_feed.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// An ad group feed.
type AdGroupFeed struct {
	// The resource name of the ad group feed.
	// Ad group feed resource names have the form:
	//
	// `customers/{customer_id}/adGroupFeeds/{ad_group_id}~{feed_id}
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The feed being linked to the ad group.
	Feed *wrappers.StringValue `protobuf:"bytes,2,opt,name=feed,proto3" json:"feed,omitempty"`
	// The ad group being linked to the feed.
	AdGroup *wrappers.StringValue `protobuf:"bytes,3,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Indicates which placeholder types the feed may populate under the connected
	// ad group. Required.
	PlaceholderTypes []enums.PlaceholderTypeEnum_PlaceholderType `protobuf:"varint,4,rep,packed,name=placeholder_types,json=placeholderTypes,proto3,enum=google.ads.googleads.v2.enums.PlaceholderTypeEnum_PlaceholderType" json:"placeholder_types,omitempty"`
	// Matching function associated with the AdGroupFeed.
	// The matching function is used to filter the set of feed items selected.
	// Required.
	MatchingFunction *common.MatchingFunction `protobuf:"bytes,5,opt,name=matching_function,json=matchingFunction,proto3" json:"matching_function,omitempty"`
	// Status of the ad group feed.
	// This field is read-only.
	Status               enums.FeedLinkStatusEnum_FeedLinkStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.FeedLinkStatusEnum_FeedLinkStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *AdGroupFeed) Reset()         { *m = AdGroupFeed{} }
func (m *AdGroupFeed) String() string { return proto.CompactTextString(m) }
func (*AdGroupFeed) ProtoMessage()    {}
func (*AdGroupFeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_8abb159f141fd625, []int{0}
}

func (m *AdGroupFeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupFeed.Unmarshal(m, b)
}
func (m *AdGroupFeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupFeed.Marshal(b, m, deterministic)
}
func (m *AdGroupFeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupFeed.Merge(m, src)
}
func (m *AdGroupFeed) XXX_Size() int {
	return xxx_messageInfo_AdGroupFeed.Size(m)
}
func (m *AdGroupFeed) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupFeed.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupFeed proto.InternalMessageInfo

func (m *AdGroupFeed) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupFeed) GetFeed() *wrappers.StringValue {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *AdGroupFeed) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupFeed) GetPlaceholderTypes() []enums.PlaceholderTypeEnum_PlaceholderType {
	if m != nil {
		return m.PlaceholderTypes
	}
	return nil
}

func (m *AdGroupFeed) GetMatchingFunction() *common.MatchingFunction {
	if m != nil {
		return m.MatchingFunction
	}
	return nil
}

func (m *AdGroupFeed) GetStatus() enums.FeedLinkStatusEnum_FeedLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.FeedLinkStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*AdGroupFeed)(nil), "google.ads.googleads.v2.resources.AdGroupFeed")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/ad_group_feed.proto", fileDescriptor_8abb159f141fd625)
}

var fileDescriptor_8abb159f141fd625 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xc7, 0xc9, 0x6e, 0x5d, 0x75, 0xaa, 0x65, 0x9b, 0xab, 0x50, 0x8a, 0x6c, 0x95, 0xc2, 0x5e,
	0xcd, 0x94, 0xf8, 0x05, 0xf1, 0xc6, 0x2c, 0xd8, 0x05, 0x51, 0x59, 0x52, 0x59, 0x44, 0x56, 0xc2,
	0x34, 0x33, 0x3b, 0x0d, 0xcd, 0x7c, 0x30, 0x93, 0x54, 0xfa, 0x3a, 0x5e, 0xfa, 0x22, 0x82, 0x8f,
	0xe2, 0x3b, 0x08, 0x92, 0x4c, 0x26, 0xda, 0x95, 0xb4, 0xbd, 0x3b, 0x7b, 0xe6, 0xff, 0x3b, 0x7b,
	0xce, 0xf9, 0x9f, 0x80, 0xe7, 0x4c, 0x4a, 0x56, 0x50, 0x84, 0x89, 0x41, 0x36, 0xac, 0xa3, 0x8b,
	0x10, 0x69, 0x6a, 0x64, 0xa5, 0x33, 0x6a, 0x10, 0x26, 0x29, 0xd3, 0xb2, 0x52, 0xe9, 0x9a, 0x52,
	0x02, 0x95, 0x96, 0xa5, 0xf4, 0x0f, 0xac, 0x16, 0x62, 0x62, 0x60, 0x87, 0xc1, 0x8b, 0x10, 0x76,
	0xd8, 0xde, 0x8b, 0xbe, 0xca, 0x99, 0xe4, 0x5c, 0x0a, 0xc4, 0x71, 0x99, 0x9d, 0xe5, 0x82, 0xa5,
	0xeb, 0x4a, 0x64, 0x65, 0x2e, 0x85, 0x2d, 0xbd, 0xf7, 0xac, 0x8f, 0xa3, 0xa2, 0xe2, 0x06, 0xd5,
	0x4d, 0xa4, 0x45, 0x2e, 0xce, 0x53, 0x53, 0xe2, 0xb2, 0x32, 0xb7, 0xa3, 0x54, 0x81, 0x33, 0x7a,
	0x26, 0x0b, 0x42, 0x75, 0x5a, 0x5e, 0x2a, 0xda, 0x52, 0x8f, 0x5a, 0xaa, 0xf9, 0x75, 0x5a, 0xad,
	0xd1, 0x57, 0x8d, 0x95, 0xa2, 0xda, 0x55, 0xdd, 0x77, 0x55, 0x55, 0x8e, 0xb0, 0x10, 0xb2, 0xc4,
	0x75, 0xa3, 0xed, 0xeb, 0xe3, 0x1f, 0x43, 0xb0, 0x1d, 0x93, 0x79, 0xbd, 0x9b, 0x63, 0x4a, 0x89,
	0xff, 0x04, 0x3c, 0x74, 0xe3, 0xa7, 0x02, 0x73, 0x1a, 0x78, 0x13, 0x6f, 0x7a, 0x3f, 0x79, 0xe0,
	0x92, 0x1f, 0x30, 0xa7, 0xfe, 0x11, 0xd8, 0xaa, 0x47, 0x08, 0x06, 0x13, 0x6f, 0xba, 0x1d, 0xee,
	0xb7, 0xdb, 0x83, 0xae, 0x03, 0x78, 0x52, 0xea, 0x5c, 0xb0, 0x25, 0x2e, 0x2a, 0x9a, 0x34, 0x4a,
	0xff, 0x25, 0xb8, 0xe7, 0x2c, 0x08, 0x86, 0xb7, 0xa0, 0xee, 0x62, 0xdb, 0x93, 0x2f, 0xc1, 0xee,
	0xe6, 0xdc, 0x26, 0xd8, 0x9a, 0x0c, 0xa7, 0x3b, 0xe1, 0x0c, 0xf6, 0x19, 0xd8, 0xec, 0x0b, 0x2e,
	0xfe, 0x72, 0x1f, 0x2f, 0x15, 0x7d, 0x23, 0x2a, 0xbe, 0x99, 0x4b, 0xc6, 0xea, 0x6a, 0xc2, 0xf8,
	0x5f, 0xc0, 0xee, 0x7f, 0xae, 0x06, 0x77, 0x9a, 0x96, 0x8f, 0x7a, 0xff, 0xd0, 0x9e, 0x03, 0x7c,
	0xdf, 0x82, 0xc7, 0x2d, 0x97, 0x8c, 0xf9, 0x46, 0xc6, 0xff, 0x04, 0x46, 0xd6, 0xf3, 0x60, 0x34,
	0xf1, 0xa6, 0x3b, 0xe1, 0xeb, 0x1b, 0x86, 0xa8, 0x4d, 0x79, 0x97, 0x8b, 0xf3, 0x93, 0x06, 0x6a,
	0x66, 0xb8, 0x9a, 0x4a, 0xda, 0x7a, 0xb3, 0xdf, 0x1e, 0x38, 0xcc, 0x24, 0x87, 0x37, 0x5e, 0xf5,
	0x6c, 0xfc, 0x8f, 0xe1, 0x8b, 0x7a, 0xfb, 0x0b, 0xef, 0xf3, 0xdb, 0x16, 0x63, 0xb2, 0xc0, 0x82,
	0x41, 0xa9, 0x19, 0x62, 0x54, 0x34, 0xde, 0xb8, 0x5b, 0x54, 0xb9, 0xb9, 0xe6, 0x13, 0x7b, 0xd5,
	0x45, 0xdf, 0x06, 0xc3, 0x79, 0x1c, 0x7f, 0x1f, 0x1c, 0xcc, 0x6d, 0xc9, 0x98, 0x18, 0x68, 0xc3,
	0x3a, 0x5a, 0x86, 0x30, 0x71, 0xca, 0x9f, 0x4e, 0xb3, 0x8a, 0x89, 0x59, 0x75, 0x9a, 0xd5, 0x32,
	0x5c, 0x75, 0x9a, 0x5f, 0x83, 0x43, 0xfb, 0x10, 0x45, 0x31, 0x31, 0x51, 0xd4, 0xa9, 0xa2, 0x68,
	0x19, 0x46, 0x51, 0xa7, 0x3b, 0x1d, 0x35, 0xcd, 0x3e, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x54,
	0x91, 0x72, 0x74, 0x0e, 0x04, 0x00, 0x00,
}
