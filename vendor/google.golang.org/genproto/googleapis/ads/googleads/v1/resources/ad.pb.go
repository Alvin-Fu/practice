// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// An ad.
type Ad struct {
	// The ID of the ad.
	Id *wrappers.Int64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The list of possible final URLs after all cross-domain redirects for the
	// ad.
	FinalUrls []*wrappers.StringValue `protobuf:"bytes,2,rep,name=final_urls,json=finalUrls,proto3" json:"final_urls,omitempty"`
	// A list of final app URLs that will be used on mobile if the user has the
	// specific app installed.
	FinalAppUrls []*common.FinalAppUrl `protobuf:"bytes,35,rep,name=final_app_urls,json=finalAppUrls,proto3" json:"final_app_urls,omitempty"`
	// The list of possible final mobile URLs after all cross-domain redirects
	// for the ad.
	FinalMobileUrls []*wrappers.StringValue `protobuf:"bytes,16,rep,name=final_mobile_urls,json=finalMobileUrls,proto3" json:"final_mobile_urls,omitempty"`
	// The URL template for constructing a tracking URL.
	TrackingUrlTemplate *wrappers.StringValue `protobuf:"bytes,12,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3" json:"tracking_url_template,omitempty"`
	// The list of mappings that can be used to substitute custom parameter tags
	// in a `tracking_url_template`, `final_urls`, or `mobile_final_urls`.
	UrlCustomParameters []*common.CustomParameter `protobuf:"bytes,10,rep,name=url_custom_parameters,json=urlCustomParameters,proto3" json:"url_custom_parameters,omitempty"`
	// The URL that appears in the ad description for some ad formats.
	DisplayUrl *wrappers.StringValue `protobuf:"bytes,4,opt,name=display_url,json=displayUrl,proto3" json:"display_url,omitempty"`
	// The type of ad.
	Type enums.AdTypeEnum_AdType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.AdTypeEnum_AdType" json:"type,omitempty"`
	// Indicates if this ad was automatically added by Google Ads and not by a
	// user. For example, this could happen when ads are automatically created as
	// suggestions for new ads based on knowledge of how existing ads are
	// performing.
	AddedByGoogleAds *wrappers.BoolValue `protobuf:"bytes,19,opt,name=added_by_google_ads,json=addedByGoogleAds,proto3" json:"added_by_google_ads,omitempty"`
	// The device preference for the ad. You can only specify a preference for
	// mobile devices. When this preference is set the ad will be preferred over
	// other ads when being displayed on a mobile device. The ad can still be
	// displayed on other device types, e.g. if no other ads are available.
	// If unspecified (no device preference), all devices are targeted.
	// This is only supported by some ad types.
	DevicePreference enums.DeviceEnum_Device `protobuf:"varint,20,opt,name=device_preference,json=devicePreference,proto3,enum=google.ads.googleads.v1.enums.DeviceEnum_Device" json:"device_preference,omitempty"`
	// Additional URLs for the ad that are tagged with a unique identifier that
	// can be referenced from other fields in the ad.
	UrlCollections []*common.UrlCollection `protobuf:"bytes,26,rep,name=url_collections,json=urlCollections,proto3" json:"url_collections,omitempty"`
	// The name of the ad. This is only used to be able to identify the ad. It
	// does not need to be unique and does not affect the served ad.
	Name *wrappers.StringValue `protobuf:"bytes,23,opt,name=name,proto3" json:"name,omitempty"`
	// If this ad is system managed, then this field will indicate the source.
	// This field is read-only.
	SystemManagedResourceSource enums.SystemManagedResourceSourceEnum_SystemManagedResourceSource `protobuf:"varint,27,opt,name=system_managed_resource_source,json=systemManagedResourceSource,proto3,enum=google.ads.googleads.v1.enums.SystemManagedResourceSourceEnum_SystemManagedResourceSource" json:"system_managed_resource_source,omitempty"`
	// Details pertinent to the ad type. Exactly one value must be set.
	//
	// Types that are valid to be assigned to AdData:
	//	*Ad_TextAd
	//	*Ad_ExpandedTextAd
	//	*Ad_CallOnlyAd
	//	*Ad_ExpandedDynamicSearchAd
	//	*Ad_HotelAd
	//	*Ad_ShoppingSmartAd
	//	*Ad_ShoppingProductAd
	//	*Ad_GmailAd
	//	*Ad_ImageAd
	//	*Ad_VideoAd
	//	*Ad_ResponsiveSearchAd
	//	*Ad_LegacyResponsiveDisplayAd
	//	*Ad_AppAd
	//	*Ad_LegacyAppInstallAd
	//	*Ad_ResponsiveDisplayAd
	//	*Ad_DisplayUploadAd
	//	*Ad_AppEngagementAd
	//	*Ad_ShoppingComparisonListingAd
	AdData               isAd_AdData `protobuf_oneof:"ad_data"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Ad) Reset()         { *m = Ad{} }
func (m *Ad) String() string { return proto.CompactTextString(m) }
func (*Ad) ProtoMessage()    {}
func (*Ad) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cfecfd7772c305a, []int{0}
}

func (m *Ad) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ad.Unmarshal(m, b)
}
func (m *Ad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ad.Marshal(b, m, deterministic)
}
func (m *Ad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ad.Merge(m, src)
}
func (m *Ad) XXX_Size() int {
	return xxx_messageInfo_Ad.Size(m)
}
func (m *Ad) XXX_DiscardUnknown() {
	xxx_messageInfo_Ad.DiscardUnknown(m)
}

var xxx_messageInfo_Ad proto.InternalMessageInfo

func (m *Ad) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Ad) GetFinalUrls() []*wrappers.StringValue {
	if m != nil {
		return m.FinalUrls
	}
	return nil
}

func (m *Ad) GetFinalAppUrls() []*common.FinalAppUrl {
	if m != nil {
		return m.FinalAppUrls
	}
	return nil
}

func (m *Ad) GetFinalMobileUrls() []*wrappers.StringValue {
	if m != nil {
		return m.FinalMobileUrls
	}
	return nil
}

func (m *Ad) GetTrackingUrlTemplate() *wrappers.StringValue {
	if m != nil {
		return m.TrackingUrlTemplate
	}
	return nil
}

func (m *Ad) GetUrlCustomParameters() []*common.CustomParameter {
	if m != nil {
		return m.UrlCustomParameters
	}
	return nil
}

func (m *Ad) GetDisplayUrl() *wrappers.StringValue {
	if m != nil {
		return m.DisplayUrl
	}
	return nil
}

func (m *Ad) GetType() enums.AdTypeEnum_AdType {
	if m != nil {
		return m.Type
	}
	return enums.AdTypeEnum_UNSPECIFIED
}

func (m *Ad) GetAddedByGoogleAds() *wrappers.BoolValue {
	if m != nil {
		return m.AddedByGoogleAds
	}
	return nil
}

func (m *Ad) GetDevicePreference() enums.DeviceEnum_Device {
	if m != nil {
		return m.DevicePreference
	}
	return enums.DeviceEnum_UNSPECIFIED
}

func (m *Ad) GetUrlCollections() []*common.UrlCollection {
	if m != nil {
		return m.UrlCollections
	}
	return nil
}

func (m *Ad) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Ad) GetSystemManagedResourceSource() enums.SystemManagedResourceSourceEnum_SystemManagedResourceSource {
	if m != nil {
		return m.SystemManagedResourceSource
	}
	return enums.SystemManagedResourceSourceEnum_UNSPECIFIED
}

type isAd_AdData interface {
	isAd_AdData()
}

type Ad_TextAd struct {
	TextAd *common.TextAdInfo `protobuf:"bytes,6,opt,name=text_ad,json=textAd,proto3,oneof"`
}

type Ad_ExpandedTextAd struct {
	ExpandedTextAd *common.ExpandedTextAdInfo `protobuf:"bytes,7,opt,name=expanded_text_ad,json=expandedTextAd,proto3,oneof"`
}

type Ad_CallOnlyAd struct {
	CallOnlyAd *common.CallOnlyAdInfo `protobuf:"bytes,13,opt,name=call_only_ad,json=callOnlyAd,proto3,oneof"`
}

type Ad_ExpandedDynamicSearchAd struct {
	ExpandedDynamicSearchAd *common.ExpandedDynamicSearchAdInfo `protobuf:"bytes,14,opt,name=expanded_dynamic_search_ad,json=expandedDynamicSearchAd,proto3,oneof"`
}

type Ad_HotelAd struct {
	HotelAd *common.HotelAdInfo `protobuf:"bytes,15,opt,name=hotel_ad,json=hotelAd,proto3,oneof"`
}

type Ad_ShoppingSmartAd struct {
	ShoppingSmartAd *common.ShoppingSmartAdInfo `protobuf:"bytes,17,opt,name=shopping_smart_ad,json=shoppingSmartAd,proto3,oneof"`
}

type Ad_ShoppingProductAd struct {
	ShoppingProductAd *common.ShoppingProductAdInfo `protobuf:"bytes,18,opt,name=shopping_product_ad,json=shoppingProductAd,proto3,oneof"`
}

type Ad_GmailAd struct {
	GmailAd *common.GmailAdInfo `protobuf:"bytes,21,opt,name=gmail_ad,json=gmailAd,proto3,oneof"`
}

type Ad_ImageAd struct {
	ImageAd *common.ImageAdInfo `protobuf:"bytes,22,opt,name=image_ad,json=imageAd,proto3,oneof"`
}

type Ad_VideoAd struct {
	VideoAd *common.VideoAdInfo `protobuf:"bytes,24,opt,name=video_ad,json=videoAd,proto3,oneof"`
}

type Ad_ResponsiveSearchAd struct {
	ResponsiveSearchAd *common.ResponsiveSearchAdInfo `protobuf:"bytes,25,opt,name=responsive_search_ad,json=responsiveSearchAd,proto3,oneof"`
}

type Ad_LegacyResponsiveDisplayAd struct {
	LegacyResponsiveDisplayAd *common.LegacyResponsiveDisplayAdInfo `protobuf:"bytes,28,opt,name=legacy_responsive_display_ad,json=legacyResponsiveDisplayAd,proto3,oneof"`
}

type Ad_AppAd struct {
	AppAd *common.AppAdInfo `protobuf:"bytes,29,opt,name=app_ad,json=appAd,proto3,oneof"`
}

type Ad_LegacyAppInstallAd struct {
	LegacyAppInstallAd *common.LegacyAppInstallAdInfo `protobuf:"bytes,30,opt,name=legacy_app_install_ad,json=legacyAppInstallAd,proto3,oneof"`
}

type Ad_ResponsiveDisplayAd struct {
	ResponsiveDisplayAd *common.ResponsiveDisplayAdInfo `protobuf:"bytes,31,opt,name=responsive_display_ad,json=responsiveDisplayAd,proto3,oneof"`
}

type Ad_DisplayUploadAd struct {
	DisplayUploadAd *common.DisplayUploadAdInfo `protobuf:"bytes,33,opt,name=display_upload_ad,json=displayUploadAd,proto3,oneof"`
}

type Ad_AppEngagementAd struct {
	AppEngagementAd *common.AppEngagementAdInfo `protobuf:"bytes,34,opt,name=app_engagement_ad,json=appEngagementAd,proto3,oneof"`
}

type Ad_ShoppingComparisonListingAd struct {
	ShoppingComparisonListingAd *common.ShoppingComparisonListingAdInfo `protobuf:"bytes,36,opt,name=shopping_comparison_listing_ad,json=shoppingComparisonListingAd,proto3,oneof"`
}

func (*Ad_TextAd) isAd_AdData() {}

func (*Ad_ExpandedTextAd) isAd_AdData() {}

func (*Ad_CallOnlyAd) isAd_AdData() {}

func (*Ad_ExpandedDynamicSearchAd) isAd_AdData() {}

func (*Ad_HotelAd) isAd_AdData() {}

func (*Ad_ShoppingSmartAd) isAd_AdData() {}

func (*Ad_ShoppingProductAd) isAd_AdData() {}

func (*Ad_GmailAd) isAd_AdData() {}

func (*Ad_ImageAd) isAd_AdData() {}

func (*Ad_VideoAd) isAd_AdData() {}

func (*Ad_ResponsiveSearchAd) isAd_AdData() {}

func (*Ad_LegacyResponsiveDisplayAd) isAd_AdData() {}

func (*Ad_AppAd) isAd_AdData() {}

func (*Ad_LegacyAppInstallAd) isAd_AdData() {}

func (*Ad_ResponsiveDisplayAd) isAd_AdData() {}

func (*Ad_DisplayUploadAd) isAd_AdData() {}

func (*Ad_AppEngagementAd) isAd_AdData() {}

func (*Ad_ShoppingComparisonListingAd) isAd_AdData() {}

func (m *Ad) GetAdData() isAd_AdData {
	if m != nil {
		return m.AdData
	}
	return nil
}

func (m *Ad) GetTextAd() *common.TextAdInfo {
	if x, ok := m.GetAdData().(*Ad_TextAd); ok {
		return x.TextAd
	}
	return nil
}

func (m *Ad) GetExpandedTextAd() *common.ExpandedTextAdInfo {
	if x, ok := m.GetAdData().(*Ad_ExpandedTextAd); ok {
		return x.ExpandedTextAd
	}
	return nil
}

func (m *Ad) GetCallOnlyAd() *common.CallOnlyAdInfo {
	if x, ok := m.GetAdData().(*Ad_CallOnlyAd); ok {
		return x.CallOnlyAd
	}
	return nil
}

func (m *Ad) GetExpandedDynamicSearchAd() *common.ExpandedDynamicSearchAdInfo {
	if x, ok := m.GetAdData().(*Ad_ExpandedDynamicSearchAd); ok {
		return x.ExpandedDynamicSearchAd
	}
	return nil
}

func (m *Ad) GetHotelAd() *common.HotelAdInfo {
	if x, ok := m.GetAdData().(*Ad_HotelAd); ok {
		return x.HotelAd
	}
	return nil
}

func (m *Ad) GetShoppingSmartAd() *common.ShoppingSmartAdInfo {
	if x, ok := m.GetAdData().(*Ad_ShoppingSmartAd); ok {
		return x.ShoppingSmartAd
	}
	return nil
}

func (m *Ad) GetShoppingProductAd() *common.ShoppingProductAdInfo {
	if x, ok := m.GetAdData().(*Ad_ShoppingProductAd); ok {
		return x.ShoppingProductAd
	}
	return nil
}

func (m *Ad) GetGmailAd() *common.GmailAdInfo {
	if x, ok := m.GetAdData().(*Ad_GmailAd); ok {
		return x.GmailAd
	}
	return nil
}

func (m *Ad) GetImageAd() *common.ImageAdInfo {
	if x, ok := m.GetAdData().(*Ad_ImageAd); ok {
		return x.ImageAd
	}
	return nil
}

func (m *Ad) GetVideoAd() *common.VideoAdInfo {
	if x, ok := m.GetAdData().(*Ad_VideoAd); ok {
		return x.VideoAd
	}
	return nil
}

func (m *Ad) GetResponsiveSearchAd() *common.ResponsiveSearchAdInfo {
	if x, ok := m.GetAdData().(*Ad_ResponsiveSearchAd); ok {
		return x.ResponsiveSearchAd
	}
	return nil
}

func (m *Ad) GetLegacyResponsiveDisplayAd() *common.LegacyResponsiveDisplayAdInfo {
	if x, ok := m.GetAdData().(*Ad_LegacyResponsiveDisplayAd); ok {
		return x.LegacyResponsiveDisplayAd
	}
	return nil
}

func (m *Ad) GetAppAd() *common.AppAdInfo {
	if x, ok := m.GetAdData().(*Ad_AppAd); ok {
		return x.AppAd
	}
	return nil
}

func (m *Ad) GetLegacyAppInstallAd() *common.LegacyAppInstallAdInfo {
	if x, ok := m.GetAdData().(*Ad_LegacyAppInstallAd); ok {
		return x.LegacyAppInstallAd
	}
	return nil
}

func (m *Ad) GetResponsiveDisplayAd() *common.ResponsiveDisplayAdInfo {
	if x, ok := m.GetAdData().(*Ad_ResponsiveDisplayAd); ok {
		return x.ResponsiveDisplayAd
	}
	return nil
}

func (m *Ad) GetDisplayUploadAd() *common.DisplayUploadAdInfo {
	if x, ok := m.GetAdData().(*Ad_DisplayUploadAd); ok {
		return x.DisplayUploadAd
	}
	return nil
}

func (m *Ad) GetAppEngagementAd() *common.AppEngagementAdInfo {
	if x, ok := m.GetAdData().(*Ad_AppEngagementAd); ok {
		return x.AppEngagementAd
	}
	return nil
}

func (m *Ad) GetShoppingComparisonListingAd() *common.ShoppingComparisonListingAdInfo {
	if x, ok := m.GetAdData().(*Ad_ShoppingComparisonListingAd); ok {
		return x.ShoppingComparisonListingAd
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Ad) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Ad_TextAd)(nil),
		(*Ad_ExpandedTextAd)(nil),
		(*Ad_CallOnlyAd)(nil),
		(*Ad_ExpandedDynamicSearchAd)(nil),
		(*Ad_HotelAd)(nil),
		(*Ad_ShoppingSmartAd)(nil),
		(*Ad_ShoppingProductAd)(nil),
		(*Ad_GmailAd)(nil),
		(*Ad_ImageAd)(nil),
		(*Ad_VideoAd)(nil),
		(*Ad_ResponsiveSearchAd)(nil),
		(*Ad_LegacyResponsiveDisplayAd)(nil),
		(*Ad_AppAd)(nil),
		(*Ad_LegacyAppInstallAd)(nil),
		(*Ad_ResponsiveDisplayAd)(nil),
		(*Ad_DisplayUploadAd)(nil),
		(*Ad_AppEngagementAd)(nil),
		(*Ad_ShoppingComparisonListingAd)(nil),
	}
}

func init() {
	proto.RegisterType((*Ad)(nil), "google.ads.googleads.v1.resources.Ad")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad.proto", fileDescriptor_2cfecfd7772c305a)
}

var fileDescriptor_2cfecfd7772c305a = []byte{
	// 1216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x97, 0xdd, 0x6e, 0xdb, 0x36,
	0x1b, 0xc7, 0x5f, 0xbb, 0x6d, 0xf2, 0x96, 0xed, 0xf2, 0xc1, 0x34, 0xab, 0x9a, 0x64, 0x5d, 0xdb,
	0xad, 0x40, 0xd7, 0x62, 0x72, 0x93, 0xae, 0x1d, 0xe0, 0xa2, 0xd8, 0xe4, 0x26, 0x6b, 0x32, 0xb4,
	0x98, 0x67, 0xd7, 0x3e, 0x28, 0xb2, 0x09, 0x8c, 0x48, 0x2b, 0x5a, 0x29, 0x92, 0x20, 0x25, 0xaf,
	0xde, 0xd1, 0x8e, 0x76, 0x1b, 0x03, 0x76, 0xb8, 0x4b, 0xd9, 0xa5, 0xec, 0x6c, 0xc0, 0x2e, 0x60,
	0x20, 0x29, 0xd1, 0x75, 0x12, 0x47, 0x3a, 0x31, 0x44, 0xf2, 0xf9, 0xff, 0xfe, 0x8f, 0x1e, 0x7e,
	0xc9, 0xe0, 0x7e, 0xcc, 0x79, 0x4c, 0x49, 0x0b, 0x61, 0xd5, 0xb2, 0x8f, 0xfa, 0x69, 0xbc, 0xdd,
	0x92, 0x44, 0xf1, 0x5c, 0x46, 0x44, 0xb5, 0x10, 0xf6, 0x85, 0xe4, 0x19, 0x87, 0xb7, 0x6d, 0x80,
	0x8f, 0xb0, 0xf2, 0x5d, 0xac, 0x3f, 0xde, 0xf6, 0x5d, 0xec, 0xc6, 0xce, 0x3c, 0x5c, 0xc4, 0xd3,
	0x94, 0xb3, 0x16, 0xc2, 0x61, 0x36, 0x11, 0x24, 0x4c, 0xd8, 0x88, 0x2b, 0x8b, 0xdd, 0x78, 0x5c,
	0xa1, 0x89, 0x72, 0x95, 0xf1, 0x34, 0x14, 0x48, 0xa2, 0x94, 0x64, 0x44, 0x16, 0xb2, 0x2a, 0xab,
	0x51, 0xc2, 0x10, 0x0d, 0x91, 0x10, 0x61, 0x2e, 0x69, 0xa1, 0x79, 0x54, 0xa1, 0xc9, 0x25, 0x0d,
	0x23, 0x4e, 0x29, 0x89, 0xb2, 0x84, 0xb3, 0x42, 0xf4, 0x60, 0x9e, 0x88, 0xb0, 0x3c, 0x55, 0xe5,
	0x2b, 0x15, 0xc1, 0xf7, 0xcf, 0x0f, 0xc6, 0x64, 0x9c, 0x44, 0x65, 0xec, 0xd7, 0xe7, 0xc7, 0xaa,
	0x89, 0xca, 0x48, 0x1a, 0xa6, 0x88, 0xa1, 0x98, 0xe0, 0x90, 0xb0, 0x2c, 0xc9, 0x26, 0xa1, 0xad,
	0x74, 0x41, 0xd8, 0x2a, 0x09, 0x22, 0x69, 0x21, 0xc6, 0x78, 0x86, 0x74, 0xde, 0x65, 0x61, 0x6f,
	0x16, 0xa3, 0xa6, 0x75, 0x94, 0x8f, 0x5a, 0x3f, 0x4b, 0x24, 0x04, 0x91, 0xc5, 0xf8, 0x9d, 0x7f,
	0xd6, 0x41, 0x33, 0xc0, 0xf0, 0x01, 0x68, 0x26, 0xd8, 0x6b, 0xdc, 0x6a, 0xdc, 0xbb, 0xb2, 0xb3,
	0x59, 0x4c, 0xac, 0x5f, 0x6a, 0xfc, 0x03, 0x96, 0x3d, 0xf9, 0x62, 0x88, 0x68, 0x4e, 0x7a, 0xcd,
	0x04, 0xc3, 0xa7, 0x00, 0xd8, 0xc2, 0xe6, 0x92, 0x2a, 0xaf, 0x79, 0xeb, 0xc2, 0xbd, 0x2b, 0x3b,
	0x5b, 0xa7, 0x44, 0xfd, 0x4c, 0x26, 0x2c, 0xb6, 0xaa, 0xcb, 0x26, 0x7e, 0x20, 0xa9, 0x82, 0xdf,
	0x83, 0xa5, 0x99, 0x59, 0x51, 0xde, 0x27, 0x06, 0xf0, 0xc0, 0x9f, 0xb7, 0xb2, 0xec, 0xbc, 0xf8,
	0xdf, 0x68, 0x55, 0x20, 0xc4, 0x40, 0xd2, 0xde, 0xd5, 0xd1, 0xb4, 0xa1, 0xe0, 0x3e, 0x58, 0xb5,
	0xc8, 0x94, 0x1f, 0x25, 0x94, 0x58, 0xea, 0x4a, 0x8d, 0xb4, 0x96, 0x8d, 0xec, 0x95, 0x51, 0x19,
	0x52, 0x17, 0xac, 0x67, 0x12, 0x45, 0x6f, 0x13, 0x16, 0x6b, 0x4a, 0x98, 0x91, 0x54, 0x50, 0x94,
	0x11, 0xef, 0xaa, 0xa9, 0xcc, 0xf9, 0xb4, 0xb5, 0x52, 0x3a, 0x90, 0xf4, 0x75, 0x21, 0x84, 0x11,
	0x58, 0x37, 0x0b, 0xea, 0xc4, 0xfa, 0x55, 0x1e, 0x30, 0xf9, 0xb5, 0xaa, 0xde, 0xfa, 0xb9, 0x11,
	0x76, 0x4b, 0x5d, 0x6f, 0x2d, 0x97, 0xf4, 0x44, 0x9f, 0x82, 0xcf, 0xc0, 0x15, 0x9c, 0x28, 0x41,
	0xd1, 0x44, 0x67, 0xed, 0x5d, 0xac, 0x91, 0x2c, 0x28, 0x04, 0x03, 0x49, 0xe1, 0x2e, 0xb8, 0xa8,
	0x57, 0xaf, 0x77, 0xe9, 0x56, 0xe3, 0xde, 0xd2, 0xce, 0xc3, 0xb9, 0x29, 0x99, 0x25, 0xe9, 0x07,
	0xf8, 0xf5, 0x44, 0x90, 0x3d, 0x96, 0xa7, 0xc5, 0x63, 0xcf, 0xa8, 0xe1, 0x01, 0x58, 0x43, 0x18,
	0x13, 0x1c, 0x1e, 0x4d, 0x42, 0x2b, 0x0b, 0x11, 0x56, 0xde, 0x9a, 0x49, 0x66, 0xe3, 0x54, 0x32,
	0x1d, 0xce, 0xa9, 0x4d, 0x65, 0xc5, 0xc8, 0x3a, 0x93, 0x17, 0x26, 0x22, 0xc0, 0x0a, 0xfe, 0x00,
	0x56, 0xed, 0x26, 0x09, 0x85, 0x24, 0x23, 0x22, 0x09, 0x8b, 0x88, 0x77, 0xad, 0x56, 0x76, 0xbb,
	0x46, 0x67, 0xb2, 0xb3, 0x8f, 0xbd, 0x15, 0x8b, 0xea, 0x3a, 0x12, 0x1c, 0x82, 0xe5, 0xd9, 0x4d,
	0xae, 0xbc, 0x0d, 0x33, 0x1b, 0x9f, 0x57, 0xcd, 0xc6, 0x40, 0xd2, 0xe7, 0x4e, 0xd5, 0x5b, 0xca,
	0xdf, 0x6f, 0x2a, 0xf8, 0x10, 0x5c, 0x64, 0x28, 0x25, 0xde, 0xf5, 0x1a, 0xf5, 0x37, 0x91, 0xf0,
	0xf7, 0x06, 0xb8, 0x79, 0x62, 0x8b, 0x97, 0xe7, 0x68, 0xb1, 0xc9, 0xbd, 0x4d, 0xf3, 0xda, 0x6f,
	0x2a, 0x5e, 0xbb, 0x6f, 0x20, 0xaf, 0x2c, 0xa3, 0x57, 0x20, 0xfa, 0xe6, 0xd7, 0xd4, 0xe2, 0x9c,
	0xf1, 0xde, 0xa6, 0x9a, 0x3f, 0x08, 0xf7, 0xc0, 0x62, 0x46, 0xde, 0x65, 0x21, 0xc2, 0xde, 0x82,
	0x79, 0xad, 0xfb, 0x55, 0x35, 0x7a, 0x4d, 0xde, 0x65, 0x01, 0x3e, 0x60, 0x23, 0xbe, 0xff, 0xbf,
	0xde, 0x42, 0x66, 0x5a, 0xf0, 0x47, 0xb0, 0x42, 0xde, 0x09, 0xc4, 0xf4, 0xfa, 0x28, 0x79, 0x8b,
	0x86, 0xb7, 0x53, 0xc5, 0xdb, 0x2b, 0x74, 0x33, 0xdc, 0x25, 0x32, 0xd3, 0x0b, 0x7b, 0xe0, 0x6a,
	0x84, 0x28, 0x0d, 0x39, 0xa3, 0x13, 0xcd, 0xfe, 0xc0, 0xb0, 0xfd, 0xca, 0xdd, 0x85, 0x28, 0xfd,
	0x8e, 0xd1, 0x89, 0xe3, 0x82, 0xc8, 0xf5, 0xc0, 0x5f, 0xc0, 0x86, 0xcb, 0x19, 0x4f, 0x18, 0x4a,
	0x93, 0x28, 0x54, 0x04, 0xc9, 0xe8, 0x58, 0x3b, 0x2c, 0x19, 0x87, 0xa7, 0x75, 0xb3, 0xdf, 0xb5,
	0x80, 0xbe, 0xd1, 0x3b, 0xbb, 0xeb, 0xe4, 0xec, 0x61, 0xb8, 0x0f, 0xfe, 0x7f, 0xcc, 0x33, 0x42,
	0xb5, 0xd3, 0xb2, 0x71, 0xaa, 0x3c, 0x1f, 0xf7, 0x75, 0xbc, 0x23, 0x2f, 0x1e, 0xdb, 0x26, 0x44,
	0x60, 0x55, 0x1d, 0x73, 0x21, 0xf4, 0x91, 0xa6, 0x52, 0x24, 0x4d, 0xe9, 0x57, 0x0d, 0xf2, 0x51,
	0x15, 0xb2, 0x5f, 0x08, 0xfb, 0x5a, 0xe7, 0xd0, 0xcb, 0x6a, 0xb6, 0x1b, 0xc6, 0x60, 0xcd, 0x59,
	0x08, 0xc9, 0x71, 0x1e, 0x19, 0x13, 0x68, 0x4c, 0x1e, 0xd7, 0x35, 0xe9, 0x5a, 0xa5, 0xb3, 0x71,
	0x69, 0xbb, 0x01, 0x5d, 0x95, 0x38, 0x45, 0x89, 0xa9, 0xca, 0x7a, 0xbd, 0xaa, 0xbc, 0xd0, 0xf1,
	0xd3, 0xaa, 0xc4, 0xb6, 0xa9, 0x49, 0x49, 0x8a, 0x62, 0x7d, 0x44, 0x79, 0x1f, 0xd6, 0x23, 0x1d,
	0xe8, 0xf8, 0x29, 0x29, 0xb1, 0x4d, 0x4d, 0x1a, 0x27, 0x98, 0x70, 0x4d, 0xf2, 0xea, 0x91, 0x86,
	0x3a, 0x7e, 0x4a, 0x1a, 0xdb, 0x26, 0xfc, 0x09, 0x5c, 0x93, 0x44, 0x09, 0xce, 0x54, 0x32, 0x26,
	0xef, 0xad, 0xb4, 0x1b, 0x86, 0xfa, 0xa4, 0x8a, 0xda, 0x73, 0xda, 0x13, 0x8b, 0x0c, 0xca, 0x53,
	0x23, 0xf0, 0xd7, 0x06, 0xd8, 0xa2, 0x24, 0x46, 0xd1, 0x24, 0x7c, 0xcf, 0xb3, 0xbc, 0x44, 0x10,
	0xf6, 0xb6, 0x8c, 0xe9, 0xb3, 0x2a, 0xd3, 0x97, 0x86, 0x31, 0xb5, 0xde, 0xb5, 0x04, 0xe7, 0x7d,
	0x83, 0xce, 0x0b, 0x80, 0x1d, 0xb0, 0xa0, 0x3f, 0x01, 0x10, 0xf6, 0x3e, 0x32, 0x5e, 0x9f, 0x55,
	0x79, 0x05, 0x42, 0x38, 0xee, 0x25, 0xa4, 0x1b, 0xf0, 0x2d, 0x58, 0x2f, 0xde, 0x42, 0xa3, 0x12,
	0xa6, 0x32, 0x7d, 0x08, 0x20, 0xec, 0xdd, 0xac, 0x57, 0x33, 0x9b, 0x7e, 0x20, 0xc4, 0x81, 0x95,
	0x4e, 0x6b, 0x46, 0x4f, 0x8d, 0xc0, 0x14, 0xac, 0x9f, 0x5d, 0xab, 0x8f, 0x8d, 0xd9, 0x97, 0xf5,
	0x27, 0xe8, 0x64, 0x95, 0xd6, 0xe4, 0x19, 0xf5, 0x41, 0x60, 0xd5, 0x5d, 0xea, 0x82, 0x72, 0x84,
	0xb5, 0xd5, 0xed, 0x7a, 0x1b, 0xb7, 0xa0, 0x0c, 0x8c, 0x6e, 0xba, 0x71, 0xf1, 0x6c, 0xb7, 0xb6,
	0xd0, 0x75, 0x23, 0x2c, 0x46, 0x31, 0x49, 0x09, 0x33, 0xdb, 0xf6, 0x4e, 0x3d, 0x8b, 0x40, 0x88,
	0x3d, 0xa7, 0x9b, 0x5a, 0xa0, 0xd9, 0x6e, 0xf8, 0x9b, 0xbe, 0xe1, 0xca, 0xc3, 0x21, 0xe2, 0xa9,
	0x40, 0x32, 0x51, 0x9c, 0x85, 0x34, 0x51, 0x99, 0xee, 0x42, 0xd8, 0xfb, 0xd4, 0x18, 0x7e, 0x55,
	0xf7, 0x9c, 0x78, 0xee, 0x20, 0x2f, 0x2d, 0xc3, 0x99, 0x6f, 0xaa, 0xf9, 0x21, 0x9d, 0xcb, 0x60,
	0x11, 0xe1, 0x10, 0xa3, 0x0c, 0x75, 0xfe, 0x6d, 0x80, 0xbb, 0x11, 0x4f, 0xfd, 0xca, 0xbf, 0x32,
	0x9d, 0xc5, 0x00, 0x77, 0xf5, 0xed, 0xdd, 0x6d, 0xbc, 0xf9, 0xb6, 0x88, 0x8e, 0x39, 0x45, 0x2c,
	0xf6, 0xb9, 0x8c, 0x5b, 0x31, 0x61, 0xe6, 0x6e, 0x2f, 0x3f, 0xdc, 0x45, 0xa2, 0xce, 0xf9, 0x0f,
	0xf5, 0xd4, 0x3d, 0xfd, 0xd1, 0xbc, 0xf0, 0x22, 0x08, 0xfe, 0x6c, 0xde, 0xb6, 0x5f, 0x3c, 0x7e,
	0x80, 0x95, 0xef, 0x3e, 0x7e, 0xfc, 0xe1, 0xb6, 0x5f, 0xde, 0xc1, 0xea, 0xaf, 0x32, 0xe6, 0x30,
	0xc0, 0xea, 0xd0, 0xc5, 0x1c, 0x0e, 0xb7, 0x0f, 0x5d, 0xcc, 0xdf, 0xcd, 0xbb, 0x76, 0xa0, 0xdd,
	0x0e, 0xb0, 0x6a, 0xb7, 0x5d, 0x54, 0xbb, 0x3d, 0xdc, 0x6e, 0xb7, 0x5d, 0xdc, 0xd1, 0x82, 0x49,
	0xf6, 0xd1, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x27, 0x6c, 0xba, 0x2d, 0xef, 0x0d, 0x00, 0x00,
}
