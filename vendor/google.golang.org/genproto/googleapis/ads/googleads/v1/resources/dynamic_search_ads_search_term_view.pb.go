// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/dynamic_search_ads_search_term_view.proto

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

// A dynamic search ads search term view.
type DynamicSearchAdsSearchTermView struct {
	// The resource name of the dynamic search ads search term view.
	// Dynamic search ads search term view resource names have the form:
	//
	//
	// `customers/{customer_id}/dynamicSearchAdsSearchTermViews/{ad_group_id}~{search_term_fp}~{headline_fp}~{landing_page_fp}~{page_url_fp}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Search term
	//
	// This field is read-only.
	SearchTerm *wrappers.StringValue `protobuf:"bytes,2,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	// The dynamically generated headline of the Dynamic Search Ad.
	//
	// This field is read-only.
	Headline *wrappers.StringValue `protobuf:"bytes,3,opt,name=headline,proto3" json:"headline,omitempty"`
	// The dynamically selected landing page URL of the impression.
	//
	// This field is read-only.
	LandingPage *wrappers.StringValue `protobuf:"bytes,4,opt,name=landing_page,json=landingPage,proto3" json:"landing_page,omitempty"`
	// The URL of page feed item served for the impression.
	//
	// This field is read-only.
	PageUrl              *wrappers.StringValue `protobuf:"bytes,5,opt,name=page_url,json=pageUrl,proto3" json:"page_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DynamicSearchAdsSearchTermView) Reset()         { *m = DynamicSearchAdsSearchTermView{} }
func (m *DynamicSearchAdsSearchTermView) String() string { return proto.CompactTextString(m) }
func (*DynamicSearchAdsSearchTermView) ProtoMessage()    {}
func (*DynamicSearchAdsSearchTermView) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ffd212889021798, []int{0}
}

func (m *DynamicSearchAdsSearchTermView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicSearchAdsSearchTermView.Unmarshal(m, b)
}
func (m *DynamicSearchAdsSearchTermView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicSearchAdsSearchTermView.Marshal(b, m, deterministic)
}
func (m *DynamicSearchAdsSearchTermView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicSearchAdsSearchTermView.Merge(m, src)
}
func (m *DynamicSearchAdsSearchTermView) XXX_Size() int {
	return xxx_messageInfo_DynamicSearchAdsSearchTermView.Size(m)
}
func (m *DynamicSearchAdsSearchTermView) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicSearchAdsSearchTermView.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicSearchAdsSearchTermView proto.InternalMessageInfo

func (m *DynamicSearchAdsSearchTermView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *DynamicSearchAdsSearchTermView) GetSearchTerm() *wrappers.StringValue {
	if m != nil {
		return m.SearchTerm
	}
	return nil
}

func (m *DynamicSearchAdsSearchTermView) GetHeadline() *wrappers.StringValue {
	if m != nil {
		return m.Headline
	}
	return nil
}

func (m *DynamicSearchAdsSearchTermView) GetLandingPage() *wrappers.StringValue {
	if m != nil {
		return m.LandingPage
	}
	return nil
}

func (m *DynamicSearchAdsSearchTermView) GetPageUrl() *wrappers.StringValue {
	if m != nil {
		return m.PageUrl
	}
	return nil
}

func init() {
	proto.RegisterType((*DynamicSearchAdsSearchTermView)(nil), "google.ads.googleads.v1.resources.DynamicSearchAdsSearchTermView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/dynamic_search_ads_search_term_view.proto", fileDescriptor_9ffd212889021798)
}

