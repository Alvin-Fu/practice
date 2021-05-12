// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/websecurityscanner/v1beta/scan_run_warning_trace.proto

package websecurityscanner

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Output only.
// Defines a warning message code.
// Next id: 6
type ScanRunWarningTrace_Code int32

const (
	// Default value is never used.
	ScanRunWarningTrace_CODE_UNSPECIFIED ScanRunWarningTrace_Code = 0
	// Indicates that a scan discovered an unexpectedly low number of URLs. This
	// is sometimes caused by complex navigation features or by using a single
	// URL for numerous pages.
	ScanRunWarningTrace_INSUFFICIENT_CRAWL_RESULTS ScanRunWarningTrace_Code = 1
	// Indicates that a scan discovered too many URLs to test, or excessive
	// redundant URLs.
	ScanRunWarningTrace_TOO_MANY_CRAWL_RESULTS ScanRunWarningTrace_Code = 2
	// Indicates that too many tests have been generated for the scan. Customer
	// should try reducing the number of starting URLs, increasing the QPS rate,
	// or narrowing down the scope of the scan using the excluded patterns.
	ScanRunWarningTrace_TOO_MANY_FUZZ_TASKS ScanRunWarningTrace_Code = 3
	// Indicates that a scan is blocked by IAP.
	ScanRunWarningTrace_BLOCKED_BY_IAP ScanRunWarningTrace_Code = 4
)

var ScanRunWarningTrace_Code_name = map[int32]string{
	0: "CODE_UNSPECIFIED",
	1: "INSUFFICIENT_CRAWL_RESULTS",
	2: "TOO_MANY_CRAWL_RESULTS",
	3: "TOO_MANY_FUZZ_TASKS",
	4: "BLOCKED_BY_IAP",
}

var ScanRunWarningTrace_Code_value = map[string]int32{
	"CODE_UNSPECIFIED":           0,
	"INSUFFICIENT_CRAWL_RESULTS": 1,
	"TOO_MANY_CRAWL_RESULTS":     2,
	"TOO_MANY_FUZZ_TASKS":        3,
	"BLOCKED_BY_IAP":             4,
}

func (x ScanRunWarningTrace_Code) String() string {
	return proto.EnumName(ScanRunWarningTrace_Code_name, int32(x))
}

func (ScanRunWarningTrace_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8ce6d2f19679e52d, []int{0, 0}
}

