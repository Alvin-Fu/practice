// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/mobile_app_category_constant.proto

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

// A mobile application category constant.
type MobileAppCategoryConstant struct {
	// The resource name of the mobile app category constant.
	// Mobile app category constant resource names have the form:
	//
	// `mobileAppCategoryConstants/{mobile_app_category_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the mobile app category constant.
	Id *wrappers.Int32Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Mobile app category name.
	Name                 *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MobileAppCategoryConstant) Reset()         { *m = MobileAppCategoryConstant{} }
func (m *MobileAppCategoryConstant) String() string { return proto.CompactTextString(m) }
func (*MobileAppCategoryConstant) ProtoMessage()    {}
func (*MobileAppCategoryConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fad794294b415b1, []int{0}
}

func (m *MobileAppCategoryConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileAppCategoryConstant.Unmarshal(m, b)
}
func (m *MobileAppCategoryConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileAppCategoryConstant.Marshal(b, m, deterministic)
}
func (m *MobileAppCategoryConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileAppCategoryConstant.Merge(m, src)
}
func (m *MobileAppCategoryConstant) XXX_Size() int {
	return xxx_messageInfo_MobileAppCategoryConstant.Size(m)
}
func (m *MobileAppCategoryConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileAppCategoryConstant.DiscardUnknown(m)
}

var xxx_messageInfo_MobileAppCategoryConstant proto.InternalMessageInfo

func (m *MobileAppCategoryConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *MobileAppCategoryConstant) GetId() *wrappers.Int32Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MobileAppCategoryConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func init() {
	proto.RegisterType((*MobileAppCategoryConstant)(nil), "google.ads.googleads.v2.resources.MobileAppCategoryConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/mobile_app_category_constant.proto", fileDescriptor_8fad794294b415b1)
}

var fileDescriptor_8fad794294b415b1 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x69, 0x27, 0x82, 0x55, 0x2f, 0x3d, 0xcd, 0x39, 0xc6, 0xa6, 0x0c, 0x06, 0x42, 0x2a,
	0xdd, 0x2d, 0x9e, 0xba, 0x09, 0x43, 0x41, 0x19, 0x13, 0x7a, 0x90, 0x42, 0xc9, 0xda, 0x18, 0x0a,
	0x6d, 0x12, 0x92, 0x6c, 0xe2, 0x33, 0xf8, 0x10, 0x82, 0x47, 0x1f, 0xc5, 0x47, 0xf1, 0x29, 0x64,
	0x49, 0x93, 0x8b, 0xa8, 0xb7, 0x3f, 0xcd, 0x2f, 0xff, 0xef, 0xd7, 0x7c, 0xc1, 0x35, 0x61, 0x8c,
	0xd4, 0x38, 0x42, 0xa5, 0x8c, 0x4c, 0xdc, 0xa5, 0x6d, 0x1c, 0x09, 0x2c, 0xd9, 0x46, 0x14, 0x58,
	0x46, 0x0d, 0x5b, 0x57, 0x35, 0xce, 0x11, 0xe7, 0x79, 0x81, 0x14, 0x26, 0x4c, 0xbc, 0xe4, 0x05,
	0xa3, 0x52, 0x21, 0xaa, 0x00, 0x17, 0x4c, 0xb1, 0x70, 0x64, 0xae, 0x02, 0x54, 0x4a, 0xe0, 0x5a,
	0xc0, 0x36, 0x06, 0xae, 0xa5, 0x37, 0x68, 0x07, 0xe9, 0x0b, 0xeb, 0xcd, 0x53, 0xf4, 0x2c, 0x10,
	0xe7, 0x58, 0x48, 0x53, 0xd1, 0xeb, 0x5b, 0x11, 0x5e, 0x45, 0x88, 0x52, 0xa6, 0x90, 0xaa, 0x18,
	0x6d, 0x4f, 0xcf, 0xde, 0xbc, 0xe0, 0xe4, 0x4e, 0x7b, 0x24, 0x9c, 0xcf, 0x5b, 0x8b, 0x79, 0x2b,
	0x11, 0x9e, 0x07, 0xc7, 0x76, 0x50, 0x4e, 0x51, 0x83, 0xbb, 0xde, 0xd0, 0x9b, 0x1c, 0xac, 0x8e,
	0xec, 0xc7, 0x7b, 0xd4, 0xe0, 0xf0, 0x22, 0xf0, 0xab, 0xb2, 0xeb, 0x0f, 0xbd, 0xc9, 0x61, 0x7c,
	0xda, 0x5a, 0x02, 0x6b, 0x03, 0x6e, 0xa8, 0x9a, 0xc6, 0x29, 0xaa, 0x37, 0x78, 0xe5, 0x57, 0x65,
	0x78, 0x19, 0xec, 0xe9, 0xa2, 0x8e, 0xc6, 0xfb, 0x3f, 0xf0, 0x07, 0x25, 0x2a, 0x4a, 0x0c, 0xaf,
	0xc9, 0xd9, 0xab, 0x1f, 0x8c, 0x0b, 0xd6, 0x80, 0x7f, 0x5f, 0x62, 0x36, 0xf8, 0xf5, 0x47, 0x96,
	0xbb, 0xfa, 0xa5, 0xf7, 0x78, 0xdb, 0x96, 0x10, 0x56, 0x23, 0x4a, 0x00, 0x13, 0x24, 0x22, 0x98,
	0xea, 0xe1, 0x76, 0x49, 0xbc, 0x92, 0x7f, 0xec, 0xec, 0xca, 0xa5, 0x77, 0xbf, 0xb3, 0x48, 0x92,
	0x0f, 0x7f, 0xb4, 0x30, 0x95, 0x49, 0x29, 0x81, 0x89, 0xbb, 0x94, 0xc6, 0x60, 0x65, 0xc9, 0x4f,
	0xcb, 0x64, 0x49, 0x29, 0x33, 0xc7, 0x64, 0x69, 0x9c, 0x39, 0xe6, 0xcb, 0x1f, 0x9b, 0x03, 0x08,
	0x93, 0x52, 0x42, 0xe8, 0x28, 0x08, 0xd3, 0x18, 0x42, 0xc7, 0xad, 0xf7, 0xb5, 0xec, 0xf4, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0xe8, 0x9a, 0x16, 0x12, 0x5f, 0x02, 0x00, 0x00,
}
