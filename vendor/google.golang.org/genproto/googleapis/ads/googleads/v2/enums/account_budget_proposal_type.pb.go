// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/account_budget_proposal_type.proto

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

// The possible types of an AccountBudgetProposal.
type AccountBudgetProposalTypeEnum_AccountBudgetProposalType int32

const (
	// Not specified.
	AccountBudgetProposalTypeEnum_UNSPECIFIED AccountBudgetProposalTypeEnum_AccountBudgetProposalType = 0
	// Used for return value only. Represents value unknown in this version.
	AccountBudgetProposalTypeEnum_UNKNOWN AccountBudgetProposalTypeEnum_AccountBudgetProposalType = 1
	// Identifies a request to create a new budget.
	AccountBudgetProposalTypeEnum_CREATE AccountBudgetProposalTypeEnum_AccountBudgetProposalType = 2
	// Identifies a request to edit an existing budget.
	AccountBudgetProposalTypeEnum_UPDATE AccountBudgetProposalTypeEnum_AccountBudgetProposalType = 3
	// Identifies a request to end a budget that has already started.
	AccountBudgetProposalTypeEnum_END AccountBudgetProposalTypeEnum_AccountBudgetProposalType = 4
	// Identifies a request to remove a budget that hasn't started yet.
	AccountBudgetProposalTypeEnum_REMOVE AccountBudgetProposalTypeEnum_AccountBudgetProposalType = 5
)

var AccountBudgetProposalTypeEnum_AccountBudgetProposalType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CREATE",
	3: "UPDATE",
	4: "END",
	5: "REMOVE",
}

var AccountBudgetProposalTypeEnum_AccountBudgetProposalType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"CREATE":      2,
	"UPDATE":      3,
	"END":         4,
	"REMOVE":      5,
}

func (x AccountBudgetProposalTypeEnum_AccountBudgetProposalType) String() string {
	return proto.EnumName(AccountBudgetProposalTypeEnum_AccountBudgetProposalType_name, int32(x))
}

func (AccountBudgetProposalTypeEnum_AccountBudgetProposalType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7c7f0c81e59de710, []int{0, 0}
}

// Message describing AccountBudgetProposal types.
type AccountBudgetProposalTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountBudgetProposalTypeEnum) Reset()         { *m = AccountBudgetProposalTypeEnum{} }
func (m *AccountBudgetProposalTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AccountBudgetProposalTypeEnum) ProtoMessage()    {}
func (*AccountBudgetProposalTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c7f0c81e59de710, []int{0}
}

func (m *AccountBudgetProposalTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBudgetProposalTypeEnum.Unmarshal(m, b)
}
func (m *AccountBudgetProposalTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBudgetProposalTypeEnum.Marshal(b, m, deterministic)
}
func (m *AccountBudgetProposalTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBudgetProposalTypeEnum.Merge(m, src)
}
func (m *AccountBudgetProposalTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AccountBudgetProposalTypeEnum.Size(m)
}
func (m *AccountBudgetProposalTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBudgetProposalTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBudgetProposalTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.AccountBudgetProposalTypeEnum_AccountBudgetProposalType", AccountBudgetProposalTypeEnum_AccountBudgetProposalType_name, AccountBudgetProposalTypeEnum_AccountBudgetProposalType_value)
	proto.RegisterType((*AccountBudgetProposalTypeEnum)(nil), "google.ads.googleads.v2.enums.AccountBudgetProposalTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/account_budget_proposal_type.proto", fileDescriptor_7c7f0c81e59de710)
}

