// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/campaign_serving_status.proto

package enums

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

// Possible serving statuses of a campaign.
type CampaignServingStatusEnum_CampaignServingStatus int32

const (
	// No value has been specified.
	CampaignServingStatusEnum_UNSPECIFIED CampaignServingStatusEnum_CampaignServingStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	CampaignServingStatusEnum_UNKNOWN CampaignServingStatusEnum_CampaignServingStatus = 1
	// Serving.
	CampaignServingStatusEnum_SERVING CampaignServingStatusEnum_CampaignServingStatus = 2
	// None.
	CampaignServingStatusEnum_NONE CampaignServingStatusEnum_CampaignServingStatus = 3
	// Ended.
	CampaignServingStatusEnum_ENDED CampaignServingStatusEnum_CampaignServingStatus = 4
	// Pending.
	CampaignServingStatusEnum_PENDING CampaignServingStatusEnum_CampaignServingStatus = 5
	// Suspended.
	CampaignServingStatusEnum_SUSPENDED CampaignServingStatusEnum_CampaignServingStatus = 6
)

var CampaignServingStatusEnum_CampaignServingStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SERVING",
	3: "NONE",
	4: "ENDED",
	5: "PENDING",
	6: "SUSPENDED",
}

var CampaignServingStatusEnum_CampaignServingStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"SERVING":     2,
	"NONE":        3,
	"ENDED":       4,
	"PENDING":     5,
	"SUSPENDED":   6,
}

func (x CampaignServingStatusEnum_CampaignServingStatus) String() string {
	return proto.EnumName(CampaignServingStatusEnum_CampaignServingStatus_name, int32(x))
}

func (CampaignServingStatusEnum_CampaignServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_40e65b6fe8387390, []int{0, 0}
}

// Message describing Campaign serving statuses.
type CampaignServingStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignServingStatusEnum) Reset()         { *m = CampaignServingStatusEnum{} }
func (m *CampaignServingStatusEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignServingStatusEnum) ProtoMessage()    {}
func (*CampaignServingStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_40e65b6fe8387390, []int{0}
}

func (m *CampaignServingStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignServingStatusEnum.Unmarshal(m, b)
}
func (m *CampaignServingStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignServingStatusEnum.Marshal(b, m, deterministic)
}
func (m *CampaignServingStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignServingStatusEnum.Merge(m, src)
}
func (m *CampaignServingStatusEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignServingStatusEnum.Size(m)
}
func (m *CampaignServingStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignServingStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignServingStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.CampaignServingStatusEnum_CampaignServingStatus", CampaignServingStatusEnum_CampaignServingStatus_name, CampaignServingStatusEnum_CampaignServingStatus_value)
	proto.RegisterType((*CampaignServingStatusEnum)(nil), "google.ads.googleads.v1.enums.CampaignServingStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/campaign_serving_status.proto", fileDescriptor_40e65b6fe8387390)
}

var fileDescriptor_40e65b6fe8387390 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0xb6, 0xdd, 0x1f, 0x5d, 0x86, 0x58, 0x0a, 0x1e, 0x1c, 0xee, 0xb0, 0x3d, 0x40, 0x4a, 0xf1,
	0x96, 0x9d, 0xba, 0x35, 0x8e, 0x21, 0x64, 0xc5, 0xb2, 0x0a, 0x52, 0x18, 0x71, 0x2d, 0xa1, 0xb0,
	0x26, 0x65, 0xe9, 0xf6, 0x1c, 0x3e, 0x83, 0x47, 0x1f, 0xc5, 0x47, 0xf1, 0xe2, 0x2b, 0x48, 0x92,
	0xad, 0xa7, 0xe9, 0xa5, 0x7c, 0xcd, 0xf7, 0x87, 0xdf, 0xf7, 0x81, 0x09, 0x13, 0x82, 0x6d, 0x73,
	0x8f, 0x66, 0xd2, 0x33, 0x50, 0xa1, 0x83, 0xef, 0xe5, 0x7c, 0x5f, 0x4a, 0x6f, 0x43, 0xcb, 0x8a,
	0x16, 0x8c, 0xaf, 0x65, 0xbe, 0x3b, 0x14, 0x9c, 0xad, 0x65, 0x4d, 0xeb, 0xbd, 0x84, 0xd5, 0x4e,
	0xd4, 0xc2, 0x1d, 0x1a, 0x07, 0xa4, 0x99, 0x84, 0x8d, 0x19, 0x1e, 0x7c, 0xa8, 0xcd, 0x83, 0xfb,
	0x53, 0x76, 0x55, 0x78, 0x94, 0x73, 0x51, 0xd3, 0xba, 0x10, 0xfc, 0x68, 0x1e, 0xbf, 0x5b, 0xe0,
	0x6e, 0x76, 0x8c, 0x8f, 0x4d, 0x7a, 0xac, 0xc3, 0x31, 0xdf, 0x97, 0x63, 0x09, 0x6e, 0xcf, 0x92,
	0xee, 0x0d, 0xe8, 0xaf, 0x48, 0x1c, 0xe1, 0xd9, 0xe2, 0x71, 0x81, 0x43, 0xe7, 0xc2, 0xed, 0x83,
	0xcb, 0x15, 0x79, 0x22, 0xcb, 0x17, 0xe2, 0x58, 0xea, 0x27, 0xc6, 0xcf, 0xc9, 0x82, 0xcc, 0x1d,
	0xdb, 0xbd, 0x02, 0x6d, 0xb2, 0x24, 0xd8, 0x69, 0xb9, 0x3d, 0xd0, 0xc1, 0x24, 0xc4, 0xa1, 0xd3,
	0x56, 0x8a, 0x08, 0x93, 0x50, 0x29, 0x3a, 0xee, 0x35, 0xe8, 0xc5, 0xab, 0x38, 0x32, 0x5c, 0x77,
	0xfa, 0x63, 0x81, 0xd1, 0x46, 0x94, 0xf0, 0xdf, 0x5a, 0xd3, 0xc1, 0xd9, 0xc3, 0x22, 0x55, 0x2a,
	0xb2, 0x5e, 0xa7, 0x47, 0x33, 0x13, 0x5b, 0xca, 0x19, 0x14, 0x3b, 0xe6, 0xb1, 0x9c, 0xeb, 0xca,
	0xa7, 0x81, 0xab, 0x42, 0xfe, 0xb1, 0xf7, 0x44, 0x7f, 0x3f, 0xec, 0xd6, 0x3c, 0x08, 0x3e, 0xed,
	0xe1, 0xdc, 0x44, 0x05, 0x99, 0x84, 0x06, 0x2a, 0x94, 0xf8, 0x50, 0x2d, 0x24, 0xbf, 0x4e, 0x7c,
	0x1a, 0x64, 0x32, 0x6d, 0xf8, 0x34, 0xf1, 0x53, 0xcd, 0x7f, 0xdb, 0x23, 0xf3, 0x88, 0x50, 0x90,
	0x49, 0x84, 0x1a, 0x05, 0x42, 0x89, 0x8f, 0x90, 0xd6, 0xbc, 0x75, 0xf5, 0x61, 0x0f, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x98, 0x3d, 0x25, 0x2d, 0x07, 0x02, 0x00, 0x00,
}
