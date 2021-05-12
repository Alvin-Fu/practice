// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/campaign_extension_setting.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A campaign extension setting.
type CampaignExtensionSetting struct {
	// The resource name of the campaign extension setting.
	// CampaignExtensionSetting resource names have the form:
	//
	//
	// `customers/{customer_id}/campaignExtensionSettings/{campaign_id}~{extension_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The extension type of the customer extension setting.
	ExtensionType enums.ExtensionTypeEnum_ExtensionType `protobuf:"varint,2,opt,name=extension_type,json=extensionType,proto3,enum=google.ads.googleads.v1.enums.ExtensionTypeEnum_ExtensionType" json:"extension_type,omitempty"`
	// The resource name of the campaign. The linked extension feed items will
	// serve under this campaign.
	// Campaign resource names have the form:
	//
	// `customers/{customer_id}/campaigns/{campaign_id}`
	Campaign *wrappers.StringValue `protobuf:"bytes,3,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The resource names of the extension feed items to serve under the campaign.
	// ExtensionFeedItem resource names have the form:
	//
	// `customers/{customer_id}/extensionFeedItems/{feed_item_id}`
	ExtensionFeedItems []*wrappers.StringValue `protobuf:"bytes,4,rep,name=extension_feed_items,json=extensionFeedItems,proto3" json:"extension_feed_items,omitempty"`
	// The device for which the extensions will serve. Optional.
	Device               enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice `protobuf:"varint,5,opt,name=device,proto3,enum=google.ads.googleads.v1.enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                `json:"-"`
	XXX_unrecognized     []byte                                                  `json:"-"`
	XXX_sizecache        int32                                                   `json:"-"`
}

func (m *CampaignExtensionSetting) Reset()         { *m = CampaignExtensionSetting{} }
func (m *CampaignExtensionSetting) String() string { return proto.CompactTextString(m) }
func (*CampaignExtensionSetting) ProtoMessage()    {}
func (*CampaignExtensionSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_be2f6987185b7cdc, []int{0}
}

func (m *CampaignExtensionSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignExtensionSetting.Unmarshal(m, b)
}
func (m *CampaignExtensionSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignExtensionSetting.Marshal(b, m, deterministic)
}
func (m *CampaignExtensionSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignExtensionSetting.Merge(m, src)
}
func (m *CampaignExtensionSetting) XXX_Size() int {
	return xxx_messageInfo_CampaignExtensionSetting.Size(m)
}
func (m *CampaignExtensionSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignExtensionSetting.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignExtensionSetting proto.InternalMessageInfo

func (m *CampaignExtensionSetting) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignExtensionSetting) GetExtensionType() enums.ExtensionTypeEnum_ExtensionType {
	if m != nil {
		return m.ExtensionType
	}
	return enums.ExtensionTypeEnum_UNSPECIFIED
}

func (m *CampaignExtensionSetting) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignExtensionSetting) GetExtensionFeedItems() []*wrappers.StringValue {
	if m != nil {
		return m.ExtensionFeedItems
	}
	return nil
}

func (m *CampaignExtensionSetting) GetDevice() enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice {
	if m != nil {
		return m.Device
	}
	return enums.ExtensionSettingDeviceEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CampaignExtensionSetting)(nil), "google.ads.googleads.v1.resources.CampaignExtensionSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/campaign_extension_setting.proto", fileDescriptor_be2f6987185b7cdc)
}

