// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1/run_asset_discovery_response.proto

package securitycenter

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

// The state of an asset discovery run.
type RunAssetDiscoveryResponse_State int32

const (
	// Asset discovery run state was unspecified.
	RunAssetDiscoveryResponse_STATE_UNSPECIFIED RunAssetDiscoveryResponse_State = 0
	// Asset discovery run completed successfully.
	RunAssetDiscoveryResponse_COMPLETED RunAssetDiscoveryResponse_State = 1
	// Asset discovery run was cancelled with tasks still pending, as another
	// run for the same organization was started with a higher priority.
	RunAssetDiscoveryResponse_SUPERSEDED RunAssetDiscoveryResponse_State = 2
	// Asset discovery run was killed and terminated.
	RunAssetDiscoveryResponse_TERMINATED RunAssetDiscoveryResponse_State = 3
)

var RunAssetDiscoveryResponse_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "COMPLETED",
	2: "SUPERSEDED",
	3: "TERMINATED",
}

var RunAssetDiscoveryResponse_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"COMPLETED":         1,
	"SUPERSEDED":        2,
	"TERMINATED":        3,
}

func (x RunAssetDiscoveryResponse_State) String() string {
	return proto.EnumName(RunAssetDiscoveryResponse_State_name, int32(x))
}

func (RunAssetDiscoveryResponse_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_421ea6ed240ae1f7, []int{0, 0}
}

