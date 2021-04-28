// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/extension_feed_item_error.proto

package errors

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

// Enum describing possible extension feed item errors.
type ExtensionFeedItemErrorEnum_ExtensionFeedItemError int32

const (
	// Enum unspecified.
	ExtensionFeedItemErrorEnum_UNSPECIFIED ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 0
	// The received error code is not known in this version.
	ExtensionFeedItemErrorEnum_UNKNOWN ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 1
	// Value is not within the accepted range.
	ExtensionFeedItemErrorEnum_VALUE_OUT_OF_RANGE ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 2
	// Url list is too long.
	ExtensionFeedItemErrorEnum_URL_LIST_TOO_LONG ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 3
	// Cannot have a geo targeting restriction without having geo targeting.
	ExtensionFeedItemErrorEnum_CANNOT_HAVE_RESTRICTION_ON_EMPTY_GEO_TARGETING ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 4
	// Cannot simultaneously set sitelink field with final urls.
	ExtensionFeedItemErrorEnum_CANNOT_SET_WITH_FINAL_URLS ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 5
	// Must set field with final urls.
	ExtensionFeedItemErrorEnum_CANNOT_SET_WITHOUT_FINAL_URLS ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 6
	// Phone number for a call extension is invalid.
	ExtensionFeedItemErrorEnum_INVALID_PHONE_NUMBER ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 7
	// Phone number for a call extension is not supported for the given country
	// code.
	ExtensionFeedItemErrorEnum_PHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 8
	// A carrier specific number in short format is not allowed for call
	// extensions.
	ExtensionFeedItemErrorEnum_CARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 9
	// Premium rate numbers are not allowed for call extensions.
	ExtensionFeedItemErrorEnum_PREMIUM_RATE_NUMBER_NOT_ALLOWED ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 10
	// Phone number type for a call extension is not allowed.
	// For example, personal number is not allowed for a call extension in
	// most regions.
	ExtensionFeedItemErrorEnum_DISALLOWED_NUMBER_TYPE ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 11
	// Phone number for a call extension does not meet domestic format
	// requirements.
	ExtensionFeedItemErrorEnum_INVALID_DOMESTIC_PHONE_NUMBER_FORMAT ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 12
	// Vanity phone numbers (i.e. those including letters) are not allowed for
	// call extensions.
	ExtensionFeedItemErrorEnum_VANITY_PHONE_NUMBER_NOT_ALLOWED ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 13
	// Call conversion action provided for a call extension is invalid.
	ExtensionFeedItemErrorEnum_INVALID_CALL_CONVERSION_ACTION ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 14
	// For a call extension, the customer is not whitelisted for call tracking.
	ExtensionFeedItemErrorEnum_CUSTOMER_NOT_WHITELISTED_FOR_CALLTRACKING ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 15
	// Call tracking is not supported for the given country for a call
	// extension.
	ExtensionFeedItemErrorEnum_CALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 16
	// Customer hasn't consented for call recording, which is required for
	// creating/updating call feed items. Please see
	// https://support.google.com/google-ads/answer/7412639.
	ExtensionFeedItemErrorEnum_CUSTOMER_CONSENT_FOR_CALL_RECORDING_REQUIRED ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 17
	// App id provided for an app extension is invalid.
	ExtensionFeedItemErrorEnum_INVALID_APP_ID ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 18
	// Quotation marks present in the review text for a review extension.
	ExtensionFeedItemErrorEnum_QUOTES_IN_REVIEW_EXTENSION_SNIPPET ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 19
	// Hyphen character present in the review text for a review extension.
	ExtensionFeedItemErrorEnum_HYPHENS_IN_REVIEW_EXTENSION_SNIPPET ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 20
	// A blacklisted review source name or url was provided for a review
	// extension.
	ExtensionFeedItemErrorEnum_REVIEW_EXTENSION_SOURCE_INELIGIBLE ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 21
	// Review source name should not be found in the review text.
	ExtensionFeedItemErrorEnum_SOURCE_NAME_IN_REVIEW_EXTENSION_TEXT ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 22
	// Inconsistent currency codes.
	ExtensionFeedItemErrorEnum_INCONSISTENT_CURRENCY_CODES ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 23
	// Price extension cannot have duplicated headers.
	ExtensionFeedItemErrorEnum_PRICE_EXTENSION_HAS_DUPLICATED_HEADERS ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 24
	// Price item cannot have duplicated header and description.
	ExtensionFeedItemErrorEnum_PRICE_ITEM_HAS_DUPLICATED_HEADER_AND_DESCRIPTION ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 25
	// Price extension has too few items.
	ExtensionFeedItemErrorEnum_PRICE_EXTENSION_HAS_TOO_FEW_ITEMS ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 26
	// Price extension has too many items.
	ExtensionFeedItemErrorEnum_PRICE_EXTENSION_HAS_TOO_MANY_ITEMS ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 27
	// The input value is not currently supported.
	ExtensionFeedItemErrorEnum_UNSUPPORTED_VALUE ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 28
	// The input value is not currently supported in the selected language of an
	// extension.
	ExtensionFeedItemErrorEnum_UNSUPPORTED_VALUE_IN_SELECTED_LANGUAGE ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 29
	// Unknown or unsupported device preference.
	ExtensionFeedItemErrorEnum_INVALID_DEVICE_PREFERENCE ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 30
	// Invalid feed item schedule end time (i.e., endHour = 24 and endMinute !=
	// 0).
	ExtensionFeedItemErrorEnum_INVALID_SCHEDULE_END ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 31
	// Date time zone does not match the account's time zone.
	ExtensionFeedItemErrorEnum_DATE_TIME_MUST_BE_IN_ACCOUNT_TIME_ZONE ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 32
	// Invalid structured snippet header.
	ExtensionFeedItemErrorEnum_INVALID_SNIPPETS_HEADER ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 33
	// Cannot operate on removed feed item.
	ExtensionFeedItemErrorEnum_CANNOT_OPERATE_ON_REMOVED_FEED_ITEM ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 34
	// Phone number not supported when call tracking enabled for country.
	ExtensionFeedItemErrorEnum_PHONE_NUMBER_NOT_SUPPORTED_WITH_CALLTRACKING_FOR_COUNTRY ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 35
	// Cannot set call_conversion_action while call_conversion_tracking_enabled
	// is set to true.
	ExtensionFeedItemErrorEnum_CONFLICTING_CALL_CONVERSION_SETTINGS ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 36
	// The type of the input extension feed item doesn't match the existing
	// extension feed item.
	ExtensionFeedItemErrorEnum_EXTENSION_TYPE_MISMATCH ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 37
	// The oneof field extension i.e. subtype of extension feed item is
	// required.
	ExtensionFeedItemErrorEnum_EXTENSION_SUBTYPE_REQUIRED ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 38
	// The referenced feed item is not mapped to a supported extension type.
	ExtensionFeedItemErrorEnum_EXTENSION_TYPE_UNSUPPORTED ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 39
	// Cannot operate on a Feed with more than one active FeedMapping.
	ExtensionFeedItemErrorEnum_CANNOT_OPERATE_ON_FEED_WITH_MULTIPLE_MAPPINGS ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 40
	// Cannot operate on a Feed that has key attributes.
	ExtensionFeedItemErrorEnum_CANNOT_OPERATE_ON_FEED_WITH_KEY_ATTRIBUTES ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 41
	// Input price is not in a valid format.
	ExtensionFeedItemErrorEnum_INVALID_PRICE_FORMAT ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 42
	// The promotion time is invalid.
	ExtensionFeedItemErrorEnum_PROMOTION_INVALID_TIME ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 43
	// This field has too many decimal places specified.
	ExtensionFeedItemErrorEnum_TOO_MANY_DECIMAL_PLACES_SPECIFIED ExtensionFeedItemErrorEnum_ExtensionFeedItemError = 44
)

