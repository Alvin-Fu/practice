// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/landing_page_view.proto

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

// A landing page view with metrics aggregated at the unexpanded final URL
// level.
type LandingPageView struct {
	// The resource name of the landing page view.
	// Landing page view resource names have the form:
	//
	// `customers/{customer_id}/landingPageViews/{unexpanded_final_url_fingerprint}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The advertiser-specified final URL.
	UnexpandedFinalUrl   *wrappers.StringValue `protobuf:"bytes,2,opt,name=unexpanded_final_url,json=unexpandedFinalUrl,proto3" json:"unexpanded_final_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LandingPageView) Reset()         { *m = LandingPageView{} }
func (m *LandingPageView) String() string { return proto.CompactTextString(m) }
func (*LandingPageView) ProtoMessage()    {}
func (*LandingPageView) Descriptor() ([]byte, []int) {
	return fileDescriptor_902f5a0985d352f9, []int{0}
}

func (m *LandingPageView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LandingPageView.Unmarshal(m, b)
}
func (m *LandingPageView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LandingPageView.Marshal(b, m, deterministic)
}
func (m *LandingPageView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LandingPageView.Merge(m, src)
}
func (m *LandingPageView) XXX_Size() int {
	return xxx_messageInfo_LandingPageView.Size(m)
}
func (m *LandingPageView) XXX_DiscardUnknown() {
	xxx_messageInfo_LandingPageView.DiscardUnknown(m)
}

var xxx_messageInfo_LandingPageView proto.InternalMessageInfo

func (m *LandingPageView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *LandingPageView) GetUnexpandedFinalUrl() *wrappers.StringValue {
	if m != nil {
		return m.UnexpandedFinalUrl
	}
	return nil
}

func init() {
	proto.RegisterType((*LandingPageView)(nil), "google.ads.googleads.v2.resources.LandingPageView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/landing_page_view.proto", fileDescriptor_902f5a0985d352f9)
}

var fileDescriptor_902f5a0985d352f9 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x6a, 0xeb, 0x30,
	0x14, 0x86, 0xb1, 0x2f, 0x5c, 0xb8, 0xbe, 0xf7, 0x52, 0x30, 0x19, 0x42, 0x08, 0x25, 0x69, 0x09,
	0x64, 0x92, 0xc1, 0x9d, 0xaa, 0x4e, 0xce, 0xd0, 0x40, 0x29, 0x21, 0xa4, 0xd4, 0x43, 0x31, 0x98,
	0x93, 0xe8, 0x44, 0x08, 0x1c, 0xc9, 0x48, 0x76, 0xd2, 0xb5, 0x4b, 0x1f, 0xa4, 0x63, 0x1f, 0xa5,
	0x8f, 0xd2, 0xa7, 0x28, 0x8e, 0x2d, 0x15, 0x3a, 0xb4, 0xdb, 0x8f, 0xce, 0xf7, 0xff, 0xe7, 0x97,
	0x14, 0x5c, 0x72, 0xa5, 0x78, 0x81, 0x11, 0x30, 0x13, 0xb5, 0xb2, 0x51, 0xfb, 0x38, 0xd2, 0x68,
	0x54, 0xad, 0x37, 0x68, 0xa2, 0x02, 0x24, 0x13, 0x92, 0xe7, 0x25, 0x70, 0xcc, 0xf7, 0x02, 0x0f,
	0xa4, 0xd4, 0xaa, 0x52, 0xe1, 0xb8, 0xe5, 0x09, 0x30, 0x43, 0x9c, 0x95, 0xec, 0x63, 0xe2, 0xac,
	0x83, 0xd3, 0x2e, 0xfd, 0x68, 0x58, 0xd7, 0xdb, 0xe8, 0xa0, 0xa1, 0x2c, 0x51, 0x9b, 0x36, 0x62,
	0x30, 0xb4, 0xdb, 0x4b, 0x11, 0x81, 0x94, 0xaa, 0x82, 0x4a, 0x28, 0xd9, 0x4d, 0xcf, 0x9e, 0xbd,
	0xe0, 0xe4, 0xb6, 0x5d, 0xbe, 0x04, 0x8e, 0xa9, 0xc0, 0x43, 0x78, 0x1e, 0xfc, 0xb7, 0xf1, 0xb9,
	0x84, 0x1d, 0xf6, 0xbd, 0x91, 0x37, 0xfd, 0xb3, 0xfa, 0x67, 0x0f, 0x17, 0xb0, 0xc3, 0x70, 0x11,
	0xf4, 0x6a, 0x89, 0x8f, 0x25, 0x48, 0x86, 0x2c, 0xdf, 0x0a, 0x09, 0x45, 0x5e, 0xeb, 0xa2, 0xef,
	0x8f, 0xbc, 0xe9, 0xdf, 0x78, 0xd8, 0xb5, 0x25, 0xb6, 0x15, 0xb9, 0xab, 0xb4, 0x90, 0x3c, 0x85,
	0xa2, 0xc6, 0x55, 0xf8, 0xe9, 0xbc, 0x6e, 0x8c, 0xf7, 0xba, 0x98, 0x3d, 0xf9, 0xc1, 0x64, 0xa3,
	0x76, 0xe4, 0xc7, 0x0b, 0xcf, 0x7a, 0x5f, 0xfa, 0x2e, 0x9b, 0x15, 0x4b, 0xef, 0xe1, 0xa6, 0xb3,
	0x72, 0x55, 0x80, 0xe4, 0x44, 0x69, 0x1e, 0x71, 0x94, 0xc7, 0x02, 0xf6, 0xd9, 0x4b, 0x61, 0xbe,
	0xf9, 0x85, 0x2b, 0xa7, 0x5e, 0xfc, 0x5f, 0xf3, 0x24, 0x79, 0xf5, 0xc7, 0xf3, 0x36, 0x32, 0x61,
	0x86, 0xb4, 0xb2, 0x51, 0x69, 0x4c, 0x56, 0x96, 0x7c, 0xb3, 0x4c, 0x96, 0x30, 0x93, 0x39, 0x26,
	0x4b, 0xe3, 0xcc, 0x31, 0xef, 0xfe, 0xa4, 0x1d, 0x50, 0x9a, 0x30, 0x43, 0xa9, 0xa3, 0x28, 0x4d,
	0x63, 0x4a, 0x1d, 0xb7, 0xfe, 0x7d, 0x2c, 0x7b, 0xf1, 0x11, 0x00, 0x00, 0xff, 0xff, 0x03, 0xbb,
	0x11, 0x52, 0x31, 0x02, 0x00, 0x00,
}
