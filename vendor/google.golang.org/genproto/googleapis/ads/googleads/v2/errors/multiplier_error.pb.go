// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/multiplier_error.proto

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

// Enum describing possible multiplier errors.
type MultiplierErrorEnum_MultiplierError int32

const (
	// Enum unspecified.
	MultiplierErrorEnum_UNSPECIFIED MultiplierErrorEnum_MultiplierError = 0
	// The received error code is not known in this version.
	MultiplierErrorEnum_UNKNOWN MultiplierErrorEnum_MultiplierError = 1
	// Multiplier value is too high
	MultiplierErrorEnum_MULTIPLIER_TOO_HIGH MultiplierErrorEnum_MultiplierError = 2
	// Multiplier value is too low
	MultiplierErrorEnum_MULTIPLIER_TOO_LOW MultiplierErrorEnum_MultiplierError = 3
	// Too many fractional digits
	MultiplierErrorEnum_TOO_MANY_FRACTIONAL_DIGITS MultiplierErrorEnum_MultiplierError = 4
	// A multiplier cannot be set for this bidding strategy
	MultiplierErrorEnum_MULTIPLIER_NOT_ALLOWED_FOR_BIDDING_STRATEGY MultiplierErrorEnum_MultiplierError = 5
	// A multiplier cannot be set when there is no base bid (e.g., content max
	// cpc)
	MultiplierErrorEnum_MULTIPLIER_NOT_ALLOWED_WHEN_BASE_BID_IS_MISSING MultiplierErrorEnum_MultiplierError = 6
	// A bid multiplier must be specified
	MultiplierErrorEnum_NO_MULTIPLIER_SPECIFIED MultiplierErrorEnum_MultiplierError = 7
	// Multiplier causes bid to exceed daily budget
	MultiplierErrorEnum_MULTIPLIER_CAUSES_BID_TO_EXCEED_DAILY_BUDGET MultiplierErrorEnum_MultiplierError = 8
	// Multiplier causes bid to exceed monthly budget
	MultiplierErrorEnum_MULTIPLIER_CAUSES_BID_TO_EXCEED_MONTHLY_BUDGET MultiplierErrorEnum_MultiplierError = 9
	// Multiplier causes bid to exceed custom budget
	MultiplierErrorEnum_MULTIPLIER_CAUSES_BID_TO_EXCEED_CUSTOM_BUDGET MultiplierErrorEnum_MultiplierError = 10
	// Multiplier causes bid to exceed maximum allowed bid
	MultiplierErrorEnum_MULTIPLIER_CAUSES_BID_TO_EXCEED_MAX_ALLOWED_BID MultiplierErrorEnum_MultiplierError = 11
	// Multiplier causes bid to become less than the minimum bid allowed
	MultiplierErrorEnum_BID_LESS_THAN_MIN_ALLOWED_BID_WITH_MULTIPLIER MultiplierErrorEnum_MultiplierError = 12
	// Multiplier type (cpc vs. cpm) needs to match campaign's bidding strategy
	MultiplierErrorEnum_MULTIPLIER_AND_BIDDING_STRATEGY_TYPE_MISMATCH MultiplierErrorEnum_MultiplierError = 13
)

var MultiplierErrorEnum_MultiplierError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "MULTIPLIER_TOO_HIGH",
	3:  "MULTIPLIER_TOO_LOW",
	4:  "TOO_MANY_FRACTIONAL_DIGITS",
	5:  "MULTIPLIER_NOT_ALLOWED_FOR_BIDDING_STRATEGY",
	6:  "MULTIPLIER_NOT_ALLOWED_WHEN_BASE_BID_IS_MISSING",
	7:  "NO_MULTIPLIER_SPECIFIED",
	8:  "MULTIPLIER_CAUSES_BID_TO_EXCEED_DAILY_BUDGET",
	9:  "MULTIPLIER_CAUSES_BID_TO_EXCEED_MONTHLY_BUDGET",
	10: "MULTIPLIER_CAUSES_BID_TO_EXCEED_CUSTOM_BUDGET",
	11: "MULTIPLIER_CAUSES_BID_TO_EXCEED_MAX_ALLOWED_BID",
	12: "BID_LESS_THAN_MIN_ALLOWED_BID_WITH_MULTIPLIER",
	13: "MULTIPLIER_AND_BIDDING_STRATEGY_TYPE_MISMATCH",
}