// Output only.
// Defines a warning trace message for ScanRun. Warning traces provide customers
// with useful information that helps make the scanning process more effective.
type ScanRunWarningTrace struct {
	// Indicates the warning code.
	Code                 ScanRunWarningTrace_Code `protobuf:"varint,1,opt,name=code,proto3,enum=google.cloud.websecurityscanner.v1beta.ScanRunWarningTrace_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ScanRunWarningTrace) Reset()         { *m = ScanRunWarningTrace{} }
func (m *ScanRunWarningTrace) String() string { return proto.CompactTextString(m) }
func (*ScanRunWarningTrace) ProtoMessage()    {}
func (*ScanRunWarningTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ce6d2f19679e52d, []int{0}
}

func (m *ScanRunWarningTrace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanRunWarningTrace.Unmarshal(m, b)
}
func (m *ScanRunWarningTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanRunWarningTrace.Marshal(b, m, deterministic)
}
func (m *ScanRunWarningTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanRunWarningTrace.Merge(m, src)
}
func (m *ScanRunWarningTrace) XXX_Size() int {
	return xxx_messageInfo_ScanRunWarningTrace.Size(m)
}
func (m *ScanRunWarningTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanRunWarningTrace.DiscardUnknown(m)
}

var xxx_messageInfo_ScanRunWarningTrace proto.InternalMessageInfo

func (m *ScanRunWarningTrace) GetCode() ScanRunWarningTrace_Code {
	if m != nil {
		return m.Code
	}
	return ScanRunWarningTrace_CODE_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("google.cloud.websecurityscanner.v1beta.ScanRunWarningTrace_Code", ScanRunWarningTrace_Code_name, ScanRunWarningTrace_Code_value)
	proto.RegisterType((*ScanRunWarningTrace)(nil), "google.cloud.websecurityscanner.v1beta.ScanRunWarningTrace")
}

func init() {
	proto.RegisterFile("google/cloud/websecurityscanner/v1beta/scan_run_warning_trace.proto", fileDescriptor_8ce6d2f19679e52d)
}

var fileDescriptor_8ce6d2f19679e52d = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd1, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x07, 0x70, 0x3b, 0x87, 0x87, 0x1c, 0x46, 0xc9, 0x44, 0xc7, 0x0e, 0x22, 0x3b, 0x0c, 0xf1,
	0x90, 0xa2, 0x1e, 0xbd, 0xd8, 0x66, 0xad, 0x94, 0xcd, 0xb6, 0x34, 0xad, 0x73, 0x63, 0x10, 0xda,
	0x2e, 0x84, 0xc1, 0x4c, 0x46, 0xd6, 0x3a, 0xfc, 0x00, 0x7e, 0x39, 0xfd, 0x38, 0x7e, 0x01, 0x69,
	0x3b, 0x3c, 0xb8, 0x81, 0x3b, 0x3e, 0xfe, 0xef, 0xfd, 0xc2, 0xcb, 0x03, 0x98, 0x4b, 0xc9, 0x97,
	0xcc, 0xc8, 0x96, 0xb2, 0x98, 0x1b, 0x1b, 0x96, 0xae, 0x59, 0x56, 0xa8, 0x45, 0xfe, 0xbe, 0xce,
	0x12, 0x21, 0x98, 0x32, 0xde, 0x6e, 0x52, 0x96, 0x27, 0x46, 0x59, 0x52, 0x55, 0x08, 0xba, 0x49,
	0x94, 0x58, 0x08, 0x4e, 0x73, 0x95, 0x64, 0x0c, 0xad, 0x94, 0xcc, 0x25, 0xec, 0xd7, 0x08, 0xaa,
	0x10, 0xb4, 0x8b, 0xa0, 0x1a, 0xe9, 0x7d, 0x6b, 0xa0, 0x4d, 0xb2, 0x44, 0x84, 0x85, 0x18, 0xd7,
	0x4c, 0x54, 0x2a, 0x30, 0x02, 0xcd, 0x4c, 0xce, 0x59, 0x47, 0xbb, 0xd4, 0xae, 0x5a, 0xb7, 0x0f,
	0xe8, 0x30, 0x0e, 0xed, 0xa1, 0x10, 0x96, 0x73, 0x16, 0x56, 0x5a, 0xef, 0x43, 0x03, 0xcd, 0xb2,
	0x84, 0xa7, 0x40, 0xc7, 0xfe, 0xc0, 0xa6, 0xb1, 0x47, 0x02, 0x1b, 0xbb, 0x8e, 0x6b, 0x0f, 0xf4,
	0x23, 0x78, 0x01, 0xba, 0xae, 0x47, 0x62, 0xc7, 0x71, 0xb1, 0x6b, 0x7b, 0x11, 0xc5, 0xa1, 0x39,
	0x1e, 0xd1, 0xd0, 0x26, 0xf1, 0x28, 0x22, 0xba, 0x06, 0xbb, 0xe0, 0x2c, 0xf2, 0x7d, 0xfa, 0x64,
	0x7a, 0x93, 0x3f, 0x59, 0x03, 0x9e, 0x83, 0xf6, 0x6f, 0xe6, 0xc4, 0xd3, 0x29, 0x8d, 0x4c, 0x32,
	0x24, 0xfa, 0x31, 0x84, 0xa0, 0x65, 0x8d, 0x7c, 0x3c, 0xb4, 0x07, 0xd4, 0x9a, 0x50, 0xd7, 0x0c,
	0xf4, 0xa6, 0xf5, 0xa5, 0x81, 0xeb, 0x4c, 0xbe, 0x1e, 0xb8, 0x95, 0xd5, 0xd9, 0xb3, 0x56, 0x50,
	0x7e, 0x73, 0xa0, 0x4d, 0x5f, 0xb6, 0x06, 0x97, 0xcb, 0x44, 0x70, 0x24, 0x15, 0x37, 0x38, 0x13,
	0xd5, 0x11, 0x8c, 0x3a, 0x4a, 0x56, 0x8b, 0xf5, 0x7f, 0xc7, 0xbc, 0xdf, 0x4d, 0x3e, 0x1b, 0xfd,
	0xc7, 0x6a, 0x7e, 0x86, 0xcb, 0xd9, 0xd9, 0x98, 0xa5, 0x64, 0xdb, 0x41, 0xea, 0x8e, 0xd9, 0x73,
	0x35, 0x9b, 0x9e, 0x54, 0xaf, 0xdd, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x57, 0x2f, 0x52,
	0x39, 0x02, 0x00, 0x00,
}
