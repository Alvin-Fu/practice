// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/mutate_error.proto

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

// Enum describing possible mutate errors.
type MutateErrorEnum_MutateError int32

const (
	// Enum unspecified.
	MutateErrorEnum_UNSPECIFIED MutateErrorEnum_MutateError = 0
	// The received error code is not known in this version.
	MutateErrorEnum_UNKNOWN MutateErrorEnum_MutateError = 1
	// Requested resource was not found.
	MutateErrorEnum_RESOURCE_NOT_FOUND MutateErrorEnum_MutateError = 3
	// Cannot mutate the same resource twice in one request.
	MutateErrorEnum_ID_EXISTS_IN_MULTIPLE_MUTATES MutateErrorEnum_MutateError = 7
	// The field's contents don't match another field that represents the same
	// data.
	MutateErrorEnum_INCONSISTENT_FIELD_VALUES MutateErrorEnum_MutateError = 8
	// Mutates are not allowed for the requested resource.
	MutateErrorEnum_MUTATE_NOT_ALLOWED MutateErrorEnum_MutateError = 9
	// The resource isn't in Google Ads. It belongs to another ads system.
	MutateErrorEnum_RESOURCE_NOT_IN_GOOGLE_ADS MutateErrorEnum_MutateError = 10
	// The resource being created already exists.
	MutateErrorEnum_RESOURCE_ALREADY_EXISTS MutateErrorEnum_MutateError = 11
)

var MutateErrorEnum_MutateError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	3:  "RESOURCE_NOT_FOUND",
	7:  "ID_EXISTS_IN_MULTIPLE_MUTATES",
	8:  "INCONSISTENT_FIELD_VALUES",
	9:  "MUTATE_NOT_ALLOWED",
	10: "RESOURCE_NOT_IN_GOOGLE_ADS",
	11: "RESOURCE_ALREADY_EXISTS",
}

var MutateErrorEnum_MutateError_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"RESOURCE_NOT_FOUND":            3,
	"ID_EXISTS_IN_MULTIPLE_MUTATES": 7,
	"INCONSISTENT_FIELD_VALUES":     8,
	"MUTATE_NOT_ALLOWED":            9,
	"RESOURCE_NOT_IN_GOOGLE_ADS":    10,
	"RESOURCE_ALREADY_EXISTS":       11,
}

func (x MutateErrorEnum_MutateError) String() string {
	return proto.EnumName(MutateErrorEnum_MutateError_name, int32(x))
}

func (MutateErrorEnum_MutateError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_69e9950c75c98711, []int{0, 0}
}

// Container for enum describing possible mutate errors.
type MutateErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateErrorEnum) Reset()         { *m = MutateErrorEnum{} }
func (m *MutateErrorEnum) String() string { return proto.CompactTextString(m) }
func (*MutateErrorEnum) ProtoMessage()    {}
func (*MutateErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_69e9950c75c98711, []int{0}
}

func (m *MutateErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateErrorEnum.Unmarshal(m, b)
}
func (m *MutateErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateErrorEnum.Marshal(b, m, deterministic)
}
func (m *MutateErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateErrorEnum.Merge(m, src)
}
func (m *MutateErrorEnum) XXX_Size() int {
	return xxx_messageInfo_MutateErrorEnum.Size(m)
}
func (m *MutateErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MutateErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.MutateErrorEnum_MutateError", MutateErrorEnum_MutateError_name, MutateErrorEnum_MutateError_value)
	proto.RegisterType((*MutateErrorEnum)(nil), "google.ads.googleads.v1.errors.MutateErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/mutate_error.proto", fileDescriptor_69e9950c75c98711)
}

