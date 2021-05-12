// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/conversion_adjustment_upload_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for
// [ConversionAdjustmentUploadService.UploadConversionAdjustments][google.ads.googleads.v1.services.ConversionAdjustmentUploadService.UploadConversionAdjustments].
type UploadConversionAdjustmentsRequest struct {
	// The ID of the customer performing the upload.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The conversion adjustments that are being uploaded.
	ConversionAdjustments []*ConversionAdjustment `protobuf:"bytes,2,rep,name=conversion_adjustments,json=conversionAdjustments,proto3" json:"conversion_adjustments,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried out
	// in one transaction if and only if they are all valid. This should always be
	// set to true.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadConversionAdjustmentsRequest) Reset()         { *m = UploadConversionAdjustmentsRequest{} }
func (m *UploadConversionAdjustmentsRequest) String() string { return proto.CompactTextString(m) }
func (*UploadConversionAdjustmentsRequest) ProtoMessage()    {}
func (*UploadConversionAdjustmentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_418d7f048cc75833, []int{0}
}

func (m *UploadConversionAdjustmentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadConversionAdjustmentsRequest.Unmarshal(m, b)
}
func (m *UploadConversionAdjustmentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadConversionAdjustmentsRequest.Marshal(b, m, deterministic)
}
func (m *UploadConversionAdjustmentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadConversionAdjustmentsRequest.Merge(m, src)
}
func (m *UploadConversionAdjustmentsRequest) XXX_Size() int {
	return xxx_messageInfo_UploadConversionAdjustmentsRequest.Size(m)
}
func (m *UploadConversionAdjustmentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadConversionAdjustmentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadConversionAdjustmentsRequest proto.InternalMessageInfo

func (m *UploadConversionAdjustmentsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *UploadConversionAdjustmentsRequest) GetConversionAdjustments() []*ConversionAdjustment {
	if m != nil {
		return m.ConversionAdjustments
	}
	return nil
}

func (m *UploadConversionAdjustmentsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *UploadConversionAdjustmentsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// Response message for
// [ConversionAdjustmentUploadService.UploadConversionAdjustments][google.ads.googleads.v1.services.ConversionAdjustmentUploadService.UploadConversionAdjustments].
type UploadConversionAdjustmentsResponse struct {
	// Errors that pertain to conversion adjustment failures in the partial
	// failure mode. Returned when all errors occur inside the adjustments. If any
	// errors occur outside the adjustments (e.g. auth errors), we return an RPC
	// level error.
	PartialFailureError *status.Status `protobuf:"bytes,1,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// Returned for successfully processed conversion adjustments. Proto will be
	// empty for rows that received an error. Results are not returned when
	// validate_only is true.
	Results              []*ConversionAdjustmentResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *UploadConversionAdjustmentsResponse) Reset()         { *m = UploadConversionAdjustmentsResponse{} }
func (m *UploadConversionAdjustmentsResponse) String() string { return proto.CompactTextString(m) }
func (*UploadConversionAdjustmentsResponse) ProtoMessage()    {}
func (*UploadConversionAdjustmentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_418d7f048cc75833, []int{1}
}

