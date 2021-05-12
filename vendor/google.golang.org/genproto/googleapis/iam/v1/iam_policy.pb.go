// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/iam/v1/iam_policy.proto

package iam

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Request message for `SetIamPolicy` method.
type SetIamPolicyRequest struct {
	// REQUIRED: The resource for which the policy is being specified.
	// See the operation documentation for the appropriate value for this field.
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// REQUIRED: The complete policy to be applied to the `resource`. The size of
	// the policy is limited to a few 10s of KB. An empty policy is a
	// valid policy but certain Cloud Platform services (such as Projects)
	// might reject them.
	Policy               *Policy  `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetIamPolicyRequest) Reset()         { *m = SetIamPolicyRequest{} }
func (m *SetIamPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*SetIamPolicyRequest) ProtoMessage()    {}
func (*SetIamPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2728eb97d748a32, []int{0}
}

func (m *SetIamPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetIamPolicyRequest.Unmarshal(m, b)
}
func (m *SetIamPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetIamPolicyRequest.Marshal(b, m, deterministic)
}
func (m *SetIamPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetIamPolicyRequest.Merge(m, src)
}
func (m *SetIamPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_SetIamPolicyRequest.Size(m)
}
func (m *SetIamPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetIamPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetIamPolicyRequest proto.InternalMessageInfo

func (m *SetIamPolicyRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *SetIamPolicyRequest) GetPolicy() *Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

// Request message for `GetIamPolicy` method.
type GetIamPolicyRequest struct {
	// REQUIRED: The resource for which the policy is being requested.
	// See the operation documentation for the appropriate value for this field.
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// OPTIONAL: A `GetPolicyOptions` object for specifying options to
	// `GetIamPolicy`. This field is only used by Cloud IAM.
	Options              *GetPolicyOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetIamPolicyRequest) Reset()         { *m = GetIamPolicyRequest{} }
func (m *GetIamPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*GetIamPolicyRequest) ProtoMessage()    {}
func (*GetIamPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2728eb97d748a32, []int{1}
}

func (m *GetIamPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIamPolicyRequest.Unmarshal(m, b)
}
func (m *GetIamPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIamPolicyRequest.Marshal(b, m, deterministic)
}
func (m *GetIamPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIamPolicyRequest.Merge(m, src)
}
func (m *GetIamPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_GetIamPolicyRequest.Size(m)
}
func (m *GetIamPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIamPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetIamPolicyRequest proto.InternalMessageInfo

func (m *GetIamPolicyRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *GetIamPolicyRequest) GetOptions() *GetPolicyOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

// Request message for `TestIamPermissions` method.
type TestIamPermissionsRequest struct {
	// REQUIRED: The resource for which the policy detail is being requested.
	// See the operation documentation for the appropriate value for this field.
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// The set of permissions to check for the `resource`. Permissions with
	// wildcards (such as '*' or 'storage.*') are not allowed. For more
	// information see
	// [IAM Overview](https://cloud.google.com/iam/docs/overview#permissions).
	Permissions          []string `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestIamPermissionsRequest) Reset()         { *m = TestIamPermissionsRequest{} }
func (m *TestIamPermissionsRequest) String() string { return proto.CompactTextString(m) }
func (*TestIamPermissionsRequest) ProtoMessage()    {}
func (*TestIamPermissionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2728eb97d748a32, []int{2}
}