var fileDescriptor_be2f6987185b7cdc = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x6e, 0xd3, 0x30,
	0x18, 0x57, 0x52, 0x98, 0x20, 0xb0, 0x1d, 0x22, 0x0e, 0xd1, 0x34, 0x50, 0x07, 0x9a, 0xd4, 0x93,
	0xad, 0x94, 0x0b, 0x32, 0x08, 0x29, 0x85, 0x31, 0xc1, 0x61, 0xaa, 0x32, 0xd4, 0x03, 0xaa, 0x14,
	0x79, 0xf5, 0x37, 0x63, 0xa9, 0xb1, 0xad, 0xd8, 0x29, 0xec, 0x15, 0x78, 0x04, 0x8e, 0x1c, 0x79,
	0x14, 0x1e, 0x85, 0xa7, 0x40, 0x8d, 0x63, 0x43, 0x85, 0xba, 0xf5, 0xf6, 0xc5, 0xbf, 0x3f, 0xdf,
	0xf7, 0xf3, 0xe7, 0x24, 0x13, 0xae, 0x14, 0x5f, 0x02, 0xa6, 0xcc, 0x60, 0x57, 0xae, 0xab, 0x55,
	0x8e, 0x1b, 0x30, 0xaa, 0x6d, 0x16, 0x60, 0xf0, 0x82, 0xd6, 0x9a, 0x0a, 0x2e, 0x2b, 0xf8, 0x6a,
	0x41, 0x1a, 0xa1, 0x64, 0x65, 0xc0, 0x5a, 0x21, 0x39, 0xd2, 0x8d, 0xb2, 0x2a, 0x3d, 0x76, 0x42,
	0x44, 0x99, 0x41, 0xc1, 0x03, 0xad, 0x72, 0x14, 0x3c, 0x0e, 0x5f, 0x6d, 0x6b, 0x03, 0xb2, 0xad,
	0x0d, 0xfe, 0xcf, 0xb9, 0x62, 0xb0, 0x12, 0x0b, 0x70, 0x0d, 0x0e, 0xc7, 0xbb, 0xaa, 0xed, 0xb5,
	0xf6, 0x9a, 0x23, 0xaf, 0xd1, 0x02, 0x53, 0x29, 0x95, 0xa5, 0x56, 0x28, 0x69, 0x7a, 0xf4, 0x49,
	0x8f, 0x76, 0x5f, 0x97, 0xed, 0x15, 0xfe, 0xd2, 0x50, 0xad, 0xa1, 0xe9, 0xf1, 0xa7, 0xdf, 0x07,
	0x49, 0xf6, 0xa6, 0xcf, 0x7d, 0xea, 0xed, 0x2f, 0xdc, 0x6c, 0xe9, 0xb3, 0x64, 0xdf, 0x27, 0xab,
	0x24, 0xad, 0x21, 0x8b, 0x86, 0xd1, 0xe8, 0x7e, 0xf9, 0xd0, 0x1f, 0x9e, 0xd3, 0x1a, 0x52, 0x48,
	0x0e, 0x36, 0xe7, 0xca, 0xe2, 0x61, 0x34, 0x3a, 0x18, 0xbf, 0x46, 0xdb, 0x6e, 0xab, 0x0b, 0x83,
	0x42, 0xb7, 0x8f, 0xd7, 0x1a, 0x4e, 0x65, 0x5b, 0x6f, 0x9e, 0x94, 0xfb, 0xf0, 0xef, 0x67, 0xfa,
	0x22, 0xb9, 0xe7, 0xf7, 0x93, 0x0d, 0x86, 0xd1, 0xe8, 0xc1, 0xf8, 0xc8, 0x37, 0xf0, 0xd9, 0xd0,
	0x85, 0x6d, 0x84, 0xe4, 0x33, 0xba, 0x6c, 0xa1, 0x0c, 0xec, 0xf4, 0x3c, 0x79, 0xf4, 0x77, 0xc0,
	0x2b, 0x00, 0x56, 0x09, 0x0b, 0xb5, 0xc9, 0xee, 0x0c, 0x07, 0xb7, 0xba, 0xa4, 0x41, 0xf9, 0x0e,
	0x80, 0xbd, 0x5f, 0xeb, 0xd2, 0xcf, 0xc9, 0x9e, 0x5b, 0x5a, 0x76, 0xb7, 0x0b, 0x3a, 0xdd, 0x35,
	0x68, 0x7f, 0xad, 0x6f, 0x3b, 0xf1, 0x66, 0xe2, 0x0d, 0xa8, 0xec, 0xfd, 0x27, 0xdf, 0xe2, 0xe4,
	0x64, 0xa1, 0x6a, 0x74, 0xeb, 0xb3, 0x9b, 0x3c, 0xde, 0xb6, 0xc3, 0xe9, 0x3a, 0xd5, 0x34, 0xfa,
	0xf4, 0xa1, 0xf7, 0xe0, 0x6a, 0x49, 0x25, 0x47, 0xaa, 0xe1, 0x98, 0x83, 0xec, 0x32, 0xfb, 0x97,
	0xa6, 0x85, 0xb9, 0xe1, 0xef, 0x78, 0x19, 0xaa, 0x1f, 0xf1, 0xe0, 0xac, 0x28, 0x7e, 0xc6, 0xc7,
	0x67, 0xce, 0xb2, 0x60, 0x06, 0xb9, 0x72, 0x5d, 0xcd, 0x72, 0x54, 0x7a, 0xe6, 0x2f, 0xcf, 0x99,
	0x17, 0xcc, 0xcc, 0x03, 0x67, 0x3e, 0xcb, 0xe7, 0x81, 0xf3, 0x3b, 0x3e, 0x71, 0x00, 0x21, 0x05,
	0x33, 0x84, 0x04, 0x16, 0x21, 0xb3, 0x9c, 0x90, 0xc0, 0xbb, 0xdc, 0xeb, 0x86, 0x7d, 0xfe, 0x27,
	0x00, 0x00, 0xff, 0xff, 0x56, 0x3a, 0xcd, 0x34, 0xc9, 0x03, 0x00, 0x00,
}
