// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/billing_setup.proto

package resources

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

// A billing setup, which associates a payments account and an advertiser. A
// billing setup is specific to one advertiser.
type BillingSetup struct {
	// The resource name of the billing setup.
	// BillingSetup resource names have the form:
	//
	// `customers/{customer_id}/billingSetups/{billing_setup_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the billing setup.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The status of the billing setup.
	Status enums.BillingSetupStatusEnum_BillingSetupStatus `protobuf:"varint,3,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.BillingSetupStatusEnum_BillingSetupStatus" json:"status,omitempty"`
	// The resource name of the payments account associated with this billing
	// setup. Payments resource names have the form:
	//
	// `customers/{customer_id}/paymentsAccounts/{payments_account_id}`
	// When setting up billing, this is used to signup with an existing payments
	// account (and then payments_account_info should not be set).
	// When getting a billing setup, this and payments_account_info will be
	// populated.
	PaymentsAccount *wrappers.StringValue `protobuf:"bytes,11,opt,name=payments_account,json=paymentsAccount,proto3" json:"payments_account,omitempty"`
	// The payments account information associated with this billing setup.
	// When setting up billing, this is used to signup with a new payments account
	// (and then payments_account should not be set).
	// When getting a billing setup, this and payments_account will be
	// populated.
	PaymentsAccountInfo *BillingSetup_PaymentsAccountInfo `protobuf:"bytes,12,opt,name=payments_account_info,json=paymentsAccountInfo,proto3" json:"payments_account_info,omitempty"`
	// When creating a new billing setup, this is when the setup should take
	// effect. NOW is the only acceptable start time if the customer doesn't have
	// any approved setups.
	//
	// When fetching an existing billing setup, this is the requested start time.
	// However, if the setup was approved (see status) after the requested start
	// time, then this is the approval time.
	//
	// Types that are valid to be assigned to StartTime:
	//	*BillingSetup_StartDateTime
	//	*BillingSetup_StartTimeType
	StartTime isBillingSetup_StartTime `protobuf_oneof:"start_time"`
	// When the billing setup ends / ended. This is either FOREVER or the start
	// time of the next scheduled billing setup.
	//
	// Types that are valid to be assigned to EndTime:
	//	*BillingSetup_EndDateTime
	//	*BillingSetup_EndTimeType
	EndTime              isBillingSetup_EndTime `protobuf_oneof:"end_time"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *BillingSetup) Reset()         { *m = BillingSetup{} }
func (m *BillingSetup) String() string { return proto.CompactTextString(m) }
func (*BillingSetup) ProtoMessage()    {}
func (*BillingSetup) Descriptor() ([]byte, []int) {
	return fileDescriptor_e355550cdf5be9da, []int{0}
}

func (m *BillingSetup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillingSetup.Unmarshal(m, b)
}
func (m *BillingSetup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillingSetup.Marshal(b, m, deterministic)
}
func (m *BillingSetup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillingSetup.Merge(m, src)
}
func (m *BillingSetup) XXX_Size() int {
	return xxx_messageInfo_BillingSetup.Size(m)
}
func (m *BillingSetup) XXX_DiscardUnknown() {
	xxx_messageInfo_BillingSetup.DiscardUnknown(m)
}

var xxx_messageInfo_BillingSetup proto.InternalMessageInfo

func (m *BillingSetup) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *BillingSetup) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *BillingSetup) GetStatus() enums.BillingSetupStatusEnum_BillingSetupStatus {
	if m != nil {
		return m.Status
	}
	return enums.BillingSetupStatusEnum_UNSPECIFIED
}

func (m *BillingSetup) GetPaymentsAccount() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsAccount
	}
	return nil
}

func (m *BillingSetup) GetPaymentsAccountInfo() *BillingSetup_PaymentsAccountInfo {
	if m != nil {
		return m.PaymentsAccountInfo
	}
	return nil
}

type isBillingSetup_StartTime interface {
	isBillingSetup_StartTime()
}

type BillingSetup_StartDateTime struct {
	StartDateTime *wrappers.StringValue `protobuf:"bytes,9,opt,name=start_date_time,json=startDateTime,proto3,oneof"`
}

type BillingSetup_StartTimeType struct {
	StartTimeType enums.TimeTypeEnum_TimeType `protobuf:"varint,10,opt,name=start_time_type,json=startTimeType,proto3,enum=google.ads.googleads.v2.enums.TimeTypeEnum_TimeType,oneof"`
}

func (*BillingSetup_StartDateTime) isBillingSetup_StartTime() {}

func (*BillingSetup_StartTimeType) isBillingSetup_StartTime() {}

func (m *BillingSetup) GetStartTime() isBillingSetup_StartTime {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *BillingSetup) GetStartDateTime() *wrappers.StringValue {
	if x, ok := m.GetStartTime().(*BillingSetup_StartDateTime); ok {
		return x.StartDateTime
	}
	return nil
}

func (m *BillingSetup) GetStartTimeType() enums.TimeTypeEnum_TimeType {
	if x, ok := m.GetStartTime().(*BillingSetup_StartTimeType); ok {
		return x.StartTimeType
	}
	return enums.TimeTypeEnum_UNSPECIFIED
}

type isBillingSetup_EndTime interface {
	isBillingSetup_EndTime()
}

type BillingSetup_EndDateTime struct {
	EndDateTime *wrappers.StringValue `protobuf:"bytes,13,opt,name=end_date_time,json=endDateTime,proto3,oneof"`
}

type BillingSetup_EndTimeType struct {
	EndTimeType enums.TimeTypeEnum_TimeType `protobuf:"varint,14,opt,name=end_time_type,json=endTimeType,proto3,enum=google.ads.googleads.v2.enums.TimeTypeEnum_TimeType,oneof"`
}

func (*BillingSetup_EndDateTime) isBillingSetup_EndTime() {}

func (*BillingSetup_EndTimeType) isBillingSetup_EndTime() {}

func (m *BillingSetup) GetEndTime() isBillingSetup_EndTime {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *BillingSetup) GetEndDateTime() *wrappers.StringValue {
	if x, ok := m.GetEndTime().(*BillingSetup_EndDateTime); ok {
		return x.EndDateTime
	}
	return nil
}

func (m *BillingSetup) GetEndTimeType() enums.TimeTypeEnum_TimeType {
	if x, ok := m.GetEndTime().(*BillingSetup_EndTimeType); ok {
		return x.EndTimeType
	}
	return enums.TimeTypeEnum_UNSPECIFIED
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BillingSetup) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BillingSetup_StartDateTime)(nil),
		(*BillingSetup_StartTimeType)(nil),
		(*BillingSetup_EndDateTime)(nil),
		(*BillingSetup_EndTimeType)(nil),
	}
}

// Container of payments account information for this billing.
type BillingSetup_PaymentsAccountInfo struct {
	// A 16 digit id used to identify the payments account associated with the
	// billing setup.
	//
	// This must be passed as a string with dashes, e.g. "1234-5678-9012-3456".
	PaymentsAccountId *wrappers.StringValue `protobuf:"bytes,1,opt,name=payments_account_id,json=paymentsAccountId,proto3" json:"payments_account_id,omitempty"`
	// The name of the payments account associated with the billing setup.
	//
	// This enables the user to specify a meaningful name for a payments account
	// to aid in reconciling monthly invoices.
	//
	// This name will be printed in the monthly invoices.
	PaymentsAccountName *wrappers.StringValue `protobuf:"bytes,2,opt,name=payments_account_name,json=paymentsAccountName,proto3" json:"payments_account_name,omitempty"`
	// A 12 digit id used to identify the payments profile associated with the
	// billing setup.
	//
	// This must be passed in as a string with dashes, e.g. "1234-5678-9012".
	PaymentsProfileId *wrappers.StringValue `protobuf:"bytes,3,opt,name=payments_profile_id,json=paymentsProfileId,proto3" json:"payments_profile_id,omitempty"`
	// The name of the payments profile associated with the billing setup.
	PaymentsProfileName *wrappers.StringValue `protobuf:"bytes,4,opt,name=payments_profile_name,json=paymentsProfileName,proto3" json:"payments_profile_name,omitempty"`
	// A secondary payments profile id present in uncommon situations, e.g.
	// when a sequential liability agreement has been arranged.
	SecondaryPaymentsProfileId *wrappers.StringValue `protobuf:"bytes,5,opt,name=secondary_payments_profile_id,json=secondaryPaymentsProfileId,proto3" json:"secondary_payments_profile_id,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}              `json:"-"`
	XXX_unrecognized           []byte                `json:"-"`
	XXX_sizecache              int32                 `json:"-"`
}

