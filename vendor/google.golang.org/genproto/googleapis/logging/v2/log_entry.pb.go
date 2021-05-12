// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/logging/v2/log_entry.proto

package logging

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	monitoredres "google.golang.org/genproto/googleapis/api/monitoredres"
	_type "google.golang.org/genproto/googleapis/logging/type"
	_ "google.golang.org/genproto/googleapis/rpc/status"
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

// An individual entry in a log.
//
type LogEntry struct {
	// Required. The resource name of the log to which this log entry belongs:
	//
	//     "projects/[PROJECT_ID]/logs/[LOG_ID]"
	//     "organizations/[ORGANIZATION_ID]/logs/[LOG_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]/logs/[LOG_ID]"
	//     "folders/[FOLDER_ID]/logs/[LOG_ID]"
	//
	// A project number may optionally be used in place of PROJECT_ID. The project
	// number is translated to its corresponding PROJECT_ID internally and the
	// `log_name` field will contain PROJECT_ID in queries and exports.
	//
	// `[LOG_ID]` must be URL-encoded within `log_name`. Example:
	// `"organizations/1234567890/logs/cloudresourcemanager.googleapis.com%2Factivity"`.
	// `[LOG_ID]` must be less than 512 characters long and can only include the
	// following characters: upper and lower case alphanumeric characters,
	// forward-slash, underscore, hyphen, and period.
	//
	// For backward compatibility, if `log_name` begins with a forward-slash, such
	// as `/projects/...`, then the log entry is ingested as usual but the
	// forward-slash is removed. Listing the log entry will not show the leading
	// slash and filtering for a log name with a leading slash will never return
	// any results.
	LogName string `protobuf:"bytes,12,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	// Required. The monitored resource that produced this log entry.
	//
	// Example: a log entry that reports a database error would be associated with
	// the monitored resource designating the particular database that reported
	// the error.
	Resource *monitoredres.MonitoredResource `protobuf:"bytes,8,opt,name=resource,proto3" json:"resource,omitempty"`
	// Optional. The log entry payload, which can be one of multiple types.
	//
	// Types that are valid to be assigned to Payload:
	//	*LogEntry_ProtoPayload
	//	*LogEntry_TextPayload
	//	*LogEntry_JsonPayload
	Payload isLogEntry_Payload `protobuf_oneof:"payload"`
	// Optional. The time the event described by the log entry occurred.  This
	// time is used to compute the log entry's age and to enforce the logs
	// retention period. If this field is omitted in a new log entry, then Logging
	// assigns it the current time.  Timestamps have nanosecond accuracy, but
	// trailing zeros in the fractional seconds might be omitted when the
	// timestamp is displayed.
	//
	// Incoming log entries should have timestamps that are no more than the [logs
	// retention period](/logging/quotas) in the past, and no more than 24 hours
	// in the future. Log entries outside those time boundaries will not be
	// available when calling `entries.list`, but those log entries can still be
	// [exported with LogSinks](/logging/docs/api/tasks/exporting-logs).
	Timestamp *timestamp.Timestamp `protobuf:"bytes,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Output only. The time the log entry was received by Logging.
	ReceiveTimestamp *timestamp.Timestamp `protobuf:"bytes,24,opt,name=receive_timestamp,json=receiveTimestamp,proto3" json:"receive_timestamp,omitempty"`
	// Optional. The severity of the log entry. The default value is
	// `LogSeverity.DEFAULT`.
	Severity _type.LogSeverity `protobuf:"varint,10,opt,name=severity,proto3,enum=google.logging.type.LogSeverity" json:"severity,omitempty"`
	// Optional. A unique identifier for the log entry. If you provide a value,
	// then Logging considers other log entries in the same project, with the same
	// `timestamp`, and with the same `insert_id` to be duplicates which can be
	// removed. If omitted in new log entries, then Logging assigns its own unique
	// identifier. The `insert_id` is also used to order log entries that have the
	// same `timestamp` value.
	InsertId string `protobuf:"bytes,4,opt,name=insert_id,json=insertId,proto3" json:"insert_id,omitempty"`
	// Optional. Information about the HTTP request associated with this log
	// entry, if applicable.
	HttpRequest *_type.HttpRequest `protobuf:"bytes,7,opt,name=http_request,json=httpRequest,proto3" json:"http_request,omitempty"`
	// Optional. A set of user-defined (key, value) data that provides additional
	// information about the log entry.
	Labels map[string]string `protobuf:"bytes,11,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Deprecated. Output only. Additional metadata about the monitored resource.
	//
	// Only `k8s_container`, `k8s_pod`, and `k8s_node` MonitoredResources have
	// this field populated for GKE versions older than 1.12.6. For GKE versions
	// 1.12.6 and above, the `metadata` field has been deprecated. The Kubernetes
	// pod labels that used to be in `metadata.userLabels` will now be present in
	// the `labels` field with a key prefix of `k8s-pod/`. The Stackdriver system
	// labels that were present in the `metadata.systemLabels` field will no
	// longer be available in the LogEntry.
	Metadata *monitoredres.MonitoredResourceMetadata `protobuf:"bytes,25,opt,name=metadata,proto3" json:"metadata,omitempty"` // Deprecated: Do not use.
	// Optional. Information about an operation associated with the log entry, if
	// applicable.
	Operation *LogEntryOperation `protobuf:"bytes,15,opt,name=operation,proto3" json:"operation,omitempty"`
	// Optional. Resource name of the trace associated with the log entry, if any.
	// If it contains a relative resource name, the name is assumed to be relative
	// to `//tracing.googleapis.com`. Example:
	// `projects/my-projectid/traces/06796866738c859f2f19b7cfb3214824`
	Trace string `protobuf:"bytes,22,opt,name=trace,proto3" json:"trace,omitempty"`
	// Optional. The span ID within the trace associated with the log entry.
	//
	// For Trace spans, this is the same format that the Trace API v2 uses: a
	// 16-character hexadecimal encoding of an 8-byte array, such as
	// <code>"000000000000004a"</code>.
	SpanId string `protobuf:"bytes,27,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// Optional. The sampling decision of the trace associated with the log entry.
	//
	// True means that the trace resource name in the `trace` field was sampled
	// for storage in a trace backend. False means that the trace was not sampled
	// for storage when this log entry was written, or the sampling decision was
	// unknown at the time. A non-sampled `trace` value is still useful as a
	// request correlation identifier. The default is False.
	TraceSampled bool `protobuf:"varint,30,opt,name=trace_sampled,json=traceSampled,proto3" json:"trace_sampled,omitempty"`
	// Optional. Source code location information associated with the log entry,
	// if any.
	SourceLocation       *LogEntrySourceLocation `protobuf:"bytes,23,opt,name=source_location,json=sourceLocation,proto3" json:"source_location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba2017251165146, []int{0}
}

