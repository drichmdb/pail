// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/google_ads_field_data_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// These are the various types a GoogleAdsService artifact may take on.
type GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType int32

const (
	// Unspecified
	GoogleAdsFieldDataTypeEnum_UNSPECIFIED GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 0
	// Unknown
	GoogleAdsFieldDataTypeEnum_UNKNOWN GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 1
	// Maps to google.protobuf.BoolValue
	//
	// Applicable operators:  =, !=
	GoogleAdsFieldDataTypeEnum_BOOLEAN GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 2
	// Maps to google.protobuf.StringValue. It can be compared using the set of
	// operators specific to dates however.
	//
	// Applicable operators:  =, <, >, <=, >=, BETWEEN, DURING, and IN
	GoogleAdsFieldDataTypeEnum_DATE GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 3
	// Maps to google.protobuf.DoubleValue
	//
	// Applicable operators:  =, !=, <, >, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_DOUBLE GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 4
	// Maps to an enum. It's specific definition can be found at type_url.
	//
	// Applicable operators:  =, !=, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_ENUM GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 5
	// Maps to google.protobuf.FloatValue
	//
	// Applicable operators:  =, !=, <, >, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_FLOAT GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 6
	// Maps to google.protobuf.Int32Value
	//
	// Applicable operators:  =, !=, <, >, <=, >=, BETWEEN, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_INT32 GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 7
	// Maps to google.protobuf.Int64Value
	//
	// Applicable operators:  =, !=, <, >, <=, >=, BETWEEN, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_INT64 GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 8
	// Maps to a protocol buffer message type. The data type's details can be
	// found in type_url.
	//
	// No operators work with MESSAGE fields.
	GoogleAdsFieldDataTypeEnum_MESSAGE GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 9
	// Maps to google.protobuf.StringValue. Represents the resource name
	// (unique id) of a resource or one of its foreign keys.
	//
	// No operators work with RESOURCE_NAME fields.
	GoogleAdsFieldDataTypeEnum_RESOURCE_NAME GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 10
	// Maps to google.protobuf.StringValue.
	//
	// Applicable operators:  =, !=, LIKE, NOT LIKE, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_STRING GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 11
)

var GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "BOOLEAN",
	3:  "DATE",
	4:  "DOUBLE",
	5:  "ENUM",
	6:  "FLOAT",
	7:  "INT32",
	8:  "INT64",
	9:  "MESSAGE",
	10: "RESOURCE_NAME",
	11: "STRING",
}
var GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType_value = map[string]int32{
	"UNSPECIFIED":   0,
	"UNKNOWN":       1,
	"BOOLEAN":       2,
	"DATE":          3,
	"DOUBLE":        4,
	"ENUM":          5,
	"FLOAT":         6,
	"INT32":         7,
	"INT64":         8,
	"MESSAGE":       9,
	"RESOURCE_NAME": 10,
	"STRING":        11,
}

func (x GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType) String() string {
	return proto.EnumName(GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType_name, int32(x))
}
func (GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_google_ads_field_data_type_ceea0e0f3ddb5014, []int{0, 0}
}

// Container holding the various data types.
type GoogleAdsFieldDataTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoogleAdsFieldDataTypeEnum) Reset()         { *m = GoogleAdsFieldDataTypeEnum{} }
func (m *GoogleAdsFieldDataTypeEnum) String() string { return proto.CompactTextString(m) }
func (*GoogleAdsFieldDataTypeEnum) ProtoMessage()    {}
func (*GoogleAdsFieldDataTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_google_ads_field_data_type_ceea0e0f3ddb5014, []int{0}
}
func (m *GoogleAdsFieldDataTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleAdsFieldDataTypeEnum.Unmarshal(m, b)
}
func (m *GoogleAdsFieldDataTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleAdsFieldDataTypeEnum.Marshal(b, m, deterministic)
}
func (dst *GoogleAdsFieldDataTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleAdsFieldDataTypeEnum.Merge(dst, src)
}
func (m *GoogleAdsFieldDataTypeEnum) XXX_Size() int {
	return xxx_messageInfo_GoogleAdsFieldDataTypeEnum.Size(m)
}
func (m *GoogleAdsFieldDataTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleAdsFieldDataTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleAdsFieldDataTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GoogleAdsFieldDataTypeEnum)(nil), "google.ads.googleads.v0.enums.GoogleAdsFieldDataTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType", GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType_name, GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/google_ads_field_data_type.proto", fileDescriptor_google_ads_field_data_type_ceea0e0f3ddb5014)
}

