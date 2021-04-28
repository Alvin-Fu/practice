// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/domain_category.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A category generated automatically by crawling a domain. If a campaign uses
// the DynamicSearchAdsSetting, then domain categories will be generated for
// the domain. The categories can be targeted using WebpageConditionInfo.
// See: https://support.google.com/google-ads/answer/2471185
type DomainCategory struct {
	// The resource name of the domain category.
	// Domain category resource names have the form:
	//
	// `customers/{customer_id}/domainCategories/{campaign_id}~{category_base64}~{language_code}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The campaign this category is recommended for.
	Campaign *wrappers.StringValue `protobuf:"bytes,2,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Recommended category for the website domain. e.g. if you have a website
	// about electronics, the categories could be "cameras", "televisions", etc.
	Category *wrappers.StringValue `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	// The language code specifying the language of the website. e.g. "en" for
	// English. The language can be specified in the DynamicSearchAdsSetting
	// required for dynamic search ads. This is the language of the pages from
	// your website that you want Google Ads to find, create ads for,
	// and match searches with.
	LanguageCode *wrappers.StringValue `protobuf:"bytes,4,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// The domain for the website. The domain can be specified in the
	// DynamicSearchAdsSetting required for dynamic search ads.
	Domain *wrappers.StringValue `protobuf:"bytes,5,opt,name=domain,proto3" json:"domain,omitempty"`
	// Fraction of pages on your site that this category matches.
	CoverageFraction *wrappers.DoubleValue `protobuf:"bytes,6,opt,name=coverage_fraction,json=coverageFraction,proto3" json:"coverage_fraction,omitempty"`
	// The position of this category in the set of categories. Lower numbers
	// indicate a better match for the domain. null indicates not recommended.
	CategoryRank *wrappers.Int64Value `protobuf:"bytes,7,opt,name=category_rank,json=categoryRank,proto3" json:"category_rank,omitempty"`
	// Indicates whether this category has sub-categories.
	HasChildren *wrappers.BoolValue `protobuf:"bytes,8,opt,name=has_children,json=hasChildren,proto3" json:"has_children,omitempty"`
	// The recommended cost per click for the category.
	RecommendedCpcBidMicros *wrappers.Int64Value `protobuf:"bytes,9,opt,name=recommended_cpc_bid_micros,json=recommendedCpcBidMicros,proto3" json:"recommended_cpc_bid_micros,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}             `json:"-"`
	XXX_unrecognized        []byte               `json:"-"`
	XXX_sizecache           int32                `json:"-"`
}

func (m *DomainCategory) Reset()         { *m = DomainCategory{} }
func (m *DomainCategory) String() string { return proto.CompactTextString(m) }
func (*DomainCategory) ProtoMessage()    {}
func (*DomainCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b5d3cffbd1a298, []int{0}
}

func (m *DomainCategory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DomainCategory.Unmarshal(m, b)
}
func (m *DomainCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DomainCategory.Marshal(b, m, deterministic)
}
func (m *DomainCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DomainCategory.Merge(m, src)
}
func (m *DomainCategory) XXX_Size() int {
	return xxx_messageInfo_DomainCategory.Size(m)
}
func (m *DomainCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_DomainCategory.DiscardUnknown(m)
}

var xxx_messageInfo_DomainCategory proto.InternalMessageInfo

func (m *DomainCategory) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *DomainCategory) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *DomainCategory) GetCategory() *wrappers.StringValue {
	if m != nil {
		return m.Category
	}
	return nil
}

func (m *DomainCategory) GetLanguageCode() *wrappers.StringValue {
	if m != nil {
		return m.LanguageCode
	}
	return nil
}

func (m *DomainCategory) GetDomain() *wrappers.StringValue {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *DomainCategory) GetCoverageFraction() *wrappers.DoubleValue {
	if m != nil {
		return m.CoverageFraction
	}
	return nil
}

func (m *DomainCategory) GetCategoryRank() *wrappers.Int64Value {
	if m != nil {
		return m.CategoryRank
	}
	return nil
}

func (m *DomainCategory) GetHasChildren() *wrappers.BoolValue {
	if m != nil {
		return m.HasChildren
	}
	return nil
}

func (m *DomainCategory) GetRecommendedCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.RecommendedCpcBidMicros
	}
	return nil
}

func init() {
	proto.RegisterType((*DomainCategory)(nil), "google.ads.googleads.v2.resources.DomainCategory")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/domain_category.proto", fileDescriptor_68b5d3cffbd1a298)
}