func (m *LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntry.Unmarshal(m, b)
}
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntry.Marshal(b, m, deterministic)
}
func (m *LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry.Merge(m, src)
}
func (m *LogEntry) XXX_Size() int {
	return xxx_messageInfo_LogEntry.Size(m)
}
func (m *LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry proto.InternalMessageInfo

func (m *LogEntry) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *LogEntry) GetResource() *monitoredres.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type isLogEntry_Payload interface {
	isLogEntry_Payload()
}

type LogEntry_ProtoPayload struct {
	ProtoPayload *any.Any `protobuf:"bytes,2,opt,name=proto_payload,json=protoPayload,proto3,oneof"`
}

type LogEntry_TextPayload struct {
	TextPayload string `protobuf:"bytes,3,opt,name=text_payload,json=textPayload,proto3,oneof"`
}

type LogEntry_JsonPayload struct {
	JsonPayload *_struct.Struct `protobuf:"bytes,6,opt,name=json_payload,json=jsonPayload,proto3,oneof"`
}

func (*LogEntry_ProtoPayload) isLogEntry_Payload() {}

func (*LogEntry_TextPayload) isLogEntry_Payload() {}

func (*LogEntry_JsonPayload) isLogEntry_Payload() {}

func (m *LogEntry) GetPayload() isLogEntry_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *LogEntry) GetProtoPayload() *any.Any {
	if x, ok := m.GetPayload().(*LogEntry_ProtoPayload); ok {
		return x.ProtoPayload
	}
	return nil
}

