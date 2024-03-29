// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/streetview/publish/v1/resources.proto

package publish

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
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

// Status of rights transfer.
type Photo_TransferStatus int32

const (
	// The status of this transfer is unspecified.
	Photo_TRANSFER_STATUS_UNKNOWN Photo_TransferStatus = 0
	// This photo has never been in a transfer.
	Photo_NEVER_TRANSFERRED Photo_TransferStatus = 1
	// This photo transfer has been initiated, but the receiver has not yet
	// responded.
	Photo_PENDING Photo_TransferStatus = 2
	// The photo transfer has been completed, and this photo has been
	// transferred to the recipient.
	Photo_COMPLETED Photo_TransferStatus = 3
	// The recipient rejected this photo transfer.
	Photo_REJECTED Photo_TransferStatus = 4
	// The photo transfer expired before the recipient took any action.
	Photo_EXPIRED Photo_TransferStatus = 5
	// The sender cancelled this photo transfer.
	Photo_CANCELLED Photo_TransferStatus = 6
	// The recipient owns this photo due to a rights transfer.
	Photo_RECEIVED_VIA_TRANSFER Photo_TransferStatus = 7
)

var Photo_TransferStatus_name = map[int32]string{
	0: "TRANSFER_STATUS_UNKNOWN",
	1: "NEVER_TRANSFERRED",
	2: "PENDING",
	3: "COMPLETED",
	4: "REJECTED",
	5: "EXPIRED",
	6: "CANCELLED",
	7: "RECEIVED_VIA_TRANSFER",
}

var Photo_TransferStatus_value = map[string]int32{
	"TRANSFER_STATUS_UNKNOWN": 0,
	"NEVER_TRANSFERRED":       1,
	"PENDING":                 2,
	"COMPLETED":               3,
	"REJECTED":                4,
	"EXPIRED":                 5,
	"CANCELLED":               6,
	"RECEIVED_VIA_TRANSFER":   7,
}

func (x Photo_TransferStatus) String() string {
	return proto.EnumName(Photo_TransferStatus_name, int32(x))
}

func (Photo_TransferStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_56f2b10d5439612a, []int{6, 0}
}

// Publication status of the photo in Google Maps.
type Photo_MapsPublishStatus int32

const (
	// The status of the photo is unknown.
	Photo_UNSPECIFIED_MAPS_PUBLISH_STATUS Photo_MapsPublishStatus = 0
	// The photo is published to the public through Google Maps.
	Photo_PUBLISHED Photo_MapsPublishStatus = 1
	// The photo has been rejected for an unknown reason.
	Photo_REJECTED_UNKNOWN Photo_MapsPublishStatus = 2
)

var Photo_MapsPublishStatus_name = map[int32]string{
	0: "UNSPECIFIED_MAPS_PUBLISH_STATUS",
	1: "PUBLISHED",
	2: "REJECTED_UNKNOWN",
}

var Photo_MapsPublishStatus_value = map[string]int32{
	"UNSPECIFIED_MAPS_PUBLISH_STATUS": 0,
	"PUBLISHED":                       1,
	"REJECTED_UNKNOWN":                2,
}

func (x Photo_MapsPublishStatus) String() string {
	return proto.EnumName(Photo_MapsPublishStatus_name, int32(x))
}

func (Photo_MapsPublishStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_56f2b10d5439612a, []int{6, 1}
}

