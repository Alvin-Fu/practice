// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/reach_plan_error.proto

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

// Enum describing possible errors from ReachPlanService.
type ReachPlanErrorEnum_ReachPlanError int32

const (
	// Enum unspecified.
	ReachPlanErrorEnum_UNSPECIFIED ReachPlanErrorEnum_ReachPlanError = 0
	// The received error code is not known in this version.
	ReachPlanErrorEnum_UNKNOWN ReachPlanErrorEnum_ReachPlanError = 1
)

var ReachPlanErrorEnum_ReachPlanError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
}

var ReachPlanErrorEnum_ReachPlanError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
}

func (x ReachPlanErrorEnum_ReachPlanError) String() string {
	return proto.EnumName(ReachPlanErrorEnum_ReachPlanError_name, int32(x))
}

func (ReachPlanErrorEnum_ReachPlanError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_965d67e16d0d10af, []int{0, 0}
}

// Container for enum describing possible errors returned from
// the ReachPlanService.
type ReachPlanErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReachPlanErrorEnum) Reset()         { *m = ReachPlanErrorEnum{} }
func (m *ReachPlanErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ReachPlanErrorEnum) ProtoMessage()    {}
func (*ReachPlanErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_965d67e16d0d10af, []int{0}
}

func (m *ReachPlanErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReachPlanErrorEnum.Unmarshal(m, b)
}
func (m *ReachPlanErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReachPlanErrorEnum.Marshal(b, m, deterministic)
}
func (m *ReachPlanErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReachPlanErrorEnum.Merge(m, src)
}
func (m *ReachPlanErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ReachPlanErrorEnum.Size(m)
}
func (m *ReachPlanErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ReachPlanErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ReachPlanErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.ReachPlanErrorEnum_ReachPlanError", ReachPlanErrorEnum_ReachPlanError_name, ReachPlanErrorEnum_ReachPlanError_value)
	proto.RegisterType((*ReachPlanErrorEnum)(nil), "google.ads.googleads.v2.errors.ReachPlanErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/reach_plan_error.proto", fileDescriptor_965d67e16d0d10af)
}

var fileDescriptor_965d67e16d0d10af = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4a, 0xc3, 0x40,
	0x18, 0x85, 0x4d, 0x04, 0x85, 0x29, 0x68, 0x89, 0x3b, 0x91, 0x2e, 0x72, 0x80, 0x7f, 0x20, 0xe2,
	0x66, 0x5c, 0xa5, 0x26, 0x96, 0x22, 0xc4, 0xa0, 0x34, 0x82, 0x04, 0xca, 0xd8, 0x84, 0x31, 0x90,
	0xce, 0x1f, 0x66, 0x62, 0x0f, 0xe4, 0xd2, 0xa3, 0x78, 0x14, 0x17, 0x9e, 0x41, 0x92, 0xdf, 0x04,
	0xba, 0xd0, 0xd5, 0x3c, 0x1e, 0xdf, 0x7b, 0xf3, 0xf8, 0xd9, 0x95, 0x42, 0x54, 0x75, 0xc9, 0x65,
	0x61, 0x39, 0xc9, 0x4e, 0xed, 0x02, 0x5e, 0x1a, 0x83, 0xc6, 0x72, 0x53, 0xca, 0xcd, 0xeb, 0xba,
	0xa9, 0xa5, 0x5e, 0xf7, 0x0e, 0x34, 0x06, 0x5b, 0xf4, 0x66, 0xc4, 0x82, 0x2c, 0x2c, 0x8c, 0x31,
	0xd8, 0x05, 0x40, 0xb1, 0xf3, 0x8b, 0xa1, 0xb6, 0xa9, 0xb8, 0xd4, 0x1a, 0x5b, 0xd9, 0x56, 0xa8,
	0x2d, 0xa5, 0xfd, 0x88, 0x79, 0x0f, 0x5d, 0x6f, 0x5a, 0x4b, 0x1d, 0x77, 0x81, 0x58, 0xbf, 0x6d,
	0x7d, 0x60, 0x27, 0xfb, 0xae, 0x77, 0xca, 0x26, 0xab, 0xe4, 0x31, 0x8d, 0x6f, 0x96, 0xb7, 0xcb,
	0x38, 0x9a, 0x1e, 0x78, 0x13, 0x76, 0xbc, 0x4a, 0xee, 0x92, 0xfb, 0xa7, 0x64, 0xea, 0xcc, 0xbf,
	0x1d, 0xe6, 0x6f, 0x70, 0x0b, 0xff, 0x4f, 0x99, 0x9f, 0xed, 0x97, 0xa6, 0xdd, 0x82, 0xd4, 0x79,
	0x8e, 0x7e, 0x63, 0x0a, 0x6b, 0xa9, 0x15, 0xa0, 0x51, 0x5c, 0x95, 0xba, 0xdf, 0x37, 0x1c, 0xa2,
	0xa9, 0xec, 0x5f, 0x77, 0xb9, 0xa6, 0xe7, 0xdd, 0x3d, 0x5c, 0x84, 0xe1, 0x87, 0x3b, 0x5b, 0x50,
	0x59, 0x58, 0x58, 0x20, 0xd9, 0xa9, 0x2c, 0x80, 0xfe, 0x4b, 0xfb, 0x39, 0x00, 0x79, 0x58, 0xd8,
	0x7c, 0x04, 0xf2, 0x2c, 0xc8, 0x09, 0xf8, 0x72, 0x7d, 0x72, 0x85, 0x08, 0x0b, 0x2b, 0xc4, 0x88,
	0x08, 0x91, 0x05, 0x42, 0x10, 0xf4, 0x72, 0xd4, 0xaf, 0xbb, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xee, 0x5a, 0x2e, 0x65, 0xb4, 0x01, 0x00, 0x00,
}
