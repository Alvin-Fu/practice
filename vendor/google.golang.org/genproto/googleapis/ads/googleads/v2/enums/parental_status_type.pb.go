// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/parental_status_type.proto

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

// The type of parental statuses (e.g. not a parent).
type ParentalStatusTypeEnum_ParentalStatusType int32

const (
	// Not specified.
	ParentalStatusTypeEnum_UNSPECIFIED ParentalStatusTypeEnum_ParentalStatusType = 0
	// Used for return value only. Represents value unknown in this version.
	ParentalStatusTypeEnum_UNKNOWN ParentalStatusTypeEnum_ParentalStatusType = 1
	// Parent.
	ParentalStatusTypeEnum_PARENT ParentalStatusTypeEnum_ParentalStatusType = 300
	// Not a parent.
	ParentalStatusTypeEnum_NOT_A_PARENT ParentalStatusTypeEnum_ParentalStatusType = 301
	// Undetermined parental status.
	ParentalStatusTypeEnum_UNDETERMINED ParentalStatusTypeEnum_ParentalStatusType = 302
)

var ParentalStatusTypeEnum_ParentalStatusType_name = map[int32]string{
	0:   "UNSPECIFIED",
	1:   "UNKNOWN",
	300: "PARENT",
	301: "NOT_A_PARENT",
	302: "UNDETERMINED",
}

var ParentalStatusTypeEnum_ParentalStatusType_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"PARENT":       300,
	"NOT_A_PARENT": 301,
	"UNDETERMINED": 302,
}

func (x ParentalStatusTypeEnum_ParentalStatusType) String() string {
	return proto.EnumName(ParentalStatusTypeEnum_ParentalStatusType_name, int32(x))
}

func (ParentalStatusTypeEnum_ParentalStatusType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_781178e175795ad1, []int{0, 0}
}

// Container for enum describing the type of demographic parental statuses.
type ParentalStatusTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParentalStatusTypeEnum) Reset()         { *m = ParentalStatusTypeEnum{} }
func (m *ParentalStatusTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ParentalStatusTypeEnum) ProtoMessage()    {}
func (*ParentalStatusTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_781178e175795ad1, []int{0}
}

func (m *ParentalStatusTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParentalStatusTypeEnum.Unmarshal(m, b)
}
func (m *ParentalStatusTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParentalStatusTypeEnum.Marshal(b, m, deterministic)
}
func (m *ParentalStatusTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParentalStatusTypeEnum.Merge(m, src)
}
func (m *ParentalStatusTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ParentalStatusTypeEnum.Size(m)
}
func (m *ParentalStatusTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ParentalStatusTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ParentalStatusTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.ParentalStatusTypeEnum_ParentalStatusType", ParentalStatusTypeEnum_ParentalStatusType_name, ParentalStatusTypeEnum_ParentalStatusType_value)
	proto.RegisterType((*ParentalStatusTypeEnum)(nil), "google.ads.googleads.v2.enums.ParentalStatusTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/parental_status_type.proto", fileDescriptor_781178e175795ad1)
}

var fileDescriptor_781178e175795ad1 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4f, 0x4a, 0x33, 0x31,
	0x1c, 0xfd, 0x3a, 0x1f, 0x54, 0x48, 0x05, 0xc7, 0x59, 0x28, 0x88, 0x5d, 0xb4, 0x07, 0xc8, 0xc0,
	0xb8, 0x91, 0xb8, 0x4a, 0x6d, 0x2c, 0x45, 0x4c, 0x87, 0x76, 0x5a, 0x41, 0x06, 0x86, 0xe8, 0x84,
	0x50, 0x68, 0x93, 0xd0, 0xa4, 0x85, 0xae, 0xbc, 0x8b, 0x4b, 0x51, 0x0f, 0xe2, 0x51, 0xc4, 0x43,
	0xc8, 0x24, 0x6d, 0x37, 0x45, 0x37, 0xe1, 0xf1, 0x7b, 0x7f, 0x78, 0x79, 0xe0, 0x52, 0x28, 0x25,
	0x66, 0x3c, 0x66, 0xa5, 0x89, 0x3d, 0xac, 0xd0, 0x2a, 0x89, 0xb9, 0x5c, 0xce, 0x4d, 0xac, 0xd9,
	0x82, 0x4b, 0xcb, 0x66, 0x85, 0xb1, 0xcc, 0x2e, 0x4d, 0x61, 0xd7, 0x9a, 0x43, 0xbd, 0x50, 0x56,
	0x45, 0x4d, 0x2f, 0x87, 0xac, 0x34, 0x70, 0xe7, 0x84, 0xab, 0x04, 0x3a, 0xe7, 0xd9, 0xf9, 0x36,
	0x58, 0x4f, 0x63, 0x26, 0xa5, 0xb2, 0xcc, 0x4e, 0x95, 0x34, 0xde, 0xdc, 0x7e, 0x06, 0x27, 0xe9,
	0x26, 0x7a, 0xe4, 0x92, 0xb3, 0xb5, 0xe6, 0x44, 0x2e, 0xe7, 0x6d, 0x0e, 0xa2, 0x7d, 0x26, 0x3a,
	0x02, 0x8d, 0x31, 0x1d, 0xa5, 0xe4, 0xba, 0x7f, 0xd3, 0x27, 0xdd, 0xf0, 0x5f, 0xd4, 0x00, 0x07,
	0x63, 0x7a, 0x4b, 0x07, 0xf7, 0x34, 0xac, 0x45, 0x0d, 0x50, 0x4f, 0xf1, 0x90, 0xd0, 0x2c, 0x7c,
	0x0b, 0xa2, 0x63, 0x70, 0x48, 0x07, 0x59, 0x81, 0x8b, 0xcd, 0xe9, 0xdd, 0x9d, 0xc6, 0xb4, 0x4b,
	0x32, 0x32, 0xbc, 0xeb, 0x53, 0xd2, 0x0d, 0x3f, 0x82, 0xce, 0x77, 0x0d, 0xb4, 0x9e, 0xd4, 0x1c,
	0xfe, 0xf9, 0x89, 0xce, 0xe9, 0x7e, 0x95, 0xb4, 0xea, 0x9f, 0xd6, 0x1e, 0x3a, 0x1b, 0xa7, 0x50,
	0x33, 0x26, 0x05, 0x54, 0x0b, 0x11, 0x0b, 0x2e, 0xdd, 0xef, 0xb6, 0x43, 0xea, 0xa9, 0xf9, 0x65,
	0xd7, 0x2b, 0xf7, 0xbe, 0x04, 0xff, 0x7b, 0x18, 0xbf, 0x06, 0xcd, 0x9e, 0x8f, 0xc2, 0xa5, 0x81,
	0x1e, 0x56, 0x68, 0x92, 0xc0, 0x6a, 0x10, 0xf3, 0xb9, 0xe5, 0x73, 0x5c, 0x9a, 0x7c, 0xc7, 0xe7,
	0x93, 0x24, 0x77, 0xfc, 0x57, 0xd0, 0xf2, 0x47, 0x84, 0x70, 0x69, 0x10, 0xda, 0x29, 0x10, 0x9a,
	0x24, 0x08, 0x39, 0xcd, 0x63, 0xdd, 0x15, 0xbb, 0xf8, 0x09, 0x00, 0x00, 0xff, 0xff, 0xff, 0xfd,
	0xc8, 0xc4, 0xef, 0x01, 0x00, 0x00,
}
