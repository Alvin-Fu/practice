// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/bigquery/datatransfer/v1/transfer.proto

package datatransfer

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// DEPRECATED. Represents data transfer type.
type TransferType int32 // Deprecated: Do not use.
const (
	// Invalid or Unknown transfer type placeholder.
	TransferType_TRANSFER_TYPE_UNSPECIFIED TransferType = 0
	// Batch data transfer.
	TransferType_BATCH TransferType = 1
	// Streaming data transfer. Streaming data source currently doesn't
	// support multiple transfer configs per project.
	TransferType_STREAMING TransferType = 2
)

var TransferType_name = map[int32]string{
	0: "TRANSFER_TYPE_UNSPECIFIED",
	1: "BATCH",
	2: "STREAMING",
}

var TransferType_value = map[string]int32{
	"TRANSFER_TYPE_UNSPECIFIED": 0,
	"BATCH":                     1,
	"STREAMING":                 2,
}

func (x TransferType) String() string {
	return proto.EnumName(TransferType_name, int32(x))
}

func (TransferType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_90c2574420b818ad, []int{0}
}

// Represents data transfer run state.
type TransferState int32

const (
	// State placeholder.
	TransferState_TRANSFER_STATE_UNSPECIFIED TransferState = 0
	// Data transfer is scheduled and is waiting to be picked up by
	// data transfer backend.
	TransferState_PENDING TransferState = 2
	// Data transfer is in progress.
	TransferState_RUNNING TransferState = 3
	// Data transfer completed successfully.
	TransferState_SUCCEEDED TransferState = 4
	// Data transfer failed.
	TransferState_FAILED TransferState = 5
	// Data transfer is cancelled.
	TransferState_CANCELLED TransferState = 6
)

var TransferState_name = map[int32]string{
	0: "TRANSFER_STATE_UNSPECIFIED",
	2: "PENDING",
	3: "RUNNING",
	4: "SUCCEEDED",
	5: "FAILED",
	6: "CANCELLED",
}

var TransferState_value = map[string]int32{
	"TRANSFER_STATE_UNSPECIFIED": 0,
	"PENDING":                    2,
	"RUNNING":                    3,
	"SUCCEEDED":                  4,
	"FAILED":                     5,
	"CANCELLED":                  6,
}

func (x TransferState) String() string {
	return proto.EnumName(TransferState_name, int32(x))
}

func (TransferState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_90c2574420b818ad, []int{1}
}

// Represents data transfer user facing message severity.
type TransferMessage_MessageSeverity int32

const (
	// No severity specified.
	TransferMessage_MESSAGE_SEVERITY_UNSPECIFIED TransferMessage_MessageSeverity = 0
	// Informational message.
	TransferMessage_INFO TransferMessage_MessageSeverity = 1
	// Warning message.
	TransferMessage_WARNING TransferMessage_MessageSeverity = 2
	// Error message.
	TransferMessage_ERROR TransferMessage_MessageSeverity = 3
)

var TransferMessage_MessageSeverity_name = map[int32]string{
	0: "MESSAGE_SEVERITY_UNSPECIFIED",
	1: "INFO",
	2: "WARNING",
	3: "ERROR",
}

var TransferMessage_MessageSeverity_value = map[string]int32{
	"MESSAGE_SEVERITY_UNSPECIFIED": 0,
	"INFO":                         1,
	"WARNING":                      2,
	"ERROR":                        3,
}

func (x TransferMessage_MessageSeverity) String() string {
	return proto.EnumName(TransferMessage_MessageSeverity_name, int32(x))
}

func (TransferMessage_MessageSeverity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_90c2574420b818ad, []int{3, 0}
}

// Options customizing the data transfer schedule.
type ScheduleOptions struct {
	// If true, automatic scheduling of data transfer runs for this configuration
	// will be disabled. The runs can be started on ad-hoc basis using
	// StartManualTransferRuns API. When automatic scheduling is disabled, the
	// TransferConfig.schedule field will be ignored.
	DisableAutoScheduling bool `protobuf:"varint,3,opt,name=disable_auto_scheduling,json=disableAutoScheduling,proto3" json:"disable_auto_scheduling,omitempty"`
	// Specifies time to start scheduling transfer runs. The first run will be
	// scheduled at or after the start time according to a recurrence pattern
	// defined in the schedule string. The start time can be changed at any
	// moment. The time when a data transfer can be trigerred manually is not
	// limited by this option.
	StartTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Defines time to stop scheduling transfer runs. A transfer run cannot be
	// scheduled at or after the end time. The end time can be changed at any
	// moment. The time when a data transfer can be trigerred manually is not
	// limited by this option.
	EndTime              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ScheduleOptions) Reset()         { *m = ScheduleOptions{} }
