// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/conversion_action_error.proto

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

// Enum describing possible conversion action errors.
type ConversionActionErrorEnum_ConversionActionError int32

const (
	// Enum unspecified.
	ConversionActionErrorEnum_UNSPECIFIED ConversionActionErrorEnum_ConversionActionError = 0
	// The received error code is not known in this version.
	ConversionActionErrorEnum_UNKNOWN ConversionActionErrorEnum_ConversionActionError = 1
	// The specified conversion action name already exists.
	ConversionActionErrorEnum_DUPLICATE_NAME ConversionActionErrorEnum_ConversionActionError = 2
	// Another conversion action with the specified app id already exists.
	ConversionActionErrorEnum_DUPLICATE_APP_ID ConversionActionErrorEnum_ConversionActionError = 3
	// Android first open action conflicts with Google play codeless download
	// action tracking the same app.
	ConversionActionErrorEnum_TWO_CONVERSION_ACTIONS_BIDDING_ON_SAME_APP_DOWNLOAD ConversionActionErrorEnum_ConversionActionError = 4
	// Android first open action conflicts with Google play codeless download
	// action tracking the same app.
	ConversionActionErrorEnum_BIDDING_ON_SAME_APP_DOWNLOAD_AS_GLOBAL_ACTION ConversionActionErrorEnum_ConversionActionError = 5
	// The attribution model cannot be set to DATA_DRIVEN because a data-driven
	// model has never been generated.
	ConversionActionErrorEnum_DATA_DRIVEN_MODEL_WAS_NEVER_GENERATED ConversionActionErrorEnum_ConversionActionError = 6
	// The attribution model cannot be set to DATA_DRIVEN because the
	// data-driven model is expired.
	ConversionActionErrorEnum_DATA_DRIVEN_MODEL_EXPIRED ConversionActionErrorEnum_ConversionActionError = 7
	// The attribution model cannot be set to DATA_DRIVEN because the
	// data-driven model is stale.
	ConversionActionErrorEnum_DATA_DRIVEN_MODEL_STALE ConversionActionErrorEnum_ConversionActionError = 8
	// The attribution model cannot be set to DATA_DRIVEN because the
	// data-driven model is unavailable or the conversion action was newly
	// added.
	ConversionActionErrorEnum_DATA_DRIVEN_MODEL_UNKNOWN ConversionActionErrorEnum_ConversionActionError = 9
)

var ConversionActionErrorEnum_ConversionActionError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "DUPLICATE_NAME",
	3: "DUPLICATE_APP_ID",
	4: "TWO_CONVERSION_ACTIONS_BIDDING_ON_SAME_APP_DOWNLOAD",
	5: "BIDDING_ON_SAME_APP_DOWNLOAD_AS_GLOBAL_ACTION",
	6: "DATA_DRIVEN_MODEL_WAS_NEVER_GENERATED",
	7: "DATA_DRIVEN_MODEL_EXPIRED",
	8: "DATA_DRIVEN_MODEL_STALE",
	9: "DATA_DRIVEN_MODEL_UNKNOWN",
}

var ConversionActionErrorEnum_ConversionActionError_value = map[string]int32{
	"UNSPECIFIED":      0,
	"UNKNOWN":          1,
	"DUPLICATE_NAME":   2,
	"DUPLICATE_APP_ID": 3,
	"TWO_CONVERSION_ACTIONS_BIDDING_ON_SAME_APP_DOWNLOAD": 4,
	"BIDDING_ON_SAME_APP_DOWNLOAD_AS_GLOBAL_ACTION":       5,
	"DATA_DRIVEN_MODEL_WAS_NEVER_GENERATED":               6,
	"DATA_DRIVEN_MODEL_EXPIRED":                           7,
	"DATA_DRIVEN_MODEL_STALE":                             8,
	"DATA_DRIVEN_MODEL_UNKNOWN":                           9,
}

func (x ConversionActionErrorEnum_ConversionActionError) String() string {
	return proto.EnumName(ConversionActionErrorEnum_ConversionActionError_name, int32(x))
}

func (ConversionActionErrorEnum_ConversionActionError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bfc676da5dfef964, []int{0, 0}
}

// Container for enum describing possible conversion action errors.
type ConversionActionErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionActionErrorEnum) Reset()         { *m = ConversionActionErrorEnum{} }
func (m *ConversionActionErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionActionErrorEnum) ProtoMessage()    {}
func (*ConversionActionErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc676da5dfef964, []int{0}
}

func (m *ConversionActionErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionActionErrorEnum.Unmarshal(m, b)
}
func (m *ConversionActionErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionActionErrorEnum.Marshal(b, m, deterministic)
}
func (m *ConversionActionErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionActionErrorEnum.Merge(m, src)
}
func (m *ConversionActionErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionActionErrorEnum.Size(m)
}
func (m *ConversionActionErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionActionErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionActionErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.ConversionActionErrorEnum_ConversionActionError", ConversionActionErrorEnum_ConversionActionError_name, ConversionActionErrorEnum_ConversionActionError_value)
	proto.RegisterType((*ConversionActionErrorEnum)(nil), "google.ads.googleads.v2.errors.ConversionActionErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/conversion_action_error.proto", fileDescriptor_bfc676da5dfef964)
}

