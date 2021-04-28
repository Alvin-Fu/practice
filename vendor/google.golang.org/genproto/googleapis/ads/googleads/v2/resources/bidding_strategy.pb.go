// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/bidding_strategy.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
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

// A bidding strategy.
type BiddingStrategy struct {
	// The resource name of the bidding strategy.
	// Bidding strategy resource names have the form:
	//
	// `customers/{customer_id}/biddingStrategies/{bidding_strategy_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the bidding strategy.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the bidding strategy.
	// All bidding strategies within an account must be named distinctly.
	//
	// The length of this string should be between 1 and 255, inclusive,
	// in UTF-8 bytes, (trimmed).
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The status of the bidding strategy.
	//
	// This field is read-only.
	Status enums.BiddingStrategyStatusEnum_BiddingStrategyStatus `protobuf:"varint,15,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.BiddingStrategyStatusEnum_BiddingStrategyStatus" json:"status,omitempty"`
	// The type of the bidding strategy.
	// Create a bidding strategy by setting the bidding scheme.
	//
	// This field is read-only.
	Type enums.BiddingStrategyTypeEnum_BiddingStrategyType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.BiddingStrategyTypeEnum_BiddingStrategyType" json:"type,omitempty"`
	// The number of campaigns attached to this bidding strategy.
	//
	// This field is read-only.
	CampaignCount *wrappers.Int64Value `protobuf:"bytes,13,opt,name=campaign_count,json=campaignCount,proto3" json:"campaign_count,omitempty"`
	// The number of non-removed campaigns attached to this bidding strategy.
	//
	// This field is read-only.
	NonRemovedCampaignCount *wrappers.Int64Value `protobuf:"bytes,14,opt,name=non_removed_campaign_count,json=nonRemovedCampaignCount,proto3" json:"non_removed_campaign_count,omitempty"`
	// The bidding scheme.
	//
	// Only one can be set.
	//
	// Types that are valid to be assigned to Scheme:
	//	*BiddingStrategy_EnhancedCpc
	//	*BiddingStrategy_PageOnePromoted
	//	*BiddingStrategy_TargetCpa
	//	*BiddingStrategy_TargetImpressionShare
	//	*BiddingStrategy_TargetOutrankShare
	//	*BiddingStrategy_TargetRoas
	//	*BiddingStrategy_TargetSpend
	Scheme               isBiddingStrategy_Scheme `protobuf_oneof:"scheme"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *BiddingStrategy) Reset()         { *m = BiddingStrategy{} }
func (m *BiddingStrategy) String() string { return proto.CompactTextString(m) }
func (*BiddingStrategy) ProtoMessage()    {}
func (*BiddingStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f01e34c4d3a0d3a, []int{0}
}

func (m *BiddingStrategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingStrategy.Unmarshal(m, b)
}
func (m *BiddingStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingStrategy.Marshal(b, m, deterministic)
}
func (m *BiddingStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingStrategy.Merge(m, src)
}
func (m *BiddingStrategy) XXX_Size() int {
	return xxx_messageInfo_BiddingStrategy.Size(m)
}
func (m *BiddingStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingStrategy proto.InternalMessageInfo

func (m *BiddingStrategy) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *BiddingStrategy) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *BiddingStrategy) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *BiddingStrategy) GetStatus() enums.BiddingStrategyStatusEnum_BiddingStrategyStatus {
	if m != nil {
		return m.Status
	}
	return enums.BiddingStrategyStatusEnum_UNSPECIFIED
}

func (m *BiddingStrategy) GetType() enums.BiddingStrategyTypeEnum_BiddingStrategyType {
	if m != nil {
		return m.Type
	}
	return enums.BiddingStrategyTypeEnum_UNSPECIFIED
}

func (m *BiddingStrategy) GetCampaignCount() *wrappers.Int64Value {
	if m != nil {
		return m.CampaignCount
	}
	return nil
}

func (m *BiddingStrategy) GetNonRemovedCampaignCount() *wrappers.Int64Value {
	if m != nil {
		return m.NonRemovedCampaignCount
	}
	return nil
}

type isBiddingStrategy_Scheme interface {
	isBiddingStrategy_Scheme()
}

type BiddingStrategy_EnhancedCpc struct {
	EnhancedCpc *common.EnhancedCpc `protobuf:"bytes,7,opt,name=enhanced_cpc,json=enhancedCpc,proto3,oneof"`
}

type BiddingStrategy_PageOnePromoted struct {
	PageOnePromoted *common.PageOnePromoted `protobuf:"bytes,8,opt,name=page_one_promoted,json=pageOnePromoted,proto3,oneof"`
}

type BiddingStrategy_TargetCpa struct {
	TargetCpa *common.TargetCpa `protobuf:"bytes,9,opt,name=target_cpa,json=targetCpa,proto3,oneof"`
}

type BiddingStrategy_TargetImpressionShare struct {
	TargetImpressionShare *common.TargetImpressionShare `protobuf:"bytes,48,opt,name=target_impression_share,json=targetImpressionShare,proto3,oneof"`
}

type BiddingStrategy_TargetOutrankShare struct {
	TargetOutrankShare *common.TargetOutrankShare `protobuf:"bytes,10,opt,name=target_outrank_share,json=targetOutrankShare,proto3,oneof"`
}

type BiddingStrategy_TargetRoas struct {
	TargetRoas *common.TargetRoas `protobuf:"bytes,11,opt,name=target_roas,json=targetRoas,proto3,oneof"`
}

type BiddingStrategy_TargetSpend struct {
	TargetSpend *common.TargetSpend `protobuf:"bytes,12,opt,name=target_spend,json=targetSpend,proto3,oneof"`
}

func (*BiddingStrategy_EnhancedCpc) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_PageOnePromoted) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetCpa) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetImpressionShare) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetOutrankShare) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetRoas) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetSpend) isBiddingStrategy_Scheme() {}

func (m *BiddingStrategy) GetScheme() isBiddingStrategy_Scheme {
	if m != nil {
		return m.Scheme
	}
	return nil
}

func (m *BiddingStrategy) GetEnhancedCpc() *common.EnhancedCpc {
	if x, ok := m.GetScheme().(*BiddingStrategy_EnhancedCpc); ok {
		return x.EnhancedCpc
	}
	return nil
}

func (m *BiddingStrategy) GetPageOnePromoted() *common.PageOnePromoted {
	if x, ok := m.GetScheme().(*BiddingStrategy_PageOnePromoted); ok {
		return x.PageOnePromoted
	}
	return nil
}

func (m *BiddingStrategy) GetTargetCpa() *common.TargetCpa {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetCpa); ok {
		return x.TargetCpa
	}
	return nil
}

func (m *BiddingStrategy) GetTargetImpressionShare() *common.TargetImpressionShare {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetImpressionShare); ok {
		return x.TargetImpressionShare
	}
	return nil
}

func (m *BiddingStrategy) GetTargetOutrankShare() *common.TargetOutrankShare {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetOutrankShare); ok {
		return x.TargetOutrankShare
	}
	return nil
}

func (m *BiddingStrategy) GetTargetRoas() *common.TargetRoas {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetRoas); ok {
		return x.TargetRoas
	}
	return nil
}

func (m *BiddingStrategy) GetTargetSpend() *common.TargetSpend {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetSpend); ok {
		return x.TargetSpend
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BiddingStrategy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BiddingStrategy_EnhancedCpc)(nil),
		(*BiddingStrategy_PageOnePromoted)(nil),
		(*BiddingStrategy_TargetCpa)(nil),
		(*BiddingStrategy_TargetImpressionShare)(nil),
		(*BiddingStrategy_TargetOutrankShare)(nil),
		(*BiddingStrategy_TargetRoas)(nil),
		(*BiddingStrategy_TargetSpend)(nil),
	}
}

func init() {
	proto.RegisterType((*BiddingStrategy)(nil), "google.ads.googleads.v2.resources.BiddingStrategy")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/bidding_strategy.proto", fileDescriptor_3f01e34c4d3a0d3a)
}

var fileDescriptor_3f01e34c4d3a0d3a = []byte{
	// 694 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xdd, 0x4e, 0xdb, 0x3c,
	0x18, 0xc7, 0xdb, 0xc2, 0xcb, 0x0b, 0x2e, 0x1f, 0x7a, 0x23, 0x5e, 0x11, 0x31, 0x34, 0xc1, 0x26,
	0x24, 0x36, 0x26, 0x07, 0x65, 0x1f, 0xda, 0xc2, 0x51, 0x5b, 0x21, 0x0a, 0xd2, 0xa0, 0x4a, 0x51,
	0x35, 0x4d, 0xdd, 0x22, 0x93, 0x3c, 0x84, 0x68, 0xc4, 0xb6, 0x6c, 0x87, 0x89, 0xc3, 0xdd, 0xca,
	0x0e, 0x77, 0x29, 0xbb, 0x94, 0x5d, 0xc3, 0x0e, 0xa6, 0x38, 0x4e, 0x10, 0x1f, 0x5d, 0xcb, 0x99,
	0x1f, 0xfb, 0xff, 0xff, 0xfd, 0xdd, 0xa7, 0x4f, 0x12, 0xf4, 0x36, 0x66, 0x2c, 0xbe, 0x00, 0x87,
	0x44, 0xd2, 0x29, 0x96, 0xf9, 0xea, 0xd2, 0x75, 0x04, 0x48, 0x96, 0x89, 0x10, 0xa4, 0x73, 0x9a,
	0x44, 0x51, 0x42, 0xe3, 0x40, 0x2a, 0x41, 0x14, 0xc4, 0x57, 0x98, 0x0b, 0xa6, 0x98, 0xb5, 0x51,
	0xc8, 0x31, 0x89, 0x24, 0xae, 0x9c, 0xf8, 0xd2, 0xc5, 0x95, 0x73, 0xf5, 0xc5, 0x28, 0x78, 0xc8,
	0xd2, 0x94, 0xd1, 0x92, 0x5c, 0x00, 0x57, 0x77, 0x47, 0xa9, 0x81, 0x66, 0xe9, 0xdd, 0x6b, 0x04,
	0x52, 0x11, 0x95, 0x49, 0x63, 0x7e, 0xf7, 0x40, 0xb3, 0xba, 0xe2, 0x60, 0xac, 0x8f, 0x8d, 0x55,
	0x57, 0xa7, 0xd9, 0x99, 0xf3, 0x55, 0x10, 0xce, 0x41, 0x94, 0xe8, 0xb5, 0x12, 0xcd, 0x13, 0x87,
	0x50, 0xca, 0x14, 0x51, 0x09, 0xa3, 0xe6, 0xf4, 0xc9, 0xef, 0x59, 0xb4, 0xd4, 0x2e, 0xe8, 0x7d,
	0x03, 0xb7, 0x9e, 0xa2, 0x85, 0xb2, 0x09, 0x01, 0x25, 0x29, 0xd8, 0xf5, 0xf5, 0xfa, 0xd6, 0x9c,
	0x3f, 0x5f, 0x6e, 0x1e, 0x91, 0x14, 0xac, 0x6d, 0xd4, 0x48, 0x22, 0x7b, 0x6a, 0xbd, 0xbe, 0xd5,
	0x74, 0x1f, 0x99, 0x0e, 0xe2, 0xf2, 0x0e, 0xf8, 0x80, 0xaa, 0x37, 0xaf, 0x06, 0xe4, 0x22, 0x03,
	0xbf, 0x91, 0x44, 0xd6, 0x0e, 0x9a, 0xd6, 0xa0, 0x69, 0x2d, 0x5f, 0xbb, 0x23, 0xef, 0x2b, 0x91,
	0xd0, 0xb8, 0xd0, 0x6b, 0xa5, 0x75, 0x86, 0x66, 0x8a, 0x06, 0xd9, 0x4b, 0xeb, 0xf5, 0xad, 0x45,
	0xf7, 0x08, 0x8f, 0xfa, 0xbf, 0x74, 0x87, 0xf0, 0xad, 0xdf, 0xd0, 0xd7, 0xde, 0x3d, 0x9a, 0xa5,
	0xf7, 0x9f, 0xf8, 0x86, 0x6e, 0x7d, 0x46, 0xd3, 0x79, 0x2f, 0xed, 0x7f, 0x74, 0xca, 0xe1, 0xc3,
	0x52, 0x4e, 0xae, 0x38, 0xdc, 0x97, 0x91, 0xef, 0xfb, 0x9a, 0x6b, 0xb5, 0xd1, 0x62, 0x48, 0x52,
	0x4e, 0x92, 0x98, 0x06, 0x21, 0xcb, 0xa8, 0xb2, 0x17, 0xc6, 0xb7, 0x6c, 0xa1, 0xb4, 0x74, 0x72,
	0x87, 0xf5, 0x01, 0xad, 0x52, 0x46, 0x03, 0x01, 0x29, 0xbb, 0x84, 0x28, 0xb8, 0xc5, 0x5b, 0x1c,
	0xcf, 0x5b, 0xa1, 0x8c, 0xfa, 0x85, 0xbb, 0x73, 0x83, 0xdc, 0x43, 0xf3, 0x40, 0xcf, 0x09, 0x0d,
	0x73, 0x2c, 0x0f, 0xed, 0x7f, 0x35, 0x6b, 0x7b, 0x64, 0x17, 0x8a, 0xc1, 0xc7, 0x7b, 0xc6, 0xd3,
	0xe1, 0x61, 0xb7, 0xe6, 0x37, 0xe1, 0xba, 0xb4, 0x3e, 0xa1, 0xff, 0x38, 0x89, 0x21, 0x60, 0x14,
	0x02, 0x2e, 0x58, 0xca, 0x14, 0x44, 0xf6, 0xac, 0xc6, 0x3a, 0xe3, 0xb0, 0x3d, 0x12, 0xc3, 0x31,
	0x85, 0x9e, 0xb1, 0x75, 0x6b, 0xfe, 0x12, 0xbf, 0xb9, 0x65, 0x1d, 0x22, 0xa4, 0x88, 0x88, 0x41,
	0x05, 0x21, 0x27, 0xf6, 0x9c, 0xe6, 0x3e, 0x1b, 0xc7, 0x3d, 0xd1, 0x8e, 0x0e, 0x27, 0xdd, 0x9a,
	0x3f, 0xa7, 0xca, 0xc2, 0x62, 0x68, 0xc5, 0xb0, 0x92, 0x94, 0x0b, 0x90, 0x32, 0x61, 0x34, 0x90,
	0xe7, 0x44, 0x80, 0xbd, 0xa3, 0xc1, 0xaf, 0x27, 0x03, 0x1f, 0x54, 0xee, 0x7e, 0x6e, 0xee, 0xd6,
	0xfc, 0xff, 0xd5, 0x7d, 0x07, 0xd6, 0x19, 0x5a, 0x36, 0x81, 0x2c, 0x53, 0x82, 0xd0, 0x2f, 0x26,
	0x0d, 0xe9, 0x34, 0x77, 0xb2, 0xb4, 0xe3, 0xc2, 0x5a, 0x46, 0x59, 0xea, 0xce, 0xae, 0xf5, 0x1e,
	0x35, 0x4d, 0x8e, 0x60, 0x44, 0xda, 0x4d, 0x8d, 0x7f, 0x3e, 0x19, 0xde, 0x67, 0x44, 0x76, 0x6b,
	0xbe, 0xe9, 0x72, 0x5e, 0xe5, 0x43, 0x62, 0x70, 0x92, 0x03, 0x8d, 0xec, 0xf9, 0xc9, 0x86, 0xa4,
	0xe0, 0xf5, 0x73, 0x4b, 0x3e, 0x24, 0xea, 0xba, 0x6c, 0xcf, 0xa2, 0x19, 0x19, 0x9e, 0x43, 0x0a,
	0xed, 0x6f, 0x0d, 0xb4, 0x19, 0xb2, 0x14, 0x8f, 0x7d, 0x19, 0xb7, 0x97, 0x6f, 0x3d, 0x63, 0xbd,
	0x7c, 0xce, 0x7b, 0xf5, 0x8f, 0x87, 0xc6, 0x1a, 0xb3, 0x0b, 0x42, 0x63, 0xcc, 0x44, 0xec, 0xc4,
	0x40, 0xf5, 0x53, 0x50, 0xbe, 0x49, 0x79, 0x22, 0xff, 0xf2, 0x81, 0xd8, 0xad, 0x56, 0xdf, 0x1b,
	0x53, 0xfb, 0xad, 0xd6, 0x8f, 0xc6, 0xc6, 0x7e, 0x81, 0x6c, 0x45, 0x12, 0x17, 0xcb, 0x7c, 0x35,
	0x70, 0xb1, 0x5f, 0x2a, 0x7f, 0x96, 0x9a, 0x61, 0x2b, 0x92, 0xc3, 0x4a, 0x33, 0x1c, 0xb8, 0xc3,
	0x4a, 0xf3, 0xab, 0xb1, 0x59, 0x1c, 0x78, 0x5e, 0x2b, 0x92, 0x9e, 0x57, 0xa9, 0x3c, 0x6f, 0xe0,
	0x7a, 0x5e, 0xa5, 0x3b, 0x9d, 0xd1, 0x97, 0x7d, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xc9,
	0xac, 0x3a, 0xcc, 0x06, 0x00, 0x00,
}