func (m *TestIamPermissionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestIamPermissionsRequest.Unmarshal(m, b)
}
func (m *TestIamPermissionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestIamPermissionsRequest.Marshal(b, m, deterministic)
}
func (m *TestIamPermissionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestIamPermissionsRequest.Merge(m, src)
}
func (m *TestIamPermissionsRequest) XXX_Size() int {
	return xxx_messageInfo_TestIamPermissionsRequest.Size(m)
}
func (m *TestIamPermissionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestIamPermissionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestIamPermissionsRequest proto.InternalMessageInfo

func (m *TestIamPermissionsRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *TestIamPermissionsRequest) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

// Response message for `TestIamPermissions` method.
type TestIamPermissionsResponse struct {
	// A subset of `TestPermissionsRequest.permissions` that the caller is
	// allowed.
	Permissions          []string `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestIamPermissionsResponse) Reset()         { *m = TestIamPermissionsResponse{} }
func (m *TestIamPermissionsResponse) String() string { return proto.CompactTextString(m) }
func (*TestIamPermissionsResponse) ProtoMessage()    {}
func (*TestIamPermissionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2728eb97d748a32, []int{3}
}

func (m *TestIamPermissionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestIamPermissionsResponse.Unmarshal(m, b)
}
func (m *TestIamPermissionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestIamPermissionsResponse.Marshal(b, m, deterministic)
}
func (m *TestIamPermissionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestIamPermissionsResponse.Merge(m, src)
}
func (m *TestIamPermissionsResponse) XXX_Size() int {
	return xxx_messageInfo_TestIamPermissionsResponse.Size(m)
}
func (m *TestIamPermissionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestIamPermissionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestIamPermissionsResponse proto.InternalMessageInfo

func (m *TestIamPermissionsResponse) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func init() {
	proto.RegisterType((*SetIamPolicyRequest)(nil), "google.iam.v1.SetIamPolicyRequest")
	proto.RegisterType((*GetIamPolicyRequest)(nil), "google.iam.v1.GetIamPolicyRequest")
	proto.RegisterType((*TestIamPermissionsRequest)(nil), "google.iam.v1.TestIamPermissionsRequest")
	proto.RegisterType((*TestIamPermissionsResponse)(nil), "google.iam.v1.TestIamPermissionsResponse")
}

func init() { proto.RegisterFile("google/iam/v1/iam_policy.proto", fileDescriptor_d2728eb97d748a32) }

var fileDescriptor_d2728eb97d748a32 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x8a, 0xd3, 0x40,
	0x18, 0x66, 0x52, 0x58, 0xed, 0xac, 0x0a, 0xa6, 0x88, 0xdd, 0xac, 0x74, 0x4b, 0x74, 0xa1, 0x0d,
	0xec, 0xc4, 0xd6, 0x93, 0x15, 0x85, 0xd4, 0x43, 0xe8, 0x41, 0x2c, 0x55, 0xf6, 0x20, 0x85, 0x65,
	0x36, 0x3b, 0xc6, 0x81, 0x4c, 0x66, 0xcc, 0x4c, 0x2b, 0x22, 0x5e, 0x3c, 0xf8, 0x02, 0xde, 0x7c,
	0x04, 0xcf, 0x3e, 0xc5, 0x5e, 0x7d, 0x81, 0x3d, 0xf8, 0x10, 0xe2, 0x49, 0x92, 0x99, 0x6e, 0x93,
	0xb6, 0x8a, 0xca, 0x9e, 0x0a, 0xff, 0xf7, 0xfd, 0xdf, 0xf7, 0x7f, 0xff, 0xdf, 0x09, 0x6c, 0xc5,
	0x9c, 0xc7, 0x09, 0xf1, 0x29, 0x66, 0xfe, 0xbc, 0x97, 0xff, 0x1c, 0x09, 0x9e, 0xd0, 0xe8, 0x2d,
	0x12, 0x19, 0x57, 0xdc, 0xbe, 0xaa, 0x71, 0x44, 0x31, 0x43, 0xf3, 0x9e, 0xb3, 0x5b, 0xa5, 0x73,
	0xa1, 0x28, 0x4f, 0xa5, 0xe6, 0x3a, 0x4e, 0x15, 0x2c, 0xeb, 0x38, 0xb7, 0x0c, 0x86, 0x05, 0xf5,
	0x71, 0x9a, 0x72, 0x85, 0xcb, 0x9d, 0x37, 0x4b, 0x68, 0x94, 0x50, 0x92, 0x2a, 0x03, 0xec, 0x95,
	0x80, 0x97, 0x94, 0x24, 0x27, 0x47, 0xc7, 0xe4, 0x15, 0x9e, 0x53, 0x9e, 0x19, 0xc2, 0x4e, 0x89,
	0x90, 0x11, 0xc9, 0x67, 0x59, 0x44, 0x34, 0xe4, 0x0a, 0xd8, 0x78, 0x46, 0xd4, 0x08, 0xb3, 0x71,
	0x31, 0xc8, 0x84, 0xbc, 0x9e, 0x11, 0xa9, 0xec, 0x7d, 0x78, 0x79, 0x41, 0x6c, 0x82, 0x36, 0xe8,
	0xd4, 0x87, 0xf5, 0xb3, 0xc0, 0xfa, 0x19, 0xd4, 0x20, 0xf0, 0x26, 0xe7, 0x90, 0xdd, 0x87, 0x5b,
	0x3a, 0x40, 0xd3, 0x6a, 0x83, 0xce, 0x76, 0xff, 0x06, 0xaa, 0x6c, 0x02, 0x69, 0xd1, 0x61, 0xed,
	0x2c, 0xb0, 0x26, 0x86, 0xe9, 0xbe, 0x81, 0x8d, 0xf0, 0xff, 0x1d, 0xef, 0xc3, 0x4b, 0x66, 0x9f,
	0xc6, 0x72, 0x6f, 0xc5, 0x32, 0x24, 0x4a, 0x0b, 0x3f, 0xd5, 0xb4, 0xc9, 0x82, 0xef, 0x52, 0xb8,
	0xf3, 0x9c, 0xc8, 0xc2, 0x99, 0x64, 0x8c, 0x4a, 0x59, 0xc0, 0xff, 0x66, 0xbf, 0x0f, 0xb7, 0xc5,
	0xb2, 0xb9, 0x69, 0xb5, 0x6b, 0x9d, 0xba, 0x8e, 0x57, 0xae, 0xbb, 0x8f, 0xa0, 0xb3, 0xc9, 0x4a,
	0x0a, 0x9e, 0x4a, 0x62, 0xb7, 0xab, 0x22, 0x20, 0x17, 0xa9, 0xf4, 0xf7, 0xbf, 0xd6, 0x60, 0x7d,
	0x14, 0x3c, 0xd1, 0x41, 0x6c, 0x05, 0xaf, 0x94, 0x6f, 0x64, 0xbb, 0x2b, 0x91, 0x37, 0x1c, 0xd0,
	0xd9, 0x7c, 0x09, 0xb7, 0xfb, 0xe1, 0xdb, 0xf7, 0x4f, 0xd6, 0x6d, 0xb7, 0x95, 0xff, 0xf7, 0xde,
	0x2d, 0x62, 0x3d, 0xf4, 0xbc, 0xf7, 0x03, 0x59, 0x52, 0x19, 0x00, 0x2f, 0x77, 0x0d, 0xff, 0xe4,
	0x1a, 0x5e, 0x88, 0x6b, 0xbc, 0xe2, 0xfa, 0x19, 0x40, 0x7b, 0x7d, 0x75, 0x76, 0x67, 0x45, 0xf8,
	0xb7, 0x87, 0x74, 0xba, 0x7f, 0xc1, 0xd4, 0x77, 0x70, 0xfd, 0x62, 0xac, 0xae, 0x7b, 0x67, 0x7d,
	0x2c, 0xb5, 0xd6, 0x35, 0x00, 0x9e, 0xd3, 0x3a, 0x0d, 0x76, 0x29, 0x66, 0x07, 0x8c, 0x28, 0x7c,
	0x80, 0x05, 0x35, 0x56, 0x58, 0x50, 0x89, 0x22, 0xce, 0x86, 0x1f, 0x01, 0xbc, 0x1e, 0x71, 0x56,
	0x9d, 0x60, 0x78, 0xed, 0x3c, 0xe0, 0x38, 0x7f, 0x72, 0x63, 0xf0, 0xe2, 0xae, 0x21, 0xc4, 0x3c,
	0xc1, 0x69, 0x8c, 0x78, 0x16, 0xfb, 0x31, 0x49, 0x8b, 0x07, 0xe9, 0x2f, 0x25, 0xcd, 0x27, 0xe2,
	0x01, 0xc5, 0xec, 0x07, 0x00, 0x5f, 0xac, 0x46, 0xa8, 0xbb, 0x1e, 0x27, 0x7c, 0x76, 0x82, 0x46,
	0x98, 0xa1, 0xc3, 0xde, 0xe9, 0xa2, 0x3a, 0x2d, 0xaa, 0xd3, 0x11, 0x66, 0xd3, 0xc3, 0xde, 0xf1,
	0x56, 0xa1, 0x75, 0xef, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0x24, 0xb5, 0x51, 0xb9, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IAMPolicyClient is the client API for IAMPolicy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IAMPolicyClient interface {
	// Sets the access control policy on the specified resource. Replaces any
	// existing policy.
	SetIamPolicy(ctx context.Context, in *SetIamPolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	// Gets the access control policy for a resource.
	// Returns an empty policy if the resource exists and does not have a policy
	// set.
	GetIamPolicy(ctx context.Context, in *GetIamPolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	// Returns permissions that a caller has on the specified resource.
	// If the resource does not exist, this will return an empty set of
	// permissions, not a NOT_FOUND error.
	//
	// Note: This operation is designed to be used for building permission-aware
	// UIs and command-line tools, not for authorization checking. This operation
	// may "fail open" without warning.
	TestIamPermissions(ctx context.Context, in *TestIamPermissionsRequest, opts ...grpc.CallOption) (*TestIamPermissionsResponse, error)
}

type iAMPolicyClient struct {
	cc *grpc.ClientConn
}

func NewIAMPolicyClient(cc *grpc.ClientConn) IAMPolicyClient {
	return &iAMPolicyClient{cc}
}

func (c *iAMPolicyClient) SetIamPolicy(ctx context.Context, in *SetIamPolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, "/google.iam.v1.IAMPolicy/SetIamPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMPolicyClient) GetIamPolicy(ctx context.Context, in *GetIamPolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, "/google.iam.v1.IAMPolicy/GetIamPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMPolicyClient) TestIamPermissions(ctx context.Context, in *TestIamPermissionsRequest, opts ...grpc.CallOption) (*TestIamPermissionsResponse, error) {
	out := new(TestIamPermissionsResponse)
	err := c.cc.Invoke(ctx, "/google.iam.v1.IAMPolicy/TestIamPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IAMPolicyServer is the server API for IAMPolicy service.
type IAMPolicyServer interface {
	// Sets the access control policy on the specified resource. Replaces any
	// existing policy.
	SetIamPolicy(context.Context, *SetIamPolicyRequest) (*Policy, error)
	// Gets the access control policy for a resource.
	// Returns an empty policy if the resource exists and does not have a policy
	// set.
	GetIamPolicy(context.Context, *GetIamPolicyRequest) (*Policy, error)
	// Returns permissions that a caller has on the specified resource.
	// If the resource does not exist, this will return an empty set of
	// permissions, not a NOT_FOUND error.
	//
	// Note: This operation is designed to be used for building permission-aware
	// UIs and command-line tools, not for authorization checking. This operation
	// may "fail open" without warning.
	TestIamPermissions(context.Context, *TestIamPermissionsRequest) (*TestIamPermissionsResponse, error)
}

// UnimplementedIAMPolicyServer can be embedded to have forward compatible implementations.
type UnimplementedIAMPolicyServer struct {
}

func (*UnimplementedIAMPolicyServer) SetIamPolicy(ctx context.Context, req *SetIamPolicyRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIamPolicy not implemented")
}
func (*UnimplementedIAMPolicyServer) GetIamPolicy(ctx context.Context, req *GetIamPolicyRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIamPolicy not implemented")
}
func (*UnimplementedIAMPolicyServer) TestIamPermissions(ctx context.Context, req *TestIamPermissionsRequest) (*TestIamPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestIamPermissions not implemented")
}

func RegisterIAMPolicyServer(s *grpc.Server, srv IAMPolicyServer) {
	s.RegisterService(&_IAMPolicy_serviceDesc, srv)
}

func _IAMPolicy_SetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMPolicyServer).SetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.v1.IAMPolicy/SetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMPolicyServer).SetIamPolicy(ctx, req.(*SetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMPolicy_GetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMPolicyServer).GetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.v1.IAMPolicy/GetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMPolicyServer).GetIamPolicy(ctx, req.(*GetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMPolicy_TestIamPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestIamPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMPolicyServer).TestIamPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.v1.IAMPolicy/TestIamPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMPolicyServer).TestIamPermissions(ctx, req.(*TestIamPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IAMPolicy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.iam.v1.IAMPolicy",
	HandlerType: (*IAMPolicyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetIamPolicy",
			Handler:    _IAMPolicy_SetIamPolicy_Handler,
		},
		{
			MethodName: "GetIamPolicy",
			Handler:    _IAMPolicy_GetIamPolicy_Handler,
		},
		{
			MethodName: "TestIamPermissions",
			Handler:    _IAMPolicy_TestIamPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/iam/v1/iam_policy.proto",
}
