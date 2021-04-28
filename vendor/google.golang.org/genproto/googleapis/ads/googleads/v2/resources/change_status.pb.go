// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/change_status.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Describes the status of returned resource.
type ChangeStatus struct {
	// The resource name of the change status.
	// Change status resource names have the form:
	//
	// `customers/{customer_id}/changeStatus/{change_status_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Time at which the most recent change has occurred on this resource.
	LastChangeDateTime *wrappers.StringValue `protobuf:"bytes,3,opt,name=last_change_date_time,json=lastChangeDateTime,proto3" json:"last_change_date_time,omitempty"`
	// Represents the type of the changed resource. This dictates what fields
	// will be set. For example, for AD_GROUP, campaign and ad_group fields will
	// be set.
	ResourceType enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType `protobuf:"varint,4,opt,name=resource_type,json=resourceType,proto3,enum=google.ads.googleads.v2.enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType" json:"resource_type,omitempty"`
	// The Campaign affected by this change.
	Campaign *wrappers.StringValue `protobuf:"bytes,5,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The AdGroup affected by this change.
	AdGroup *wrappers.StringValue `protobuf:"bytes,6,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Represents the status of the changed resource.
	ResourceStatus enums.ChangeStatusOperationEnum_ChangeStatusOperation `protobuf:"varint,8,opt,name=resource_status,json=resourceStatus,proto3,enum=google.ads.googleads.v2.enums.ChangeStatusOperationEnum_ChangeStatusOperation" json:"resource_status,omitempty"`
	// The AdGroupAd affected by this change.
	AdGroupAd *wrappers.StringValue `protobuf:"bytes,9,opt,name=ad_group_ad,json=adGroupAd,proto3" json:"ad_group_ad,omitempty"`
	// The AdGroupCriterion affected by this change.
	AdGroupCriterion *wrappers.StringValue `protobuf:"bytes,10,opt,name=ad_group_criterion,json=adGroupCriterion,proto3" json:"ad_group_criterion,omitempty"`
	// The CampaignCriterion affected by this change.
	CampaignCriterion *wrappers.StringValue `protobuf:"bytes,11,opt,name=campaign_criterion,json=campaignCriterion,proto3" json:"campaign_criterion,omitempty"`
	// The Feed affected by this change.
	Feed *wrappers.StringValue `protobuf:"bytes,12,opt,name=feed,proto3" json:"feed,omitempty"`
	// The FeedItem affected by this change.
	FeedItem *wrappers.StringValue `protobuf:"bytes,13,opt,name=feed_item,json=feedItem,proto3" json:"feed_item,omitempty"`
	// The AdGroupFeed affected by this change.
	AdGroupFeed *wrappers.StringValue `protobuf:"bytes,14,opt,name=ad_group_feed,json=adGroupFeed,proto3" json:"ad_group_feed,omitempty"`
	// The CampaignFeed affected by this change.
	CampaignFeed *wrappers.StringValue `protobuf:"bytes,15,opt,name=campaign_feed,json=campaignFeed,proto3" json:"campaign_feed,omitempty"`
	// The AdGroupBidModifier affected by this change.
	AdGroupBidModifier   *wrappers.StringValue `protobuf:"bytes,16,opt,name=ad_group_bid_modifier,json=adGroupBidModifier,proto3" json:"ad_group_bid_modifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ChangeStatus) Reset()         { *m = ChangeStatus{} }
func (m *ChangeStatus) String() string { return proto.CompactTextString(m) }
func (*ChangeStatus) ProtoMessage()    {}
func (*ChangeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_edc122555564308c, []int{0}
}

func (m *ChangeStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeStatus.Unmarshal(m, b)
}
func (m *ChangeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeStatus.Marshal(b, m, deterministic)
}
func (m *ChangeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeStatus.Merge(m, src)
}
func (m *ChangeStatus) XXX_Size() int {
	return xxx_messageInfo_ChangeStatus.Size(m)
}
func (m *ChangeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeStatus proto.InternalMessageInfo

func (m *ChangeStatus) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ChangeStatus) GetLastChangeDateTime() *wrappers.StringValue {
	if m != nil {
		return m.LastChangeDateTime
	}
	return nil
}

func (m *ChangeStatus) GetResourceType() enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType {
	if m != nil {
		return m.ResourceType
	}
	return enums.ChangeStatusResourceTypeEnum_UNSPECIFIED
}

func (m *ChangeStatus) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *ChangeStatus) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *ChangeStatus) GetResourceStatus() enums.ChangeStatusOperationEnum_ChangeStatusOperation {
	if m != nil {
		return m.ResourceStatus
	}
	return enums.ChangeStatusOperationEnum_UNSPECIFIED
}

func (m *ChangeStatus) GetAdGroupAd() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupAd
	}
	return nil
}

func (m *ChangeStatus) GetAdGroupCriterion() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupCriterion
	}
	return nil
}

func (m *ChangeStatus) GetCampaignCriterion() *wrappers.StringValue {
	if m != nil {
		return m.CampaignCriterion
	}
	return nil
}

func (m *ChangeStatus) GetFeed() *wrappers.StringValue {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *ChangeStatus) GetFeedItem() *wrappers.StringValue {
	if m != nil {
		return m.FeedItem
	}
	return nil
}

func (m *ChangeStatus) GetAdGroupFeed() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupFeed
	}
	return nil
}

func (m *ChangeStatus) GetCampaignFeed() *wrappers.StringValue {
	if m != nil {
		return m.CampaignFeed
	}
	return nil
}

func (m *ChangeStatus) GetAdGroupBidModifier() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupBidModifier
	}
	return nil
}

func init() {
	proto.RegisterType((*ChangeStatus)(nil), "google.ads.googleads.v2.resources.ChangeStatus")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/change_status.proto", fileDescriptor_edc122555564308c)
}

var fileDescriptor_edc122555564308c = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0x14, 0x3f,
	0x14, 0x67, 0xda, 0xfe, 0xfb, 0xdf, 0xcd, 0xee, 0xf6, 0x23, 0x50, 0x18, 0x4a, 0x91, 0x56, 0x29,
	0xf4, 0x2a, 0x23, 0x23, 0xa2, 0x4e, 0x05, 0x9d, 0x56, 0x2d, 0x56, 0x6c, 0xcb, 0xb6, 0xec, 0x85,
	0x2c, 0x0c, 0xe9, 0x26, 0x1d, 0x03, 0x3b, 0xc9, 0x90, 0x64, 0x5a, 0xfa, 0x00, 0xbe, 0x88, 0x97,
	0x3e, 0x8a, 0x8f, 0xe2, 0x43, 0x88, 0x4c, 0x26, 0x49, 0x17, 0x65, 0x75, 0x7a, 0xb5, 0x67, 0xf6,
	0xfc, 0xbe, 0xce, 0xc9, 0x4c, 0xc0, 0xd3, 0x5c, 0x88, 0x7c, 0x4a, 0x23, 0x4c, 0x54, 0xd4, 0x94,
	0x75, 0x75, 0x1d, 0x47, 0x92, 0x2a, 0x51, 0xc9, 0x09, 0x55, 0xd1, 0xe4, 0x33, 0xe6, 0x39, 0xcd,
	0x94, 0xc6, 0xba, 0x52, 0xa8, 0x94, 0x42, 0x0b, 0xb8, 0xd3, 0x60, 0x11, 0x26, 0x0a, 0x79, 0x1a,
	0xba, 0x8e, 0x91, 0xa7, 0x6d, 0xee, 0xcf, 0x53, 0xa6, 0xbc, 0x2a, 0x7e, 0x53, 0xcd, 0x44, 0x49,
	0x25, 0xd6, 0x4c, 0xf0, 0x46, 0x7f, 0xf3, 0xd5, 0x7d, 0xc8, 0xce, 0x33, 0xd3, 0xb7, 0x25, 0xb5,
	0x02, 0x0f, 0xac, 0x80, 0x79, 0xba, 0xac, 0xae, 0xa2, 0x1b, 0x89, 0xcb, 0x92, 0x4a, 0x3b, 0xc0,
	0xe6, 0x96, 0x33, 0x28, 0x59, 0x84, 0x39, 0x17, 0xda, 0xb8, 0xdb, 0xee, 0xc3, 0x2f, 0x1d, 0xd0,
	0x3f, 0x34, 0x1e, 0xe7, 0xc6, 0x02, 0x3e, 0x02, 0x03, 0xef, 0xc2, 0x71, 0x41, 0xc3, 0x60, 0x3b,
	0xd8, 0xeb, 0x0e, 0xfb, 0xee, 0xcf, 0x13, 0x5c, 0x50, 0x78, 0x0a, 0x36, 0xa6, 0x58, 0xe9, 0xcc,
	0xa6, 0x23, 0x58, 0xd3, 0x4c, 0xb3, 0x82, 0x86, 0x8b, 0xdb, 0xc1, 0x5e, 0x2f, 0xde, 0xb2, 0x9b,
	0x42, 0x2e, 0x13, 0x3a, 0xd7, 0x92, 0xf1, 0x7c, 0x84, 0xa7, 0x15, 0x1d, 0xc2, 0x9a, 0xda, 0x78,
	0xbe, 0xc1, 0x9a, 0x5e, 0xb0, 0x82, 0xc2, 0xdb, 0x19, 0xd7, 0x7a, 0xb6, 0x70, 0x69, 0x3b, 0xd8,
	0x5b, 0x89, 0x2f, 0xd0, 0xbc, 0xed, 0x9b, 0xed, 0xa0, 0xd9, 0xe4, 0x43, 0xcb, 0xbf, 0xb8, 0x2d,
	0xe9, 0x5b, 0x5e, 0x15, 0x73, 0x9b, 0x77, 0xb3, 0xd4, 0x4f, 0xf0, 0x39, 0xe8, 0x4c, 0x70, 0x51,
	0x62, 0x96, 0xf3, 0xf0, 0xbf, 0x16, 0xf1, 0x3d, 0x1a, 0x3e, 0x03, 0x1d, 0x4c, 0xb2, 0x5c, 0x8a,
	0xaa, 0x0c, 0x97, 0x5b, 0x30, 0xff, 0xc7, 0xe4, 0xa8, 0x06, 0xc3, 0x1b, 0xb0, 0xea, 0xa7, 0x6d,
	0x4e, 0x36, 0xec, 0x98, 0x79, 0x4f, 0xee, 0x31, 0xef, 0xa9, 0x7b, 0x91, 0xfe, 0x18, 0xd6, 0x77,
	0x86, 0x2b, 0xce, 0xc6, 0x1e, 0xee, 0x4b, 0xd0, 0x73, 0x89, 0x33, 0x4c, 0xc2, 0x6e, 0x8b, 0xd0,
	0x5d, 0x1b, 0x3a, 0x25, 0xf0, 0x18, 0x40, 0xcf, 0x9e, 0x48, 0xa6, 0xa9, 0x64, 0x82, 0x87, 0xa0,
	0x85, 0xc8, 0x9a, 0x15, 0x39, 0x74, 0x2c, 0xf8, 0x01, 0x40, 0xb7, 0xc7, 0x19, 0xad, 0x5e, 0x0b,
	0xad, 0x75, 0xc7, 0xbb, 0x13, 0x7b, 0x0c, 0x96, 0xae, 0x28, 0x25, 0x61, 0xbf, 0x05, 0xdd, 0x20,
	0xe1, 0x0b, 0xd0, 0xad, 0x7f, 0x33, 0xa6, 0x69, 0x11, 0x0e, 0xda, 0x9c, 0x7a, 0x0d, 0x7f, 0xaf,
	0x69, 0x01, 0x5f, 0x83, 0x81, 0xdf, 0x82, 0x71, 0x5d, 0x69, 0x41, 0xef, 0xd9, 0x05, 0xbc, 0xab,
	0xcd, 0x53, 0x30, 0xf0, 0xb3, 0x1b, 0x85, 0xd5, 0x16, 0x0a, 0x7d, 0x47, 0x31, 0x12, 0xa7, 0x60,
	0xc3, 0x87, 0xb8, 0x64, 0x24, 0x2b, 0x04, 0x61, 0x57, 0x8c, 0xca, 0x70, 0xad, 0xcd, 0x07, 0x68,
	0xc3, 0x1c, 0x30, 0xf2, 0xd1, 0xf2, 0x0e, 0x7e, 0x06, 0x60, 0x77, 0x22, 0x0a, 0xf4, 0xcf, 0xdb,
	0xee, 0x60, 0x7d, 0xf6, 0x55, 0x3b, 0xab, 0xf5, 0xcf, 0x82, 0x4f, 0xc7, 0x96, 0x97, 0x8b, 0x29,
	0xe6, 0x39, 0x12, 0x32, 0x8f, 0x72, 0xca, 0x8d, 0xbb, 0xbb, 0xd5, 0x4a, 0xa6, 0xfe, 0x72, 0xf7,
	0xee, 0xfb, 0xea, 0xeb, 0xc2, 0xe2, 0x51, 0x9a, 0x7e, 0x5b, 0xd8, 0x39, 0x6a, 0x24, 0x53, 0xa2,
	0x50, 0x53, 0xd6, 0xd5, 0x28, 0x46, 0xee, 0x5b, 0x56, 0xdf, 0x1d, 0x66, 0x9c, 0x12, 0x35, 0xf6,
	0x98, 0xf1, 0x28, 0x1e, 0x7b, 0xcc, 0x8f, 0x85, 0xdd, 0xa6, 0x91, 0x24, 0x29, 0x51, 0x49, 0xe2,
	0x51, 0x49, 0x32, 0x8a, 0x93, 0xc4, 0xe3, 0x2e, 0x97, 0x4d, 0xd8, 0x27, 0xbf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x47, 0x4a, 0x30, 0xf6, 0x27, 0x06, 0x00, 0x00,
}