// Upload reference for media files.
type UploadRef struct {
	// Required. An upload reference should be unique for each user. It follows
	// the form:
	// "https://streetviewpublish.googleapis.com/media/user/{account_id}/photo/{upload_reference}"
	UploadUrl            string   `protobuf:"bytes,1,opt,name=upload_url,json=uploadUrl,proto3" json:"upload_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadRef) Reset()         { *m = UploadRef{} }
func (m *UploadRef) String() string { return proto.CompactTextString(m) }
func (*UploadRef) ProtoMessage()    {}
func (*UploadRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f2b10d5439612a, []int{0}
}

func (m *UploadRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRef.Unmarshal(m, b)
}
func (m *UploadRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRef.Marshal(b, m, deterministic)
}
func (m *UploadRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRef.Merge(m, src)
}
func (m *UploadRef) XXX_Size() int {
	return xxx_messageInfo_UploadRef.Size(m)
}
func (m *UploadRef) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRef.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRef proto.InternalMessageInfo

func (m *UploadRef) GetUploadUrl() string {
	if m != nil {
		return m.UploadUrl
	}
	return ""
}

// Identifier for a [Photo][google.streetview.publish.v1.Photo].
type PhotoId struct {
	// Required. A unique identifier for a photo.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhotoId) Reset()         { *m = PhotoId{} }
func (m *PhotoId) String() string { return proto.CompactTextString(m) }
func (*PhotoId) ProtoMessage()    {}
func (*PhotoId) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f2b10d5439612a, []int{1}
}

func (m *PhotoId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhotoId.Unmarshal(m, b)
}
func (m *PhotoId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhotoId.Marshal(b, m, deterministic)
}
func (m *PhotoId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhotoId.Merge(m, src)
}
func (m *PhotoId) XXX_Size() int {
	return xxx_messageInfo_PhotoId.Size(m)
}
func (m *PhotoId) XXX_DiscardUnknown() {
	xxx_messageInfo_PhotoId.DiscardUnknown(m)
}

var xxx_messageInfo_PhotoId proto.InternalMessageInfo

func (m *PhotoId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Level information containing level number and its corresponding name.
type Level struct {
	// Floor number, used for ordering. 0 indicates the ground level, 1 indicates
	// the first level above ground level, -1 indicates the first level under
	// ground level. Non-integer values are OK.
	Number float64 `protobuf:"fixed64,1,opt,name=number,proto3" json:"number,omitempty"`
	// Required. A name assigned to this Level, restricted to 3 characters.
	// Consider how the elevator buttons would be labeled for this level if there
	// was an elevator.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Level) Reset()         { *m = Level{} }
func (m *Level) String() string { return proto.CompactTextString(m) }
func (*Level) ProtoMessage()    {}
func (*Level) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f2b10d5439612a, []int{2}
}

func (m *Level) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Level.Unmarshal(m, b)
}
func (m *Level) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Level.Marshal(b, m, deterministic)
}
func (m *Level) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Level.Merge(m, src)
}
func (m *Level) XXX_Size() int {
	return xxx_messageInfo_Level.Size(m)
}
func (m *Level) XXX_DiscardUnknown() {
	xxx_messageInfo_Level.DiscardUnknown(m)
}

var xxx_messageInfo_Level proto.InternalMessageInfo

func (m *Level) GetNumber() float64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Level) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Raw pose measurement for an entity.
type Pose struct {
	// Latitude and longitude pair of the pose, as explained here:
	// https://cloud.google.com/datastore/docs/reference/rest/Shared.Types/LatLng
	// When creating a [Photo][google.streetview.publish.v1.Photo], if the
	// latitude and longitude pair are not provided, the geolocation from the
	// exif header is used. A latitude and longitude pair not provided in the
	// photo or exif header causes the create photo process to fail.
	LatLngPair *latlng.LatLng `protobuf:"bytes,1,opt,name=lat_lng_pair,json=latLngPair,proto3" json:"lat_lng_pair,omitempty"`
	// Altitude of the pose in meters above WGS84 ellipsoid.
	// NaN indicates an unmeasured quantity.
	Altitude float64 `protobuf:"fixed64,2,opt,name=altitude,proto3" json:"altitude,omitempty"`
	// Compass heading, measured at the center of the photo in degrees clockwise
	// from North. Value must be >=0 and <360.
	// NaN indicates an unmeasured quantity.
	Heading float64 `protobuf:"fixed64,3,opt,name=heading,proto3" json:"heading,omitempty"`
	// Pitch, measured at the center of the photo in degrees. Value must be >=-90
	// and <= 90. A value of -90 means looking directly down, and a value of 90
	// means looking directly up.
	// NaN indicates an unmeasured quantity.
	Pitch float64 `protobuf:"fixed64,4,opt,name=pitch,proto3" json:"pitch,omitempty"`
	// Roll, measured in degrees. Value must be >= 0 and <360. A value of 0
	// means level with the horizon.
	// NaN indicates an unmeasured quantity.
	Roll float64 `protobuf:"fixed64,5,opt,name=roll,proto3" json:"roll,omitempty"`
	// Level (the floor in a building) used to configure vertical navigation.
	Level *Level `protobuf:"bytes,7,opt,name=level,proto3" json:"level,omitempty"`
	// The estimated horizontal accuracy of this pose in meters with 68%
	// confidence (one standard deviation). For example, on Android, this value is
	// available from this method:
	// https://developer.android.com/reference/android/location/Location#getAccuracy().
	// Other platforms have different methods of obtaining similar accuracy
	// estimations.
	AccuracyMeters       float32  `protobuf:"fixed32,9,opt,name=accuracy_meters,json=accuracyMeters,proto3" json:"accuracy_meters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pose) Reset()         { *m = Pose{} }
func (m *Pose) String() string { return proto.CompactTextString(m) }
func (*Pose) ProtoMessage()    {}
func (*Pose) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f2b10d5439612a, []int{3}
}

