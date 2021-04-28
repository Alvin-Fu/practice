// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/landing_page_view.proto

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
	return fileDescriptor_0dcfa31285588a3a, []int{0}
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
	proto.RegisterType((*LandingPageView)(nil), "google.ads.googleads.v1.resources.LandingPageView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/landing_page_view.proto", fileDescriptor_0dcfa31285588a3a)
}

var fileDescriptor_0dcfa31285588a3a = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x69, 0x05, 0xc1, 0xaa, 0x08, 0x65, 0x87, 0x31, 0x86, 0x6c, 0xca, 0x60, 0xa7, 0x94,
	0xea, 0xc9, 0x78, 0xea, 0x0e, 0x0e, 0x44, 0xc6, 0x98, 0xd8, 0x83, 0x14, 0xca, 0xb7, 0xe5, 0x5b,
	0x08, 0x74, 0x49, 0x49, 0xda, 0xcd, 0xab, 0x17, 0x7f, 0x88, 0x47, 0x7f, 0x8a, 0x3f, 0xc5, 0x5f,
	0x21, 0x5d, 0xdb, 0x08, 0x1e, 0xf4, 0xf6, 0xb6, 0xdf, 0xf3, 0xbe, 0xdf, 0x9b, 0xc4, 0xbb, 0xe1,
	0x4a, 0xf1, 0x0c, 0x03, 0x60, 0x26, 0xa8, 0x65, 0xa5, 0xb6, 0x61, 0xa0, 0xd1, 0xa8, 0x52, 0xaf,
	0xd0, 0x04, 0x19, 0x48, 0x26, 0x24, 0x4f, 0x73, 0xe0, 0x98, 0x6e, 0x05, 0xee, 0x48, 0xae, 0x55,
	0xa1, 0xfc, 0x61, 0xcd, 0x13, 0x60, 0x86, 0x58, 0x2b, 0xd9, 0x86, 0xc4, 0x5a, 0x7b, 0xfd, 0x36,
	0x3d, 0x17, 0x01, 0x48, 0xa9, 0x0a, 0x28, 0x84, 0x92, 0xa6, 0x0e, 0xe8, 0x9d, 0x37, 0xd3, 0xfd,
	0xd7, 0xb2, 0x5c, 0x07, 0x3b, 0x0d, 0x79, 0x8e, 0xba, 0x99, 0x5f, 0xbc, 0x39, 0xde, 0xd9, 0x43,
	0xbd, 0x7c, 0x0e, 0x1c, 0x63, 0x81, 0x3b, 0xff, 0xd2, 0x3b, 0x6d, 0xe3, 0x53, 0x09, 0x1b, 0xec,
	0x3a, 0x03, 0x67, 0x7c, 0xb4, 0x38, 0x69, 0x7f, 0xce, 0x60, 0x83, 0xfe, 0xcc, 0xeb, 0x94, 0x12,
	0x5f, 0x72, 0x90, 0x0c, 0x59, 0xba, 0x16, 0x12, 0xb2, 0xb4, 0xd4, 0x59, 0xd7, 0x1d, 0x38, 0xe3,
	0xe3, 0xab, 0x7e, 0xd3, 0x96, 0xb4, 0x7b, 0xc9, 0x63, 0xa1, 0x85, 0xe4, 0x31, 0x64, 0x25, 0x2e,
	0xfc, 0x1f, 0xe7, 0x5d, 0x65, 0x7c, 0xd2, 0xd9, 0xe4, 0xd5, 0xf5, 0x46, 0x2b, 0xb5, 0x21, 0xff,
	0x1e, 0x78, 0xd2, 0xf9, 0xd5, 0x77, 0x5e, 0xad, 0x98, 0x3b, 0xcf, 0xf7, 0x8d, 0x95, 0xab, 0x0c,
	0x24, 0x27, 0x4a, 0xf3, 0x80, 0xa3, 0xdc, 0x17, 0x68, 0xaf, 0x3d, 0x17, 0xe6, 0x8f, 0x57, 0xb8,
	0xb5, 0xea, 0xdd, 0x3d, 0x98, 0x46, 0xd1, 0x87, 0x3b, 0x9c, 0xd6, 0x91, 0x11, 0x33, 0xa4, 0x96,
	0x95, 0x8a, 0x43, 0xb2, 0x68, 0xc9, 0xcf, 0x96, 0x49, 0x22, 0x66, 0x12, 0xcb, 0x24, 0x71, 0x98,
	0x58, 0xe6, 0xcb, 0x1d, 0xd5, 0x03, 0x4a, 0x23, 0x66, 0x28, 0xb5, 0x14, 0xa5, 0x71, 0x48, 0xa9,
	0xe5, 0x96, 0x87, 0xfb, 0xb2, 0xd7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x72, 0xa7, 0xcc, 0x68,
	0x31, 0x02, 0x00, 0x00,
}
