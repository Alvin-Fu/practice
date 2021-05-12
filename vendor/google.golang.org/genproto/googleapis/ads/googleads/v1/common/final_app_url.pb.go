// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/final_app_url.proto

package common

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

// A URL for deep linking into an app for the given operating system.
type FinalAppUrl struct {
	// The operating system targeted by this URL. Required.
	OsType enums.AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType `protobuf:"varint,1,opt,name=os_type,json=osType,proto3,enum=google.ads.googleads.v1.enums.AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType" json:"os_type,omitempty"`
	// The app deep link URL. Deep links specify a location in an app that
	// corresponds to the content you'd like to show, and should be of the form
	// {scheme}://{host_path}
	// The scheme identifies which app to open. For your app, you can use a custom
	// scheme that starts with the app's name. The host and path specify the
	// unique location in the app where your content exists.
	// Example: "exampleapp://productid_1234". Required.
	Url                  *wrappers.StringValue `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FinalAppUrl) Reset()         { *m = FinalAppUrl{} }
func (m *FinalAppUrl) String() string { return proto.CompactTextString(m) }
func (*FinalAppUrl) ProtoMessage()    {}
func (*FinalAppUrl) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b7b2ec96b577e51, []int{0}
}

func (m *FinalAppUrl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FinalAppUrl.Unmarshal(m, b)
}
func (m *FinalAppUrl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FinalAppUrl.Marshal(b, m, deterministic)
}
func (m *FinalAppUrl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinalAppUrl.Merge(m, src)
}
func (m *FinalAppUrl) XXX_Size() int {
	return xxx_messageInfo_FinalAppUrl.Size(m)
}
func (m *FinalAppUrl) XXX_DiscardUnknown() {
	xxx_messageInfo_FinalAppUrl.DiscardUnknown(m)
}

var xxx_messageInfo_FinalAppUrl proto.InternalMessageInfo

func (m *FinalAppUrl) GetOsType() enums.AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType {
	if m != nil {
		return m.OsType
	}
	return enums.AppUrlOperatingSystemTypeEnum_UNSPECIFIED
}

func (m *FinalAppUrl) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

func init() {
	proto.RegisterType((*FinalAppUrl)(nil), "google.ads.googleads.v1.common.FinalAppUrl")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/final_app_url.proto", fileDescriptor_8b7b2ec96b577e51)
}

var fileDescriptor_8b7b2ec96b577e51 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xbf, 0x6a, 0xf3, 0x30,
	0x14, 0xc5, 0xb1, 0x03, 0xf9, 0xc0, 0x81, 0x8f, 0x92, 0x29, 0x84, 0x10, 0x42, 0xa6, 0x4c, 0x12,
	0x4e, 0x37, 0x75, 0x72, 0xfa, 0x27, 0x63, 0x43, 0xd2, 0x7a, 0x28, 0x06, 0xa3, 0xc4, 0x8a, 0x30,
	0xd8, 0xba, 0x42, 0x92, 0x53, 0xf2, 0x3a, 0x1d, 0x3b, 0xf4, 0x41, 0xfa, 0x28, 0xed, 0x4b, 0x14,
	0x4b, 0x76, 0xe8, 0xe2, 0x4e, 0xba, 0xf6, 0xfd, 0x9d, 0x73, 0x8f, 0xae, 0x82, 0x25, 0x07, 0xe0,
	0x05, 0xc3, 0x34, 0xd3, 0xd8, 0x95, 0x75, 0x75, 0x0a, 0xf1, 0x01, 0xca, 0x12, 0x04, 0x3e, 0xe6,
	0x82, 0x16, 0x29, 0x95, 0x32, 0xad, 0x54, 0x81, 0xa4, 0x02, 0x03, 0xc3, 0xa9, 0x03, 0x11, 0xcd,
	0x34, 0xba, 0x68, 0xd0, 0x29, 0x44, 0x4e, 0x33, 0x8e, 0xba, 0x3c, 0x99, 0xa8, 0x4a, 0x8d, 0x1b,
	0xb3, 0x14, 0x24, 0x53, 0xd4, 0xe4, 0x82, 0xa7, 0xfa, 0xac, 0x0d, 0x2b, 0x53, 0x73, 0x96, 0xcc,
	0x8d, 0x18, 0x4f, 0x5a, 0x0b, 0x99, 0x63, 0x2a, 0x04, 0x18, 0x6a, 0x72, 0x10, 0xba, 0xe9, 0x36,
	0x01, 0xb0, 0xfd, 0xda, 0x57, 0x47, 0xfc, 0xaa, 0xa8, 0x94, 0x4c, 0x35, 0xfd, 0xf9, 0x87, 0x17,
	0x0c, 0x1e, 0xea, 0xe0, 0x91, 0x94, 0xcf, 0xaa, 0x18, 0x42, 0xf0, 0x0f, 0xb4, 0xb5, 0x1f, 0x79,
	0x33, 0x6f, 0xf1, 0x7f, 0x19, 0xa3, 0xae, 0x2b, 0xd8, 0x88, 0xc8, 0xe9, 0x1e, 0xdb, 0x80, 0x3b,
	0x9b, 0xef, 0xe9, 0x2c, 0xd9, 0xbd, 0xa8, 0xca, 0xee, 0xee, 0xb6, 0x0f, 0xba, 0x3e, 0x87, 0x28,
	0xe8, 0x55, 0xaa, 0x18, 0xf9, 0x33, 0x6f, 0x31, 0x58, 0x4e, 0xda, 0x61, 0x6d, 0x5c, 0xb4, 0x33,
	0x2a, 0x17, 0x3c, 0xa6, 0x45, 0xc5, 0xb6, 0x35, 0xb8, 0xfa, 0xf6, 0x82, 0xf9, 0x01, 0x4a, 0xf4,
	0xf7, 0x62, 0x57, 0x57, 0xbf, 0x2e, 0xb5, 0xa9, 0xcd, 0x36, 0xde, 0xcb, 0x5d, 0xa3, 0xe1, 0x50,
	0x50, 0xc1, 0x11, 0x28, 0x8e, 0x39, 0x13, 0x76, 0x54, 0xbb, 0x7c, 0x99, 0xeb, 0xae, 0xf7, 0xbd,
	0x71, 0xc7, 0x9b, 0xdf, 0x5b, 0x47, 0xd1, 0xbb, 0x3f, 0x5d, 0x3b, 0xb3, 0x28, 0xd3, 0xc8, 0x95,
	0x75, 0x15, 0x87, 0xe8, 0xd6, 0x62, 0x9f, 0x2d, 0x90, 0x44, 0x99, 0x4e, 0x2e, 0x40, 0x12, 0x87,
	0x89, 0x03, 0xbe, 0xfc, 0xb9, 0xfb, 0x4b, 0x48, 0x94, 0x69, 0x42, 0x2e, 0x08, 0x21, 0x71, 0x48,
	0x88, 0x83, 0xf6, 0x7d, 0x9b, 0xee, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0x12, 0xd0, 0x1f, 0x4d,
	0x7c, 0x02, 0x00, 0x00,
}
