// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/resultstore/v2/download_metadata.proto

package resultstore

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

// The download metadata for an invocation
type DownloadMetadata struct {
	// The name of the download metadata.  Its format will be:
	// invocations/${INVOCATION_ID}/downloadMetadata
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Indicates the upload status of the invocation, whether it is
	// post-processing, or immutable, etc.
	UploadStatus         UploadStatus `protobuf:"varint,2,opt,name=upload_status,json=uploadStatus,proto3,enum=google.devtools.resultstore.v2.UploadStatus" json:"upload_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DownloadMetadata) Reset()         { *m = DownloadMetadata{} }
func (m *DownloadMetadata) String() string { return proto.CompactTextString(m) }
func (*DownloadMetadata) ProtoMessage()    {}
func (*DownloadMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c835a18f2d71117, []int{0}
}

func (m *DownloadMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadMetadata.Unmarshal(m, b)
}
func (m *DownloadMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadMetadata.Marshal(b, m, deterministic)
}
func (m *DownloadMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadMetadata.Merge(m, src)
}
func (m *DownloadMetadata) XXX_Size() int {
	return xxx_messageInfo_DownloadMetadata.Size(m)
}
func (m *DownloadMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadMetadata proto.InternalMessageInfo

func (m *DownloadMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DownloadMetadata) GetUploadStatus() UploadStatus {
	if m != nil {
		return m.UploadStatus
	}
	return UploadStatus_UPLOAD_STATUS_UNSPECIFIED
}

func init() {
	proto.RegisterType((*DownloadMetadata)(nil), "google.devtools.resultstore.v2.DownloadMetadata")
}

func init() {
	proto.RegisterFile("google/devtools/resultstore/v2/download_metadata.proto", fileDescriptor_9c835a18f2d71117)
}

var fileDescriptor_9c835a18f2d71117 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xb1, 0x4b, 0xc6, 0x30,
	0x10, 0xc5, 0x89, 0x88, 0x60, 0x50, 0x91, 0x4c, 0xc5, 0x41, 0x4a, 0xa7, 0x82, 0x92, 0x40, 0x05,
	0x17, 0x37, 0x71, 0x71, 0x10, 0xb4, 0xe2, 0xe2, 0x22, 0x67, 0x13, 0x82, 0x90, 0xe4, 0x6a, 0x72,
	0xa9, 0xf8, 0xdf, 0x0b, 0x69, 0x85, 0x2e, 0xdf, 0xd7, 0xed, 0x1e, 0x77, 0xbf, 0xf7, 0xde, 0xf1,
	0x5b, 0x8b, 0x68, 0x9d, 0x51, 0xda, 0x4c, 0x84, 0xe8, 0x92, 0x8a, 0x26, 0x65, 0x47, 0x89, 0x30,
	0x1a, 0x35, 0x75, 0x4a, 0xe3, 0x4f, 0x70, 0x08, 0xfa, 0xc3, 0x1b, 0x02, 0x0d, 0x04, 0x72, 0x8c,
	0x48, 0x28, 0x2e, 0x67, 0x4e, 0xfe, 0x73, 0x72, 0xc5, 0xc9, 0xa9, 0xbb, 0xb8, 0xda, 0xf0, 0x1d,
	0xd0, 0x7b, 0x0c, 0xb3, 0x59, 0xf3, 0xcb, 0xcf, 0x1f, 0x96, 0x9c, 0xa7, 0x25, 0x46, 0x08, 0x7e,
	0x18, 0xc0, 0x9b, 0x8a, 0xd5, 0xac, 0x3d, 0xee, 0xcb, 0x2c, 0x5e, 0xf8, 0x69, 0x1e, 0x4b, 0x9b,
	0x44, 0x40, 0x39, 0x55, 0x07, 0x35, 0x6b, 0xcf, 0xba, 0x6b, 0xb9, 0xbf, 0x8c, 0x7c, 0x2b, 0xd0,
	0x6b, 0x61, 0xfa, 0x93, 0xbc, 0x52, 0xf7, 0xdf, 0xbc, 0x19, 0xd0, 0x6f, 0x18, 0x3c, 0xb3, 0xf7,
	0xc7, 0xe5, 0xc2, 0xa2, 0x83, 0x60, 0x25, 0x46, 0xab, 0xac, 0x09, 0xa5, 0xbe, 0x9a, 0x57, 0x30,
	0x7e, 0xa5, 0x5d, 0xef, 0xde, 0xad, 0xe4, 0xe7, 0x51, 0xa1, 0x6e, 0xfe, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x41, 0xce, 0x56, 0x29, 0x7b, 0x01, 0x00, 0x00,
}
