// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/v3/metric.proto

package monitoring

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/distribution"
	_ "google.golang.org/genproto/googleapis/api/label"
	metric "google.golang.org/genproto/googleapis/api/metric"
	monitoredres "google.golang.org/genproto/googleapis/api/monitoredres"
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

// A single data point in a time series.
type Point struct {
	// The time interval to which the data point applies.  For `GAUGE` metrics,
	// the start time is optional, but if it is supplied, it must equal the
	// end time.  For `DELTA` metrics, the start
	// and end time should specify a non-zero interval, with subsequent points
	// specifying contiguous and non-overlapping intervals.  For `CUMULATIVE`
	// metrics, the start and end time should specify a non-zero interval, with
	// subsequent points specifying the same start time and increasing end times,
	// until an event resets the cumulative value to zero and sets a new start
	// time for the following points.
	Interval *TimeInterval `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
	// The value of the data point.
	Value                *TypedValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76199a3d2c4c21e, []int{0}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetInterval() *TimeInterval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *Point) GetValue() *TypedValue {
	if m != nil {
		return m.Value
	}
	return nil
}

// A collection of data points that describes the time-varying values
// of a metric. A time series is identified by a combination of a
// fully-specified monitored resource and a fully-specified metric.
// This type is used for both listing and creating time series.
type TimeSeries struct {
	// The associated metric. A fully-specified metric used to identify the time
	// series.
	Metric *metric.Metric `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
	// The associated monitored resource.  Custom metrics can use only certain
	// monitored resource types in their time series data.
	Resource *monitoredres.MonitoredResource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	// Output only. The associated monitored resource metadata. When reading a
	// a timeseries, this field will include metadata labels that are explicitly
	// named in the reduction. When creating a timeseries, this field is ignored.
	Metadata *monitoredres.MonitoredResourceMetadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The metric kind of the time series. When listing time series, this metric
	// kind might be different from the metric kind of the associated metric if
	// this time series is an alignment or reduction of other time series.
	//
	// When creating a time series, this field is optional. If present, it must be
	// the same as the metric kind of the associated metric. If the associated
	// metric's descriptor must be auto-created, then this field specifies the
	// metric kind of the new descriptor and must be either `GAUGE` (the default)
	// or `CUMULATIVE`.
	MetricKind metric.MetricDescriptor_MetricKind `protobuf:"varint,3,opt,name=metric_kind,json=metricKind,proto3,enum=google.api.MetricDescriptor_MetricKind" json:"metric_kind,omitempty"`
	// The value type of the time series. When listing time series, this value
	// type might be different from the value type of the associated metric if
	// this time series is an alignment or reduction of other time series.
	//
	// When creating a time series, this field is optional. If present, it must be
	// the same as the type of the data in the `points` field.
	ValueType metric.MetricDescriptor_ValueType `protobuf:"varint,4,opt,name=value_type,json=valueType,proto3,enum=google.api.MetricDescriptor_ValueType" json:"value_type,omitempty"`
	// The data points of this time series. When listing time series, points are
	// returned in reverse time order.
	//
	// When creating a time series, this field must contain exactly one point and
	// the point's type must be the same as the value type of the associated
	// metric. If the associated metric's descriptor must be auto-created, then
	// the value type of the descriptor is determined by the point's type, which
	// must be `BOOL`, `INT64`, `DOUBLE`, or `DISTRIBUTION`.
	Points               []*Point `protobuf:"bytes,5,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeSeries) Reset()         { *m = TimeSeries{} }
func (m *TimeSeries) String() string { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()    {}
func (*TimeSeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76199a3d2c4c21e, []int{1}
}

func (m *TimeSeries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSeries.Unmarshal(m, b)
}
func (m *TimeSeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSeries.Marshal(b, m, deterministic)
}
func (m *TimeSeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeries.Merge(m, src)
}
func (m *TimeSeries) XXX_Size() int {
	return xxx_messageInfo_TimeSeries.Size(m)
}
func (m *TimeSeries) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeries.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeries proto.InternalMessageInfo

func (m *TimeSeries) GetMetric() *metric.Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func (m *TimeSeries) GetResource() *monitoredres.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *TimeSeries) GetMetadata() *monitoredres.MonitoredResourceMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TimeSeries) GetMetricKind() metric.MetricDescriptor_MetricKind {
	if m != nil {
		return m.MetricKind
	}
	return metric.MetricDescriptor_METRIC_KIND_UNSPECIFIED
}

func (m *TimeSeries) GetValueType() metric.MetricDescriptor_ValueType {
	if m != nil {
		return m.ValueType
	}
	return metric.MetricDescriptor_VALUE_TYPE_UNSPECIFIED
}

func (m *TimeSeries) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

func init() {
	proto.RegisterType((*Point)(nil), "google.monitoring.v3.Point")
	proto.RegisterType((*TimeSeries)(nil), "google.monitoring.v3.TimeSeries")
}

func init() { proto.RegisterFile("google/monitoring/v3/metric.proto", fileDescriptor_c76199a3d2c4c21e) }

var fileDescriptor_c76199a3d2c4c21e = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x49, 0xd7, 0xd6, 0x3a, 0x05, 0x2f, 0x06, 0xd1, 0x50, 0x59, 0xa8, 0x15, 0xb5, 0x78,
	0x91, 0x40, 0x03, 0x82, 0x08, 0x0b, 0xae, 0x8a, 0x8a, 0x2c, 0x94, 0x51, 0x7a, 0x21, 0x85, 0x32,
	0x4d, 0x0e, 0xe1, 0x60, 0xe6, 0x0f, 0x93, 0x69, 0x60, 0xaf, 0x7c, 0x18, 0xef, 0x7c, 0x14, 0x9f,
	0xc9, 0x0b, 0xc9, 0xcc, 0xa4, 0xed, 0x62, 0xb7, 0x77, 0xc9, 0x7c, 0xbf, 0xef, 0x7c, 0x67, 0xce,
	0x1c, 0xf2, 0xa4, 0x54, 0xaa, 0xac, 0x20, 0x15, 0x4a, 0xa2, 0x55, 0x06, 0x65, 0x99, 0x36, 0x59,
	0x2a, 0xc0, 0x1a, 0xcc, 0x13, 0x6d, 0x94, 0x55, 0xf4, 0x81, 0x47, 0x92, 0x3d, 0x92, 0x34, 0xd9,
	0xf8, 0x3c, 0x18, 0xb9, 0xc6, 0xb4, 0xc0, 0xda, 0x1a, 0xdc, 0x6c, 0x2d, 0x2a, 0xe9, 0x4d, 0xe3,
	0x87, 0x07, 0x72, 0xc5, 0x37, 0x50, 0x85, 0xf3, 0x47, 0x07, 0xe7, 0x87, 0x29, 0xe3, 0xa7, 0x87,
	0x82, 0x4f, 0x82, 0x62, 0x6d, 0xa0, 0x56, 0x5b, 0x93, 0x43, 0x80, 0x8e, 0x77, 0x9b, 0x2b, 0x21,
	0xba, 0xe0, 0xe9, 0x4f, 0xd2, 0x5f, 0x28, 0x94, 0x96, 0x5e, 0x90, 0x21, 0x4a, 0x0b, 0xa6, 0xe1,
	0x55, 0x1c, 0x4d, 0xa2, 0xd9, 0x68, 0x3e, 0x4d, 0x8e, 0xdd, 0x24, 0xf9, 0x86, 0x02, 0x3e, 0x07,
	0x92, 0xed, 0x3c, 0xf4, 0x15, 0xe9, 0x37, 0xbc, 0xda, 0x42, 0xdc, 0x73, 0xe6, 0xc9, 0x2d, 0xe6,
	0x6b, 0x0d, 0xc5, 0xb2, 0xe5, 0x98, 0xc7, 0xa7, 0x7f, 0x7b, 0x84, 0xb4, 0x25, 0xbf, 0x82, 0x41,
	0xa8, 0xe9, 0x4b, 0x32, 0xf0, 0xf7, 0x0c, 0x4d, 0xd0, 0xae, 0x0e, 0xd7, 0x98, 0x5c, 0x39, 0x85,
	0x05, 0x82, 0xbe, 0x26, 0xc3, 0xee, 0xc2, 0x21, 0xf5, 0xfc, 0x06, 0xdd, 0x8d, 0x85, 0x05, 0x88,
	0xed, 0x70, 0xfa, 0x96, 0x0c, 0x05, 0x58, 0x5e, 0x70, 0xcb, 0xe3, 0xbb, 0xce, 0xfa, 0xec, 0xa4,
	0xf5, 0x2a, 0xc0, 0x6c, 0x67, 0xa3, 0x9f, 0xc8, 0xc8, 0xf7, 0xb1, 0xfe, 0x81, 0xb2, 0x88, 0xcf,
	0x26, 0xd1, 0xec, 0xfe, 0xfc, 0xc5, 0xff, 0xed, 0xbe, 0x87, 0x3a, 0x37, 0xa8, 0xad, 0x32, 0xe1,
	0xe0, 0x0b, 0xca, 0x82, 0x11, 0xb1, 0xfb, 0xa6, 0x1f, 0x08, 0x71, 0xb3, 0x58, 0xdb, 0x6b, 0x0d,
	0xf1, 0x1d, 0x57, 0xe8, 0xf9, 0xc9, 0x42, 0x6e, 0x82, 0xed, 0x2c, 0xd9, 0xbd, 0xa6, 0xfb, 0xa4,
	0x19, 0x19, 0xe8, 0xf6, 0x29, 0xeb, 0xb8, 0x3f, 0x39, 0x9b, 0x8d, 0xe6, 0x8f, 0x8f, 0x3f, 0x81,
	0x7b, 0x6e, 0x16, 0xd0, 0xcb, 0x5f, 0x11, 0x89, 0x73, 0x25, 0x8e, 0xa2, 0x97, 0x23, 0x1f, 0xbc,
	0x68, 0x37, 0x65, 0x11, 0x7d, 0xbf, 0x08, 0x50, 0xa9, 0x2a, 0x2e, 0xcb, 0x44, 0x99, 0x32, 0x2d,
	0x41, 0xba, 0x3d, 0x4a, 0xbd, 0xc4, 0x35, 0xd6, 0x37, 0xb7, 0xed, 0xcd, 0xfe, 0xef, 0x77, 0x6f,
	0xfc, 0xd1, 0x17, 0x78, 0x57, 0xa9, 0x6d, 0xd1, 0x0d, 0xb9, 0xcd, 0x5a, 0x66, 0x7f, 0x3a, 0x71,
	0xe5, 0xc4, 0xd5, 0x5e, 0x5c, 0x2d, 0xb3, 0xcd, 0xc0, 0x85, 0x64, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x06, 0x37, 0xbb, 0x92, 0x7f, 0x03, 0x00, 0x00,
}
