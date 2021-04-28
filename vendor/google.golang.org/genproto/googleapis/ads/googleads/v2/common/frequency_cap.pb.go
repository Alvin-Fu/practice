// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/common/frequency_cap.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// A rule specifying the maximum number of times an ad (or some set of ads) can
// be shown to a user over a particular time period.
type FrequencyCapEntry struct {
	// The key of a particular frequency cap. There can be no more
	// than one frequency cap with the same key.
	Key *FrequencyCapKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Maximum number of events allowed during the time range by this cap.
	Cap                  *wrappers.Int32Value `protobuf:"bytes,2,opt,name=cap,proto3" json:"cap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FrequencyCapEntry) Reset()         { *m = FrequencyCapEntry{} }
func (m *FrequencyCapEntry) String() string { return proto.CompactTextString(m) }
func (*FrequencyCapEntry) ProtoMessage()    {}
func (*FrequencyCapEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_b39111d225f0ecfa, []int{0}
}

func (m *FrequencyCapEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrequencyCapEntry.Unmarshal(m, b)
}
func (m *FrequencyCapEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrequencyCapEntry.Marshal(b, m, deterministic)
}
func (m *FrequencyCapEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrequencyCapEntry.Merge(m, src)
}
func (m *FrequencyCapEntry) XXX_Size() int {
	return xxx_messageInfo_FrequencyCapEntry.Size(m)
}
func (m *FrequencyCapEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_FrequencyCapEntry.DiscardUnknown(m)
}

var xxx_messageInfo_FrequencyCapEntry proto.InternalMessageInfo

func (m *FrequencyCapEntry) GetKey() *FrequencyCapKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *FrequencyCapEntry) GetCap() *wrappers.Int32Value {
	if m != nil {
		return m.Cap
	}
	return nil
}

// A group of fields used as keys for a frequency cap.
// There can be no more than one frequency cap with the same key.
type FrequencyCapKey struct {
	// The level on which the cap is to be applied (e.g. ad group ad, ad group).
	// The cap is applied to all the entities of this level.
	Level enums.FrequencyCapLevelEnum_FrequencyCapLevel `protobuf:"varint,1,opt,name=level,proto3,enum=google.ads.googleads.v2.enums.FrequencyCapLevelEnum_FrequencyCapLevel" json:"level,omitempty"`
	// The type of event that the cap applies to (e.g. impression).
	EventType enums.FrequencyCapEventTypeEnum_FrequencyCapEventType `protobuf:"varint,3,opt,name=event_type,json=eventType,proto3,enum=google.ads.googleads.v2.enums.FrequencyCapEventTypeEnum_FrequencyCapEventType" json:"event_type,omitempty"`
	// Unit of time the cap is defined at (e.g. day, week).
	TimeUnit enums.FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit `protobuf:"varint,2,opt,name=time_unit,json=timeUnit,proto3,enum=google.ads.googleads.v2.enums.FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit" json:"time_unit,omitempty"`
	// Number of time units the cap lasts.
	TimeLength           *wrappers.Int32Value `protobuf:"bytes,4,opt,name=time_length,json=timeLength,proto3" json:"time_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FrequencyCapKey) Reset()         { *m = FrequencyCapKey{} }
func (m *FrequencyCapKey) String() string { return proto.CompactTextString(m) }
func (*FrequencyCapKey) ProtoMessage()    {}
func (*FrequencyCapKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_b39111d225f0ecfa, []int{1}
}

func (m *FrequencyCapKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrequencyCapKey.Unmarshal(m, b)
}
func (m *FrequencyCapKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrequencyCapKey.Marshal(b, m, deterministic)
}
func (m *FrequencyCapKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrequencyCapKey.Merge(m, src)
}
func (m *FrequencyCapKey) XXX_Size() int {
	return xxx_messageInfo_FrequencyCapKey.Size(m)
}
func (m *FrequencyCapKey) XXX_DiscardUnknown() {
	xxx_messageInfo_FrequencyCapKey.DiscardUnknown(m)
}

var xxx_messageInfo_FrequencyCapKey proto.InternalMessageInfo

func (m *FrequencyCapKey) GetLevel() enums.FrequencyCapLevelEnum_FrequencyCapLevel {
	if m != nil {
		return m.Level
	}
	return enums.FrequencyCapLevelEnum_UNSPECIFIED
}

func (m *FrequencyCapKey) GetEventType() enums.FrequencyCapEventTypeEnum_FrequencyCapEventType {
	if m != nil {
		return m.EventType
	}
	return enums.FrequencyCapEventTypeEnum_UNSPECIFIED
}

func (m *FrequencyCapKey) GetTimeUnit() enums.FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit {
	if m != nil {
		return m.TimeUnit
	}
	return enums.FrequencyCapTimeUnitEnum_UNSPECIFIED
}

func (m *FrequencyCapKey) GetTimeLength() *wrappers.Int32Value {
	if m != nil {
		return m.TimeLength
	}
	return nil
}

func init() {
	proto.RegisterType((*FrequencyCapEntry)(nil), "google.ads.googleads.v2.common.FrequencyCapEntry")
	proto.RegisterType((*FrequencyCapKey)(nil), "google.ads.googleads.v2.common.FrequencyCapKey")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/common/frequency_cap.proto", fileDescriptor_b39111d225f0ecfa)
}

var fileDescriptor_b39111d225f0ecfa = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x6a, 0xd4, 0x40,
	0x18, 0x27, 0x1b, 0x15, 0x3b, 0x85, 0x8a, 0x39, 0x2d, 0x55, 0x8a, 0xec, 0xc9, 0x8b, 0x33, 0x30,
	0x3d, 0x08, 0x69, 0x2f, 0x69, 0xdd, 0x16, 0x71, 0x91, 0x12, 0xea, 0x1e, 0x24, 0xb0, 0x4c, 0x93,
	0xaf, 0x31, 0x98, 0xcc, 0x8c, 0xc9, 0x64, 0x25, 0x0f, 0x20, 0xbe, 0x87, 0x47, 0x1f, 0xc5, 0x47,
	0x11, 0x1f, 0x42, 0x66, 0x26, 0x13, 0x5d, 0x97, 0xb5, 0xe4, 0x94, 0x5f, 0xe6, 0xfb, 0xfd, 0xf9,
	0xf2, 0x7d, 0x13, 0x44, 0x73, 0x21, 0xf2, 0x12, 0x08, 0xcb, 0x1a, 0x62, 0xa1, 0x46, 0x6b, 0x4a,
	0x52, 0x51, 0x55, 0x82, 0x93, 0xdb, 0x1a, 0x3e, 0xb5, 0xc0, 0xd3, 0x6e, 0x95, 0x32, 0x89, 0x65,
	0x2d, 0x94, 0x08, 0x8e, 0x2c, 0x11, 0xb3, 0xac, 0xc1, 0x83, 0x06, 0xaf, 0x29, 0xb6, 0x9a, 0xc3,
	0xd3, 0x5d, 0x9e, 0xc0, 0xdb, 0xaa, 0xd9, 0xb4, 0x5c, 0xc1, 0x1a, 0xb8, 0x5a, 0xa9, 0x4e, 0x82,
	0x75, 0x3f, 0x7c, 0x39, 0x46, 0x5d, 0xc2, 0x1a, 0xca, 0x5e, 0x78, 0x32, 0x46, 0xa8, 0x8a, 0x0a,
	0x56, 0x2d, 0x2f, 0x54, 0x2f, 0xee, 0xbf, 0x89, 0x98, 0xb7, 0x9b, 0xf6, 0x96, 0x7c, 0xae, 0x99,
	0x94, 0x50, 0x37, 0x7d, 0xfd, 0xa9, 0x33, 0x97, 0x05, 0x61, 0x9c, 0x0b, 0xc5, 0x54, 0x21, 0x78,
	0x5f, 0x9d, 0x7d, 0xf1, 0xd0, 0xe3, 0x0b, 0xe7, 0x7f, 0xce, 0xe4, 0x9c, 0xab, 0xba, 0x0b, 0x22,
	0xe4, 0x7f, 0x84, 0x6e, 0xea, 0x3d, 0xf3, 0x9e, 0xef, 0x53, 0x82, 0xff, 0x3f, 0x35, 0xfc, 0xb7,
	0xfe, 0x0d, 0x74, 0xb1, 0xd6, 0x06, 0x2f, 0x90, 0x9f, 0x32, 0x39, 0x9d, 0x18, 0x8b, 0x27, 0xce,
	0xc2, 0x35, 0x89, 0x5f, 0x73, 0x75, 0x4c, 0x97, 0xac, 0x6c, 0x21, 0xd6, 0xbc, 0xd9, 0x57, 0x1f,
	0x3d, 0xfa, 0xc7, 0x27, 0x48, 0xd0, 0x7d, 0x33, 0x25, 0xd3, 0xc7, 0x01, 0xbd, 0xd8, 0xd9, 0x87,
	0x19, 0xd3, 0x46, 0x1b, 0x0b, 0xad, 0x9b, 0xf3, 0xb6, 0xda, 0x3e, 0x8d, 0xad, 0x69, 0x50, 0x21,
	0xf4, 0x67, 0x83, 0x53, 0xdf, 0x44, 0xbc, 0x1d, 0x11, 0x31, 0xd7, 0xe2, 0xeb, 0x4e, 0xc2, 0x56,
	0xcc, 0x50, 0x89, 0xf7, 0xc0, 0xc1, 0xa0, 0x40, 0x7b, 0xc3, 0xe6, 0xcc, 0x54, 0x0e, 0xe8, 0x62,
	0x44, 0xda, 0x75, 0x51, 0xc1, 0x3b, 0x5e, 0xa8, 0xad, 0x30, 0x57, 0x88, 0x1f, 0xaa, 0x1e, 0x05,
	0xa7, 0x68, 0xdf, 0x44, 0x95, 0xc0, 0x73, 0xf5, 0x61, 0x7a, 0xef, 0xee, 0x15, 0x20, 0xcd, 0x5f,
	0x18, 0xfa, 0xd9, 0x2f, 0x0f, 0xcd, 0x52, 0x51, 0xdd, 0xb1, 0xf4, 0xb3, 0x8d, 0x5b, 0x73, 0xa5,
	0x3d, 0xaf, 0xbc, 0xf7, 0xaf, 0x7a, 0x51, 0x2e, 0x4a, 0xc6, 0x73, 0x2c, 0xea, 0x9c, 0xe4, 0xc0,
	0x4d, 0xa2, 0xbb, 0xd8, 0xb2, 0x68, 0x76, 0xfd, 0xb2, 0x27, 0xf6, 0xf1, 0x6d, 0xe2, 0x5f, 0x46,
	0xd1, 0xf7, 0xc9, 0xd1, 0xa5, 0x35, 0x8b, 0xb2, 0x06, 0x5b, 0xa8, 0xd1, 0x92, 0xe2, 0x73, 0x43,
	0xfb, 0xe1, 0x08, 0x49, 0x94, 0x35, 0xc9, 0x40, 0x48, 0x96, 0x34, 0xb1, 0x84, 0x9f, 0x93, 0x99,
	0x3d, 0x0d, 0xc3, 0x28, 0x6b, 0xc2, 0x70, 0xa0, 0x84, 0xe1, 0x92, 0x86, 0xa1, 0x25, 0xdd, 0x3c,
	0x30, 0xdd, 0x1d, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x19, 0x50, 0xd8, 0x4f, 0x04, 0x00,
	0x00,
}