var fileDescriptor_bfc676da5dfef964 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x6d, 0x56, 0x77, 0x75, 0x16, 0x74, 0x18, 0x14, 0xd9, 0x55, 0xf7, 0x50, 0xf0, 0xe0,
	0xc1, 0x04, 0xb3, 0x07, 0x21, 0x7a, 0x79, 0xcd, 0x8c, 0x61, 0x30, 0x9d, 0x09, 0x49, 0x9a, 0x8a,
	0x14, 0x86, 0xd8, 0x94, 0x50, 0xd8, 0xcd, 0x94, 0x4c, 0xed, 0x07, 0xf2, 0xe8, 0x47, 0xf1, 0xe6,
	0xd7, 0xf0, 0xe2, 0xc9, 0xbb, 0x24, 0xb3, 0xad, 0x87, 0xad, 0x3d, 0xcd, 0x9f, 0xf7, 0xfe, 0xbf,
	0x7f, 0xc2, 0x7b, 0x0f, 0xbd, 0xaf, 0xb5, 0xae, 0xaf, 0x16, 0x5e, 0x59, 0x19, 0xcf, 0xca, 0x4e,
	0x6d, 0x7c, 0x6f, 0xd1, 0xb6, 0xba, 0x35, 0xde, 0x5c, 0x37, 0x9b, 0x45, 0x6b, 0x96, 0xba, 0x51,
	0xe5, 0x7c, 0xdd, 0x3d, 0x7d, 0xc3, 0x5d, 0xb5, 0x7a, 0xad, 0xc9, 0x85, 0x45, 0xdc, 0xb2, 0x32,
	0xee, 0x8e, 0x76, 0x37, 0xbe, 0x6b, 0xe9, 0xf3, 0xe7, 0xdb, 0xf4, 0xd5, 0xd2, 0x2b, 0x9b, 0x46,
	0xaf, 0xcb, 0x2e, 0xc2, 0x58, 0x7a, 0xf8, 0xdb, 0x41, 0x67, 0xe1, 0x2e, 0x1f, 0xfa, 0x78, 0xd6,
	0x81, 0xac, 0xf9, 0x7a, 0x3d, 0xfc, 0xe9, 0xa0, 0x27, 0x7b, 0xbb, 0xe4, 0x11, 0x3a, 0x9d, 0x88,
	0x2c, 0x61, 0x21, 0xff, 0xc0, 0x19, 0xc5, 0x77, 0xc8, 0x29, 0x3a, 0x99, 0x88, 0x8f, 0x42, 0x4e,
	0x05, 0x1e, 0x10, 0x82, 0x1e, 0xd2, 0x49, 0x12, 0xf3, 0x10, 0x72, 0xa6, 0x04, 0x8c, 0x19, 0x76,
	0xc8, 0x63, 0x84, 0xff, 0xd5, 0x20, 0x49, 0x14, 0xa7, 0xf8, 0x88, 0xbc, 0x45, 0x97, 0xf9, 0x54,
	0xaa, 0x50, 0x8a, 0x82, 0xa5, 0x19, 0x97, 0x42, 0x41, 0x98, 0x73, 0x29, 0x32, 0x35, 0xe2, 0x94,
	0x72, 0x11, 0x29, 0x29, 0x54, 0x06, 0x63, 0x8b, 0x50, 0x39, 0x15, 0xb1, 0x04, 0x8a, 0xef, 0x92,
	0x37, 0xe8, 0xf5, 0x21, 0x87, 0x82, 0x4c, 0x45, 0xb1, 0x1c, 0x41, 0x7c, 0x13, 0x88, 0xef, 0x91,
	0x57, 0xe8, 0x25, 0x85, 0x1c, 0x14, 0x4d, 0x79, 0xc1, 0x84, 0x1a, 0x4b, 0xca, 0x62, 0x35, 0x85,
	0x4c, 0x09, 0x56, 0xb0, 0x54, 0x45, 0x4c, 0xb0, 0x14, 0x72, 0x46, 0xf1, 0x31, 0x79, 0x81, 0xce,
	0x6e, 0x5b, 0xd9, 0xa7, 0x84, 0xa7, 0x8c, 0xe2, 0x13, 0xf2, 0x0c, 0x3d, 0xbd, 0xdd, 0xce, 0x72,
	0x88, 0x19, 0xbe, 0xbf, 0x9f, 0xdd, 0xce, 0xe6, 0xc1, 0xe8, 0xcf, 0x00, 0x0d, 0xe7, 0xfa, 0xda,
	0x3d, 0xbc, 0xb6, 0xd1, 0xf9, 0xde, 0xb9, 0x27, 0xdd, 0xd2, 0x92, 0xc1, 0x67, 0x7a, 0x43, 0xd7,
	0xfa, 0xaa, 0x6c, 0x6a, 0x57, 0xb7, 0xb5, 0x57, 0x2f, 0x9a, 0x7e, 0xa5, 0xdb, 0x13, 0x5a, 0x2d,
	0xcd, 0xff, 0x2e, 0xea, 0x9d, 0x7d, 0xbe, 0x39, 0x47, 0x11, 0xc0, 0x77, 0xe7, 0x22, 0xb2, 0x61,
	0x50, 0x19, 0xd7, 0xca, 0x4e, 0x15, 0xbe, 0xdb, 0x7f, 0xd2, 0xfc, 0xd8, 0x1a, 0x66, 0x50, 0x99,
	0xd9, 0xce, 0x30, 0x2b, 0xfc, 0x99, 0x35, 0xfc, 0x72, 0x86, 0xb6, 0x1a, 0x04, 0x50, 0x99, 0x20,
	0xd8, 0x59, 0x82, 0xa0, 0xf0, 0x83, 0xc0, 0x9a, 0xbe, 0x1c, 0xf7, 0x7f, 0x77, 0xf9, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x3e, 0x6e, 0xc6, 0x2b, 0xee, 0x02, 0x00, 0x00,
}
