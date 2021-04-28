// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/customer_client_link_error.proto

package errors

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

// Enum describing possible CustomerClientLink errors.
type CustomerClientLinkErrorEnum_CustomerClientLinkError int32

const (
	// Enum unspecified.
	CustomerClientLinkErrorEnum_UNSPECIFIED CustomerClientLinkErrorEnum_CustomerClientLinkError = 0
	// The received error code is not known in this version.
	CustomerClientLinkErrorEnum_UNKNOWN CustomerClientLinkErrorEnum_CustomerClientLinkError = 1
	// Trying to manage a client that already in being managed by customer.
	CustomerClientLinkErrorEnum_CLIENT_ALREADY_INVITED_BY_THIS_MANAGER CustomerClientLinkErrorEnum_CustomerClientLinkError = 2
	// Already managed by some other manager in the hierarchy.
	CustomerClientLinkErrorEnum_CLIENT_ALREADY_MANAGED_IN_HIERARCHY CustomerClientLinkErrorEnum_CustomerClientLinkError = 3
	// Attempt to create a cycle in the hierarchy.
	CustomerClientLinkErrorEnum_CYCLIC_LINK_NOT_ALLOWED CustomerClientLinkErrorEnum_CustomerClientLinkError = 4
	// Managed accounts has the maximum number of linked accounts.
	CustomerClientLinkErrorEnum_CUSTOMER_HAS_TOO_MANY_ACCOUNTS CustomerClientLinkErrorEnum_CustomerClientLinkError = 5
	// Invitor has the maximum pending invitations.
	CustomerClientLinkErrorEnum_CLIENT_HAS_TOO_MANY_INVITATIONS CustomerClientLinkErrorEnum_CustomerClientLinkError = 6
	// Attempt to change hidden status of a link that is not active.
	CustomerClientLinkErrorEnum_CANNOT_HIDE_OR_UNHIDE_MANAGER_ACCOUNTS CustomerClientLinkErrorEnum_CustomerClientLinkError = 7
	// Parent manager account has the maximum number of linked accounts.
	CustomerClientLinkErrorEnum_CUSTOMER_HAS_TOO_MANY_ACCOUNTS_AT_MANAGER CustomerClientLinkErrorEnum_CustomerClientLinkError = 8
)

var CustomerClientLinkErrorEnum_CustomerClientLinkError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CLIENT_ALREADY_INVITED_BY_THIS_MANAGER",
	3: "CLIENT_ALREADY_MANAGED_IN_HIERARCHY",
	4: "CYCLIC_LINK_NOT_ALLOWED",
	5: "CUSTOMER_HAS_TOO_MANY_ACCOUNTS",
	6: "CLIENT_HAS_TOO_MANY_INVITATIONS",
	7: "CANNOT_HIDE_OR_UNHIDE_MANAGER_ACCOUNTS",
	8: "CUSTOMER_HAS_TOO_MANY_ACCOUNTS_AT_MANAGER",
}

var CustomerClientLinkErrorEnum_CustomerClientLinkError_value = map[string]int32{
	"UNSPECIFIED":                               0,
	"UNKNOWN":                                   1,
	"CLIENT_ALREADY_INVITED_BY_THIS_MANAGER":    2,
	"CLIENT_ALREADY_MANAGED_IN_HIERARCHY":       3,
	"CYCLIC_LINK_NOT_ALLOWED":                   4,
	"CUSTOMER_HAS_TOO_MANY_ACCOUNTS":            5,
	"CLIENT_HAS_TOO_MANY_INVITATIONS":           6,
	"CANNOT_HIDE_OR_UNHIDE_MANAGER_ACCOUNTS":    7,
	"CUSTOMER_HAS_TOO_MANY_ACCOUNTS_AT_MANAGER": 8,
}

func (x CustomerClientLinkErrorEnum_CustomerClientLinkError) String() string {
	return proto.EnumName(CustomerClientLinkErrorEnum_CustomerClientLinkError_name, int32(x))
}

func (CustomerClientLinkErrorEnum_CustomerClientLinkError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_05999249da6418bd, []int{0, 0}
}

// Container for enum describing possible CustomeClientLink errors.
type CustomerClientLinkErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerClientLinkErrorEnum) Reset()         { *m = CustomerClientLinkErrorEnum{} }
func (m *CustomerClientLinkErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CustomerClientLinkErrorEnum) ProtoMessage()    {}
func (*CustomerClientLinkErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_05999249da6418bd, []int{0}
}

func (m *CustomerClientLinkErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerClientLinkErrorEnum.Unmarshal(m, b)
}
func (m *CustomerClientLinkErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerClientLinkErrorEnum.Marshal(b, m, deterministic)
}
func (m *CustomerClientLinkErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerClientLinkErrorEnum.Merge(m, src)
}
func (m *CustomerClientLinkErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CustomerClientLinkErrorEnum.Size(m)
}
func (m *CustomerClientLinkErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerClientLinkErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerClientLinkErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.CustomerClientLinkErrorEnum_CustomerClientLinkError", CustomerClientLinkErrorEnum_CustomerClientLinkError_name, CustomerClientLinkErrorEnum_CustomerClientLinkError_value)
	proto.RegisterType((*CustomerClientLinkErrorEnum)(nil), "google.ads.googleads.v1.errors.CustomerClientLinkErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/customer_client_link_error.proto", fileDescriptor_05999249da6418bd)
}