func (m *UploadConversionAdjustmentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadConversionAdjustmentsResponse.Unmarshal(m, b)
}
func (m *UploadConversionAdjustmentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadConversionAdjustmentsResponse.Marshal(b, m, deterministic)
}
func (m *UploadConversionAdjustmentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadConversionAdjustmentsResponse.Merge(m, src)
}
func (m *UploadConversionAdjustmentsResponse) XXX_Size() int {
	return xxx_messageInfo_UploadConversionAdjustmentsResponse.Size(m)
}
func (m *UploadConversionAdjustmentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadConversionAdjustmentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadConversionAdjustmentsResponse proto.InternalMessageInfo

func (m *UploadConversionAdjustmentsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *UploadConversionAdjustmentsResponse) GetResults() []*ConversionAdjustmentResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// A conversion adjustment.
type ConversionAdjustment struct {
	// Resource name of the conversion action associated with this conversion
	// adjustment. Note: Although this resource name consists of a customer id and
	// a conversion action id, validation will ignore the customer id and use the
	// conversion action id as the sole identifier of the conversion action.
	ConversionAction *wrappers.StringValue `protobuf:"bytes,3,opt,name=conversion_action,json=conversionAction,proto3" json:"conversion_action,omitempty"`
	// The date time at which the adjustment occurred. Must be after the
	// conversion_date_time. The timezone must be specified. The format is
	// "yyyy-mm-dd hh:mm:ss+|-hh:mm", e.g. "2019-01-01 12:32:45-08:00".
	AdjustmentDateTime *wrappers.StringValue `protobuf:"bytes,4,opt,name=adjustment_date_time,json=adjustmentDateTime,proto3" json:"adjustment_date_time,omitempty"`
	// The adjustment type.
	AdjustmentType enums.ConversionAdjustmentTypeEnum_ConversionAdjustmentType `protobuf:"varint,5,opt,name=adjustment_type,json=adjustmentType,proto3,enum=google.ads.googleads.v1.enums.ConversionAdjustmentTypeEnum_ConversionAdjustmentType" json:"adjustment_type,omitempty"`
	// Information needed to restate the conversion's value.
	// Required for restatements. Should not be supplied for retractions. An error
	// will be returned if provided for a retraction.
	RestatementValue *RestatementValue `protobuf:"bytes,6,opt,name=restatement_value,json=restatementValue,proto3" json:"restatement_value,omitempty"`
	// Identifies the conversion to be adjusted.
	//
	// Types that are valid to be assigned to ConversionIdentifier:
	//	*ConversionAdjustment_GclidDateTimePair
	//	*ConversionAdjustment_OrderId
	ConversionIdentifier isConversionAdjustment_ConversionIdentifier `protobuf_oneof:"conversion_identifier"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *ConversionAdjustment) Reset()         { *m = ConversionAdjustment{} }
func (m *ConversionAdjustment) String() string { return proto.CompactTextString(m) }
func (*ConversionAdjustment) ProtoMessage()    {}
func (*ConversionAdjustment) Descriptor() ([]byte, []int) {
	return fileDescriptor_418d7f048cc75833, []int{2}
}

func (m *ConversionAdjustment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionAdjustment.Unmarshal(m, b)
}
func (m *ConversionAdjustment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionAdjustment.Marshal(b, m, deterministic)
}
func (m *ConversionAdjustment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionAdjustment.Merge(m, src)
}
func (m *ConversionAdjustment) XXX_Size() int {
	return xxx_messageInfo_ConversionAdjustment.Size(m)
}
func (m *ConversionAdjustment) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionAdjustment.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionAdjustment proto.InternalMessageInfo

func (m *ConversionAdjustment) GetConversionAction() *wrappers.StringValue {
	if m != nil {
		return m.ConversionAction
	}
	return nil
}

func (m *ConversionAdjustment) GetAdjustmentDateTime() *wrappers.StringValue {
	if m != nil {
		return m.AdjustmentDateTime
	}
	return nil
}

func (m *ConversionAdjustment) GetAdjustmentType() enums.ConversionAdjustmentTypeEnum_ConversionAdjustmentType {
	if m != nil {
		return m.AdjustmentType
	}
	return enums.ConversionAdjustmentTypeEnum_UNSPECIFIED
}

func (m *ConversionAdjustment) GetRestatementValue() *RestatementValue {
	if m != nil {
		return m.RestatementValue
	}
	return nil
}

type isConversionAdjustment_ConversionIdentifier interface {
	isConversionAdjustment_ConversionIdentifier()
}

type ConversionAdjustment_GclidDateTimePair struct {
	GclidDateTimePair *GclidDateTimePair `protobuf:"bytes,1,opt,name=gclid_date_time_pair,json=gclidDateTimePair,proto3,oneof"`
}

type ConversionAdjustment_OrderId struct {
	OrderId *wrappers.StringValue `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3,oneof"`
}

func (*ConversionAdjustment_GclidDateTimePair) isConversionAdjustment_ConversionIdentifier() {}

func (*ConversionAdjustment_OrderId) isConversionAdjustment_ConversionIdentifier() {}

func (m *ConversionAdjustment) GetConversionIdentifier() isConversionAdjustment_ConversionIdentifier {
	if m != nil {
		return m.ConversionIdentifier
	}
	return nil
}

func (m *ConversionAdjustment) GetGclidDateTimePair() *GclidDateTimePair {
	if x, ok := m.GetConversionIdentifier().(*ConversionAdjustment_GclidDateTimePair); ok {
		return x.GclidDateTimePair
	}
	return nil
}

func (m *ConversionAdjustment) GetOrderId() *wrappers.StringValue {
	if x, ok := m.GetConversionIdentifier().(*ConversionAdjustment_OrderId); ok {
		return x.OrderId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ConversionAdjustment) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ConversionAdjustment_GclidDateTimePair)(nil),
		(*ConversionAdjustment_OrderId)(nil),
	}
}