var MultiplierErrorEnum_MultiplierError_value = map[string]int32{
	"UNSPECIFIED":                0,
	"UNKNOWN":                    1,
	"MULTIPLIER_TOO_HIGH":        2,
	"MULTIPLIER_TOO_LOW":         3,
	"TOO_MANY_FRACTIONAL_DIGITS": 4,
	"MULTIPLIER_NOT_ALLOWED_FOR_BIDDING_STRATEGY":     5,
	"MULTIPLIER_NOT_ALLOWED_WHEN_BASE_BID_IS_MISSING": 6,
	"NO_MULTIPLIER_SPECIFIED":                         7,
	"MULTIPLIER_CAUSES_BID_TO_EXCEED_DAILY_BUDGET":    8,
	"MULTIPLIER_CAUSES_BID_TO_EXCEED_MONTHLY_BUDGET":  9,
	"MULTIPLIER_CAUSES_BID_TO_EXCEED_CUSTOM_BUDGET":   10,
	"MULTIPLIER_CAUSES_BID_TO_EXCEED_MAX_ALLOWED_BID": 11,
	"BID_LESS_THAN_MIN_ALLOWED_BID_WITH_MULTIPLIER":   12,
	"MULTIPLIER_AND_BIDDING_STRATEGY_TYPE_MISMATCH":   13,
}

func (x MultiplierErrorEnum_MultiplierError) String() string {
	return proto.EnumName(MultiplierErrorEnum_MultiplierError_name, int32(x))
}

func (MultiplierErrorEnum_MultiplierError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3c59bb3bff2d3683, []int{0, 0}
}

// Container for enum describing possible multiplier errors.
type MultiplierErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiplierErrorEnum) Reset()         { *m = MultiplierErrorEnum{} }
func (m *MultiplierErrorEnum) String() string { return proto.CompactTextString(m) }
func (*MultiplierErrorEnum) ProtoMessage()    {}
func (*MultiplierErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c59bb3bff2d3683, []int{0}
}

func (m *MultiplierErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiplierErrorEnum.Unmarshal(m, b)
}
func (m *MultiplierErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiplierErrorEnum.Marshal(b, m, deterministic)
}
func (m *MultiplierErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiplierErrorEnum.Merge(m, src)
}
func (m *MultiplierErrorEnum) XXX_Size() int {
	return xxx_messageInfo_MultiplierErrorEnum.Size(m)
}
func (m *MultiplierErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiplierErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MultiplierErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.MultiplierErrorEnum_MultiplierError", MultiplierErrorEnum_MultiplierError_name, MultiplierErrorEnum_MultiplierError_value)
	proto.RegisterType((*MultiplierErrorEnum)(nil), "google.ads.googleads.v2.errors.MultiplierErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/multiplier_error.proto", fileDescriptor_3c59bb3bff2d3683)
}

