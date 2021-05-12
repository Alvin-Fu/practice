// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/extension_feed_item.proto

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

// An extension feed item.
type ExtensionFeedItem struct {
	// The resource name of the extension feed item.
	// Extension feed item resource names have the form:
	//
	// `customers/{customer_id}/extensionFeedItems/{feed_item_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of this feed item. Read-only.
	Id *wrappers.Int64Value `protobuf:"bytes,24,opt,name=id,proto3" json:"id,omitempty"`
	// The extension type of the extension feed item.
	// This field is read-only.
	ExtensionType enums.ExtensionTypeEnum_ExtensionType `protobuf:"varint,13,opt,name=extension_type,json=extensionType,proto3,enum=google.ads.googleads.v2.enums.ExtensionTypeEnum_ExtensionType" json:"extension_type,omitempty"`
	// Start time in which this feed item is effective and can begin serving. The
	// time is in the customer's time zone.
	// The format is "YYYY-MM-DD HH:MM:SS".
	// Examples: "2018-03-05 09:15:00" or "2018-02-01 14:34:30"
	StartDateTime *wrappers.StringValue `protobuf:"bytes,5,opt,name=start_date_time,json=startDateTime,proto3" json:"start_date_time,omitempty"`
	// End time in which this feed item is no longer effective and will stop
	// serving. The time is in the customer's time zone.
	// The format is "YYYY-MM-DD HH:MM:SS".
	// Examples: "2018-03-05 09:15:00" or "2018-02-01 14:34:30"
	EndDateTime *wrappers.StringValue `protobuf:"bytes,6,opt,name=end_date_time,json=endDateTime,proto3" json:"end_date_time,omitempty"`
	// List of non-overlapping schedules specifying all time intervals
	// for which the feed item may serve. There can be a maximum of 6 schedules
	// per day.
	AdSchedules []*common.AdScheduleInfo `protobuf:"bytes,16,rep,name=ad_schedules,json=adSchedules,proto3" json:"ad_schedules,omitempty"`
	// The targeted device.
	Device enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice `protobuf:"varint,17,opt,name=device,proto3,enum=google.ads.googleads.v2.enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice" json:"device,omitempty"`
	// The targeted geo target constant.
	TargetedGeoTargetConstant *wrappers.StringValue `protobuf:"bytes,20,opt,name=targeted_geo_target_constant,json=targetedGeoTargetConstant,proto3" json:"targeted_geo_target_constant,omitempty"`
	// The targeted keyword.
	TargetedKeyword *common.KeywordInfo `protobuf:"bytes,22,opt,name=targeted_keyword,json=targetedKeyword,proto3" json:"targeted_keyword,omitempty"`
	// Status of the feed item.
	// This field is read-only.
	Status enums.FeedItemStatusEnum_FeedItemStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.FeedItemStatusEnum_FeedItemStatus" json:"status,omitempty"`
	// Extension type.
	//
	// Types that are valid to be assigned to Extension:
	//	*ExtensionFeedItem_SitelinkFeedItem
	//	*ExtensionFeedItem_StructuredSnippetFeedItem
	//	*ExtensionFeedItem_AppFeedItem
	//	*ExtensionFeedItem_CallFeedItem
	//	*ExtensionFeedItem_CalloutFeedItem
	//	*ExtensionFeedItem_TextMessageFeedItem
	//	*ExtensionFeedItem_PriceFeedItem
	//	*ExtensionFeedItem_PromotionFeedItem
	//	*ExtensionFeedItem_LocationFeedItem
	//	*ExtensionFeedItem_AffiliateLocationFeedItem
	//	*ExtensionFeedItem_HotelCalloutFeedItem
	Extension isExtensionFeedItem_Extension `protobuf_oneof:"extension"`
	// Targeting at either the campaign or ad group level. Feed items that target
	// a campaign or ad group will only serve with that resource.
	//
	// Types that are valid to be assigned to ServingResourceTargeting:
	//	*ExtensionFeedItem_TargetedCampaign
	//	*ExtensionFeedItem_TargetedAdGroup
	ServingResourceTargeting isExtensionFeedItem_ServingResourceTargeting `protobuf_oneof:"serving_resource_targeting"`
	XXX_NoUnkeyedLiteral     struct{}                                     `json:"-"`
	XXX_unrecognized         []byte                                       `json:"-"`
	XXX_sizecache            int32                                        `json:"-"`
}