var fileDescriptor_7c7f0c81e59de710 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x6a, 0xc2, 0x30,
	0x14, 0x86, 0xd7, 0xba, 0x29, 0xc4, 0x8b, 0x95, 0xde, 0x6d, 0xcc, 0x81, 0x3e, 0x40, 0x0a, 0xdd,
	0x5d, 0x76, 0xb3, 0x54, 0x33, 0x91, 0xb1, 0x5a, 0x9c, 0x76, 0x30, 0x0a, 0x12, 0x6d, 0x17, 0x04,
	0x4d, 0x82, 0x69, 0x05, 0x9f, 0x60, 0xef, 0xb1, 0xcb, 0x3d, 0xca, 0x1e, 0x65, 0xb7, 0x7b, 0x81,
	0x91, 0x44, 0xbd, 0xeb, 0x6e, 0xc2, 0x47, 0xce, 0x39, 0xff, 0x7f, 0xce, 0x0f, 0x1e, 0x98, 0x10,
	0x6c, 0x5d, 0x04, 0x34, 0x57, 0x81, 0x45, 0x4d, 0xbb, 0x30, 0x28, 0x78, 0xb5, 0x51, 0x01, 0x5d,
	0x2e, 0x45, 0xc5, 0xcb, 0xf9, 0xa2, 0xca, 0x59, 0x51, 0xce, 0xe5, 0x56, 0x48, 0xa1, 0xe8, 0x7a,
	0x5e, 0xee, 0x65, 0x01, 0xe5, 0x56, 0x94, 0xc2, 0xef, 0xd8, 0x31, 0x48, 0x73, 0x05, 0x4f, 0x0a,
	0x70, 0x17, 0x42, 0xa3, 0x70, 0x7d, 0x73, 0x34, 0x90, 0xab, 0x80, 0x72, 0x2e, 0x4a, 0x5a, 0xae,
	0x04, 0x57, 0x76, 0xb8, 0xf7, 0xe1, 0x80, 0x0e, 0xb6, 0x1e, 0x91, 0xb1, 0x48, 0x0e, 0x0e, 0xd3,
	0xbd, 0x2c, 0x08, 0xaf, 0x36, 0xbd, 0x77, 0x70, 0x55, 0xdb, 0xe0, 0x5f, 0x82, 0xf6, 0x2c, 0x7e,
	0x49, 0x48, 0x7f, 0xf4, 0x38, 0x22, 0x03, 0xef, 0xcc, 0x6f, 0x83, 0xd6, 0x2c, 0x7e, 0x8a, 0xc7,
	0xaf, 0xb1, 0xe7, 0xf8, 0x00, 0x34, 0xfb, 0x13, 0x82, 0xa7, 0xc4, 0x73, 0x35, 0xcf, 0x92, 0x81,
	0xe6, 0x86, 0xdf, 0x02, 0x0d, 0x12, 0x0f, 0xbc, 0x73, 0xfd, 0x39, 0x21, 0xcf, 0xe3, 0x94, 0x78,
	0x17, 0xd1, 0xaf, 0x03, 0xba, 0x4b, 0xb1, 0x81, 0xff, 0x5e, 0x13, 0xdd, 0xd6, 0xee, 0x92, 0xe8,
	0x7b, 0x12, 0xe7, 0x2d, 0x3a, 0x08, 0x30, 0xb1, 0xa6, 0x9c, 0x41, 0xb1, 0x65, 0x01, 0x2b, 0xb8,
	0xb9, 0xf6, 0x18, 0xb0, 0x5c, 0xa9, 0x9a, 0xbc, 0xef, 0xcd, 0xfb, 0xe9, 0x36, 0x86, 0x18, 0x7f,
	0xb9, 0x9d, 0xa1, 0x95, 0xc2, 0xb9, 0x82, 0x16, 0x35, 0xa5, 0x21, 0xd4, 0xc1, 0xa8, 0xef, 0x63,
	0x3d, 0xc3, 0xb9, 0xca, 0x4e, 0xf5, 0x2c, 0x0d, 0x33, 0x53, 0xff, 0x71, 0xbb, 0xf6, 0x13, 0x21,
	0x9c, 0x2b, 0x84, 0x4e, 0x1d, 0x08, 0xa5, 0x21, 0x42, 0xa6, 0x67, 0xd1, 0x34, 0x8b, 0xdd, 0xfd,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x46, 0x87, 0x1a, 0x07, 0x02, 0x00, 0x00,
}