func (m *Pose) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pose.Unmarshal(m, b)
}
func (m *Pose) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pose.Marshal(b, m, deterministic)
}
func (m *Pose) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pose.Merge(m, src)
}
func (m *Pose) XXX_Size() int {
	return xxx_messageInfo_Pose.Size(m)
}
func (m *Pose) XXX_DiscardUnknown() {
	xxx_messageInfo_Pose.DiscardUnknown(m)
}

var xxx_messageInfo_Pose proto.InternalMessageInfo

func (m *Pose) GetLatLngPair() *latlng.LatLng {
	if m != nil {
		return m.LatLngPair
	}
	return nil
}

func (m *Pose) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Pose) GetHeading() float64 {
	if m != nil {
		return m.Heading
	}
	return 0
}

func (m *Pose) GetPitch() float64 {
	if m != nil {
		return m.Pitch
	}
	return 0
}

func (m *Pose) GetRoll() float64 {
	if m != nil {
		return m.Roll
	}
	return 0
}

func (m *Pose) GetLevel() *Level {
	if m != nil {
		return m.Level
	}
	return nil
}

func (m *Pose) GetAccuracyMeters() float32 {
	if m != nil {
		return m.AccuracyMeters
	}
	return 0
}

// Place metadata for an entity.
type Place struct {
	// Place identifier, as described in
	// https://developers.google.com/places/place-id.
	PlaceId string `protobuf:"bytes,1,opt,name=place_id,json=placeId,proto3" json:"place_id,omitempty"`
	// Output-only. The name of the place, localized to the language_code.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Output-only. The language_code that the name is localized with. This should
	// be the language_code specified in the request, but may be a fallback.
	LanguageCode         string   `protobuf:"bytes,3,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Place) Reset()         { *m = Place{} }
func (m *Place) String() string { return proto.CompactTextString(m) }
func (*Place) ProtoMessage()    {}
func (*Place) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f2b10d5439612a, []int{4}
}

func (m *Place) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Place.Unmarshal(m, b)
}
func (m *Place) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Place.Marshal(b, m, deterministic)
}
func (m *Place) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Place.Merge(m, src)
}
func (m *Place) XXX_Size() int {
	return xxx_messageInfo_Place.Size(m)
}
func (m *Place) XXX_DiscardUnknown() {
	xxx_messageInfo_Place.DiscardUnknown(m)
}

var xxx_messageInfo_Place proto.InternalMessageInfo

func (m *Place) GetPlaceId() string {
	if m != nil {
		return m.PlaceId
	}
	return ""
}

func (m *Place) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Place) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

// A connection is the link from a source photo to a destination photo.
type Connection struct {
	// Required. The destination of the connection from the containing photo to
	// another photo.
	Target               *PhotoId `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}
