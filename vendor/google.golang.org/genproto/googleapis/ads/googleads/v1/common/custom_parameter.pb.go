// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/custom_parameter.proto

package common

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

// A mapping that can be used by custom parameter tags in a
// `tracking_url_template`, `final_urls`, or `mobile_final_urls`.
type CustomParameter struct {
	// The key matching the parameter tag name.
	Key *wrappers.StringValue `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The value to be substituted.
	Value                *wrappers.StringValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CustomParameter) Reset()         { *m = CustomParameter{} }
func (m *CustomParameter) String() string { return proto.CompactTextString(m) }
func (*CustomParameter) ProtoMessage()    {}
func (*CustomParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d23b9a41ba81714, []int{0}
}

func (m *CustomParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomParameter.Unmarshal(m, b)
}
func (m *CustomParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomParameter.Marshal(b, m, deterministic)
}
func (m *CustomParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomParameter.Merge(m, src)
}
func (m *CustomParameter) XXX_Size() int {
	return xxx_messageInfo_CustomParameter.Size(m)
}
func (m *CustomParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomParameter.DiscardUnknown(m)
}

var xxx_messageInfo_CustomParameter proto.InternalMessageInfo

func (m *CustomParameter) GetKey() *wrappers.StringValue {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *CustomParameter) GetValue() *wrappers.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*CustomParameter)(nil), "google.ads.googleads.v1.common.CustomParameter")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/custom_parameter.proto", fileDescriptor_6d23b9a41ba81714)
}

var fileDescriptor_6d23b9a41ba81714 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4a, 0xf4, 0x40,
	0x14, 0x85, 0x49, 0x96, 0xff, 0x2f, 0x62, 0x21, 0x04, 0x8b, 0x65, 0x59, 0x16, 0x49, 0x65, 0x75,
	0x87, 0xac, 0xd8, 0x8c, 0x55, 0x76, 0x85, 0x6d, 0x83, 0x42, 0x0a, 0x09, 0xc8, 0x6c, 0x32, 0x0e,
	0xc1, 0x64, 0xee, 0x30, 0x33, 0x89, 0xf8, 0x3a, 0x96, 0x3e, 0x8a, 0x8f, 0x62, 0xe3, 0x2b, 0x48,
	0x32, 0x49, 0x0a, 0x41, 0xb1, 0xca, 0xb9, 0xc9, 0x77, 0xce, 0xb9, 0xb9, 0xc1, 0x95, 0x40, 0x14,
	0x35, 0x27, 0xac, 0x34, 0xc4, 0xc9, 0x5e, 0x75, 0x31, 0x29, 0xb0, 0x69, 0x50, 0x92, 0xa2, 0x35,
	0x16, 0x9b, 0x07, 0xc5, 0x34, 0x6b, 0xb8, 0xe5, 0x1a, 0x94, 0x46, 0x8b, 0xe1, 0xc6, 0xb1, 0xc0,
	0x4a, 0x03, 0xb3, 0x0d, 0xba, 0x18, 0x9c, 0x6d, 0xb5, 0x9e, 0x62, 0x55, 0x45, 0x98, 0x94, 0x68,
	0x99, 0xad, 0x50, 0x1a, 0xe7, 0x5e, 0x8d, 0x6e, 0x32, 0x4c, 0xc7, 0xf6, 0x91, 0x3c, 0x6b, 0xa6,
	0x14, 0xd7, 0xe3, 0xf7, 0xa8, 0x0d, 0x4e, 0xf7, 0x43, 0x6f, 0x3a, 0xd5, 0x86, 0x10, 0x2c, 0x9e,
	0xf8, 0xcb, 0xd2, 0x3b, 0xf7, 0x2e, 0x4e, 0xb6, 0xeb, 0xb1, 0x13, 0xa6, 0x00, 0xb8, 0xb3, 0xba,
	0x92, 0x22, 0x63, 0x75, 0xcb, 0x6f, 0x7b, 0x30, 0xdc, 0x06, 0xff, 0xba, 0x7e, 0x5a, 0xfa, 0x7f,
	0x70, 0x38, 0x74, 0xf7, 0xe9, 0x05, 0x51, 0x81, 0x0d, 0xfc, 0xfe, 0x6f, 0xbb, 0xb3, 0x6f, 0xbb,
	0xa5, 0x7d, 0x64, 0xea, 0xdd, 0xdf, 0x8c, 0x3e, 0x81, 0x35, 0x93, 0x02, 0x50, 0x0b, 0x22, 0xb8,
	0x1c, 0x0a, 0xa7, 0xd3, 0xaa, 0xca, 0xfc, 0x74, 0xe9, 0x6b, 0xf7, 0x78, 0xf5, 0x17, 0x87, 0x24,
	0x79, 0xf3, 0x37, 0x07, 0x17, 0x96, 0x94, 0x06, 0x9c, 0xec, 0x55, 0x16, 0xc3, 0x7e, 0xc0, 0xde,
	0x27, 0x20, 0x4f, 0x4a, 0x93, 0xcf, 0x40, 0x9e, 0xc5, 0xb9, 0x03, 0x3e, 0xfc, 0xc8, 0xbd, 0xa5,
	0x34, 0x29, 0x0d, 0xa5, 0x33, 0x42, 0x69, 0x16, 0x53, 0xea, 0xa0, 0xe3, 0xff, 0x61, 0xbb, 0xcb,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0x34, 0x2b, 0xad, 0x06, 0x02, 0x00, 0x00,
}
