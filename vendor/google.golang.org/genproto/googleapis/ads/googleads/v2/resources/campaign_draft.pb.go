// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/campaign_draft.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// A campaign draft.
type CampaignDraft struct {
	// The resource name of the campaign draft.
	// Campaign draft resource names have the form:
	//
	// `customers/{customer_id}/campaignDrafts/{base_campaign_id}~{draft_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the draft.
	//
	// This field is read-only.
	DraftId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=draft_id,json=draftId,proto3" json:"draft_id,omitempty"`
	// The base campaign to which the draft belongs.
	BaseCampaign *wrappers.StringValue `protobuf:"bytes,3,opt,name=base_campaign,json=baseCampaign,proto3" json:"base_campaign,omitempty"`
	// The name of the campaign draft.
	//
	// This field is required and should not be empty when creating new
	// campaign drafts.
	//
	// It must not contain any null (code point 0x0), NL line feed
	// (code point 0xA) or carriage return (code point 0xD) characters.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Resource name of the Campaign that results from overlaying the draft
	// changes onto the base campaign.
	//
	// This field is read-only.
	DraftCampaign *wrappers.StringValue `protobuf:"bytes,5,opt,name=draft_campaign,json=draftCampaign,proto3" json:"draft_campaign,omitempty"`
	// The status of the campaign draft. This field is read-only.
	//
	// When a new campaign draft is added, the status defaults to PROPOSED.
	Status enums.CampaignDraftStatusEnum_CampaignDraftStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.CampaignDraftStatusEnum_CampaignDraftStatus" json:"status,omitempty"`
	// Whether there is an experiment based on this draft currently serving.
	HasExperimentRunning *wrappers.BoolValue `protobuf:"bytes,7,opt,name=has_experiment_running,json=hasExperimentRunning,proto3" json:"has_experiment_running,omitempty"`
	// The resource name of the long-running operation that can be used to poll
	// for completion of draft promotion. This is only set if the draft promotion
	// is in progress or finished.
	LongRunningOperation *wrappers.StringValue `protobuf:"bytes,8,opt,name=long_running_operation,json=longRunningOperation,proto3" json:"long_running_operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CampaignDraft) Reset()         { *m = CampaignDraft{} }
func (m *CampaignDraft) String() string { return proto.CompactTextString(m) }
func (*CampaignDraft) ProtoMessage()    {}
func (*CampaignDraft) Descriptor() ([]byte, []int) {
	return fileDescriptor_75e3b9d8638a5c8e, []int{0}
}

func (m *CampaignDraft) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignDraft.Unmarshal(m, b)
}
func (m *CampaignDraft) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignDraft.Marshal(b, m, deterministic)
}
func (m *CampaignDraft) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignDraft.Merge(m, src)
}
func (m *CampaignDraft) XXX_Size() int {
	return xxx_messageInfo_CampaignDraft.Size(m)
}
func (m *CampaignDraft) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignDraft.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignDraft proto.InternalMessageInfo

func (m *CampaignDraft) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignDraft) GetDraftId() *wrappers.Int64Value {
	if m != nil {
		return m.DraftId
	}
	return nil
}

func (m *CampaignDraft) GetBaseCampaign() *wrappers.StringValue {
	if m != nil {
		return m.BaseCampaign
	}
	return nil
}

func (m *CampaignDraft) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CampaignDraft) GetDraftCampaign() *wrappers.StringValue {
	if m != nil {
		return m.DraftCampaign
	}
	return nil
}

func (m *CampaignDraft) GetStatus() enums.CampaignDraftStatusEnum_CampaignDraftStatus {
	if m != nil {
		return m.Status
	}
	return enums.CampaignDraftStatusEnum_UNSPECIFIED
}

func (m *CampaignDraft) GetHasExperimentRunning() *wrappers.BoolValue {
	if m != nil {
		return m.HasExperimentRunning
	}
	return nil
}

func (m *CampaignDraft) GetLongRunningOperation() *wrappers.StringValue {
	if m != nil {
		return m.LongRunningOperation
	}
	return nil
}

func init() {
	proto.RegisterType((*CampaignDraft)(nil), "google.ads.googleads.v2.resources.CampaignDraft")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/campaign_draft.proto", fileDescriptor_75e3b9d8638a5c8e)
}

var fileDescriptor_75e3b9d8638a5c8e = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6a, 0xdb, 0x30,
	0x18, 0xc7, 0x71, 0x9a, 0xa5, 0x9d, 0xd6, 0xf4, 0x20, 0x4a, 0x31, 0x59, 0x19, 0xe9, 0x46, 0x21,
	0x27, 0x79, 0x78, 0x23, 0x30, 0xef, 0xe4, 0x74, 0xa5, 0xb4, 0x87, 0x2d, 0xb8, 0x90, 0xc3, 0x08,
	0x18, 0x25, 0x56, 0x55, 0x83, 0x2d, 0x09, 0x49, 0xee, 0xf6, 0x3c, 0x3b, 0xee, 0xb4, 0xe7, 0xd8,
	0xa3, 0xec, 0x25, 0x36, 0x2c, 0x59, 0x82, 0xae, 0xcb, 0x9a, 0xdb, 0x67, 0xe9, 0xff, 0xfb, 0xeb,
	0xaf, 0xef, 0x93, 0xc1, 0x94, 0x72, 0x4e, 0x2b, 0x12, 0xe1, 0x42, 0x45, 0xb6, 0x6c, 0xab, 0xbb,
	0x38, 0x92, 0x44, 0xf1, 0x46, 0xae, 0x89, 0x8a, 0xd6, 0xb8, 0x16, 0xb8, 0xa4, 0x2c, 0x2f, 0x24,
	0xbe, 0xd1, 0x48, 0x48, 0xae, 0x39, 0x3c, 0xb1, 0x62, 0x84, 0x0b, 0x85, 0x3c, 0x87, 0xee, 0x62,
	0xe4, 0xb9, 0xd1, 0xbb, 0x4d, 0xd6, 0x84, 0x35, 0xf5, 0xdf, 0xb6, 0xb9, 0xd2, 0x58, 0x37, 0xca,
	0xba, 0x8f, 0x5e, 0x74, 0xa8, 0xf9, 0x5a, 0x35, 0x37, 0xd1, 0x17, 0x89, 0x85, 0x20, 0xd2, 0xed,
	0x1f, 0x3b, 0x6b, 0x51, 0x46, 0x98, 0x31, 0xae, 0xb1, 0x2e, 0x39, 0xeb, 0x76, 0x5f, 0xfe, 0xe8,
	0x83, 0xe1, 0x59, 0xe7, 0xfe, 0xa1, 0x35, 0x87, 0xaf, 0xc0, 0xd0, 0xe5, 0xca, 0x19, 0xae, 0x49,
	0x18, 0x8c, 0x83, 0xc9, 0xd3, 0x6c, 0xdf, 0x2d, 0x7e, 0xc4, 0x35, 0x81, 0x53, 0xb0, 0x67, 0xa3,
	0x94, 0x45, 0xd8, 0x1b, 0x07, 0x93, 0x67, 0xf1, 0xf3, 0xee, 0x6a, 0xc8, 0xe5, 0x40, 0x97, 0x4c,
	0x4f, 0xdf, 0x2e, 0x70, 0xd5, 0x90, 0x6c, 0xd7, 0x88, 0x2f, 0x0b, 0x98, 0x82, 0xe1, 0x0a, 0x2b,
	0x92, 0xbb, 0x0b, 0x85, 0x3b, 0x06, 0x3e, 0x7e, 0x00, 0x5f, 0x6b, 0x59, 0x32, 0x6a, 0xe9, 0xfd,
	0x16, 0x71, 0x21, 0xe1, 0x6b, 0xd0, 0x37, 0xb1, 0xfa, 0x5b, 0x90, 0x46, 0x09, 0xcf, 0xc0, 0x81,
	0x0d, 0xeb, 0x4f, 0x7d, 0xb2, 0x05, 0x3b, 0x34, 0x8c, 0x3f, 0x76, 0x05, 0x06, 0xb6, 0xed, 0xe1,
	0x60, 0x1c, 0x4c, 0x0e, 0xe2, 0x2b, 0xb4, 0x69, 0xaa, 0x66, 0x64, 0xe8, 0x5e, 0x53, 0xaf, 0x0d,
	0x79, 0xce, 0x9a, 0xfa, 0x5f, 0xeb, 0x59, 0xe7, 0x0c, 0xe7, 0xe0, 0xe8, 0x16, 0xab, 0x9c, 0x7c,
	0x15, 0x44, 0x96, 0x35, 0x61, 0x3a, 0x97, 0x0d, 0x63, 0x25, 0xa3, 0xe1, 0xae, 0x09, 0x3c, 0x7a,
	0x10, 0x78, 0xc6, 0x79, 0x65, 0xe3, 0x1e, 0xde, 0x62, 0x75, 0xee, 0xc1, 0xcc, 0x72, 0x30, 0x03,
	0x47, 0x15, 0x67, 0xd4, 0xf9, 0xe4, 0x5c, 0x10, 0x69, 0xe6, 0x1f, 0xee, 0x6d, 0xd1, 0x82, 0xc3,
	0x96, 0xed, 0xac, 0x3e, 0x39, 0x72, 0xf6, 0x3b, 0x00, 0xa7, 0x6b, 0x5e, 0xa3, 0x47, 0x5f, 0xf5,
	0x0c, 0xde, 0xbb, 0xec, 0xbc, 0x3d, 0x62, 0x1e, 0x7c, 0xbe, 0xea, 0x40, 0xca, 0x2b, 0xcc, 0x28,
	0xe2, 0x92, 0x46, 0x94, 0x30, 0x13, 0xc0, 0xbd, 0x7d, 0x51, 0xaa, 0xff, 0xfc, 0x65, 0xef, 0x7d,
	0xf5, 0xad, 0xb7, 0x73, 0x91, 0xa6, 0xdf, 0x7b, 0x27, 0x17, 0xd6, 0x32, 0x2d, 0x14, 0xb2, 0x65,
	0x5b, 0x2d, 0x62, 0x94, 0x39, 0xe5, 0x4f, 0xa7, 0x59, 0xa6, 0x85, 0x5a, 0x7a, 0xcd, 0x72, 0x11,
	0x2f, 0xbd, 0xe6, 0x57, 0xef, 0xd4, 0x6e, 0x24, 0x49, 0x5a, 0xa8, 0x24, 0xf1, 0xaa, 0x24, 0x59,
	0xc4, 0x49, 0xe2, 0x75, 0xab, 0x81, 0x09, 0xfb, 0xe6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89,
	0xac, 0x9b, 0xd2, 0x11, 0x04, 0x00, 0x00,
}
