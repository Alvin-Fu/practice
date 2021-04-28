// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/keyword_plan.proto

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

// A Keyword Planner plan.
// Max number of saved keyword plans: 10000.
// It's possible to remove plans if limit is reached.
type KeywordPlan struct {
	// The resource name of the Keyword Planner plan.
	// KeywordPlan resource names have the form:
	//
	// `customers/{customer_id}/keywordPlans/{kp_plan_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the keyword plan.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the keyword plan.
	//
	// This field is required and should not be empty when creating new keyword
	// plans.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The date period used for forecasting the plan.
	ForecastPeriod       *KeywordPlanForecastPeriod `protobuf:"bytes,4,opt,name=forecast_period,json=forecastPeriod,proto3" json:"forecast_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *KeywordPlan) Reset()         { *m = KeywordPlan{} }
func (m *KeywordPlan) String() string { return proto.CompactTextString(m) }
func (*KeywordPlan) ProtoMessage()    {}
func (*KeywordPlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_db2ef87e79a4b462, []int{0}
}

func (m *KeywordPlan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlan.Unmarshal(m, b)
}
func (m *KeywordPlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlan.Marshal(b, m, deterministic)
}
func (m *KeywordPlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlan.Merge(m, src)
}
func (m *KeywordPlan) XXX_Size() int {
	return xxx_messageInfo_KeywordPlan.Size(m)
}
func (m *KeywordPlan) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlan.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlan proto.InternalMessageInfo

func (m *KeywordPlan) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *KeywordPlan) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *KeywordPlan) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *KeywordPlan) GetForecastPeriod() *KeywordPlanForecastPeriod {
	if m != nil {
		return m.ForecastPeriod
	}
	return nil
}

// The forecasting period associated with the keyword plan.
type KeywordPlanForecastPeriod struct {
	// Required. The date used for forecasting the Plan.
	//
	// Types that are valid to be assigned to Interval:
	//	*KeywordPlanForecastPeriod_DateInterval
	//	*KeywordPlanForecastPeriod_DateRange
	Interval             isKeywordPlanForecastPeriod_Interval `protobuf_oneof:"interval"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *KeywordPlanForecastPeriod) Reset()         { *m = KeywordPlanForecastPeriod{} }
func (m *KeywordPlanForecastPeriod) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanForecastPeriod) ProtoMessage()    {}
func (*KeywordPlanForecastPeriod) Descriptor() ([]byte, []int) {
	return fileDescriptor_db2ef87e79a4b462, []int{1}
}

func (m *KeywordPlanForecastPeriod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanForecastPeriod.Unmarshal(m, b)
}
func (m *KeywordPlanForecastPeriod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanForecastPeriod.Marshal(b, m, deterministic)
}
func (m *KeywordPlanForecastPeriod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanForecastPeriod.Merge(m, src)
}
func (m *KeywordPlanForecastPeriod) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanForecastPeriod.Size(m)
}
func (m *KeywordPlanForecastPeriod) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanForecastPeriod.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanForecastPeriod proto.InternalMessageInfo

type isKeywordPlanForecastPeriod_Interval interface {
	isKeywordPlanForecastPeriod_Interval()
}

type KeywordPlanForecastPeriod_DateInterval struct {
	DateInterval enums.KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval `protobuf:"varint,1,opt,name=date_interval,json=dateInterval,proto3,enum=google.ads.googleads.v2.enums.KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval,oneof"`
}

type KeywordPlanForecastPeriod_DateRange struct {
	DateRange *common.DateRange `protobuf:"bytes,2,opt,name=date_range,json=dateRange,proto3,oneof"`
}

func (*KeywordPlanForecastPeriod_DateInterval) isKeywordPlanForecastPeriod_Interval() {}

func (*KeywordPlanForecastPeriod_DateRange) isKeywordPlanForecastPeriod_Interval() {}