var fileDescriptor_69e9950c75c98711 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x8e, 0xd3, 0x30,
	0x18, 0x85, 0x49, 0x47, 0x62, 0xc0, 0x5d, 0x8c, 0xe5, 0x05, 0x88, 0x81, 0xa9, 0x44, 0x0f, 0xe0,
	0x28, 0x62, 0x67, 0x56, 0x9e, 0xda, 0xad, 0x2c, 0x5c, 0xbb, 0xaa, 0x93, 0x0e, 0xa0, 0x4a, 0x56,
	0x20, 0x51, 0x54, 0x69, 0x1a, 0x57, 0x71, 0x66, 0x0e, 0xc4, 0x92, 0xa3, 0x70, 0x04, 0x8e, 0x00,
	0x12, 0x67, 0x40, 0x89, 0xa7, 0x51, 0x59, 0x30, 0xab, 0xbc, 0xd8, 0xdf, 0x7b, 0xfe, 0xf5, 0x7e,
	0x90, 0x54, 0xce, 0x55, 0xb7, 0x65, 0x9c, 0x17, 0x3e, 0x0e, 0xb2, 0x53, 0xf7, 0x49, 0x5c, 0x36,
	0x8d, 0x6b, 0x7c, 0xbc, 0xbf, 0x6b, 0xf3, 0xb6, 0xb4, 0xfd, 0x1f, 0x3e, 0x34, 0xae, 0x75, 0x68,
	0x12, 0x38, 0x9c, 0x17, 0x1e, 0x0f, 0x16, 0x7c, 0x9f, 0xe0, 0x60, 0xb9, 0x7c, 0x73, 0x8c, 0x3c,
	0xec, 0xe2, 0xbc, 0xae, 0x5d, 0x9b, 0xb7, 0x3b, 0x57, 0xfb, 0xe0, 0x9e, 0xfe, 0x89, 0xc0, 0xc5,
	0xb2, 0x0f, 0xe5, 0x1d, 0xce, 0xeb, 0xbb, 0xfd, 0xf4, 0x67, 0x04, 0xc6, 0x27, 0x67, 0xe8, 0x02,
	0x8c, 0x33, 0x65, 0x56, 0x7c, 0x26, 0xe6, 0x82, 0x33, 0xf8, 0x04, 0x8d, 0xc1, 0x79, 0xa6, 0x3e,
	0x28, 0x7d, 0xa3, 0x60, 0x84, 0x5e, 0x00, 0xb4, 0xe6, 0x46, 0x67, 0xeb, 0x19, 0xb7, 0x4a, 0xa7,
	0x76, 0xae, 0x33, 0xc5, 0xe0, 0x19, 0x7a, 0x0b, 0xae, 0x04, 0xb3, 0xfc, 0xa3, 0x30, 0xa9, 0xb1,
	0x42, 0xd9, 0x65, 0x26, 0x53, 0xb1, 0x92, 0xdc, 0x2e, 0xb3, 0x94, 0xa6, 0xdc, 0xc0, 0x73, 0x74,
	0x05, 0x5e, 0x09, 0x35, 0xd3, 0xca, 0x08, 0x93, 0x72, 0x95, 0xda, 0xb9, 0xe0, 0x92, 0xd9, 0x0d,
	0x95, 0x19, 0x37, 0xf0, 0x59, 0x97, 0x1c, 0xd8, 0x3e, 0x97, 0x4a, 0xa9, 0x6f, 0x38, 0x83, 0xcf,
	0xd1, 0x04, 0x5c, 0xfe, 0xf3, 0xa2, 0x50, 0x76, 0xa1, 0xf5, 0x42, 0x72, 0x4b, 0x99, 0x81, 0x00,
	0xbd, 0x06, 0x2f, 0x87, 0x7b, 0x2a, 0xd7, 0x9c, 0xb2, 0x4f, 0x0f, 0x73, 0xc0, 0xf1, 0xf5, 0xef,
	0x08, 0x4c, 0xbf, 0xba, 0x3d, 0x7e, 0xbc, 0xb5, 0x6b, 0x78, 0x52, 0xc0, 0xaa, 0x6b, 0x6a, 0x15,
	0x7d, 0x66, 0x0f, 0x9e, 0xca, 0xdd, 0xe6, 0x75, 0x85, 0x5d, 0x53, 0xc5, 0x55, 0x59, 0xf7, 0x3d,
	0x1e, 0x97, 0x75, 0xd8, 0xf9, 0xff, 0xed, 0xee, 0x7d, 0xf8, 0x7c, 0x1b, 0x9d, 0x2d, 0x28, 0xfd,
	0x3e, 0x9a, 0x2c, 0x42, 0x18, 0x2d, 0x3c, 0x0e, 0xb2, 0x53, 0x9b, 0x04, 0xf7, 0x4f, 0xfa, 0x1f,
	0x47, 0x60, 0x4b, 0x0b, 0xbf, 0x1d, 0x80, 0xed, 0x26, 0xd9, 0x06, 0xe0, 0xd7, 0x68, 0x1a, 0x4e,
	0x09, 0xa1, 0x85, 0x27, 0x64, 0x40, 0x08, 0xd9, 0x24, 0x84, 0x04, 0xe8, 0xcb, 0xd3, 0x7e, 0xba,
	0x77, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xde, 0x4c, 0xb9, 0x2c, 0x58, 0x02, 0x00, 0x00,
}
