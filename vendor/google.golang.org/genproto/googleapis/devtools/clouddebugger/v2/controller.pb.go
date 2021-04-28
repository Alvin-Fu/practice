// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/clouddebugger/v2/controller.proto

package clouddebugger

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request to register a debuggee.
type RegisterDebuggeeRequest struct {
	// Required. Debuggee information to register.
	// The fields `project`, `uniquifier`, `description` and `agent_version`
	// of the debuggee must be set.
	Debuggee             *Debuggee `protobuf:"bytes,1,opt,name=debuggee,proto3" json:"debuggee,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RegisterDebuggeeRequest) Reset()         { *m = RegisterDebuggeeRequest{} }
func (m *RegisterDebuggeeRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterDebuggeeRequest) ProtoMessage()    {}
func (*RegisterDebuggeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_694192a34270926f, []int{0}
}

func (m *RegisterDebuggeeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterDebuggeeRequest.Unmarshal(m, b)
}
func (m *RegisterDebuggeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterDebuggeeRequest.Marshal(b, m, deterministic)
}
func (m *RegisterDebuggeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterDebuggeeRequest.Merge(m, src)
}
func (m *RegisterDebuggeeRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterDebuggeeRequest.Size(m)
}
func (m *RegisterDebuggeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterDebuggeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterDebuggeeRequest proto.InternalMessageInfo

func (m *RegisterDebuggeeRequest) GetDebuggee() *Debuggee {
	if m != nil {
		return m.Debuggee
	}
	return nil
}

// Response for registering a debuggee.
type RegisterDebuggeeResponse struct {
	// Debuggee resource.
	// The field `id` is guaranteed to be set (in addition to the echoed fields).
	// If the field `is_disabled` is set to `true`, the agent should disable
	// itself by removing all breakpoints and detaching from the application.
	// It should however continue to poll `RegisterDebuggee` until reenabled.
	Debuggee             *Debuggee `protobuf:"bytes,1,opt,name=debuggee,proto3" json:"debuggee,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RegisterDebuggeeResponse) Reset()         { *m = RegisterDebuggeeResponse{} }
func (m *RegisterDebuggeeResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterDebuggeeResponse) ProtoMessage()    {}
func (*RegisterDebuggeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_694192a34270926f, []int{1}
}

func (m *RegisterDebuggeeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterDebuggeeResponse.Unmarshal(m, b)
}
func (m *RegisterDebuggeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterDebuggeeResponse.Marshal(b, m, deterministic)
}
func (m *RegisterDebuggeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterDebuggeeResponse.Merge(m, src)
}
func (m *RegisterDebuggeeResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterDebuggeeResponse.Size(m)
}
func (m *RegisterDebuggeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterDebuggeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterDebuggeeResponse proto.InternalMessageInfo

func (m *RegisterDebuggeeResponse) GetDebuggee() *Debuggee {
	if m != nil {
		return m.Debuggee
	}
	return nil
}

// Request to list active breakpoints.
type ListActiveBreakpointsRequest struct {
	// Required. Identifies the debuggee.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId,proto3" json:"debuggee_id,omitempty"`
	// A token that, if specified, blocks the method call until the list
	// of active breakpoints has changed, or a server-selected timeout has
	// expired. The value should be set from the `next_wait_token` field in
	// the last response. The initial value should be set to `"init"`.
	WaitToken string `protobuf:"bytes,2,opt,name=wait_token,json=waitToken,proto3" json:"wait_token,omitempty"`
	// If set to `true` (recommended), returns `google.rpc.Code.OK` status and
	// sets the `wait_expired` response field to `true` when the server-selected
	// timeout has expired.
	//
	// If set to `false` (deprecated), returns `google.rpc.Code.ABORTED` status
	// when the server-selected timeout has expired.
	SuccessOnTimeout     bool     `protobuf:"varint,3,opt,name=success_on_timeout,json=successOnTimeout,proto3" json:"success_on_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListActiveBreakpointsRequest) Reset()         { *m = ListActiveBreakpointsRequest{} }
func (m *ListActiveBreakpointsRequest) String() string { return proto.CompactTextString(m) }
func (*ListActiveBreakpointsRequest) ProtoMessage()    {}
func (*ListActiveBreakpointsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_694192a34270926f, []int{2}
}

func (m *ListActiveBreakpointsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListActiveBreakpointsRequest.Unmarshal(m, b)
}
func (m *ListActiveBreakpointsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListActiveBreakpointsRequest.Marshal(b, m, deterministic)
}
func (m *ListActiveBreakpointsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListActiveBreakpointsRequest.Merge(m, src)
}
func (m *ListActiveBreakpointsRequest) XXX_Size() int {
	return xxx_messageInfo_ListActiveBreakpointsRequest.Size(m)
}
func (m *ListActiveBreakpointsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListActiveBreakpointsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListActiveBreakpointsRequest proto.InternalMessageInfo

func (m *ListActiveBreakpointsRequest) GetDebuggeeId() string {
	if m != nil {
		return m.DebuggeeId
	}
	return ""
}

func (m *ListActiveBreakpointsRequest) GetWaitToken() string {
	if m != nil {
		return m.WaitToken
	}
	return ""
}

func (m *ListActiveBreakpointsRequest) GetSuccessOnTimeout() bool {
	if m != nil {
		return m.SuccessOnTimeout
	}
	return false
}

// Response for listing active breakpoints.
type ListActiveBreakpointsResponse struct {
	// List of all active breakpoints.
	// The fields `id` and `location` are guaranteed to be set on each breakpoint.
	Breakpoints []*Breakpoint `protobuf:"bytes,1,rep,name=breakpoints,proto3" json:"breakpoints,omitempty"`
	// A token that can be used in the next method call to block until
	// the list of breakpoints changes.
	NextWaitToken string `protobuf:"bytes,2,opt,name=next_wait_token,json=nextWaitToken,proto3" json:"next_wait_token,omitempty"`
	// If set to `true`, indicates that there is no change to the
	// list of active breakpoints and the server-selected timeout has expired.
	// The `breakpoints` field would be empty and should be ignored.
	WaitExpired          bool     `protobuf:"varint,3,opt,name=wait_expired,json=waitExpired,proto3" json:"wait_expired,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListActiveBreakpointsResponse) Reset()         { *m = ListActiveBreakpointsResponse{} }
func (m *ListActiveBreakpointsResponse) String() string { return proto.CompactTextString(m) }
func (*ListActiveBreakpointsResponse) ProtoMessage()    {}
func (*ListActiveBreakpointsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_694192a34270926f, []int{3}
}

func (m *ListActiveBreakpointsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListActiveBreakpointsResponse.Unmarshal(m, b)
}
func (m *ListActiveBreakpointsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListActiveBreakpointsResponse.Marshal(b, m, deterministic)
}
func (m *ListActiveBreakpointsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListActiveBreakpointsResponse.Merge(m, src)
}
func (m *ListActiveBreakpointsResponse) XXX_Size() int {
	return xxx_messageInfo_ListActiveBreakpointsResponse.Size(m)
}
func (m *ListActiveBreakpointsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListActiveBreakpointsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListActiveBreakpointsResponse proto.InternalMessageInfo

func (m *ListActiveBreakpointsResponse) GetBreakpoints() []*Breakpoint {
	if m != nil {
		return m.Breakpoints
	}
	return nil
}

func (m *ListActiveBreakpointsResponse) GetNextWaitToken() string {
	if m != nil {
		return m.NextWaitToken
	}
	return ""
}

func (m *ListActiveBreakpointsResponse) GetWaitExpired() bool {
	if m != nil {
		return m.WaitExpired
	}
	return false
}

// Request to update an active breakpoint.
type UpdateActiveBreakpointRequest struct {
	// Required. Identifies the debuggee being debugged.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId,proto3" json:"debuggee_id,omitempty"`
	// Required. Updated breakpoint information.
	// The field `id` must be set.
	// The agent must echo all Breakpoint specification fields in the update.
	Breakpoint           *Breakpoint `protobuf:"bytes,2,opt,name=breakpoint,proto3" json:"breakpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateActiveBreakpointRequest) Reset()         { *m = UpdateActiveBreakpointRequest{} }
func (m *UpdateActiveBreakpointRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateActiveBreakpointRequest) ProtoMessage()    {}
func (*UpdateActiveBreakpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_694192a34270926f, []int{4}
}

func (m *UpdateActiveBreakpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateActiveBreakpointRequest.Unmarshal(m, b)
}
func (m *UpdateActiveBreakpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateActiveBreakpointRequest.Marshal(b, m, deterministic)
}
func (m *UpdateActiveBreakpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateActiveBreakpointRequest.Merge(m, src)
}
func (m *UpdateActiveBreakpointRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateActiveBreakpointRequest.Size(m)
}
func (m *UpdateActiveBreakpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateActiveBreakpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateActiveBreakpointRequest proto.InternalMessageInfo

func (m *UpdateActiveBreakpointRequest) GetDebuggeeId() string {
	if m != nil {
		return m.DebuggeeId
	}
	return ""
}

func (m *UpdateActiveBreakpointRequest) GetBreakpoint() *Breakpoint {
	if m != nil {
		return m.Breakpoint
	}
	return nil
}

// Response for updating an active breakpoint.
// The message is defined to allow future extensions.
type UpdateActiveBreakpointResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateActiveBreakpointResponse) Reset()         { *m = UpdateActiveBreakpointResponse{} }
func (m *UpdateActiveBreakpointResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateActiveBreakpointResponse) ProtoMessage()    {}
func (*UpdateActiveBreakpointResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_694192a34270926f, []int{5}
}

func (m *UpdateActiveBreakpointResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateActiveBreakpointResponse.Unmarshal(m, b)
}
func (m *UpdateActiveBreakpointResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateActiveBreakpointResponse.Marshal(b, m, deterministic)
}
func (m *UpdateActiveBreakpointResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateActiveBreakpointResponse.Merge(m, src)
}
func (m *UpdateActiveBreakpointResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateActiveBreakpointResponse.Size(m)
}
func (m *UpdateActiveBreakpointResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateActiveBreakpointResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateActiveBreakpointResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegisterDebuggeeRequest)(nil), "google.devtools.clouddebugger.v2.RegisterDebuggeeRequest")
	proto.RegisterType((*RegisterDebuggeeResponse)(nil), "google.devtools.clouddebugger.v2.RegisterDebuggeeResponse")
	proto.RegisterType((*ListActiveBreakpointsRequest)(nil), "google.devtools.clouddebugger.v2.ListActiveBreakpointsRequest")
	proto.RegisterType((*ListActiveBreakpointsResponse)(nil), "google.devtools.clouddebugger.v2.ListActiveBreakpointsResponse")
	proto.RegisterType((*UpdateActiveBreakpointRequest)(nil), "google.devtools.clouddebugger.v2.UpdateActiveBreakpointRequest")
	proto.RegisterType((*UpdateActiveBreakpointResponse)(nil), "google.devtools.clouddebugger.v2.UpdateActiveBreakpointResponse")
}

func init() {
	proto.RegisterFile("google/devtools/clouddebugger/v2/controller.proto", fileDescriptor_694192a34270926f)
}

var fileDescriptor_694192a34270926f = []byte{
	// 707 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x6a, 0xd4, 0x40,
	0x18, 0x25, 0x5b, 0x28, 0xed, 0xac, 0xd2, 0x12, 0xb0, 0x5d, 0x62, 0xab, 0xdb, 0x50, 0x4a, 0xa9,
	0x6b, 0x06, 0xa3, 0x20, 0xae, 0xa0, 0x66, 0xab, 0x56, 0xf1, 0xaf, 0x2e, 0xb5, 0x82, 0x54, 0x42,
	0x36, 0x99, 0x4d, 0x87, 0x66, 0x67, 0x62, 0x66, 0xb2, 0x5b, 0x29, 0x45, 0xd0, 0x27, 0x10, 0x6f,
	0x7c, 0x06, 0x1f, 0xc1, 0x8b, 0x5e, 0x78, 0xd9, 0x4b, 0xbd, 0xeb, 0x95, 0x88, 0x0f, 0x22, 0x49,
	0x26, 0x9b, 0x6c, 0xdb, 0xed, 0xb6, 0xeb, 0xe5, 0x9e, 0xf3, 0x7d, 0xe7, 0x3b, 0x67, 0xf6, 0x9b,
	0x0c, 0xb8, 0xe6, 0x52, 0xea, 0x7a, 0x08, 0x3a, 0xa8, 0xcd, 0x29, 0xf5, 0x18, 0xb4, 0x3d, 0x1a,
	0x3a, 0x0e, 0x6a, 0x84, 0xae, 0x8b, 0x02, 0xd8, 0xd6, 0xa1, 0x4d, 0x09, 0x0f, 0xa8, 0xe7, 0xa1,
	0x40, 0xf3, 0x03, 0xca, 0xa9, 0x5c, 0x4e, 0x5a, 0xb4, 0xb4, 0x45, 0xeb, 0x69, 0xd1, 0xda, 0xba,
	0x32, 0x23, 0x44, 0x2d, 0x1f, 0x43, 0x8b, 0x10, 0xca, 0x2d, 0x8e, 0x29, 0x61, 0x49, 0xbf, 0x32,
	0x9d, 0x63, 0x6d, 0x0f, 0x23, 0xc2, 0x05, 0x71, 0x39, 0x47, 0x34, 0x31, 0xf2, 0x1c, 0xb3, 0x81,
	0x36, 0xad, 0x36, 0xa6, 0x62, 0xb2, 0x72, 0x65, 0xa0, 0x59, 0xc7, 0xe2, 0x96, 0x28, 0xbe, 0x28,
	0x8a, 0xe3, 0x5f, 0x8d, 0xb0, 0x09, 0x51, 0xcb, 0xe7, 0xef, 0x13, 0x52, 0x6d, 0x82, 0xe9, 0x3a,
	0x72, 0x31, 0xe3, 0x28, 0xb8, 0x9f, 0xb4, 0xa3, 0x3a, 0x7a, 0x17, 0x22, 0xc6, 0xe5, 0x27, 0x60,
	0x4c, 0x28, 0xa2, 0x92, 0x54, 0x96, 0x16, 0x8b, 0xfa, 0x92, 0x36, 0x28, 0xb1, 0x96, 0x8a, 0xd4,
	0x46, 0x7e, 0x1b, 0x85, 0x7a, 0x57, 0x40, 0x6d, 0x80, 0xd2, 0xd1, 0x39, 0xcc, 0xa7, 0x84, 0x21,
	0xf9, 0xe1, 0xff, 0x0c, 0xca, 0xcd, 0xf8, 0x2c, 0x81, 0x99, 0xa7, 0x98, 0x71, 0xc3, 0xe6, 0xb8,
	0x8d, 0x6a, 0x01, 0xb2, 0xb6, 0x7c, 0x8a, 0x09, 0x67, 0x69, 0xa2, 0x79, 0x50, 0x4c, 0x8b, 0x4d,
	0xec, 0xc4, 0xb3, 0xc6, 0x13, 0xa3, 0x20, 0xc5, 0x1f, 0x3b, 0xf2, 0x2c, 0x00, 0x1d, 0x0b, 0x73,
	0x93, 0xd3, 0x2d, 0x44, 0x4a, 0x85, 0xa8, 0xa8, 0x3e, 0x1e, 0x21, 0x6b, 0x11, 0x20, 0x57, 0x80,
	0xcc, 0x42, 0xdb, 0x46, 0x8c, 0x99, 0x94, 0x98, 0x1c, 0xb7, 0x10, 0x0d, 0x79, 0x69, 0xa4, 0x2c,
	0x2d, 0x8e, 0xd5, 0x27, 0x05, 0xf3, 0x82, 0xac, 0x25, 0xb8, 0xfa, 0x5d, 0x02, 0xb3, 0x7d, 0x3c,
	0x89, 0xf4, 0xcf, 0x41, 0xb1, 0x91, 0xc1, 0x25, 0xa9, 0x3c, 0xb2, 0x58, 0xd4, 0x2b, 0x83, 0x0f,
	0x20, 0xd3, 0xaa, 0xe7, 0x05, 0xe4, 0x05, 0x30, 0x41, 0xd0, 0x36, 0x37, 0x8f, 0x64, 0x38, 0x1f,
	0xc1, 0xaf, 0xbb, 0x39, 0xe6, 0xc0, 0xb9, 0xb8, 0x04, 0x6d, 0xfb, 0x38, 0x40, 0x8e, 0x48, 0x50,
	0x8c, 0xb0, 0x07, 0x09, 0xa4, 0x7e, 0x95, 0xc0, 0xec, 0x2b, 0xdf, 0xb1, 0x38, 0x3a, 0x6c, 0xff,
	0x6c, 0x27, 0xfa, 0x12, 0x80, 0xcc, 0x61, 0xec, 0xe6, 0x8c, 0x09, 0x85, 0x64, 0x26, 0xa2, 0x96,
	0xc1, 0xa5, 0x7e, 0xce, 0x92, 0x73, 0xd5, 0xf7, 0x46, 0x41, 0x71, 0xb9, 0x7b, 0x65, 0x75, 0xf9,
	0x87, 0x04, 0x26, 0x0f, 0xaf, 0xa0, 0x7c, 0x6b, 0xb0, 0x8b, 0x3e, 0xd7, 0x43, 0xa9, 0x0e, 0xd3,
	0x9a, 0x78, 0x53, 0x6f, 0x1e, 0x18, 0xdd, 0xb5, 0xfd, 0xf8, 0xeb, 0xef, 0x97, 0xc2, 0x82, 0x3a,
	0xd7, 0xfb, 0x89, 0x81, 0x29, 0xcd, 0x60, 0x20, 0x54, 0xaa, 0xd2, 0x92, 0xfc, 0x47, 0x02, 0x17,
	0x8e, 0x5d, 0x27, 0xf9, 0xce, 0x60, 0x3b, 0x27, 0xdd, 0x0d, 0xe5, 0xee, 0xd0, 0xfd, 0x22, 0xd3,
	0xa3, 0x03, 0x23, 0xbf, 0x0b, 0x71, 0xac, 0x1b, 0xb2, 0xde, 0x37, 0xd6, 0x4e, 0xae, 0x78, 0x17,
	0xe6, 0x37, 0xf8, 0x53, 0x01, 0x4c, 0x1d, 0xff, 0xe7, 0xca, 0xa7, 0x70, 0x79, 0xe2, 0xc2, 0x2a,
	0xf7, 0x86, 0x17, 0x10, 0x39, 0x9b, 0x07, 0xc6, 0x54, 0xce, 0x7a, 0x25, 0x73, 0x1e, 0x47, 0x5e,
	0x51, 0x6a, 0x67, 0x8f, 0x0c, 0x77, 0xb2, 0x1f, 0x1a, 0x76, 0x76, 0xab, 0xd2, 0x92, 0xf2, 0x61,
	0xdf, 0x98, 0xe9, 0xf5, 0x96, 0x38, 0xb7, 0x7c, 0xcc, 0x34, 0x9b, 0xb6, 0x7e, 0x1a, 0x6f, 0x37,
	0x39, 0xf7, 0x59, 0x15, 0xc2, 0x4e, 0xa7, 0x73, 0x88, 0x84, 0x56, 0xc8, 0x37, 0x93, 0x57, 0xe1,
	0xaa, 0xef, 0x59, 0xbc, 0x49, 0x83, 0x56, 0xe5, 0x54, 0xe5, 0x66, 0x3a, 0xae, 0xb6, 0x27, 0x81,
	0x79, 0x9b, 0xb6, 0x06, 0x1e, 0x58, 0x6d, 0x22, 0xbb, 0x66, 0xab, 0xd1, 0xa3, 0xb2, 0x2a, 0xbd,
	0x79, 0x26, 0x9a, 0x5c, 0xea, 0x59, 0xc4, 0xd5, 0x68, 0xe0, 0x42, 0x17, 0x91, 0xf8, 0xc9, 0x81,
	0xd9, 0xe8, 0xfe, 0xef, 0xd7, 0xed, 0x1e, 0xe0, 0x5b, 0xa1, 0xb4, 0x92, 0xe8, 0x2d, 0x47, 0x70,
	0xfa, 0xed, 0x0f, 0xb4, 0x75, 0x7d, 0x3f, 0xa5, 0x36, 0x62, 0x6a, 0x23, 0xa5, 0x36, 0xd6, 0xf5,
	0xc6, 0x68, 0x3c, 0xef, 0xfa, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0xef, 0x57, 0x76, 0xdb,
	0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Controller2Client is the client API for Controller2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Controller2Client interface {
	// Registers the debuggee with the controller service.
	//
	// All agents attached to the same application must call this method with
	// exactly the same request content to get back the same stable `debuggee_id`.
	// Agents should call this method again whenever `google.rpc.Code.NOT_FOUND`
	// is returned from any controller method.
	//
	// This protocol allows the controller service to disable debuggees, recover
	// from data loss, or change the `debuggee_id` format. Agents must handle
	// `debuggee_id` value changing upon re-registration.
	RegisterDebuggee(ctx context.Context, in *RegisterDebuggeeRequest, opts ...grpc.CallOption) (*RegisterDebuggeeResponse, error)
	// Returns the list of all active breakpoints for the debuggee.
	//
	// The breakpoint specification (`location`, `condition`, and `expressions`
	// fields) is semantically immutable, although the field values may
	// change. For example, an agent may update the location line number
	// to reflect the actual line where the breakpoint was set, but this
	// doesn't change the breakpoint semantics.
	//
	// This means that an agent does not need to check if a breakpoint has changed
	// when it encounters the same breakpoint on a successive call.
	// Moreover, an agent should remember the breakpoints that are completed
	// until the controller removes them from the active list to avoid
	// setting those breakpoints again.
	ListActiveBreakpoints(ctx context.Context, in *ListActiveBreakpointsRequest, opts ...grpc.CallOption) (*ListActiveBreakpointsResponse, error)
	// Updates the breakpoint state or mutable fields.
	// The entire Breakpoint message must be sent back to the controller service.
	//
	// Updates to active breakpoint fields are only allowed if the new value
	// does not change the breakpoint specification. Updates to the `location`,
	// `condition` and `expressions` fields should not alter the breakpoint
	// semantics. These may only make changes such as canonicalizing a value
	// or snapping the location to the correct line of code.
	UpdateActiveBreakpoint(ctx context.Context, in *UpdateActiveBreakpointRequest, opts ...grpc.CallOption) (*UpdateActiveBreakpointResponse, error)
}

type controller2Client struct {
	cc *grpc.ClientConn
}

func NewController2Client(cc *grpc.ClientConn) Controller2Client {
	return &controller2Client{cc}
}

func (c *controller2Client) RegisterDebuggee(ctx context.Context, in *RegisterDebuggeeRequest, opts ...grpc.CallOption) (*RegisterDebuggeeResponse, error) {
	out := new(RegisterDebuggeeResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Controller2/RegisterDebuggee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controller2Client) ListActiveBreakpoints(ctx context.Context, in *ListActiveBreakpointsRequest, opts ...grpc.CallOption) (*ListActiveBreakpointsResponse, error) {
	out := new(ListActiveBreakpointsResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Controller2/ListActiveBreakpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controller2Client) UpdateActiveBreakpoint(ctx context.Context, in *UpdateActiveBreakpointRequest, opts ...grpc.CallOption) (*UpdateActiveBreakpointResponse, error) {
	out := new(UpdateActiveBreakpointResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Controller2/UpdateActiveBreakpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Controller2Server is the server API for Controller2 service.
type Controller2Server interface {
	// Registers the debuggee with the controller service.
	//
	// All agents attached to the same application must call this method with
	// exactly the same request content to get back the same stable `debuggee_id`.
	// Agents should call this method again whenever `google.rpc.Code.NOT_FOUND`
	// is returned from any controller method.
	//
	// This protocol allows the controller service to disable debuggees, recover
	// from data loss, or change the `debuggee_id` format. Agents must handle
	// `debuggee_id` value changing upon re-registration.
	RegisterDebuggee(context.Context, *RegisterDebuggeeRequest) (*RegisterDebuggeeResponse, error)
	// Returns the list of all active breakpoints for the debuggee.
	//
	// The breakpoint specification (`location`, `condition`, and `expressions`
	// fields) is semantically immutable, although the field values may
	// change. For example, an agent may update the location line number
	// to reflect the actual line where the breakpoint was set, but this
	// doesn't change the breakpoint semantics.
	//
	// This means that an agent does not need to check if a breakpoint has changed
	// when it encounters the same breakpoint on a successive call.
	// Moreover, an agent should remember the breakpoints that are completed
	// until the controller removes them from the active list to avoid
	// setting those breakpoints again.
	ListActiveBreakpoints(context.Context, *ListActiveBreakpointsRequest) (*ListActiveBreakpointsResponse, error)
	// Updates the breakpoint state or mutable fields.
	// The entire Breakpoint message must be sent back to the controller service.
	//
	// Updates to active breakpoint fields are only allowed if the new value
	// does not change the breakpoint specification. Updates to the `location`,
	// `condition` and `expressions` fields should not alter the breakpoint
	// semantics. These may only make changes such as canonicalizing a value
	// or snapping the location to the correct line of code.
	UpdateActiveBreakpoint(context.Context, *UpdateActiveBreakpointRequest) (*UpdateActiveBreakpointResponse, error)
}

// UnimplementedController2Server can be embedded to have forward compatible implementations.
type UnimplementedController2Server struct {
}

func (*UnimplementedController2Server) RegisterDebuggee(ctx context.Context, req *RegisterDebuggeeRequest) (*RegisterDebuggeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDebuggee not implemented")
}
func (*UnimplementedController2Server) ListActiveBreakpoints(ctx context.Context, req *ListActiveBreakpointsRequest) (*ListActiveBreakpointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActiveBreakpoints not implemented")
}
func (*UnimplementedController2Server) UpdateActiveBreakpoint(ctx context.Context, req *UpdateActiveBreakpointRequest) (*UpdateActiveBreakpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateActiveBreakpoint not implemented")
}

func RegisterController2Server(s *grpc.Server, srv Controller2Server) {
	s.RegisterService(&_Controller2_serviceDesc, srv)
}

func _Controller2_RegisterDebuggee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDebuggeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Controller2Server).RegisterDebuggee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Controller2/RegisterDebuggee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Controller2Server).RegisterDebuggee(ctx, req.(*RegisterDebuggeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller2_ListActiveBreakpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActiveBreakpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Controller2Server).ListActiveBreakpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Controller2/ListActiveBreakpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Controller2Server).ListActiveBreakpoints(ctx, req.(*ListActiveBreakpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller2_UpdateActiveBreakpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateActiveBreakpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Controller2Server).UpdateActiveBreakpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Controller2/UpdateActiveBreakpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Controller2Server).UpdateActiveBreakpoint(ctx, req.(*UpdateActiveBreakpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Controller2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.clouddebugger.v2.Controller2",
	HandlerType: (*Controller2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDebuggee",
			Handler:    _Controller2_RegisterDebuggee_Handler,
		},
		{
			MethodName: "ListActiveBreakpoints",
			Handler:    _Controller2_ListActiveBreakpoints_Handler,
		},
		{
			MethodName: "UpdateActiveBreakpoint",
			Handler:    _Controller2_UpdateActiveBreakpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/clouddebugger/v2/controller.proto",
}