// Response of asset discovery run
type RunAssetDiscoveryResponse struct {
	// The state of an asset discovery run.
	State RunAssetDiscoveryResponse_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.securitycenter.v1.RunAssetDiscoveryResponse_State" json:"state,omitempty"`
	// The duration between asset discovery run start and end
	Duration             *duration.Duration `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RunAssetDiscoveryResponse) Reset()         { *m = RunAssetDiscoveryResponse{} }
func (m *RunAssetDiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*RunAssetDiscoveryResponse) ProtoMessage()    {}
func (*RunAssetDiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_421ea6ed240ae1f7, []int{0}
}

func (m *RunAssetDiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunAssetDiscoveryResponse.Unmarshal(m, b)
}
func (m *RunAssetDiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunAssetDiscoveryResponse.Marshal(b, m, deterministic)
}
func (m *RunAssetDiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunAssetDiscoveryResponse.Merge(m, src)
}
func (m *RunAssetDiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_RunAssetDiscoveryResponse.Size(m)
}
func (m *RunAssetDiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RunAssetDiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RunAssetDiscoveryResponse proto.InternalMessageInfo

func (m *RunAssetDiscoveryResponse) GetState() RunAssetDiscoveryResponse_State {
	if m != nil {
		return m.State
	}
	return RunAssetDiscoveryResponse_STATE_UNSPECIFIED
}

func (m *RunAssetDiscoveryResponse) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.securitycenter.v1.RunAssetDiscoveryResponse_State", RunAssetDiscoveryResponse_State_name, RunAssetDiscoveryResponse_State_value)
	proto.RegisterType((*RunAssetDiscoveryResponse)(nil), "google.cloud.securitycenter.v1.RunAssetDiscoveryResponse")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1/run_asset_discovery_response.proto", fileDescriptor_421ea6ed240ae1f7)
}

var fileDescriptor_421ea6ed240ae1f7 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xd1, 0xca, 0xd3, 0x30,
	0x18, 0x86, 0x6d, 0xe5, 0x17, 0x8d, 0xf8, 0x33, 0x0b, 0xc2, 0x36, 0x64, 0xce, 0x1d, 0xed, 0x28,
	0xa1, 0x13, 0x4f, 0xea, 0x81, 0xd4, 0x36, 0xca, 0x60, 0x9b, 0xa5, 0xed, 0x76, 0x20, 0x83, 0x92,
	0x75, 0xb1, 0x14, 0xb6, 0xa4, 0x24, 0xe9, 0x60, 0xb7, 0xe4, 0xa5, 0x78, 0x09, 0x5e, 0x82, 0x57,
	0xe0, 0xa1, 0x34, 0x69, 0x85, 0x0a, 0xd3, 0xc3, 0xf6, 0x7d, 0xbe, 0xe7, 0x7b, 0x3f, 0x08, 0xf0,
	0x0b, 0xce, 0x8b, 0x13, 0x45, 0xf9, 0x89, 0xd7, 0x47, 0x24, 0x69, 0x5e, 0x8b, 0x52, 0x5d, 0x73,
	0xca, 0x14, 0x15, 0xe8, 0xe2, 0x22, 0x51, 0xb3, 0x8c, 0x48, 0x49, 0x55, 0x76, 0x2c, 0x65, 0xce,
	0x2f, 0x54, 0x5c, 0x33, 0x41, 0x65, 0xc5, 0x99, 0xa4, 0xb0, 0x12, 0x5c, 0x71, 0x67, 0x62, 0x14,
	0x50, 0x2b, 0x60, 0x5f, 0x01, 0x2f, 0xee, 0xb8, 0xcd, 0x91, 0xa6, 0x0f, 0xf5, 0x57, 0x74, 0xac,
	0x05, 0x51, 0x25, 0x67, 0x66, 0x7e, 0xfc, 0xea, 0xef, 0x5c, 0x95, 0x67, 0x2a, 0x15, 0x39, 0x57,
	0x2d, 0xf0, 0xb2, 0x05, 0x48, 0x55, 0x22, 0xc2, 0x18, 0x57, 0x7a, 0x5a, 0x9a, 0x74, 0xf6, 0xcb,
	0x02, 0xa3, 0xb8, 0x66, 0x7e, 0x53, 0x32, 0xec, 0x3a, 0xc6, 0x6d, 0x45, 0x67, 0x0b, 0xee, 0xa4,
	0x22, 0x8a, 0x0e, 0xad, 0xa9, 0x35, 0xbf, 0x5f, 0xbc, 0x87, 0xff, 0x2e, 0x0b, 0x6f, 0x9a, 0x60,
	0xd2, 0x68, 0x62, 0x63, 0x73, 0xde, 0x82, 0xc7, 0xdd, 0x15, 0x43, 0x7b, 0x6a, 0xcd, 0x9f, 0x2e,
	0x46, 0x9d, 0xb9, 0x3b, 0x03, 0x86, 0x2d, 0x10, 0xff, 0x41, 0x67, 0x6b, 0x70, 0xa7, 0x35, 0xce,
	0x0b, 0xf0, 0x3c, 0x49, 0xfd, 0x14, 0x67, 0xdb, 0x4d, 0x12, 0xe1, 0x60, 0xf9, 0x71, 0x89, 0xc3,
	0xc1, 0x03, 0xe7, 0x19, 0x78, 0x12, 0x7c, 0x5e, 0x47, 0x2b, 0x9c, 0xe2, 0x70, 0x60, 0x39, 0xf7,
	0x00, 0x24, 0xdb, 0x08, 0xc7, 0x09, 0x0e, 0x71, 0x38, 0xb0, 0x9b, 0xef, 0x14, 0xc7, 0xeb, 0xe5,
	0xc6, 0x6f, 0xf2, 0x87, 0x1f, 0x7e, 0x58, 0x60, 0x96, 0xf3, 0xf3, 0x7f, 0x6e, 0x8a, 0xac, 0x2f,
	0xab, 0x96, 0x28, 0xf8, 0x89, 0xb0, 0x02, 0x72, 0x51, 0xa0, 0x82, 0x32, 0xdd, 0x14, 0x99, 0x88,
	0x54, 0xa5, 0xbc, 0xf5, 0x08, 0xde, 0xf5, 0xff, 0x7c, 0xb3, 0x27, 0x9f, 0x8c, 0x2e, 0xd0, 0x0b,
	0x93, 0x36, 0x0d, 0xcc, 0xc2, 0x9d, 0xfb, 0xbd, 0x03, 0xf6, 0x1a, 0xd8, 0xf7, 0x81, 0xfd, 0xce,
	0xfd, 0x69, 0xbf, 0x36, 0x80, 0xe7, 0x69, 0xc2, 0xf3, 0xfa, 0x88, 0xe7, 0xed, 0xdc, 0xc3, 0x23,
	0x5d, 0xef, 0xcd, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x50, 0x42, 0xf9, 0x08, 0xa2, 0x02, 0x00,
	0x00,
}
