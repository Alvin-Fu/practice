// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/keyword_plan_campaign.proto

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

// A Keyword Plan campaign.
// Max number of keyword plan campaigns per plan allowed: 1.
type KeywordPlanCampaign struct {
	// The resource name of the Keyword Plan campaign.
	// KeywordPlanCampaign resource names have the form:
	//
	// `customers/{customer_id}/keywordPlanCampaigns/{kp_campaign_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The keyword plan this campaign belongs to.
	KeywordPlan *wrappers.StringValue `protobuf:"bytes,2,opt,name=keyword_plan,json=keywordPlan,proto3" json:"keyword_plan,omitempty"`
	// The ID of the Keyword Plan campaign.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the Keyword Plan campaign.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The languages targeted for the Keyword Plan campaign.
	// Max allowed: 1.
	LanguageConstants []*wrappers.StringValue `protobuf:"bytes,5,rep,name=language_constants,json=languageConstants,proto3" json:"language_constants,omitempty"`
	// Targeting network.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	KeywordPlanNetwork enums.KeywordPlanNetworkEnum_KeywordPlanNetwork `protobuf:"varint,6,opt,name=keyword_plan_network,json=keywordPlanNetwork,proto3,enum=google.ads.googleads.v2.enums.KeywordPlanNetworkEnum_KeywordPlanNetwork" json:"keyword_plan_network,omitempty"`
	// A default max cpc bid in micros, and in the account currency, for all ad
	// groups under the campaign.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	CpcBidMicros *wrappers.Int64Value `protobuf:"bytes,7,opt,name=cpc_bid_micros,json=cpcBidMicros,proto3" json:"cpc_bid_micros,omitempty"`
	// The geo targets.
	// Max number allowed: 20.
	GeoTargets           []*KeywordPlanGeoTarget `protobuf:"bytes,8,rep,name=geo_targets,json=geoTargets,proto3" json:"geo_targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeywordPlanCampaign) Reset()         { *m = KeywordPlanCampaign{} }
func (m *KeywordPlanCampaign) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanCampaign) ProtoMessage()    {}
func (*KeywordPlanCampaign) Descriptor() ([]byte, []int) {
	return fileDescriptor_d287701d9b77d16a, []int{0}
}

func (m *KeywordPlanCampaign) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanCampaign.Unmarshal(m, b)
}
func (m *KeywordPlanCampaign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanCampaign.Marshal(b, m, deterministic)
}
func (m *KeywordPlanCampaign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanCampaign.Merge(m, src)
}
func (m *KeywordPlanCampaign) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanCampaign.Size(m)
}
func (m *KeywordPlanCampaign) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanCampaign.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanCampaign proto.InternalMessageInfo

func (m *KeywordPlanCampaign) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *KeywordPlanCampaign) GetKeywordPlan() *wrappers.StringValue {
	if m != nil {
		return m.KeywordPlan
	}
	return nil
}

func (m *KeywordPlanCampaign) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *KeywordPlanCampaign) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *KeywordPlanCampaign) GetLanguageConstants() []*wrappers.StringValue {
	if m != nil {
		return m.LanguageConstants
	}
	return nil
}

func (m *KeywordPlanCampaign) GetKeywordPlanNetwork() enums.KeywordPlanNetworkEnum_KeywordPlanNetwork {
	if m != nil {
		return m.KeywordPlanNetwork
	}
	return enums.KeywordPlanNetworkEnum_UNSPECIFIED
}

func (m *KeywordPlanCampaign) GetCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidMicros
	}
	return nil
}

func (m *KeywordPlanCampaign) GetGeoTargets() []*KeywordPlanGeoTarget {
	if m != nil {
		return m.GeoTargets
	}
	return nil
}

// A geo target.
// Next ID: 3
type KeywordPlanGeoTarget struct {
	// Required. The resource name of the geo target.
	GeoTargetConstant    *wrappers.StringValue `protobuf:"bytes,1,opt,name=geo_target_constant,json=geoTargetConstant,proto3" json:"geo_target_constant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *KeywordPlanGeoTarget) Reset()         { *m = KeywordPlanGeoTarget{} }
func (m *KeywordPlanGeoTarget) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanGeoTarget) ProtoMessage()    {}
func (*KeywordPlanGeoTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_d287701d9b77d16a, []int{1}
}