var fileDescriptor_68b5d3cffbd1a298 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x61, 0x6b, 0xd4, 0x30,
	0x1c, 0xc6, 0xb9, 0xdb, 0x3c, 0xb7, 0xec, 0x4e, 0xb4, 0xbe, 0xb0, 0x9c, 0x43, 0x36, 0x65, 0xb0,
	0x57, 0x29, 0xd4, 0xa1, 0x52, 0x11, 0xec, 0xdd, 0x70, 0x4c, 0x50, 0xc6, 0x09, 0x87, 0xc8, 0x41,
	0xc9, 0x25, 0x59, 0x2e, 0xac, 0xcd, 0xbf, 0x24, 0xed, 0x89, 0xef, 0xfc, 0x2c, 0xbe, 0xf4, 0x93,
	0x88, 0x1f, 0xc5, 0x4f, 0x21, 0x6d, 0x92, 0xa2, 0x0c, 0xdd, 0xbd, 0xfb, 0xd3, 0x3c, 0xbf, 0xe7,
	0x79, 0x92, 0x26, 0xe8, 0xb9, 0x00, 0x10, 0x39, 0x8f, 0x08, 0x33, 0x91, 0x1d, 0x9b, 0x69, 0x1d,
	0x47, 0x9a, 0x1b, 0xa8, 0x35, 0xe5, 0x26, 0x62, 0x50, 0x10, 0xa9, 0x32, 0x4a, 0x2a, 0x2e, 0x40,
	0x7f, 0xc1, 0xa5, 0x86, 0x0a, 0x82, 0x43, 0xab, 0xc6, 0x84, 0x19, 0xdc, 0x81, 0x78, 0x1d, 0xe3,
	0x0e, 0x1c, 0x3f, 0x72, 0xde, 0x2d, 0xb0, 0xac, 0x2f, 0xa3, 0xcf, 0x9a, 0x94, 0x25, 0xd7, 0xc6,
	0x5a, 0x8c, 0xf7, 0x7d, 0x76, 0x29, 0x23, 0xa2, 0x14, 0x54, 0xa4, 0x92, 0xa0, 0xdc, 0xea, 0xe3,
	0x1f, 0xdb, 0xe8, 0xce, 0x69, 0x1b, 0x3d, 0x75, 0xc9, 0xc1, 0x13, 0x34, 0xf2, 0xee, 0x99, 0x22,
	0x05, 0x0f, 0x7b, 0x07, 0xbd, 0xe3, 0xdd, 0xd9, 0xd0, 0x7f, 0x7c, 0x4f, 0x0a, 0x1e, 0xbc, 0x40,
	0x3b, 0x94, 0x14, 0x25, 0x91, 0x42, 0x85, 0xfd, 0x83, 0xde, 0xf1, 0x5e, 0xbc, 0xef, 0x0a, 0x62,
	0x5f, 0x04, 0x7f, 0xa8, 0xb4, 0x54, 0x62, 0x4e, 0xf2, 0x9a, 0xcf, 0x3a, 0xb5, 0x25, 0x6d, 0x54,
	0xb8, 0xb5, 0x19, 0xe9, 0x8a, 0xa5, 0x68, 0x94, 0x13, 0x25, 0x6a, 0x22, 0x78, 0x46, 0x81, 0xf1,
	0x70, 0x7b, 0x03, 0x7c, 0xe8, 0x91, 0x29, 0x30, 0x1e, 0x9c, 0xa0, 0x81, 0x3d, 0xe8, 0xf0, 0xd6,
	0x06, 0xac, 0xd3, 0x06, 0xe7, 0xe8, 0x1e, 0x85, 0x35, 0xd7, 0x4d, 0xf0, 0xa5, 0x26, 0xb4, 0x39,
	0xc0, 0x70, 0xf0, 0x0f, 0x83, 0x53, 0xa8, 0x97, 0x39, 0xb7, 0x06, 0x77, 0x3d, 0xf6, 0xc6, 0x51,
	0xc1, 0x6b, 0x34, 0xf2, 0xfb, 0xc9, 0x34, 0x51, 0x57, 0xe1, 0xed, 0xd6, 0xe6, 0xe1, 0x35, 0x9b,
	0x73, 0x55, 0x3d, 0x3b, 0x71, 0x5b, 0xf0, 0xc4, 0x8c, 0xa8, 0xab, 0xe0, 0x15, 0x1a, 0xae, 0x88,
	0xc9, 0xe8, 0x4a, 0xe6, 0x4c, 0x73, 0x15, 0xee, 0xb4, 0x06, 0xe3, 0x6b, 0x06, 0x13, 0x80, 0xdc,
	0xf2, 0x7b, 0x2b, 0x62, 0xa6, 0x4e, 0x1e, 0x7c, 0x44, 0x63, 0xcd, 0x29, 0x14, 0x05, 0x57, 0x8c,
	0xb3, 0x8c, 0x96, 0x34, 0x5b, 0x4a, 0x96, 0x15, 0x92, 0x6a, 0x30, 0xe1, 0xee, 0xcd, 0x6d, 0x1e,
	0xfc, 0x81, 0x4f, 0x4b, 0x3a, 0x91, 0xec, 0x5d, 0xcb, 0x4e, 0xbe, 0xf6, 0xd1, 0x11, 0x85, 0x02,
	0xdf, 0x78, 0x65, 0x27, 0xf7, 0xff, 0xbe, 0x71, 0x17, 0x4d, 0xca, 0x45, 0xef, 0xd3, 0x5b, 0x47,
	0x0a, 0x68, 0xfe, 0x19, 0x06, 0x2d, 0x22, 0xc1, 0x55, 0xdb, 0xc1, 0xbf, 0x9a, 0x52, 0x9a, 0xff,
	0x3c, 0xa2, 0x97, 0xdd, 0xf4, 0xad, 0xbf, 0x75, 0x96, 0xa6, 0xdf, 0xfb, 0x87, 0x67, 0xd6, 0x32,
	0x65, 0x06, 0xdb, 0xb1, 0x99, 0xe6, 0x31, 0x9e, 0x79, 0xe5, 0x4f, 0xaf, 0x59, 0xa4, 0xcc, 0x2c,
	0x3a, 0xcd, 0x62, 0x1e, 0x2f, 0x3a, 0xcd, 0xaf, 0xfe, 0x91, 0x5d, 0x48, 0x92, 0x94, 0x99, 0x24,
	0xe9, 0x54, 0x49, 0x32, 0x8f, 0x93, 0xa4, 0xd3, 0x2d, 0x07, 0x6d, 0xd9, 0xa7, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xa0, 0x28, 0x56, 0xa8, 0xf0, 0x03, 0x00, 0x00,
}
