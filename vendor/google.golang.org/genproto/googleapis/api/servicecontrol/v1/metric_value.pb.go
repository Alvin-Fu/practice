// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/servicecontrol/v1/metric_value.proto

package servicecontrol

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/type/money"
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

// Represents a single metric value.
type MetricValue struct {
	// The labels describing the metric value.
	// See comments on
	// [google.api.servicecontrol.v1.Operation.labels][google.api.servicecontrol.v1.Operation.labels]
	// for the overriding relationship.
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The start of the time period over which this metric value's measurement
	// applies. The time period has different semantics for different metric
	// types (cumulative, delta, and gauge). See the metric definition
	// documentation in the service configuration for details.
	StartTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The end of the time period over which this metric value's measurement
	// applies.
	EndTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// The value. The type of value used in the request must
	// agree with the metric definition in the service configuration, otherwise
	// the MetricValue is rejected.
	//
	// Types that are valid to be assigned to Value:
	//	*MetricValue_BoolValue
	//	*MetricValue_Int64Value
	//	*MetricValue_DoubleValue
	//	*MetricValue_StringValue
	//	*MetricValue_DistributionValue
	Value                isMetricValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MetricValue) Reset()         { *m = MetricValue{} }
func (m *MetricValue) String() string { return proto.CompactTextString(m) }
func (*MetricValue) ProtoMessage()    {}
func (*MetricValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8818c371cfc5a8d3, []int{0}
}

func (m *MetricValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricValue.Unmarshal(m, b)
}
func (m *MetricValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricValue.Marshal(b, m, deterministic)
}
func (m *MetricValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricValue.Merge(m, src)
}
func (m *MetricValue) XXX_Size() int {
	return xxx_messageInfo_MetricValue.Size(m)
}
func (m *MetricValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricValue.DiscardUnknown(m)
}

var xxx_messageInfo_MetricValue proto.InternalMessageInfo

func (m *MetricValue) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *MetricValue) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *MetricValue) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type isMetricValue_Value interface {
	isMetricValue_Value()
}

type MetricValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type MetricValue_Int64Value struct {
	Int64Value int64 `protobuf:"varint,5,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type MetricValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,6,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type MetricValue_StringValue struct {
	StringValue string `protobuf:"bytes,7,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type MetricValue_DistributionValue struct {
	DistributionValue *Distribution `protobuf:"bytes,8,opt,name=distribution_value,json=distributionValue,proto3,oneof"`
}

func (*MetricValue_BoolValue) isMetricValue_Value() {}

func (*MetricValue_Int64Value) isMetricValue_Value() {}

func (*MetricValue_DoubleValue) isMetricValue_Value() {}

func (*MetricValue_StringValue) isMetricValue_Value() {}

func (*MetricValue_DistributionValue) isMetricValue_Value() {}

