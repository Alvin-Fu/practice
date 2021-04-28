// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad_group_criterion_label.proto

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

// A relationship between an ad group criterion and a label.
type AdGroupCriterionLabel struct {
	// The resource name of the ad group criterion label.
	// Ad group criterion label resource names have the form:
	//
	// `customers/{customer_id}/adGroupCriterionLabels/{ad_group_id}~{criterion_id}~{label_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ad group criterion to which the label is attached.
	AdGroupCriterion *wrappers.StringValue `protobuf:"bytes,2,opt,name=ad_group_criterion,json=adGroupCriterion,proto3" json:"ad_group_criterion,omitempty"`
	// The label assigned to the ad group criterion.
	Label                *wrappers.StringValue `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AdGroupCriterionLabel) Reset()         { *m = AdGroupCriterionLabel{} }
func (m *AdGroupCriterionLabel) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterionLabel) ProtoMessage()    {}
func (*AdGroupCriterionLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_2165ddac79b87814, []int{0}
}

func (m *AdGroupCriterionLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterionLabel.Unmarshal(m, b)
}
func (m *AdGroupCriterionLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterionLabel.Marshal(b, m, deterministic)
}
func (m *AdGroupCriterionLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterionLabel.Merge(m, src)
}
func (m *AdGroupCriterionLabel) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterionLabel.Size(m)
}
func (m *AdGroupCriterionLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterionLabel.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterionLabel proto.InternalMessageInfo

func (m *AdGroupCriterionLabel) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupCriterionLabel) GetAdGroupCriterion() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupCriterion
	}
	return nil
}

func (m *AdGroupCriterionLabel) GetLabel() *wrappers.StringValue {
	if m != nil {
		return m.Label
	}
	return nil
}

func init() {
	proto.RegisterType((*AdGroupCriterionLabel)(nil), "google.ads.googleads.v1.resources.AdGroupCriterionLabel")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad_group_criterion_label.proto", fileDescriptor_2165ddac79b87814)
}

var fileDescriptor_2165ddac79b87814 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x69, 0x87, 0x82, 0x55, 0x41, 0x0a, 0xc2, 0x18, 0x43, 0x36, 0x65, 0xb0, 0x53, 0x42,
	0xe7, 0x2d, 0x5e, 0xec, 0x3c, 0x0c, 0x86, 0xc8, 0x98, 0xd0, 0x83, 0x14, 0x4a, 0xb6, 0xc4, 0x50,
	0x68, 0x93, 0x90, 0xb4, 0xf3, 0x11, 0x7c, 0x0f, 0x8f, 0x3e, 0x83, 0x4f, 0xe0, 0xa3, 0xf8, 0x14,
	0xd2, 0xa6, 0xc9, 0x41, 0x45, 0xbd, 0x7d, 0xed, 0xff, 0x97, 0xef, 0xfb, 0xfe, 0x49, 0x70, 0xcd,
	0x84, 0x60, 0x05, 0x85, 0x98, 0x68, 0x68, 0x64, 0xa3, 0x76, 0x11, 0x54, 0x54, 0x8b, 0x5a, 0x6d,
	0xa9, 0x86, 0x98, 0x64, 0x4c, 0x89, 0x5a, 0x66, 0x5b, 0x95, 0x57, 0x54, 0xe5, 0x82, 0x67, 0x05,
	0xde, 0xd0, 0x02, 0x48, 0x25, 0x2a, 0x11, 0x8e, 0xcd, 0x31, 0x80, 0x89, 0x06, 0xce, 0x01, 0xec,
	0x22, 0xe0, 0x1c, 0x06, 0x43, 0x1b, 0x22, 0x73, 0x88, 0x39, 0x17, 0x15, 0xae, 0x72, 0xc1, 0xb5,
	0x31, 0x18, 0x9c, 0x75, 0xd3, 0xf6, 0x6b, 0x53, 0x3f, 0xc2, 0x27, 0x85, 0xa5, 0xa4, 0xaa, 0x9b,
	0x9f, 0xbf, 0x79, 0xc1, 0x69, 0x4c, 0x16, 0x4d, 0x85, 0x1b, 0xdb, 0xe0, 0xb6, 0x29, 0x10, 0x5e,
	0x04, 0xc7, 0x36, 0x24, 0xe3, 0xb8, 0xa4, 0x7d, 0x6f, 0xe4, 0x4d, 0x0f, 0xd6, 0x47, 0xf6, 0xe7,
	0x1d, 0x2e, 0x69, 0xb8, 0x0c, 0xc2, 0xef, 0x1b, 0xf4, 0xfd, 0x91, 0x37, 0x3d, 0x9c, 0x0d, 0xbb,
	0xc6, 0xc0, 0x66, 0x83, 0xfb, 0x4a, 0xe5, 0x9c, 0x25, 0xb8, 0xa8, 0xe9, 0xfa, 0x04, 0x7f, 0x49,
	0x0d, 0x67, 0xc1, 0x5e, 0xbb, 0x7a, 0xbf, 0xf7, 0x8f, 0xe3, 0x06, 0x9d, 0x3f, 0xfb, 0xc1, 0x64,
	0x2b, 0x4a, 0xf0, 0xe7, 0x35, 0xcd, 0x07, 0x3f, 0x6e, 0xb9, 0x6a, 0xbc, 0x57, 0xde, 0xc3, 0xb2,
	0x33, 0x60, 0xa2, 0xc0, 0x9c, 0x01, 0xa1, 0x18, 0x64, 0x94, 0xb7, 0xc9, 0xf6, 0xe5, 0x64, 0xae,
	0x7f, 0x79, 0xc8, 0x2b, 0xa7, 0x5e, 0xfc, 0xde, 0x22, 0x8e, 0x5f, 0xfd, 0xf1, 0xc2, 0x58, 0xc6,
	0x44, 0x03, 0x23, 0x1b, 0x95, 0x44, 0x60, 0x6d, 0xc9, 0x77, 0xcb, 0xa4, 0x31, 0xd1, 0xa9, 0x63,
	0xd2, 0x24, 0x4a, 0x1d, 0xf3, 0xe1, 0x4f, 0xcc, 0x00, 0xa1, 0x98, 0x68, 0x84, 0x1c, 0x85, 0x50,
	0x12, 0x21, 0xe4, 0xb8, 0xcd, 0x7e, 0x5b, 0xf6, 0xf2, 0x33, 0x00, 0x00, 0xff, 0xff, 0x47, 0x81,
	0x4c, 0x85, 0x74, 0x02, 0x00, 0x00,
}