var fileDescriptor_05999249da6418bd = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0x77, 0x53, 0x6d, 0x65, 0x7a, 0x30, 0xe4, 0x52, 0xb0, 0x65, 0x85, 0x14, 0x14, 0x05, 0x13,
	0x82, 0xb7, 0x78, 0x90, 0xd9, 0xc9, 0xb8, 0x19, 0x9a, 0x4e, 0x96, 0xfc, 0xd9, 0x12, 0x59, 0x78,
	0xc4, 0x4d, 0x08, 0xa1, 0xbb, 0x99, 0x25, 0x93, 0xf6, 0x03, 0x79, 0xf4, 0xa3, 0x08, 0x7e, 0x11,
	0xc1, 0xb3, 0x57, 0x49, 0x66, 0x77, 0xc5, 0xc2, 0xf6, 0x94, 0x1f, 0x79, 0xbf, 0x7f, 0xc9, 0x7b,
	0xe8, 0x53, 0x25, 0x44, 0xb5, 0x2a, 0xed, 0xbc, 0x90, 0xb6, 0x82, 0x3d, 0xba, 0x77, 0xec, 0xb2,
	0x6d, 0x45, 0x2b, 0xed, 0xe5, 0x9d, 0xec, 0xc4, 0xba, 0x6c, 0x61, 0xb9, 0xaa, 0xcb, 0xa6, 0x83,
	0x55, 0xdd, 0xdc, 0xc2, 0x30, 0xb3, 0x36, 0xad, 0xe8, 0x84, 0x31, 0x56, 0x2a, 0x2b, 0x2f, 0xa4,
	0xb5, 0x37, 0xb0, 0xee, 0x1d, 0x4b, 0x19, 0xbc, 0xbc, 0xd8, 0x05, 0x6c, 0x6a, 0x3b, 0x6f, 0x1a,
	0xd1, 0xe5, 0x5d, 0x2d, 0x1a, 0xa9, 0xd4, 0xe6, 0x6f, 0x0d, 0x9d, 0x93, 0x6d, 0x04, 0x19, 0x12,
	0x82, 0xba, 0xb9, 0xa5, 0xbd, 0x94, 0x36, 0x77, 0x6b, 0xf3, 0xa7, 0x86, 0xce, 0x0e, 0xcc, 0x8d,
	0x17, 0xe8, 0x34, 0xe5, 0xf1, 0x8c, 0x12, 0xf6, 0x99, 0x51, 0x4f, 0x7f, 0x62, 0x9c, 0xa2, 0x93,
	0x94, 0x5f, 0xf1, 0xf0, 0x86, 0xeb, 0x23, 0xe3, 0x1d, 0x7a, 0x4d, 0x02, 0x46, 0x79, 0x02, 0x38,
	0x88, 0x28, 0xf6, 0x32, 0x60, 0x7c, 0xce, 0x12, 0xea, 0xc1, 0x24, 0x83, 0xc4, 0x67, 0x31, 0x5c,
	0x63, 0x8e, 0xa7, 0x34, 0xd2, 0x35, 0xe3, 0x0d, 0xba, 0x7c, 0xc0, 0x55, 0x33, 0x0f, 0x18, 0x07,
	0x9f, 0xd1, 0x08, 0x47, 0xc4, 0xcf, 0xf4, 0x23, 0xe3, 0x1c, 0x9d, 0x91, 0x8c, 0x04, 0x8c, 0x40,
	0xc0, 0xf8, 0x15, 0xf0, 0xb0, 0x57, 0x04, 0xe1, 0x0d, 0xf5, 0xf4, 0xa7, 0x86, 0x89, 0xc6, 0x24,
	0x8d, 0x93, 0xf0, 0x9a, 0x46, 0xe0, 0xe3, 0x18, 0x92, 0x30, 0xec, 0x7d, 0x32, 0xc0, 0x84, 0x84,
	0x29, 0x4f, 0x62, 0xfd, 0x99, 0x71, 0x89, 0x5e, 0x6d, 0x93, 0xfe, 0x63, 0x0c, 0xd5, 0x70, 0xc2,
	0x42, 0x1e, 0xeb, 0xc7, 0x43, 0x75, 0xcc, 0x7b, 0x73, 0x9f, 0x79, 0x14, 0xc2, 0x08, 0x52, 0x3e,
	0xa0, 0x6d, 0xe3, 0x7f, 0x86, 0x27, 0xc6, 0x7b, 0xf4, 0xf6, 0xf1, 0x50, 0xc0, 0xc9, 0xfe, 0x4b,
	0x9f, 0x4f, 0xfe, 0x8c, 0x90, 0xb9, 0x14, 0x6b, 0xeb, 0xf1, 0xa5, 0x4d, 0x2e, 0x0e, 0xfc, 0xf3,
	0x59, 0xbf, 0xb4, 0xd9, 0xe8, 0x8b, 0xb7, 0xd5, 0x57, 0x62, 0x95, 0x37, 0x95, 0x25, 0xda, 0xca,
	0xae, 0xca, 0x66, 0x58, 0xe9, 0xee, 0x8a, 0x36, 0xb5, 0x3c, 0x74, 0x54, 0x1f, 0xd5, 0xe3, 0x9b,
	0x76, 0x34, 0xc5, 0xf8, 0xbb, 0x36, 0x9e, 0x2a, 0x33, 0x5c, 0x48, 0x4b, 0xc1, 0x1e, 0xcd, 0x1d,
	0x6b, 0x88, 0x94, 0x3f, 0x76, 0x84, 0x05, 0x2e, 0xe4, 0x62, 0x4f, 0x58, 0xcc, 0x9d, 0x85, 0x22,
	0xfc, 0xd2, 0x4c, 0xf5, 0xd6, 0x75, 0x71, 0x21, 0x5d, 0x77, 0x4f, 0x71, 0xdd, 0xb9, 0xe3, 0xba,
	0x8a, 0xf4, 0xf5, 0x78, 0x68, 0xf7, 0xe1, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x0f, 0x8b,
	0x6a, 0xf1, 0x02, 0x00, 0x00,
}
