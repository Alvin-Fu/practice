// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/search_term_match_type.proto

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

// Possible match types for a keyword triggering an ad, including variants.
type SearchTermMatchTypeEnum_SearchTermMatchType int32

const (
	// Not specified.
	SearchTermMatchTypeEnum_UNSPECIFIED SearchTermMatchTypeEnum_SearchTermMatchType = 0
	// Used for return value only. Represents value unknown in this version.
	SearchTermMatchTypeEnum_UNKNOWN SearchTermMatchTypeEnum_SearchTermMatchType = 1
	// Broad match.
	SearchTermMatchTypeEnum_BROAD SearchTermMatchTypeEnum_SearchTermMatchType = 2
	// Exact match.
	SearchTermMatchTypeEnum_EXACT SearchTermMatchTypeEnum_SearchTermMatchType = 3
	// Phrase match.
	SearchTermMatchTypeEnum_PHRASE SearchTermMatchTypeEnum_SearchTermMatchType = 4
	// Exact match (close variant).
	SearchTermMatchTypeEnum_NEAR_EXACT SearchTermMatchTypeEnum_SearchTermMatchType = 5
	// Phrase match (close variant).
	SearchTermMatchTypeEnum_NEAR_PHRASE SearchTermMatchTypeEnum_SearchTermMatchType = 6
)

var SearchTermMatchTypeEnum_SearchTermMatchType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "BROAD",
	3: "EXACT",
	4: "PHRASE",
	5: "NEAR_EXACT",
	6: "NEAR_PHRASE",
}

var SearchTermMatchTypeEnum_SearchTermMatchType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"BROAD":       2,
	"EXACT":       3,
	"PHRASE":      4,
	"NEAR_EXACT":  5,
	"NEAR_PHRASE": 6,
}

func (x SearchTermMatchTypeEnum_SearchTermMatchType) String() string {
	return proto.EnumName(SearchTermMatchTypeEnum_SearchTermMatchType_name, int32(x))
}

func (SearchTermMatchTypeEnum_SearchTermMatchType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3e3d2a8f5a86c467, []int{0, 0}
}

// Container for enum describing match types for a keyword triggering an ad.
type SearchTermMatchTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchTermMatchTypeEnum) Reset()         { *m = SearchTermMatchTypeEnum{} }
func (m *SearchTermMatchTypeEnum) String() string { return proto.CompactTextString(m) }
func (*SearchTermMatchTypeEnum) ProtoMessage()    {}
func (*SearchTermMatchTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e3d2a8f5a86c467, []int{0}
}

func (m *SearchTermMatchTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchTermMatchTypeEnum.Unmarshal(m, b)
}
func (m *SearchTermMatchTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchTermMatchTypeEnum.Marshal(b, m, deterministic)
}
func (m *SearchTermMatchTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchTermMatchTypeEnum.Merge(m, src)
}
func (m *SearchTermMatchTypeEnum) XXX_Size() int {
	return xxx_messageInfo_SearchTermMatchTypeEnum.Size(m)
}
func (m *SearchTermMatchTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchTermMatchTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SearchTermMatchTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.SearchTermMatchTypeEnum_SearchTermMatchType", SearchTermMatchTypeEnum_SearchTermMatchType_name, SearchTermMatchTypeEnum_SearchTermMatchType_value)
	proto.RegisterType((*SearchTermMatchTypeEnum)(nil), "google.ads.googleads.v2.enums.SearchTermMatchTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/search_term_match_type.proto", fileDescriptor_3e3d2a8f5a86c467)
}