func (m *ExtensionFeedItem) Reset()         { *m = ExtensionFeedItem{} }
func (m *ExtensionFeedItem) String() string { return proto.CompactTextString(m) }
func (*ExtensionFeedItem) ProtoMessage()    {}
func (*ExtensionFeedItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_84144755971e2200, []int{0}
}

func (m *ExtensionFeedItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionFeedItem.Unmarshal(m, b)
}
func (m *ExtensionFeedItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionFeedItem.Marshal(b, m, deterministic)
}
func (m *ExtensionFeedItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionFeedItem.Merge(m, src)
}
func (m *ExtensionFeedItem) XXX_Size() int {
	return xxx_messageInfo_ExtensionFeedItem.Size(m)
}
func (m *ExtensionFeedItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionFeedItem.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionFeedItem proto.InternalMessageInfo

func (m *ExtensionFeedItem) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ExtensionFeedItem) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ExtensionFeedItem) GetExtensionType() enums.ExtensionTypeEnum_ExtensionType {
	if m != nil {
		return m.ExtensionType
	}
	return enums.ExtensionTypeEnum_UNSPECIFIED
}

func (m *ExtensionFeedItem) GetStartDateTime() *wrappers.StringValue {
	if m != nil {
		return m.StartDateTime
	}
	return nil
}

func (m *ExtensionFeedItem) GetEndDateTime() *wrappers.StringValue {
	if m != nil {
		return m.EndDateTime
	}
	return nil
}

func (m *ExtensionFeedItem) GetAdSchedules() []*common.AdScheduleInfo {
	if m != nil {
		return m.AdSchedules
	}
	return nil
}

func (m *ExtensionFeedItem) GetDevice() enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice {
	if m != nil {
		return m.Device
	}
	return enums.FeedItemTargetDeviceEnum_UNSPECIFIED
}

func (m *ExtensionFeedItem) GetTargetedGeoTargetConstant() *wrappers.StringValue {
	if m != nil {
		return m.TargetedGeoTargetConstant
	}
	return nil
}

func (m *ExtensionFeedItem) GetTargetedKeyword() *common.KeywordInfo {
	if m != nil {
		return m.TargetedKeyword
	}
	return nil
}

func (m *ExtensionFeedItem) GetStatus() enums.FeedItemStatusEnum_FeedItemStatus {
	if m != nil {
		return m.Status
	}
	return enums.FeedItemStatusEnum_UNSPECIFIED
}

type isExtensionFeedItem_Extension interface {
	isExtensionFeedItem_Extension()
}

type ExtensionFeedItem_SitelinkFeedItem struct {
	SitelinkFeedItem *common.SitelinkFeedItem `protobuf:"bytes,2,opt,name=sitelink_feed_item,json=sitelinkFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_StructuredSnippetFeedItem struct {
	StructuredSnippetFeedItem *common.StructuredSnippetFeedItem `protobuf:"bytes,3,opt,name=structured_snippet_feed_item,json=structuredSnippetFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_AppFeedItem struct {
	AppFeedItem *common.AppFeedItem `protobuf:"bytes,7,opt,name=app_feed_item,json=appFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_CallFeedItem struct {
	CallFeedItem *common.CallFeedItem `protobuf:"bytes,8,opt,name=call_feed_item,json=callFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_CalloutFeedItem struct {
	CalloutFeedItem *common.CalloutFeedItem `protobuf:"bytes,9,opt,name=callout_feed_item,json=calloutFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_TextMessageFeedItem struct {
	TextMessageFeedItem *common.TextMessageFeedItem `protobuf:"bytes,10,opt,name=text_message_feed_item,json=textMessageFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_PriceFeedItem struct {
	PriceFeedItem *common.PriceFeedItem `protobuf:"bytes,11,opt,name=price_feed_item,json=priceFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_PromotionFeedItem struct {
	PromotionFeedItem *common.PromotionFeedItem `protobuf:"bytes,12,opt,name=promotion_feed_item,json=promotionFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_LocationFeedItem struct {
	LocationFeedItem *common.LocationFeedItem `protobuf:"bytes,14,opt,name=location_feed_item,json=locationFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_AffiliateLocationFeedItem struct {
	AffiliateLocationFeedItem *common.AffiliateLocationFeedItem `protobuf:"bytes,15,opt,name=affiliate_location_feed_item,json=affiliateLocationFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_HotelCalloutFeedItem struct {
	HotelCalloutFeedItem *common.HotelCalloutFeedItem `protobuf:"bytes,23,opt,name=hotel_callout_feed_item,json=hotelCalloutFeedItem,proto3,oneof"`
}