func (m *KeywordPlanGeoTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanGeoTarget.Unmarshal(m, b)
}
func (m *KeywordPlanGeoTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanGeoTarget.Marshal(b, m, deterministic)
}
func (m *KeywordPlanGeoTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanGeoTarget.Merge(m, src)
}
func (m *KeywordPlanGeoTarget) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanGeoTarget.Size(m)
}
func (m *KeywordPlanGeoTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanGeoTarget.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanGeoTarget proto.InternalMessageInfo

func (m *KeywordPlanGeoTarget) GetGeoTargetConstant() *wrappers.StringValue {
	if m != nil {
		return m.GeoTargetConstant
	}
	return nil
}

func init() {
	proto.RegisterType((*KeywordPlanCampaign)(nil), "google.ads.googleads.v2.resources.KeywordPlanCampaign")
	proto.RegisterType((*KeywordPlanGeoTarget)(nil), "google.ads.googleads.v2.resources.KeywordPlanGeoTarget")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/keyword_plan_campaign.proto", fileDescriptor_d287701d9b77d16a)
}

var fileDescriptor_d287701d9b77d16a = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x8a, 0x13, 0x31,
	0x14, 0x86, 0xe9, 0x6c, 0x5d, 0x35, 0xad, 0x0b, 0x66, 0xf7, 0x62, 0x58, 0x17, 0xe9, 0xae, 0x2c,
	0x14, 0x84, 0x8c, 0x8c, 0xa2, 0x32, 0x22, 0x32, 0x5d, 0xa4, 0xea, 0xea, 0x52, 0xaa, 0x14, 0x91,
	0xc2, 0x90, 0x4e, 0x62, 0x08, 0xed, 0x24, 0x43, 0x92, 0xd9, 0xa2, 0xf7, 0xbe, 0x88, 0x97, 0x3e,
	0x85, 0xd7, 0x3e, 0x8a, 0x4f, 0x21, 0xcd, 0x4c, 0x66, 0x8b, 0xdb, 0x75, 0xbc, 0x3b, 0x93, 0xfc,
	0xdf, 0x39, 0xff, 0x9c, 0x73, 0x02, 0x9e, 0x33, 0x29, 0xd9, 0x82, 0x06, 0x98, 0xe8, 0xa0, 0x0c,
	0x57, 0xd1, 0x79, 0x18, 0x28, 0xaa, 0x65, 0xa1, 0x52, 0xaa, 0x83, 0x39, 0xfd, 0xb2, 0x94, 0x8a,
	0x24, 0xf9, 0x02, 0x8b, 0x24, 0xc5, 0x59, 0x8e, 0x39, 0x13, 0x28, 0x57, 0xd2, 0x48, 0x78, 0x58,
	0x32, 0x08, 0x13, 0x8d, 0x6a, 0x1c, 0x9d, 0x87, 0xa8, 0xc6, 0xf7, 0x9f, 0x5e, 0x55, 0x81, 0x8a,
	0x22, 0xfb, 0x2b, 0xbb, 0xa0, 0x66, 0x29, 0xd5, 0xbc, 0x4c, 0xbe, 0x7f, 0xb7, 0x22, 0xed, 0xd7,
	0xac, 0xf8, 0x1c, 0x2c, 0x15, 0xce, 0x73, 0xaa, 0x74, 0x75, 0x7f, 0xe0, 0x32, 0xe7, 0x3c, 0xc0,
	0x42, 0x48, 0x83, 0x0d, 0x97, 0xa2, 0xba, 0x3d, 0xfa, 0xd9, 0x06, 0xbb, 0xa7, 0x65, 0xf2, 0xd1,
	0x02, 0x8b, 0x93, 0xca, 0x38, 0xbc, 0x07, 0x6e, 0x39, 0x73, 0x89, 0xc0, 0x19, 0xf5, 0x5b, 0xbd,
	0x56, 0xff, 0xe6, 0xb8, 0xeb, 0x0e, 0xcf, 0x70, 0x46, 0xe1, 0x0b, 0xd0, 0x5d, 0x37, 0xe6, 0x7b,
	0xbd, 0x56, 0xbf, 0x13, 0x1e, 0x54, 0xff, 0x88, 0x9c, 0x23, 0xf4, 0xde, 0x28, 0x2e, 0xd8, 0x04,
	0x2f, 0x0a, 0x3a, 0xee, 0xcc, 0x2f, 0xaa, 0xc1, 0xfb, 0xc0, 0xe3, 0xc4, 0xdf, 0xb2, 0xd8, 0x9d,
	0x4b, 0xd8, 0x6b, 0x61, 0x1e, 0x3f, 0x2a, 0x29, 0x8f, 0x13, 0xf8, 0x00, 0xb4, 0xad, 0x93, 0xf6,
	0x7f, 0x54, 0xb1, 0x4a, 0x78, 0x0a, 0xe0, 0x02, 0x0b, 0x56, 0x60, 0x46, 0x93, 0x54, 0x0a, 0x6d,
	0xb0, 0x30, 0xda, 0xbf, 0xd6, 0xdb, 0x6a, 0xe4, 0x6f, 0x3b, 0xee, 0xc4, 0x61, 0xf0, 0x2b, 0xd8,
	0xdb, 0x34, 0x05, 0x7f, 0xbb, 0xd7, 0xea, 0xef, 0x84, 0xaf, 0xd0, 0x55, 0x33, 0xb6, 0x03, 0x44,
	0x6b, 0x3d, 0x3e, 0x2b, 0xc1, 0x97, 0xa2, 0xc8, 0x36, 0x1c, 0x8f, 0xe1, 0xfc, 0xd2, 0x19, 0x8c,
	0xc1, 0x4e, 0x9a, 0xa7, 0xc9, 0x8c, 0x93, 0x24, 0xe3, 0xa9, 0x92, 0xda, 0xbf, 0xde, 0xdc, 0xb3,
	0x6e, 0x9a, 0xa7, 0x03, 0x4e, 0xde, 0x59, 0x00, 0x7e, 0x04, 0x1d, 0x46, 0x65, 0x62, 0xb0, 0x62,
	0xd4, 0x68, 0xff, 0x86, 0x6d, 0xc2, 0x13, 0xd4, 0xb8, 0x99, 0xeb, 0x16, 0x87, 0x54, 0x7e, 0xb0,
	0xfc, 0x18, 0x30, 0x17, 0xea, 0x23, 0x02, 0xf6, 0x36, 0x69, 0xe0, 0x5b, 0xb0, 0x7b, 0x51, 0xb1,
	0xee, 0xbf, 0x5d, 0xa4, 0xc6, 0xf6, 0xd7, 0xe9, 0x5d, 0xff, 0x07, 0xdf, 0x3c, 0x70, 0x9c, 0xca,
	0xac, 0xd9, 0xf0, 0xc0, 0xdf, 0xb0, 0xcf, 0xa3, 0x55, 0x95, 0x51, 0xeb, 0xd3, 0x9b, 0x0a, 0x67,
	0x72, 0x35, 0x60, 0x24, 0x15, 0x0b, 0x18, 0x15, 0xd6, 0x83, 0x7b, 0x76, 0x39, 0xd7, 0xff, 0x78,
	0xe7, 0xcf, 0xea, 0xe8, 0xbb, 0xb7, 0x35, 0x8c, 0xe3, 0x1f, 0xde, 0xe1, 0xb0, 0x4c, 0x19, 0x13,
	0x8d, 0xca, 0x70, 0x15, 0x4d, 0x42, 0x34, 0x76, 0xca, 0x5f, 0x4e, 0x33, 0x8d, 0x89, 0x9e, 0xd6,
	0x9a, 0xe9, 0x24, 0x9c, 0xd6, 0x9a, 0xdf, 0xde, 0x71, 0x79, 0x11, 0x45, 0x31, 0xd1, 0x51, 0x54,
	0xab, 0xa2, 0x68, 0x12, 0x46, 0x51, 0xad, 0x9b, 0x6d, 0x5b, 0xb3, 0x0f, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x4f, 0x90, 0xa8, 0x8f, 0x93, 0x04, 0x00, 0x00,
}
