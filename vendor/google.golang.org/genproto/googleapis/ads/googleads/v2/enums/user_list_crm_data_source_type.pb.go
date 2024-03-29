// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/user_list_crm_data_source_type.proto

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

// Enum describing possible user list crm data source type.
type UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType int32

const (
	// Not specified.
	UserListCrmDataSourceTypeEnum_UNSPECIFIED UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType = 0
	// Used for return value only. Represents value unknown in this version.
	UserListCrmDataSourceTypeEnum_UNKNOWN UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType = 1
	// The uploaded data is first-party data.
	UserListCrmDataSourceTypeEnum_FIRST_PARTY UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType = 2
	// The uploaded data is from a third-party credit bureau.
	UserListCrmDataSourceTypeEnum_THIRD_PARTY_CREDIT_BUREAU UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType = 3
	// The uploaded data is from a third-party voter file.
	UserListCrmDataSourceTypeEnum_THIRD_PARTY_VOTER_FILE UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType = 4
)

var UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "FIRST_PARTY",
	3: "THIRD_PARTY_CREDIT_BUREAU",
	4: "THIRD_PARTY_VOTER_FILE",
}

var UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType_value = map[string]int32{
	"UNSPECIFIED":               0,
	"UNKNOWN":                   1,
	"FIRST_PARTY":               2,
	"THIRD_PARTY_CREDIT_BUREAU": 3,
	"THIRD_PARTY_VOTER_FILE":    4,
}

func (x UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType) String() string {
	return proto.EnumName(UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType_name, int32(x))
}

func (UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_23ec3967d6ce16a1, []int{0, 0}
}

// Indicates source of Crm upload data.
type UserListCrmDataSourceTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListCrmDataSourceTypeEnum) Reset()         { *m = UserListCrmDataSourceTypeEnum{} }
func (m *UserListCrmDataSourceTypeEnum) String() string { return proto.CompactTextString(m) }
func (*UserListCrmDataSourceTypeEnum) ProtoMessage()    {}
func (*UserListCrmDataSourceTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_23ec3967d6ce16a1, []int{0}
}

func (m *UserListCrmDataSourceTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListCrmDataSourceTypeEnum.Unmarshal(m, b)
}
func (m *UserListCrmDataSourceTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListCrmDataSourceTypeEnum.Marshal(b, m, deterministic)
}
func (m *UserListCrmDataSourceTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListCrmDataSourceTypeEnum.Merge(m, src)
}
func (m *UserListCrmDataSourceTypeEnum) XXX_Size() int {
	return xxx_messageInfo_UserListCrmDataSourceTypeEnum.Size(m)
}
func (m *UserListCrmDataSourceTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListCrmDataSourceTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListCrmDataSourceTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType", UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType_name, UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType_value)
	proto.RegisterType((*UserListCrmDataSourceTypeEnum)(nil), "google.ads.googleads.v2.enums.UserListCrmDataSourceTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/user_list_crm_data_source_type.proto", fileDescriptor_23ec3967d6ce16a1)
}

