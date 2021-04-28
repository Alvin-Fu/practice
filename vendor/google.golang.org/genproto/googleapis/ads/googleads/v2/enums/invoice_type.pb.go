// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/invoice_type.proto

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

// The possible type of invoices.
type InvoiceTypeEnum_InvoiceType int32

const (
	// Not specified.
	InvoiceTypeEnum_UNSPECIFIED InvoiceTypeEnum_InvoiceType = 0
	// Used for return value only. Represents value unknown in this version.
	InvoiceTypeEnum_UNKNOWN InvoiceTypeEnum_InvoiceType = 1
	// An invoice with a negative amount. The account receives a credit.
	InvoiceTypeEnum_CREDIT_MEMO InvoiceTypeEnum_InvoiceType = 2
	// An invoice with a positive amount. The account owes a balance.
	InvoiceTypeEnum_INVOICE InvoiceTypeEnum_InvoiceType = 3
)

var InvoiceTypeEnum_InvoiceType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CREDIT_MEMO",
	3: "INVOICE",
}

var InvoiceTypeEnum_InvoiceType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"CREDIT_MEMO": 2,
	"INVOICE":     3,
}

func (x InvoiceTypeEnum_InvoiceType) String() string {
	return proto.EnumName(InvoiceTypeEnum_InvoiceType_name, int32(x))
}

func (InvoiceTypeEnum_InvoiceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9b9c6249f36ed5a4, []int{0, 0}
}

// Container for enum describing the type of invoices.
type InvoiceTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvoiceTypeEnum) Reset()         { *m = InvoiceTypeEnum{} }
func (m *InvoiceTypeEnum) String() string { return proto.CompactTextString(m) }
func (*InvoiceTypeEnum) ProtoMessage()    {}
func (*InvoiceTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9c6249f36ed5a4, []int{0}
}

func (m *InvoiceTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvoiceTypeEnum.Unmarshal(m, b)
}
func (m *InvoiceTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvoiceTypeEnum.Marshal(b, m, deterministic)
}
func (m *InvoiceTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvoiceTypeEnum.Merge(m, src)
}
func (m *InvoiceTypeEnum) XXX_Size() int {
	return xxx_messageInfo_InvoiceTypeEnum.Size(m)
}
func (m *InvoiceTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_InvoiceTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_InvoiceTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.InvoiceTypeEnum_InvoiceType", InvoiceTypeEnum_InvoiceType_name, InvoiceTypeEnum_InvoiceType_value)
	proto.RegisterType((*InvoiceTypeEnum)(nil), "google.ads.googleads.v2.enums.InvoiceTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/invoice_type.proto", fileDescriptor_9b9c6249f36ed5a4)
}

var fileDescriptor_9b9c6249f36ed5a4 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x18, 0x77, 0x1d, 0x28, 0x64, 0x87, 0x95, 0x1d, 0xc5, 0x1d, 0xb6, 0x07, 0x48, 0xa4, 0xde, 0xe2,
	0xa9, 0xed, 0xe2, 0x08, 0xb2, 0xb6, 0xe8, 0x56, 0x41, 0x0a, 0xa3, 0xae, 0x21, 0x14, 0xd6, 0xa4,
	0x2c, 0x5d, 0x61, 0xaf, 0xe3, 0xd1, 0x47, 0xf1, 0x3d, 0xbc, 0xf8, 0x14, 0x92, 0xc4, 0x96, 0x5d,
	0xf4, 0x12, 0x7e, 0x7c, 0xbf, 0x3f, 0xf9, 0x7d, 0x1f, 0xb8, 0xe5, 0x52, 0xf2, 0x3d, 0x43, 0x79,
	0xa1, 0x90, 0x85, 0x1a, 0xb5, 0x1e, 0x62, 0xe2, 0x58, 0x29, 0x54, 0x8a, 0x56, 0x96, 0x3b, 0xb6,
	0x6d, 0x4e, 0x35, 0x83, 0xf5, 0x41, 0x36, 0x72, 0x32, 0xb5, 0x32, 0x98, 0x17, 0x0a, 0xf6, 0x0e,
	0xd8, 0x7a, 0xd0, 0x38, 0xae, 0x6f, 0xba, 0xc0, 0xba, 0x44, 0xb9, 0x10, 0xb2, 0xc9, 0x9b, 0x52,
	0x0a, 0x65, 0xcd, 0xf3, 0x0c, 0x8c, 0xa9, 0x8d, 0x5c, 0x9f, 0x6a, 0x46, 0xc4, 0xb1, 0x9a, 0x53,
	0x30, 0x3a, 0x1b, 0x4d, 0xc6, 0x60, 0xb4, 0x89, 0x9e, 0x13, 0x12, 0xd2, 0x07, 0x4a, 0x16, 0xee,
	0xc5, 0x64, 0x04, 0xae, 0x36, 0xd1, 0x63, 0x14, 0xbf, 0x44, 0xee, 0x40, 0xb3, 0xe1, 0x13, 0x59,
	0xd0, 0xf5, 0x76, 0x45, 0x56, 0xb1, 0xeb, 0x68, 0x96, 0x46, 0x69, 0x4c, 0x43, 0xe2, 0x0e, 0x83,
	0xaf, 0x01, 0x98, 0xed, 0x64, 0x05, 0xff, 0x6d, 0x18, 0xb8, 0x67, 0xdf, 0x25, 0xba, 0x55, 0x32,
	0x78, 0x0d, 0x7e, 0x2d, 0x5c, 0xee, 0x73, 0xc1, 0xa1, 0x3c, 0x70, 0xc4, 0x99, 0x30, 0x9d, 0xbb,
	0xb3, 0xd4, 0xa5, 0xfa, 0xe3, 0x4a, 0xf7, 0xe6, 0x7d, 0x77, 0x86, 0x4b, 0xdf, 0xff, 0x70, 0xa6,
	0x4b, 0x1b, 0xe5, 0x17, 0x0a, 0x5a, 0xa8, 0x51, 0xea, 0x41, 0xbd, 0xad, 0xfa, 0xec, 0xf8, 0xcc,
	0x2f, 0x54, 0xd6, 0xf3, 0x59, 0xea, 0x65, 0x86, 0xff, 0x76, 0x66, 0x76, 0x88, 0xb1, 0x5f, 0x28,
	0x8c, 0x7b, 0x05, 0xc6, 0xa9, 0x87, 0xb1, 0xd1, 0xbc, 0x5d, 0x9a, 0x62, 0x77, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xd9, 0x24, 0x76, 0xb3, 0xbd, 0x01, 0x00, 0x00,
}