// Contains information needed to restate a conversion's value.
type RestatementValue struct {
	// The restated conversion value. This is the value of the conversion after
	// restatement. For example, to change the value of a conversion from 100 to
	// 70, an adjusted value of 70 should be reported.
	AdjustedValue *wrappers.DoubleValue `protobuf:"bytes,1,opt,name=adjusted_value,json=adjustedValue,proto3" json:"adjusted_value,omitempty"`
	// The currency of the restated value. If not provided, then the default
	// currency from the conversion action is used, and if that is not set then
	// the account currency is used. This is the ISO 4217 3-character currency
	// code e.g. USD or EUR.
	CurrencyCode         *wrappers.StringValue `protobuf:"bytes,2,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RestatementValue) Reset()         { *m = RestatementValue{} }
func (m *RestatementValue) String() string { return proto.CompactTextString(m) }
func (*RestatementValue) ProtoMessage()    {}
func (*RestatementValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_418d7f048cc75833, []int{3}
}

func (m *RestatementValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestatementValue.Unmarshal(m, b)
}
func (m *RestatementValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestatementValue.Marshal(b, m, deterministic)
}
func (m *RestatementValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestatementValue.Merge(m, src)
}
func (m *RestatementValue) XXX_Size() int {
	return xxx_messageInfo_RestatementValue.Size(m)
}
func (m *RestatementValue) XXX_DiscardUnknown() {
	xxx_messageInfo_RestatementValue.DiscardUnknown(m)
}

var xxx_messageInfo_RestatementValue proto.InternalMessageInfo

func (m *RestatementValue) GetAdjustedValue() *wrappers.DoubleValue {
	if m != nil {
		return m.AdjustedValue
	}
	return nil
}

func (m *RestatementValue) GetCurrencyCode() *wrappers.StringValue {
	if m != nil {
		return m.CurrencyCode
	}
	return nil
}

// Uniquely identifies a conversion that was reported without an order ID
// specified.
type GclidDateTimePair struct {
	// Google click ID (gclid) associated with the original conversion for this
	// adjustment.
	Gclid *wrappers.StringValue `protobuf:"bytes,1,opt,name=gclid,proto3" json:"gclid,omitempty"`
	// The date time at which the original conversion for this adjustment
	// occurred. The timezone must be specified. The format is "yyyy-mm-dd
	// hh:mm:ss+|-hh:mm", e.g. "2019-01-01 12:32:45-08:00".
	ConversionDateTime   *wrappers.StringValue `protobuf:"bytes,2,opt,name=conversion_date_time,json=conversionDateTime,proto3" json:"conversion_date_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GclidDateTimePair) Reset()         { *m = GclidDateTimePair{} }
func (m *GclidDateTimePair) String() string { return proto.CompactTextString(m) }
func (*GclidDateTimePair) ProtoMessage()    {}
func (*GclidDateTimePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_418d7f048cc75833, []int{4}
}

func (m *GclidDateTimePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GclidDateTimePair.Unmarshal(m, b)
}
func (m *GclidDateTimePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GclidDateTimePair.Marshal(b, m, deterministic)
}
func (m *GclidDateTimePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GclidDateTimePair.Merge(m, src)
}
func (m *GclidDateTimePair) XXX_Size() int {
	return xxx_messageInfo_GclidDateTimePair.Size(m)
}
func (m *GclidDateTimePair) XXX_DiscardUnknown() {
	xxx_messageInfo_GclidDateTimePair.DiscardUnknown(m)
}

var xxx_messageInfo_GclidDateTimePair proto.InternalMessageInfo

func (m *GclidDateTimePair) GetGclid() *wrappers.StringValue {
	if m != nil {
		return m.Gclid
	}
	return nil
}

func (m *GclidDateTimePair) GetConversionDateTime() *wrappers.StringValue {
	if m != nil {
		return m.ConversionDateTime
	}
	return nil
}

// Information identifying a successfully processed ConversionAdjustment.
type ConversionAdjustmentResult struct {
	// Resource name of the conversion action associated with this conversion
	// adjustment.
	ConversionAction *wrappers.StringValue `protobuf:"bytes,3,opt,name=conversion_action,json=conversionAction,proto3" json:"conversion_action,omitempty"`
	// The date time at which the adjustment occurred. The format is
	// "yyyy-mm-dd hh:mm:ss+|-hh:mm", e.g. "2019-01-01 12:32:45-08:00".
	AdjustmentDateTime *wrappers.StringValue `protobuf:"bytes,4,opt,name=adjustment_date_time,json=adjustmentDateTime,proto3" json:"adjustment_date_time,omitempty"`
	// The adjustment type.
	AdjustmentType enums.ConversionAdjustmentTypeEnum_ConversionAdjustmentType `protobuf:"varint,5,opt,name=adjustment_type,json=adjustmentType,proto3,enum=google.ads.googleads.v1.enums.ConversionAdjustmentTypeEnum_ConversionAdjustmentType" json:"adjustment_type,omitempty"`
	// Identifies the conversion that was adjusted.
	//
	// Types that are valid to be assigned to ConversionIdentifier:
	//	*ConversionAdjustmentResult_GclidDateTimePair
	//	*ConversionAdjustmentResult_OrderId
	ConversionIdentifier isConversionAdjustmentResult_ConversionIdentifier `protobuf_oneof:"conversion_identifier"`
	XXX_NoUnkeyedLiteral struct{}                                          `json:"-"`
	XXX_unrecognized     []byte                                            `json:"-"`
	XXX_sizecache        int32                                             `json:"-"`
}

func (m *ConversionAdjustmentResult) Reset()         { *m = ConversionAdjustmentResult{} }
func (m *ConversionAdjustmentResult) String() string { return proto.CompactTextString(m) }
func (*ConversionAdjustmentResult) ProtoMessage()    {}
func (*ConversionAdjustmentResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_418d7f048cc75833, []int{5}
}

func (m *ConversionAdjustmentResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionAdjustmentResult.Unmarshal(m, b)
}
func (m *ConversionAdjustmentResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionAdjustmentResult.Marshal(b, m, deterministic)
}
func (m *ConversionAdjustmentResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionAdjustmentResult.Merge(m, src)
}
func (m *ConversionAdjustmentResult) XXX_Size() int {
	return xxx_messageInfo_ConversionAdjustmentResult.Size(m)
}
func (m *ConversionAdjustmentResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionAdjustmentResult.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionAdjustmentResult proto.InternalMessageInfo

func (m *ConversionAdjustmentResult) GetConversionAction() *wrappers.StringValue {
	if m != nil {
		return m.ConversionAction
	}
	return nil
}

func (m *ConversionAdjustmentResult) GetAdjustmentDateTime() *wrappers.StringValue {
	if m != nil {
		return m.AdjustmentDateTime
	}
	return nil
}

func (m *ConversionAdjustmentResult) GetAdjustmentType() enums.ConversionAdjustmentTypeEnum_ConversionAdjustmentType {
	if m != nil {
		return m.AdjustmentType
	}
	return enums.ConversionAdjustmentTypeEnum_UNSPECIFIED
}

type isConversionAdjustmentResult_ConversionIdentifier interface {
	isConversionAdjustmentResult_ConversionIdentifier()
}

type ConversionAdjustmentResult_GclidDateTimePair struct {
	GclidDateTimePair *GclidDateTimePair `protobuf:"bytes,1,opt,name=gclid_date_time_pair,json=gclidDateTimePair,proto3,oneof"`
}

type ConversionAdjustmentResult_OrderId struct {
	OrderId *wrappers.StringValue `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3,oneof"`
}

func (*ConversionAdjustmentResult_GclidDateTimePair) isConversionAdjustmentResult_ConversionIdentifier() {
}

func (*ConversionAdjustmentResult_OrderId) isConversionAdjustmentResult_ConversionIdentifier() {}

func (m *ConversionAdjustmentResult) GetConversionIdentifier() isConversionAdjustmentResult_ConversionIdentifier {
	if m != nil {
		return m.ConversionIdentifier
	}
	return nil
}

func (m *ConversionAdjustmentResult) GetGclidDateTimePair() *GclidDateTimePair {
	if x, ok := m.GetConversionIdentifier().(*ConversionAdjustmentResult_GclidDateTimePair); ok {
		return x.GclidDateTimePair
	}
	return nil
}

func (m *ConversionAdjustmentResult) GetOrderId() *wrappers.StringValue {
	if x, ok := m.GetConversionIdentifier().(*ConversionAdjustmentResult_OrderId); ok {
		return x.OrderId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ConversionAdjustmentResult) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ConversionAdjustmentResult_GclidDateTimePair)(nil),
		(*ConversionAdjustmentResult_OrderId)(nil),
	}
}