var ExtensionFeedItemErrorEnum_ExtensionFeedItemError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "VALUE_OUT_OF_RANGE",
	3:  "URL_LIST_TOO_LONG",
	4:  "CANNOT_HAVE_RESTRICTION_ON_EMPTY_GEO_TARGETING",
	5:  "CANNOT_SET_WITH_FINAL_URLS",
	6:  "CANNOT_SET_WITHOUT_FINAL_URLS",
	7:  "INVALID_PHONE_NUMBER",
	8:  "PHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY",
	9:  "CARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED",
	10: "PREMIUM_RATE_NUMBER_NOT_ALLOWED",
	11: "DISALLOWED_NUMBER_TYPE",
	12: "INVALID_DOMESTIC_PHONE_NUMBER_FORMAT",
	13: "VANITY_PHONE_NUMBER_NOT_ALLOWED",
	14: "INVALID_CALL_CONVERSION_ACTION",
	15: "CUSTOMER_NOT_WHITELISTED_FOR_CALLTRACKING",
	16: "CALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY",
	17: "CUSTOMER_CONSENT_FOR_CALL_RECORDING_REQUIRED",
	18: "INVALID_APP_ID",
	19: "QUOTES_IN_REVIEW_EXTENSION_SNIPPET",
	20: "HYPHENS_IN_REVIEW_EXTENSION_SNIPPET",
	21: "REVIEW_EXTENSION_SOURCE_INELIGIBLE",
	22: "SOURCE_NAME_IN_REVIEW_EXTENSION_TEXT",
	23: "INCONSISTENT_CURRENCY_CODES",
	24: "PRICE_EXTENSION_HAS_DUPLICATED_HEADERS",
	25: "PRICE_ITEM_HAS_DUPLICATED_HEADER_AND_DESCRIPTION",
	26: "PRICE_EXTENSION_HAS_TOO_FEW_ITEMS",
	27: "PRICE_EXTENSION_HAS_TOO_MANY_ITEMS",
	28: "UNSUPPORTED_VALUE",
	29: "UNSUPPORTED_VALUE_IN_SELECTED_LANGUAGE",
	30: "INVALID_DEVICE_PREFERENCE",
	31: "INVALID_SCHEDULE_END",
	32: "DATE_TIME_MUST_BE_IN_ACCOUNT_TIME_ZONE",
	33: "INVALID_SNIPPETS_HEADER",
	34: "CANNOT_OPERATE_ON_REMOVED_FEED_ITEM",
	35: "PHONE_NUMBER_NOT_SUPPORTED_WITH_CALLTRACKING_FOR_COUNTRY",
	36: "CONFLICTING_CALL_CONVERSION_SETTINGS",
	37: "EXTENSION_TYPE_MISMATCH",
	38: "EXTENSION_SUBTYPE_REQUIRED",
	39: "EXTENSION_TYPE_UNSUPPORTED",
	40: "CANNOT_OPERATE_ON_FEED_WITH_MULTIPLE_MAPPINGS",
	41: "CANNOT_OPERATE_ON_FEED_WITH_KEY_ATTRIBUTES",
	42: "INVALID_PRICE_FORMAT",
	43: "PROMOTION_INVALID_TIME",
	44: "TOO_MANY_DECIMAL_PLACES_SPECIFIED",
}