func (m *BillingSetup_PaymentsAccountInfo) Reset()         { *m = BillingSetup_PaymentsAccountInfo{} }
func (m *BillingSetup_PaymentsAccountInfo) String() string { return proto.CompactTextString(m) }
func (*BillingSetup_PaymentsAccountInfo) ProtoMessage()    {}
func (*BillingSetup_PaymentsAccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e355550cdf5be9da, []int{0, 0}
}

func (m *BillingSetup_PaymentsAccountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillingSetup_PaymentsAccountInfo.Unmarshal(m, b)
}
func (m *BillingSetup_PaymentsAccountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillingSetup_PaymentsAccountInfo.Marshal(b, m, deterministic)
}
func (m *BillingSetup_PaymentsAccountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillingSetup_PaymentsAccountInfo.Merge(m, src)
}
func (m *BillingSetup_PaymentsAccountInfo) XXX_Size() int {
	return xxx_messageInfo_BillingSetup_PaymentsAccountInfo.Size(m)
}
func (m *BillingSetup_PaymentsAccountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BillingSetup_PaymentsAccountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BillingSetup_PaymentsAccountInfo proto.InternalMessageInfo

func (m *BillingSetup_PaymentsAccountInfo) GetPaymentsAccountId() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsAccountId
	}
	return nil
}

