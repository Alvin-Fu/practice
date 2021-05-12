// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/talent/v4beta1/company_service.proto

package talent

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// The Request of the CreateCompany method.
type CreateCompanyRequest struct {
	// Required. Resource name of the tenant under which the company is created.
	//
	// The format is "projects/{project_id}/tenants/{tenant_id}", for example,
	// "projects/foo/tenant/bar". If tenant id is unspecified, a default tenant
	// is created, for example, "projects/foo".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The company to be created.
	Company              *Company `protobuf:"bytes,2,opt,name=company,proto3" json:"company,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCompanyRequest) Reset()         { *m = CreateCompanyRequest{} }
func (m *CreateCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCompanyRequest) ProtoMessage()    {}
func (*CreateCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c499e6cf60b8944, []int{0}
}

func (m *CreateCompanyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCompanyRequest.Unmarshal(m, b)
}
func (m *CreateCompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCompanyRequest.Marshal(b, m, deterministic)
}
func (m *CreateCompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCompanyRequest.Merge(m, src)
}
func (m *CreateCompanyRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCompanyRequest.Size(m)
}
func (m *CreateCompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCompanyRequest proto.InternalMessageInfo

func (m *CreateCompanyRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateCompanyRequest) GetCompany() *Company {
	if m != nil {
		return m.Company
	}
	return nil
}

// Request for getting a company by name.
type GetCompanyRequest struct {
	// Required. The resource name of the company to be retrieved.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for
	// example, "projects/api-test-project/tenants/foo/companies/bar".
	//
	// If tenant id is unspecified, the default tenant is used, for
	// example, "projects/api-test-project/companies/bar".
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCompanyRequest) Reset()         { *m = GetCompanyRequest{} }
func (m *GetCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*GetCompanyRequest) ProtoMessage()    {}
func (*GetCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c499e6cf60b8944, []int{1}
}

func (m *GetCompanyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCompanyRequest.Unmarshal(m, b)
}
func (m *GetCompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCompanyRequest.Marshal(b, m, deterministic)
}
func (m *GetCompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCompanyRequest.Merge(m, src)
}
func (m *GetCompanyRequest) XXX_Size() int {
	return xxx_messageInfo_GetCompanyRequest.Size(m)
}
func (m *GetCompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCompanyRequest proto.InternalMessageInfo

func (m *GetCompanyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request for updating a specified company.
type UpdateCompanyRequest struct {
	// Required. The company resource to replace the current resource in the
	// system.
	Company *Company `protobuf:"bytes,1,opt,name=company,proto3" json:"company,omitempty"`
	// Strongly recommended for the best service experience.
	//
	// If
	// [update_mask][google.cloud.talent.v4beta1.UpdateCompanyRequest.update_mask]
	// is provided, only the specified fields in
	// [company][google.cloud.talent.v4beta1.UpdateCompanyRequest.company] are
	// updated. Otherwise all the fields are updated.
	//
	// A field mask to specify the company fields to be updated. Only
	// top level fields of [Company][google.cloud.talent.v4beta1.Company] are
	// supported.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateCompanyRequest) Reset()         { *m = UpdateCompanyRequest{} }
func (m *UpdateCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCompanyRequest) ProtoMessage()    {}
func (*UpdateCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c499e6cf60b8944, []int{2}
}

func (m *UpdateCompanyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCompanyRequest.Unmarshal(m, b)
}
func (m *UpdateCompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCompanyRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCompanyRequest.Merge(m, src)
}
func (m *UpdateCompanyRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCompanyRequest.Size(m)
}
func (m *UpdateCompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCompanyRequest proto.InternalMessageInfo

func (m *UpdateCompanyRequest) GetCompany() *Company {
	if m != nil {
		return m.Company
	}
	return nil
}

func (m *UpdateCompanyRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// Request to delete a company.
type DeleteCompanyRequest struct {
	// Required. The resource name of the company to be deleted.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for
	// example, "projects/foo/tenants/bar/companies/baz".
	//
	// If tenant id is unspecified, the default tenant is used, for
	// example, "projects/foo/companies/bar".
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCompanyRequest) Reset()         { *m = DeleteCompanyRequest{} }
func (m *DeleteCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCompanyRequest) ProtoMessage()    {}
func (*DeleteCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c499e6cf60b8944, []int{3}
}

func (m *DeleteCompanyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCompanyRequest.Unmarshal(m, b)
}
func (m *DeleteCompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCompanyRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCompanyRequest.Merge(m, src)
}
func (m *DeleteCompanyRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCompanyRequest.Size(m)
}
func (m *DeleteCompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCompanyRequest proto.InternalMessageInfo

func (m *DeleteCompanyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// List companies for which the client has ACL visibility.
type ListCompaniesRequest struct {
	// Required. Resource name of the tenant under which the company is created.
	//
	// The format is "projects/{project_id}/tenants/{tenant_id}", for example,
	// "projects/foo/tenant/bar".
	//
	// If tenant id is unspecified, the default tenant will be used, for
	// example, "projects/foo".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The starting indicator from which to return results.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The maximum number of companies to be returned, at most 100.
	// Default is 100 if a non-positive number is provided.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Set to true if the companies requested must have open jobs.
	//
	// Defaults to false.
	//
	// If true, at most
	// [page_size][google.cloud.talent.v4beta1.ListCompaniesRequest.page_size] of
	// companies are fetched, among which only those with open jobs are returned.
	RequireOpenJobs      bool     `protobuf:"varint,4,opt,name=require_open_jobs,json=requireOpenJobs,proto3" json:"require_open_jobs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCompaniesRequest) Reset()         { *m = ListCompaniesRequest{} }
func (m *ListCompaniesRequest) String() string { return proto.CompactTextString(m) }
func (*ListCompaniesRequest) ProtoMessage()    {}
func (*ListCompaniesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c499e6cf60b8944, []int{4}
}

func (m *ListCompaniesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCompaniesRequest.Unmarshal(m, b)
}
func (m *ListCompaniesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCompaniesRequest.Marshal(b, m, deterministic)
}
func (m *ListCompaniesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCompaniesRequest.Merge(m, src)
}
func (m *ListCompaniesRequest) XXX_Size() int {
	return xxx_messageInfo_ListCompaniesRequest.Size(m)
}
func (m *ListCompaniesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCompaniesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCompaniesRequest proto.InternalMessageInfo

func (m *ListCompaniesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListCompaniesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListCompaniesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListCompaniesRequest) GetRequireOpenJobs() bool {
	if m != nil {
		return m.RequireOpenJobs
	}
	return false
}

// The List companies response object.
type ListCompaniesResponse struct {
	// Companies for the current client.
	Companies []*Company `protobuf:"bytes,1,rep,name=companies,proto3" json:"companies,omitempty"`
	// A token to retrieve the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Additional information for the API invocation, such as the request
	// tracking id.
	Metadata             *ResponseMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListCompaniesResponse) Reset()         { *m = ListCompaniesResponse{} }
func (m *ListCompaniesResponse) String() string { return proto.CompactTextString(m) }
func (*ListCompaniesResponse) ProtoMessage()    {}
func (*ListCompaniesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c499e6cf60b8944, []int{5}
}

func (m *ListCompaniesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCompaniesResponse.Unmarshal(m, b)
}
func (m *ListCompaniesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCompaniesResponse.Marshal(b, m, deterministic)
}
func (m *ListCompaniesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCompaniesResponse.Merge(m, src)
}
func (m *ListCompaniesResponse) XXX_Size() int {
	return xxx_messageInfo_ListCompaniesResponse.Size(m)
}
func (m *ListCompaniesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCompaniesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCompaniesResponse proto.InternalMessageInfo

func (m *ListCompaniesResponse) GetCompanies() []*Company {
	if m != nil {
		return m.Companies
	}
	return nil
}

func (m *ListCompaniesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func (m *ListCompaniesResponse) GetMetadata() *ResponseMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateCompanyRequest)(nil), "google.cloud.talent.v4beta1.CreateCompanyRequest")
	proto.RegisterType((*GetCompanyRequest)(nil), "google.cloud.talent.v4beta1.GetCompanyRequest")
	proto.RegisterType((*UpdateCompanyRequest)(nil), "google.cloud.talent.v4beta1.UpdateCompanyRequest")
	proto.RegisterType((*DeleteCompanyRequest)(nil), "google.cloud.talent.v4beta1.DeleteCompanyRequest")
	proto.RegisterType((*ListCompaniesRequest)(nil), "google.cloud.talent.v4beta1.ListCompaniesRequest")
	proto.RegisterType((*ListCompaniesResponse)(nil), "google.cloud.talent.v4beta1.ListCompaniesResponse")
}

func init() {
	proto.RegisterFile("google/cloud/talent/v4beta1/company_service.proto", fileDescriptor_1c499e6cf60b8944)
}

var fileDescriptor_1c499e6cf60b8944 = []byte{
	// 799 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x4f, 0x6f, 0xeb, 0x44,
	0x10, 0xd7, 0x26, 0xa5, 0x34, 0x5b, 0x85, 0xaa, 0xdb, 0x40, 0xa3, 0x14, 0xd4, 0xc8, 0x42, 0x55,
	0x08, 0xad, 0x97, 0xa6, 0x1c, 0x50, 0x2b, 0x0e, 0x49, 0xf8, 0x23, 0x10, 0x15, 0x95, 0x53, 0x2e,
	0x15, 0x52, 0xb4, 0x76, 0xa6, 0x89, 0x5b, 0x7b, 0xd7, 0xb5, 0x37, 0x6d, 0x29, 0xea, 0x01, 0xbe,
	0x02, 0x17, 0x24, 0x0e, 0x48, 0x7c, 0x10, 0x8e, 0x1c, 0x38, 0xbc, 0x43, 0xdf, 0xed, 0x9d, 0x9e,
	0x9e, 0xde, 0x07, 0x79, 0xb2, 0xbd, 0xce, 0xff, 0x26, 0x7e, 0xbd, 0x79, 0x67, 0xe6, 0x37, 0xfb,
	0xfb, 0xcd, 0xec, 0x4c, 0x82, 0xf7, 0xbb, 0x42, 0x74, 0x1d, 0xa0, 0x96, 0x23, 0xfa, 0x1d, 0x2a,
	0x99, 0x03, 0x5c, 0xd2, 0xeb, 0xcf, 0x4d, 0x90, 0x6c, 0x9f, 0x5a, 0xc2, 0xf5, 0x18, 0xff, 0xa5,
	0x1d, 0x80, 0x7f, 0x6d, 0x5b, 0xa0, 0x7b, 0xbe, 0x90, 0x82, 0x6c, 0xc5, 0x10, 0x3d, 0x82, 0xe8,
	0x31, 0x44, 0x57, 0x90, 0xd2, 0x87, 0x2a, 0x1f, 0xf3, 0x6c, 0xca, 0x38, 0x17, 0x92, 0x49, 0x5b,
	0xf0, 0x20, 0x86, 0x96, 0x36, 0x47, 0xbc, 0x96, 0x63, 0x87, 0xc0, 0xd8, 0xb1, 0x3d, 0xe2, 0x38,
	0xb7, 0xc1, 0xe9, 0xb4, 0x4d, 0xe8, 0xb1, 0x6b, 0x5b, 0xf8, 0x2a, 0xa0, 0xb2, 0x80, 0xa7, 0x2b,
	0xb8, 0x8a, 0xfc, 0x24, 0x85, 0x22, 0x15, 0xaa, 0x94, 0xd0, 0xe8, 0x64, 0xf6, 0xcf, 0x29, 0xb8,
	0x9e, 0x4c, 0x9c, 0xe5, 0x49, 0x67, 0xcc, 0xcb, 0x65, 0xc1, 0x65, 0x1c, 0xa1, 0xdd, 0xe2, 0x42,
	0xd3, 0x07, 0x26, 0xa1, 0x19, 0x67, 0x35, 0xe0, 0xaa, 0x0f, 0x81, 0x24, 0x5b, 0x78, 0xd9, 0x63,
	0x3e, 0x70, 0x59, 0x44, 0x65, 0x54, 0xc9, 0x35, 0xb2, 0x2f, 0xeb, 0x19, 0x43, 0x99, 0x48, 0x13,
	0xbf, 0xab, 0x48, 0x14, 0x33, 0x65, 0x54, 0x59, 0xad, 0x7d, 0xac, 0xcf, 0xa9, 0xa7, 0xae, 0x52,
	0xc7, 0x39, 0x12, 0xa4, 0xb6, 0x8b, 0xd7, 0xbf, 0x05, 0x39, 0x71, 0xed, 0x26, 0x5e, 0xe2, 0xcc,
	0x85, 0xd1, 0x4b, 0x23, 0x83, 0xf6, 0x27, 0xc2, 0x85, 0x9f, 0xbc, 0xce, 0x34, 0xd1, 0x11, 0x2e,
	0xe8, 0xa9, 0x5c, 0xc8, 0x11, 0x5e, 0xed, 0x47, 0xc9, 0xa3, 0xd2, 0x28, 0x51, 0xa5, 0x24, 0x51,
	0x52, 0x3d, 0xfd, 0x9b, 0xb0, 0x7a, 0xc7, 0x2c, 0xb8, 0x34, 0x70, 0x1c, 0x1e, 0x7e, 0x6b, 0x14,
	0x17, 0xbe, 0x02, 0x07, 0xa6, 0x98, 0x3d, 0xaa, 0xe5, 0x2f, 0x84, 0x0b, 0x3f, 0xd8, 0x81, 0xd2,
	0x6e, 0x43, 0x90, 0xaa, 0xe8, 0x1f, 0x61, 0xec, 0xb1, 0x2e, 0xb4, 0xa5, 0xb8, 0x04, 0x1e, 0x51,
	0xcc, 0x19, 0xb9, 0xd0, 0x72, 0x1a, 0x1a, 0xc8, 0x16, 0x8e, 0x0e, 0xed, 0xc0, 0xbe, 0x83, 0x62,
	0xb6, 0x8c, 0x2a, 0xef, 0x18, 0x2b, 0xa1, 0xa1, 0x65, 0xdf, 0x01, 0xa9, 0xe2, 0x75, 0x1f, 0xae,
	0xfa, 0xb6, 0x0f, 0x6d, 0xe1, 0x01, 0x6f, 0x5f, 0x08, 0x33, 0x28, 0x2e, 0x95, 0x51, 0x65, 0xc5,
	0x58, 0x53, 0x8e, 0x1f, 0x3d, 0xe0, 0xdf, 0x0b, 0x33, 0xd0, 0x9e, 0x21, 0xfc, 0xfe, 0x04, 0xbb,
	0xc0, 0x13, 0x3c, 0x00, 0xd2, 0xc0, 0x39, 0x2b, 0x31, 0x16, 0x51, 0x39, 0x9b, 0xb6, 0xd8, 0xc6,
	0x10, 0x46, 0x76, 0xf0, 0x1a, 0x87, 0x5b, 0xd9, 0x9e, 0x92, 0x92, 0x0f, 0xcd, 0x27, 0x03, 0x39,
	0xdf, 0xe1, 0x15, 0x17, 0x24, 0xeb, 0x30, 0xc9, 0x22, 0x35, 0xab, 0xb5, 0xbd, 0xb9, 0x57, 0x25,
	0x24, 0x8f, 0x15, 0xc8, 0x18, 0xc0, 0x6b, 0x7f, 0xe7, 0xf0, 0x7b, 0x8a, 0x49, 0x2b, 0x5e, 0x02,
	0xe4, 0x01, 0xe1, 0xfc, 0xd8, 0xb3, 0x27, 0xfb, 0xf3, 0x85, 0xcc, 0x18, 0x91, 0x52, 0x2a, 0xed,
	0x5a, 0xef, 0xf7, 0xe7, 0xaf, 0xff, 0xc8, 0x98, 0xda, 0x67, 0x83, 0xf9, 0xfd, 0x35, 0x6e, 0xe8,
	0x97, 0x9e, 0x2f, 0x2e, 0xc0, 0x92, 0x01, 0xad, 0x52, 0x09, 0x9c, 0xf1, 0xf0, 0xeb, 0x9e, 0x0e,
	0x4a, 0x75, 0x88, 0xaa, 0x67, 0x9f, 0x6a, 0x3b, 0x73, 0x60, 0xe3, 0xc1, 0xe4, 0x3f, 0x84, 0xf1,
	0x70, 0xa2, 0x88, 0x3e, 0x97, 0xde, 0xd4, 0xe8, 0xa5, 0x94, 0x63, 0x46, 0x72, 0x7e, 0x26, 0x23,
	0x72, 0xc2, 0x37, 0x3d, 0x53, 0xcc, 0x90, 0x1e, 0xad, 0xde, 0x9f, 0x55, 0xc8, 0xce, 0xe3, 0x98,
	0xd1, 0x48, 0xf2, 0x0a, 0xe1, 0xfc, 0xd8, 0xac, 0x2f, 0xe8, 0xce, 0xac, 0xbd, 0x90, 0x52, 0xce,
	0x6d, 0x24, 0xc7, 0xaf, 0x7d, 0x31, 0xa4, 0x96, 0xac, 0xd7, 0x74, 0xb2, 0xc2, 0x2e, 0x1d, 0xd4,
	0xf4, 0xc5, 0xf0, 0x09, 0x10, 0xf9, 0x17, 0xe1, 0xfc, 0xd8, 0xda, 0x58, 0x20, 0x72, 0xd6, 0x8a,
	0x29, 0x7d, 0x30, 0xb5, 0xa2, 0xbe, 0x0e, 0xb7, 0x7f, 0xd2, 0xa5, 0xea, 0x13, 0xba, 0x54, 0x4d,
	0xdb, 0xa5, 0x17, 0x08, 0xe7, 0xc7, 0xf6, 0xc4, 0x02, 0x01, 0xb3, 0x36, 0x5e, 0xa9, 0xf6, 0x36,
	0x90, 0x78, 0xc2, 0x67, 0x3d, 0xc1, 0x74, 0x13, 0x35, 0xfe, 0x04, 0xe7, 0x8d, 0x53, 0xc9, 0xf9,
	0xbf, 0xbe, 0x11, 0x2e, 0x49, 0x45, 0x8f, 0x79, 0x76, 0xa0, 0x5b, 0xc2, 0x7d, 0xa8, 0xb7, 0x7a,
	0x52, 0x7a, 0xc1, 0x21, 0xa5, 0x37, 0x37, 0x37, 0x13, 0x4e, 0xca, 0xfa, 0xb2, 0x17, 0xff, 0x64,
	0xef, 0x79, 0x0e, 0x93, 0xe7, 0xc2, 0x77, 0x77, 0x17, 0x85, 0x87, 0x97, 0x34, 0x7e, 0x43, 0x78,
	0xdb, 0x12, 0xee, 0xbc, 0x5a, 0x34, 0x36, 0xc6, 0x57, 0xd8, 0x49, 0xd8, 0xf0, 0x13, 0x74, 0x56,
	0x57, 0x98, 0xae, 0x70, 0x18, 0xef, 0xea, 0xc2, 0xef, 0xd2, 0x2e, 0xf0, 0xe8, 0x39, 0xd0, 0xe1,
	0x7d, 0x33, 0xff, 0x48, 0x1c, 0xc5, 0xc7, 0x7f, 0x32, 0xd9, 0xe6, 0x69, 0xcb, 0x5c, 0x8e, 0x30,
	0x07, 0x6f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x46, 0x7b, 0xe0, 0xcb, 0x4d, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CompanyServiceClient is the client API for CompanyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompanyServiceClient interface {
	// Creates a new company entity.
	CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*Company, error)
	// Retrieves specified company.
	GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*Company, error)
	// Updates specified company.
	UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*Company, error)
	// Deletes specified company.
	// Prerequisite: The company has no jobs associated with it.
	DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Lists all companies associated with the project.
	ListCompanies(ctx context.Context, in *ListCompaniesRequest, opts ...grpc.CallOption) (*ListCompaniesResponse, error)
}