func (m *KeywordPlanForecastPeriod) GetInterval() isKeywordPlanForecastPeriod_Interval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *KeywordPlanForecastPeriod) GetDateInterval() enums.KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval {
	if x, ok := m.GetInterval().(*KeywordPlanForecastPeriod_DateInterval); ok {
		return x.DateInterval
	}
	return enums.KeywordPlanForecastIntervalEnum_UNSPECIFIED
}

func (m *KeywordPlanForecastPeriod) GetDateRange() *common.DateRange {
	if x, ok := m.GetInterval().(*KeywordPlanForecastPeriod_DateRange); ok {
		return x.DateRange
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*KeywordPlanForecastPeriod) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*KeywordPlanForecastPeriod_DateInterval)(nil),
		(*KeywordPlanForecastPeriod_DateRange)(nil),
	}
}

func init() {
	proto.RegisterType((*KeywordPlan)(nil), "google.ads.googleads.v2.resources.KeywordPlan")
	proto.RegisterType((*KeywordPlanForecastPeriod)(nil), "google.ads.googleads.v2.resources.KeywordPlanForecastPeriod")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/keyword_plan.proto", fileDescriptor_db2ef87e79a4b462)
}

var fileDescriptor_db2ef87e79a4b462 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xc7, 0x9b, 0xb4, 0x88, 0x9d, 0x7e, 0x28, 0xb9, 0x5a, 0x6b, 0x91, 0xb6, 0x52, 0xa8, 0x0a,
	0x13, 0x89, 0xc5, 0x8b, 0xe8, 0x4d, 0x16, 0xb5, 0x1f, 0x82, 0x2c, 0x11, 0xf6, 0xa2, 0x2c, 0x2c,
	0xd3, 0x9d, 0xb3, 0x21, 0x98, 0xcc, 0x84, 0x99, 0xc9, 0x16, 0x2f, 0x7d, 0x15, 0x2f, 0x7d, 0x14,
	0x1f, 0xc5, 0x17, 0xd0, 0x1b, 0x41, 0x32, 0x5f, 0xb4, 0xd8, 0x74, 0xef, 0xce, 0xd9, 0xf9, 0x9d,
	0xff, 0xff, 0x7c, 0x64, 0xd1, 0x71, 0xc1, 0x79, 0x51, 0x41, 0x4c, 0xa8, 0x8c, 0x4d, 0xd8, 0x45,
	0x8b, 0x24, 0x16, 0x20, 0x79, 0x2b, 0x66, 0x20, 0xe3, 0x2f, 0xf0, 0xf5, 0x8a, 0x0b, 0x3a, 0x6d,
	0x2a, 0xc2, 0x70, 0x23, 0xb8, 0xe2, 0xd1, 0xbe, 0x41, 0x31, 0xa1, 0x12, 0xfb, 0x2a, 0xbc, 0x48,
	0xb0, 0xaf, 0xda, 0x79, 0xde, 0x27, 0x3c, 0xe3, 0x75, 0xcd, 0x59, 0x4c, 0x89, 0x02, 0x69, 0xe4,
	0x76, 0x86, 0x7d, 0x2c, 0xb0, 0xb6, 0xbe, 0xd9, 0xc0, 0x74, 0xce, 0x05, 0xcc, 0x88, 0x54, 0xd3,
	0x92, 0x29, 0x10, 0x0b, 0x52, 0x59, 0x8d, 0x27, 0x56, 0x43, 0x67, 0x97, 0xed, 0x3c, 0xbe, 0x12,
	0xa4, 0x69, 0x40, 0x38, 0x8f, 0x5d, 0xe7, 0xd1, 0x94, 0x31, 0x61, 0x8c, 0x2b, 0xa2, 0x4a, 0xce,
	0xec, 0xeb, 0xc1, 0x9f, 0x00, 0x6d, 0x7c, 0x34, 0x36, 0xa3, 0x8a, 0xb0, 0xe8, 0x29, 0xda, 0x72,
	0xa3, 0x4c, 0x19, 0xa9, 0x61, 0x10, 0xec, 0x05, 0x47, 0xeb, 0xf9, 0xa6, 0xfb, 0xf1, 0x13, 0xa9,
	0x21, 0x7a, 0x81, 0xc2, 0x92, 0x0e, 0xc2, 0xbd, 0xe0, 0x68, 0x23, 0x79, 0x6c, 0xf7, 0x80, 0x9d,
	0x3f, 0x3e, 0x63, 0xea, 0xf5, 0xf1, 0x98, 0x54, 0x2d, 0xe4, 0x61, 0x49, 0xa3, 0x97, 0x68, 0x4d,
	0x0b, 0xad, 0x6a, 0x7c, 0xf7, 0x3f, 0xfc, 0xb3, 0x12, 0x25, 0x2b, 0x0c, 0xaf, 0xc9, 0x08, 0xd0,
	0x03, 0x3f, 0x6c, 0x03, 0xa2, 0xe4, 0x74, 0xb0, 0xa6, 0x8b, 0xdf, 0xe2, 0xa5, 0xeb, 0xc7, 0xd7,
	0x86, 0xf9, 0x60, 0x45, 0x46, 0x5a, 0x23, 0xdf, 0x9e, 0xdf, 0xc8, 0x0f, 0x7e, 0x07, 0xe8, 0x51,
	0x2f, 0x1d, 0x7d, 0x0b, 0xd0, 0x56, 0x77, 0x2a, 0xbf, 0x6e, 0xbd, 0x89, 0xed, 0xe4, 0xa2, 0xb7,
	0x07, 0x7d, 0xb3, 0xdb, 0xfc, 0xcf, 0xac, 0xc2, 0x7b, 0xd6, 0xd6, 0x77, 0xbd, 0x9f, 0xae, 0xe4,
	0x9b, 0x9d, 0xa5, 0xcb, 0xa3, 0x73, 0x84, 0x74, 0x0b, 0x82, 0xb0, 0x02, 0xec, 0xbe, 0x9f, 0xf5,
	0xfa, 0x9b, 0xef, 0x0b, 0xbf, 0x23, 0x0a, 0xf2, 0xae, 0xe0, 0x74, 0x25, 0x5f, 0xa7, 0x2e, 0x19,
	0x22, 0x74, 0xdf, 0x4d, 0x32, 0xfc, 0x1b, 0xa0, 0xc3, 0x19, 0xaf, 0x97, 0x6f, 0x73, 0xf8, 0xf0,
	0x5a, 0xbb, 0xa3, 0xee, 0x62, 0xa3, 0xe0, 0xe2, 0xdc, 0x96, 0x15, 0xbc, 0x22, 0xac, 0xc0, 0x5c,
	0x14, 0x71, 0x01, 0x4c, 0xdf, 0xd3, 0x7d, 0xc4, 0x4d, 0x29, 0xef, 0xf8, 0x63, 0xbd, 0xf1, 0xd1,
	0xf7, 0x70, 0xf5, 0x24, 0xcb, 0x7e, 0x84, 0xfb, 0x27, 0x46, 0x32, 0xa3, 0x12, 0x9b, 0xb0, 0x8b,
	0xc6, 0x09, 0xce, 0x1d, 0xf9, 0xd3, 0x31, 0x93, 0x8c, 0xca, 0x89, 0x67, 0x26, 0xe3, 0x64, 0xe2,
	0x99, 0x5f, 0xe1, 0xa1, 0x79, 0x48, 0xd3, 0x8c, 0xca, 0x34, 0xf5, 0x54, 0x9a, 0x8e, 0x93, 0x34,
	0xf5, 0xdc, 0xe5, 0x3d, 0xdd, 0xec, 0xab, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x35, 0x7d,
	0x8b, 0x04, 0x04, 0x00, 0x00,
}
