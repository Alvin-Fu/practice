// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/user_list_combined_rule_operator.proto

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

// Enum describing possible user list combined rule operators.
type UserListCombinedRuleOperatorEnum_UserListCombinedRuleOperator int32

const (
	// Not specified.
	UserListCombinedRuleOperatorEnum_UNSPECIFIED UserListCombinedRuleOperatorEnum_UserListCombinedRuleOperator = 0
	// Used for return value only. Represents value unknown in this version.
	UserListCombinedRuleOperatorEnum_UNKNOWN UserListCombinedRuleOperatorEnum_UserListCombinedRuleOperator = 1
	// A AND B.
	UserListCombinedRuleOperatorEnum_AND UserListCombinedRuleOperatorEnum_UserListCombinedRuleOperator = 2
	// A AND NOT B.
	UserListCombinedRuleOperatorEnum_AND_NOT UserListCombinedRuleOperatorEnum_UserListCombinedRuleOperator = 3
)

var UserListCombinedRuleOperatorEnum_UserListCombinedRuleOperator_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AND",
	3: "AND_NOT",
}

var UserListCombinedRuleOperatorEnum_UserListCombinedRuleOperator_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"AND":         2,
	"AND_NOT":     3,
}

func (x UserListCombinedRuleOperatorEnum_UserListCombinedRuleOperator) String() string {
	return proto.EnumName(UserListCombinedRuleOperatorEnum_UserListCombinedRuleOperator_name, int32(x))
}

func (UserListCombinedRuleOperatorEnum_UserListCombinedRuleOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c43031e2c954a0ca, []int{0, 0}
}

// Logical operator connecting two rules.
type UserListCombinedRuleOperatorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListCombinedRuleOperatorEnum) Reset()         { *m = UserListCombinedRuleOperatorEnum{} }
func (m *UserListCombinedRuleOperatorEnum) String() string { return proto.CompactTextString(m) }
func (*UserListCombinedRuleOperatorEnum) ProtoMessage()    {}
func (*UserListCombinedRuleOperatorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_c43031e2c954a0ca, []int{0}
}

func (m *UserListCombinedRuleOperatorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListCombinedRuleOperatorEnum.Unmarshal(m, b)
}
func (m *UserListCombinedRuleOperatorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListCombinedRuleOperatorEnum.Marshal(b, m, deterministic)
}
func (m *UserListCombinedRuleOperatorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListCombinedRuleOperatorEnum.Merge(m, src)
}
func (m *UserListCombinedRuleOperatorEnum) XXX_Size() int {
	return xxx_messageInfo_UserListCombinedRuleOperatorEnum.Size(m)
}
func (m *UserListCombinedRuleOperatorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListCombinedRuleOperatorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListCombinedRuleOperatorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.UserListCombinedRuleOperatorEnum_UserListCombinedRuleOperator", UserListCombinedRuleOperatorEnum_UserListCombinedRuleOperator_name, UserListCombinedRuleOperatorEnum_UserListCombinedRuleOperator_value)
	proto.RegisterType((*UserListCombinedRuleOperatorEnum)(nil), "google.ads.googleads.v1.enums.UserListCombinedRuleOperatorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/user_list_combined_rule_operator.proto", fileDescriptor_c43031e2c954a0ca)
}

var fileDescriptor_c43031e2c954a0ca = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xc1, 0x6a, 0xea, 0x40,
	0x14, 0x7d, 0x46, 0x78, 0xc2, 0xb8, 0x68, 0xc8, 0xb2, 0x28, 0x54, 0x3f, 0x60, 0x42, 0xe8, 0x6e,
	0xba, 0x1a, 0x8d, 0x15, 0x69, 0x19, 0xc5, 0x56, 0x0b, 0x25, 0x10, 0xa2, 0x19, 0x42, 0x20, 0x99,
	0x1b, 0xe6, 0x26, 0x7e, 0x50, 0x97, 0xfd, 0x94, 0x7e, 0x4a, 0x3f, 0xa0, 0xeb, 0x92, 0x19, 0x75,
	0xd7, 0x6c, 0x86, 0xc3, 0x9c, 0x73, 0xcf, 0xb9, 0xf7, 0x90, 0x30, 0x03, 0xc8, 0x0a, 0xe9, 0x27,
	0x29, 0xfa, 0x16, 0xb6, 0xe8, 0x14, 0xf8, 0x52, 0x35, 0x25, 0xfa, 0x0d, 0x4a, 0x1d, 0x17, 0x39,
	0xd6, 0xf1, 0x11, 0xca, 0x43, 0xae, 0x64, 0x1a, 0xeb, 0xa6, 0x90, 0x31, 0x54, 0x52, 0x27, 0x35,
	0x68, 0x5a, 0x69, 0xa8, 0xc1, 0x1b, 0xdb, 0x51, 0x9a, 0xa4, 0x48, 0xaf, 0x2e, 0xf4, 0x14, 0x50,
	0xe3, 0x72, 0x3b, 0xba, 0x84, 0x54, 0xb9, 0x9f, 0x28, 0x05, 0x75, 0x52, 0xe7, 0xa0, 0xd0, 0x0e,
	0x4f, 0x4f, 0xe4, 0x6e, 0x87, 0x52, 0x3f, 0xe7, 0x58, 0xcf, 0xcf, 0x21, 0xdb, 0xa6, 0x90, 0xeb,
	0x73, 0xc4, 0x42, 0x35, 0xe5, 0x74, 0x4b, 0x46, 0x5d, 0x1a, 0xef, 0x86, 0x0c, 0x77, 0xe2, 0x65,
	0xb3, 0x98, 0xaf, 0x1e, 0x57, 0x8b, 0xd0, 0xfd, 0xe7, 0x0d, 0xc9, 0x60, 0x27, 0x9e, 0xc4, 0xfa,
	0x4d, 0xb8, 0x3d, 0x6f, 0x40, 0xfa, 0x5c, 0x84, 0xae, 0xd3, 0xfe, 0x72, 0x11, 0xc6, 0x62, 0xfd,
	0xea, 0xf6, 0x67, 0x3f, 0x3d, 0x32, 0x39, 0x42, 0x49, 0x3b, 0x77, 0x9f, 0x4d, 0xba, 0x72, 0x37,
	0xed, 0x01, 0x9b, 0xde, 0xfb, 0xec, 0xec, 0x91, 0x41, 0x91, 0xa8, 0x8c, 0x82, 0xce, 0xfc, 0x4c,
	0x2a, 0x73, 0xde, 0xa5, 0xd5, 0x2a, 0xc7, 0x3f, 0x4a, 0x7e, 0x30, 0xef, 0x87, 0xd3, 0x5f, 0x72,
	0xfe, 0xe9, 0x8c, 0x97, 0xd6, 0x8a, 0xa7, 0x48, 0x2d, 0x6c, 0xd1, 0x3e, 0xa0, 0x6d, 0x0f, 0xf8,
	0x75, 0xe1, 0x23, 0x9e, 0x62, 0x74, 0xe5, 0xa3, 0x7d, 0x10, 0x19, 0xfe, 0xdb, 0x99, 0xd8, 0x4f,
	0xc6, 0x78, 0x8a, 0x8c, 0x5d, 0x15, 0x8c, 0xed, 0x03, 0xc6, 0x8c, 0xe6, 0xf0, 0xdf, 0x2c, 0x76,
	0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x53, 0xc2, 0xa7, 0xfc, 0x01, 0x00, 0x00,
}