var fileDescriptor_3e3d2a8f5a86c467 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdf, 0x4a, 0xf3, 0x30,
	0x18, 0xc6, 0xbf, 0x76, 0xdf, 0x26, 0x66, 0xa0, 0xa5, 0x1e, 0x28, 0xe2, 0x0e, 0xb6, 0x0b, 0x48,
	0xa1, 0x9e, 0xc5, 0xa3, 0x74, 0xab, 0x73, 0x88, 0x5d, 0xd9, 0x3f, 0x45, 0x0a, 0x23, 0xae, 0xa1,
	0x0e, 0x96, 0xa4, 0x34, 0xdd, 0x60, 0x97, 0xe1, 0x2d, 0x78, 0xe8, 0xa5, 0x78, 0x29, 0x82, 0xf7,
	0x20, 0x49, 0xb6, 0x1e, 0x4d, 0x4f, 0xca, 0xaf, 0x79, 0xde, 0xe7, 0xe1, 0x7d, 0x1f, 0x80, 0x32,
	0x21, 0xb2, 0x15, 0xf5, 0x48, 0x2a, 0x3d, 0x83, 0x8a, 0x36, 0xbe, 0x47, 0xf9, 0x9a, 0x49, 0x4f,
	0x52, 0x52, 0x2c, 0x5e, 0xe7, 0x25, 0x2d, 0xd8, 0x9c, 0x91, 0x52, 0xe1, 0x36, 0xa7, 0x30, 0x2f,
	0x44, 0x29, 0xdc, 0x96, 0x31, 0x40, 0x92, 0x4a, 0x58, 0x79, 0xe1, 0xc6, 0x87, 0xda, 0x7b, 0x79,
	0xb5, 0x8f, 0xce, 0x97, 0x1e, 0xe1, 0x5c, 0x94, 0xa4, 0x5c, 0x0a, 0x2e, 0x8d, 0xb9, 0xf3, 0x66,
	0x81, 0xf3, 0xb1, 0x4e, 0x9f, 0xd0, 0x82, 0x3d, 0xa8, 0xec, 0xc9, 0x36, 0xa7, 0x21, 0x5f, 0xb3,
	0xce, 0x06, 0x9c, 0x1d, 0x90, 0xdc, 0x53, 0xd0, 0x9c, 0x46, 0xe3, 0x38, 0xec, 0x0e, 0x6e, 0x07,
	0x61, 0xcf, 0xf9, 0xe7, 0x36, 0xc1, 0xd1, 0x34, 0xba, 0x8f, 0x86, 0x8f, 0x91, 0x63, 0xb9, 0xc7,
	0xa0, 0x1e, 0x8c, 0x86, 0xb8, 0xe7, 0xd8, 0x0a, 0xc3, 0x27, 0xdc, 0x9d, 0x38, 0x35, 0x17, 0x80,
	0x46, 0x7c, 0x37, 0xc2, 0xe3, 0xd0, 0xf9, 0xef, 0x9e, 0x00, 0x10, 0x85, 0x78, 0x34, 0x37, 0x5a,
	0x5d, 0xe5, 0xe9, 0xff, 0xdd, 0x40, 0x23, 0xf8, 0xb6, 0x40, 0x7b, 0x21, 0x18, 0xfc, 0xf3, 0xae,
	0xe0, 0xe2, 0xc0, 0x6e, 0xb1, 0xba, 0x29, 0xb6, 0x9e, 0x83, 0x9d, 0x35, 0x13, 0x2b, 0xc2, 0x33,
	0x28, 0x8a, 0xcc, 0xcb, 0x28, 0xd7, 0x17, 0xef, 0xeb, 0xcd, 0x97, 0xf2, 0x97, 0xb6, 0x6f, 0xf4,
	0xf7, 0xdd, 0xae, 0xf5, 0x31, 0xfe, 0xb0, 0x5b, 0x7d, 0x13, 0x85, 0x53, 0x09, 0x0d, 0x2a, 0x9a,
	0xf9, 0x50, 0x55, 0x24, 0x3f, 0xf7, 0x7a, 0x82, 0x53, 0x99, 0x54, 0x7a, 0x32, 0xf3, 0x13, 0xad,
	0x7f, 0xd9, 0x6d, 0xf3, 0x88, 0x10, 0x4e, 0x25, 0x42, 0xd5, 0x04, 0x42, 0x33, 0x1f, 0x21, 0x3d,
	0xf3, 0xd2, 0xd0, 0x8b, 0x5d, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xcb, 0x85, 0x38, 0x05,
	0x02, 0x00, 0x00,
}