func (*ExtensionFeedItem_SitelinkFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_StructuredSnippetFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_AppFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_CallFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_CalloutFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_TextMessageFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_PriceFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_PromotionFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_LocationFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_AffiliateLocationFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_HotelCalloutFeedItem) isExtensionFeedItem_Extension() {}

func (m *ExtensionFeedItem) GetExtension() isExtensionFeedItem_Extension {
	if m != nil {
		return m.Extension
	}
	return nil
}

func (m *ExtensionFeedItem) GetSitelinkFeedItem() *common.SitelinkFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_SitelinkFeedItem); ok {
		return x.SitelinkFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetStructuredSnippetFeedItem() *common.StructuredSnippetFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_StructuredSnippetFeedItem); ok {
		return x.StructuredSnippetFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetAppFeedItem() *common.AppFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_AppFeedItem); ok {
		return x.AppFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetCallFeedItem() *common.CallFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_CallFeedItem); ok {
		return x.CallFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetCalloutFeedItem() *common.CalloutFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_CalloutFeedItem); ok {
		return x.CalloutFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetTextMessageFeedItem() *common.TextMessageFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_TextMessageFeedItem); ok {
		return x.TextMessageFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetPriceFeedItem() *common.PriceFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_PriceFeedItem); ok {
		return x.PriceFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetPromotionFeedItem() *common.PromotionFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_PromotionFeedItem); ok {
		return x.PromotionFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetLocationFeedItem() *common.LocationFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_LocationFeedItem); ok {
		return x.LocationFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetAffiliateLocationFeedItem() *common.AffiliateLocationFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_AffiliateLocationFeedItem); ok {
		return x.AffiliateLocationFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetHotelCalloutFeedItem() *common.HotelCalloutFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_HotelCalloutFeedItem); ok {
		return x.HotelCalloutFeedItem
	}
	return nil
}

type isExtensionFeedItem_ServingResourceTargeting interface {
	isExtensionFeedItem_ServingResourceTargeting()
}

type ExtensionFeedItem_TargetedCampaign struct {
	TargetedCampaign *wrappers.StringValue `protobuf:"bytes,18,opt,name=targeted_campaign,json=targetedCampaign,proto3,oneof"`
}

type ExtensionFeedItem_TargetedAdGroup struct {
	TargetedAdGroup *wrappers.StringValue `protobuf:"bytes,19,opt,name=targeted_ad_group,json=targetedAdGroup,proto3,oneof"`
}

func (*ExtensionFeedItem_TargetedCampaign) isExtensionFeedItem_ServingResourceTargeting() {}

func (*ExtensionFeedItem_TargetedAdGroup) isExtensionFeedItem_ServingResourceTargeting() {}

func (m *ExtensionFeedItem) GetServingResourceTargeting() isExtensionFeedItem_ServingResourceTargeting {
	if m != nil {
		return m.ServingResourceTargeting
	}
	return nil
}

func (m *ExtensionFeedItem) GetTargetedCampaign() *wrappers.StringValue {
	if x, ok := m.GetServingResourceTargeting().(*ExtensionFeedItem_TargetedCampaign); ok {
		return x.TargetedCampaign
	}
	return nil
}

func (m *ExtensionFeedItem) GetTargetedAdGroup() *wrappers.StringValue {
	if x, ok := m.GetServingResourceTargeting().(*ExtensionFeedItem_TargetedAdGroup); ok {
		return x.TargetedAdGroup
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExtensionFeedItem) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExtensionFeedItem_SitelinkFeedItem)(nil),
		(*ExtensionFeedItem_StructuredSnippetFeedItem)(nil),
		(*ExtensionFeedItem_AppFeedItem)(nil),
		(*ExtensionFeedItem_CallFeedItem)(nil),
		(*ExtensionFeedItem_CalloutFeedItem)(nil),
		(*ExtensionFeedItem_TextMessageFeedItem)(nil),
		(*ExtensionFeedItem_PriceFeedItem)(nil),
		(*ExtensionFeedItem_PromotionFeedItem)(nil),
		(*ExtensionFeedItem_LocationFeedItem)(nil),
		(*ExtensionFeedItem_AffiliateLocationFeedItem)(nil),
		(*ExtensionFeedItem_HotelCalloutFeedItem)(nil),
		(*ExtensionFeedItem_TargetedCampaign)(nil),
		(*ExtensionFeedItem_TargetedAdGroup)(nil),
	}
}

