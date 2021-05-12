// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/user_list_string_rule_item_operator.proto

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

// Enum describing possible user list string rule item operators.
type UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator int32

const (
	// Not specified.
	UserListStringRuleItemOperatorEnum_UNSPECIFIED UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 0
	// Used for return value only. Represents value unknown in this version.
	UserListStringRuleItemOperatorEnum_UNKNOWN UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 1
	// Contains.
	UserListStringRuleItemOperatorEnum_CONTAINS UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 2
	// Equals.
	UserListStringRuleItemOperatorEnum_EQUALS UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 3
	// Starts with.
	UserListStringRuleItemOperatorEnum_STARTS_WITH UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 4
	// Ends with.
	UserListStringRuleItemOperatorEnum_ENDS_WITH UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 5
	// Not equals.
	UserListStringRuleItemOperatorEnum_NOT_EQUALS UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 6
	// Not contains.
	UserListStringRuleItemOperatorEnum_NOT_CONTAINS UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 7
	// Not starts with.
	UserListStringRuleItemOperatorEnum_NOT_STARTS_WITH UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 8
	// Not ends with.
	UserListStringRuleItemOperatorEnum_NOT_ENDS_WITH UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 9
)

var UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CONTAINS",
	3: "EQUALS",
	4: "STARTS_WITH",
	5: "ENDS_WITH",
	6: "NOT_EQUALS",
	7: "NOT_CONTAINS",
	8: "NOT_STARTS_WITH",
	9: "NOT_ENDS_WITH",
}

var UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator_value = map[string]int32{
	"UNSPECIFIED":     0,
	"UNKNOWN":         1,
	"CONTAINS":        2,
	"EQUALS":          3,
	"STARTS_WITH":     4,
	"ENDS_WITH":       5,
	"NOT_EQUALS":      6,
	"NOT_CONTAINS":    7,
	"NOT_STARTS_WITH": 8,
	"NOT_ENDS_WITH":   9,
}

func (x UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator) String() string {
	return proto.EnumName(UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator_name, int32(x))
}

func (UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_71706bdcf45a4e95, []int{0, 0}
}

// Supported rule operator for string type.
type UserListStringRuleItemOperatorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListStringRuleItemOperatorEnum) Reset()         { *m = UserListStringRuleItemOperatorEnum{} }
func (m *UserListStringRuleItemOperatorEnum) String() string { return proto.CompactTextString(m) }
func (*UserListStringRuleItemOperatorEnum) ProtoMessage()    {}
func (*UserListStringRuleItemOperatorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_71706bdcf45a4e95, []int{0}
}

func (m *UserListStringRuleItemOperatorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListStringRuleItemOperatorEnum.Unmarshal(m, b)
}
func (m *UserListStringRuleItemOperatorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListStringRuleItemOperatorEnum.Marshal(b, m, deterministic)
}
func (m *UserListStringRuleItemOperatorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListStringRuleItemOperatorEnum.Merge(m, src)
}
func (m *UserListStringRuleItemOperatorEnum) XXX_Size() int {
	return xxx_messageInfo_UserListStringRuleItemOperatorEnum.Size(m)
}
func (m *UserListStringRuleItemOperatorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListStringRuleItemOperatorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListStringRuleItemOperatorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator", UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator_name, UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator_value)
	proto.RegisterType((*UserListStringRuleItemOperatorEnum)(nil), "google.ads.googleads.v1.enums.UserListStringRuleItemOperatorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/user_list_string_rule_item_operator.proto", fileDescriptor_71706bdcf45a4e95)
}

