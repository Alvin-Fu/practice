// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/income_range_type.proto

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

// The type of demographic income ranges (e.g. between 0% to 50%).
type IncomeRangeTypeEnum_IncomeRangeType int32

const (
	// Not specified.
	IncomeRangeTypeEnum_UNSPECIFIED IncomeRangeTypeEnum_IncomeRangeType = 0
	// Used for return value only. Represents value unknown in this version.
	IncomeRangeTypeEnum_UNKNOWN IncomeRangeTypeEnum_IncomeRangeType = 1
	// 0%-50%.
	IncomeRangeTypeEnum_INCOME_RANGE_0_50 IncomeRangeTypeEnum_IncomeRangeType = 510001
	// 50% to 60%.
	IncomeRangeTypeEnum_INCOME_RANGE_50_60 IncomeRangeTypeEnum_IncomeRangeType = 510002
	// 60% to 70%.
	IncomeRangeTypeEnum_INCOME_RANGE_60_70 IncomeRangeTypeEnum_IncomeRangeType = 510003
	// 70% to 80%.
	IncomeRangeTypeEnum_INCOME_RANGE_70_80 IncomeRangeTypeEnum_IncomeRangeType = 510004
	// 80% to 90%.
	IncomeRangeTypeEnum_INCOME_RANGE_80_90 IncomeRangeTypeEnum_IncomeRangeType = 510005
	// Greater than 90%.
	IncomeRangeTypeEnum_INCOME_RANGE_90_UP IncomeRangeTypeEnum_IncomeRangeType = 510006
	// Undetermined income range.
	IncomeRangeTypeEnum_INCOME_RANGE_UNDETERMINED IncomeRangeTypeEnum_IncomeRangeType = 510000
)

var IncomeRangeTypeEnum_IncomeRangeType_name = map[int32]string{
	0:      "UNSPECIFIED",
	1:      "UNKNOWN",
	510001: "INCOME_RANGE_0_50",
	510002: "INCOME_RANGE_50_60",
	510003: "INCOME_RANGE_60_70",
	510004: "INCOME_RANGE_70_80",
	510005: "INCOME_RANGE_80_90",
	510006: "INCOME_RANGE_90_UP",
	510000: "INCOME_RANGE_UNDETERMINED",
}

var IncomeRangeTypeEnum_IncomeRangeType_value = map[string]int32{
	"UNSPECIFIED":               0,
	"UNKNOWN":                   1,
	"INCOME_RANGE_0_50":         510001,
	"INCOME_RANGE_50_60":        510002,
	"INCOME_RANGE_60_70":        510003,
	"INCOME_RANGE_70_80":        510004,
	"INCOME_RANGE_80_90":        510005,
	"INCOME_RANGE_90_UP":        510006,
	"INCOME_RANGE_UNDETERMINED": 510000,
}

func (x IncomeRangeTypeEnum_IncomeRangeType) String() string {
	return proto.EnumName(IncomeRangeTypeEnum_IncomeRangeType_name, int32(x))
}

func (IncomeRangeTypeEnum_IncomeRangeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_170475cd2f3d5831, []int{0, 0}
}

// Container for enum describing the type of demographic income ranges.
type IncomeRangeTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncomeRangeTypeEnum) Reset()         { *m = IncomeRangeTypeEnum{} }
func (m *IncomeRangeTypeEnum) String() string { return proto.CompactTextString(m) }
func (*IncomeRangeTypeEnum) ProtoMessage()    {}
func (*IncomeRangeTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_170475cd2f3d5831, []int{0}
}

func (m *IncomeRangeTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncomeRangeTypeEnum.Unmarshal(m, b)
}
func (m *IncomeRangeTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncomeRangeTypeEnum.Marshal(b, m, deterministic)
}
func (m *IncomeRangeTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncomeRangeTypeEnum.Merge(m, src)
}
func (m *IncomeRangeTypeEnum) XXX_Size() int {
	return xxx_messageInfo_IncomeRangeTypeEnum.Size(m)
}
func (m *IncomeRangeTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_IncomeRangeTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_IncomeRangeTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.IncomeRangeTypeEnum_IncomeRangeType", IncomeRangeTypeEnum_IncomeRangeType_name, IncomeRangeTypeEnum_IncomeRangeType_value)
	proto.RegisterType((*IncomeRangeTypeEnum)(nil), "google.ads.googleads.v2.enums.IncomeRangeTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/income_range_type.proto", fileDescriptor_170475cd2f3d5831)
}

