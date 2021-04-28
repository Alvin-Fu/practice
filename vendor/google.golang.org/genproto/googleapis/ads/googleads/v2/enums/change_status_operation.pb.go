// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/change_status_operation.proto

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

// Status of the changed resource
type ChangeStatusOperationEnum_ChangeStatusOperation int32

const (
	// No value has been specified.
	ChangeStatusOperationEnum_UNSPECIFIED ChangeStatusOperationEnum_ChangeStatusOperation = 0
	// Used for return value only. Represents an unclassified resource unknown
	// in this version.
	ChangeStatusOperationEnum_UNKNOWN ChangeStatusOperationEnum_ChangeStatusOperation = 1
	// The resource was created.
	ChangeStatusOperationEnum_ADDED ChangeStatusOperationEnum_ChangeStatusOperation = 2
	// The resource was modified.
	ChangeStatusOperationEnum_CHANGED ChangeStatusOperationEnum_ChangeStatusOperation = 3
	// The resource was removed.
	ChangeStatusOperationEnum_REMOVED ChangeStatusOperationEnum_ChangeStatusOperation = 4
)

var ChangeStatusOperationEnum_ChangeStatusOperation_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ADDED",
	3: "CHANGED",
	4: "REMOVED",
}

var ChangeStatusOperationEnum_ChangeStatusOperation_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ADDED":       2,
	"CHANGED":     3,
	"REMOVED":     4,
}

func (x ChangeStatusOperationEnum_ChangeStatusOperation) String() string {
	return proto.EnumName(ChangeStatusOperationEnum_ChangeStatusOperation_name, int32(x))
}

func (ChangeStatusOperationEnum_ChangeStatusOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4e1869597aca0ca1, []int{0, 0}
}

// Container for enum describing operations for the ChangeStatus resource.
type ChangeStatusOperationEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeStatusOperationEnum) Reset()         { *m = ChangeStatusOperationEnum{} }
func (m *ChangeStatusOperationEnum) String() string { return proto.CompactTextString(m) }
func (*ChangeStatusOperationEnum) ProtoMessage()    {}
func (*ChangeStatusOperationEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e1869597aca0ca1, []int{0}
}

func (m *ChangeStatusOperationEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeStatusOperationEnum.Unmarshal(m, b)
}
func (m *ChangeStatusOperationEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeStatusOperationEnum.Marshal(b, m, deterministic)
}
func (m *ChangeStatusOperationEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeStatusOperationEnum.Merge(m, src)
}
func (m *ChangeStatusOperationEnum) XXX_Size() int {
	return xxx_messageInfo_ChangeStatusOperationEnum.Size(m)
}
func (m *ChangeStatusOperationEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeStatusOperationEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeStatusOperationEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.ChangeStatusOperationEnum_ChangeStatusOperation", ChangeStatusOperationEnum_ChangeStatusOperation_name, ChangeStatusOperationEnum_ChangeStatusOperation_value)
	proto.RegisterType((*ChangeStatusOperationEnum)(nil), "google.ads.googleads.v2.enums.ChangeStatusOperationEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/change_status_operation.proto", fileDescriptor_4e1869597aca0ca1)
}

var fileDescriptor_4e1869597aca0ca1 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4f, 0xfa, 0x30,
	0x18, 0xc6, 0xff, 0x8c, 0xbf, 0x1a, 0xcb, 0xc1, 0x65, 0x89, 0x07, 0x89, 0x1c, 0xe0, 0x03, 0xb4,
	0xc9, 0xbc, 0x95, 0x53, 0xa1, 0x15, 0x89, 0x71, 0x10, 0x09, 0x33, 0x21, 0x4b, 0x48, 0x65, 0x4b,
	0x25, 0x81, 0x76, 0x59, 0x37, 0xfc, 0x3e, 0x1e, 0xfd, 0x28, 0x7e, 0x14, 0x2f, 0x7e, 0x05, 0xd3,
	0xd6, 0xed, 0x84, 0x5e, 0x9a, 0xa7, 0xef, 0xf3, 0xfe, 0x9e, 0xbe, 0x7d, 0xc1, 0x50, 0x28, 0x25,
	0x76, 0x19, 0xe2, 0xa9, 0x46, 0x4e, 0x1a, 0x75, 0x08, 0x51, 0x26, 0xab, 0xbd, 0x46, 0x9b, 0x17,
	0x2e, 0x45, 0xb6, 0xd6, 0x25, 0x2f, 0x2b, 0xbd, 0x56, 0x79, 0x56, 0xf0, 0x72, 0xab, 0x24, 0xcc,
	0x0b, 0x55, 0xaa, 0xa0, 0xe7, 0x08, 0xc8, 0x53, 0x0d, 0x1b, 0x18, 0x1e, 0x42, 0x68, 0xe1, 0xee,
	0x75, 0x9d, 0x9d, 0x6f, 0x11, 0x97, 0x52, 0x95, 0x96, 0xd5, 0x0e, 0x1e, 0xbc, 0x82, 0xab, 0xb1,
	0x4d, 0x5f, 0xd8, 0xf0, 0x59, 0x9d, 0xcd, 0x64, 0xb5, 0x1f, 0xac, 0xc0, 0xe5, 0x51, 0x33, 0xb8,
	0x00, 0x9d, 0x65, 0xb4, 0x98, 0xb3, 0xf1, 0xf4, 0x76, 0xca, 0xa8, 0xff, 0x2f, 0xe8, 0x80, 0xb3,
	0x65, 0x74, 0x1f, 0xcd, 0x9e, 0x22, 0xbf, 0x15, 0x9c, 0x83, 0x13, 0x42, 0x29, 0xa3, 0xbe, 0x67,
	0xea, 0xe3, 0x3b, 0x12, 0x4d, 0x18, 0xf5, 0xdb, 0xe6, 0xf2, 0xc8, 0x1e, 0x66, 0x31, 0xa3, 0xfe,
	0xff, 0xd1, 0x57, 0x0b, 0xf4, 0x37, 0x6a, 0x0f, 0xff, 0x1c, 0x7e, 0xd4, 0x3d, 0xfa, 0xfe, 0xdc,
	0x8c, 0x3e, 0x6f, 0xad, 0x46, 0x3f, 0xb0, 0x50, 0x3b, 0x2e, 0x05, 0x54, 0x85, 0x40, 0x22, 0x93,
	0xf6, 0x63, 0xf5, 0x1a, 0xf3, 0xad, 0xfe, 0x65, 0xab, 0x43, 0x7b, 0xbe, 0x79, 0xed, 0x09, 0x21,
	0xef, 0x5e, 0x6f, 0xe2, 0xa2, 0x48, 0xaa, 0xa1, 0x93, 0x46, 0xc5, 0x21, 0x34, 0x8b, 0xd0, 0x1f,
	0xb5, 0x9f, 0x90, 0x54, 0x27, 0x8d, 0x9f, 0xc4, 0x61, 0x62, 0xfd, 0x4f, 0xaf, 0xef, 0x8a, 0x18,
	0x93, 0x54, 0x63, 0xdc, 0x74, 0x60, 0x1c, 0x87, 0x18, 0xdb, 0x9e, 0xe7, 0x53, 0x3b, 0xd8, 0xcd,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0xf0, 0xb9, 0x3c, 0xed, 0x01, 0x00, 0x00,
}