var fileDescriptor_71706bdcf45a4e95 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x41, 0x8f, 0x94, 0x30,
	0x14, 0x16, 0x56, 0x67, 0x77, 0xbb, 0xbb, 0x2e, 0xd6, 0x9b, 0x71, 0x4d, 0x76, 0xbc, 0x97, 0x10,
	0x6f, 0xf5, 0xd4, 0xd9, 0xc1, 0x91, 0x38, 0x29, 0xe3, 0x00, 0x33, 0x89, 0x21, 0x21, 0x28, 0x0d,
	0x21, 0x81, 0x96, 0xb4, 0x65, 0x7e, 0x90, 0x47, 0x7f, 0x86, 0x47, 0xff, 0x89, 0xfe, 0x02, 0x8f,
	0x86, 0xc2, 0x10, 0x2f, 0x3b, 0x17, 0xf2, 0xbd, 0xf7, 0xbe, 0xef, 0x7b, 0xf4, 0x7b, 0x60, 0x55,
	0x0a, 0x51, 0xd6, 0xcc, 0xcd, 0x0b, 0xe5, 0x0e, 0xb0, 0x47, 0x07, 0xcf, 0x65, 0xbc, 0x6b, 0x94,
	0xdb, 0x29, 0x26, 0xb3, 0xba, 0x52, 0x3a, 0x53, 0x5a, 0x56, 0xbc, 0xcc, 0x64, 0x57, 0xb3, 0xac,
	0xd2, 0xac, 0xc9, 0x44, 0xcb, 0x64, 0xae, 0x85, 0x44, 0xad, 0x14, 0x5a, 0xc0, 0xbb, 0x41, 0x8d,
	0xf2, 0x42, 0xa1, 0xc9, 0x08, 0x1d, 0x3c, 0x64, 0x8c, 0x5e, 0xbd, 0x3e, 0xee, 0x69, 0x2b, 0x37,
	0xe7, 0x5c, 0xe8, 0x5c, 0x57, 0x82, 0xab, 0x41, 0x3c, 0xff, 0x6d, 0x81, 0x79, 0xa2, 0x98, 0x5c,
	0x57, 0x4a, 0x47, 0x66, 0xd1, 0xb6, 0xab, 0x59, 0xa0, 0x59, 0x13, 0x8e, 0x5b, 0x7c, 0xde, 0x35,
	0xf3, 0x9f, 0x16, 0x78, 0x73, 0x9a, 0x06, 0x6f, 0xc1, 0x55, 0x42, 0xa3, 0x8d, 0xff, 0x10, 0x7c,
	0x08, 0xfc, 0xa5, 0xf3, 0x04, 0x5e, 0x81, 0xf3, 0x84, 0x7e, 0xa2, 0xe1, 0x9e, 0x3a, 0x16, 0xbc,
	0x06, 0x17, 0x0f, 0x21, 0x8d, 0x49, 0x40, 0x23, 0xc7, 0x86, 0x00, 0xcc, 0xfc, 0xcf, 0x09, 0x59,
	0x47, 0xce, 0x59, 0xaf, 0x8b, 0x62, 0xb2, 0x8d, 0xa3, 0x6c, 0x1f, 0xc4, 0x1f, 0x9d, 0xa7, 0xf0,
	0x06, 0x5c, 0xfa, 0x74, 0x39, 0x96, 0xcf, 0xe0, 0x73, 0x00, 0x68, 0x18, 0x67, 0x23, 0x7f, 0x06,
	0x1d, 0x70, 0xdd, 0xd7, 0x93, 0xdb, 0x39, 0x7c, 0x09, 0x6e, 0xfb, 0xce, 0xff, 0x2e, 0x17, 0xf0,
	0x05, 0xb8, 0x31, 0xb2, 0xc9, 0xe9, 0x72, 0xf1, 0xd7, 0x02, 0xf7, 0xdf, 0x44, 0x83, 0x4e, 0xe6,
	0xb5, 0x78, 0x7b, 0xfa, 0x9d, 0x9b, 0x3e, 0xb6, 0x8d, 0xf5, 0x65, 0x31, 0xba, 0x94, 0xa2, 0xce,
	0x79, 0x89, 0x84, 0x2c, 0xdd, 0x92, 0x71, 0x13, 0xea, 0xf1, 0x9c, 0x6d, 0xa5, 0x1e, 0xb9, 0xee,
	0x7b, 0xf3, 0xfd, 0x6e, 0x9f, 0xad, 0x08, 0xf9, 0x61, 0xdf, 0xad, 0x06, 0x2b, 0x52, 0x28, 0x34,
	0xc0, 0x1e, 0xed, 0x3c, 0xd4, 0x47, 0xaf, 0x7e, 0x1d, 0xe7, 0x29, 0x29, 0x54, 0x3a, 0xcd, 0xd3,
	0x9d, 0x97, 0x9a, 0xf9, 0x1f, 0xfb, 0x7e, 0x68, 0x62, 0x4c, 0x0a, 0x85, 0xf1, 0xc4, 0xc0, 0x78,
	0xe7, 0x61, 0x6c, 0x38, 0x5f, 0x67, 0xe6, 0xc7, 0xde, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xf8,
	0xb7, 0x29, 0xc0, 0x75, 0x02, 0x00, 0x00,
}