func init() {
	proto.RegisterType((*UploadConversionAdjustmentsRequest)(nil), "google.ads.googleads.v1.services.UploadConversionAdjustmentsRequest")
	proto.RegisterType((*UploadConversionAdjustmentsResponse)(nil), "google.ads.googleads.v1.services.UploadConversionAdjustmentsResponse")
	proto.RegisterType((*ConversionAdjustment)(nil), "google.ads.googleads.v1.services.ConversionAdjustment")
	proto.RegisterType((*RestatementValue)(nil), "google.ads.googleads.v1.services.RestatementValue")
	proto.RegisterType((*GclidDateTimePair)(nil), "google.ads.googleads.v1.services.GclidDateTimePair")
	proto.RegisterType((*ConversionAdjustmentResult)(nil), "google.ads.googleads.v1.services.ConversionAdjustmentResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/conversion_adjustment_upload_service.proto", fileDescriptor_418d7f048cc75833)
}

var fileDescriptor_418d7f048cc75833 = []byte{
	// 845 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0x4f, 0x8f, 0xdb, 0x44,
	0x14, 0xaf, 0xbd, 0xfd, 0xc7, 0x6c, 0x77, 0xdb, 0x0c, 0x5b, 0x1a, 0x85, 0x0a, 0x82, 0x5b, 0x41,
	0xd4, 0x83, 0xad, 0xa4, 0x12, 0x52, 0x0d, 0x54, 0xca, 0x6e, 0xb6, 0xbb, 0x2b, 0x24, 0x58, 0x79,
	0x97, 0x1c, 0x50, 0x24, 0x6b, 0xd6, 0x33, 0xb1, 0x06, 0xd9, 0x33, 0x66, 0x66, 0x1c, 0x14, 0xa1,
	0x5e, 0x7a, 0xe7, 0x84, 0x04, 0x12, 0x47, 0x8e, 0x7c, 0x0f, 0x2e, 0x48, 0x9c, 0xf8, 0x0a, 0x9c,
	0x38, 0x21, 0x3e, 0x01, 0xf2, 0xd8, 0x93, 0xa4, 0x21, 0xae, 0x2b, 0x38, 0x21, 0xf5, 0x36, 0x7e,
	0x7f, 0x7e, 0xef, 0xf7, 0x7e, 0xf3, 0x9e, 0xc6, 0xe0, 0xe3, 0x98, 0xf3, 0x38, 0x21, 0x1e, 0xc2,
	0xd2, 0x2b, 0x8f, 0xc5, 0x69, 0xd6, 0xf7, 0x24, 0x11, 0x33, 0x1a, 0x11, 0xe9, 0x45, 0x9c, 0xcd,
	0x88, 0x90, 0x94, 0xb3, 0x10, 0xe1, 0x2f, 0x72, 0xa9, 0x52, 0xc2, 0x54, 0x98, 0x67, 0x09, 0x47,
	0x38, 0xac, 0xa2, 0xdc, 0x4c, 0x70, 0xc5, 0x61, 0xb7, 0x44, 0x70, 0x11, 0x96, 0xee, 0x02, 0xcc,
	0x9d, 0xf5, 0x5d, 0x03, 0xd6, 0x79, 0x5c, 0x57, 0x8e, 0xb0, 0x3c, 0xad, 0xab, 0xa5, 0xe6, 0x59,
	0x55, 0xa1, 0x73, 0xd7, 0xe4, 0x67, 0xd4, 0x43, 0x8c, 0x71, 0x85, 0x14, 0xe5, 0x4c, 0x56, 0xde,
	0xb7, 0x2a, 0xaf, 0xfe, 0xba, 0xc8, 0xa7, 0xde, 0x57, 0x02, 0x65, 0x19, 0x11, 0xc6, 0x7f, 0xa7,
	0xf2, 0x8b, 0x2c, 0xf2, 0xa4, 0x42, 0x2a, 0xaf, 0x1c, 0xce, 0x33, 0x1b, 0x38, 0x9f, 0xe9, 0x8e,
	0x0e, 0x16, 0x0c, 0x86, 0x0b, 0x02, 0x32, 0x20, 0x5f, 0xe6, 0x44, 0x2a, 0xf8, 0x36, 0xd8, 0x8e,
	0x72, 0xa9, 0x78, 0x4a, 0x44, 0x48, 0x71, 0xdb, 0xea, 0x5a, 0xbd, 0xd7, 0x02, 0x60, 0x4c, 0x27,
	0x18, 0xa6, 0xe0, 0x8d, 0x8d, 0x2d, 0xc8, 0xb6, 0xdd, 0xdd, 0xea, 0x6d, 0x0f, 0xde, 0x77, 0x9b,
	0x14, 0x72, 0x37, 0x11, 0x08, 0x6e, 0x47, 0x9b, 0x68, 0xc1, 0xf7, 0xc0, 0xcd, 0x0c, 0x09, 0x45,
	0x51, 0x12, 0x4e, 0x11, 0x4d, 0x72, 0x41, 0xda, 0x5b, 0x5d, 0xab, 0x77, 0x3d, 0xd8, 0xad, 0xcc,
	0x4f, 0x4a, 0x2b, 0xbc, 0x07, 0x76, 0x66, 0x28, 0xa1, 0x18, 0x29, 0x12, 0x72, 0x96, 0xcc, 0xdb,
	0x97, 0x75, 0xd8, 0x0d, 0x63, 0xfc, 0x94, 0x25, 0x73, 0xe7, 0x67, 0x0b, 0xdc, 0x7b, 0xa1, 0x08,
	0x32, 0xe3, 0x4c, 0x12, 0xf8, 0x04, 0xdc, 0x5e, 0xab, 0x1a, 0x12, 0x21, 0xb8, 0xd0, 0x7a, 0x6c,
	0x0f, 0xa0, 0xe9, 0x51, 0x64, 0x91, 0x7b, 0xa6, 0x55, 0x0e, 0x5e, 0x7f, 0x9e, 0xcf, 0x61, 0x11,
	0x0e, 0xc7, 0xe0, 0x9a, 0x20, 0x32, 0x4f, 0x16, 0xea, 0x7c, 0xf8, 0x2f, 0xd5, 0xd1, 0x20, 0x81,
	0x01, 0x73, 0x7e, 0xbd, 0x0c, 0xf6, 0x36, 0xc5, 0xc1, 0x13, 0xd0, 0x5a, 0xbd, 0x9d, 0xa8, 0x18,
	0x1d, 0x2d, 0xd8, 0xf6, 0xe0, 0xae, 0x29, 0x6d, 0x46, 0xc7, 0x3d, 0x53, 0x82, 0xb2, 0x78, 0x8c,
	0x92, 0x9c, 0x04, 0xb7, 0x56, 0xe4, 0xd7, 0x59, 0xf0, 0x13, 0xb0, 0xb7, 0x32, 0xa0, 0x5a, 0x57,
	0x45, 0x53, 0xa2, 0x75, 0x6d, 0x42, 0x83, 0xcb, 0xcc, 0x11, 0x52, 0xe4, 0x9c, 0xa6, 0x04, 0x3e,
	0x05, 0x37, 0xd7, 0x06, 0xbe, 0x7d, 0xa5, 0x6b, 0xf5, 0x76, 0x07, 0xe7, 0xb5, 0x9a, 0xe8, 0x8d,
	0xd9, 0x28, 0xc8, 0xf9, 0x3c, 0x23, 0x87, 0x2c, 0x4f, 0x6b, 0x9d, 0xc1, 0x2e, 0x7a, 0xee, 0x1b,
	0x86, 0xa0, 0x25, 0x48, 0xb1, 0x11, 0x44, 0xd7, 0x9f, 0x15, 0x3c, 0xdb, 0x57, 0x75, 0x2f, 0x83,
	0xe6, 0x4b, 0x09, 0x96, 0xa9, 0x95, 0x5e, 0x62, 0xcd, 0x02, 0xa7, 0x60, 0x2f, 0x8e, 0x12, 0x8a,
	0x97, 0x52, 0x85, 0x19, 0xa2, 0x66, 0x64, 0x1e, 0x36, 0xd7, 0x38, 0x2a, 0xb2, 0x8d, 0x5c, 0xa7,
	0x88, 0x8a, 0xe3, 0x4b, 0x41, 0x2b, 0x5e, 0x37, 0xc2, 0x47, 0xe0, 0x3a, 0x17, 0xb8, 0x5c, 0x4f,
	0xbb, 0xf9, 0x2e, 0x8e, 0x2f, 0x05, 0xd7, 0x74, 0xfc, 0x09, 0xde, 0xbf, 0x03, 0x56, 0xb6, 0x2c,
	0xa4, 0x98, 0x30, 0x45, 0xa7, 0x94, 0x08, 0xe7, 0x07, 0x0b, 0xdc, 0x5a, 0x6f, 0x11, 0x1e, 0x80,
	0x4a, 0x43, 0x82, 0x2b, 0xb9, 0xac, 0x9a, 0x72, 0x23, 0x9e, 0x5f, 0x24, 0xa4, 0x14, 0x66, 0xc7,
	0xe4, 0x94, 0x20, 0x43, 0xb0, 0x13, 0xe5, 0x42, 0x10, 0x16, 0xcd, 0xc3, 0x88, 0x63, 0xf2, 0x32,
	0x94, 0x83, 0x1b, 0x26, 0xe5, 0x80, 0x63, 0xe2, 0x7c, 0x6f, 0x81, 0xd6, 0x3f, 0xb4, 0x81, 0x03,
	0x70, 0x45, 0x6b, 0x53, 0x4b, 0x6a, 0x15, 0xb0, 0x0c, 0x2d, 0x46, 0x7a, 0xa5, 0xff, 0xe5, 0x48,
	0xbf, 0x0c, 0x27, 0xb8, 0xcc, 0x34, 0x3c, 0x9c, 0xbf, 0xb6, 0x40, 0xa7, 0x7e, 0x5d, 0x5f, 0x2d,
	0x63, 0xed, 0x32, 0xfe, 0x8f, 0x77, 0x65, 0xf0, 0x9d, 0x0d, 0xde, 0xd9, 0xd4, 0x68, 0xf9, 0xae,
	0x9c, 0x95, 0x0c, 0xe1, 0x9f, 0x16, 0x78, 0xf3, 0x05, 0x2f, 0x0d, 0x1c, 0x35, 0xf7, 0xd8, 0xfc,
	0x5a, 0x77, 0x0e, 0xff, 0x23, 0x4a, 0xf9, 0xdc, 0x39, 0xa3, 0x67, 0xbf, 0xfd, 0xfe, 0xad, 0xfd,
	0xd8, 0x79, 0x54, 0xfc, 0xa6, 0x98, 0xb7, 0x5e, 0x7a, 0x5f, 0xaf, 0xfc, 0x09, 0x7c, 0xf4, 0xe0,
	0xa9, 0x9f, 0xd7, 0x43, 0xf9, 0xd6, 0x83, 0xfd, 0x6f, 0x6c, 0x70, 0x3f, 0xe2, 0x69, 0x23, 0xa5,
	0xfd, 0x77, 0x1b, 0xe5, 0x3b, 0x2d, 0x2e, 0xe7, 0xd4, 0xfa, 0xfc, 0xb8, 0xc2, 0x8a, 0x79, 0x82,
	0x58, 0xec, 0x72, 0x11, 0x7b, 0x31, 0x61, 0xfa, 0xea, 0xcc, 0xbf, 0x55, 0x46, 0x65, 0xfd, 0x9f,
	0xdd, 0x07, 0xe6, 0xf0, 0xa3, 0xbd, 0x75, 0x34, 0x1c, 0xfe, 0x64, 0x77, 0x8f, 0x4a, 0xc0, 0x21,
	0x96, 0x6e, 0x79, 0x2c, 0x4e, 0xe3, 0xbe, 0x5b, 0x15, 0x96, 0xbf, 0x98, 0x90, 0xc9, 0x10, 0xcb,
	0xc9, 0x22, 0x64, 0x32, 0xee, 0x4f, 0x4c, 0xc8, 0x1f, 0xf6, 0xfd, 0xd2, 0xee, 0xfb, 0x43, 0x2c,
	0x7d, 0x7f, 0x11, 0xe4, 0xfb, 0xe3, 0xbe, 0xef, 0x9b, 0xb0, 0x8b, 0xab, 0x9a, 0xe7, 0xc3, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x24, 0xe2, 0x40, 0x64, 0x80, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConversionAdjustmentUploadServiceClient is the client API for ConversionAdjustmentUploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConversionAdjustmentUploadServiceClient interface {
	// Processes the given conversion adjustments.
	UploadConversionAdjustments(ctx context.Context, in *UploadConversionAdjustmentsRequest, opts ...grpc.CallOption) (*UploadConversionAdjustmentsResponse, error)
}

type conversionAdjustmentUploadServiceClient struct {
	cc *grpc.ClientConn
}

func NewConversionAdjustmentUploadServiceClient(cc *grpc.ClientConn) ConversionAdjustmentUploadServiceClient {
	return &conversionAdjustmentUploadServiceClient{cc}
}

func (c *conversionAdjustmentUploadServiceClient) UploadConversionAdjustments(ctx context.Context, in *UploadConversionAdjustmentsRequest, opts ...grpc.CallOption) (*UploadConversionAdjustmentsResponse, error) {
	out := new(UploadConversionAdjustmentsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ConversionAdjustmentUploadService/UploadConversionAdjustments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversionAdjustmentUploadServiceServer is the server API for ConversionAdjustmentUploadService service.
type ConversionAdjustmentUploadServiceServer interface {
	// Processes the given conversion adjustments.
	UploadConversionAdjustments(context.Context, *UploadConversionAdjustmentsRequest) (*UploadConversionAdjustmentsResponse, error)
}

// UnimplementedConversionAdjustmentUploadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConversionAdjustmentUploadServiceServer struct {
}

func (*UnimplementedConversionAdjustmentUploadServiceServer) UploadConversionAdjustments(ctx context.Context, req *UploadConversionAdjustmentsRequest) (*UploadConversionAdjustmentsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method UploadConversionAdjustments not implemented")
}

func RegisterConversionAdjustmentUploadServiceServer(s *grpc.Server, srv ConversionAdjustmentUploadServiceServer) {
	s.RegisterService(&_ConversionAdjustmentUploadService_serviceDesc, srv)
}

func _ConversionAdjustmentUploadService_UploadConversionAdjustments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadConversionAdjustmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionAdjustmentUploadServiceServer).UploadConversionAdjustments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ConversionAdjustmentUploadService/UploadConversionAdjustments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionAdjustmentUploadServiceServer).UploadConversionAdjustments(ctx, req.(*UploadConversionAdjustmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConversionAdjustmentUploadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.ConversionAdjustmentUploadService",
	HandlerType: (*ConversionAdjustmentUploadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadConversionAdjustments",
			Handler:    _ConversionAdjustmentUploadService_UploadConversionAdjustments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/conversion_adjustment_upload_service.proto",
}
