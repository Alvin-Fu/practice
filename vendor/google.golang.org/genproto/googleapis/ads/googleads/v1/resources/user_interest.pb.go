// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/user_interest.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
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

// A user interest: a particular interest-based vertical to be targeted.
type UserInterest struct {
	// The resource name of the user interest.
	// User interest resource names have the form:
	//
	// `customers/{customer_id}/userInterests/{user_interest_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Taxonomy type of the user interest.
	TaxonomyType enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType `protobuf:"varint,2,opt,name=taxonomy_type,json=taxonomyType,proto3,enum=google.ads.googleads.v1.enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType" json:"taxonomy_type,omitempty"`
	// The ID of the user interest.
	UserInterestId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=user_interest_id,json=userInterestId,proto3" json:"user_interest_id,omitempty"`
	// The name of the user interest.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The parent of the user interest.
	UserInterestParent *wrappers.StringValue `protobuf:"bytes,5,opt,name=user_interest_parent,json=userInterestParent,proto3" json:"user_interest_parent,omitempty"`
	// True if the user interest is launched to all channels and locales.
	LaunchedToAll *wrappers.BoolValue `protobuf:"bytes,6,opt,name=launched_to_all,json=launchedToAll,proto3" json:"launched_to_all,omitempty"`
	// Availability information of the user interest.
	Availabilities       []*common.CriterionCategoryAvailability `protobuf:"bytes,7,rep,name=availabilities,proto3" json:"availabilities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *UserInterest) Reset()         { *m = UserInterest{} }
func (m *UserInterest) String() string { return proto.CompactTextString(m) }
func (*UserInterest) ProtoMessage()    {}
func (*UserInterest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01746905ff80465b, []int{0}
}

func (m *UserInterest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInterest.Unmarshal(m, b)
}
func (m *UserInterest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInterest.Marshal(b, m, deterministic)
}
func (m *UserInterest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInterest.Merge(m, src)
}
func (m *UserInterest) XXX_Size() int {
	return xxx_messageInfo_UserInterest.Size(m)
}
func (m *UserInterest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInterest.DiscardUnknown(m)
}

var xxx_messageInfo_UserInterest proto.InternalMessageInfo

func (m *UserInterest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *UserInterest) GetTaxonomyType() enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType {
	if m != nil {
		return m.TaxonomyType
	}
	return enums.UserInterestTaxonomyTypeEnum_UNSPECIFIED
}

func (m *UserInterest) GetUserInterestId() *wrappers.Int64Value {
	if m != nil {
		return m.UserInterestId
	}
	return nil
}

func (m *UserInterest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UserInterest) GetUserInterestParent() *wrappers.StringValue {
	if m != nil {
		return m.UserInterestParent
	}
	return nil
}

func (m *UserInterest) GetLaunchedToAll() *wrappers.BoolValue {
	if m != nil {
		return m.LaunchedToAll
	}
	return nil
}

func (m *UserInterest) GetAvailabilities() []*common.CriterionCategoryAvailability {
	if m != nil {
		return m.Availabilities
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInterest)(nil), "google.ads.googleads.v1.resources.UserInterest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/user_interest.proto", fileDescriptor_01746905ff80465b)
}

var fileDescriptor_01746905ff80465b = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x26, 0x6d, 0x5d, 0x31, 0xdb, 0x56, 0x0d, 0x5e, 0x84, 0xba, 0x48, 0x57, 0x59, 0xe8, 0xd5,
	0xc4, 0xd6, 0x9f, 0x8b, 0x88, 0x48, 0xba, 0x2e, 0x4b, 0xbd, 0x58, 0x4a, 0xad, 0xbd, 0x90, 0x42,
	0x98, 0x26, 0xc7, 0x38, 0x30, 0x99, 0x13, 0x66, 0x26, 0xd5, 0xbc, 0xce, 0x5e, 0xfa, 0x28, 0x3e,
	0x8a, 0x0f, 0x21, 0xd2, 0xfc, 0x91, 0xee, 0x52, 0xf5, 0xee, 0x24, 0xe7, 0xfb, 0x39, 0x67, 0xe6,
	0x1b, 0xf3, 0x55, 0x84, 0x18, 0x71, 0x70, 0x68, 0xa8, 0x9c, 0xa2, 0xdc, 0x55, 0xdb, 0xb1, 0x23,
	0x41, 0x61, 0x2a, 0x03, 0x50, 0x4e, 0xaa, 0x40, 0xfa, 0x4c, 0x68, 0x90, 0xa0, 0x34, 0x49, 0x24,
	0x6a, 0xb4, 0x4e, 0x0b, 0x2c, 0xa1, 0xa1, 0x22, 0x35, 0x8d, 0x6c, 0xc7, 0xa4, 0xa6, 0x0d, 0xde,
	0x1f, 0x52, 0x0e, 0x30, 0x8e, 0x51, 0x38, 0x81, 0x64, 0x1a, 0x24, 0x43, 0xe1, 0x07, 0x54, 0x43,
	0x84, 0x32, 0xf3, 0xe9, 0x96, 0x32, 0x4e, 0x37, 0x8c, 0x33, 0x9d, 0x15, 0x46, 0x83, 0x77, 0x87,
	0x54, 0x40, 0xa4, 0xf1, 0x8d, 0xd9, 0x7c, 0x4d, 0xbf, 0xa3, 0xc0, 0x38, 0xf3, 0x75, 0x96, 0x40,
	0x29, 0x70, 0x52, 0x09, 0x24, 0xcc, 0xa1, 0x42, 0xa0, 0xa6, 0x9a, 0xa1, 0x50, 0x65, 0xf7, 0x49,
	0xd9, 0xcd, 0xbf, 0x36, 0xe9, 0x17, 0xe7, 0x9b, 0xa4, 0x49, 0x02, 0xb2, 0xec, 0x3f, 0xbd, 0xee,
	0x98, 0xdd, 0x4f, 0x0a, 0xe4, 0xac, 0xb4, 0xb0, 0x9e, 0x99, 0xbd, 0x6a, 0x45, 0x5f, 0xd0, 0x18,
	0x6c, 0x63, 0x68, 0x8c, 0xee, 0x2d, 0xba, 0xd5, 0xcf, 0x2b, 0x1a, 0x83, 0x95, 0x99, 0xbd, 0xbd,
	0x51, 0xec, 0xd6, 0xd0, 0x18, 0xf5, 0x27, 0x4b, 0x72, 0xe8, 0xd4, 0xf2, 0x65, 0x48, 0xd3, 0x68,
	0x59, 0xf2, 0x97, 0x59, 0x02, 0x17, 0x22, 0x8d, 0x0f, 0x36, 0x17, 0x5d, 0xdd, 0xf8, 0xb2, 0x2e,
	0xcc, 0x07, 0xfb, 0x67, 0xc2, 0x42, 0xbb, 0x3d, 0x34, 0x46, 0xc7, 0x93, 0xc7, 0x95, 0x7b, 0xb5,
	0x2b, 0x99, 0x09, 0xfd, 0xfa, 0xe5, 0x8a, 0xf2, 0x14, 0x16, 0xfd, 0xb4, 0x21, 0x3f, 0x0b, 0xad,
	0xe7, 0x66, 0x27, 0xdf, 0xae, 0x93, 0x53, 0x4f, 0x6e, 0x51, 0x3f, 0x6a, 0xc9, 0x44, 0x54, 0x70,
	0x73, 0xa4, 0x75, 0x65, 0x3e, 0xda, 0x37, 0x4e, 0xa8, 0x04, 0xa1, 0xed, 0x3b, 0xff, 0xa1, 0x60,
	0x35, 0xdd, 0xe7, 0x39, 0xcf, 0x9a, 0x9a, 0xf7, 0x39, 0x4d, 0x45, 0xf0, 0x15, 0x42, 0x5f, 0xa3,
	0x4f, 0x39, 0xb7, 0x8f, 0x72, 0xa9, 0xc1, 0x2d, 0xa9, 0x29, 0x22, 0x2f, 0x84, 0x7a, 0x15, 0x65,
	0x89, 0x1e, 0xe7, 0x16, 0x98, 0xfd, 0x46, 0xa4, 0x18, 0x28, 0xfb, 0xee, 0xb0, 0x3d, 0x3a, 0x9e,
	0xbc, 0x3d, 0x78, 0x11, 0x45, 0x36, 0xc9, 0x79, 0x95, 0xcd, 0xf3, 0x32, 0x9a, 0x5e, 0x23, 0x99,
	0x8b, 0x1b, 0xa2, 0xd3, 0xdf, 0x86, 0x79, 0x16, 0x60, 0x4c, 0xfe, 0xf9, 0x26, 0xa6, 0x0f, 0x9b,
	0xb7, 0x38, 0xdf, 0xcd, 0x3f, 0x37, 0x3e, 0x7f, 0x28, 0x79, 0x11, 0x72, 0x2a, 0x22, 0x82, 0x32,
	0x72, 0x22, 0x10, 0xf9, 0x76, 0x55, 0xe4, 0x13, 0xa6, 0xfe, 0xf2, 0x42, 0xdf, 0xd4, 0xd5, 0x75,
	0xab, 0x7d, 0xe9, 0x79, 0x3f, 0x5a, 0xa7, 0x97, 0x85, 0xa4, 0x17, 0x2a, 0x52, 0x94, 0xbb, 0x6a,
	0x35, 0x26, 0x8b, 0x0a, 0xf9, 0xb3, 0xc2, 0xac, 0xbd, 0x50, 0xad, 0x6b, 0xcc, 0x7a, 0x35, 0x5e,
	0xd7, 0x98, 0x5f, 0xad, 0xb3, 0xa2, 0xe1, 0xba, 0x5e, 0xa8, 0x5c, 0xb7, 0x46, 0xb9, 0xee, 0x6a,
	0xec, 0xba, 0x35, 0x6e, 0x73, 0x94, 0x0f, 0xfb, 0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10,
	0x33, 0x6a, 0x5e, 0x4d, 0x04, 0x00, 0x00,
}