func init() {
	proto.RegisterType((*ExtensionFeedItem)(nil), "google.ads.googleads.v2.resources.ExtensionFeedItem")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/extension_feed_item.proto", fileDescriptor_84144755971e2200)
}

var fileDescriptor_84144755971e2200 = []byte{
	// 969 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xc7, 0x67, 0x67, 0xcb, 0x16, 0x3a, 0x8e, 0x63, 0xa6, 0x48, 0x95, 0xcc, 0x18, 0xd2, 0x0d,
	0x05, 0x02, 0x74, 0x95, 0x36, 0x37, 0x18, 0x30, 0x07, 0x18, 0xea, 0x24, 0x5d, 0x92, 0xb6, 0x1b,
	0x52, 0xdb, 0xf0, 0x86, 0x21, 0x85, 0xc6, 0x8a, 0xc7, 0x0a, 0x57, 0x89, 0x14, 0x44, 0x2a, 0x4d,
	0x80, 0xdd, 0xed, 0x4d, 0x76, 0xb9, 0xd7, 0xd8, 0xdd, 0x1e, 0x65, 0x4f, 0x51, 0x88, 0xfa, 0x08,
	0x1d, 0xc7, 0x95, 0xef, 0xc8, 0x23, 0xfe, 0x7f, 0xff, 0xc3, 0xc3, 0x0f, 0x11, 0xed, 0xfb, 0x42,
	0xf8, 0x01, 0x38, 0x84, 0x4a, 0x27, 0x6b, 0xa6, 0xad, 0xcb, 0xae, 0x13, 0x83, 0x14, 0x49, 0xec,
	0x81, 0x74, 0xe0, 0x4a, 0x01, 0x97, 0x4c, 0x70, 0x77, 0x02, 0x40, 0x5d, 0xa6, 0x20, 0xb4, 0xa3,
	0x58, 0x28, 0x81, 0x1f, 0x64, 0x0a, 0x9b, 0x50, 0x69, 0x97, 0x62, 0xfb, 0xb2, 0x6b, 0x97, 0xe2,
	0xed, 0xc7, 0xf3, 0xf8, 0x9e, 0x08, 0x43, 0xc1, 0x1d, 0x2f, 0x66, 0x0a, 0x62, 0x46, 0x32, 0xe2,
	0xb6, 0x53, 0x31, 0xbc, 0xcc, 0x45, 0xe6, 0x82, 0xee, 0x3c, 0x01, 0xf0, 0x24, 0x34, 0x73, 0x57,
	0xd7, 0x11, 0xe4, 0x9a, 0xbd, 0x0f, 0x6b, 0xca, 0x59, 0xba, 0x52, 0x11, 0x95, 0x14, 0x4e, 0xfb,
	0x8b, 0xaa, 0x14, 0x89, 0x7d, 0x50, 0x2e, 0x85, 0x4b, 0xe6, 0x15, 0x96, 0x5f, 0xe4, 0x62, 0xdd,
	0x7b, 0x93, 0x4c, 0x9c, 0x77, 0x31, 0x89, 0x22, 0x88, 0x0b, 0x78, 0xa7, 0x80, 0x47, 0xcc, 0x21,
	0x9c, 0x0b, 0x45, 0xd4, 0xcd, 0x24, 0xbf, 0xfc, 0x77, 0x1d, 0xb5, 0x9f, 0x15, 0x33, 0xf9, 0x11,
	0x80, 0x9e, 0x2a, 0x08, 0xf1, 0x57, 0xa8, 0x59, 0xd4, 0xd9, 0xe5, 0x24, 0x04, 0xab, 0xb6, 0x53,
	0xdb, 0x5d, 0x19, 0xac, 0x16, 0xc1, 0x9f, 0x49, 0x08, 0xf8, 0x11, 0xaa, 0x33, 0x6a, 0x59, 0x3b,
	0xb5, 0xdd, 0x46, 0xf7, 0xf3, 0x7c, 0x91, 0xec, 0x22, 0x0b, 0xfb, 0x94, 0xab, 0xef, 0xf6, 0xc6,
	0x24, 0x48, 0x60, 0x50, 0x67, 0x14, 0x03, 0x5a, 0x9b, 0x2e, 0x98, 0xd5, 0xdc, 0xa9, 0xed, 0xae,
	0x75, 0x7f, 0xb0, 0xe7, 0x2d, 0xb4, 0x9e, 0xbb, 0x5d, 0xe6, 0x36, 0xba, 0x8e, 0xe0, 0x19, 0x4f,
	0xc2, 0xe9, 0xc8, 0xa0, 0x09, 0x66, 0x17, 0x1f, 0xa1, 0x96, 0x54, 0x24, 0x56, 0x2e, 0x25, 0x0a,
	0x5c, 0xc5, 0x42, 0xb0, 0x3e, 0xd1, 0x09, 0x76, 0x66, 0x12, 0x1c, 0xaa, 0x98, 0x71, 0x3f, 0xcb,
	0xb0, 0xa9, 0x45, 0x47, 0x44, 0xc1, 0x88, 0x85, 0x80, 0x9f, 0xa2, 0x26, 0x70, 0x6a, 0x30, 0x96,
	0x17, 0x60, 0x34, 0x80, 0xd3, 0x92, 0xf0, 0x0a, 0xad, 0x12, 0xea, 0x4a, 0xef, 0x02, 0x68, 0x12,
	0x80, 0xb4, 0xd6, 0x77, 0x96, 0x76, 0x1b, 0x5d, 0x7b, 0xee, 0x64, 0xb3, 0x3d, 0x68, 0xf7, 0xe9,
	0x30, 0x97, 0x9c, 0xf2, 0x89, 0x18, 0x34, 0x48, 0xd9, 0x97, 0x98, 0xa2, 0xe5, 0x6c, 0xdd, 0xad,
	0xb6, 0xae, 0xdc, 0xcb, 0x8a, 0xca, 0x15, 0x8b, 0x39, 0xd2, 0x7b, 0xe6, 0x48, 0x4b, 0x75, 0x01,
	0xef, 0xfa, 0x30, 0xc8, 0xd9, 0xf8, 0x35, 0xea, 0x64, 0x9b, 0x0c, 0xa8, 0xeb, 0x83, 0x28, 0x76,
	0x9c, 0x27, 0xb8, 0x54, 0x84, 0x2b, 0xeb, 0xde, 0x02, 0x95, 0xd8, 0x2a, 0x08, 0xc7, 0x20, 0x32,
	0x93, 0xc3, 0x5c, 0x8e, 0xc7, 0x68, 0xbd, 0xc4, 0xbf, 0x85, 0xeb, 0x77, 0x22, 0xa6, 0xd6, 0xa6,
	0x46, 0x3e, 0xaa, 0xaa, 0xcd, 0x8b, 0x6c, 0xb8, 0x2e, 0x4c, 0xab, 0x80, 0xe4, 0x41, 0xfc, 0x2b,
	0x5a, 0xce, 0x4e, 0x94, 0xf5, 0xb1, 0x2e, 0xce, 0xd3, 0x05, 0x8b, 0x33, 0xd4, 0xa2, 0xa9, 0xb2,
	0x64, 0xa1, 0x41, 0xce, 0xc3, 0xbf, 0x23, 0x2c, 0x99, 0x82, 0x80, 0xf1, 0xb7, 0x37, 0x97, 0x94,
	0x55, 0xd7, 0x39, 0x7f, 0x53, 0x95, 0xf3, 0x30, 0x57, 0x16, 0xec, 0x93, 0x8f, 0x06, 0xeb, 0xf2,
	0x56, 0x0c, 0xff, 0x89, 0x3a, 0x52, 0xc5, 0x89, 0xa7, 0x92, 0x18, 0xa8, 0x2b, 0x39, 0x8b, 0x22,
	0x50, 0x86, 0xd7, 0x92, 0xf6, 0xfa, 0xbe, 0xd2, 0xab, 0x64, 0x0c, 0x33, 0x84, 0x61, 0xba, 0x25,
	0xe7, 0x7d, 0xc4, 0xaf, 0x50, 0x93, 0x44, 0x91, 0x61, 0xf7, 0xe9, 0x62, 0xcb, 0xd1, 0x8f, 0x22,
	0xc3, 0xa0, 0x41, 0x6e, 0xba, 0x78, 0x84, 0xd6, 0x3c, 0x12, 0x04, 0x06, 0xf3, 0x33, 0xcd, 0xfc,
	0xba, 0x8a, 0x79, 0x48, 0x82, 0xc0, 0x80, 0xae, 0x7a, 0x46, 0x1f, 0xbf, 0x46, 0xed, 0xb4, 0x2f,
	0x12, 0xb3, 0x36, 0x2b, 0x1a, 0xec, 0x2c, 0x02, 0x16, 0x89, 0x59, 0x91, 0x96, 0x37, 0x1d, 0xc2,
	0x7f, 0xa0, 0x4d, 0x05, 0x57, 0xca, 0x0d, 0x41, 0x4a, 0xe2, 0x83, 0xe1, 0x81, 0xb4, 0xc7, 0x93,
	0x2a, 0x8f, 0x11, 0x5c, 0xa9, 0x9f, 0x32, 0xb1, 0xe1, 0xb3, 0xa1, 0x66, 0xc3, 0xf8, 0x17, 0xd4,
	0x8a, 0x62, 0xe6, 0x99, 0x26, 0x0d, 0x6d, 0xf2, 0xb8, 0xca, 0xe4, 0x2c, 0x95, 0x19, 0xf8, 0x66,
	0x64, 0x06, 0xb0, 0x87, 0x36, 0xa2, 0x58, 0x84, 0x42, 0x4d, 0xfd, 0x52, 0xad, 0x55, 0x0d, 0xff,
	0xb6, 0x1a, 0x9e, 0x4b, 0x0d, 0x83, 0x76, 0x74, 0x3b, 0x98, 0x9e, 0x88, 0x40, 0x78, 0xe4, 0x96,
	0xc7, 0xda, 0x62, 0x27, 0xe2, 0x65, 0xae, 0x34, 0x4f, 0x44, 0x70, 0x2b, 0x96, 0x9e, 0x08, 0x32,
	0x99, 0xb0, 0x80, 0xa5, 0x17, 0xf0, 0x1d, 0x5e, 0xad, 0xc5, 0x4e, 0x44, 0xbf, 0x60, 0xdc, 0x61,
	0xba, 0x45, 0xe6, 0x7d, 0xc4, 0x21, 0xba, 0x7f, 0x21, 0x14, 0x04, 0xee, 0xec, 0x76, 0xbb, 0xaf,
	0x8d, 0xf7, 0xaa, 0x8c, 0x4f, 0x52, 0xf9, 0xec, 0x9e, 0xbb, 0x77, 0x71, 0x47, 0x1c, 0xbf, 0x40,
	0xed, 0xf2, 0x4a, 0xf4, 0x48, 0x18, 0x11, 0xe6, 0x73, 0x0b, 0x57, 0x5f, 0xb3, 0x27, 0xb5, 0x41,
	0x79, 0x97, 0x1e, 0xe6, 0x3a, 0xfc, 0xdc, 0x80, 0x11, 0xea, 0xfa, 0xb1, 0x48, 0x22, 0x6b, 0x63,
	0x21, 0x58, 0x79, 0xa7, 0xf6, 0xe9, 0x71, 0x2a, 0x3b, 0x68, 0xa0, 0x95, 0xf2, 0xe7, 0x7a, 0xd0,
	0x41, 0xdb, 0x12, 0xe2, 0x4b, 0xc6, 0x7d, 0xb7, 0x7c, 0x19, 0x64, 0x02, 0xc6, 0xfd, 0x83, 0xbf,
	0xea, 0xe8, 0xa1, 0x27, 0x42, 0xbb, 0xf2, 0xd1, 0x76, 0xb0, 0x39, 0xf3, 0xd8, 0x38, 0x4b, 0xd3,
	0x39, 0xab, 0xfd, 0xf6, 0x3c, 0x17, 0xfb, 0x22, 0x20, 0xdc, 0xb7, 0x45, 0xec, 0x3b, 0x3e, 0x70,
	0x9d, 0x6c, 0xf1, 0x28, 0x8a, 0x98, 0xfc, 0xc0, 0x6b, 0x72, 0xbf, 0x6c, 0xfd, 0x5d, 0x5f, 0x3a,
	0xee, 0xf7, 0xff, 0xa9, 0x3f, 0x38, 0xce, 0x90, 0x7d, 0x2a, 0xed, 0xac, 0x99, 0xb6, 0xc6, 0x5d,
	0x7b, 0x50, 0x8c, 0xfc, 0xaf, 0x18, 0x73, 0xde, 0xa7, 0xf2, 0xbc, 0x1c, 0x73, 0x3e, 0xee, 0x9e,
	0x97, 0x63, 0xfe, 0xaf, 0x3f, 0xcc, 0x3e, 0xf4, 0x7a, 0x7d, 0x2a, 0x7b, 0xbd, 0x72, 0x54, 0xaf,
	0x37, 0xee, 0xf6, 0x7a, 0xe5, 0xb8, 0x37, 0xcb, 0x3a, 0xd9, 0x27, 0xef, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xef, 0x14, 0xa3, 0x5c, 0xf9, 0x0a, 0x00, 0x00,
}