func (m *ScheduleOptions) String() string { return proto.CompactTextString(m) }
func (*ScheduleOptions) ProtoMessage()    {}
func (*ScheduleOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_90c2574420b818ad, []int{0}
}

func (m *ScheduleOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduleOptions.Unmarshal(m, b)
}
func (m *ScheduleOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduleOptions.Marshal(b, m, deterministic)
}
func (m *ScheduleOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleOptions.Merge(m, src)
}
func (m *ScheduleOptions) XXX_Size() int {
	return xxx_messageInfo_ScheduleOptions.Size(m)
}
func (m *ScheduleOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleOptions proto.InternalMessageInfo

func (m *ScheduleOptions) GetDisableAutoScheduling() bool {
	if m != nil {
		return m.DisableAutoScheduling
	}
	return false
}

func (m *ScheduleOptions) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ScheduleOptions) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// Represents a data transfer configuration. A transfer configuration
// contains all metadata needed to perform a data transfer. For example,
// `destination_dataset_id` specifies where data should be stored.
// When a new transfer configuration is created, the specified
// `destination_dataset_id` is created when needed and shared with the
// appropriate data source service account.
type TransferConfig struct {
	// The resource name of the transfer config.
	// Transfer config names have the form of
	// `projects/{project_id}/locations/{region}/transferConfigs/{config_id}`.
	// The name is automatically generated based on the config_id specified in
	// CreateTransferConfigRequest along with project_id and region. If config_id
	// is not provided, usually a uuid, even though it is not guaranteed or
	// required, will be generated for config_id.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The desination of the transfer config.
	//
	// Types that are valid to be assigned to Destination:
	//	*TransferConfig_DestinationDatasetId
	Destination isTransferConfig_Destination `protobuf_oneof:"destination"`
	// User specified display name for the data transfer.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Data source id. Cannot be changed once data transfer is created.
	DataSourceId string `protobuf:"bytes,5,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	// Data transfer specific parameters.
	Params *_struct.Struct `protobuf:"bytes,9,opt,name=params,proto3" json:"params,omitempty"`
	// Data transfer schedule.
	// If the data source does not support a custom schedule, this should be
	// empty. If it is empty, the default value for the data source will be
	// used.
	// The specified times are in UTC.
	// Examples of valid format:
	// `1st,3rd monday of month 15:30`,
	// `every wed,fri of jan,jun 13:15`, and
	// `first sunday of quarter 00:00`.
	// See more explanation about the format here:
	// https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
	// NOTE: the granularity should be at least 8 hours, or less frequent.
	Schedule string `protobuf:"bytes,7,opt,name=schedule,proto3" json:"schedule,omitempty"`
	// Options customizing the data transfer schedule.
	ScheduleOptions *ScheduleOptions `protobuf:"bytes,24,opt,name=schedule_options,json=scheduleOptions,proto3" json:"schedule_options,omitempty"`
	// The number of days to look back to automatically refresh the data.
	// For example, if `data_refresh_window_days = 10`, then every day
	// BigQuery reingests data for [today-10, today-1], rather than ingesting data
	// for just [today-1].
	// Only valid if the data source supports the feature. Set the value to  0
	// to use the default value.
	DataRefreshWindowDays int32 `protobuf:"varint,12,opt,name=data_refresh_window_days,json=dataRefreshWindowDays,proto3" json:"data_refresh_window_days,omitempty"`
	// Is this config disabled. When set to true, no runs are scheduled
	// for a given transfer.
	Disabled bool `protobuf:"varint,13,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Output only. Data transfer modification time. Ignored by server on input.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Next time when data transfer will run.
	NextRunTime *timestamp.Timestamp `protobuf:"bytes,8,opt,name=next_run_time,json=nextRunTime,proto3" json:"next_run_time,omitempty"`
	// Output only. State of the most recently updated transfer run.
	State TransferState `protobuf:"varint,10,opt,name=state,proto3,enum=google.cloud.bigquery.datatransfer.v1.TransferState" json:"state,omitempty"`
	// Deprecated. Unique ID of the user on whose behalf transfer is done.
	UserId int64 `protobuf:"varint,11,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Output only. Region in which BigQuery dataset is located.
	DatasetRegion        string   `protobuf:"bytes,14,opt,name=dataset_region,json=datasetRegion,proto3" json:"dataset_region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferConfig) Reset()         { *m = TransferConfig{} }
func (m *TransferConfig) String() string { return proto.CompactTextString(m) }
func (*TransferConfig) ProtoMessage()    {}
func (*TransferConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_90c2574420b818ad, []int{1}
}

func (m *TransferConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferConfig.Unmarshal(m, b)
}
func (m *TransferConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferConfig.Marshal(b, m, deterministic)
}
func (m *TransferConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferConfig.Merge(m, src)
}
func (m *TransferConfig) XXX_Size() int {
	return xxx_messageInfo_TransferConfig.Size(m)
}
func (m *TransferConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TransferConfig proto.InternalMessageInfo

func (m *TransferConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isTransferConfig_Destination interface {
	isTransferConfig_Destination()
}

type TransferConfig_DestinationDatasetId struct {
	DestinationDatasetId string `protobuf:"bytes,2,opt,name=destination_dataset_id,json=destinationDatasetId,proto3,oneof"`
}

func (*TransferConfig_DestinationDatasetId) isTransferConfig_Destination() {}

func (m *TransferConfig) GetDestination() isTransferConfig_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *TransferConfig) GetDestinationDatasetId() string {
	if x, ok := m.GetDestination().(*TransferConfig_DestinationDatasetId); ok {
		return x.DestinationDatasetId
	}
	return ""
}

func (m *TransferConfig) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *TransferConfig) GetDataSourceId() string {
	if m != nil {
		return m.DataSourceId
	}
	return ""
}

func (m *TransferConfig) GetParams() *_struct.Struct {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *TransferConfig) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

func (m *TransferConfig) GetScheduleOptions() *ScheduleOptions {
	if m != nil {
		return m.ScheduleOptions
	}
	return nil
}

func (m *TransferConfig) GetDataRefreshWindowDays() int32 {
	if m != nil {
		return m.DataRefreshWindowDays
	}
	return 0
}

func (m *TransferConfig) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *TransferConfig) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *TransferConfig) GetNextRunTime() *timestamp.Timestamp {
	if m != nil {
		return m.NextRunTime
	}
	return nil
}