var ExtensionFeedItemErrorEnum_ExtensionFeedItemError_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"VALUE_OUT_OF_RANGE": 2,
	"URL_LIST_TOO_LONG":  3,
	"CANNOT_HAVE_RESTRICTION_ON_EMPTY_GEO_TARGETING":           4,
	"CANNOT_SET_WITH_FINAL_URLS":                               5,
	"CANNOT_SET_WITHOUT_FINAL_URLS":                            6,
	"INVALID_PHONE_NUMBER":                                     7,
	"PHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY":                   8,
	"CARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED":                9,
	"PREMIUM_RATE_NUMBER_NOT_ALLOWED":                          10,
	"DISALLOWED_NUMBER_TYPE":                                   11,
	"INVALID_DOMESTIC_PHONE_NUMBER_FORMAT":                     12,
	"VANITY_PHONE_NUMBER_NOT_ALLOWED":                          13,
	"INVALID_CALL_CONVERSION_ACTION":                           14,
	"CUSTOMER_NOT_WHITELISTED_FOR_CALLTRACKING":                15,
	"CALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY":                   16,
	"CUSTOMER_CONSENT_FOR_CALL_RECORDING_REQUIRED":             17,
	"INVALID_APP_ID":                                           18,
	"QUOTES_IN_REVIEW_EXTENSION_SNIPPET":                       19,
	"HYPHENS_IN_REVIEW_EXTENSION_SNIPPET":                      20,
	"REVIEW_EXTENSION_SOURCE_INELIGIBLE":                       21,
	"SOURCE_NAME_IN_REVIEW_EXTENSION_TEXT":                     22,
	"INCONSISTENT_CURRENCY_CODES":                              23,
	"PRICE_EXTENSION_HAS_DUPLICATED_HEADERS":                   24,
	"PRICE_ITEM_HAS_DUPLICATED_HEADER_AND_DESCRIPTION":         25,
	"PRICE_EXTENSION_HAS_TOO_FEW_ITEMS":                        26,
	"PRICE_EXTENSION_HAS_TOO_MANY_ITEMS":                       27,
	"UNSUPPORTED_VALUE":                                        28,
	"UNSUPPORTED_VALUE_IN_SELECTED_LANGUAGE":                   29,
	"INVALID_DEVICE_PREFERENCE":                                30,
	"INVALID_SCHEDULE_END":                                     31,
	"DATE_TIME_MUST_BE_IN_ACCOUNT_TIME_ZONE":                   32,
	"INVALID_SNIPPETS_HEADER":                                  33,
	"CANNOT_OPERATE_ON_REMOVED_FEED_ITEM":                      34,
	"PHONE_NUMBER_NOT_SUPPORTED_WITH_CALLTRACKING_FOR_COUNTRY": 35,
	"CONFLICTING_CALL_CONVERSION_SETTINGS":                     36,
	"EXTENSION_TYPE_MISMATCH":                                  37,
	"EXTENSION_SUBTYPE_REQUIRED":                               38,
	"EXTENSION_TYPE_UNSUPPORTED":                               39,
	"CANNOT_OPERATE_ON_FEED_WITH_MULTIPLE_MAPPINGS":            40,
	"CANNOT_OPERATE_ON_FEED_WITH_KEY_ATTRIBUTES":               41,
	"INVALID_PRICE_FORMAT":                                     42,
	"PROMOTION_INVALID_TIME":                                   43,
	"TOO_MANY_DECIMAL_PLACES_SPECIFIED":                        44,
}

