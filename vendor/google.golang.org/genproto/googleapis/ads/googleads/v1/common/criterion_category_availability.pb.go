// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/criterion_category_availability.proto

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

// Information of category availability, per advertising channel.
type CriterionCategoryAvailability struct {
	// Channel types and subtypes that are available to the category.
	Channel *CriterionCategoryChannelAvailability `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	// Locales that are available to the category for the channel.
	Locale               []*CriterionCategoryLocaleAvailability `protobuf:"bytes,2,rep,name=locale,proto3" json:"locale,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *CriterionCategoryAvailability) Reset()         { *m = CriterionCategoryAvailability{} }
func (m *CriterionCategoryAvailability) String() string { return proto.CompactTextString(m) }
func (*CriterionCategoryAvailability) ProtoMessage()    {}
func (*CriterionCategoryAvailability) Descriptor() ([]byte, []int) {
	return fileDescriptor_241d9930057f9832, []int{0}
}

func (m *CriterionCategoryAvailability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriterionCategoryAvailability.Unmarshal(m, b)
}
func (m *CriterionCategoryAvailability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriterionCategoryAvailability.Marshal(b, m, deterministic)
}
func (m *CriterionCategoryAvailability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriterionCategoryAvailability.Merge(m, src)
}
func (m *CriterionCategoryAvailability) XXX_Size() int {
	return xxx_messageInfo_CriterionCategoryAvailability.Size(m)
}
func (m *CriterionCategoryAvailability) XXX_DiscardUnknown() {
	xxx_messageInfo_CriterionCategoryAvailability.DiscardUnknown(m)
}

var xxx_messageInfo_CriterionCategoryAvailability proto.InternalMessageInfo

func (m *CriterionCategoryAvailability) GetChannel() *CriterionCategoryChannelAvailability {
	if m != nil {
		return m.Channel
	}
	return nil
}

func (m *CriterionCategoryAvailability) GetLocale() []*CriterionCategoryLocaleAvailability {
	if m != nil {
		return m.Locale
	}
	return nil
}

// Information of advertising channel type and subtypes a category is available
// in.
type CriterionCategoryChannelAvailability struct {
	// Format of the channel availability. Can be ALL_CHANNELS (the rest of the
	// fields will not be set), CHANNEL_TYPE (only advertising_channel_type type
	// will be set, the category is available to all sub types under it) or
	// CHANNEL_TYPE_AND_SUBTYPES (advertising_channel_type,
	// advertising_channel_sub_type, and include_default_channel_sub_type will all
	// be set).
	AvailabilityMode enums.CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode `protobuf:"varint,1,opt,name=availability_mode,json=availabilityMode,proto3,enum=google.ads.googleads.v1.enums.CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode" json:"availability_mode,omitempty"`
	// Channel type the category is available to.
	AdvertisingChannelType enums.AdvertisingChannelTypeEnum_AdvertisingChannelType `protobuf:"varint,2,opt,name=advertising_channel_type,json=advertisingChannelType,proto3,enum=google.ads.googleads.v1.enums.AdvertisingChannelTypeEnum_AdvertisingChannelType" json:"advertising_channel_type,omitempty"`
	// Channel subtypes under the channel type the category is available to.
	AdvertisingChannelSubType []enums.AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType `protobuf:"varint,3,rep,packed,name=advertising_channel_sub_type,json=advertisingChannelSubType,proto3,enum=google.ads.googleads.v1.enums.AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType" json:"advertising_channel_sub_type,omitempty"`
	// Whether default channel sub type is included. For example,
	// advertising_channel_type being DISPLAY and include_default_channel_sub_type
	// being false means that the default display campaign where channel sub type
	// is not set is not included in this availability configuration.
	IncludeDefaultChannelSubType *wrappers.BoolValue `protobuf:"bytes,4,opt,name=include_default_channel_sub_type,json=includeDefaultChannelSubType,proto3" json:"include_default_channel_sub_type,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}            `json:"-"`
	XXX_unrecognized             []byte              `json:"-"`
	XXX_sizecache                int32               `json:"-"`
}

func (m *CriterionCategoryChannelAvailability) Reset()         { *m = CriterionCategoryChannelAvailability{} }
func (m *CriterionCategoryChannelAvailability) String() string { return proto.CompactTextString(m) }
func (*CriterionCategoryChannelAvailability) ProtoMessage()    {}
func (*CriterionCategoryChannelAvailability) Descriptor() ([]byte, []int) {
	return fileDescriptor_241d9930057f9832, []int{1}
}

func (m *CriterionCategoryChannelAvailability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriterionCategoryChannelAvailability.Unmarshal(m, b)
}
func (m *CriterionCategoryChannelAvailability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriterionCategoryChannelAvailability.Marshal(b, m, deterministic)
}
func (m *CriterionCategoryChannelAvailability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriterionCategoryChannelAvailability.Merge(m, src)
}
func (m *CriterionCategoryChannelAvailability) XXX_Size() int {
	return xxx_messageInfo_CriterionCategoryChannelAvailability.Size(m)
}
func (m *CriterionCategoryChannelAvailability) XXX_DiscardUnknown() {
	xxx_messageInfo_CriterionCategoryChannelAvailability.DiscardUnknown(m)
}

var xxx_messageInfo_CriterionCategoryChannelAvailability proto.InternalMessageInfo

func (m *CriterionCategoryChannelAvailability) GetAvailabilityMode() enums.CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode {
	if m != nil {
		return m.AvailabilityMode
	}
	return enums.CriterionCategoryChannelAvailabilityModeEnum_UNSPECIFIED
}

func (m *CriterionCategoryChannelAvailability) GetAdvertisingChannelType() enums.AdvertisingChannelTypeEnum_AdvertisingChannelType {
	if m != nil {
		return m.AdvertisingChannelType
	}
	return enums.AdvertisingChannelTypeEnum_UNSPECIFIED
}

func (m *CriterionCategoryChannelAvailability) GetAdvertisingChannelSubType() []enums.AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType {
	if m != nil {
		return m.AdvertisingChannelSubType
	}
	return nil
}

func (m *CriterionCategoryChannelAvailability) GetIncludeDefaultChannelSubType() *wrappers.BoolValue {
	if m != nil {
		return m.IncludeDefaultChannelSubType
	}
	return nil
}

// Information about which locales a category is available in.
type CriterionCategoryLocaleAvailability struct {
	// Format of the locale availability. Can be LAUNCHED_TO_ALL (both country and
	// language will be empty), COUNTRY (only country will be set), LANGUAGE (only
	// language wil be set), COUNTRY_AND_LANGUAGE (both country and language will
	// be set).
	AvailabilityMode enums.CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode `protobuf:"varint,1,opt,name=availability_mode,json=availabilityMode,proto3,enum=google.ads.googleads.v1.enums.CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode" json:"availability_mode,omitempty"`
	// Code of the country.
	CountryCode *wrappers.StringValue `protobuf:"bytes,2,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	// Code of the language.
	LanguageCode         *wrappers.StringValue `protobuf:"bytes,3,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CriterionCategoryLocaleAvailability) Reset()         { *m = CriterionCategoryLocaleAvailability{} }
func (m *CriterionCategoryLocaleAvailability) String() string { return proto.CompactTextString(m) }
func (*CriterionCategoryLocaleAvailability) ProtoMessage()    {}
func (*CriterionCategoryLocaleAvailability) Descriptor() ([]byte, []int) {
	return fileDescriptor_241d9930057f9832, []int{2}
}

func (m *CriterionCategoryLocaleAvailability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriterionCategoryLocaleAvailability.Unmarshal(m, b)
}
func (m *CriterionCategoryLocaleAvailability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriterionCategoryLocaleAvailability.Marshal(b, m, deterministic)
}
func (m *CriterionCategoryLocaleAvailability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriterionCategoryLocaleAvailability.Merge(m, src)
}
func (m *CriterionCategoryLocaleAvailability) XXX_Size() int {
	return xxx_messageInfo_CriterionCategoryLocaleAvailability.Size(m)
}
func (m *CriterionCategoryLocaleAvailability) XXX_DiscardUnknown() {
	xxx_messageInfo_CriterionCategoryLocaleAvailability.DiscardUnknown(m)
}

var xxx_messageInfo_CriterionCategoryLocaleAvailability proto.InternalMessageInfo

func (m *CriterionCategoryLocaleAvailability) GetAvailabilityMode() enums.CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode {
	if m != nil {
		return m.AvailabilityMode
	}
	return enums.CriterionCategoryLocaleAvailabilityModeEnum_UNSPECIFIED
}

func (m *CriterionCategoryLocaleAvailability) GetCountryCode() *wrappers.StringValue {
	if m != nil {
		return m.CountryCode
	}
	return nil
}

func (m *CriterionCategoryLocaleAvailability) GetLanguageCode() *wrappers.StringValue {
	if m != nil {
		return m.LanguageCode
	}
	return nil
}

func init() {
	proto.RegisterType((*CriterionCategoryAvailability)(nil), "google.ads.googleads.v1.common.CriterionCategoryAvailability")
	proto.RegisterType((*CriterionCategoryChannelAvailability)(nil), "google.ads.googleads.v1.common.CriterionCategoryChannelAvailability")
	proto.RegisterType((*CriterionCategoryLocaleAvailability)(nil), "google.ads.googleads.v1.common.CriterionCategoryLocaleAvailability")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/criterion_category_availability.proto", fileDescriptor_241d9930057f9832)
}

var fileDescriptor_241d9930057f9832 = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x95, 0x74, 0xb5, 0x48, 0xee, 0xb2, 0x82, 0x1c, 0x50, 0xa9, 0xca, 0xaa, 0x0a, 0x1c,
	0x7a, 0x72, 0xd4, 0x72, 0x0b, 0x48, 0x90, 0xa6, 0x68, 0x2f, 0x20, 0xaa, 0x2e, 0xea, 0x01, 0x2a,
	0x22, 0x27, 0xf1, 0x06, 0x4b, 0x8e, 0x1d, 0x25, 0x4e, 0x51, 0x5f, 0x81, 0x17, 0xe0, 0x88, 0xe0,
	0xc8, 0xa3, 0xc0, 0x99, 0x1b, 0x4f, 0xc0, 0x3b, 0x20, 0xa1, 0xc4, 0x4e, 0xe9, 0x6e, 0x9b, 0x6c,
	0x7b, 0xaa, 0xeb, 0x99, 0xf9, 0xe6, 0x9f, 0xc9, 0xd8, 0x06, 0x93, 0x88, 0xf3, 0x88, 0x62, 0x0b,
	0x85, 0x99, 0x25, 0x97, 0xc5, 0x6a, 0x39, 0xb4, 0x02, 0x1e, 0xc7, 0x9c, 0x59, 0x41, 0x4a, 0x04,
	0x4e, 0x09, 0x67, 0x5e, 0x80, 0x04, 0x8e, 0x78, 0xba, 0xf2, 0xd0, 0x12, 0x11, 0x8a, 0x7c, 0x42,
	0x89, 0x58, 0xc1, 0x24, 0xe5, 0x82, 0x1b, 0x67, 0x32, 0x14, 0xa2, 0x30, 0x83, 0x6b, 0x0a, 0x5c,
	0x0e, 0xa1, 0xa4, 0x74, 0x9f, 0xd7, 0x65, 0xc1, 0x2c, 0x8f, 0x33, 0x0b, 0x85, 0x4b, 0x9c, 0x0a,
	0x92, 0x11, 0x16, 0x79, 0xc1, 0x07, 0xc4, 0x18, 0xa6, 0x5e, 0x96, 0xfb, 0x9e, 0x58, 0x25, 0x58,
	0x66, 0xe8, 0x3e, 0x3d, 0x9c, 0xb0, 0x11, 0x3d, 0x6d, 0x8e, 0xde, 0x51, 0x64, 0x05, 0xd9, 0x2c,
	0xd6, 0x8b, 0x79, 0x58, 0x11, 0x5f, 0x1f, 0x4c, 0xa4, 0x3c, 0x40, 0x14, 0xd7, 0x02, 0x7b, 0x15,
	0x30, 0x21, 0x16, 0x62, 0x8c, 0x0b, 0x24, 0x08, 0x67, 0x99, 0xb2, 0xaa, 0x06, 0x5b, 0xe5, 0x3f,
	0x3f, 0xbf, 0xb4, 0x3e, 0xa6, 0x28, 0x49, 0x70, 0xaa, 0xec, 0xe6, 0x6f, 0x0d, 0x3c, 0x70, 0xab,
	0x9c, 0xae, 0x4a, 0xe9, 0x6c, 0xa4, 0x32, 0xde, 0x83, 0x5b, 0xaa, 0xa6, 0x8e, 0xd6, 0xd7, 0x06,
	0xed, 0xd1, 0x04, 0x36, 0x7f, 0x34, 0xb8, 0xc5, 0x73, 0x65, 0xfc, 0x26, 0x76, 0x56, 0x41, 0x8d,
	0x77, 0xe0, 0x58, 0x56, 0xd8, 0xd1, 0xfb, 0xad, 0x41, 0x7b, 0xe4, 0x1e, 0x8c, 0x7f, 0x59, 0x86,
	0x5f, 0xa1, 0x2b, 0xa4, 0xf9, 0xeb, 0x08, 0x3c, 0xda, 0x47, 0x8e, 0xf1, 0x55, 0x03, 0x77, 0xb7,
	0x3a, 0x5c, 0x16, 0x7c, 0x3a, 0x12, 0xb5, 0x8a, 0xca, 0x6f, 0xb6, 0x57, 0xbd, 0xaf, 0x78, 0x88,
	0x5f, 0xb0, 0x3c, 0xde, 0xdb, 0x79, 0x76, 0x07, 0x5d, 0xdb, 0x31, 0x3e, 0x69, 0xa0, 0x53, 0x37,
	0xaf, 0x1d, 0xbd, 0x94, 0x3a, 0xbd, 0x41, 0xaa, 0xf3, 0x3f, 0x5c, 0xe5, 0x7d, 0xb3, 0x4a, 0xa4,
	0xb0, 0xdd, 0xa6, 0xd9, 0x3d, 0xb4, 0x73, 0xdf, 0xf8, 0xac, 0x81, 0x5e, 0xd3, 0xf1, 0xeb, 0xb4,
	0xfa, 0xad, 0xc1, 0xe9, 0x68, 0x7e, 0xb0, 0xa0, 0x8b, 0xdc, 0x6f, 0xd0, 0xa4, 0xac, 0xb3, 0xfb,
	0xa8, 0xce, 0x64, 0xf8, 0xa0, 0x4f, 0x58, 0x40, 0xf3, 0x10, 0x7b, 0x21, 0xbe, 0x44, 0x39, 0x15,
	0xdb, 0xe2, 0x8e, 0xca, 0x49, 0xee, 0x56, 0xe2, 0xaa, 0xd3, 0x01, 0xc7, 0x9c, 0xd3, 0x39, 0xa2,
	0x39, 0x9e, 0xf5, 0x14, 0x63, 0x22, 0x11, 0x57, 0x73, 0x98, 0x3f, 0x75, 0xf0, 0x70, 0x8f, 0x39,
	0x34, 0xbe, 0x34, 0x8c, 0x55, 0x7a, 0xe8, 0x58, 0x6d, 0xf3, 0xeb, 0xa7, 0x6a, 0xb7, 0xef, 0x8e,
	0xa1, 0x7a, 0x06, 0x4e, 0x02, 0x9e, 0x33, 0x51, 0x5c, 0x5d, 0x85, 0x36, 0xbd, 0xec, 0x4c, 0x6f,
	0xab, 0x33, 0x17, 0x22, 0x25, 0x2c, 0x92, 0xbd, 0x69, 0xab, 0x08, 0xb7, 0x00, 0x38, 0xe0, 0x36,
	0x45, 0x2c, 0xca, 0x51, 0x84, 0x25, 0xa1, 0xb5, 0x07, 0xe1, 0xa4, 0x0a, 0x29, 0x10, 0xe3, 0xbf,
	0x1a, 0x30, 0x03, 0x1e, 0xdf, 0x70, 0xf0, 0xc7, 0x66, 0xe3, 0x45, 0x35, 0x2d, 0xf2, 0x4c, 0xb5,
	0xb7, 0xea, 0x61, 0x82, 0x11, 0x2f, 0x72, 0x40, 0x9e, 0x46, 0x56, 0x84, 0x59, 0xa9, 0xa2, 0xba,
	0x70, 0x13, 0x92, 0xd5, 0xbd, 0x5b, 0x4f, 0xe4, 0xcf, 0x37, 0xbd, 0x75, 0xee, 0x38, 0xdf, 0xf5,
	0xb3, 0x73, 0x09, 0x73, 0xc2, 0x0c, 0xca, 0x65, 0xb1, 0x9a, 0x0f, 0xa1, 0x5b, 0xba, 0xfd, 0xa8,
	0x1c, 0x16, 0x4e, 0x98, 0x2d, 0xd6, 0x0e, 0x8b, 0xf9, 0x70, 0x21, 0x1d, 0xfe, 0xe8, 0xa6, 0xdc,
	0xb5, 0x6d, 0x27, 0xcc, 0x6c, 0x7b, 0xed, 0x62, 0xdb, 0xf3, 0xa1, 0x6d, 0x4b, 0x27, 0xff, 0xb8,
	0x54, 0xf7, 0xf8, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0x74, 0x4f, 0x48, 0x54, 0x07, 0x00,
	0x00,
}