var fileDescriptor_170475cd2f3d5831 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x5d, 0xca, 0xd3, 0x40,
	0x18, 0x85, 0x4d, 0x04, 0x85, 0xf9, 0x2e, 0xbe, 0x31, 0x0a, 0xfe, 0xe0, 0x47, 0x68, 0x17, 0x30,
	0x19, 0x22, 0xfd, 0x1b, 0xaf, 0xd2, 0x76, 0x2c, 0x41, 0x3a, 0x0d, 0xb5, 0xa9, 0x20, 0x81, 0x21,
	0x36, 0x21, 0x04, 0x9a, 0x99, 0xd0, 0x49, 0x0b, 0xbd, 0x76, 0x03, 0x5d, 0x83, 0x97, 0x5e, 0xf9,
	0xbf, 0x08, 0x77, 0xa2, 0x6e, 0x42, 0x92, 0xd8, 0x42, 0x4b, 0xf4, 0x66, 0x38, 0xbc, 0xcf, 0x7b,
	0x0e, 0xc3, 0x79, 0x41, 0x27, 0x91, 0x32, 0x59, 0xc7, 0x56, 0x18, 0x29, 0xab, 0x96, 0xa5, 0xda,
	0xd9, 0x56, 0x2c, 0xb6, 0x99, 0xb2, 0x52, 0xb1, 0x92, 0x59, 0xcc, 0x37, 0xa1, 0x48, 0x62, 0x5e,
	0xec, 0xf3, 0x18, 0xe5, 0x1b, 0x59, 0x48, 0xe3, 0xa6, 0xde, 0x45, 0x61, 0xa4, 0xd0, 0xc9, 0x86,
	0x76, 0x36, 0xaa, 0x6c, 0x4f, 0x9e, 0x1e, 0x53, 0xf3, 0xd4, 0x0a, 0x85, 0x90, 0x45, 0x58, 0xa4,
	0x52, 0xa8, 0xda, 0xdc, 0x7e, 0xa7, 0x83, 0xfb, 0x6e, 0x15, 0x3c, 0x2f, 0x73, 0x17, 0xfb, 0x3c,
	0xa6, 0x62, 0x9b, 0xb5, 0x7f, 0x6b, 0xe0, 0xfa, 0x62, 0x6e, 0x5c, 0x83, 0x2b, 0x9f, 0xbd, 0xf2,
	0xe8, 0xc8, 0x7d, 0xe1, 0xd2, 0x31, 0xbc, 0x65, 0x5c, 0x81, 0xbb, 0x3e, 0x7b, 0xc9, 0x66, 0xaf,
	0x19, 0xd4, 0x8c, 0x87, 0xe0, 0x9e, 0xcb, 0x46, 0xb3, 0x29, 0xe5, 0x73, 0x87, 0x4d, 0x28, 0xc7,
	0xbc, 0x83, 0xe1, 0xa7, 0x83, 0x69, 0x3c, 0x02, 0xc6, 0x19, 0xe8, 0x60, 0xde, 0xc5, 0xf0, 0x73,
	0x03, 0xe9, 0x62, 0xde, 0xc3, 0xf0, 0x4b, 0x03, 0xe9, 0x61, 0xde, 0xc7, 0xf0, 0x6b, 0x03, 0xe9,
	0x63, 0x3e, 0xc0, 0xf0, 0x5b, 0x03, 0x19, 0x60, 0xee, 0x7b, 0xf0, 0xfb, 0xc1, 0x34, 0x4c, 0xf0,
	0xf8, 0x8c, 0xf8, 0x6c, 0x4c, 0x17, 0x74, 0x3e, 0x75, 0x19, 0x1d, 0xc3, 0x8f, 0x07, 0x73, 0xf8,
	0x53, 0x03, 0xad, 0x95, 0xcc, 0xd0, 0x7f, 0x9b, 0x1c, 0x3e, 0xb8, 0x28, 0xc4, 0x2b, 0x1b, 0xf4,
	0xb4, 0x37, 0xc3, 0xbf, 0xb6, 0x44, 0xae, 0x43, 0x91, 0x20, 0xb9, 0x49, 0xac, 0x24, 0x16, 0x55,
	0xbf, 0xc7, 0x3b, 0xe6, 0xa9, 0xfa, 0xc7, 0x59, 0x9f, 0x57, 0xef, 0x7b, 0xfd, 0xf6, 0xc4, 0x71,
	0x3e, 0xe8, 0x37, 0x93, 0x3a, 0xca, 0x89, 0x14, 0xaa, 0x65, 0xa9, 0x96, 0x36, 0x2a, 0x8f, 0xa2,
	0x7e, 0x1c, 0x79, 0xe0, 0x44, 0x2a, 0x38, 0xf1, 0x60, 0x69, 0x07, 0x15, 0xff, 0xa5, 0xb7, 0xea,
	0x21, 0x21, 0x4e, 0xa4, 0x08, 0x39, 0x6d, 0x10, 0xb2, 0xb4, 0x09, 0xa9, 0x76, 0xde, 0xde, 0xa9,
	0x3e, 0xf6, 0xec, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0x2e, 0x7c, 0xf4, 0x6e, 0x02, 0x00,
	0x00,
}