var fileDescriptor_9ffd212889021798 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6a, 0xdc, 0x30,
	0x14, 0x66, 0x9c, 0xfe, 0xa4, 0x9a, 0x74, 0xe3, 0x95, 0x09, 0x21, 0x24, 0x0d, 0x81, 0xac, 0x64,
	0xdc, 0x2e, 0x5a, 0x54, 0x4a, 0x71, 0x28, 0x04, 0x5a, 0x28, 0xc3, 0xa4, 0xf5, 0xa2, 0x18, 0xcc,
	0x8b, 0xf5, 0xaa, 0x08, 0x6c, 0xc9, 0x48, 0xf6, 0x0c, 0xbd, 0x45, 0xcf, 0xd0, 0x65, 0x4f, 0xd0,
	0x33, 0xf4, 0x28, 0x3d, 0x45, 0xb1, 0x65, 0x89, 0xae, 0xa6, 0xb3, 0xfb, 0xec, 0xf7, 0xfd, 0x49,
	0x7a, 0xe4, 0x83, 0xd0, 0x5a, 0x34, 0x98, 0x02, 0xb7, 0xa9, 0x83, 0x23, 0xda, 0x64, 0xa9, 0x41,
	0xab, 0x07, 0x53, 0xa3, 0x4d, 0xf9, 0x37, 0x05, 0xad, 0xac, 0x2b, 0x8b, 0x60, 0xea, 0xfb, 0x0a,
	0xb8, 0xf5, 0xb0, 0x47, 0xd3, 0x56, 0x1b, 0x89, 0x5b, 0xda, 0x19, 0xdd, 0xeb, 0xf8, 0xdc, 0x39,
	0x50, 0xe0, 0x96, 0x06, 0x33, 0xba, 0xc9, 0x68, 0x30, 0x3b, 0x3e, 0xf1, 0x79, 0x9d, 0x4c, 0x41,
	0x29, 0xdd, 0x43, 0x2f, 0xb5, 0xb2, 0xce, 0xe0, 0xf8, 0x74, 0x9e, 0x4e, 0x5f, 0x77, 0xc3, 0xd7,
	0x74, 0x6b, 0xa0, 0xeb, 0xd0, 0xcc, 0xf3, 0x67, 0xbf, 0x22, 0x72, 0xfa, 0xce, 0xd5, 0xb9, 0x9d,
	0x2a, 0xe4, 0xdc, 0x3a, 0xf0, 0x09, 0x4d, 0x5b, 0x48, 0xdc, 0xc6, 0x17, 0xe4, 0xa9, 0x4f, 0xab,
	0x14, 0xb4, 0x98, 0x2c, 0xce, 0x16, 0x57, 0x4f, 0xd6, 0x47, 0xfe, 0xe7, 0x47, 0x68, 0x31, 0x7e,
	0x43, 0x96, 0xff, 0x1c, 0x21, 0x89, 0xce, 0x16, 0x57, 0xcb, 0xe7, 0x27, 0x73, 0x67, 0xea, 0xd3,
	0xe9, 0x6d, 0x6f, 0xa4, 0x12, 0x05, 0x34, 0x03, 0xae, 0x89, 0x0d, 0x39, 0xf1, 0x2b, 0x72, 0x78,
	0x8f, 0xc0, 0x1b, 0xa9, 0x30, 0x39, 0xd8, 0x43, 0x1b, 0xd8, 0xf1, 0x5b, 0x72, 0xd4, 0x80, 0xe2,
	0x52, 0x89, 0xaa, 0x03, 0x81, 0xc9, 0x83, 0x3d, 0xd4, 0xcb, 0x59, 0xb1, 0x02, 0x81, 0xf1, 0x4b,
	0x72, 0x38, 0x0a, 0xab, 0xc1, 0x34, 0xc9, 0xc3, 0x3d, 0xc4, 0x8f, 0x47, 0xf6, 0x67, 0xd3, 0x5c,
	0x7f, 0x8f, 0xc8, 0x65, 0xad, 0x5b, 0xfa, 0xdf, 0x27, 0xba, 0xbe, 0xd8, 0x7d, 0xc3, 0xab, 0x31,
	0x66, 0xb5, 0xf8, 0xf2, 0x7e, 0x76, 0x12, 0xba, 0x01, 0x25, 0xa8, 0x36, 0x22, 0x15, 0xa8, 0xa6,
	0x12, 0x7e, 0x93, 0x3a, 0x69, 0x77, 0x2c, 0xd6, 0xeb, 0x80, 0x7e, 0x44, 0x07, 0x37, 0x79, 0xfe,
	0x33, 0x3a, 0xbf, 0x71, 0x96, 0x39, 0xb7, 0xd4, 0xc1, 0x11, 0x15, 0x19, 0x5d, 0x7b, 0xe6, 0x6f,
	0xcf, 0x29, 0x73, 0x6e, 0xcb, 0xc0, 0x29, 0x8b, 0xac, 0x0c, 0x9c, 0x3f, 0xd1, 0xa5, 0x1b, 0x30,
	0x96, 0x73, 0xcb, 0x58, 0x60, 0x31, 0x56, 0x64, 0x8c, 0x05, 0xde, 0xdd, 0xa3, 0xa9, 0xec, 0x8b,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x72, 0x5b, 0xf8, 0xe3, 0x04, 0x03, 0x00, 0x00,
}
