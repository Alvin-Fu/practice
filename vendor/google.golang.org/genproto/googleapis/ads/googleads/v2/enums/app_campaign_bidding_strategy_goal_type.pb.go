// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/app_campaign_bidding_strategy_goal_type.proto

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

// Goal type of App campaign BiddingStrategy.
type AppCampaignBiddingStrategyGoalTypeEnum_AppCampaignBiddingStrategyGoalType int32

const (
	// Not specified.
	AppCampaignBiddingStrategyGoalTypeEnum_UNSPECIFIED AppCampaignBiddingStrategyGoalTypeEnum_AppCampaignBiddingStrategyGoalType = 0
	// Used for return value only. Represents value unknown in this version.
	AppCampaignBiddingStrategyGoalTypeEnum_UNKNOWN AppCampaignBiddingStrategyGoalTypeEnum_AppCampaignBiddingStrategyGoalType = 1
	// Aim to maximize the number of app installs. The cpa bid is the
	// target cost per install.
	AppCampaignBiddingStrategyGoalTypeEnum_OPTIMIZE_INSTALLS_TARGET_INSTALL_COST AppCampaignBiddingStrategyGoalTypeEnum_AppCampaignBiddingStrategyGoalType = 2
	// Aim to maximize the long term number of selected in-app conversions from
	// app installs. The cpa bid is the target cost per install.
	AppCampaignBiddingStrategyGoalTypeEnum_OPTIMIZE_IN_APP_CONVERSIONS_TARGET_INSTALL_COST AppCampaignBiddingStrategyGoalTypeEnum_AppCampaignBiddingStrategyGoalType = 3
	// Aim to maximize the long term number of selected in-app conversions from
	// app installs. The cpa bid is the target cost per in-app conversion. Note
	// that the actual cpa may seem higher than the target cpa at first, since
	// the long term conversions haven’t happened yet.
	AppCampaignBiddingStrategyGoalTypeEnum_OPTIMIZE_IN_APP_CONVERSIONS_TARGET_CONVERSION_COST AppCampaignBiddingStrategyGoalTypeEnum_AppCampaignBiddingStrategyGoalType = 4
	// Aim to maximize all conversions' value, i.e. install + selected in-app
	// conversions while achieving or exceeding target return on advertising
	// spend.
	AppCampaignBiddingStrategyGoalTypeEnum_OPTIMIZE_RETURN_ON_ADVERTISING_SPEND AppCampaignBiddingStrategyGoalTypeEnum_AppCampaignBiddingStrategyGoalType = 5
)

var AppCampaignBiddingStrategyGoalTypeEnum_AppCampaignBiddingStrategyGoalType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "OPTIMIZE_INSTALLS_TARGET_INSTALL_COST",
	3: "OPTIMIZE_IN_APP_CONVERSIONS_TARGET_INSTALL_COST",
	4: "OPTIMIZE_IN_APP_CONVERSIONS_TARGET_CONVERSION_COST",
	5: "OPTIMIZE_RETURN_ON_ADVERTISING_SPEND",
}

var AppCampaignBiddingStrategyGoalTypeEnum_AppCampaignBiddingStrategyGoalType_value = map[string]int32{
	"UNSPECIFIED":                           0,
	"UNKNOWN":                               1,
	"OPTIMIZE_INSTALLS_TARGET_INSTALL_COST": 2,
	"OPTIMIZE_IN_APP_CONVERSIONS_TARGET_INSTALL_COST":    3,
	"OPTIMIZE_IN_APP_CONVERSIONS_TARGET_CONVERSION_COST": 4,
	"OPTIMIZE_RETURN_ON_ADVERTISING_SPEND":               5,
}

func (x AppCampaignBiddingStrategyGoalTypeEnum_AppCampaignBiddingStrategyGoalType) String() string {
	return proto.EnumName(AppCampaignBiddingStrategyGoalTypeEnum_AppCampaignBiddingStrategyGoalType_name, int32(x))
}

