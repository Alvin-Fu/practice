// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/mobile_device_constant.proto

package resources

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

// A mobile device constant.
type MobileDeviceConstant struct {
	// The resource name of the mobile device constant.
	// Mobile device constant resource names have the form:
	//
	// `mobileDeviceConstants/{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the mobile device constant.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the mobile device.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The manufacturer of the mobile device.
	ManufacturerName *wrappers.StringValue `protobuf:"bytes,4,opt,name=manufacturer_name,json=manufacturerName,proto3" json:"manufacturer_name,omitempty"`
	// The operating system of the mobile device.
	OperatingSystemName *wrappers.StringValue `protobuf:"bytes,5,opt,name=operating_system_name,json=operatingSystemName,proto3" json:"operating_system_name,omitempty"`
	// The type of mobile device.
	Type                 enums.MobileDeviceTypeEnum_MobileDeviceType `protobuf:"varint,6,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.MobileDeviceTypeEnum_MobileDeviceType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *MobileDeviceConstant) Reset()         { *m = MobileDeviceConstant{} }
func (m *MobileDeviceConstant) String() string { return proto.CompactTextString(m) }
func (*MobileDeviceConstant) ProtoMessage()    {}
func (*MobileDeviceConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f49c94eff431093, []int{0}
}

func (m *MobileDeviceConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileDeviceConstant.Unmarshal(m, b)
}
func (m *MobileDeviceConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileDeviceConstant.Marshal(b, m, deterministic)
}
func (m *MobileDeviceConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileDeviceConstant.Merge(m, src)
}
func (m *MobileDeviceConstant) XXX_Size() int {
	return xxx_messageInfo_MobileDeviceConstant.Size(m)
}
func (m *MobileDeviceConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileDeviceConstant.DiscardUnknown(m)
}

var xxx_messageInfo_MobileDeviceConstant proto.InternalMessageInfo

func (m *MobileDeviceConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *MobileDeviceConstant) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MobileDeviceConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MobileDeviceConstant) GetManufacturerName() *wrappers.StringValue {
	if m != nil {
		return m.ManufacturerName
	}
	return nil
}

func (m *MobileDeviceConstant) GetOperatingSystemName() *wrappers.StringValue {
	if m != nil {
		return m.OperatingSystemName
	}
	return nil
}

func (m *MobileDeviceConstant) GetType() enums.MobileDeviceTypeEnum_MobileDeviceType {
	if m != nil {
		return m.Type
	}
	return enums.MobileDeviceTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*MobileDeviceConstant)(nil), "google.ads.googleads.v1.resources.MobileDeviceConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/mobile_device_constant.proto", fileDescriptor_5f49c94eff431093)
}

var fileDescriptor_5f49c94eff431093 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x6a, 0x14, 0x31,
	0x14, 0xc6, 0x99, 0xd9, 0xb5, 0x60, 0xfc, 0x83, 0x8e, 0x0a, 0x6b, 0x2d, 0xb2, 0x55, 0x0a, 0x0b,
	0x42, 0xc6, 0xa9, 0xd2, 0x8b, 0x11, 0x84, 0xa9, 0x95, 0x52, 0x41, 0x59, 0xb6, 0xb2, 0x88, 0x2c,
	0x2c, 0xd9, 0xc9, 0x69, 0x08, 0x6c, 0xfe, 0x90, 0x64, 0x56, 0xf6, 0x05, 0xbc, 0xf7, 0x15, 0xbc,
	0xf4, 0x51, 0x7c, 0x14, 0x9f, 0x42, 0x36, 0x99, 0x09, 0xc5, 0x5a, 0xed, 0xdd, 0x99, 0x9c, 0xef,
	0xfb, 0xe5, 0xcb, 0x9c, 0x83, 0x5e, 0x33, 0xa5, 0xd8, 0x12, 0x72, 0x42, 0x6d, 0x1e, 0xca, 0x4d,
	0xb5, 0x2a, 0x72, 0x03, 0x56, 0x35, 0xa6, 0x06, 0x9b, 0x0b, 0xb5, 0xe0, 0x4b, 0x98, 0x53, 0x58,
	0xf1, 0x1a, 0xe6, 0xb5, 0x92, 0xd6, 0x11, 0xe9, 0xb0, 0x36, 0xca, 0xa9, 0x6c, 0x37, 0x98, 0x30,
	0xa1, 0x16, 0x47, 0x3f, 0x5e, 0x15, 0x38, 0xfa, 0xb7, 0x0f, 0x2e, 0xbb, 0x02, 0x64, 0x23, 0xfe,
	0xc4, 0xbb, 0xb5, 0x86, 0x80, 0xde, 0xde, 0xe9, 0x7c, 0x9a, 0xe7, 0x44, 0x4a, 0xe5, 0x88, 0xe3,
	0x4a, 0xda, 0xb6, 0xfb, 0xb8, 0xed, 0xfa, 0xaf, 0x45, 0x73, 0x96, 0x7f, 0x31, 0x44, 0x6b, 0x30,
	0x6d, 0xff, 0xc9, 0xb7, 0x1e, 0xba, 0xff, 0xde, 0xa3, 0x8f, 0x3c, 0xf9, 0x4d, 0x9b, 0x3b, 0x7b,
	0x8a, 0x6e, 0x75, 0xd9, 0xe6, 0x92, 0x08, 0x18, 0x24, 0xc3, 0x64, 0x74, 0x7d, 0x72, 0xb3, 0x3b,
	0xfc, 0x40, 0x04, 0x64, 0xcf, 0x50, 0xca, 0xe9, 0x20, 0x1d, 0x26, 0xa3, 0x1b, 0xfb, 0x8f, 0xda,
	0x87, 0xe1, 0xee, 0x2a, 0x7c, 0x22, 0xdd, 0xc1, 0xcb, 0x29, 0x59, 0x36, 0x30, 0x49, 0x39, 0xcd,
	0x9e, 0xa3, 0xbe, 0x07, 0xf5, 0xbc, 0x7c, 0xe7, 0x82, 0xfc, 0xd4, 0x19, 0x2e, 0x59, 0xd0, 0x7b,
	0x65, 0x76, 0x82, 0xee, 0x0a, 0x22, 0x9b, 0x33, 0x52, 0xbb, 0xc6, 0x80, 0x09, 0x39, 0xfa, 0x57,
	0xb0, 0xdf, 0x39, 0x6f, 0xf3, 0x49, 0xc7, 0xe8, 0x81, 0xd2, 0x60, 0x88, 0xe3, 0x92, 0xcd, 0xed,
	0xda, 0x3a, 0x10, 0x01, 0x77, 0xed, 0x0a, 0xb8, 0x7b, 0xd1, 0x7a, 0xea, 0x9d, 0x9e, 0xf8, 0x09,
	0xf5, 0x37, 0x53, 0x18, 0x6c, 0x0d, 0x93, 0xd1, 0xed, 0xfd, 0x23, 0x7c, 0xd9, 0x84, 0xfd, 0xf8,
	0xf0, 0xf9, 0x7f, 0xfc, 0x71, 0xad, 0xe1, 0xad, 0x6c, 0xc4, 0x85, 0xc3, 0x89, 0x27, 0x1e, 0x7e,
	0x4d, 0xd1, 0x5e, 0xad, 0x04, 0xfe, 0xef, 0xce, 0x1c, 0x3e, 0xfc, 0xdb, 0xe8, 0xc6, 0x9b, 0x27,
	0x8c, 0x93, 0xcf, 0xef, 0x5a, 0x3f, 0x53, 0x4b, 0x22, 0x19, 0x56, 0x86, 0xe5, 0x0c, 0xa4, 0x7f,
	0x60, 0xb7, 0x60, 0x9a, 0xdb, 0x7f, 0xac, 0xf4, 0xab, 0x58, 0x7d, 0x4f, 0x7b, 0xc7, 0x55, 0xf5,
	0x23, 0xdd, 0x3d, 0x0e, 0xc8, 0x8a, 0x5a, 0x1c, 0xca, 0x4d, 0x35, 0x2d, 0xf0, 0xa4, 0x53, 0xfe,
	0xec, 0x34, 0xb3, 0x8a, 0xda, 0x59, 0xd4, 0xcc, 0xa6, 0xc5, 0x2c, 0x6a, 0x7e, 0xa5, 0x7b, 0xa1,
	0x51, 0x96, 0x15, 0xb5, 0x65, 0x19, 0x55, 0x65, 0x39, 0x2d, 0xca, 0x32, 0xea, 0x16, 0x5b, 0x3e,
	0xec, 0x8b, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x3f, 0x5b, 0x38, 0x7e, 0x03, 0x00, 0x00,
}
