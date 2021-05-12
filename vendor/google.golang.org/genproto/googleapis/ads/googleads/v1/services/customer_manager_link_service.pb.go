// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/customer_manager_link_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// Request message for
// [CustomerManagerLinkService.GetCustomerManagerLink][google.ads.googleads.v1.services.CustomerManagerLinkService.GetCustomerManagerLink].
type GetCustomerManagerLinkRequest struct {
	// The resource name of the CustomerManagerLink to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerManagerLinkRequest) Reset()         { *m = GetCustomerManagerLinkRequest{} }
func (m *GetCustomerManagerLinkRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerManagerLinkRequest) ProtoMessage()    {}
func (*GetCustomerManagerLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa2b7db8eabbebe, []int{0}
}

func (m *GetCustomerManagerLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerManagerLinkRequest.Unmarshal(m, b)
}
func (m *GetCustomerManagerLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerManagerLinkRequest.Marshal(b, m, deterministic)
}
func (m *GetCustomerManagerLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerManagerLinkRequest.Merge(m, src)
}
func (m *GetCustomerManagerLinkRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerManagerLinkRequest.Size(m)
}
func (m *GetCustomerManagerLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerManagerLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerManagerLinkRequest proto.InternalMessageInfo

func (m *GetCustomerManagerLinkRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [CustomerManagerLinkService.MutateCustomerManagerLink][google.ads.googleads.v1.services.CustomerManagerLinkService.MutateCustomerManagerLink].
type MutateCustomerManagerLinkRequest struct {
	// The ID of the customer whose customer manager links are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual customer manager links.
	Operations           []*CustomerManagerLinkOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *MutateCustomerManagerLinkRequest) Reset()         { *m = MutateCustomerManagerLinkRequest{} }
func (m *MutateCustomerManagerLinkRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerManagerLinkRequest) ProtoMessage()    {}
func (*MutateCustomerManagerLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa2b7db8eabbebe, []int{1}
}

func (m *MutateCustomerManagerLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerManagerLinkRequest.Unmarshal(m, b)
}
func (m *MutateCustomerManagerLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerManagerLinkRequest.Marshal(b, m, deterministic)
}
func (m *MutateCustomerManagerLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerManagerLinkRequest.Merge(m, src)
}
func (m *MutateCustomerManagerLinkRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerManagerLinkRequest.Size(m)
}
func (m *MutateCustomerManagerLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerManagerLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerManagerLinkRequest proto.InternalMessageInfo

func (m *MutateCustomerManagerLinkRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCustomerManagerLinkRequest) GetOperations() []*CustomerManagerLinkOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// Updates the status of a CustomerManagerLink.
// The following actions are possible:
// 1. Update operation with status ACTIVE accepts a pending invitation.
// 2. Update operation with status REFUSED declines a pending invitation.
// 3. Update operation with status INACTIVE terminates link to manager.
type CustomerManagerLinkOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CustomerManagerLinkOperation_Update
	Operation            isCustomerManagerLinkOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *CustomerManagerLinkOperation) Reset()         { *m = CustomerManagerLinkOperation{} }
func (m *CustomerManagerLinkOperation) String() string { return proto.CompactTextString(m) }
func (*CustomerManagerLinkOperation) ProtoMessage()    {}
func (*CustomerManagerLinkOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa2b7db8eabbebe, []int{2}
}

func (m *CustomerManagerLinkOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerManagerLinkOperation.Unmarshal(m, b)
}
func (m *CustomerManagerLinkOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerManagerLinkOperation.Marshal(b, m, deterministic)
}
func (m *CustomerManagerLinkOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerManagerLinkOperation.Merge(m, src)
}
func (m *CustomerManagerLinkOperation) XXX_Size() int {
	return xxx_messageInfo_CustomerManagerLinkOperation.Size(m)
}
func (m *CustomerManagerLinkOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerManagerLinkOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerManagerLinkOperation proto.InternalMessageInfo

func (m *CustomerManagerLinkOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCustomerManagerLinkOperation_Operation interface {
	isCustomerManagerLinkOperation_Operation()
}

type CustomerManagerLinkOperation_Update struct {
	Update *resources.CustomerManagerLink `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

func (*CustomerManagerLinkOperation_Update) isCustomerManagerLinkOperation_Operation() {}

func (m *CustomerManagerLinkOperation) GetOperation() isCustomerManagerLinkOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CustomerManagerLinkOperation) GetUpdate() *resources.CustomerManagerLink {
	if x, ok := m.GetOperation().(*CustomerManagerLinkOperation_Update); ok {
		return x.Update
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CustomerManagerLinkOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CustomerManagerLinkOperation_Update)(nil),
	}
}

// Response message for a CustomerManagerLink mutate.
type MutateCustomerManagerLinkResponse struct {
	// A result that identifies the resource affected by the mutate request.
	Results              []*MutateCustomerManagerLinkResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *MutateCustomerManagerLinkResponse) Reset()         { *m = MutateCustomerManagerLinkResponse{} }
func (m *MutateCustomerManagerLinkResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerManagerLinkResponse) ProtoMessage()    {}
func (*MutateCustomerManagerLinkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa2b7db8eabbebe, []int{3}
}

func (m *MutateCustomerManagerLinkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerManagerLinkResponse.Unmarshal(m, b)
}
func (m *MutateCustomerManagerLinkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerManagerLinkResponse.Marshal(b, m, deterministic)
}
func (m *MutateCustomerManagerLinkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerManagerLinkResponse.Merge(m, src)
}
func (m *MutateCustomerManagerLinkResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerManagerLinkResponse.Size(m)
}
func (m *MutateCustomerManagerLinkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerManagerLinkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerManagerLinkResponse proto.InternalMessageInfo

func (m *MutateCustomerManagerLinkResponse) GetResults() []*MutateCustomerManagerLinkResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the customer manager link mutate.
type MutateCustomerManagerLinkResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCustomerManagerLinkResult) Reset()         { *m = MutateCustomerManagerLinkResult{} }
func (m *MutateCustomerManagerLinkResult) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerManagerLinkResult) ProtoMessage()    {}
func (*MutateCustomerManagerLinkResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa2b7db8eabbebe, []int{4}
}

func (m *MutateCustomerManagerLinkResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerManagerLinkResult.Unmarshal(m, b)
}
func (m *MutateCustomerManagerLinkResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerManagerLinkResult.Marshal(b, m, deterministic)
}
func (m *MutateCustomerManagerLinkResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerManagerLinkResult.Merge(m, src)
}
func (m *MutateCustomerManagerLinkResult) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerManagerLinkResult.Size(m)
}
func (m *MutateCustomerManagerLinkResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerManagerLinkResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerManagerLinkResult proto.InternalMessageInfo

func (m *MutateCustomerManagerLinkResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCustomerManagerLinkRequest)(nil), "google.ads.googleads.v1.services.GetCustomerManagerLinkRequest")
	proto.RegisterType((*MutateCustomerManagerLinkRequest)(nil), "google.ads.googleads.v1.services.MutateCustomerManagerLinkRequest")
	proto.RegisterType((*CustomerManagerLinkOperation)(nil), "google.ads.googleads.v1.services.CustomerManagerLinkOperation")
	proto.RegisterType((*MutateCustomerManagerLinkResponse)(nil), "google.ads.googleads.v1.services.MutateCustomerManagerLinkResponse")
	proto.RegisterType((*MutateCustomerManagerLinkResult)(nil), "google.ads.googleads.v1.services.MutateCustomerManagerLinkResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/customer_manager_link_service.proto", fileDescriptor_aaa2b7db8eabbebe)
}

var fileDescriptor_aaa2b7db8eabbebe = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0x36, 0xa9, 0xac, 0xec, 0x54, 0x2f, 0x73, 0x90, 0x5a, 0x56, 0x1a, 0xe3, 0x1e, 0x4a, 0x0f,
	0x13, 0x5a, 0x61, 0xd1, 0x2c, 0xad, 0xa4, 0x5b, 0xb6, 0x2b, 0x58, 0x2d, 0x11, 0x7a, 0xd0, 0x62,
	0x99, 0x6d, 0x66, 0x43, 0x68, 0x92, 0x89, 0x99, 0x49, 0x2f, 0xcb, 0x82, 0x78, 0xf0, 0x0f, 0xf8,
	0x0f, 0xf4, 0xe6, 0x4f, 0xf0, 0x27, 0x78, 0x13, 0xcf, 0xde, 0x3c, 0xed, 0xaf, 0x90, 0x64, 0x32,
	0xb1, 0x42, 0xd3, 0xca, 0xee, 0xed, 0x65, 0xf2, 0xde, 0xf7, 0xbe, 0xef, 0x9b, 0xf7, 0x06, 0x0c,
	0x5c, 0x4a, 0x5d, 0x9f, 0x18, 0xd8, 0x61, 0x86, 0x08, 0xd3, 0x68, 0xd9, 0x36, 0x18, 0x89, 0x97,
	0xde, 0x9c, 0x30, 0x63, 0x9e, 0x30, 0x4e, 0x03, 0x12, 0xcf, 0x02, 0x1c, 0x62, 0x97, 0xc4, 0x33,
	0xdf, 0x0b, 0x17, 0xb3, 0xfc, 0x37, 0x8a, 0x62, 0xca, 0x29, 0xd4, 0x44, 0x29, 0xc2, 0x0e, 0x43,
	0x05, 0x0a, 0x5a, 0xb6, 0x91, 0x44, 0xa9, 0x77, 0xcb, 0xfa, 0xc4, 0x84, 0xd1, 0x24, 0x2e, 0x6d,
	0x24, 0x1a, 0xd4, 0xf7, 0x64, 0x79, 0xe4, 0x19, 0x38, 0x0c, 0x29, 0xc7, 0xdc, 0xa3, 0x21, 0xcb,
	0xff, 0xe6, 0xed, 0x8d, 0xec, 0xeb, 0x34, 0x39, 0x33, 0xce, 0x3c, 0xe2, 0x3b, 0xb3, 0x00, 0xb3,
	0xbc, 0x5e, 0x1f, 0x80, 0xfb, 0x43, 0xc2, 0x8f, 0xf2, 0x0e, 0x23, 0xd1, 0xe0, 0xb9, 0x17, 0x2e,
	0x6c, 0xf2, 0x2e, 0x21, 0x8c, 0xc3, 0x87, 0xe0, 0x8e, 0x64, 0x32, 0x0b, 0x71, 0x40, 0x6a, 0x8a,
	0xa6, 0x34, 0x77, 0xed, 0xdb, 0xf2, 0xf0, 0x05, 0x0e, 0x88, 0xfe, 0x45, 0x01, 0xda, 0x28, 0xe1,
	0x98, 0x93, 0x0d, 0x48, 0x0d, 0x50, 0x2d, 0x94, 0x78, 0x4e, 0x8e, 0x03, 0xe4, 0xd1, 0x33, 0x07,
	0xbe, 0x05, 0x80, 0x46, 0x24, 0x16, 0x0a, 0x6a, 0xaa, 0x56, 0x69, 0x56, 0x3b, 0x3d, 0xb4, 0xcd,
	0x41, 0xb4, 0xa6, 0xe5, 0x4b, 0x09, 0x63, 0xaf, 0x20, 0xea, 0xdf, 0x14, 0xb0, 0xb7, 0x29, 0x19,
	0x1e, 0x82, 0x6a, 0x12, 0x39, 0x98, 0x93, 0xcc, 0xa1, 0xda, 0x4d, 0x4d, 0x69, 0x56, 0x3b, 0x75,
	0xc9, 0x40, 0x9a, 0x88, 0x8e, 0x53, 0x13, 0x47, 0x98, 0x2d, 0x6c, 0x20, 0xd2, 0xd3, 0x18, 0x8e,
	0xc1, 0x8e, 0xf8, 0xaa, 0xa9, 0x59, 0xdd, 0x41, 0x29, 0xf3, 0xe2, 0x66, 0xd7, 0x51, 0x3f, 0xb9,
	0x61, 0xe7, 0x38, 0xfd, 0x2a, 0xd8, 0x2d, 0xd8, 0xeb, 0xef, 0x15, 0xf0, 0x60, 0x83, 0xc5, 0x2c,
	0xa2, 0x21, 0x23, 0xf0, 0x0d, 0xb8, 0x15, 0x13, 0x96, 0xf8, 0x9c, 0xd5, 0x94, 0xcc, 0x3f, 0x6b,
	0xbb, 0x7f, 0x9b, 0x50, 0x13, 0x9f, 0xdb, 0x12, 0x51, 0x3f, 0x06, 0x8d, 0x2d, 0xb9, 0xff, 0x35,
	0x2d, 0x9d, 0x5f, 0x15, 0x50, 0x5f, 0x03, 0xf1, 0x4a, 0x10, 0x82, 0x3f, 0x14, 0x70, 0x77, 0xfd,
	0x4c, 0xc2, 0xa7, 0xdb, 0xd5, 0x6c, 0x9c, 0xe6, 0xfa, 0x15, 0x2f, 0x45, 0xef, 0x7d, 0xf8, 0xf9,
	0xfb, 0x93, 0xfa, 0x18, 0x1e, 0xa4, 0x9b, 0x79, 0xfe, 0x8f, 0xc4, 0xae, 0x9c, 0x60, 0x66, 0xb4,
	0x8a, 0x55, 0x5d, 0xa9, 0x65, 0x46, 0xeb, 0x02, 0x5e, 0x2a, 0xe0, 0x5e, 0xa9, 0x77, 0xb0, 0x7f,
	0xad, 0x4b, 0x12, 0xca, 0x8e, 0xae, 0x77, 0xd1, 0xd9, 0xf8, 0xe8, 0x83, 0x4c, 0x66, 0x4f, 0x7f,
	0x92, 0xca, 0xfc, 0xab, 0xeb, 0x7c, 0x65, 0x6f, 0xbb, 0xad, 0x8b, 0xb5, 0x2a, 0xcd, 0x20, 0x43,
	0x37, 0x95, 0x56, 0xff, 0xa3, 0x0a, 0xf6, 0xe7, 0x34, 0xd8, 0x4a, 0xa8, 0xdf, 0x28, 0x9f, 0x82,
	0x71, 0xba, 0x6c, 0x63, 0xe5, 0xf5, 0x49, 0x0e, 0xe2, 0x52, 0x1f, 0x87, 0x2e, 0xa2, 0xb1, 0x6b,
	0xb8, 0x24, 0xcc, 0x56, 0x51, 0x3e, 0x97, 0x91, 0xc7, 0xca, 0x5f, 0xe9, 0x43, 0x19, 0x7c, 0x56,
	0x2b, 0x43, 0xcb, 0xfa, 0xaa, 0x6a, 0x43, 0x01, 0x68, 0x39, 0x0c, 0x89, 0x30, 0x8d, 0x26, 0x6d,
	0x94, 0x37, 0x66, 0xdf, 0x65, 0xca, 0xd4, 0x72, 0xd8, 0xb4, 0x48, 0x99, 0x4e, 0xda, 0x53, 0x99,
	0x72, 0xa9, 0xee, 0x8b, 0x73, 0xd3, 0xb4, 0x1c, 0x66, 0x9a, 0x45, 0x92, 0x69, 0x4e, 0xda, 0xa6,
	0x29, 0xd3, 0x4e, 0x77, 0x32, 0x9e, 0x8f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xb5, 0x81,
	0xa2, 0x4c, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerManagerLinkServiceClient is the client API for CustomerManagerLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerManagerLinkServiceClient interface {
	// Returns the requested CustomerManagerLink in full detail.
	GetCustomerManagerLink(ctx context.Context, in *GetCustomerManagerLinkRequest, opts ...grpc.CallOption) (*resources.CustomerManagerLink, error)
	// Creates or updates customer manager links. Operation statuses are returned.
	MutateCustomerManagerLink(ctx context.Context, in *MutateCustomerManagerLinkRequest, opts ...grpc.CallOption) (*MutateCustomerManagerLinkResponse, error)
}

type customerManagerLinkServiceClient struct {
	cc *grpc.ClientConn
}

func NewCustomerManagerLinkServiceClient(cc *grpc.ClientConn) CustomerManagerLinkServiceClient {
	return &customerManagerLinkServiceClient{cc}
}

func (c *customerManagerLinkServiceClient) GetCustomerManagerLink(ctx context.Context, in *GetCustomerManagerLinkRequest, opts ...grpc.CallOption) (*resources.CustomerManagerLink, error) {
	out := new(resources.CustomerManagerLink)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CustomerManagerLinkService/GetCustomerManagerLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerManagerLinkServiceClient) MutateCustomerManagerLink(ctx context.Context, in *MutateCustomerManagerLinkRequest, opts ...grpc.CallOption) (*MutateCustomerManagerLinkResponse, error) {
	out := new(MutateCustomerManagerLinkResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CustomerManagerLinkService/MutateCustomerManagerLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerManagerLinkServiceServer is the server API for CustomerManagerLinkService service.
type CustomerManagerLinkServiceServer interface {
	// Returns the requested CustomerManagerLink in full detail.
	GetCustomerManagerLink(context.Context, *GetCustomerManagerLinkRequest) (*resources.CustomerManagerLink, error)
	// Creates or updates customer manager links. Operation statuses are returned.
	MutateCustomerManagerLink(context.Context, *MutateCustomerManagerLinkRequest) (*MutateCustomerManagerLinkResponse, error)
}

// UnimplementedCustomerManagerLinkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerManagerLinkServiceServer struct {
}

func (*UnimplementedCustomerManagerLinkServiceServer) GetCustomerManagerLink(ctx context.Context, req *GetCustomerManagerLinkRequest) (*resources.CustomerManagerLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerManagerLink not implemented")
}
func (*UnimplementedCustomerManagerLinkServiceServer) MutateCustomerManagerLink(ctx context.Context, req *MutateCustomerManagerLinkRequest) (*MutateCustomerManagerLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerManagerLink not implemented")
}

func RegisterCustomerManagerLinkServiceServer(s *grpc.Server, srv CustomerManagerLinkServiceServer) {
	s.RegisterService(&_CustomerManagerLinkService_serviceDesc, srv)
}

func _CustomerManagerLinkService_GetCustomerManagerLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerManagerLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerManagerLinkServiceServer).GetCustomerManagerLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CustomerManagerLinkService/GetCustomerManagerLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerManagerLinkServiceServer).GetCustomerManagerLink(ctx, req.(*GetCustomerManagerLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerManagerLinkService_MutateCustomerManagerLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerManagerLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerManagerLinkServiceServer).MutateCustomerManagerLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CustomerManagerLinkService/MutateCustomerManagerLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerManagerLinkServiceServer).MutateCustomerManagerLink(ctx, req.(*MutateCustomerManagerLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerManagerLinkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CustomerManagerLinkService",
	HandlerType: (*CustomerManagerLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerManagerLink",
			Handler:    _CustomerManagerLinkService_GetCustomerManagerLink_Handler,
		},
		{
			MethodName: "MutateCustomerManagerLink",
			Handler:    _CustomerManagerLinkService_MutateCustomerManagerLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/customer_manager_link_service.proto",
}