func (m *BillingSetup_PaymentsAccountInfo) GetPaymentsAccountName() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsAccountName
	}
	return nil
}

func (m *BillingSetup_PaymentsAccountInfo) GetPaymentsProfileId() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsProfileId
	}
	return nil
}

func (m *BillingSetup_PaymentsAccountInfo) GetPaymentsProfileName() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsProfileName
	}
	return nil
}

func (m *BillingSetup_PaymentsAccountInfo) GetSecondaryPaymentsProfileId() *wrappers.StringValue {
	if m != nil {
		return m.SecondaryPaymentsProfileId
	}
	return nil
}

func init() {
	proto.RegisterType((*BillingSetup)(nil), "google.ads.googleads.v2.resources.BillingSetup")
	proto.RegisterType((*BillingSetup_PaymentsAccountInfo)(nil), "google.ads.googleads.v2.resources.BillingSetup.PaymentsAccountInfo")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/billing_setup.proto", fileDescriptor_e355550cdf5be9da)
}

var fileDescriptor_e355550cdf5be9da = []byte{
	// 630 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0x80, 0x67, 0x67, 0x2b, 0xab, 0x9a, 0xb4, 0xab, 0xcb, 0xc0, 0x64, 0xdd, 0x68, 0x37, 0x0a,
	0x85, 0x31, 0x19, 0xbc, 0x6e, 0x0c, 0xef, 0x2a, 0xd9, 0x4f, 0xdb, 0x31, 0x46, 0x70, 0x4b, 0x2e,
	0x4a, 0x98, 0xa7, 0x46, 0x8a, 0x11, 0xc4, 0x92, 0xb1, 0xe4, 0x96, 0x3c, 0xcd, 0x60, 0x97, 0x7b,
	0x80, 0x3d, 0xc4, 0x1e, 0x65, 0x0f, 0x31, 0x86, 0x25, 0xcb, 0x4d, 0xdd, 0xb4, 0x49, 0xd9, 0x9d,
	0x24, 0x9f, 0xf3, 0xf1, 0xe9, 0x9c, 0xa3, 0x04, 0xbc, 0x8a, 0x39, 0x8f, 0xc7, 0xc4, 0x43, 0x58,
	0x78, 0x7a, 0x59, 0xac, 0xce, 0x7c, 0x2f, 0x23, 0x82, 0xe7, 0xd9, 0x90, 0x08, 0xef, 0x94, 0x8e,
	0xc7, 0x94, 0xc5, 0x91, 0x20, 0x32, 0x4f, 0x61, 0x9a, 0x71, 0xc9, 0x9d, 0x6d, 0x1d, 0x0b, 0x11,
	0x16, 0xb0, 0x4a, 0x83, 0x67, 0x3e, 0xac, 0xd2, 0xda, 0x6f, 0xae, 0x23, 0x13, 0x96, 0x27, 0x35,
	0x6a, 0x24, 0x24, 0x92, 0xb9, 0xd0, 0xf0, 0xf6, 0x8b, 0x9b, 0x33, 0x25, 0x4d, 0x48, 0x24, 0x27,
	0x29, 0x29, 0xc3, 0x9f, 0x94, 0xe1, 0x6a, 0x77, 0x9a, 0x8f, 0xbc, 0xf3, 0x0c, 0xa5, 0x29, 0xc9,
	0x0c, 0x6e, 0xd3, 0xe0, 0x52, 0xea, 0x21, 0xc6, 0xb8, 0x44, 0x92, 0x72, 0x56, 0x7e, 0x7d, 0xfa,
	0x7d, 0x19, 0x34, 0xbb, 0xda, 0xe5, 0xa8, 0x50, 0x71, 0x9e, 0x81, 0x96, 0xb9, 0x44, 0xc4, 0x50,
	0x42, 0x5c, 0x6b, 0xcb, 0xda, 0x5d, 0x0e, 0x9b, 0xe6, 0xf0, 0x0b, 0x4a, 0x88, 0xf3, 0x1c, 0xd8,
	0x14, 0xbb, 0xf6, 0x96, 0xb5, 0xbb, 0xe2, 0x3f, 0x2a, 0x2b, 0x00, 0x8d, 0x00, 0x3c, 0x64, 0xf2,
	0xf5, 0x5e, 0x1f, 0x8d, 0x73, 0x12, 0xda, 0x14, 0x3b, 0xdf, 0xc0, 0x92, 0xbe, 0x9f, 0xdb, 0xd8,
	0xb2, 0x76, 0x57, 0xfd, 0x03, 0x78, 0x5d, 0xf5, 0xd4, 0x05, 0xe1, 0xb4, 0xce, 0x91, 0x4a, 0xfc,
	0xc0, 0xf2, 0x64, 0xc6, 0x71, 0x58, 0x72, 0x9d, 0x7d, 0xf0, 0x20, 0x45, 0x93, 0x84, 0x30, 0x29,
	0x22, 0x34, 0x1c, 0xf2, 0x9c, 0x49, 0x77, 0x45, 0xc9, 0x6d, 0x5e, 0x91, 0x3b, 0x92, 0x19, 0x65,
	0xb1, 0xb6, 0x5b, 0x33, 0x59, 0x1d, 0x9d, 0xe4, 0x9c, 0x83, 0x87, 0x75, 0x50, 0x44, 0xd9, 0x88,
	0xbb, 0x4d, 0x45, 0x7b, 0x07, 0xe7, 0xf6, 0xfd, 0x92, 0x26, 0xec, 0x5d, 0xe6, 0x1f, 0xb2, 0x11,
	0x0f, 0x37, 0xd2, 0xab, 0x87, 0xce, 0x47, 0xb0, 0x26, 0x24, 0xca, 0x64, 0x84, 0x91, 0x24, 0x51,
	0xd1, 0x62, 0x77, 0x79, 0xfe, 0x05, 0x0e, 0xee, 0x84, 0x2d, 0x95, 0xf6, 0x1e, 0x49, 0x72, 0x4c,
	0x13, 0xe2, 0x7c, 0x35, 0x9c, 0x6a, 0x4a, 0x5c, 0xa0, 0x8a, 0xbe, 0x37, 0xa7, 0xe8, 0x45, 0xf6,
	0xf1, 0x24, 0x25, 0xaa, 0xd4, 0x66, 0x53, 0xf1, 0xcd, 0x81, 0xd3, 0x05, 0x2d, 0xc2, 0xf0, 0x94,
	0x65, 0x6b, 0x01, 0x4b, 0x2b, 0x5c, 0x21, 0x0c, 0x57, 0x8e, 0x27, 0x9a, 0x71, 0x61, 0xb8, 0xfa,
	0x1f, 0x86, 0x9a, 0x6d, 0xb6, 0xed, 0x5f, 0x0d, 0xb0, 0x31, 0xa3, 0xe8, 0xce, 0x67, 0xb0, 0x71,
	0xb5, 0xb1, 0x58, 0xcd, 0xf6, 0xbc, 0x21, 0x59, 0xaf, 0xf7, 0x0b, 0x3b, 0xbd, 0x19, 0x63, 0xa2,
	0xde, 0x8a, 0xbd, 0x00, 0xaf, 0xde, 0x7f, 0xf5, 0xa0, 0xa6, 0xfd, 0xd2, 0x8c, 0x8f, 0xe8, 0x98,
	0x14, 0x7e, 0x8d, 0xdb, 0xf8, 0xf5, 0x74, 0x5e, 0xcd, 0xcf, 0xd0, 0x94, 0xdf, 0xdd, 0xdb, 0xf8,
	0x95, 0x3c, 0xe5, 0x17, 0x81, 0xc7, 0x82, 0x0c, 0x39, 0xc3, 0x28, 0x9b, 0x44, 0xb3, 0x4c, 0xef,
	0x2d, 0x40, 0x6e, 0x57, 0x88, 0x5e, 0x5d, 0xb9, 0xdb, 0x04, 0xe0, 0x62, 0x70, 0xbb, 0x00, 0xdc,
	0x37, 0x23, 0xd2, 0xfd, 0x6b, 0x81, 0x9d, 0x21, 0x4f, 0xe6, 0x3f, 0xbd, 0xee, 0xfa, 0xf4, 0xdb,
	0xeb, 0x15, 0x06, 0x3d, 0xeb, 0xe4, 0x53, 0x99, 0x17, 0xf3, 0x31, 0x62, 0x31, 0xe4, 0x59, 0xec,
	0xc5, 0x84, 0x29, 0x3f, 0xf3, 0xeb, 0x9a, 0x52, 0x71, 0xc3, 0x1f, 0xc0, 0xdb, 0x6a, 0xf5, 0xc3,
	0x6e, 0xec, 0x77, 0x3a, 0x3f, 0xed, 0xed, 0x7d, 0x8d, 0xec, 0x60, 0x01, 0xf5, 0xb2, 0x58, 0xf5,
	0x7d, 0x18, 0x9a, 0xc8, 0xdf, 0x26, 0x66, 0xd0, 0xc1, 0x62, 0x50, 0xc5, 0x0c, 0xfa, 0xfe, 0xa0,
	0x8a, 0xf9, 0x63, 0xef, 0xe8, 0x0f, 0x41, 0xd0, 0xc1, 0x22, 0x08, 0xaa, 0xa8, 0x20, 0xe8, 0xfb,
	0x41, 0x50, 0xc5, 0x9d, 0x2e, 0x29, 0xd9, 0x97, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x94,
	0xca, 0xc7, 0xac, 0x06, 0x00, 0x00,
}