func (m *TransferConfig) GetState() TransferState {
	if m != nil {
		return m.State
	}
	return TransferState_TRANSFER_STATE_UNSPECIFIED
}

func (m *TransferConfig) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *TransferConfig) GetDatasetRegion() string {
	if m != nil {
		return m.DatasetRegion
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TransferConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TransferConfig_DestinationDatasetId)(nil),
	}
}

// Represents a data transfer run.
type TransferRun struct {
	// The resource name of the transfer run.
	// Transfer run names have the form
	// `projects/{project_id}/locations/{location}/transferConfigs/{config_id}/runs/{run_id}`.
	// The name is ignored when creating a transfer run.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Minimum time after which a transfer run can be started.
	ScheduleTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	// For batch transfer runs, specifies the date and time of the data should be
	// ingested.
	RunTime *timestamp.Timestamp `protobuf:"bytes,10,opt,name=run_time,json=runTime,proto3" json:"run_time,omitempty"`
	// Status of the transfer run.
	ErrorStatus *status.Status `protobuf:"bytes,21,opt,name=error_status,json=errorStatus,proto3" json:"error_status,omitempty"`
	// Output only. Time when transfer run was started.
	// Parameter ignored by server for input requests.
	StartTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Output only. Time when transfer run ended.
	// Parameter ignored by server for input requests.
	EndTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Output only. Last time the data transfer run state was updated.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Data transfer specific parameters.
	Params *_struct.Struct `protobuf:"bytes,9,opt,name=params,proto3" json:"params,omitempty"`
	// Data transfer destination.
	//
	// Types that are valid to be assigned to Destination:
	//	*TransferRun_DestinationDatasetId
	Destination isTransferRun_Destination `protobuf_oneof:"destination"`
	// Output only. Data source id.
	DataSourceId string `protobuf:"bytes,7,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	// Data transfer run state. Ignored for input requests.
	State TransferState `protobuf:"varint,8,opt,name=state,proto3,enum=google.cloud.bigquery.datatransfer.v1.TransferState" json:"state,omitempty"`
	// Deprecated. Unique ID of the user on whose behalf transfer is done.
	UserId int64 `protobuf:"varint,11,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Output only. Describes the schedule of this transfer run if it was
	// created as part of a regular schedule. For batch transfer runs that are
	// scheduled manually, this is empty.
	// NOTE: the system might choose to delay the schedule depending on the
	// current load, so `schedule_time` doesn't always match this.
	Schedule             string   `protobuf:"bytes,12,opt,name=schedule,proto3" json:"schedule,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferRun) Reset()         { *m = TransferRun{} }
func (m *TransferRun) String() string { return proto.CompactTextString(m) }
func (*TransferRun) ProtoMessage()    {}
func (*TransferRun) Descriptor() ([]byte, []int) {
	return fileDescriptor_90c2574420b818ad, []int{2}
}

func (m *TransferRun) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferRun.Unmarshal(m, b)
}
func (m *TransferRun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferRun.Marshal(b, m, deterministic)
}
func (m *TransferRun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferRun.Merge(m, src)
}
func (m *TransferRun) XXX_Size() int {
	return xxx_messageInfo_TransferRun.Size(m)
}
func (m *TransferRun) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferRun.DiscardUnknown(m)
}

var xxx_messageInfo_TransferRun proto.InternalMessageInfo

func (m *TransferRun) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TransferRun) GetScheduleTime() *timestamp.Timestamp {
	if m != nil {
		return m.ScheduleTime
	}
	return nil
}

func (m *TransferRun) GetRunTime() *timestamp.Timestamp {
	if m != nil {
		return m.RunTime
	}
	return nil
}

func (m *TransferRun) GetErrorStatus() *status.Status {
	if m != nil {
		return m.ErrorStatus
	}
	return nil
}

func (m *TransferRun) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TransferRun) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *TransferRun) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *TransferRun) GetParams() *_struct.Struct {
	if m != nil {
		return m.Params
	}
	return nil
}

type isTransferRun_Destination interface {
	isTransferRun_Destination()
}

type TransferRun_DestinationDatasetId struct {
	DestinationDatasetId string `protobuf:"bytes,2,opt,name=destination_dataset_id,json=destinationDatasetId,proto3,oneof"`
}

func (*TransferRun_DestinationDatasetId) isTransferRun_Destination() {}

func (m *TransferRun) GetDestination() isTransferRun_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *TransferRun) GetDestinationDatasetId() string {
	if x, ok := m.GetDestination().(*TransferRun_DestinationDatasetId); ok {
		return x.DestinationDatasetId
	}
	return ""
}

func (m *TransferRun) GetDataSourceId() string {
	if m != nil {
		return m.DataSourceId
	}
	return ""
}

func (m *TransferRun) GetState() TransferState {
	if m != nil {
		return m.State
	}
	return TransferState_TRANSFER_STATE_UNSPECIFIED
}

func (m *TransferRun) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *TransferRun) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TransferRun) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TransferRun_DestinationDatasetId)(nil),
	}
}

// Represents a user facing message for a particular data transfer run.
type TransferMessage struct {
	// Time when message was logged.
	MessageTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=message_time,json=messageTime,proto3" json:"message_time,omitempty"`
	// Message severity.
	Severity TransferMessage_MessageSeverity `protobuf:"varint,2,opt,name=severity,proto3,enum=google.cloud.bigquery.datatransfer.v1.TransferMessage_MessageSeverity" json:"severity,omitempty"`
	// Message text.
	MessageText          string   `protobuf:"bytes,3,opt,name=message_text,json=messageText,proto3" json:"message_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferMessage) Reset()         { *m = TransferMessage{} }
func (m *TransferMessage) String() string { return proto.CompactTextString(m) }
func (*TransferMessage) ProtoMessage()    {}
func (*TransferMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_90c2574420b818ad, []int{3}
}

func (m *TransferMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferMessage.Unmarshal(m, b)
}
func (m *TransferMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferMessage.Marshal(b, m, deterministic)
}
func (m *TransferMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferMessage.Merge(m, src)
}
func (m *TransferMessage) XXX_Size() int {
	return xxx_messageInfo_TransferMessage.Size(m)
}
func (m *TransferMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TransferMessage proto.InternalMessageInfo

func (m *TransferMessage) GetMessageTime() *timestamp.Timestamp {
	if m != nil {
		return m.MessageTime
	}
	return nil
}

func (m *TransferMessage) GetSeverity() TransferMessage_MessageSeverity {
	if m != nil {
		return m.Severity
	}
	return TransferMessage_MESSAGE_SEVERITY_UNSPECIFIED
}

func (m *TransferMessage) GetMessageText() string {
	if m != nil {
		return m.MessageText
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.cloud.bigquery.datatransfer.v1.TransferType", TransferType_name, TransferType_value)
	proto.RegisterEnum("google.cloud.bigquery.datatransfer.v1.TransferState", TransferState_name, TransferState_value)
	proto.RegisterEnum("google.cloud.bigquery.datatransfer.v1.TransferMessage_MessageSeverity", TransferMessage_MessageSeverity_name, TransferMessage_MessageSeverity_value)
	proto.RegisterType((*ScheduleOptions)(nil), "google.cloud.bigquery.datatransfer.v1.ScheduleOptions")
	proto.RegisterType((*TransferConfig)(nil), "google.cloud.bigquery.datatransfer.v1.TransferConfig")
	proto.RegisterType((*TransferRun)(nil), "google.cloud.bigquery.datatransfer.v1.TransferRun")
	proto.RegisterType((*TransferMessage)(nil), "google.cloud.bigquery.datatransfer.v1.TransferMessage")
}

func init() {
	proto.RegisterFile("google/cloud/bigquery/datatransfer/v1/transfer.proto", fileDescriptor_90c2574420b818ad)
}

var fileDescriptor_90c2574420b818ad = []byte{
	// 1151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x0e, 0x25, 0xeb, 0x6f, 0x28, 0xd9, 0xc2, 0xa2, 0xa9, 0x19, 0x21, 0x6d, 0x54, 0xa3, 0x01,
	0x9c, 0x1c, 0x48, 0xd8, 0x75, 0x52, 0x34, 0x41, 0x7f, 0x28, 0x89, 0x76, 0xd4, 0x26, 0xb2, 0xb3,
	0xa4, 0x13, 0xa4, 0x30, 0x40, 0x50, 0xe2, 0x9a, 0x61, 0x20, 0x91, 0xec, 0x2e, 0xe9, 0x44, 0x08,
	0x72, 0xee, 0xa1, 0x6f, 0xd1, 0x63, 0x0f, 0x7d, 0x81, 0x9e, 0x7b, 0xe9, 0xa5, 0xaf, 0xd0, 0x5b,
	0x81, 0x3c, 0x45, 0xb1, 0x4b, 0x52, 0x91, 0xe5, 0xb4, 0x92, 0x7b, 0xd2, 0xce, 0xce, 0x7c, 0xdf,
	0xce, 0xee, 0x7c, 0x33, 0x14, 0xec, 0x79, 0x61, 0xe8, 0x8d, 0x89, 0x36, 0x1a, 0x87, 0x89, 0xab,
	0x0d, 0x7d, 0xef, 0x87, 0x84, 0xd0, 0xa9, 0xe6, 0x3a, 0xb1, 0x13, 0x53, 0x27, 0x60, 0xa7, 0x84,
	0x6a, 0x67, 0x3b, 0x5a, 0xbe, 0x56, 0x23, 0x1a, 0xc6, 0x21, 0xba, 0x99, 0xa2, 0x54, 0x81, 0x52,
	0x73, 0x94, 0x3a, 0x8f, 0x52, 0xcf, 0x76, 0x5a, 0x37, 0x32, 0x72, 0x27, 0xf2, 0xb5, 0x53, 0x9f,
	0x8c, 0x5d, 0x7b, 0x48, 0x9e, 0x3b, 0x67, 0x7e, 0x98, 0xf1, 0xb4, 0xae, 0xcd, 0x05, 0x50, 0xc2,
	0xc2, 0x84, 0x8e, 0x48, 0xe6, 0xba, 0x9e, 0xb9, 0x84, 0x35, 0x4c, 0x4e, 0x35, 0x16, 0xd3, 0x64,
	0x14, 0x67, 0xde, 0x1b, 0x8b, 0xde, 0xd8, 0x9f, 0x10, 0x16, 0x3b, 0x93, 0x28, 0x0b, 0xd8, 0xcc,
	0x02, 0x68, 0x34, 0xd2, 0x58, 0xec, 0xc4, 0x09, 0x4b, 0x1d, 0x5b, 0xbf, 0x49, 0xb0, 0x61, 0x8e,
	0x9e, 0x13, 0x37, 0x19, 0x93, 0xc3, 0x28, 0xf6, 0xc3, 0x80, 0xa1, 0xbb, 0xb0, 0xe9, 0xfa, 0xcc,
	0x19, 0x8e, 0x89, 0xed, 0x24, 0x71, 0x68, 0xb3, 0xd4, 0xef, 0x07, 0x9e, 0x52, 0x6c, 0x4b, 0xdb,
	0x55, 0x7c, 0x35, 0x73, 0xeb, 0x49, 0x1c, 0x9a, 0x33, 0x27, 0xfa, 0x02, 0x80, 0xc5, 0x0e, 0x8d,
	0x6d, 0x7e, 0xba, 0x22, 0xb5, 0xa5, 0x6d, 0x79, 0xb7, 0xa5, 0x66, 0x6f, 0x93, 0xa7, 0xa6, 0x5a,
	0x79, 0x6a, 0xb8, 0x26, 0xa2, 0xb9, 0x8d, 0xee, 0x40, 0x95, 0x04, 0x6e, 0x0a, 0x2c, 0x2c, 0x05,
	0x56, 0x48, 0xe0, 0x72, 0x6b, 0xeb, 0xf7, 0x32, 0xac, 0x5b, 0xd9, 0x0b, 0x77, 0xc3, 0xe0, 0xd4,
	0xf7, 0x10, 0x82, 0xb5, 0xc0, 0xc9, 0x8e, 0xaf, 0x61, 0xb1, 0x46, 0x77, 0xe1, 0x43, 0x97, 0xb0,
	0xd8, 0x0f, 0x1c, 0x7e, 0x41, 0x9b, 0xd7, 0x85, 0x91, 0xd8, 0xf6, 0x5d, 0x71, 0x56, 0xed, 0xc1,
	0x15, 0xfc, 0xc1, 0x9c, 0xbf, 0x97, 0xba, 0xfb, 0x2e, 0xfa, 0x04, 0xea, 0xae, 0xcf, 0xa2, 0xb1,
	0x33, 0xb5, 0x05, 0x67, 0x51, 0x70, 0xca, 0xd9, 0xde, 0x80, 0x53, 0x7f, 0x0a, 0xeb, 0x9c, 0xce,
	0x4e, 0x8b, 0xc5, 0x29, 0x4b, 0x22, 0xa8, 0xce, 0x77, 0x4d, 0xb1, 0xd9, 0x77, 0x91, 0x06, 0xe5,
	0xc8, 0xa1, 0xce, 0x84, 0x29, 0x35, 0x71, 0xb9, 0xcd, 0x0b, 0x97, 0x33, 0x45, 0x39, 0x71, 0x16,
	0x86, 0x5a, 0x50, 0xcd, 0x5e, 0x9d, 0x28, 0x15, 0x41, 0x38, 0xb3, 0x91, 0x03, 0xcd, 0x7c, 0x6d,
	0x87, 0x69, 0xc9, 0x14, 0x45, 0xd0, 0xde, 0x55, 0x57, 0x12, 0xa2, 0xba, 0x50, 0x70, 0xbc, 0xc1,
	0x16, 0x14, 0xf0, 0x39, 0x28, 0xe2, 0x56, 0x94, 0x9c, 0x52, 0xc2, 0x9e, 0xdb, 0x2f, 0xfd, 0xc0,
	0x0d, 0x5f, 0xda, 0xae, 0x33, 0x65, 0x4a, 0xbd, 0x2d, 0x6d, 0x97, 0xf0, 0x55, 0xee, 0xc7, 0xa9,
	0xfb, 0xa9, 0xf0, 0xf6, 0x9c, 0xa9, 0xc8, 0x3b, 0xd3, 0x86, 0xab, 0x34, 0x84, 0x56, 0x66, 0x36,
	0xfa, 0x06, 0xe4, 0x24, 0x72, 0x9d, 0x98, 0xa4, 0x65, 0x5e, 0x5b, 0x56, 0xe6, 0x4e, 0xf1, 0x2f,
	0xbd, 0x88, 0x21, 0xc5, 0x08, 0x95, 0x74, 0xa1, 0x11, 0x90, 0x57, 0xb1, 0x4d, 0x93, 0x20, 0xe5,
	0xa8, 0xae, 0xc6, 0x21, 0x73, 0x14, 0x4e, 0x02, 0x41, 0x72, 0x08, 0x25, 0xde, 0x01, 0x44, 0x81,
	0xb6, 0xb4, 0xbd, 0xbe, 0xbb, 0xb7, 0xe2, 0x9b, 0xe5, 0x32, 0x33, 0x39, 0x36, 0xa5, 0x4d, 0x79,
	0xd0, 0x26, 0x54, 0x12, 0x46, 0x28, 0xaf, 0xbd, 0xdc, 0x96, 0xb6, 0x8b, 0xb8, 0xcc, 0xcd, 0xbe,
	0x8b, 0x6e, 0xa7, 0xda, 0xe0, 0x52, 0xa3, 0xc4, 0xf3, 0xc3, 0x40, 0x59, 0xe7, 0xa5, 0x4c, 0xc1,
	0x8d, 0xcc, 0x85, 0x85, 0xe7, 0xde, 0xe4, 0xad, 0xfe, 0x02, 0x76, 0xf3, 0xd3, 0xcf, 0x1d, 0x9e,
	0xe6, 0xe7, 0x44, 0x3e, 0x53, 0x47, 0xe1, 0x44, 0x5b, 0xd0, 0xfb, 0x5e, 0x44, 0xc3, 0x17, 0x64,
	0x14, 0x33, 0xed, 0x75, 0xb6, 0x7a, 0x33, 0x1b, 0x4f, 0x69, 0x08, 0xd3, 0x5e, 0xe7, 0x1b, 0xf6,
	0x48, 0xec, 0xbc, 0xe9, 0x34, 0x40, 0x9e, 0x53, 0xfc, 0xd6, 0x9f, 0x65, 0x90, 0x73, 0x5e, 0x9c,
	0x04, 0xef, 0x6d, 0xa2, 0xaf, 0xa1, 0x31, 0x93, 0x9d, 0x78, 0xfc, 0xe2, 0xd2, 0x3e, 0xad, 0xe7,
	0x80, 0xbc, 0xc7, 0x67, 0x85, 0x83, 0xe5, 0x3d, 0x4e, 0xb3, 0x7a, 0xdd, 0x81, 0x3a, 0xa1, 0x34,
	0xa4, 0x76, 0x3a, 0xb7, 0x94, 0xab, 0x02, 0x8a, 0x72, 0x28, 0x8d, 0x46, 0xaa, 0x29, 0x3c, 0x58,
	0x16, 0x71, 0xa9, 0x81, 0xbe, 0x3a, 0x37, 0x8c, 0x56, 0x14, 0xdb, 0xdc, 0x44, 0xba, 0x37, 0x37,
	0x91, 0x4a, 0xab, 0xa1, 0xf3, 0xb1, 0xb4, 0xa8, 0xf4, 0xf2, 0xe5, 0x95, 0xbe, 0xb7, 0xe2, 0xc0,
	0x48, 0x91, 0xf9, 0xd4, 0xb8, 0xff, 0xdf, 0x73, 0x4e, 0x04, 0xff, 0xeb, 0xb0, 0xbb, 0x75, 0x61,
	0x92, 0x55, 0xde, 0xa9, 0xf5, 0xfc, 0x38, 0xfb, 0x36, 0x6f, 0xa1, 0xea, 0xff, 0x6f, 0xa1, 0xa5,
	0xdd, 0x73, 0x63, 0x6e, 0x04, 0xd6, 0xdf, 0x65, 0x32, 0xdb, 0xbc, 0xf7, 0x93, 0xf4, 0x56, 0xff,
	0x51, 0x02, 0xed, 0x32, 0x4d, 0xc3, 0xc5, 0x6d, 0xbd, 0xa7, 0x63, 0xc6, 0xe1, 0x48, 0x3c, 0x07,
	0xd3, 0x5e, 0xe7, 0xcb, 0x15, 0xda, 0x48, 0xa3, 0x09, 0x47, 0xd0, 0x24, 0xb8, 0xd0, 0x51, 0xbf,
	0x16, 0x60, 0x23, 0x3f, 0xf4, 0x11, 0x61, 0xcc, 0xf1, 0x08, 0xfa, 0x12, 0xea, 0x93, 0x74, 0xb9,
	0xea, 0x17, 0x52, 0xce, 0xe2, 0x85, 0x26, 0x86, 0x50, 0x65, 0xe4, 0x8c, 0x50, 0x3f, 0x9e, 0x8a,
	0x7a, 0xae, 0xef, 0xee, 0x5f, 0xf2, 0xe1, 0xb3, 0x44, 0xd4, 0xec, 0xd7, 0xcc, 0xd8, 0xf0, 0x8c,
	0x97, 0x7f, 0xf1, 0x66, 0x29, 0x92, 0x57, 0x71, 0xfe, 0xc5, 0xcb, 0xd3, 0x20, 0xaf, 0xe2, 0xad,
	0x63, 0xd8, 0x58, 0xc0, 0xa3, 0x36, 0x5c, 0x7f, 0x64, 0x98, 0xa6, 0x7e, 0x60, 0xd8, 0xa6, 0xf1,
	0xc4, 0xc0, 0x7d, 0xeb, 0x99, 0x7d, 0x3c, 0x30, 0x8f, 0x8c, 0x6e, 0x7f, 0xbf, 0x6f, 0xf4, 0x9a,
	0x57, 0x50, 0x15, 0xd6, 0xfa, 0x83, 0xfd, 0xc3, 0xa6, 0x84, 0x64, 0xa8, 0x3c, 0xd5, 0xf1, 0xa0,
	0x3f, 0x38, 0x68, 0x16, 0x50, 0x0d, 0x4a, 0x06, 0xc6, 0x87, 0xb8, 0x59, 0xbc, 0xfd, 0x1d, 0xd4,
	0xf3, 0x34, 0xad, 0x69, 0x44, 0xd0, 0x47, 0x70, 0xcd, 0xc2, 0xfa, 0xc0, 0xdc, 0x37, 0xb0, 0x6d,
	0x3d, 0x3b, 0x32, 0x16, 0x08, 0x6b, 0x50, 0xea, 0xe8, 0x56, 0xf7, 0x41, 0x53, 0x42, 0x0d, 0xa8,
	0x99, 0x16, 0x36, 0xf4, 0x47, 0x82, 0xb3, 0x55, 0x50, 0xa4, 0xdb, 0x0c, 0x1a, 0xe7, 0xc4, 0x86,
	0x3e, 0x86, 0xd6, 0x8c, 0xcd, 0xb4, 0x74, 0x6b, 0x91, 0x4e, 0x86, 0xca, 0x91, 0x31, 0xe8, 0xa5,
	0x59, 0xc9, 0x50, 0xc1, 0xc7, 0x03, 0x91, 0x62, 0x51, 0xb0, 0x1f, 0x77, 0xbb, 0x86, 0xd1, 0x33,
	0x7a, 0xcd, 0x35, 0x04, 0x50, 0xde, 0xd7, 0xfb, 0x0f, 0x8d, 0x5e, 0xb3, 0xc4, 0x5d, 0x5d, 0x7d,
	0xd0, 0x35, 0x1e, 0x72, 0xb3, 0xdc, 0xf9, 0x5b, 0x82, 0x5b, 0xa3, 0x70, 0xb2, 0x5a, 0x4d, 0x3a,
	0xb3, 0x04, 0x8f, 0x78, 0xd9, 0x8f, 0xa4, 0xef, 0x1f, 0x67, 0x38, 0x2f, 0x1c, 0x3b, 0x81, 0xa7,
	0x86, 0xd4, 0xd3, 0x3c, 0x12, 0x08, 0x51, 0x68, 0xef, 0xd4, 0xbc, 0xe4, 0x9f, 0xe9, 0xfd, 0x79,
	0xfb, 0xe7, 0x42, 0xe9, 0xa0, 0xdb, 0xe9, 0x59, 0xbf, 0x14, 0x6e, 0x1e, 0xa4, 0xdc, 0x5d, 0x91,
	0x53, 0xc7, 0xf7, 0x1e, 0x8b, 0x9c, 0x78, 0xf3, 0xe7, 0x69, 0xa8, 0x4f, 0x76, 0xfe, 0xc8, 0xe3,
	0x4e, 0x44, 0xdc, 0x49, 0x1e, 0x77, 0x32, 0x1f, 0x77, 0xf2, 0x64, 0x67, 0x58, 0x16, 0x59, 0x7d,
	0xf6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xf7, 0x51, 0xb3, 0x2e, 0x0b, 0x00, 0x00,
}