func (x ExtensionFeedItemErrorEnum_ExtensionFeedItemError) String() string {
	return proto.EnumName(ExtensionFeedItemErrorEnum_ExtensionFeedItemError_name, int32(x))
}

func (ExtensionFeedItemErrorEnum_ExtensionFeedItemError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_881736da1600dac9, []int{0, 0}
}

// Container for enum describing possible extension feed item error.
type ExtensionFeedItemErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtensionFeedItemErrorEnum) Reset()         { *m = ExtensionFeedItemErrorEnum{} }
func (m *ExtensionFeedItemErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ExtensionFeedItemErrorEnum) ProtoMessage()    {}
func (*ExtensionFeedItemErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_881736da1600dac9, []int{0}
}

func (m *ExtensionFeedItemErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionFeedItemErrorEnum.Unmarshal(m, b)
}
func (m *ExtensionFeedItemErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionFeedItemErrorEnum.Marshal(b, m, deterministic)
}
func (m *ExtensionFeedItemErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionFeedItemErrorEnum.Merge(m, src)
}
func (m *ExtensionFeedItemErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ExtensionFeedItemErrorEnum.Size(m)
}
func (m *ExtensionFeedItemErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionFeedItemErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionFeedItemErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.ExtensionFeedItemErrorEnum_ExtensionFeedItemError", ExtensionFeedItemErrorEnum_ExtensionFeedItemError_name, ExtensionFeedItemErrorEnum_ExtensionFeedItemError_value)
	proto.RegisterType((*ExtensionFeedItemErrorEnum)(nil), "google.ads.googleads.v2.errors.ExtensionFeedItemErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/extension_feed_item_error.proto", fileDescriptor_881736da1600dac9)
}