var fileDescriptor_23ec3967d6ce16a1 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x6a, 0xab, 0x40,
	0x14, 0x7d, 0x9a, 0x47, 0x0b, 0x93, 0x45, 0xc5, 0x45, 0x21, 0xa1, 0x16, 0x92, 0x0f, 0x18, 0xc1,
	0xee, 0xa6, 0x2b, 0x8d, 0x26, 0x95, 0x06, 0x23, 0x46, 0x2d, 0x2d, 0xc2, 0x30, 0x8d, 0x83, 0x08,
	0xd1, 0x11, 0x67, 0x0c, 0xe4, 0x03, 0xfa, 0x23, 0xdd, 0x14, 0xfa, 0x29, 0xfd, 0x94, 0x6e, 0xfb,
	0x03, 0x45, 0x6d, 0x42, 0x37, 0xe9, 0x66, 0x38, 0xcc, 0x39, 0xf7, 0xdc, 0x7b, 0xcf, 0x05, 0x56,
	0xc6, 0x58, 0xb6, 0xa5, 0x3a, 0x49, 0xb9, 0xde, 0xc3, 0x16, 0xed, 0x0c, 0x9d, 0x96, 0x4d, 0xc1,
	0xf5, 0x86, 0xd3, 0x1a, 0x6f, 0x73, 0x2e, 0xf0, 0xa6, 0x2e, 0x70, 0x4a, 0x04, 0xc1, 0x9c, 0x35,
	0xf5, 0x86, 0x62, 0xb1, 0xaf, 0x28, 0xac, 0x6a, 0x26, 0x98, 0xaa, 0xf5, 0x85, 0x90, 0xa4, 0x1c,
	0x1e, 0x3d, 0xe0, 0xce, 0x80, 0x9d, 0xc7, 0xf8, 0xea, 0xd0, 0xa2, 0xca, 0x75, 0x52, 0x96, 0x4c,
	0x10, 0x91, 0xb3, 0x92, 0xf7, 0xc5, 0xd3, 0x37, 0x09, 0x68, 0x11, 0xa7, 0xf5, 0x32, 0xe7, 0x62,
	0x56, 0x17, 0x36, 0x11, 0x64, 0xdd, 0x75, 0x08, 0xf7, 0x15, 0x75, 0xca, 0xa6, 0x98, 0xbe, 0x48,
	0x60, 0x74, 0x52, 0xa1, 0x5e, 0x80, 0x61, 0xe4, 0xad, 0x7d, 0x67, 0xe6, 0xce, 0x5d, 0xc7, 0x56,
	0xfe, 0xa9, 0x43, 0x70, 0x1e, 0x79, 0xf7, 0xde, 0xea, 0xc1, 0x53, 0xa4, 0x96, 0x9d, 0xbb, 0xc1,
	0x3a, 0xc4, 0xbe, 0x19, 0x84, 0x8f, 0x8a, 0xac, 0x6a, 0x60, 0x14, 0xde, 0xb9, 0x81, 0xdd, 0x7f,
	0xe0, 0x59, 0xe0, 0xd8, 0x6e, 0x88, 0xad, 0x28, 0x70, 0xcc, 0x48, 0x19, 0xa8, 0x63, 0x70, 0xf9,
	0x9b, 0x8e, 0x57, 0xa1, 0x13, 0xe0, 0xb9, 0xbb, 0x74, 0x94, 0xff, 0xd6, 0x97, 0x04, 0x26, 0x1b,
	0x56, 0xc0, 0x3f, 0xb7, 0xb5, 0xae, 0x4f, 0x8e, 0xea, 0xb7, 0xfb, 0xfa, 0xd2, 0xd3, 0x4f, 0xe4,
	0x30, 0x63, 0x5b, 0x52, 0x66, 0x90, 0xd5, 0x99, 0x9e, 0xd1, 0xb2, 0x4b, 0xe3, 0x70, 0x82, 0x2a,
	0xe7, 0x27, 0x2e, 0x72, 0xdb, 0xbd, 0xaf, 0xf2, 0x60, 0x61, 0x9a, 0xef, 0xb2, 0xb6, 0xe8, 0xad,
	0xcc, 0x94, 0xc3, 0x1e, 0xb6, 0x28, 0x36, 0x60, 0x1b, 0x1c, 0xff, 0x38, 0xf0, 0x89, 0x99, 0xf2,
	0xe4, 0xc8, 0x27, 0xb1, 0x91, 0x74, 0xfc, 0xa7, 0x3c, 0xe9, 0x3f, 0x11, 0x32, 0x53, 0x8e, 0xd0,
	0x51, 0x81, 0x50, 0x6c, 0x20, 0xd4, 0x69, 0x9e, 0xcf, 0xba, 0xc1, 0x6e, 0xbe, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xdf, 0xd8, 0x79, 0xb4, 0x29, 0x02, 0x00, 0x00,
}