func (AppCampaignBiddingStrategyGoalTypeEnum_AppCampaignBiddingStrategyGoalType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77777a81d427ac7f, []int{0, 0}
}

// Container for enum describing goal towards which the bidding strategy of an
// app campaign should optimize for.
type AppCampaignBiddingStrategyGoalTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppCampaignBiddingStrategyGoalTypeEnum) Reset() {
	*m = AppCampaignBiddingStrategyGoalTypeEnum{}
}
func (m *AppCampaignBiddingStrategyGoalTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AppCampaignBiddingStrategyGoalTypeEnum) ProtoMessage()    {}
func (*AppCampaignBiddingStrategyGoalTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_77777a81d427ac7f, []int{0}
}

func (m *AppCampaignBiddingStrategyGoalTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppCampaignBiddingStrategyGoalTypeEnum.Unmarshal(m, b)
}
func (m *AppCampaignBiddingStrategyGoalTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppCampaignBiddingStrategyGoalTypeEnum.Marshal(b, m, deterministic)
}
func (m *AppCampaignBiddingStrategyGoalTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppCampaignBiddingStrategyGoalTypeEnum.Merge(m, src)
}
func (m *AppCampaignBiddingStrategyGoalTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AppCampaignBiddingStrategyGoalTypeEnum.Size(m)
}
func (m *AppCampaignBiddingStrategyGoalTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AppCampaignBiddingStrategyGoalTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AppCampaignBiddingStrategyGoalTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.AppCampaignBiddingStrategyGoalTypeEnum_AppCampaignBiddingStrategyGoalType", AppCampaignBiddingStrategyGoalTypeEnum_AppCampaignBiddingStrategyGoalType_name, AppCampaignBiddingStrategyGoalTypeEnum_AppCampaignBiddingStrategyGoalType_value)
	proto.RegisterType((*AppCampaignBiddingStrategyGoalTypeEnum)(nil), "google.ads.googleads.v2.enums.AppCampaignBiddingStrategyGoalTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/app_campaign_bidding_strategy_goal_type.proto", fileDescriptor_77777a81d427ac7f)
}

