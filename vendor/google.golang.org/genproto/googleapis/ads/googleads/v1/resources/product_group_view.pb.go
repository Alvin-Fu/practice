// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/product_group_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// A product group view.
type ProductGroupView struct {
	// The resource name of the product group view.
	// Product group view resource names have the form:
	//
	// `customers/{customer_id}/productGroupViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductGroupView) Reset()         { *m = ProductGroupView{} }
func (m *ProductGroupView) String() string { return proto.CompactTextString(m) }
func (*ProductGroupView) ProtoMessage()    {}
func (*ProductGroupView) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3a6978f2d357ed9, []int{0}
}

func (m *ProductGroupView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductGroupView.Unmarshal(m, b)
}
func (m *ProductGroupView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductGroupView.Marshal(b, m, deterministic)
}
func (m *ProductGroupView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductGroupView.Merge(m, src)
}
func (m *ProductGroupView) XXX_Size() int {
	return xxx_messageInfo_ProductGroupView.Size(m)
}
func (m *ProductGroupView) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductGroupView.DiscardUnknown(m)
}

var xxx_messageInfo_ProductGroupView proto.InternalMessageInfo

func (m *ProductGroupView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*ProductGroupView)(nil), "google.ads.googleads.v1.resources.ProductGroupView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/product_group_view.proto", fileDescriptor_e3a6978f2d357ed9)
}

var fileDescriptor_e3a6978f2d357ed9 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xb1, 0x4a, 0xc4, 0x30,
	0x18, 0xc7, 0x69, 0x05, 0xc1, 0xa2, 0x20, 0x07, 0x82, 0x88, 0x83, 0xa7, 0x1c, 0x38, 0x25, 0x04,
	0x07, 0x21, 0x4e, 0xb9, 0xa5, 0xe0, 0x20, 0xe5, 0x86, 0x0e, 0x52, 0x28, 0xb1, 0x09, 0x21, 0x70,
	0xcd, 0x17, 0x92, 0xb6, 0xb7, 0xfb, 0x28, 0x8e, 0x3e, 0x8a, 0x8f, 0xe2, 0x53, 0x48, 0x2f, 0x26,
	0x83, 0x83, 0xb7, 0xfd, 0x49, 0x7e, 0xff, 0xdf, 0xf7, 0xf1, 0x15, 0x54, 0x01, 0xa8, 0xad, 0xc4,
	0x5c, 0x78, 0x1c, 0xe2, 0x9c, 0x26, 0x82, 0x9d, 0xf4, 0x30, 0xba, 0x4e, 0x7a, 0x6c, 0x1d, 0x88,
	0xb1, 0x1b, 0x5a, 0xe5, 0x60, 0xb4, 0xed, 0xa4, 0xe5, 0x0e, 0x59, 0x07, 0x03, 0x2c, 0x96, 0xa1,
	0x80, 0xb8, 0xf0, 0x28, 0x75, 0xd1, 0x44, 0x50, 0xea, 0x5e, 0x5d, 0x47, 0xbd, 0xd5, 0x98, 0x1b,
	0x03, 0x03, 0x1f, 0x34, 0x18, 0x1f, 0x04, 0xb7, 0x8f, 0xc5, 0x79, 0x15, 0xe4, 0xe5, 0xec, 0xae,
	0xb5, 0xdc, 0x2d, 0xee, 0x8a, 0xb3, 0x58, 0x6f, 0x0d, 0xef, 0xe5, 0x65, 0x76, 0x93, 0xdd, 0x9f,
	0x6c, 0x4e, 0xe3, 0xe3, 0x0b, 0xef, 0xe5, 0xfa, 0x3d, 0x2f, 0x56, 0x1d, 0xf4, 0xe8, 0xe0, 0x02,
	0xeb, 0x8b, 0xbf, 0x03, 0xaa, 0x79, 0x72, 0x95, 0xbd, 0x3e, 0xff, 0x76, 0x15, 0x6c, 0xb9, 0x51,
	0x08, 0x9c, 0xc2, 0x4a, 0x9a, 0xfd, 0x5e, 0xf1, 0x10, 0x56, 0xfb, 0x7f, 0xee, 0xf2, 0x94, 0xd2,
	0x47, 0x7e, 0x54, 0x32, 0xf6, 0x99, 0x2f, 0xcb, 0xa0, 0x64, 0xc2, 0xa3, 0x10, 0xe7, 0x54, 0x13,
	0xb4, 0x89, 0xe4, 0x57, 0x64, 0x1a, 0x26, 0x7c, 0x93, 0x98, 0xa6, 0x26, 0x4d, 0x62, 0xbe, 0xf3,
	0x55, 0xf8, 0xa0, 0x94, 0x09, 0x4f, 0x69, 0xa2, 0x28, 0xad, 0x09, 0xa5, 0x89, 0x7b, 0x3b, 0xde,
	0x2f, 0xfb, 0xf0, 0x13, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x25, 0x41, 0x4a, 0xc3, 0x01, 0x00, 0x00,
}
