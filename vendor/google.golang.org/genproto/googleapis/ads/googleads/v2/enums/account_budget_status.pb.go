// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/account_budget_status.proto

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

// The possible statuses of an AccountBudget.
type AccountBudgetStatusEnum_AccountBudgetStatus int32

const (
	// Not specified.
	AccountBudgetStatusEnum_UNSPECIFIED AccountBudgetStatusEnum_AccountBudgetStatus = 0
	// Used for return value only. Represents value unknown in this version.
	AccountBudgetStatusEnum_UNKNOWN AccountBudgetStatusEnum_AccountBudgetStatus = 1
	// The account budget is pending approval.
	AccountBudgetStatusEnum_PENDING AccountBudgetStatusEnum_AccountBudgetStatus = 2
	// The account budget has been approved.
	AccountBudgetStatusEnum_APPROVED AccountBudgetStatusEnum_AccountBudgetStatus = 3
	// The account budget has been cancelled by the user.
	AccountBudgetStatusEnum_CANCELLED AccountBudgetStatusEnum_AccountBudgetStatus = 4
)

var AccountBudgetStatusEnum_AccountBudgetStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PENDING",
	3: "APPROVED",
	4: "CANCELLED",
}

var AccountBudgetStatusEnum_AccountBudgetStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"PENDING":     2,
	"APPROVED":    3,
	"CANCELLED":   4,
}

func (x AccountBudgetStatusEnum_AccountBudgetStatus) String() string {
	return proto.EnumName(AccountBudgetStatusEnum_AccountBudgetStatus_name, int32(x))
}

func (AccountBudgetStatusEnum_AccountBudgetStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6121ddbe27d5d48b, []int{0, 0}
}

// Message describing AccountBudget statuses.
type AccountBudgetStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountBudgetStatusEnum) Reset()         { *m = AccountBudgetStatusEnum{} }
func (m *AccountBudgetStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AccountBudgetStatusEnum) ProtoMessage()    {}
func (*AccountBudgetStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_6121ddbe27d5d48b, []int{0}
}

func (m *AccountBudgetStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBudgetStatusEnum.Unmarshal(m, b)
}
func (m *AccountBudgetStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBudgetStatusEnum.Marshal(b, m, deterministic)
}
func (m *AccountBudgetStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBudgetStatusEnum.Merge(m, src)
}
func (m *AccountBudgetStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AccountBudgetStatusEnum.Size(m)
}
func (m *AccountBudgetStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBudgetStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBudgetStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.AccountBudgetStatusEnum_AccountBudgetStatus", AccountBudgetStatusEnum_AccountBudgetStatus_name, AccountBudgetStatusEnum_AccountBudgetStatus_value)
	proto.RegisterType((*AccountBudgetStatusEnum)(nil), "google.ads.googleads.v2.enums.AccountBudgetStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/account_budget_status.proto", fileDescriptor_6121ddbe27d5d48b)
}

var fileDescriptor_6121ddbe27d5d48b = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcd, 0x4a, 0xc3, 0x30,
	0x00, 0x76, 0x9d, 0xf8, 0x93, 0x29, 0x96, 0x7a, 0x50, 0xc4, 0x1d, 0xb6, 0x07, 0x48, 0xa1, 0x9e,
	0x8c, 0xa7, 0x74, 0x8d, 0x63, 0x38, 0xb2, 0xe2, 0x58, 0x05, 0xa9, 0x8c, 0x6c, 0x29, 0x61, 0xb0,
	0x25, 0x63, 0x49, 0x87, 0xcf, 0xe3, 0xd1, 0x47, 0xf1, 0x51, 0x04, 0xdf, 0x41, 0x9a, 0x6c, 0x3b,
	0x4d, 0x2f, 0xe1, 0x4b, 0xbe, 0x1f, 0xbe, 0x7c, 0xe0, 0x5e, 0x28, 0x25, 0xe6, 0x45, 0xc8, 0xb8,
	0x0e, 0x1d, 0xac, 0xd0, 0x3a, 0x0a, 0x0b, 0x59, 0x2e, 0x74, 0xc8, 0xa6, 0x53, 0x55, 0x4a, 0x33,
	0x9e, 0x94, 0x5c, 0x14, 0x66, 0xac, 0x0d, 0x33, 0xa5, 0x86, 0xcb, 0x95, 0x32, 0x2a, 0x68, 0x3a,
	0x3d, 0x64, 0x5c, 0xc3, 0x9d, 0x15, 0xae, 0x23, 0x68, 0xad, 0x37, 0xb7, 0xdb, 0xe4, 0xe5, 0x2c,
	0x64, 0x52, 0x2a, 0xc3, 0xcc, 0x4c, 0xc9, 0x8d, 0xb9, 0xfd, 0x0e, 0xae, 0xb0, 0xcb, 0x8e, 0x6d,
	0xf4, 0xd0, 0x26, 0x13, 0x59, 0x2e, 0xda, 0x6f, 0xe0, 0x72, 0x0f, 0x15, 0x5c, 0x80, 0xc6, 0x88,
	0x0e, 0x53, 0xd2, 0xe9, 0x3d, 0xf6, 0x48, 0xe2, 0x1f, 0x04, 0x0d, 0x70, 0x3c, 0xa2, 0x4f, 0x74,
	0xf0, 0x42, 0xfd, 0x5a, 0x75, 0x49, 0x09, 0x4d, 0x7a, 0xb4, 0xeb, 0x7b, 0xc1, 0x19, 0x38, 0xc1,
	0x69, 0xfa, 0x3c, 0xc8, 0x48, 0xe2, 0xd7, 0x83, 0x73, 0x70, 0xda, 0xc1, 0xb4, 0x43, 0xfa, 0x7d,
	0x92, 0xf8, 0x87, 0xf1, 0x4f, 0x0d, 0xb4, 0xa6, 0x6a, 0x01, 0xff, 0x6d, 0x1f, 0x5f, 0xef, 0xa9,
	0x90, 0x56, 0xcd, 0xd3, 0xda, 0x6b, 0xbc, 0xb1, 0x0a, 0x35, 0x67, 0x52, 0x40, 0xb5, 0x12, 0xa1,
	0x28, 0xa4, 0xfd, 0xd7, 0x76, 0xc3, 0xe5, 0x4c, 0xff, 0x31, 0xe9, 0x83, 0x3d, 0x3f, 0xbc, 0x7a,
	0x17, 0xe3, 0x4f, 0xaf, 0xd9, 0x75, 0x51, 0x98, 0x6b, 0xe8, 0x60, 0x85, 0xb2, 0x08, 0x56, 0x4b,
	0xe8, 0xaf, 0x2d, 0x9f, 0x63, 0xae, 0xf3, 0x1d, 0x9f, 0x67, 0x51, 0x6e, 0xf9, 0x6f, 0xaf, 0xe5,
	0x1e, 0x11, 0xc2, 0x5c, 0x23, 0xb4, 0x53, 0x20, 0x94, 0x45, 0x08, 0x59, 0xcd, 0xe4, 0xc8, 0x16,
	0xbb, 0xfb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x74, 0x52, 0xd4, 0x33, 0xea, 0x01, 0x00, 0x00,
}