func (m *MetricValue) GetValue() isMetricValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MetricValue) GetBoolValue() bool {
	if x, ok := m.GetValue().(*MetricValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *MetricValue) GetInt64Value() int64 {
	if x, ok := m.GetValue().(*MetricValue_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *MetricValue) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*MetricValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *MetricValue) GetStringValue() string {
	if x, ok := m.GetValue().(*MetricValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *MetricValue) GetDistributionValue() *Distribution {
	if x, ok := m.GetValue().(*MetricValue_DistributionValue); ok {
		return x.DistributionValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MetricValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MetricValue_BoolValue)(nil),
		(*MetricValue_Int64Value)(nil),
		(*MetricValue_DoubleValue)(nil),
		(*MetricValue_StringValue)(nil),
		(*MetricValue_DistributionValue)(nil),
	}
}

// Represents a set of metric values in the same metric.
// Each metric value in the set should have a unique combination of start time,
// end time, and label values.
type MetricValueSet struct {
	// The metric name defined in the service configuration.
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName,proto3" json:"metric_name,omitempty"`
	// The values in this metric.
	MetricValues         []*MetricValue `protobuf:"bytes,2,rep,name=metric_values,json=metricValues,proto3" json:"metric_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MetricValueSet) Reset()         { *m = MetricValueSet{} }
func (m *MetricValueSet) String() string { return proto.CompactTextString(m) }
func (*MetricValueSet) ProtoMessage()    {}
func (*MetricValueSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_8818c371cfc5a8d3, []int{1}
}

func (m *MetricValueSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricValueSet.Unmarshal(m, b)
}
func (m *MetricValueSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricValueSet.Marshal(b, m, deterministic)
}
func (m *MetricValueSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricValueSet.Merge(m, src)
}
func (m *MetricValueSet) XXX_Size() int {
	return xxx_messageInfo_MetricValueSet.Size(m)
}
func (m *MetricValueSet) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricValueSet.DiscardUnknown(m)
}

var xxx_messageInfo_MetricValueSet proto.InternalMessageInfo

func (m *MetricValueSet) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

func (m *MetricValueSet) GetMetricValues() []*MetricValue {
	if m != nil {
		return m.MetricValues
	}
	return nil
}

func init() {
	proto.RegisterType((*MetricValue)(nil), "google.api.servicecontrol.v1.MetricValue")
	proto.RegisterMapType((map[string]string)(nil), "google.api.servicecontrol.v1.MetricValue.LabelsEntry")
	proto.RegisterType((*MetricValueSet)(nil), "google.api.servicecontrol.v1.MetricValueSet")
}

func init() {
	proto.RegisterFile("google/api/servicecontrol/v1/metric_value.proto", fileDescriptor_8818c371cfc5a8d3)
}

var fileDescriptor_8818c371cfc5a8d3 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x3b, 0x8d, 0xdb, 0x1f, 0x2f, 0xab, 0x68, 0x14, 0x0c, 0x65, 0xa1, 0x71, 0xbd, 0x44,
	0x0f, 0x13, 0x76, 0x75, 0xc5, 0xd5, 0x5b, 0x51, 0x28, 0xe2, 0x2e, 0x4b, 0x14, 0x0f, 0x7a, 0x58,
	0x26, 0xed, 0x33, 0x0c, 0x26, 0x33, 0x21, 0x33, 0x2d, 0xf4, 0xe8, 0xcd, 0x3f, 0xd9, 0x8b, 0x20,
	0xf3, 0xa3, 0x9a, 0x82, 0xd4, 0xbd, 0xe5, 0x7d, 0xf3, 0xfd, 0xbc, 0x7c, 0x27, 0xef, 0x0d, 0x64,
	0xa5, 0x94, 0x65, 0x85, 0x19, 0x6b, 0x78, 0xa6, 0xb0, 0x5d, 0xf3, 0x05, 0x2e, 0xa4, 0xd0, 0xad,
	0xac, 0xb2, 0xf5, 0x49, 0x56, 0xa3, 0x6e, 0xf9, 0xe2, 0x7a, 0xcd, 0xaa, 0x15, 0xd2, 0xa6, 0x95,
	0x5a, 0x46, 0x47, 0x0e, 0xa0, 0xac, 0xe1, 0x74, 0x17, 0xa0, 0xeb, 0x93, 0xc9, 0x51, 0xa7, 0x1d,
	0x13, 0x42, 0x6a, 0xa6, 0xb9, 0x14, 0xca, 0xb1, 0x93, 0xfd, 0x1f, 0x5b, 0x72, 0xa5, 0x5b, 0x5e,
	0xac, 0x0c, 0xe1, 0x81, 0xa9, 0x07, 0x6c, 0x55, 0xac, 0xbe, 0x66, 0x9a, 0xd7, 0xa8, 0x34, 0xab,
	0x1b, 0x6f, 0x78, 0xe8, 0x0d, 0x7a, 0xd3, 0x60, 0x56, 0x4b, 0x81, 0x1b, 0xf7, 0xe2, 0xf8, 0x57,
	0x00, 0xe1, 0x85, 0x4d, 0xff, 0xc9, 0x84, 0x8f, 0x2e, 0x60, 0x50, 0xb1, 0x02, 0x2b, 0x15, 0x93,
	0x24, 0x48, 0xc3, 0xd3, 0x33, 0xba, 0xef, 0x1c, 0xb4, 0x83, 0xd2, 0xf7, 0x96, 0x7b, 0x2b, 0x74,
	0xbb, 0xc9, 0x7d, 0x93, 0xe8, 0x1c, 0x40, 0x69, 0xd6, 0xea, 0x6b, 0x13, 0x28, 0xee, 0x27, 0x24,
	0x0d, 0x4f, 0x27, 0xdb, 0x96, 0xdb, 0xb4, 0xf4, 0xe3, 0x36, 0x6d, 0x3e, 0xb6, 0x6e, 0x53, 0x47,
	0x67, 0x30, 0x42, 0xb1, 0x74, 0x60, 0xf0, 0x5f, 0x70, 0x88, 0x62, 0x69, 0xb1, 0x29, 0x40, 0x21,
	0x65, 0xe5, 0x66, 0x11, 0xdf, 0x4a, 0x48, 0x3a, 0x9a, 0xf7, 0xf2, 0xb1, 0xd1, 0xdc, 0x09, 0x1f,
	0x41, 0xc8, 0x85, 0x7e, 0xf1, 0xdc, 0x3b, 0x0e, 0x12, 0x92, 0x06, 0xf3, 0x5e, 0x0e, 0x56, 0x74,
	0x96, 0xc7, 0x70, 0xb8, 0x94, 0xab, 0xa2, 0x42, 0xef, 0x19, 0x24, 0x24, 0x25, 0xf3, 0x5e, 0x1e,
	0x3a, 0xf5, 0x8f, 0xc9, 0xcc, 0x41, 0x94, 0xde, 0x34, 0x4c, 0x48, 0x3a, 0x36, 0x26, 0xa7, 0x3a,
	0xd3, 0x17, 0x88, 0xba, 0xe3, 0xf2, 0xd6, 0x91, 0x3d, 0xce, 0xd3, 0xfd, 0xbf, 0xf6, 0x4d, 0x87,
	0x9b, 0xf7, 0xf2, 0x7b, 0xdd, 0x3e, 0xb6, 0xf9, 0xe4, 0x1c, 0xc2, 0xce, 0x3f, 0x8f, 0xee, 0x42,
	0xf0, 0x0d, 0x37, 0x31, 0x31, 0x39, 0x72, 0xf3, 0x18, 0x3d, 0x80, 0x03, 0xf7, 0xc1, 0xbe, 0xd5,
	0x5c, 0xf1, 0xaa, 0xff, 0x92, 0xcc, 0x86, 0xfe, 0xcd, 0xf1, 0x77, 0x02, 0x77, 0x3a, 0x43, 0xfc,
	0x80, 0x3a, 0x9a, 0x42, 0xe8, 0xf7, 0x59, 0xb0, 0x1a, 0x7d, 0x3f, 0x70, 0xd2, 0x25, 0xab, 0x31,
	0xba, 0x84, 0xdb, 0xdd, 0x85, 0x57, 0x71, 0xdf, 0xae, 0xca, 0x93, 0x1b, 0xaf, 0x4a, 0x7e, 0x58,
	0xff, 0x2d, 0xd4, 0xec, 0x07, 0x81, 0x64, 0x21, 0xeb, 0xbd, 0xf8, 0xec, 0xfe, 0x6e, 0xca, 0x2b,
	0xb3, 0x02, 0x57, 0xe4, 0xf3, 0x3b, 0x0f, 0x95, 0xb2, 0x62, 0xa2, 0xa4, 0xb2, 0x2d, 0xb3, 0x12,
	0x85, 0x5d, 0x10, 0x7f, 0x8b, 0x58, 0xc3, 0xd5, 0xbf, 0x6f, 0xd2, 0xeb, 0x5d, 0xe5, 0x27, 0x21,
	0xc5, 0xc0, 0x92, 0xcf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x00, 0x1a, 0xde, 0xef, 0x03,
	0x00, 0x00,
}
