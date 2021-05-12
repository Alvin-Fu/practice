// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/access_invitation_error.proto

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

// Enum describing possible AccessInvitation errors.
type AccessInvitationErrorEnum_AccessInvitationError int32

const (
	// Enum unspecified.
	AccessInvitationErrorEnum_UNSPECIFIED AccessInvitationErrorEnum_AccessInvitationError = 0
	// The received error code is not known in this version.
	AccessInvitationErrorEnum_UNKNOWN AccessInvitationErrorEnum_AccessInvitationError = 1
	// The email address is invalid for sending an invitation.
	AccessInvitationErrorEnum_INVALID_EMAIL_ADDRESS AccessInvitationErrorEnum_AccessInvitationError = 2
	// Email address already has access to this customer.
	AccessInvitationErrorEnum_EMAIL_ADDRESS_ALREADY_HAS_ACCESS AccessInvitationErrorEnum_AccessInvitationError = 3
)

var AccessInvitationErrorEnum_AccessInvitationError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "INVALID_EMAIL_ADDRESS",
	3: "EMAIL_ADDRESS_ALREADY_HAS_ACCESS",
}

var AccessInvitationErrorEnum_AccessInvitationError_value = map[string]int32{
	"UNSPECIFIED":                      0,
	"UNKNOWN":                          1,
	"INVALID_EMAIL_ADDRESS":            2,
	"EMAIL_ADDRESS_ALREADY_HAS_ACCESS": 3,
}

func (x AccessInvitationErrorEnum_AccessInvitationError) String() string {
	return proto.EnumName(AccessInvitationErrorEnum_AccessInvitationError_name, int32(x))
}

func (AccessInvitationErrorEnum_AccessInvitationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e255edc7bfd41f33, []int{0, 0}
}

// Container for enum describing possible AccessInvitation errors.
type AccessInvitationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessInvitationErrorEnum) Reset()         { *m = AccessInvitationErrorEnum{} }
func (m *AccessInvitationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AccessInvitationErrorEnum) ProtoMessage()    {}
func (*AccessInvitationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_e255edc7bfd41f33, []int{0}
}

func (m *AccessInvitationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessInvitationErrorEnum.Unmarshal(m, b)
}
func (m *AccessInvitationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessInvitationErrorEnum.Marshal(b, m, deterministic)
}
func (m *AccessInvitationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessInvitationErrorEnum.Merge(m, src)
}
func (m *AccessInvitationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AccessInvitationErrorEnum.Size(m)
}
func (m *AccessInvitationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessInvitationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AccessInvitationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.AccessInvitationErrorEnum_AccessInvitationError", AccessInvitationErrorEnum_AccessInvitationError_name, AccessInvitationErrorEnum_AccessInvitationError_value)
	proto.RegisterType((*AccessInvitationErrorEnum)(nil), "google.ads.googleads.v2.errors.AccessInvitationErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/access_invitation_error.proto", fileDescriptor_e255edc7bfd41f33)
}

var fileDescriptor_e255edc7bfd41f33 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xff, 0x49, 0xe1, 0x2f, 0x6c, 0x0f, 0x96, 0x40, 0x0f, 0x2d, 0x52, 0x24, 0x78, 0xde,
	0x40, 0xbc, 0xad, 0x5e, 0xa6, 0xcd, 0x5a, 0x83, 0x35, 0x96, 0x86, 0x46, 0x94, 0xc0, 0xb2, 0x36,
	0x61, 0x09, 0xb4, 0xbb, 0x25, 0x5b, 0xf3, 0x22, 0xbe, 0x81, 0x47, 0x1f, 0xc5, 0x47, 0xf1, 0xee,
	0x5d, 0x92, 0xb5, 0x01, 0xa1, 0x7a, 0xda, 0x8f, 0x99, 0xef, 0xf7, 0xed, 0xcc, 0xa0, 0x4b, 0xa1,
	0x94, 0x58, 0xe7, 0x1e, 0xcf, 0xb4, 0x67, 0x64, 0xad, 0x2a, 0xdf, 0xcb, 0xcb, 0x52, 0x95, 0xda,
	0xe3, 0xab, 0x55, 0xae, 0x35, 0x2b, 0x64, 0x55, 0xec, 0xf8, 0xae, 0x50, 0x92, 0x35, 0x0d, 0xbc,
	0x2d, 0xd5, 0x4e, 0x39, 0x23, 0x83, 0x60, 0x9e, 0x69, 0xdc, 0xd2, 0xb8, 0xf2, 0xb1, 0xa1, 0x87,
	0x27, 0xfb, 0xf4, 0x6d, 0xe1, 0x71, 0x29, 0x95, 0x89, 0xd0, 0x86, 0x76, 0x5f, 0x2c, 0x34, 0x80,
	0x26, 0x3f, 0x6c, 0xe3, 0x69, 0x0d, 0x52, 0xf9, 0xbc, 0x71, 0x2b, 0xd4, 0x3f, 0xd8, 0x74, 0x8e,
	0x51, 0x77, 0x19, 0xc5, 0x73, 0x3a, 0x09, 0xaf, 0x42, 0x1a, 0xf4, 0xfe, 0x39, 0x5d, 0x74, 0xb4,
	0x8c, 0x6e, 0xa2, 0xbb, 0xfb, 0xa8, 0x67, 0x39, 0x03, 0xd4, 0x0f, 0xa3, 0x04, 0x66, 0x61, 0xc0,
	0xe8, 0x2d, 0x84, 0x33, 0x06, 0x41, 0xb0, 0xa0, 0x71, 0xdc, 0xb3, 0x9d, 0x33, 0x74, 0xfa, 0xa3,
	0xc4, 0x60, 0xb6, 0xa0, 0x10, 0x3c, 0xb0, 0x6b, 0x88, 0x19, 0x4c, 0x26, 0xb5, 0xab, 0x33, 0xfe,
	0xb4, 0x90, 0xbb, 0x52, 0x1b, 0xfc, 0xf7, 0x6a, 0xe3, 0xe1, 0xc1, 0xe1, 0xe6, 0xf5, 0x62, 0x73,
	0xeb, 0x31, 0xf8, 0xa6, 0x85, 0x5a, 0x73, 0x29, 0xb0, 0x2a, 0x85, 0x27, 0x72, 0xd9, 0xac, 0xbd,
	0x3f, 0xf3, 0xb6, 0xd0, 0xbf, 0x5d, 0xfd, 0xc2, 0x3c, 0xaf, 0x76, 0x67, 0x0a, 0xf0, 0x66, 0x8f,
	0xa6, 0x26, 0x0c, 0x32, 0x8d, 0x8d, 0xac, 0x55, 0xe2, 0xe3, 0xe6, 0x4b, 0xfd, 0xbe, 0x37, 0xa4,
	0x90, 0xe9, 0xb4, 0x35, 0xa4, 0x89, 0x9f, 0x1a, 0xc3, 0x87, 0xed, 0x9a, 0x2a, 0x21, 0x90, 0x69,
	0x42, 0x5a, 0x0b, 0x21, 0x89, 0x4f, 0x88, 0x31, 0x3d, 0xfd, 0x6f, 0xa6, 0x3b, 0xff, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0xbf, 0x11, 0x2d, 0xba, 0x12, 0x02, 0x00, 0x00,
}