var fileDescriptor_google_ads_field_data_type_ceea0e0f3ddb5014 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcb, 0x0e, 0x93, 0x40,
	0x14, 0x15, 0xfa, 0x9e, 0xc6, 0x38, 0xb2, 0x70, 0xa1, 0xe9, 0xa2, 0xfd, 0x80, 0x81, 0x58, 0xe3,
	0x62, 0x4c, 0x4c, 0x86, 0x32, 0x25, 0xc4, 0x76, 0x68, 0x0a, 0xd4, 0xc4, 0x90, 0x90, 0xd1, 0x41,
	0xd2, 0xa4, 0x05, 0xd2, 0xa1, 0x4d, 0xfa, 0x3b, 0x2e, 0x5d, 0xf9, 0x1d, 0x6e, 0xfc, 0x0f, 0x57,
	0x7e, 0x82, 0x19, 0x68, 0x59, 0x55, 0x37, 0xe4, 0x70, 0xce, 0xbd, 0xe7, 0xce, 0x3d, 0x17, 0xbc,
	0xcf, 0x8a, 0x22, 0x3b, 0xa4, 0x26, 0x17, 0xd2, 0x6c, 0xa0, 0x42, 0x17, 0xcb, 0x4c, 0xf3, 0xf3,
	0xf1, 0x4e, 0x25, 0x5c, 0xc8, 0xe4, 0xeb, 0x3e, 0x3d, 0x88, 0x44, 0xf0, 0x8a, 0x27, 0xd5, 0xb5,
	0x4c, 0x51, 0x79, 0x2a, 0xaa, 0xc2, 0x98, 0x34, 0x15, 0x88, 0x0b, 0x89, 0xda, 0x7e, 0x74, 0xb1,
	0x50, 0xdd, 0x3f, 0xfb, 0xa5, 0x81, 0x97, 0x6e, 0x4d, 0x13, 0x21, 0x97, 0xca, 0xc1, 0xe1, 0x15,
	0x0f, 0xaf, 0x65, 0x4a, 0xf3, 0xf3, 0x71, 0xf6, 0x43, 0x03, 0x2f, 0x1e, 0xcb, 0xc6, 0x33, 0x30,
	0x8e, 0x58, 0xb0, 0xa1, 0x0b, 0x6f, 0xe9, 0x51, 0x07, 0x3e, 0x31, 0xc6, 0x60, 0x10, 0xb1, 0x0f,
	0xcc, 0xff, 0xc8, 0xa0, 0xa6, 0x7e, 0x6c, 0xdf, 0x5f, 0x51, 0xc2, 0xa0, 0x6e, 0x0c, 0x41, 0xd7,
	0x21, 0x21, 0x85, 0x1d, 0x03, 0x80, 0xbe, 0xe3, 0x47, 0xf6, 0x8a, 0xc2, 0xae, 0x62, 0x29, 0x8b,
	0xd6, 0xb0, 0x67, 0x8c, 0x40, 0x6f, 0xb9, 0xf2, 0x49, 0x08, 0xfb, 0x0a, 0x7a, 0x2c, 0x9c, 0xbf,
	0x86, 0x83, 0x1b, 0x7c, 0xfb, 0x06, 0x0e, 0x95, 0xdb, 0x9a, 0x06, 0x01, 0x71, 0x29, 0x1c, 0x19,
	0xcf, 0xc1, 0xd3, 0x2d, 0x0d, 0xfc, 0x68, 0xbb, 0xa0, 0x09, 0x23, 0x6b, 0x0a, 0x81, 0xb2, 0x0d,
	0xc2, 0xad, 0xc7, 0x5c, 0x38, 0xb6, 0xff, 0x68, 0x60, 0xfa, 0xa5, 0x38, 0xa2, 0xff, 0xee, 0x6d,
	0xbf, 0x7a, 0xbc, 0xd5, 0x46, 0x65, 0xb6, 0xd1, 0x3e, 0xd9, 0xb7, 0xee, 0xac, 0x38, 0xf0, 0x3c,
	0x43, 0xc5, 0x29, 0x33, 0xb3, 0x34, 0xaf, 0x13, 0xbd, 0x5f, 0xa1, 0xdc, 0xcb, 0x7f, 0x1c, 0xe5,
	0x5d, 0xfd, 0xfd, 0xa6, 0x77, 0x5c, 0x42, 0xbe, 0xeb, 0x93, 0x66, 0x12, 0x22, 0x42, 0xa2, 0x76,
	0x28, 0xda, 0x59, 0x48, 0x05, 0x2c, 0x7f, 0xde, 0xf5, 0x98, 0x08, 0x19, 0xb7, 0x7a, 0xbc, 0xb3,
	0xe2, 0x5a, 0xff, 0xad, 0x4f, 0x1b, 0x12, 0x63, 0x22, 0x24, 0xc6, 0x6d, 0x05, 0xc6, 0x3b, 0x0b,
	0xe3, 0xba, 0xe6, 0x73, 0xbf, 0x7e, 0xd8, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x72,
	0x84, 0x1c, 0x2c, 0x02, 0x00, 0x00,
}