func (*Connection) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f2b10d5439612a, []int{5}
}

func (m *Connection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connection.Unmarshal(m, b)
}
func (m *Connection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connection.Marshal(b, m, deterministic)
}
func (m *Connection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connection.Merge(m, src)
}
func (m *Connection) XXX_Size() int {
	return xxx_messageInfo_Connection.Size(m)
}
func (m *Connection) XXX_DiscardUnknown() {
	xxx_messageInfo_Connection.DiscardUnknown(m)
}

var xxx_messageInfo_Connection proto.InternalMessageInfo

func (m *Connection) GetTarget() *PhotoId {
	if m != nil {
		return m.Target
	}
	return nil
}

// Photo is used to store 360 photos along with photo metadata.
type Photo struct {
	// Required when updating a photo. Output only when creating a photo.
	// Identifier for the photo, which is unique among all photos in
	// Google.
	PhotoId *PhotoId `protobuf:"bytes,1,opt,name=photo_id,json=photoId,proto3" json:"photo_id,omitempty"`
	// Required when creating a photo. Input only. The resource URL where the
	// photo bytes are uploaded to.
	UploadReference *UploadRef `protobuf:"bytes,2,opt,name=upload_reference,json=uploadReference,proto3" json:"upload_reference,omitempty"`
	// Output only. The download URL for the photo bytes. This field is set only
	// when
	// [GetPhotoRequest.view][google.streetview.publish.v1.GetPhotoRequest.view]
	// is set to
	// [PhotoView.INCLUDE_DOWNLOAD_URL][google.streetview.publish.v1.PhotoView.INCLUDE_DOWNLOAD_URL].
	DownloadUrl string `protobuf:"bytes,3,opt,name=download_url,json=downloadUrl,proto3" json:"download_url,omitempty"`
	// Output only. The thumbnail URL for showing a preview of the given photo.
	ThumbnailUrl string `protobuf:"bytes,9,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
	// Output only. The share link for the photo.
	ShareLink string `protobuf:"bytes,11,opt,name=share_link,json=shareLink,proto3" json:"share_link,omitempty"`
	// Pose of the photo.
	Pose *Pose `protobuf:"bytes,4,opt,name=pose,proto3" json:"pose,omitempty"`
	// Connections to other photos. A connection represents the link from this
	// photo to another photo.
	Connections []*Connection `protobuf:"bytes,5,rep,name=connections,proto3" json:"connections,omitempty"`
	// Absolute time when the photo was captured.
	// When the photo has no exif timestamp, this is used to set a timestamp in
	// the photo metadata.
	CaptureTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=capture_time,json=captureTime,proto3" json:"capture_time,omitempty"`
	// Places where this photo belongs.
	Places []*Place `protobuf:"bytes,7,rep,name=places,proto3" json:"places,omitempty"`
	// Output only. View count of the photo.
	ViewCount int64 `protobuf:"varint,10,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
	// Output only. Status of rights transfer on this photo.
	TransferStatus Photo_TransferStatus `protobuf:"varint,12,opt,name=transfer_status,json=transferStatus,proto3,enum=google.streetview.publish.v1.Photo_TransferStatus" json:"transfer_status,omitempty"`
	// Output only. Status in Google Maps, whether this photo was published or
	// rejected.
	MapsPublishStatus    Photo_MapsPublishStatus `protobuf:"varint,13,opt,name=maps_publish_status,json=mapsPublishStatus,proto3,enum=google.streetview.publish.v1.Photo_MapsPublishStatus" json:"maps_publish_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Photo) Reset()         { *m = Photo{} }
func (m *Photo) String() string { return proto.CompactTextString(m) }
func (*Photo) ProtoMessage()    {}
func (*Photo) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f2b10d5439612a, []int{6}
}

func (m *Photo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Photo.Unmarshal(m, b)
}
func (m *Photo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Photo.Marshal(b, m, deterministic)
}
func (m *Photo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Photo.Merge(m, src)
}
func (m *Photo) XXX_Size() int {
	return xxx_messageInfo_Photo.Size(m)
}
func (m *Photo) XXX_DiscardUnknown() {
	xxx_messageInfo_Photo.DiscardUnknown(m)
}