var fileDescriptor_3c59bb3bff2d3683 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0xc7, 0xbf, 0x76, 0xfd, 0x36, 0x70, 0x41, 0xb3, 0x3c, 0xc4, 0xa4, 0x81, 0x7a, 0xd1, 0x5b,
	0x20, 0x81, 0x4c, 0xdc, 0x84, 0x2b, 0x27, 0x76, 0x13, 0x8b, 0xc4, 0xae, 0x6a, 0xa7, 0x5d, 0x51,
	0x25, 0xab, 0xd0, 0x2a, 0xaa, 0xd4, 0x26, 0x55, 0xd2, 0xed, 0x81, 0xb8, 0xe4, 0x05, 0x78, 0x07,
	0x6e, 0x78, 0x0f, 0x6e, 0x78, 0x05, 0x94, 0x9a, 0x76, 0x55, 0x60, 0xec, 0x2a, 0x27, 0xc7, 0xff,
	0xdf, 0x39, 0x3e, 0x47, 0x7f, 0x83, 0xb7, 0x69, 0x9e, 0xa7, 0xcb, 0xb9, 0x3d, 0x9d, 0x95, 0xb6,
	0x09, 0xab, 0xe8, 0xc6, 0xb1, 0xe7, 0x45, 0x91, 0x17, 0xa5, 0xbd, 0xba, 0x5e, 0x6e, 0x16, 0xeb,
	0xe5, 0x62, 0x5e, 0xe8, 0x6d, 0xc6, 0x5a, 0x17, 0xf9, 0x26, 0x47, 0x1d, 0xa3, 0xb5, 0xa6, 0xb3,
	0xd2, 0xda, 0x63, 0xd6, 0x8d, 0x63, 0x19, 0xec, 0xe2, 0xf9, 0xae, 0xec, 0x7a, 0x61, 0x4f, 0xb3,
	0x2c, 0xdf, 0x4c, 0x37, 0x8b, 0x3c, 0x2b, 0x0d, 0xdd, 0xfd, 0xde, 0x02, 0x67, 0xf1, 0xbe, 0x30,
	0xad, 0x10, 0x9a, 0x5d, 0xaf, 0xba, 0x5f, 0x5b, 0xe0, 0xb4, 0x96, 0x47, 0xa7, 0xa0, 0x9d, 0x70,
	0xd9, 0xa7, 0x3e, 0xeb, 0x31, 0x4a, 0xe0, 0x7f, 0xa8, 0x0d, 0x4e, 0x12, 0xfe, 0x9e, 0x8b, 0x11,
	0x87, 0x0d, 0x74, 0x0e, 0xce, 0xe2, 0x24, 0x52, 0xac, 0x1f, 0x31, 0x3a, 0xd0, 0x4a, 0x08, 0x1d,
	0xb2, 0x20, 0x84, 0x4d, 0xf4, 0x14, 0xa0, 0xda, 0x41, 0x24, 0x46, 0xf0, 0x08, 0x75, 0xc0, 0x45,
	0xf5, 0x13, 0x63, 0x3e, 0xd6, 0xbd, 0x01, 0xf6, 0x15, 0x13, 0x1c, 0x47, 0x9a, 0xb0, 0x80, 0x29,
	0x09, 0x5b, 0xc8, 0x06, 0x2f, 0x0e, 0x38, 0x2e, 0x94, 0xc6, 0x51, 0x24, 0x46, 0x94, 0xe8, 0x9e,
	0x18, 0x68, 0x8f, 0x11, 0xc2, 0x78, 0xa0, 0xa5, 0x1a, 0x60, 0x45, 0x83, 0x31, 0xfc, 0x1f, 0x5d,
	0x02, 0xfb, 0x0e, 0x60, 0x14, 0x52, 0xae, 0x3d, 0x2c, 0x69, 0x85, 0x69, 0x26, 0x75, 0xcc, 0xa4,
	0x64, 0x3c, 0x80, 0xc7, 0xe8, 0x19, 0x38, 0xe7, 0x42, 0x1f, 0x70, 0xb7, 0x03, 0x9e, 0xa0, 0xd7,
	0xe0, 0xe5, 0xc1, 0x89, 0x8f, 0x13, 0x49, 0xe5, 0xb6, 0x84, 0x12, 0x9a, 0x5e, 0xf9, 0x94, 0x12,
	0x4d, 0x30, 0x8b, 0xc6, 0xda, 0x4b, 0x48, 0x40, 0x15, 0x7c, 0x80, 0x1c, 0x60, 0xdd, 0x47, 0xc4,
	0x82, 0xab, 0xf0, 0x96, 0x79, 0x88, 0xde, 0x80, 0x57, 0xf7, 0x31, 0x7e, 0x22, 0x95, 0x88, 0x77,
	0x08, 0xa8, 0x8d, 0xfa, 0xf7, 0x36, 0xf8, 0x6a, 0x3f, 0xbf, 0xc7, 0x08, 0x6c, 0x57, 0x7d, 0x2a,
	0x49, 0x44, 0xa5, 0xd4, 0x2a, 0xc4, 0x5c, 0xc7, 0x8c, 0x1f, 0x4a, 0xf4, 0x88, 0xa9, 0xf0, 0x60,
	0x15, 0xf0, 0x51, 0xed, 0x6a, 0x98, 0x93, 0x3f, 0xf6, 0xae, 0xd5, 0xb8, 0x4f, 0xab, 0x7d, 0xc6,
	0x58, 0xf9, 0x21, 0x7c, 0xec, 0xfd, 0x6c, 0x80, 0xee, 0xa7, 0x7c, 0x65, 0xfd, 0xdb, 0x96, 0xde,
	0x93, 0x9a, 0xbb, 0xfa, 0x95, 0x1d, 0xfb, 0x8d, 0x0f, 0xe4, 0x37, 0x97, 0xe6, 0xcb, 0x69, 0x96,
	0x5a, 0x79, 0x91, 0xda, 0xe9, 0x3c, 0xdb, 0x9a, 0x75, 0xf7, 0x2a, 0xd6, 0x8b, 0xf2, 0xae, 0x47,
	0xf2, 0xce, 0x7c, 0x3e, 0x37, 0x8f, 0x02, 0x8c, 0xbf, 0x34, 0x3b, 0x81, 0x29, 0x86, 0x67, 0xa5,
	0x65, 0xc2, 0x2a, 0x1a, 0x3a, 0xd6, 0xb6, 0x65, 0xf9, 0x6d, 0x27, 0x98, 0xe0, 0x59, 0x39, 0xd9,
	0x0b, 0x26, 0x43, 0x67, 0x62, 0x04, 0x3f, 0x9a, 0x5d, 0x93, 0x75, 0x5d, 0x3c, 0x2b, 0x5d, 0x77,
	0x2f, 0x71, 0xdd, 0xa1, 0xe3, 0xba, 0x46, 0xf4, 0xf1, 0x78, 0x7b, 0xbb, 0xcb, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x41, 0x92, 0x70, 0x6b, 0xc1, 0x03, 0x00, 0x00,
}