func (m *LogEntry) GetTextPayload() string {
	if x, ok := m.GetPayload().(*LogEntry_TextPayload); ok {
		return x.TextPayload
	}
	return ""
}

func (m *LogEntry) GetJsonPayload() *_struct.Struct {
	if x, ok := m.GetPayload().(*LogEntry_JsonPayload); ok {
		return x.JsonPayload
	}
	return nil
}

func (m *LogEntry) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *LogEntry) GetReceiveTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.ReceiveTimestamp
	}
	return nil
}

func (m *LogEntry) GetSeverity() _type.LogSeverity {
	if m != nil {
		return m.Severity
	}
	return _type.LogSeverity_DEFAULT
}

func (m *LogEntry) GetInsertId() string {
	if m != nil {
		return m.InsertId
	}
	return ""
}

func (m *LogEntry) GetHttpRequest() *_type.HttpRequest {
	if m != nil {
		return m.HttpRequest
	}
	return nil
}

func (m *LogEntry) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// Deprecated: Do not use.
func (m *LogEntry) GetMetadata() *monitoredres.MonitoredResourceMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *LogEntry) GetOperation() *LogEntryOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *LogEntry) GetTrace() string {
	if m != nil {
		return m.Trace
	}
	return ""
}

func (m *LogEntry) GetSpanId() string {
	if m != nil {
		return m.SpanId
	}
	return ""
}

func (m *LogEntry) GetTraceSampled() bool {
	if m != nil {
		return m.TraceSampled
	}
	return false
}

func (m *LogEntry) GetSourceLocation() *LogEntrySourceLocation {
	if m != nil {
		return m.SourceLocation
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LogEntry) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LogEntry_ProtoPayload)(nil),
		(*LogEntry_TextPayload)(nil),
		(*LogEntry_JsonPayload)(nil),
	}
}

