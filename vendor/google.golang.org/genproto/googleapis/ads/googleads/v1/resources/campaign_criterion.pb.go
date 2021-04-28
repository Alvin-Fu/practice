// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/campaign_criterion.proto

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

// A campaign criterion.
type CampaignCriterion struct {
	// The resource name of the campaign criterion.
	// Campaign criterion resource names have the form:
	//
	// `customers/{customer_id}/campaignCriteria/{campaign_id}~{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The campaign to which the criterion belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,4,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The ID of the criterion.
	//
	// This field is ignored during mutate.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,5,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The modifier for the bids when the criterion matches. The modifier must be
	// in the range: 0.1 - 10.0. Most targetable criteria types support modifiers.
	// Use 0 to opt out of a Device type.
	BidModifier *wrappers.FloatValue `protobuf:"bytes,14,opt,name=bid_modifier,json=bidModifier,proto3" json:"bid_modifier,omitempty"`
	// Whether to target (`false`) or exclude (`true`) the criterion.
	Negative *wrappers.BoolValue `protobuf:"bytes,7,opt,name=negative,proto3" json:"negative,omitempty"`
	// The type of the criterion.
	Type enums.CriterionTypeEnum_CriterionType `protobuf:"varint,6,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.CriterionTypeEnum_CriterionType" json:"type,omitempty"`
	// The status of the criterion.
	Status enums.CampaignCriterionStatusEnum_CampaignCriterionStatus `protobuf:"varint,35,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.CampaignCriterionStatusEnum_CampaignCriterionStatus" json:"status,omitempty"`
	// The campaign criterion.
	//
	// Exactly one must be set.
	//
	// Types that are valid to be assigned to Criterion:
	//	*CampaignCriterion_Keyword
	//	*CampaignCriterion_Placement
	//	*CampaignCriterion_MobileAppCategory
	//	*CampaignCriterion_MobileApplication
	//	*CampaignCriterion_Location
	//	*CampaignCriterion_Device
	//	*CampaignCriterion_AdSchedule
	//	*CampaignCriterion_AgeRange
	//	*CampaignCriterion_Gender
	//	*CampaignCriterion_IncomeRange
	//	*CampaignCriterion_ParentalStatus
	//	*CampaignCriterion_UserList
	//	*CampaignCriterion_YoutubeVideo
	//	*CampaignCriterion_YoutubeChannel
	//	*CampaignCriterion_Proximity
	//	*CampaignCriterion_Topic
	//	*CampaignCriterion_ListingScope
	//	*CampaignCriterion_Language
	//	*CampaignCriterion_IpBlock
	//	*CampaignCriterion_ContentLabel
	//	*CampaignCriterion_Carrier
	//	*CampaignCriterion_UserInterest
	//	*CampaignCriterion_Webpage
	//	*CampaignCriterion_OperatingSystemVersion
	//	*CampaignCriterion_MobileDevice
	//	*CampaignCriterion_LocationGroup
	Criterion            isCampaignCriterion_Criterion `protobuf_oneof:"criterion"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CampaignCriterion) Reset()         { *m = CampaignCriterion{} }
func (m *CampaignCriterion) String() string { return proto.CompactTextString(m) }
func (*CampaignCriterion) ProtoMessage()    {}
func (*CampaignCriterion) Descriptor() ([]byte, []int) {
	return fileDescriptor_66974bcead75dadf, []int{0}
}

func (m *CampaignCriterion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignCriterion.Unmarshal(m, b)
}
func (m *CampaignCriterion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignCriterion.Marshal(b, m, deterministic)
}
func (m *CampaignCriterion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignCriterion.Merge(m, src)
}
func (m *CampaignCriterion) XXX_Size() int {
	return xxx_messageInfo_CampaignCriterion.Size(m)
}
func (m *CampaignCriterion) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignCriterion.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignCriterion proto.InternalMessageInfo

func (m *CampaignCriterion) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignCriterion) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignCriterion) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *CampaignCriterion) GetBidModifier() *wrappers.FloatValue {
	if m != nil {
		return m.BidModifier
	}
	return nil
}

func (m *CampaignCriterion) GetNegative() *wrappers.BoolValue {
	if m != nil {
		return m.Negative
	}
	return nil
}

func (m *CampaignCriterion) GetType() enums.CriterionTypeEnum_CriterionType {
	if m != nil {
		return m.Type
	}
	return enums.CriterionTypeEnum_UNSPECIFIED
}

func (m *CampaignCriterion) GetStatus() enums.CampaignCriterionStatusEnum_CampaignCriterionStatus {
	if m != nil {
		return m.Status
	}
	return enums.CampaignCriterionStatusEnum_UNSPECIFIED
}

type isCampaignCriterion_Criterion interface {
	isCampaignCriterion_Criterion()
}

type CampaignCriterion_Keyword struct {
	Keyword *common.KeywordInfo `protobuf:"bytes,8,opt,name=keyword,proto3,oneof"`
}

type CampaignCriterion_Placement struct {
	Placement *common.PlacementInfo `protobuf:"bytes,9,opt,name=placement,proto3,oneof"`
}

type CampaignCriterion_MobileAppCategory struct {
	MobileAppCategory *common.MobileAppCategoryInfo `protobuf:"bytes,10,opt,name=mobile_app_category,json=mobileAppCategory,proto3,oneof"`
}

type CampaignCriterion_MobileApplication struct {
	MobileApplication *common.MobileApplicationInfo `protobuf:"bytes,11,opt,name=mobile_application,json=mobileApplication,proto3,oneof"`
}

type CampaignCriterion_Location struct {
	Location *common.LocationInfo `protobuf:"bytes,12,opt,name=location,proto3,oneof"`
}

type CampaignCriterion_Device struct {
	Device *common.DeviceInfo `protobuf:"bytes,13,opt,name=device,proto3,oneof"`
}

type CampaignCriterion_AdSchedule struct {
	AdSchedule *common.AdScheduleInfo `protobuf:"bytes,15,opt,name=ad_schedule,json=adSchedule,proto3,oneof"`
}

type CampaignCriterion_AgeRange struct {
	AgeRange *common.AgeRangeInfo `protobuf:"bytes,16,opt,name=age_range,json=ageRange,proto3,oneof"`
}

type CampaignCriterion_Gender struct {
	Gender *common.GenderInfo `protobuf:"bytes,17,opt,name=gender,proto3,oneof"`
}

type CampaignCriterion_IncomeRange struct {
	IncomeRange *common.IncomeRangeInfo `protobuf:"bytes,18,opt,name=income_range,json=incomeRange,proto3,oneof"`
}

type CampaignCriterion_ParentalStatus struct {
	ParentalStatus *common.ParentalStatusInfo `protobuf:"bytes,19,opt,name=parental_status,json=parentalStatus,proto3,oneof"`
}

type CampaignCriterion_UserList struct {
	UserList *common.UserListInfo `protobuf:"bytes,22,opt,name=user_list,json=userList,proto3,oneof"`
}

type CampaignCriterion_YoutubeVideo struct {
	YoutubeVideo *common.YouTubeVideoInfo `protobuf:"bytes,20,opt,name=youtube_video,json=youtubeVideo,proto3,oneof"`
}

type CampaignCriterion_YoutubeChannel struct {
	YoutubeChannel *common.YouTubeChannelInfo `protobuf:"bytes,21,opt,name=youtube_channel,json=youtubeChannel,proto3,oneof"`
}

type CampaignCriterion_Proximity struct {
	Proximity *common.ProximityInfo `protobuf:"bytes,23,opt,name=proximity,proto3,oneof"`
}

type CampaignCriterion_Topic struct {
	Topic *common.TopicInfo `protobuf:"bytes,24,opt,name=topic,proto3,oneof"`
}

type CampaignCriterion_ListingScope struct {
	ListingScope *common.ListingScopeInfo `protobuf:"bytes,25,opt,name=listing_scope,json=listingScope,proto3,oneof"`
}

type CampaignCriterion_Language struct {
	Language *common.LanguageInfo `protobuf:"bytes,26,opt,name=language,proto3,oneof"`
}

type CampaignCriterion_IpBlock struct {
	IpBlock *common.IpBlockInfo `protobuf:"bytes,27,opt,name=ip_block,json=ipBlock,proto3,oneof"`
}

type CampaignCriterion_ContentLabel struct {
	ContentLabel *common.ContentLabelInfo `protobuf:"bytes,28,opt,name=content_label,json=contentLabel,proto3,oneof"`
}

type CampaignCriterion_Carrier struct {
	Carrier *common.CarrierInfo `protobuf:"bytes,29,opt,name=carrier,proto3,oneof"`
}

type CampaignCriterion_UserInterest struct {
	UserInterest *common.UserInterestInfo `protobuf:"bytes,30,opt,name=user_interest,json=userInterest,proto3,oneof"`
}

type CampaignCriterion_Webpage struct {
	Webpage *common.WebpageInfo `protobuf:"bytes,31,opt,name=webpage,proto3,oneof"`
}

type CampaignCriterion_OperatingSystemVersion struct {
	OperatingSystemVersion *common.OperatingSystemVersionInfo `protobuf:"bytes,32,opt,name=operating_system_version,json=operatingSystemVersion,proto3,oneof"`
}

type CampaignCriterion_MobileDevice struct {
	MobileDevice *common.MobileDeviceInfo `protobuf:"bytes,33,opt,name=mobile_device,json=mobileDevice,proto3,oneof"`
}

type CampaignCriterion_LocationGroup struct {
	LocationGroup *common.LocationGroupInfo `protobuf:"bytes,34,opt,name=location_group,json=locationGroup,proto3,oneof"`
}

func (*CampaignCriterion_Keyword) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Placement) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_MobileAppCategory) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_MobileApplication) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Location) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Device) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_AdSchedule) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_AgeRange) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Gender) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_IncomeRange) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_ParentalStatus) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_UserList) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_YoutubeVideo) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_YoutubeChannel) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Proximity) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Topic) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_ListingScope) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Language) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_IpBlock) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_ContentLabel) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Carrier) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_UserInterest) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Webpage) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_OperatingSystemVersion) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_MobileDevice) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_LocationGroup) isCampaignCriterion_Criterion() {}

func (m *CampaignCriterion) GetCriterion() isCampaignCriterion_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (m *CampaignCriterion) GetKeyword() *common.KeywordInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Keyword); ok {
		return x.Keyword
	}
	return nil
}

func (m *CampaignCriterion) GetPlacement() *common.PlacementInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Placement); ok {
		return x.Placement
	}
	return nil
}

func (m *CampaignCriterion) GetMobileAppCategory() *common.MobileAppCategoryInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_MobileAppCategory); ok {
		return x.MobileAppCategory
	}
	return nil
}

func (m *CampaignCriterion) GetMobileApplication() *common.MobileApplicationInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_MobileApplication); ok {
		return x.MobileApplication
	}
	return nil
}

func (m *CampaignCriterion) GetLocation() *common.LocationInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Location); ok {
		return x.Location
	}
	return nil
}

func (m *CampaignCriterion) GetDevice() *common.DeviceInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Device); ok {
		return x.Device
	}
	return nil
}

func (m *CampaignCriterion) GetAdSchedule() *common.AdScheduleInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_AdSchedule); ok {
		return x.AdSchedule
	}
	return nil
}

func (m *CampaignCriterion) GetAgeRange() *common.AgeRangeInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_AgeRange); ok {
		return x.AgeRange
	}
	return nil
}

func (m *CampaignCriterion) GetGender() *common.GenderInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Gender); ok {
		return x.Gender
	}
	return nil
}

func (m *CampaignCriterion) GetIncomeRange() *common.IncomeRangeInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_IncomeRange); ok {
		return x.IncomeRange
	}
	return nil
}

func (m *CampaignCriterion) GetParentalStatus() *common.ParentalStatusInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_ParentalStatus); ok {
		return x.ParentalStatus
	}
	return nil
}

func (m *CampaignCriterion) GetUserList() *common.UserListInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_UserList); ok {
		return x.UserList
	}
	return nil
}

func (m *CampaignCriterion) GetYoutubeVideo() *common.YouTubeVideoInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_YoutubeVideo); ok {
		return x.YoutubeVideo
	}
	return nil
}

func (m *CampaignCriterion) GetYoutubeChannel() *common.YouTubeChannelInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_YoutubeChannel); ok {
		return x.YoutubeChannel
	}
	return nil
}

func (m *CampaignCriterion) GetProximity() *common.ProximityInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Proximity); ok {
		return x.Proximity
	}
	return nil
}

func (m *CampaignCriterion) GetTopic() *common.TopicInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Topic); ok {
		return x.Topic
	}
	return nil
}

func (m *CampaignCriterion) GetListingScope() *common.ListingScopeInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_ListingScope); ok {
		return x.ListingScope
	}
	return nil
}

func (m *CampaignCriterion) GetLanguage() *common.LanguageInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Language); ok {
		return x.Language
	}
	return nil
}

func (m *CampaignCriterion) GetIpBlock() *common.IpBlockInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_IpBlock); ok {
		return x.IpBlock
	}
	return nil
}

func (m *CampaignCriterion) GetContentLabel() *common.ContentLabelInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_ContentLabel); ok {
		return x.ContentLabel
	}
	return nil
}

func (m *CampaignCriterion) GetCarrier() *common.CarrierInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Carrier); ok {
		return x.Carrier
	}
	return nil
}

func (m *CampaignCriterion) GetUserInterest() *common.UserInterestInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_UserInterest); ok {
		return x.UserInterest
	}
	return nil
}

func (m *CampaignCriterion) GetWebpage() *common.WebpageInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Webpage); ok {
		return x.Webpage
	}
	return nil
}

func (m *CampaignCriterion) GetOperatingSystemVersion() *common.OperatingSystemVersionInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_OperatingSystemVersion); ok {
		return x.OperatingSystemVersion
	}
	return nil
}

func (m *CampaignCriterion) GetMobileDevice() *common.MobileDeviceInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_MobileDevice); ok {
		return x.MobileDevice
	}
	return nil
}

func (m *CampaignCriterion) GetLocationGroup() *common.LocationGroupInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_LocationGroup); ok {
		return x.LocationGroup
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignCriterion) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignCriterion_Keyword)(nil),
		(*CampaignCriterion_Placement)(nil),
		(*CampaignCriterion_MobileAppCategory)(nil),
		(*CampaignCriterion_MobileApplication)(nil),
		(*CampaignCriterion_Location)(nil),
		(*CampaignCriterion_Device)(nil),
		(*CampaignCriterion_AdSchedule)(nil),
		(*CampaignCriterion_AgeRange)(nil),
		(*CampaignCriterion_Gender)(nil),
		(*CampaignCriterion_IncomeRange)(nil),
		(*CampaignCriterion_ParentalStatus)(nil),
		(*CampaignCriterion_UserList)(nil),
		(*CampaignCriterion_YoutubeVideo)(nil),
		(*CampaignCriterion_YoutubeChannel)(nil),
		(*CampaignCriterion_Proximity)(nil),
		(*CampaignCriterion_Topic)(nil),
		(*CampaignCriterion_ListingScope)(nil),
		(*CampaignCriterion_Language)(nil),
		(*CampaignCriterion_IpBlock)(nil),
		(*CampaignCriterion_ContentLabel)(nil),
		(*CampaignCriterion_Carrier)(nil),
		(*CampaignCriterion_UserInterest)(nil),
		(*CampaignCriterion_Webpage)(nil),
		(*CampaignCriterion_OperatingSystemVersion)(nil),
		(*CampaignCriterion_MobileDevice)(nil),
		(*CampaignCriterion_LocationGroup)(nil),
	}
}

func init() {
	proto.RegisterType((*CampaignCriterion)(nil), "google.ads.googleads.v1.resources.CampaignCriterion")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/campaign_criterion.proto", fileDescriptor_66974bcead75dadf)
}

var fileDescriptor_66974bcead75dadf = []byte{
	// 1113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x97, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0x80, 0x49, 0x68, 0x53, 0x67, 0xed, 0xa4, 0x64, 0x4b, 0xc3, 0x92, 0x86, 0x92, 0xb4, 0xd3,
	0x99, 0xf0, 0x53, 0x19, 0x07, 0xe8, 0x30, 0x66, 0xe8, 0x8c, 0xe3, 0x42, 0xea, 0x36, 0x81, 0xe0,
	0x04, 0x77, 0xe8, 0x84, 0xd1, 0xac, 0xa5, 0x13, 0x75, 0xa9, 0xb4, 0xab, 0x59, 0x49, 0x0e, 0xbe,
	0x86, 0x27, 0xe1, 0x92, 0x47, 0xe1, 0x51, 0x78, 0x0a, 0x66, 0xff, 0x14, 0xa7, 0x49, 0x90, 0xb8,
	0x93, 0xce, 0x9e, 0xef, 0xdb, 0xb3, 0xc7, 0xd2, 0xae, 0x8c, 0xba, 0x91, 0x10, 0x51, 0x0c, 0x6d,
	0x1a, 0x66, 0x6d, 0x73, 0xa9, 0xae, 0x26, 0x9d, 0xb6, 0x84, 0x4c, 0x14, 0x32, 0x80, 0xac, 0x1d,
	0xd0, 0x24, 0xa5, 0x2c, 0xe2, 0x7e, 0x20, 0x59, 0x0e, 0x92, 0x09, 0xee, 0xa5, 0x52, 0xe4, 0x02,
	0x6f, 0x1a, 0xc0, 0xa3, 0x61, 0xe6, 0x95, 0xac, 0x37, 0xe9, 0x78, 0x25, 0xbb, 0xf6, 0xf0, 0x2a,
	0x7d, 0x20, 0x92, 0x44, 0xf0, 0xb6, 0x55, 0x52, 0x63, 0x5c, 0xfb, 0xe6, 0xaa, 0x74, 0xe0, 0x45,
	0x72, 0x59, 0x25, 0x7e, 0x96, 0xd3, 0xbc, 0xc8, 0x2c, 0xbe, 0x5d, 0x81, 0x97, 0x54, 0x3e, 0x4d,
	0xc1, 0x32, 0xeb, 0x8e, 0x49, 0x59, 0x9b, 0x72, 0x2e, 0x72, 0x9a, 0x33, 0xc1, 0x9d, 0xf1, 0xae,
	0x1d, 0xd5, 0x77, 0xe3, 0xe2, 0xa4, 0x7d, 0x2a, 0x69, 0x9a, 0x82, 0xb4, 0xe3, 0xf7, 0xfe, 0xb8,
	0x8d, 0x56, 0xfa, 0xb6, 0xaa, 0xbe, 0xd3, 0xe3, 0xfb, 0x68, 0xc9, 0xb5, 0xc0, 0xe7, 0x34, 0x01,
	0x32, 0xb7, 0x31, 0xb7, 0xb5, 0x38, 0x6c, 0xb9, 0xe0, 0xf7, 0x34, 0x01, 0xfc, 0x15, 0x6a, 0xb8,
	0xf5, 0x90, 0x6b, 0x1b, 0x73, 0x5b, 0xcd, 0xed, 0x75, 0xdb, 0x45, 0xcf, 0xcd, 0xe6, 0x1d, 0xe6,
	0x92, 0xf1, 0x68, 0x44, 0xe3, 0x02, 0x86, 0x65, 0x36, 0x7e, 0x8c, 0x5a, 0x67, 0x4b, 0x61, 0x21,
	0xb9, 0xae, 0xe9, 0x3b, 0x17, 0xe8, 0x01, 0xcf, 0x1f, 0x7d, 0x61, 0xe0, 0x66, 0x09, 0x0c, 0x42,
	0xc5, 0x8f, 0x59, 0xe8, 0x27, 0x22, 0x64, 0x27, 0x0c, 0x24, 0x59, 0xbe, 0x82, 0xff, 0x2e, 0x16,
	0x34, 0xb7, 0xfc, 0x98, 0x85, 0xfb, 0x36, 0x1f, 0x3f, 0x42, 0x0d, 0x0e, 0x11, 0xcd, 0xd9, 0x04,
	0xc8, 0x0d, 0xcd, 0xae, 0x5d, 0x60, 0x77, 0x84, 0x88, 0x6d, 0xdd, 0x2e, 0x17, 0x0f, 0xd1, 0x35,
	0xd5, 0x78, 0xb2, 0xb0, 0x31, 0xb7, 0xb5, 0xbc, 0xfd, 0xd8, 0xbb, 0xea, 0xf1, 0xd1, 0xbf, 0x96,
	0x57, 0xb6, 0xf3, 0x68, 0x9a, 0xc2, 0xb7, 0xbc, 0x48, 0xce, 0x47, 0x86, 0xda, 0x85, 0x7f, 0x45,
	0x0b, 0xe6, 0x11, 0x20, 0xf7, 0xb5, 0x75, 0x58, 0x65, 0x7d, 0xf3, 0xc7, 0x3a, 0xd4, 0xb4, 0xf1,
	0x5f, 0x3e, 0x36, 0xb4, 0x33, 0xe0, 0x5d, 0x74, 0xe3, 0x35, 0x4c, 0x4f, 0x85, 0x0c, 0x49, 0x43,
	0x2f, 0xfb, 0x93, 0x2b, 0x27, 0x33, 0x8f, 0xb7, 0xf7, 0xdc, 0xa4, 0x0f, 0xf8, 0x89, 0x78, 0xfa,
	0xd6, 0xd0, 0xd1, 0x78, 0x1f, 0x2d, 0xa6, 0x31, 0x0d, 0x20, 0x01, 0x9e, 0x93, 0x45, 0xad, 0x7a,
	0x58, 0xa5, 0x3a, 0x70, 0x80, 0x95, 0x9d, 0x19, 0x70, 0x84, 0x6e, 0x25, 0x62, 0xcc, 0x62, 0xf0,
	0x69, 0x9a, 0xfa, 0x01, 0xcd, 0x21, 0x12, 0x72, 0x4a, 0x90, 0x16, 0x7f, 0x59, 0x25, 0xde, 0xd7,
	0x68, 0x2f, 0x4d, 0xfb, 0x16, 0xb4, 0x13, 0xac, 0x24, 0x6f, 0x0e, 0xe0, 0x13, 0x84, 0xcf, 0x26,
	0x8a, 0x59, 0xa0, 0x5f, 0x15, 0xd2, 0xfc, 0x9f, 0xf3, 0x38, 0xf0, 0xc2, 0x3c, 0x6e, 0x00, 0x3f,
	0x43, 0x8d, 0x58, 0x58, 0x7b, 0x4b, 0xdb, 0x3f, 0xad, 0xb2, 0xef, 0x89, 0x73, 0xd2, 0x92, 0xc7,
	0x4f, 0xd0, 0x42, 0x08, 0x13, 0x16, 0x00, 0x59, 0xd2, 0xa6, 0x8f, 0xab, 0x4c, 0x4f, 0x74, 0xb6,
	0xf5, 0x58, 0x16, 0xff, 0x88, 0x9a, 0x34, 0xf4, 0xb3, 0xe0, 0x15, 0x84, 0x45, 0x0c, 0xe4, 0xa6,
	0x56, 0x79, 0x55, 0xaa, 0x5e, 0x78, 0x68, 0x09, 0xab, 0x43, 0xb4, 0x8c, 0xe0, 0xe7, 0x68, 0x91,
	0x46, 0xe0, 0x4b, 0xca, 0x23, 0x20, 0xef, 0xd4, 0x5b, 0x65, 0x2f, 0x82, 0xa1, 0xca, 0x77, 0xab,
	0xa4, 0xf6, 0x5e, 0xad, 0x32, 0x02, 0x1e, 0x82, 0x24, 0x2b, 0xf5, 0x56, 0xb9, 0xab, 0xb3, 0xdd,
	0x2a, 0x0d, 0x8b, 0x8f, 0x50, 0x8b, 0xf1, 0x40, 0x24, 0xae, 0x2a, 0xac, 0x5d, 0xed, 0x2a, 0xd7,
	0x40, 0x33, 0xb3, 0x85, 0x35, 0xd9, 0x59, 0x08, 0xff, 0x82, 0x6e, 0xa6, 0x54, 0x02, 0xcf, 0x69,
	0x6c, 0xb7, 0x6b, 0x72, 0x4b, 0x8b, 0xb7, 0x2b, 0x9f, 0x79, 0x8b, 0x99, 0xf7, 0xd0, 0xba, 0x97,
	0xd3, 0x73, 0x51, 0xd5, 0xc7, 0x22, 0x03, 0xe9, 0xc7, 0x2c, 0xcb, 0xc9, 0x6a, 0xbd, 0x3e, 0xfe,
	0x94, 0x81, 0xdc, 0x63, 0x99, 0x7b, 0x97, 0x1a, 0x85, 0xbd, 0xc7, 0x2f, 0xd0, 0xd2, 0x54, 0x14,
	0x79, 0x31, 0x06, 0x7f, 0xc2, 0x42, 0x10, 0xe4, 0x5d, 0x2d, 0xfc, 0xac, 0x4a, 0xf8, 0xb3, 0x28,
	0x8e, 0x8a, 0x31, 0x8c, 0x14, 0x63, 0xa5, 0x2d, 0x2b, 0xd2, 0x31, 0xd5, 0x04, 0x27, 0x0e, 0x5e,
	0x51, 0xce, 0x21, 0x26, 0xb7, 0xeb, 0x35, 0xc1, 0xaa, 0xfb, 0x86, 0x72, 0x4d, 0xb0, 0x32, 0x1b,
	0xd5, 0x3b, 0x8a, 0x14, 0xbf, 0xb1, 0x84, 0xe5, 0x53, 0xf2, 0x5e, 0xcd, 0x1d, 0xc5, 0x01, 0xe5,
	0x8e, 0xe2, 0x02, 0xb8, 0x87, 0xae, 0xe7, 0x22, 0x65, 0x01, 0x21, 0x5a, 0xf5, 0x51, 0x95, 0xea,
	0x48, 0x25, 0x5b, 0x8d, 0x21, 0x55, 0x27, 0xd5, 0x2f, 0xc2, 0x78, 0xe4, 0x67, 0x81, 0x48, 0x81,
	0xbc, 0x5f, 0xaf, 0x93, 0x7b, 0x06, 0x3a, 0x54, 0x8c, 0xeb, 0x64, 0x3c, 0x13, 0xd3, 0x9b, 0x03,
	0xe5, 0x51, 0x41, 0x23, 0x20, 0x6b, 0x35, 0x37, 0x07, 0x9b, 0x5f, 0x6e, 0x0e, 0xf6, 0x1e, 0x3f,
	0x45, 0x0d, 0x96, 0xfa, 0xe3, 0x58, 0x04, 0xaf, 0xc9, 0x9d, 0x7a, 0x5b, 0xfa, 0x20, 0xdd, 0x51,
	0xe9, 0x6e, 0x4b, 0x67, 0xe6, 0x56, 0x2d, 0x37, 0x10, 0x3c, 0x07, 0x9e, 0xfb, 0x31, 0x1d, 0x43,
	0x4c, 0xd6, 0xeb, 0x2d, 0xb7, 0x6f, 0xa0, 0x3d, 0xc5, 0xb8, 0xe5, 0x06, 0x33, 0x31, 0x75, 0xe8,
	0x04, 0x54, 0x4a, 0x75, 0x4e, 0x7f, 0x50, 0xaf, 0xc2, 0xbe, 0x49, 0x77, 0x15, 0x5a, 0x5a, 0x55,
	0xa8, 0xdf, 0x13, 0xc6, 0x73, 0x90, 0x90, 0xe5, 0xe4, 0x6e, 0xbd, 0x0a, 0xd5, 0xbb, 0x32, 0xb0,
	0x8c, 0xab, 0xb0, 0x98, 0x89, 0xa9, 0x0a, 0x4f, 0x61, 0x9c, 0xaa, 0xdf, 0xe3, 0xc3, 0x7a, 0x15,
	0xbe, 0x30, 0xe9, 0xae, 0x42, 0x4b, 0xe3, 0x09, 0x22, 0x22, 0x05, 0x49, 0xcd, 0x43, 0x33, 0xcd,
	0x72, 0x48, 0xfc, 0x09, 0xc8, 0x4c, 0x1d, 0x03, 0x1b, 0xda, 0xdc, 0xad, 0x32, 0xff, 0xe0, 0xf8,
	0x43, 0x8d, 0x8f, 0x0c, 0x6d, 0x27, 0x5a, 0x15, 0x97, 0x8e, 0xaa, 0xce, 0xd8, 0x63, 0xcd, 0x9e,
	0x14, 0x9b, 0xf5, 0x3a, 0x63, 0x4e, 0xb4, 0x73, 0xe7, 0x45, 0x2b, 0x99, 0x89, 0xe1, 0x97, 0x68,
	0xd9, 0x9d, 0x43, 0x7e, 0x24, 0x45, 0x91, 0x92, 0x7b, 0xda, 0xdc, 0xa9, 0x7b, 0x9a, 0xed, 0x2a,
	0xc8, 0xaa, 0x97, 0xe2, 0xd9, 0xe0, 0x4e, 0x13, 0x2d, 0x96, 0xdf, 0x74, 0x3b, 0xbf, 0xcf, 0xa3,
	0x07, 0x81, 0x48, 0xbc, 0xca, 0x0f, 0xf2, 0x9d, 0xd5, 0x0b, 0x1f, 0x39, 0x07, 0xea, 0x93, 0xed,
	0x60, 0xee, 0xe5, 0x33, 0x0b, 0x47, 0x42, 0xbd, 0x1e, 0x9e, 0x90, 0x51, 0x3b, 0x02, 0xae, 0x3f,
	0xe8, 0xdc, 0xc7, 0x74, 0xca, 0xb2, 0xff, 0xf8, 0xa3, 0xf0, 0x75, 0x79, 0xf5, 0xe7, 0xfc, 0xdb,
	0xbb, 0xbd, 0xde, 0x5f, 0xf3, 0x9b, 0xbb, 0x46, 0xd9, 0x0b, 0x33, 0xcf, 0x5c, 0xaa, 0xab, 0x51,
	0xc7, 0x1b, 0xba, 0xcc, 0xbf, 0x5d, 0xce, 0x71, 0x2f, 0xcc, 0x8e, 0xcb, 0x9c, 0xe3, 0x51, 0xe7,
	0xb8, 0xcc, 0xf9, 0x67, 0xfe, 0x81, 0x19, 0xe8, 0x76, 0x7b, 0x61, 0xd6, 0xed, 0x96, 0x59, 0xdd,
	0xee, 0xa8, 0xd3, 0xed, 0x96, 0x79, 0xe3, 0x05, 0x5d, 0xec, 0xe7, 0xff, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x89, 0x75, 0x44, 0x6c, 0xd4, 0x0c, 0x00, 0x00,
}