var xxx_messageInfo_Photo proto.InternalMessageInfo

func (m *Photo) GetPhotoId() *PhotoId {
	if m != nil {
		return m.PhotoId
	}
	return nil
}

func (m *Photo) GetUploadReference() *UploadRef {
	if m != nil {
		return m.UploadReference
	}
	return nil
}

func (m *Photo) GetDownloadUrl() string {
	if m != nil {
		return m.DownloadUrl
	}
	return ""
}

func (m *Photo) GetThumbnailUrl() string {
	if m != nil {
		return m.ThumbnailUrl
	}
	return ""
}

func (m *Photo) GetShareLink() string {
	if m != nil {
		return m.ShareLink
	}
	return ""
}

func (m *Photo) GetPose() *Pose {
	if m != nil {
		return m.Pose
	}
	return nil
}

func (m *Photo) GetConnections() []*Connection {
	if m != nil {
		return m.Connections
	}
	return nil
}

func (m *Photo) GetCaptureTime() *timestamp.Timestamp {
	if m != nil {
		return m.CaptureTime
	}
	return nil
}

func (m *Photo) GetPlaces() []*Place {
	if m != nil {
		return m.Places
	}
	return nil
}

func (m *Photo) GetViewCount() int64 {
	if m != nil {
		return m.ViewCount
	}
	return 0
}

func (m *Photo) GetTransferStatus() Photo_TransferStatus {
	if m != nil {
		return m.TransferStatus
	}
	return Photo_TRANSFER_STATUS_UNKNOWN
}

func (m *Photo) GetMapsPublishStatus() Photo_MapsPublishStatus {
	if m != nil {
		return m.MapsPublishStatus
	}
	return Photo_UNSPECIFIED_MAPS_PUBLISH_STATUS
}

func init() {
	proto.RegisterEnum("google.streetview.publish.v1.Photo_TransferStatus", Photo_TransferStatus_name, Photo_TransferStatus_value)
	proto.RegisterEnum("google.streetview.publish.v1.Photo_MapsPublishStatus", Photo_MapsPublishStatus_name, Photo_MapsPublishStatus_value)
	proto.RegisterType((*UploadRef)(nil), "google.streetview.publish.v1.UploadRef")
	proto.RegisterType((*PhotoId)(nil), "google.streetview.publish.v1.PhotoId")
	proto.RegisterType((*Level)(nil), "google.streetview.publish.v1.Level")
	proto.RegisterType((*Pose)(nil), "google.streetview.publish.v1.Pose")
	proto.RegisterType((*Place)(nil), "google.streetview.publish.v1.Place")
	proto.RegisterType((*Connection)(nil), "google.streetview.publish.v1.Connection")
	proto.RegisterType((*Photo)(nil), "google.streetview.publish.v1.Photo")
}

func init() {
	proto.RegisterFile("google/streetview/publish/v1/resources.proto", fileDescriptor_56f2b10d5439612a)
}

