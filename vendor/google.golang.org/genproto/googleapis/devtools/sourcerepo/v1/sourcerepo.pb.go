// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/sourcerepo/v1/sourcerepo.proto

package sourcerepo

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	v1 "google.golang.org/genproto/googleapis/iam/v1"
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

// A repository (or repo) is a Git repository storing versioned source content.
type Repo struct {
	// Resource name of the repository, of the form
	// `projects/<project>/repos/<repo>`.  The repo name may contain slashes.
	// eg, `projects/myproject/repos/name/with/slash`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The disk usage of the repo, in bytes. Read-only field. Size is only
	// returned by GetRepo.
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	// URL to clone the repository from Google Cloud Source Repositories.
	// Read-only field.
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	// How this repository mirrors a repository managed by another service.
	// Read-only field.
	MirrorConfig         *MirrorConfig `protobuf:"bytes,4,opt,name=mirror_config,json=mirrorConfig,proto3" json:"mirror_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Repo) Reset()         { *m = Repo{} }
func (m *Repo) String() string { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()    {}
func (*Repo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e2c2c58455430a2, []int{0}
}

func (m *Repo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repo.Unmarshal(m, b)
}
func (m *Repo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repo.Marshal(b, m, deterministic)
}
func (m *Repo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repo.Merge(m, src)
}
func (m *Repo) XXX_Size() int {
	return xxx_messageInfo_Repo.Size(m)
}
func (m *Repo) XXX_DiscardUnknown() {
	xxx_messageInfo_Repo.DiscardUnknown(m)
}

var xxx_messageInfo_Repo proto.InternalMessageInfo

func (m *Repo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Repo) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Repo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Repo) GetMirrorConfig() *MirrorConfig {
	if m != nil {
		return m.MirrorConfig
	}
	return nil
}

// Configuration to automatically mirror a repository from another
// hosting service, for example GitHub or BitBucket.
type MirrorConfig struct {
	// URL of the main repository at the other hosting service.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// ID of the webhook listening to updates to trigger mirroring.
	// Removing this webhook from the other hosting service will stop
	// Google Cloud Source Repositories from receiving notifications,
	// and thereby disabling mirroring.
	WebhookId string `protobuf:"bytes,2,opt,name=webhook_id,json=webhookId,proto3" json:"webhook_id,omitempty"`
	// ID of the SSH deploy key at the other hosting service.
	// Removing this key from the other service would deauthorize
	// Google Cloud Source Repositories from mirroring.
	DeployKeyId          string   `protobuf:"bytes,3,opt,name=deploy_key_id,json=deployKeyId,proto3" json:"deploy_key_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MirrorConfig) Reset()         { *m = MirrorConfig{} }
func (m *MirrorConfig) String() string { return proto.CompactTextString(m) }
func (*MirrorConfig) ProtoMessage()    {}
func (*MirrorConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e2c2c58455430a2, []int{1}
}

func (m *MirrorConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MirrorConfig.Unmarshal(m, b)
}
func (m *MirrorConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MirrorConfig.Marshal(b, m, deterministic)
}
func (m *MirrorConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MirrorConfig.Merge(m, src)
}
func (m *MirrorConfig) XXX_Size() int {
	return xxx_messageInfo_MirrorConfig.Size(m)
}
func (m *MirrorConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MirrorConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MirrorConfig proto.InternalMessageInfo

func (m *MirrorConfig) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *MirrorConfig) GetWebhookId() string {
	if m != nil {
		return m.WebhookId
	}
	return ""
}

func (m *MirrorConfig) GetDeployKeyId() string {
	if m != nil {
		return m.DeployKeyId
	}
	return ""
}

// Request for GetRepo.
type GetRepoRequest struct {
	// The name of the requested repository. Values are of the form
	// `projects/<project>/repos/<repo>`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRepoRequest) Reset()         { *m = GetRepoRequest{} }
func (m *GetRepoRequest) String() string { return proto.CompactTextString(m) }
func (*GetRepoRequest) ProtoMessage()    {}
func (*GetRepoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e2c2c58455430a2, []int{2}
}

func (m *GetRepoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRepoRequest.Unmarshal(m, b)
}
func (m *GetRepoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRepoRequest.Marshal(b, m, deterministic)
}
func (m *GetRepoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRepoRequest.Merge(m, src)
}
func (m *GetRepoRequest) XXX_Size() int {
	return xxx_messageInfo_GetRepoRequest.Size(m)
}
func (m *GetRepoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRepoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRepoRequest proto.InternalMessageInfo

func (m *GetRepoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request for ListRepos.
type ListReposRequest struct {
	// The project ID whose repos should be listed. Values are of the form
	// `projects/<project>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Maximum number of repositories to return; between 1 and 500.
	// If not set or zero, defaults to 100 at the server.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Resume listing repositories where a prior ListReposResponse
	// left off. This is an opaque token that must be obtained from
	// a recent, prior ListReposResponse's next_page_token field.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListReposRequest) Reset()         { *m = ListReposRequest{} }
func (m *ListReposRequest) String() string { return proto.CompactTextString(m) }
func (*ListReposRequest) ProtoMessage()    {}
func (*ListReposRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e2c2c58455430a2, []int{3}
}

func (m *ListReposRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReposRequest.Unmarshal(m, b)
}
func (m *ListReposRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReposRequest.Marshal(b, m, deterministic)
}
func (m *ListReposRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReposRequest.Merge(m, src)
}
func (m *ListReposRequest) XXX_Size() int {
	return xxx_messageInfo_ListReposRequest.Size(m)
}
func (m *ListReposRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReposRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListReposRequest proto.InternalMessageInfo

func (m *ListReposRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListReposRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListReposRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Response for ListRepos.  The size is not set in the returned repositories.
type ListReposResponse struct {
	// The listed repos.
	Repos []*Repo `protobuf:"bytes,1,rep,name=repos,proto3" json:"repos,omitempty"`
	// If non-empty, additional repositories exist within the project. These
	// can be retrieved by including this value in the next ListReposRequest's
	// page_token field.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListReposResponse) Reset()         { *m = ListReposResponse{} }
func (m *ListReposResponse) String() string { return proto.CompactTextString(m) }
func (*ListReposResponse) ProtoMessage()    {}
func (*ListReposResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e2c2c58455430a2, []int{4}
}

func (m *ListReposResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReposResponse.Unmarshal(m, b)
}
func (m *ListReposResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReposResponse.Marshal(b, m, deterministic)
}
func (m *ListReposResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReposResponse.Merge(m, src)
}
func (m *ListReposResponse) XXX_Size() int {
	return xxx_messageInfo_ListReposResponse.Size(m)
}
func (m *ListReposResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReposResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListReposResponse proto.InternalMessageInfo

func (m *ListReposResponse) GetRepos() []*Repo {
	if m != nil {
		return m.Repos
	}
	return nil
}

func (m *ListReposResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request for CreateRepo
type CreateRepoRequest struct {
	// The project in which to create the repo. Values are of the form
	// `projects/<project>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The repo to create.  Only name should be set; setting other fields
	// is an error.  The project in the name should match the parent field.
	Repo                 *Repo    `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRepoRequest) Reset()         { *m = CreateRepoRequest{} }
func (m *CreateRepoRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRepoRequest) ProtoMessage()    {}
func (*CreateRepoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e2c2c58455430a2, []int{5}
}

func (m *CreateRepoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRepoRequest.Unmarshal(m, b)
}
func (m *CreateRepoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRepoRequest.Marshal(b, m, deterministic)
}
func (m *CreateRepoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRepoRequest.Merge(m, src)
}
func (m *CreateRepoRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRepoRequest.Size(m)
}
func (m *CreateRepoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRepoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRepoRequest proto.InternalMessageInfo

func (m *CreateRepoRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateRepoRequest) GetRepo() *Repo {
	if m != nil {
		return m.Repo
	}
	return nil
}

// Request for DeleteRepo.
type DeleteRepoRequest struct {
	// The name of the repo to delete. Values are of the form
	// `projects/<project>/repos/<repo>`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRepoRequest) Reset()         { *m = DeleteRepoRequest{} }
func (m *DeleteRepoRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRepoRequest) ProtoMessage()    {}
func (*DeleteRepoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e2c2c58455430a2, []int{6}
}

func (m *DeleteRepoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRepoRequest.Unmarshal(m, b)
}
func (m *DeleteRepoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRepoRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRepoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRepoRequest.Merge(m, src)
}
func (m *DeleteRepoRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRepoRequest.Size(m)
}
func (m *DeleteRepoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRepoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRepoRequest proto.InternalMessageInfo

func (m *DeleteRepoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Repo)(nil), "google.devtools.sourcerepo.v1.Repo")
	proto.RegisterType((*MirrorConfig)(nil), "google.devtools.sourcerepo.v1.MirrorConfig")
	proto.RegisterType((*GetRepoRequest)(nil), "google.devtools.sourcerepo.v1.GetRepoRequest")
	proto.RegisterType((*ListReposRequest)(nil), "google.devtools.sourcerepo.v1.ListReposRequest")
	proto.RegisterType((*ListReposResponse)(nil), "google.devtools.sourcerepo.v1.ListReposResponse")
	proto.RegisterType((*CreateRepoRequest)(nil), "google.devtools.sourcerepo.v1.CreateRepoRequest")
	proto.RegisterType((*DeleteRepoRequest)(nil), "google.devtools.sourcerepo.v1.DeleteRepoRequest")
}

func init() {
	proto.RegisterFile("google/devtools/sourcerepo/v1/sourcerepo.proto", fileDescriptor_0e2c2c58455430a2)
}

var fileDescriptor_0e2c2c58455430a2 = []byte{
	// 743 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xd1, 0x6e, 0xd3, 0x4a,
	0x10, 0xd5, 0x36, 0x69, 0x7b, 0x33, 0x4d, 0x6f, 0xdb, 0x95, 0x6e, 0x15, 0xa5, 0x37, 0x55, 0xae,
	0x7b, 0x29, 0x21, 0x15, 0x36, 0x2d, 0xa0, 0x8a, 0x20, 0x24, 0xd4, 0x82, 0xa2, 0x0a, 0x90, 0xa2,
	0xb4, 0x4f, 0xbc, 0x44, 0x4e, 0x32, 0x35, 0xa6, 0xb6, 0xd7, 0x78, 0x37, 0x81, 0x80, 0x0a, 0x52,
	0xa5, 0xbe, 0x23, 0xfa, 0x19, 0x7c, 0x0e, 0xbf, 0xc0, 0x47, 0xf0, 0x88, 0x76, 0x6d, 0x37, 0x4e,
	0x13, 0x12, 0xbf, 0xed, 0xce, 0x9c, 0x99, 0x73, 0xf6, 0xec, 0x78, 0x0d, 0xba, 0xc5, 0x98, 0xe5,
	0xa0, 0xd1, 0xc5, 0xbe, 0x60, 0xcc, 0xe1, 0x06, 0x67, 0xbd, 0xa0, 0x83, 0x01, 0xfa, 0xcc, 0xe8,
	0xef, 0x26, 0x76, 0xba, 0x1f, 0x30, 0xc1, 0x68, 0x29, 0xc4, 0xeb, 0x31, 0x5e, 0x4f, 0x20, 0xfa,
	0xbb, 0xc5, 0x7f, 0xa3, 0x76, 0xa6, 0x6f, 0x1b, 0xa6, 0xe7, 0x31, 0x61, 0x0a, 0x9b, 0x79, 0x3c,
	0x2c, 0x2e, 0x6e, 0x46, 0x59, 0xdb, 0x74, 0x65, 0x73, 0xdb, 0x74, 0x5b, 0x3e, 0x73, 0xec, 0xce,
	0x20, 0xca, 0x17, 0x47, 0xf3, 0x23, 0xb9, 0x8d, 0x28, 0xa7, 0x76, 0xed, 0xde, 0xa9, 0x81, 0xae,
	0x2f, 0xa2, 0xa4, 0xf6, 0x8d, 0x40, 0xb6, 0x89, 0x3e, 0xa3, 0x14, 0xb2, 0x9e, 0xe9, 0x62, 0x81,
	0x94, 0x49, 0x25, 0xd7, 0x54, 0x6b, 0x19, 0xe3, 0xf6, 0x47, 0x2c, 0xcc, 0x95, 0x49, 0x25, 0xd3,
	0x54, 0x6b, 0xba, 0x0a, 0x99, 0x5e, 0xe0, 0x14, 0x32, 0x0a, 0x26, 0x97, 0xb4, 0x01, 0xcb, 0xae,
	0x1d, 0x04, 0x2c, 0x68, 0x75, 0x98, 0x77, 0x6a, 0x5b, 0x85, 0x6c, 0x99, 0x54, 0x96, 0xf6, 0x76,
	0xf4, 0xa9, 0x07, 0xd6, 0x5f, 0xa9, 0x9a, 0x43, 0x55, 0xd2, 0xcc, 0xbb, 0x89, 0x9d, 0xd6, 0x81,
	0x7c, 0x32, 0x1b, 0x73, 0x92, 0x21, 0x67, 0x09, 0xe0, 0x3d, 0xb6, 0xdf, 0x30, 0x76, 0xd6, 0xb2,
	0xbb, 0x4a, 0x5f, 0xae, 0x99, 0x8b, 0x22, 0x47, 0x5d, 0xaa, 0xc1, 0x72, 0x17, 0x7d, 0x87, 0x0d,
	0x5a, 0x67, 0x38, 0x90, 0x88, 0x50, 0xee, 0x52, 0x18, 0x7c, 0x81, 0x83, 0xa3, 0xae, 0xf6, 0x3f,
	0xfc, 0x5d, 0x47, 0x21, 0xcf, 0xde, 0xc4, 0x77, 0x3d, 0xe4, 0x62, 0x92, 0x05, 0x5a, 0x1b, 0x56,
	0x5f, 0xda, 0x5c, 0xc1, 0xf8, 0x14, 0x1c, 0xdd, 0x80, 0x9c, 0x6f, 0x5a, 0xd8, 0xba, 0xf6, 0x6b,
	0xbe, 0xf9, 0x97, 0x0c, 0x1c, 0x4b, 0xcf, 0x4a, 0x00, 0x2a, 0x29, 0xd8, 0x19, 0x7a, 0x91, 0x16,
	0x05, 0x3f, 0x91, 0x01, 0xad, 0x0f, 0x6b, 0x09, 0x0e, 0xee, 0x33, 0x8f, 0x23, 0x7d, 0x04, 0xf3,
	0xd2, 0x29, 0x5e, 0x20, 0xe5, 0x4c, 0x65, 0x69, 0x6f, 0x6b, 0x86, 0x9b, 0xea, 0x1c, 0x61, 0x05,
	0xdd, 0x86, 0x15, 0x0f, 0x3f, 0x88, 0x56, 0x82, 0x33, 0x74, 0x68, 0x59, 0x86, 0x1b, 0xd7, 0xbc,
	0x5d, 0x58, 0x3b, 0x0c, 0xd0, 0x14, 0x98, 0x34, 0x61, 0x1d, 0x16, 0x7c, 0x33, 0x40, 0x4f, 0x44,
	0xc7, 0x8b, 0x76, 0x74, 0x1f, 0xb2, 0xb2, 0xbb, 0xea, 0x94, 0x52, 0x8e, 0x2a, 0xd0, 0x6e, 0xc3,
	0xda, 0x33, 0x74, 0x70, 0x94, 0x65, 0x82, 0x85, 0x7b, 0xbf, 0x16, 0x01, 0x8e, 0x55, 0x17, 0x35,
	0x90, 0x57, 0x04, 0x72, 0xd7, 0xb6, 0x50, 0x63, 0x06, 0xe1, 0xcd, 0x4b, 0x2a, 0xde, 0x4b, 0x5f,
	0x10, 0x3a, 0xae, 0x6d, 0x5d, 0xfc, 0xf8, 0x79, 0x35, 0x57, 0xa2, 0x1b, 0xf2, 0x0b, 0xfa, 0x24,
	0x25, 0x3d, 0xf1, 0x03, 0xf6, 0x16, 0x3b, 0x82, 0x1b, 0xd5, 0x73, 0x23, 0xf4, 0xf6, 0x92, 0xc0,
	0x62, 0x34, 0x36, 0xf4, 0xee, 0x0c, 0x8a, 0xd1, 0xf1, 0x2a, 0xa6, 0xf1, 0x4c, 0xdb, 0x56, 0x22,
	0xca, 0x74, 0x73, 0x92, 0x88, 0x50, 0x83, 0x51, 0xad, 0x9e, 0xd3, 0xaf, 0x04, 0x60, 0x78, 0x79,
	0x74, 0xd6, 0x69, 0xc7, 0xee, 0x39, 0x9d, 0x9a, 0x1d, 0xa5, 0xe6, 0x96, 0x56, 0x52, 0x6a, 0xc2,
	0x49, 0x18, 0x37, 0xa5, 0xa6, 0x2e, 0x9a, 0x7e, 0x06, 0x18, 0x5e, 0xf4, 0x4c, 0x45, 0x63, 0x33,
	0x51, 0x5c, 0x8f, 0x2b, 0xe2, 0x87, 0x4a, 0x7f, 0x2e, 0x1f, 0xaa, 0xd8, 0x92, 0xea, 0x2c, 0x4b,
	0x2e, 0x09, 0xe4, 0x8f, 0x51, 0x1c, 0x99, 0x6e, 0x43, 0x3d, 0x7f, 0x54, 0x8b, 0x1b, 0xda, 0xa6,
	0x2b, 0x29, 0x93, 0xc9, 0x98, 0xf4, 0x9f, 0x1b, 0x98, 0x30, 0xab, 0xd5, 0x14, 0xe7, 0x03, 0xcd,
	0x50, 0x9c, 0x01, 0x86, 0xda, 0x27, 0xf2, 0xd6, 0x78, 0xa2, 0x6d, 0x8d, 0x54, 0xe9, 0x05, 0x81,
	0x7c, 0x7d, 0x9a, 0x8e, 0x7a, 0x7a, 0x1d, 0xfb, 0x4a, 0xc7, 0x2e, 0x4d, 0xa3, 0xc3, 0x4a, 0x72,
	0x7e, 0x27, 0x40, 0x4f, 0x90, 0xab, 0x08, 0x06, 0xae, 0xcd, 0xb9, 0xfc, 0x9b, 0xd0, 0xca, 0x0d,
	0x9a, 0x71, 0x48, 0x2c, 0xe8, 0x4e, 0x0a, 0x64, 0xf4, 0xe1, 0x3c, 0x55, 0x22, 0x6b, 0xda, 0xc3,
	0x14, 0x22, 0xc5, 0x58, 0x9b, 0x1a, 0xa9, 0x1e, 0x7c, 0x81, 0xff, 0x3a, 0xcc, 0x9d, 0x3e, 0x31,
	0x07, 0x2b, 0xc3, 0xc7, 0xa1, 0x21, 0x27, 0xa4, 0x41, 0x5e, 0xd7, 0xa3, 0x0a, 0x8b, 0x39, 0xa6,
	0x67, 0xe9, 0x2c, 0xb0, 0x0c, 0x0b, 0x3d, 0x35, 0x3f, 0x46, 0x98, 0x32, 0x7d, 0x9b, 0xff, 0xe1,
	0x17, 0xfd, 0x78, 0xb8, 0x6b, 0x2f, 0xa8, 0x9a, 0xfb, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x31,
	0x75, 0x14, 0x03, 0xd5, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SourceRepoClient is the client API for SourceRepo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SourceRepoClient interface {
	// Returns all repos belonging to a project. The sizes of the repos are
	// not set by ListRepos.  To get the size of a repo, use GetRepo.
	ListRepos(ctx context.Context, in *ListReposRequest, opts ...grpc.CallOption) (*ListReposResponse, error)
	// Returns information about a repo.
	GetRepo(ctx context.Context, in *GetRepoRequest, opts ...grpc.CallOption) (*Repo, error)
	// Creates a repo in the given project with the given name.
	//
	// If the named repository already exists, `CreateRepo` returns
	// `ALREADY_EXISTS`.
	CreateRepo(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*Repo, error)
	// Deletes a repo.
	DeleteRepo(ctx context.Context, in *DeleteRepoRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Sets the access control policy on the specified resource. Replaces any
	// existing policy.
	SetIamPolicy(ctx context.Context, in *v1.SetIamPolicyRequest, opts ...grpc.CallOption) (*v1.Policy, error)
	// Gets the access control policy for a resource.
	// Returns an empty policy if the resource exists and does not have a policy
	// set.
	GetIamPolicy(ctx context.Context, in *v1.GetIamPolicyRequest, opts ...grpc.CallOption) (*v1.Policy, error)
	// Returns permissions that a caller has on the specified resource.
	// If the resource does not exist, this will return an empty set of
	// permissions, not a NOT_FOUND error.
	TestIamPermissions(ctx context.Context, in *v1.TestIamPermissionsRequest, opts ...grpc.CallOption) (*v1.TestIamPermissionsResponse, error)
}

type sourceRepoClient struct {
	cc *grpc.ClientConn
}

func NewSourceRepoClient(cc *grpc.ClientConn) SourceRepoClient {
	return &sourceRepoClient{cc}
}

func (c *sourceRepoClient) ListRepos(ctx context.Context, in *ListReposRequest, opts ...grpc.CallOption) (*ListReposResponse, error) {
	out := new(ListReposResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.sourcerepo.v1.SourceRepo/ListRepos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceRepoClient) GetRepo(ctx context.Context, in *GetRepoRequest, opts ...grpc.CallOption) (*Repo, error) {
	out := new(Repo)
	err := c.cc.Invoke(ctx, "/google.devtools.sourcerepo.v1.SourceRepo/GetRepo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceRepoClient) CreateRepo(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*Repo, error) {
	out := new(Repo)
	err := c.cc.Invoke(ctx, "/google.devtools.sourcerepo.v1.SourceRepo/CreateRepo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceRepoClient) DeleteRepo(ctx context.Context, in *DeleteRepoRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.devtools.sourcerepo.v1.SourceRepo/DeleteRepo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceRepoClient) SetIamPolicy(ctx context.Context, in *v1.SetIamPolicyRequest, opts ...grpc.CallOption) (*v1.Policy, error) {
	out := new(v1.Policy)
	err := c.cc.Invoke(ctx, "/google.devtools.sourcerepo.v1.SourceRepo/SetIamPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceRepoClient) GetIamPolicy(ctx context.Context, in *v1.GetIamPolicyRequest, opts ...grpc.CallOption) (*v1.Policy, error) {
	out := new(v1.Policy)
	err := c.cc.Invoke(ctx, "/google.devtools.sourcerepo.v1.SourceRepo/GetIamPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceRepoClient) TestIamPermissions(ctx context.Context, in *v1.TestIamPermissionsRequest, opts ...grpc.CallOption) (*v1.TestIamPermissionsResponse, error) {
	out := new(v1.TestIamPermissionsResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.sourcerepo.v1.SourceRepo/TestIamPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SourceRepoServer is the server API for SourceRepo service.
type SourceRepoServer interface {
	// Returns all repos belonging to a project. The sizes of the repos are
	// not set by ListRepos.  To get the size of a repo, use GetRepo.
	ListRepos(context.Context, *ListReposRequest) (*ListReposResponse, error)
	// Returns information about a repo.
	GetRepo(context.Context, *GetRepoRequest) (*Repo, error)
	// Creates a repo in the given project with the given name.
	//
	// If the named repository already exists, `CreateRepo` returns
	// `ALREADY_EXISTS`.
	CreateRepo(context.Context, *CreateRepoRequest) (*Repo, error)
	// Deletes a repo.
	DeleteRepo(context.Context, *DeleteRepoRequest) (*empty.Empty, error)
	// Sets the access control policy on the specified resource. Replaces any
	// existing policy.
	SetIamPolicy(context.Context, *v1.SetIamPolicyRequest) (*v1.Policy, error)
	// Gets the access control policy for a resource.
	// Returns an empty policy if the resource exists and does not have a policy
	// set.
	GetIamPolicy(context.Context, *v1.GetIamPolicyRequest) (*v1.Policy, error)
	// Returns permissions that a caller has on the specified resource.
	// If the resource does not exist, this will return an empty set of
	// permissions, not a NOT_FOUND error.
	TestIamPermissions(context.Context, *v1.TestIamPermissionsRequest) (*v1.TestIamPermissionsResponse, error)
}

// UnimplementedSourceRepoServer can be embedded to have forward compatible implementations.
type UnimplementedSourceRepoServer struct {
}

func (*UnimplementedSourceRepoServer) ListRepos(ctx context.Context, req *ListReposRequest) (*ListReposResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepos not implemented")
}
func (*UnimplementedSourceRepoServer) GetRepo(ctx context.Context, req *GetRepoRequest) (*Repo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepo not implemented")
}
func (*UnimplementedSourceRepoServer) CreateRepo(ctx context.Context, req *CreateRepoRequest) (*Repo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepo not implemented")
}
func (*UnimplementedSourceRepoServer) DeleteRepo(ctx context.Context, req *DeleteRepoRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRepo not implemented")
}
func (*UnimplementedSourceRepoServer) SetIamPolicy(ctx context.Context, req *v1.SetIamPolicyRequest) (*v1.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIamPolicy not implemented")
}
func (*UnimplementedSourceRepoServer) GetIamPolicy(ctx context.Context, req *v1.GetIamPolicyRequest) (*v1.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIamPolicy not implemented")
}
func (*UnimplementedSourceRepoServer) TestIamPermissions(ctx context.Context, req *v1.TestIamPermissionsRequest) (*v1.TestIamPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestIamPermissions not implemented")
}

func RegisterSourceRepoServer(s *grpc.Server, srv SourceRepoServer) {
	s.RegisterService(&_SourceRepo_serviceDesc, srv)
}

func _SourceRepo_ListRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReposRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceRepoServer).ListRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.sourcerepo.v1.SourceRepo/ListRepos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceRepoServer).ListRepos(ctx, req.(*ListReposRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourceRepo_GetRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceRepoServer).GetRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.sourcerepo.v1.SourceRepo/GetRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceRepoServer).GetRepo(ctx, req.(*GetRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourceRepo_CreateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceRepoServer).CreateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.sourcerepo.v1.SourceRepo/CreateRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceRepoServer).CreateRepo(ctx, req.(*CreateRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourceRepo_DeleteRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceRepoServer).DeleteRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.sourcerepo.v1.SourceRepo/DeleteRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceRepoServer).DeleteRepo(ctx, req.(*DeleteRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourceRepo_SetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceRepoServer).SetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.sourcerepo.v1.SourceRepo/SetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceRepoServer).SetIamPolicy(ctx, req.(*v1.SetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourceRepo_GetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceRepoServer).GetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.sourcerepo.v1.SourceRepo/GetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceRepoServer).GetIamPolicy(ctx, req.(*v1.GetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourceRepo_TestIamPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.TestIamPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceRepoServer).TestIamPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.sourcerepo.v1.SourceRepo/TestIamPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceRepoServer).TestIamPermissions(ctx, req.(*v1.TestIamPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SourceRepo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.sourcerepo.v1.SourceRepo",
	HandlerType: (*SourceRepoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRepos",
			Handler:    _SourceRepo_ListRepos_Handler,
		},
		{
			MethodName: "GetRepo",
			Handler:    _SourceRepo_GetRepo_Handler,
		},
		{
			MethodName: "CreateRepo",
			Handler:    _SourceRepo_CreateRepo_Handler,
		},
		{
			MethodName: "DeleteRepo",
			Handler:    _SourceRepo_DeleteRepo_Handler,
		},
		{
			MethodName: "SetIamPolicy",
			Handler:    _SourceRepo_SetIamPolicy_Handler,
		},
		{
			MethodName: "GetIamPolicy",
			Handler:    _SourceRepo_GetIamPolicy_Handler,
		},
		{
			MethodName: "TestIamPermissions",
			Handler:    _SourceRepo_TestIamPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/sourcerepo/v1/sourcerepo.proto",
}