var fileDescriptor_881736da1600dac9 = []byte{
	// 1056 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xdd, 0x8e, 0x53, 0x37,
	0x10, 0x2e, 0x4b, 0x0b, 0xad, 0x29, 0x60, 0x5c, 0x58, 0x60, 0x17, 0x96, 0x12, 0xfe, 0x29, 0x24,
	0x34, 0xed, 0x45, 0x95, 0x56, 0x95, 0x1c, 0x9f, 0x49, 0x62, 0xe1, 0x63, 0x1b, 0xdb, 0x27, 0x21,
	0x68, 0xa5, 0xd1, 0xb6, 0x49, 0xa3, 0x95, 0xd8, 0x64, 0xb5, 0x49, 0x51, 0x9f, 0xa2, 0x0f, 0xd1,
	0xcb, 0x4a, 0x7d, 0x91, 0x3e, 0x4a, 0x1f, 0xa0, 0xd7, 0x95, 0x7d, 0x92, 0x90, 0xed, 0x86, 0xbd,
	0xca, 0xc9, 0xf8, 0x9b, 0xff, 0x6f, 0x66, 0xc8, 0x8f, 0xa3, 0xc9, 0x64, 0xf4, 0x76, 0x58, 0xdb,
	0x1b, 0x4c, 0x6b, 0xe5, 0x67, 0xfc, 0x7a, 0x57, 0xaf, 0x0d, 0x8f, 0x8e, 0x26, 0x47, 0xd3, 0xda,
	0xf0, 0xb7, 0xd9, 0x70, 0x3c, 0xdd, 0x9f, 0x8c, 0xf1, 0x97, 0xe1, 0x70, 0x80, 0xfb, 0xb3, 0xe1,
	0x01, 0xa6, 0xa7, 0xea, 0xe1, 0xd1, 0x64, 0x36, 0x61, 0x3b, 0xa5, 0x52, 0x75, 0x6f, 0x30, 0xad,
	0x2e, 0xf5, 0xab, 0xef, 0xea, 0xd5, 0x52, 0x7f, 0xeb, 0xd6, 0xc2, 0xfe, 0xe1, 0x7e, 0x6d, 0x6f,
	0x3c, 0x9e, 0xcc, 0xf6, 0x66, 0xfb, 0x93, 0xf1, 0xb4, 0xd4, 0xae, 0xfc, 0x75, 0x91, 0x6c, 0xc1,
	0xc2, 0x43, 0x6b, 0x38, 0x1c, 0xc8, 0xd9, 0xf0, 0x00, 0xa2, 0x26, 0x8c, 0x7f, 0x3d, 0xa8, 0xfc,
	0x7e, 0x91, 0x6c, 0xae, 0x7f, 0x66, 0x97, 0xc9, 0x85, 0x42, 0x7b, 0x0b, 0x42, 0xb6, 0x24, 0x64,
	0xf4, 0x23, 0x76, 0x81, 0x9c, 0x2f, 0xf4, 0x4b, 0x6d, 0x7a, 0x9a, 0x9e, 0x61, 0x9b, 0x84, 0x75,
	0xb9, 0x2a, 0x00, 0x4d, 0x11, 0xd0, 0xb4, 0xd0, 0x71, 0xdd, 0x06, 0xba, 0xc1, 0xae, 0x91, 0x2b,
	0x85, 0x53, 0xa8, 0xa4, 0x0f, 0x18, 0x8c, 0x41, 0x65, 0x74, 0x9b, 0x9e, 0x65, 0x75, 0x52, 0x15,
	0x5c, 0x6b, 0x13, 0xb0, 0xc3, 0xbb, 0x80, 0x0e, 0x7c, 0x70, 0x52, 0x04, 0x69, 0x34, 0x1a, 0x8d,
	0x90, 0xdb, 0xd0, 0xc7, 0x36, 0x18, 0x0c, 0xdc, 0xb5, 0x21, 0x48, 0xdd, 0xa6, 0x1f, 0xb3, 0x1d,
	0xb2, 0x35, 0xd7, 0xf1, 0x10, 0xb0, 0x27, 0x43, 0x07, 0x5b, 0x52, 0x73, 0x85, 0x85, 0x53, 0x9e,
	0x7e, 0xc2, 0xee, 0x92, 0xdb, 0xff, 0x7b, 0x8f, 0xb1, 0xac, 0x40, 0xce, 0xb1, 0x1b, 0xe4, 0xaa,
	0xd4, 0x5d, 0xae, 0x64, 0x86, 0xb6, 0x63, 0x34, 0xa0, 0x2e, 0xf2, 0x26, 0x38, 0x7a, 0x9e, 0x3d,
	0x25, 0x0f, 0x57, 0x25, 0x98, 0xcc, 0x14, 0xd6, 0x1a, 0x17, 0x20, 0xc3, 0x96, 0x71, 0x28, 0x4c,
	0xa1, 0x83, 0xeb, 0xd3, 0x4f, 0xd9, 0x73, 0xf2, 0x44, 0x70, 0xe7, 0x24, 0x38, 0x9c, 0xd7, 0x43,
	0xa0, 0xef, 0x18, 0x17, 0x56, 0x95, 0xb9, 0x52, 0xa6, 0x07, 0x19, 0xfd, 0x8c, 0xdd, 0x23, 0x77,
	0xac, 0x83, 0x5c, 0x16, 0x39, 0x3a, 0x1e, 0x60, 0x1d, 0x88, 0xb0, 0x2d, 0xb2, 0x99, 0x49, 0x3f,
	0xff, 0xbf, 0x80, 0x84, 0xbe, 0x05, 0x7a, 0x81, 0x3d, 0x26, 0xf7, 0x17, 0x51, 0x67, 0x26, 0x07,
	0x1f, 0xa4, 0x38, 0x16, 0x7e, 0x0c, 0x2f, 0xe7, 0x81, 0x7e, 0x1e, 0x5d, 0x75, 0xb9, 0x96, 0xa1,
	0x8f, 0x27, 0x92, 0x59, 0xb8, 0xba, 0xc8, 0x2a, 0x64, 0x67, 0x61, 0x4e, 0x70, 0xa5, 0x50, 0x18,
	0xdd, 0x05, 0xe7, 0x63, 0xed, 0x79, 0x6a, 0x01, 0xbd, 0x94, 0x52, 0x2c, 0x7c, 0x30, 0xf9, 0x5c,
	0xbb, 0xd7, 0x91, 0x01, 0x62, 0x13, 0x17, 0xc5, 0xe0, 0x4a, 0x05, 0xc7, 0xc5, 0xcb, 0xd8, 0x9a,
	0xcb, 0xb1, 0x7a, 0xab, 0x92, 0x53, 0xaa, 0x47, 0xd9, 0x0b, 0xf2, 0x6c, 0x69, 0x5a, 0x18, 0xed,
	0x41, 0x87, 0xa5, 0x49, 0x74, 0x20, 0x8c, 0xcb, 0xa2, 0x09, 0x07, 0xaf, 0x0a, 0xe9, 0x20, 0xa3,
	0x57, 0x18, 0x23, 0x97, 0x16, 0x01, 0x73, 0x6b, 0x51, 0x66, 0x94, 0xb1, 0x87, 0xa4, 0xf2, 0xaa,
	0x30, 0x01, 0x3c, 0x4a, 0x8d, 0x0e, 0xba, 0x12, 0x7a, 0x08, 0xaf, 0x03, 0xe8, 0x94, 0x87, 0xd7,
	0xd2, 0x5a, 0x08, 0xf4, 0x0b, 0xf6, 0x88, 0xdc, 0xeb, 0xf4, 0x6d, 0x07, 0xf4, 0xe9, 0xc0, 0xab,
	0xd1, 0xe0, 0xc9, 0x57, 0x53, 0x38, 0x01, 0x28, 0x35, 0x28, 0xd9, 0x96, 0x4d, 0x05, 0xf4, 0x5a,
	0x6c, 0xc6, 0x5c, 0xac, 0x79, 0x0e, 0x6b, 0x8d, 0x06, 0x78, 0x1d, 0xe8, 0x26, 0xbb, 0x43, 0xb6,
	0xa5, 0x8e, 0x19, 0xc6, 0xb2, 0xe9, 0x80, 0xa2, 0x70, 0x0e, 0xb4, 0xe8, 0xa3, 0x30, 0x19, 0x78,
	0x7a, 0x3d, 0x71, 0xce, 0x49, 0x01, 0x2b, 0xaa, 0x1d, 0xee, 0x31, 0x2b, 0xac, 0x92, 0x82, 0xc7,
	0xca, 0x75, 0x80, 0x67, 0xe0, 0x3c, 0xbd, 0xc1, 0xbe, 0x25, 0x2f, 0x4a, 0xac, 0x0c, 0x90, 0xaf,
	0x87, 0x21, 0xd7, 0x19, 0x66, 0xe0, 0x85, 0x93, 0x36, 0xb5, 0xf1, 0x26, 0x7b, 0x40, 0xee, 0xae,
	0xf3, 0x10, 0x07, 0xb1, 0x05, 0xbd, 0x64, 0xcb, 0xd3, 0xad, 0x98, 0xfb, 0x87, 0x60, 0x39, 0xd7,
	0xfd, 0x39, 0x6e, 0x3b, 0x0d, 0xb3, 0x7e, 0xdf, 0xd7, 0x34, 0xf0, 0xf4, 0x56, 0xcc, 0xe3, 0x84,
	0x38, 0x16, 0xc6, 0x83, 0x02, 0x11, 0x25, 0x8a, 0xeb, 0x76, 0xc1, 0xdb, 0x40, 0x6f, 0xb3, 0xdb,
	0xe4, 0xe6, 0x92, 0xcb, 0xd0, 0x8d, 0x3e, 0xad, 0x83, 0x16, 0xc4, 0xc2, 0x00, 0xdd, 0x59, 0x1d,
	0x50, 0x2f, 0x3a, 0x90, 0x15, 0x0a, 0x10, 0x74, 0x46, 0xef, 0x44, 0x27, 0x59, 0x9c, 0x9e, 0x20,
	0x73, 0xc0, 0xbc, 0xf0, 0x01, 0x9b, 0xc9, 0x09, 0x17, 0x89, 0x5a, 0xa5, 0xfc, 0x8d, 0xd1, 0x40,
	0xbf, 0x64, 0xdb, 0xe4, 0xfa, 0xd2, 0x4a, 0xd9, 0x60, 0x3f, 0xaf, 0x11, 0xbd, 0x1b, 0x19, 0x31,
	0x5f, 0x13, 0xc6, 0x42, 0x1a, 0x48, 0x13, 0x7b, 0x98, 0x9b, 0x6e, 0xa4, 0x2a, 0x40, 0x96, 0xd2,
	0xa5, 0x15, 0xf6, 0x03, 0xf9, 0xee, 0x94, 0x95, 0x90, 0xf6, 0xcf, 0x31, 0xd2, 0xaf, 0xd2, 0xfc,
	0x5e, 0xe4, 0x89, 0x30, 0xba, 0xa5, 0xe2, 0x56, 0xd3, 0xed, 0x13, 0x93, 0xe6, 0x21, 0x44, 0xb9,
	0xa7, 0xf7, 0x63, 0xb4, 0x2b, 0xdc, 0xe9, 0x5b, 0xc0, 0x5c, 0xfa, 0x9c, 0x07, 0xd1, 0xa1, 0x0f,
	0xe2, 0xd2, 0x5b, 0xe1, 0x63, 0xd1, 0x4c, 0xef, 0xcb, 0xd9, 0x78, 0x78, 0xfc, 0x3d, 0x3d, 0xae,
	0xb4, 0x82, 0x3e, 0x62, 0x5f, 0x93, 0xe7, 0x27, 0xb3, 0x4d, 0x59, 0xa6, 0x04, 0xf2, 0x42, 0x05,
	0x69, 0x15, 0x60, 0xce, 0xad, 0x4d, 0xf1, 0x3c, 0x66, 0x55, 0xf2, 0xf4, 0x34, 0x95, 0x97, 0xd0,
	0x47, 0x1e, 0x82, 0x93, 0xcd, 0x22, 0x80, 0xa7, 0x4f, 0x8e, 0x2d, 0xd5, 0xc4, 0xa2, 0xf9, 0x3a,
	0x7a, 0x1a, 0x97, 0x9a, 0x75, 0x26, 0x37, 0x69, 0xaf, 0x2f, 0x30, 0xb1, 0x51, 0xf4, 0xab, 0x48,
	0xcd, 0x25, 0xbf, 0x32, 0x10, 0x32, 0xe7, 0x0a, 0xad, 0xe2, 0x02, 0x3c, 0xbe, 0x3f, 0x32, 0xcf,
	0x9a, 0xff, 0x9e, 0x21, 0x95, 0x9f, 0x27, 0x07, 0xd5, 0xd3, 0x8f, 0x5e, 0x73, 0x7b, 0xfd, 0xd1,
	0xb2, 0xf1, 0xe6, 0xd9, 0x33, 0x6f, 0xb2, 0xb9, 0xfa, 0x68, 0xf2, 0x76, 0x6f, 0x3c, 0xaa, 0x4e,
	0x8e, 0x46, 0xb5, 0xd1, 0x70, 0x9c, 0x2e, 0xe2, 0xe2, 0x06, 0x1f, 0xee, 0x4f, 0x3f, 0x74, 0x92,
	0xbf, 0x2f, 0x7f, 0xfe, 0xd8, 0x38, 0xdb, 0xe6, 0xfc, 0xcf, 0x8d, 0x9d, 0x76, 0x69, 0x8c, 0x0f,
	0xa6, 0xd5, 0xf2, 0x33, 0x7e, 0x75, 0xeb, 0xd5, 0xe4, 0x72, 0xfa, 0xf7, 0x02, 0xb0, 0xcb, 0x07,
	0xd3, 0xdd, 0x25, 0x60, 0xb7, 0x5b, 0xdf, 0x2d, 0x01, 0xff, 0x6c, 0x54, 0x4a, 0x69, 0xa3, 0xc1,
	0x07, 0xd3, 0x46, 0x63, 0x09, 0x69, 0x34, 0xba, 0xf5, 0x46, 0xa3, 0x04, 0xfd, 0x74, 0x2e, 0x45,
	0xf7, 0xcd, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x13, 0xf0, 0xc2, 0x8a, 0x2f, 0x08, 0x00, 0x00,
}