var fileDescriptor_77777a81d427ac7f = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4f, 0x6b, 0xdb, 0x30,
	0x14, 0x5f, 0xdc, 0xfd, 0x01, 0xf5, 0xb0, 0xe0, 0xe3, 0x58, 0x0f, 0x0d, 0xfb, 0x7b, 0x91, 0xc1,
	0x85, 0x1d, 0xb4, 0x93, 0x92, 0x68, 0x46, 0xb4, 0x93, 0x8d, 0xa5, 0x78, 0x50, 0x02, 0x42, 0xad,
	0x8d, 0x30, 0x24, 0x92, 0x88, 0xdc, 0x42, 0xee, 0xfb, 0x24, 0x3b, 0xee, 0xb2, 0xef, 0xb1, 0x8f,
	0xb2, 0xcf, 0xb0, 0xc3, 0x88, 0x95, 0x78, 0x3b, 0x6c, 0x6b, 0x2f, 0xe6, 0xe7, 0xf7, 0x7e, 0x7f,
	0xc4, 0x7b, 0x0f, 0x9c, 0x6b, 0x6b, 0xf5, 0xaa, 0x49, 0x54, 0xed, 0x93, 0x00, 0x77, 0xe8, 0x36,
	0x4d, 0x1a, 0x73, 0xb3, 0xf6, 0x89, 0x72, 0x4e, 0x5e, 0xab, 0xb5, 0x53, 0xad, 0x36, 0xf2, 0xaa,
	0xad, 0xeb, 0xd6, 0x68, 0xe9, 0xbb, 0x8d, 0xea, 0x1a, 0xbd, 0x95, 0xda, 0xaa, 0x95, 0xec, 0xb6,
	0xae, 0x81, 0x6e, 0x63, 0x3b, 0x1b, 0x9f, 0x04, 0x07, 0xa8, 0x6a, 0x0f, 0x07, 0x33, 0x78, 0x9b,
	0xc2, 0xde, 0xec, 0xd9, 0xf3, 0x43, 0x96, 0x6b, 0x13, 0x65, 0x8c, 0xed, 0x54, 0xd7, 0x5a, 0xe3,
	0x83, 0x78, 0xf2, 0x2d, 0x02, 0xaf, 0xb0, 0x73, 0xb3, 0x7d, 0xda, 0x34, 0x84, 0xf1, 0x7d, 0x56,
	0x66, 0xd5, 0x4a, 0x6c, 0x5d, 0x43, 0xcc, 0xcd, 0x7a, 0xf2, 0x39, 0x02, 0x93, 0xbb, 0xa9, 0xf1,
	0x53, 0x70, 0xbc, 0x60, 0xbc, 0x20, 0x33, 0xfa, 0x81, 0x92, 0xf9, 0xf8, 0x41, 0x7c, 0x0c, 0x9e,
	0x2c, 0xd8, 0x39, 0xcb, 0x3f, 0xb1, 0xf1, 0x28, 0x7e, 0x0b, 0x5e, 0xe6, 0x85, 0xa0, 0x1f, 0xe9,
	0x25, 0x91, 0x94, 0x71, 0x81, 0x2f, 0x2e, 0xb8, 0x14, 0xb8, 0xcc, 0x88, 0x38, 0xfc, 0xcb, 0x59,
	0xce, 0xc5, 0x38, 0x8a, 0xcf, 0x40, 0xf2, 0x07, 0x55, 0xe2, 0xa2, 0x90, 0xb3, 0x9c, 0x55, 0xa4,
	0xe4, 0x34, 0x67, 0x7f, 0x17, 0x1d, 0xc5, 0xef, 0x40, 0x7a, 0x0f, 0xd1, 0xef, 0x52, 0xd0, 0x3d,
	0x8c, 0xdf, 0x80, 0x17, 0x83, 0xae, 0x24, 0x62, 0x51, 0x32, 0x99, 0x33, 0x89, 0xe7, 0x15, 0x29,
	0x05, 0xe5, 0x94, 0x65, 0x92, 0x17, 0x84, 0xcd, 0xc7, 0x8f, 0xa6, 0x3f, 0x47, 0xe0, 0xf4, 0xda,
	0xae, 0xe1, 0x7f, 0xa7, 0x3e, 0x7d, 0x7d, 0xf7, 0xa4, 0x8a, 0xdd, 0x02, 0x8a, 0xd1, 0xe5, 0x74,
	0xef, 0xa4, 0xed, 0x4a, 0x19, 0x0d, 0xed, 0x46, 0x27, 0xba, 0x31, 0xfd, 0x7a, 0x0e, 0xc7, 0xe1,
	0x5a, 0xff, 0x8f, 0x5b, 0x79, 0xdf, 0x7f, 0xbf, 0x44, 0x47, 0x19, 0xc6, 0x5f, 0xa3, 0x93, 0x2c,
	0x58, 0xe1, 0xda, 0xc3, 0x00, 0x77, 0xa8, 0x4a, 0xe1, 0x6e, 0x81, 0xfe, 0xfb, 0xa1, 0xbf, 0xc4,
	0xb5, 0x5f, 0x0e, 0xfd, 0x65, 0x95, 0x2e, 0xfb, 0xfe, 0x8f, 0xe8, 0x34, 0x14, 0x11, 0xc2, 0xb5,
	0x47, 0x68, 0x60, 0x20, 0x54, 0xa5, 0x08, 0xf5, 0x9c, 0xab, 0xc7, 0xfd, 0xc3, 0xce, 0x7e, 0x05,
	0x00, 0x00, 0xff, 0xff, 0xb2, 0xf3, 0x3c, 0xb4, 0xc3, 0x02, 0x00, 0x00,
}