type companyServiceClient struct {
	cc *grpc.ClientConn
}

func NewCompanyServiceClient(cc *grpc.ClientConn) CompanyServiceClient {
	return &companyServiceClient{cc}
}

func (c *companyServiceClient) CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*Company, error) {
	out := new(Company)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.CompanyService/CreateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*Company, error) {
	out := new(Company)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.CompanyService/GetCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*Company, error) {
	out := new(Company)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.CompanyService/UpdateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.CompanyService/DeleteCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) ListCompanies(ctx context.Context, in *ListCompaniesRequest, opts ...grpc.CallOption) (*ListCompaniesResponse, error) {
	out := new(ListCompaniesResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.CompanyService/ListCompanies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyServiceServer is the server API for CompanyService service.
type CompanyServiceServer interface {
	// Creates a new company entity.
	CreateCompany(context.Context, *CreateCompanyRequest) (*Company, error)
	// Retrieves specified company.
	GetCompany(context.Context, *GetCompanyRequest) (*Company, error)
	// Updates specified company.
	UpdateCompany(context.Context, *UpdateCompanyRequest) (*Company, error)
	// Deletes specified company.
	// Prerequisite: The company has no jobs associated with it.
	DeleteCompany(context.Context, *DeleteCompanyRequest) (*empty.Empty, error)
	// Lists all companies associated with the project.
	ListCompanies(context.Context, *ListCompaniesRequest) (*ListCompaniesResponse, error)
}

// UnimplementedCompanyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCompanyServiceServer struct {
}

func (*UnimplementedCompanyServiceServer) CreateCompany(ctx context.Context, req *CreateCompanyRequest) (*Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}
func (*UnimplementedCompanyServiceServer) GetCompany(ctx context.Context, req *GetCompanyRequest) (*Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompany not implemented")
}
func (*UnimplementedCompanyServiceServer) UpdateCompany(ctx context.Context, req *UpdateCompanyRequest) (*Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (*UnimplementedCompanyServiceServer) DeleteCompany(ctx context.Context, req *DeleteCompanyRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompany not implemented")
}
func (*UnimplementedCompanyServiceServer) ListCompanies(ctx context.Context, req *ListCompaniesRequest) (*ListCompaniesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCompanies not implemented")
}

func RegisterCompanyServiceServer(s *grpc.Server, srv CompanyServiceServer) {
	s.RegisterService(&_CompanyService_serviceDesc, srv)
}

func _CompanyService_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.CompanyService/CreateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).CreateCompany(ctx, req.(*CreateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_GetCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).GetCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.CompanyService/GetCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).GetCompany(ctx, req.(*GetCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.CompanyService/UpdateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).UpdateCompany(ctx, req.(*UpdateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_DeleteCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).DeleteCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.CompanyService/DeleteCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).DeleteCompany(ctx, req.(*DeleteCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_ListCompanies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCompaniesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).ListCompanies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.CompanyService/ListCompanies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).ListCompanies(ctx, req.(*ListCompaniesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CompanyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.talent.v4beta1.CompanyService",
	HandlerType: (*CompanyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCompany",
			Handler:    _CompanyService_CreateCompany_Handler,
		},
		{
			MethodName: "GetCompany",
			Handler:    _CompanyService_GetCompany_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _CompanyService_UpdateCompany_Handler,
		},
		{
			MethodName: "DeleteCompany",
			Handler:    _CompanyService_DeleteCompany_Handler,
		},
		{
			MethodName: "ListCompanies",
			Handler:    _CompanyService_ListCompanies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/talent/v4beta1/company_service.proto",
}