// Additional information about a potentially long-running operation with which
// a log entry is associated.
type LogEntryOperation struct {
	// Optional. An arbitrary operation identifier. Log entries with the same
	// identifier are assumed to be part of the same operation.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Optional. An arbitrary producer identifier. The combination of `id` and
	// `producer` must be globally unique. Examples for `producer`:
	// `"MyDivision.MyBigCompany.com"`, `"github.com/MyProject/MyApplication"`.
	Producer string `protobuf:"bytes,2,opt,name=producer,proto3" json:"producer,omitempty"`
	// Optional. Set this to True if this is the first log entry in the operation.
	First bool `protobuf:"varint,3,opt,name=first,proto3" json:"first,omitempty"`
	// Optional. Set this to True if this is the last log entry in the operation.
	Last                 bool     `protobuf:"varint,4,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogEntryOperation) Reset()         { *m = LogEntryOperation{} }
func (m *LogEntryOperation) String() string { return proto.CompactTextString(m) }
func (*LogEntryOperation) ProtoMessage()    {}
func (*LogEntryOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba2017251165146, []int{1}
}

func (m *LogEntryOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntryOperation.Unmarshal(m, b)
}
func (m *LogEntryOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntryOperation.Marshal(b, m, deterministic)
}
func (m *LogEntryOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntryOperation.Merge(m, src)
}
func (m *LogEntryOperation) XXX_Size() int {
	return xxx_messageInfo_LogEntryOperation.Size(m)
}
func (m *LogEntryOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntryOperation.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntryOperation proto.InternalMessageInfo

func (m *LogEntryOperation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LogEntryOperation) GetProducer() string {
	if m != nil {
		return m.Producer
	}
	return ""
}

func (m *LogEntryOperation) GetFirst() bool {
	if m != nil {
		return m.First
	}
	return false
}

func (m *LogEntryOperation) GetLast() bool {
	if m != nil {
		return m.Last
	}
	return false
}

// Additional information about the source code location that produced the log
// entry.
type LogEntrySourceLocation struct {
	// Optional. Source file name. Depending on the runtime environment, this
	// might be a simple name or a fully-qualified name.
	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	// Optional. Line within the source file. 1-based; 0 indicates no line number
	// available.
	Line int64 `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	// Optional. Human-readable name of the function or method being invoked, with
	// optional context such as the class or package name. This information may be
	// used in contexts such as the logs viewer, where a file and line number are
	// less meaningful. The format can vary by language. For example:
	// `qual.if.ied.Class.method` (Java), `dir/package.func` (Go), `function`
	// (Python).
	Function             string   `protobuf:"bytes,3,opt,name=function,proto3" json:"function,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogEntrySourceLocation) Reset()         { *m = LogEntrySourceLocation{} }
func (m *LogEntrySourceLocation) String() string { return proto.CompactTextString(m) }
func (*LogEntrySourceLocation) ProtoMessage()    {}
func (*LogEntrySourceLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba2017251165146, []int{2}
}

func (m *LogEntrySourceLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntrySourceLocation.Unmarshal(m, b)
}
func (m *LogEntrySourceLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntrySourceLocation.Marshal(b, m, deterministic)
}
func (m *LogEntrySourceLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntrySourceLocation.Merge(m, src)
}
func (m *LogEntrySourceLocation) XXX_Size() int {
	return xxx_messageInfo_LogEntrySourceLocation.Size(m)
}
func (m *LogEntrySourceLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntrySourceLocation.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntrySourceLocation proto.InternalMessageInfo

func (m *LogEntrySourceLocation) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *LogEntrySourceLocation) GetLine() int64 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *LogEntrySourceLocation) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func init() {
	proto.RegisterType((*LogEntry)(nil), "google.logging.v2.LogEntry")
	proto.RegisterMapType((map[string]string)(nil), "google.logging.v2.LogEntry.LabelsEntry")
	proto.RegisterType((*LogEntryOperation)(nil), "google.logging.v2.LogEntryOperation")
	proto.RegisterType((*LogEntrySourceLocation)(nil), "google.logging.v2.LogEntrySourceLocation")
}

func init() { proto.RegisterFile("google/logging/v2/log_entry.proto", fileDescriptor_8ba2017251165146) }

var fileDescriptor_8ba2017251165146 = []byte{
	// 785 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdb, 0x8e, 0xdb, 0x36,
	0x10, 0x8d, 0xec, 0x74, 0x2d, 0xd3, 0xde, 0x4d, 0x96, 0x48, 0x63, 0xad, 0xd3, 0x8b, 0xbb, 0xe9,
	0xc5, 0x7d, 0x91, 0x00, 0xf7, 0x65, 0xd3, 0x04, 0x28, 0xea, 0x20, 0xc8, 0x06, 0x70, 0xda, 0x80,
	0x2e, 0xf2, 0x50, 0x18, 0x10, 0xb8, 0x12, 0xad, 0xb0, 0x95, 0x48, 0x95, 0xa4, 0x8c, 0xfa, 0x53,
	0xfa, 0x0b, 0xfd, 0x94, 0xfe, 0x43, 0xff, 0xa3, 0x8f, 0x05, 0x47, 0x94, 0xbc, 0xb5, 0x03, 0xf7,
	0x6d, 0x86, 0x73, 0xce, 0xcc, 0xe1, 0xd1, 0xd0, 0x46, 0x9f, 0x65, 0x52, 0x66, 0x39, 0x8b, 0x72,
	0x99, 0x65, 0x5c, 0x64, 0xd1, 0x66, 0x66, 0xc3, 0x98, 0x09, 0xa3, 0xb6, 0x61, 0xa9, 0xa4, 0x91,
	0xf8, 0xbc, 0x86, 0x84, 0x0e, 0x12, 0x6e, 0x66, 0xe3, 0xc7, 0x8e, 0x45, 0x4b, 0x1e, 0x15, 0x52,
	0x70, 0x23, 0x15, 0x4b, 0x63, 0xc5, 0xb4, 0xac, 0x54, 0xc2, 0x6a, 0xde, 0xf8, 0xcb, 0xbd, 0xd6,
	0x66, 0x5b, 0xb2, 0xe8, 0x9d, 0x31, 0x65, 0xac, 0xd8, 0x6f, 0x15, 0xd3, 0xe6, 0x18, 0xce, 0x8a,
	0xd0, 0x6c, 0xc3, 0x14, 0x37, 0x4e, 0xc7, 0xf8, 0xc2, 0xe1, 0x20, 0xbb, 0xa9, 0xd6, 0x11, 0x15,
	0x4d, 0xe9, 0xa3, 0xfd, 0x92, 0x36, 0xaa, 0x4a, 0x9a, 0x01, 0x9f, 0xee, 0x57, 0x0d, 0x2f, 0x98,
	0x36, 0xb4, 0x28, 0x1d, 0x60, 0xe4, 0x00, 0xaa, 0x4c, 0x22, 0x6d, 0xa8, 0xa9, 0xf4, 0x5e, 0x5f,
	0x7b, 0x4f, 0x2a, 0x84, 0x34, 0xd4, 0x70, 0x29, 0x5c, 0xf5, 0xf2, 0xef, 0x1e, 0xf2, 0x17, 0x32,
	0x7b, 0x61, 0xbd, 0xc2, 0x17, 0xc8, 0xb7, 0x9a, 0x05, 0x2d, 0x58, 0x30, 0x9c, 0x78, 0xd3, 0x3e,
	0xe9, 0xe5, 0x32, 0xfb, 0x81, 0x16, 0x0c, 0x3f, 0x41, 0x7e, 0x63, 0x4d, 0xe0, 0x4f, 0xbc, 0xe9,
	0x60, 0xf6, 0x71, 0xe8, 0x3c, 0xa5, 0x25, 0x0f, 0x5f, 0x37, 0x06, 0x12, 0x07, 0x22, 0x2d, 0x1c,
	0x3f, 0x45, 0xa7, 0x30, 0x2b, 0x2e, 0xe9, 0x36, 0x97, 0x34, 0x0d, 0x3a, 0xc0, 0x7f, 0xd0, 0xf0,
	0x9b, 0x2b, 0x85, 0xdf, 0x8b, 0xed, 0xf5, 0x1d, 0x32, 0x84, 0xfc, 0x4d, 0x8d, 0xc5, 0x8f, 0xd1,
	0xd0, 0xb0, 0xdf, 0x4d, 0xcb, 0xed, 0x5a, 0x59, 0xd7, 0x77, 0xc8, 0xc0, 0x9e, 0x36, 0xa0, 0x67,
	0x68, 0xf8, 0x8b, 0x96, 0xa2, 0x05, 0x9d, 0xc0, 0x80, 0xd1, 0xc1, 0x80, 0x25, 0x38, 0x6a, 0xd9,
	0x16, 0xde, 0xb0, 0xaf, 0x50, 0xbf, 0x35, 0x33, 0xe8, 0x03, 0x75, 0x7c, 0x40, 0xfd, 0xa9, 0x41,
	0x90, 0x1d, 0x18, 0xbf, 0x44, 0xe7, 0x8a, 0x25, 0x8c, 0x6f, 0x58, 0xbc, 0xeb, 0x10, 0xfc, 0x6f,
	0x87, 0xfb, 0x8e, 0xd4, 0x9e, 0xe0, 0x67, 0xc8, 0x6f, 0x16, 0x25, 0x40, 0x13, 0x6f, 0x7a, 0x36,
	0x9b, 0x84, 0x7b, 0x1b, 0x6b, 0x37, 0x2a, 0x5c, 0xc8, 0x6c, 0xe9, 0x70, 0xa4, 0x65, 0xe0, 0x47,
	0xa8, 0xcf, 0x85, 0x66, 0xca, 0xc4, 0x3c, 0x0d, 0xee, 0xc2, 0x77, 0xf3, 0xeb, 0x83, 0x57, 0x29,
	0x7e, 0x8e, 0x86, 0xb7, 0xf7, 0x35, 0xe8, 0x81, 0xbc, 0xf7, 0xb7, 0xbf, 0x36, 0xa6, 0x24, 0x35,
	0x8e, 0x0c, 0xde, 0xed, 0x12, 0xfc, 0x1d, 0x3a, 0xc9, 0xe9, 0x0d, 0xcb, 0x75, 0x30, 0x98, 0x74,
	0xa7, 0x83, 0xd9, 0x57, 0xe1, 0xc1, 0x7b, 0x0a, 0x9b, 0x2d, 0x0a, 0x17, 0x80, 0x84, 0x98, 0x38,
	0x1a, 0x7e, 0x81, 0xfc, 0x82, 0x19, 0x9a, 0x52, 0x43, 0x83, 0x0b, 0x50, 0xf0, 0xc5, 0xd1, 0xf5,
	0x79, 0xed, 0xc0, 0xf3, 0x4e, 0xe0, 0x91, 0x96, 0x8a, 0xe7, 0xa8, 0x2f, 0x4b, 0xa6, 0x60, 0x83,
	0x83, 0x7b, 0xd0, 0xe7, 0xf3, 0x23, 0x52, 0x7e, 0x6c, 0xb0, 0x64, 0x47, 0xc3, 0x0f, 0xd0, 0x07,
	0x46, 0xd1, 0x84, 0x05, 0x0f, 0xc1, 0xa9, 0x3a, 0xc1, 0x23, 0xd4, 0xd3, 0x25, 0x15, 0xd6, 0xc1,
	0x47, 0x70, 0x7e, 0x62, 0xd3, 0x57, 0x76, 0x01, 0x4f, 0x01, 0x11, 0x6b, 0x5a, 0x94, 0x39, 0x4b,
	0x83, 0x4f, 0x26, 0xde, 0xd4, 0x27, 0x43, 0x38, 0x5c, 0xd6, 0x67, 0x98, 0xa0, 0x7b, 0xb5, 0xee,
	0x38, 0x97, 0x49, 0xad, 0x6e, 0x04, 0xea, 0xbe, 0x3e, 0xa2, 0x6e, 0x09, 0x8c, 0x85, 0x23, 0x90,
	0x33, 0xfd, 0x9f, 0x7c, 0xfc, 0x04, 0x0d, 0x6e, 0x39, 0x89, 0xef, 0xa3, 0xee, 0xaf, 0x6c, 0x1b,
	0x78, 0x20, 0xce, 0x86, 0xf6, 0x22, 0x1b, 0x9a, 0x57, 0x0c, 0xde, 0x53, 0x9f, 0xd4, 0xc9, 0xb7,
	0x9d, 0x2b, 0x6f, 0xde, 0x47, 0x3d, 0xf7, 0x14, 0x2e, 0x39, 0x3a, 0x3f, 0x70, 0x03, 0x9f, 0xa1,
	0x0e, 0x4f, 0x5d, 0xab, 0x0e, 0x4f, 0xf1, 0x18, 0xf9, 0xa5, 0x92, 0x69, 0x95, 0x30, 0xe5, 0x9a,
	0xb5, 0xb9, 0x9d, 0xb2, 0xe6, 0x4a, 0x1b, 0x78, 0x79, 0x3e, 0xa9, 0x13, 0x8c, 0xd1, 0xdd, 0x9c,
	0x6a, 0x03, 0xdb, 0xe6, 0x13, 0x88, 0x2f, 0x57, 0xe8, 0xe1, 0xfb, 0xaf, 0x66, 0xd1, 0x6b, 0x9e,
	0x33, 0x37, 0x11, 0x62, 0xe8, 0xc0, 0x45, 0x2d, 0xbe, 0x4b, 0x20, 0xb6, 0x3a, 0xd6, 0x95, 0x48,
	0xc0, 0xbf, 0x6e, 0xad, 0xa3, 0xc9, 0xe7, 0x7f, 0x78, 0xe8, 0xc3, 0x44, 0x16, 0x87, 0x7e, 0xce,
	0x4f, 0x9b, 0xa9, 0x6f, 0xe0, 0x87, 0xc3, 0xfb, 0xf9, 0xca, 0x61, 0x32, 0x99, 0x53, 0x91, 0x85,
	0x52, 0x65, 0x51, 0xc6, 0x04, 0x3c, 0xc4, 0xa8, 0x2e, 0xd1, 0x92, 0xeb, 0x5b, 0x7f, 0x17, 0x4f,
	0x5d, 0xf8, 0x8f, 0xe7, 0xfd, 0xd9, 0x19, 0xbd, 0xac, 0xd9, 0xcf, 0x73, 0x59, 0xa5, 0xf6, 0x63,
	0xc1, 0x9c, 0xb7, 0xb3, 0xbf, 0x9a, 0xca, 0x0a, 0x2a, 0x2b, 0x57, 0x59, 0xbd, 0x9d, 0xdd, 0x9c,
	0x40, 0xef, 0x6f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x32, 0xc1, 0x34, 0x19, 0x89, 0x06, 0x00,
	0x00,
}