var fileDescriptor_56f2b10d5439612a = []byte{
	// 940 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x6f, 0x6f, 0xdb, 0xb6,
	0x13, 0xae, 0x9d, 0xd8, 0x8e, 0xcf, 0x4e, 0xe2, 0xb0, 0xed, 0xef, 0xa7, 0x64, 0x2b, 0x9a, 0x29,
	0x18, 0x6a, 0x0c, 0x83, 0x84, 0xba, 0xe8, 0x80, 0xa1, 0x28, 0xb0, 0xc4, 0x56, 0x37, 0xb7, 0x8e,
	0x2b, 0xd0, 0x76, 0x36, 0xac, 0x18, 0x04, 0x5a, 0x62, 0x64, 0xa1, 0x34, 0x29, 0x50, 0x54, 0x82,
	0xbe, 0xdc, 0xeb, 0x7d, 0x8e, 0x7d, 0xbd, 0x7d, 0x86, 0x41, 0x14, 0xe5, 0xae, 0x5b, 0x90, 0xe6,
	0x95, 0xef, 0x9e, 0x7b, 0x1e, 0xde, 0x1f, 0x9e, 0x29, 0xf8, 0x36, 0x16, 0x22, 0x66, 0xd4, 0xcd,
	0x94, 0xa4, 0x54, 0x5d, 0x25, 0xf4, 0xda, 0x4d, 0xf3, 0x25, 0x4b, 0xb2, 0x95, 0x7b, 0xf5, 0xd4,
	0x95, 0x34, 0x13, 0xb9, 0x0c, 0x69, 0xe6, 0xa4, 0x52, 0x28, 0x81, 0xbe, 0x2c, 0xd9, 0xce, 0x47,
	0xb6, 0x63, 0xd8, 0xce, 0xd5, 0xd3, 0x23, 0x13, 0x75, 0x49, 0x9a, 0xb8, 0x84, 0x73, 0xa1, 0x88,
	0x4a, 0x04, 0x37, 0xda, 0xa3, 0xc7, 0x26, 0xaa, 0xbd, 0x65, 0x7e, 0xe9, 0xaa, 0x64, 0x4d, 0x33,
	0x45, 0xd6, 0xa9, 0x21, 0x58, 0x86, 0xa0, 0x3e, 0xa4, 0xd4, 0x65, 0x44, 0x31, 0x1e, 0x97, 0x11,
	0xfb, 0x1b, 0x68, 0x2f, 0x52, 0x26, 0x48, 0x84, 0xe9, 0x25, 0x7a, 0x04, 0x90, 0x6b, 0x27, 0xc8,
	0x25, 0xb3, 0x6a, 0xc7, 0xb5, 0x7e, 0x1b, 0xb7, 0x4b, 0x64, 0x21, 0x99, 0x7d, 0x08, 0x2d, 0x7f,
	0x25, 0x94, 0x18, 0x47, 0x68, 0x0f, 0xea, 0x49, 0x64, 0x18, 0xf5, 0x24, 0xb2, 0x9f, 0x41, 0x63,
	0x42, 0xaf, 0x28, 0x43, 0xff, 0x83, 0x26, 0xcf, 0xd7, 0x4b, 0x2a, 0x75, 0xb0, 0x86, 0x8d, 0x87,
	0x10, 0x6c, 0x73, 0xb2, 0xa6, 0x56, 0x5d, 0x4b, 0xb4, 0x6d, 0xff, 0x5e, 0x87, 0x6d, 0x5f, 0x64,
	0x14, 0x3d, 0x87, 0x2e, 0x23, 0x2a, 0x60, 0x3c, 0x0e, 0x52, 0x92, 0x94, 0xd2, 0xce, 0xe0, 0xbe,
	0x63, 0x46, 0x52, 0x54, 0xed, 0x4c, 0x88, 0x9a, 0xf0, 0x18, 0x03, 0xd3, 0xbf, 0x3e, 0x49, 0x24,
	0x3a, 0x82, 0x1d, 0xc2, 0x54, 0xa2, 0xf2, 0xa8, 0x3c, 0xb7, 0x86, 0x37, 0x3e, 0xb2, 0xa0, 0xb5,
	0xa2, 0x24, 0x4a, 0x78, 0x6c, 0x6d, 0xe9, 0x50, 0xe5, 0xa2, 0x07, 0xd0, 0x48, 0x13, 0x15, 0xae,
	0xac, 0x6d, 0x8d, 0x97, 0x4e, 0x51, 0x9f, 0x14, 0x8c, 0x59, 0x0d, 0x0d, 0x6a, 0x1b, 0x7d, 0x0f,
	0x0d, 0x56, 0x34, 0x65, 0xb5, 0x74, 0x3d, 0x27, 0xce, 0x6d, 0x57, 0xe4, 0xe8, 0xfe, 0x71, 0xa9,
	0x40, 0x4f, 0x60, 0x9f, 0x84, 0x61, 0x2e, 0x49, 0xf8, 0x21, 0x58, 0x53, 0x45, 0x65, 0x66, 0xb5,
	0x8f, 0x6b, 0xfd, 0x3a, 0xde, 0xab, 0xe0, 0x73, 0x8d, 0xda, 0xef, 0xa0, 0xe1, 0x33, 0x12, 0x52,
	0x74, 0x08, 0x3b, 0x69, 0x61, 0x04, 0x9b, 0xb9, 0xb6, 0xb4, 0x3f, 0x8e, 0x6e, 0x9a, 0x1d, 0x3a,
	0x81, 0x5d, 0x46, 0x78, 0x9c, 0x93, 0x98, 0x06, 0xa1, 0x88, 0xa8, 0xee, 0xb2, 0x8d, 0xbb, 0x15,
	0x38, 0x14, 0x11, 0xb5, 0xdf, 0x00, 0x0c, 0x05, 0xe7, 0x34, 0x2c, 0x96, 0x05, 0xbd, 0x84, 0xa6,
	0x22, 0x32, 0xa6, 0xca, 0xcc, 0xf7, 0xeb, 0xdb, 0xfb, 0x31, 0x57, 0x8d, 0x8d, 0xc8, 0xfe, 0xab,
	0x05, 0x0d, 0x8d, 0xa1, 0x1f, 0x60, 0x27, 0x2d, 0x8c, 0xaa, 0xd4, 0x3b, 0x1f, 0xd5, 0x4a, 0xcd,
	0xfa, 0x60, 0xe8, 0x99, 0x45, 0x93, 0xf4, 0x92, 0x4a, 0xca, 0xc3, 0xb2, 0xbb, 0xce, 0xe0, 0xc9,
	0xed, 0x27, 0x6d, 0x76, 0x15, 0xef, 0xe7, 0x95, 0x59, 0xea, 0xd1, 0x57, 0xd0, 0x8d, 0xc4, 0x35,
	0xdf, 0xac, 0x6f, 0x39, 0x90, 0x4e, 0x85, 0x2d, 0x24, 0x2b, 0x86, 0xa6, 0x56, 0xf9, 0x7a, 0xc9,
	0x49, 0xc2, 0x34, 0xa7, 0x5d, 0x0e, 0x6d, 0x03, 0x16, 0xa4, 0x47, 0x00, 0xd9, 0x8a, 0x48, 0x1a,
	0xb0, 0x84, 0xbf, 0xb7, 0x3a, 0xe5, 0x9f, 0x40, 0x23, 0x93, 0x84, 0xbf, 0x47, 0xdf, 0xc1, 0x76,
	0x2a, 0x32, 0xaa, 0xb7, 0xa7, 0x33, 0xb0, 0x3f, 0xd3, 0xb8, 0xc8, 0x28, 0xd6, 0x7c, 0xf4, 0x1a,
	0x3a, 0xe1, 0xe6, 0x2e, 0x32, 0xab, 0x71, 0xbc, 0xd5, 0xef, 0x0c, 0xfa, 0xb7, 0xcb, 0x3f, 0x5e,
	0x1e, 0xfe, 0xa7, 0x18, 0xbd, 0x84, 0x6e, 0x48, 0x52, 0x95, 0x4b, 0x1a, 0x14, 0xff, 0x74, 0xab,
	0xa9, 0x6b, 0x39, 0xaa, 0x0e, 0xab, 0x9e, 0x01, 0x67, 0x5e, 0x3d, 0x03, 0xb8, 0x63, 0xf8, 0x05,
	0x82, 0x5e, 0x40, 0x53, 0xaf, 0x56, 0x66, 0xb5, 0x74, 0x15, 0x9f, 0x59, 0x6c, 0xbd, 0x9f, 0xd8,
	0x48, 0x8a, 0xf1, 0x14, 0x84, 0x20, 0x14, 0x39, 0x57, 0x16, 0x1c, 0xd7, 0xfa, 0x5b, 0xb8, 0x5d,
	0x20, 0xc3, 0x02, 0x40, 0xef, 0x60, 0x5f, 0x49, 0xc2, 0xb3, 0x4b, 0x2a, 0x83, 0x4c, 0x11, 0x95,
	0x67, 0x56, 0xf7, 0xb8, 0xd6, 0xdf, 0x1b, 0x0c, 0xee, 0xb0, 0x22, 0xce, 0xdc, 0x48, 0x67, 0x5a,
	0x89, 0xf7, 0xd4, 0x27, 0x3e, 0xa2, 0x70, 0x7f, 0x4d, 0xd2, 0x2c, 0x30, 0xba, 0x2a, 0xc1, 0xae,
	0x4e, 0xf0, 0xfc, 0x2e, 0x09, 0xce, 0x49, 0x9a, 0xf9, 0x25, 0x68, 0x72, 0x1c, 0xac, 0xff, 0x0d,
	0xd9, 0x7f, 0xd6, 0x60, 0xef, 0xd3, 0x4a, 0xd0, 0x17, 0xf0, 0xff, 0x39, 0x3e, 0x9d, 0xce, 0x5e,
	0x79, 0x38, 0x98, 0xcd, 0x4f, 0xe7, 0x8b, 0x59, 0xb0, 0x98, 0xbe, 0x99, 0xbe, 0xfd, 0x79, 0xda,
	0xbb, 0x87, 0x1e, 0xc2, 0xc1, 0xd4, 0xbb, 0xf0, 0x70, 0x50, 0x51, 0xb0, 0x37, 0xea, 0xd5, 0x50,
	0x07, 0x5a, 0xbe, 0x37, 0x1d, 0x8d, 0xa7, 0x3f, 0xf6, 0xea, 0x68, 0x17, 0xda, 0xc3, 0xb7, 0xe7,
	0xfe, 0xc4, 0x9b, 0x7b, 0xa3, 0xde, 0x16, 0xea, 0xc2, 0x0e, 0xf6, 0x5e, 0x7b, 0xc3, 0xc2, 0xdb,
	0x2e, 0x98, 0xde, 0x2f, 0xfe, 0xb8, 0x90, 0x35, 0x34, 0xf3, 0x74, 0x3a, 0xf4, 0x26, 0x13, 0x6f,
	0xd4, 0x6b, 0xa2, 0x43, 0x78, 0x88, 0xbd, 0xa1, 0x37, 0xbe, 0xf0, 0x46, 0xc1, 0xc5, 0xf8, 0x74,
	0x93, 0xa3, 0xd7, 0xb2, 0x7f, 0x83, 0x83, 0xff, 0xf4, 0x83, 0x4e, 0xe0, 0xf1, 0x62, 0x3a, 0xf3,
	0xbd, 0xe1, 0xf8, 0xd5, 0xd8, 0x1b, 0x05, 0xe7, 0xa7, 0xfe, 0x2c, 0xf0, 0x17, 0x67, 0x93, 0xf1,
	0xec, 0x27, 0x53, 0x79, 0xef, 0x5e, 0x91, 0xc3, 0x60, 0xba, 0xd2, 0x07, 0xd0, 0xab, 0xaa, 0xd9,
	0xb4, 0x55, 0x3f, 0xfb, 0xa3, 0x06, 0xfd, 0x50, 0xac, 0xab, 0xb1, 0xc6, 0x54, 0x38, 0x79, 0x1c,
	0xde, 0x3c, 0xde, 0xb3, 0xa3, 0x99, 0x86, 0x2f, 0x12, 0x7a, 0x6d, 0xea, 0xc1, 0xd5, 0x07, 0xee,
	0xd7, 0x61, 0x75, 0x82, 0x28, 0x5e, 0x27, 0x47, 0xc8, 0xd8, 0x8d, 0x29, 0xd7, 0x5b, 0xea, 0x96,
	0x21, 0x92, 0x26, 0xd9, 0xcd, 0xdf, 0xc9, 0x17, 0xc6, 0x5c, 0x36, 0x35, 0xff, 0xd9, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x6d, 0x23, 0x7d, 0x07, 0x56, 0x07, 0x00, 0x00,
}
