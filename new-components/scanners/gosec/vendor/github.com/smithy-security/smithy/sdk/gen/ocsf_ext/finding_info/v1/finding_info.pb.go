// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: ocsf_ext/finding_info/v1/finding_info.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TargetType specifies the target type.
type DataSource_TargetType int32

const (
	DataSource_TARGET_TYPE_UNSPECIFIED DataSource_TargetType = 0
	DataSource_TARGET_TYPE_REPOSITORY  DataSource_TargetType = 1
)

// Enum value maps for DataSource_TargetType.
var (
	DataSource_TargetType_name = map[int32]string{
		0: "TARGET_TYPE_UNSPECIFIED",
		1: "TARGET_TYPE_REPOSITORY",
	}
	DataSource_TargetType_value = map[string]int32{
		"TARGET_TYPE_UNSPECIFIED": 0,
		"TARGET_TYPE_REPOSITORY":  1,
	}
)

func (x DataSource_TargetType) Enum() *DataSource_TargetType {
	p := new(DataSource_TargetType)
	*p = x
	return p
}

func (x DataSource_TargetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataSource_TargetType) Descriptor() protoreflect.EnumDescriptor {
	return file_ocsf_ext_finding_info_v1_finding_info_proto_enumTypes[0].Descriptor()
}

func (DataSource_TargetType) Type() protoreflect.EnumType {
	return &file_ocsf_ext_finding_info_v1_finding_info_proto_enumTypes[0]
}

func (x DataSource_TargetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataSource_TargetType.Descriptor instead.
func (DataSource_TargetType) EnumDescriptor() ([]byte, []int) {
	return file_ocsf_ext_finding_info_v1_finding_info_proto_rawDescGZIP(), []int{0, 0}
}

// URISchema specifies the URI schema.
// For example:
// - purl: pkg:npm/%40angular/animation@12.3.1" -> "pkg"
// - file: file://main.go -> "file"
type DataSource_URISchema int32

const (
	DataSource_URI_SCHEMA_UNSPECIFIED DataSource_URISchema = 0
	DataSource_URI_SCHEMA_FILE        DataSource_URISchema = 1
	DataSource_URI_SCHEMA_PURL        DataSource_URISchema = 2
)

// Enum value maps for DataSource_URISchema.
var (
	DataSource_URISchema_name = map[int32]string{
		0: "URI_SCHEMA_UNSPECIFIED",
		1: "URI_SCHEMA_FILE",
		2: "URI_SCHEMA_PURL",
	}
	DataSource_URISchema_value = map[string]int32{
		"URI_SCHEMA_UNSPECIFIED": 0,
		"URI_SCHEMA_FILE":        1,
		"URI_SCHEMA_PURL":        2,
	}
)

func (x DataSource_URISchema) Enum() *DataSource_URISchema {
	p := new(DataSource_URISchema)
	*p = x
	return p
}

func (x DataSource_URISchema) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataSource_URISchema) Descriptor() protoreflect.EnumDescriptor {
	return file_ocsf_ext_finding_info_v1_finding_info_proto_enumTypes[1].Descriptor()
}

func (DataSource_URISchema) Type() protoreflect.EnumType {
	return &file_ocsf_ext_finding_info_v1_finding_info_proto_enumTypes[1]
}

func (x DataSource_URISchema) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataSource_URISchema.Descriptor instead.
func (DataSource_URISchema) EnumDescriptor() ([]byte, []int) {
	return file_ocsf_ext_finding_info_v1_finding_info_proto_rawDescGZIP(), []int{0, 1}
}

// DataSource is used to define Data Sources described on https://schema.ocsf.io/1.3.0/objects/finding_info.
type DataSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetType DataSource_TargetType `protobuf:"varint,1,opt,name=target_type,json=targetType,proto3,enum=smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource_TargetType" json:"target_type,omitempty"`
	Uri        *DataSource_URI       `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	// Types that are assignable to LocationData:
	//
	//	*DataSource_FileFindingLocationData_
	//	*DataSource_PurlFindingLocationData_
	LocationData isDataSource_LocationData `protobuf_oneof:"location_data"`
}

func (x *DataSource) Reset() {
	*x = DataSource{}
	mi := &file_ocsf_ext_finding_info_v1_finding_info_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSource) ProtoMessage() {}

func (x *DataSource) ProtoReflect() protoreflect.Message {
	mi := &file_ocsf_ext_finding_info_v1_finding_info_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSource.ProtoReflect.Descriptor instead.
func (*DataSource) Descriptor() ([]byte, []int) {
	return file_ocsf_ext_finding_info_v1_finding_info_proto_rawDescGZIP(), []int{0}
}

func (x *DataSource) GetTargetType() DataSource_TargetType {
	if x != nil {
		return x.TargetType
	}
	return DataSource_TARGET_TYPE_UNSPECIFIED
}

func (x *DataSource) GetUri() *DataSource_URI {
	if x != nil {
		return x.Uri
	}
	return nil
}

func (m *DataSource) GetLocationData() isDataSource_LocationData {
	if m != nil {
		return m.LocationData
	}
	return nil
}

func (x *DataSource) GetFileFindingLocationData() *DataSource_FileFindingLocationData {
	if x, ok := x.GetLocationData().(*DataSource_FileFindingLocationData_); ok {
		return x.FileFindingLocationData
	}
	return nil
}

func (x *DataSource) GetPurlFindingLocationData() *DataSource_PurlFindingLocationData {
	if x, ok := x.GetLocationData().(*DataSource_PurlFindingLocationData_); ok {
		return x.PurlFindingLocationData
	}
	return nil
}

type isDataSource_LocationData interface {
	isDataSource_LocationData()
}

type DataSource_FileFindingLocationData_ struct {
	FileFindingLocationData *DataSource_FileFindingLocationData `protobuf:"bytes,3,opt,name=file_finding_location_data,json=fileFindingLocationData,proto3,oneof"`
}

type DataSource_PurlFindingLocationData_ struct {
	PurlFindingLocationData *DataSource_PurlFindingLocationData `protobuf:"bytes,4,opt,name=purl_finding_location_data,json=purlFindingLocationData,proto3,oneof"`
}

func (*DataSource_FileFindingLocationData_) isDataSource_LocationData() {}

func (*DataSource_PurlFindingLocationData_) isDataSource_LocationData() {}

// URI specifies uri content.
type DataSource_URI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UriSchema DataSource_URISchema `protobuf:"varint,1,opt,name=uri_schema,json=uriSchema,proto3,enum=smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource_URISchema" json:"uri_schema,omitempty"`
	Path      string               `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DataSource_URI) Reset() {
	*x = DataSource_URI{}
	mi := &file_ocsf_ext_finding_info_v1_finding_info_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataSource_URI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSource_URI) ProtoMessage() {}

func (x *DataSource_URI) ProtoReflect() protoreflect.Message {
	mi := &file_ocsf_ext_finding_info_v1_finding_info_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSource_URI.ProtoReflect.Descriptor instead.
func (*DataSource_URI) Descriptor() ([]byte, []int) {
	return file_ocsf_ext_finding_info_v1_finding_info_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DataSource_URI) GetUriSchema() DataSource_URISchema {
	if x != nil {
		return x.UriSchema
	}
	return DataSource_URI_SCHEMA_UNSPECIFIED
}

func (x *DataSource_URI) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// FileFindingLocationData specifies data associated with the physical location of a finding in a file.
type DataSource_FileFindingLocationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartLine   uint32 `protobuf:"varint,3,opt,name=start_line,json=startLine,proto3" json:"start_line,omitempty"`
	EndLine     uint32 `protobuf:"varint,4,opt,name=end_line,json=endLine,proto3" json:"end_line,omitempty"`
	StartColumn uint32 `protobuf:"varint,5,opt,name=start_column,json=startColumn,proto3" json:"start_column,omitempty"`
	EndColumn   uint32 `protobuf:"varint,6,opt,name=end_column,json=endColumn,proto3" json:"end_column,omitempty"`
}

func (x *DataSource_FileFindingLocationData) Reset() {
	*x = DataSource_FileFindingLocationData{}
	mi := &file_ocsf_ext_finding_info_v1_finding_info_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataSource_FileFindingLocationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSource_FileFindingLocationData) ProtoMessage() {}

func (x *DataSource_FileFindingLocationData) ProtoReflect() protoreflect.Message {
	mi := &file_ocsf_ext_finding_info_v1_finding_info_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSource_FileFindingLocationData.ProtoReflect.Descriptor instead.
func (*DataSource_FileFindingLocationData) Descriptor() ([]byte, []int) {
	return file_ocsf_ext_finding_info_v1_finding_info_proto_rawDescGZIP(), []int{0, 1}
}

func (x *DataSource_FileFindingLocationData) GetStartLine() uint32 {
	if x != nil {
		return x.StartLine
	}
	return 0
}

func (x *DataSource_FileFindingLocationData) GetEndLine() uint32 {
	if x != nil {
		return x.EndLine
	}
	return 0
}

func (x *DataSource_FileFindingLocationData) GetStartColumn() uint32 {
	if x != nil {
		return x.StartColumn
	}
	return 0
}

func (x *DataSource_FileFindingLocationData) GetEndColumn() uint32 {
	if x != nil {
		return x.EndColumn
	}
	return 0
}

// PurlFindingLocationData specifies data associated with the logical location of a finding in a purl.
type DataSource_PurlFindingLocationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DataSource_PurlFindingLocationData) Reset() {
	*x = DataSource_PurlFindingLocationData{}
	mi := &file_ocsf_ext_finding_info_v1_finding_info_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataSource_PurlFindingLocationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSource_PurlFindingLocationData) ProtoMessage() {}

func (x *DataSource_PurlFindingLocationData) ProtoReflect() protoreflect.Message {
	mi := &file_ocsf_ext_finding_info_v1_finding_info_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSource_PurlFindingLocationData.ProtoReflect.Descriptor instead.
func (*DataSource_PurlFindingLocationData) Descriptor() ([]byte, []int) {
	return file_ocsf_ext_finding_info_v1_finding_info_proto_rawDescGZIP(), []int{0, 2}
}

var File_ocsf_ext_finding_info_v1_finding_info_proto protoreflect.FileDescriptor

var file_ocsf_ext_finding_info_v1_finding_info_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6f, 0x63, 0x73, 0x66, 0x5f, 0x65, 0x78, 0x74, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x73,
	0x6d, 0x69, 0x74, 0x68, 0x79, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x73,
	0x6d, 0x69, 0x74, 0x68, 0x79, 0x2e, 0x6f, 0x63, 0x73, 0x66, 0x5f, 0x65, 0x78, 0x74, 0x2e, 0x66,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0xd1,
	0x07, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x67, 0x0a,
	0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x46, 0x2e, 0x73, 0x6d, 0x69, 0x74, 0x68, 0x79, 0x2e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x6d, 0x69, 0x74, 0x68, 0x79, 0x2e, 0x6f, 0x63, 0x73, 0x66,
	0x5f, 0x65, 0x78, 0x74, 0x2e, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x51, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x73, 0x6d, 0x69, 0x74, 0x68, 0x79, 0x2e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x6d, 0x69, 0x74, 0x68, 0x79, 0x2e, 0x6f, 0x63, 0x73,
	0x66, 0x5f, 0x65, 0x78, 0x74, 0x2e, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x55, 0x52, 0x49, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x92, 0x01, 0x0a, 0x1a, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x53,
	0x2e, 0x73, 0x6d, 0x69, 0x74, 0x68, 0x79, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x73, 0x6d, 0x69, 0x74, 0x68, 0x79, 0x2e, 0x6f, 0x63, 0x73, 0x66, 0x5f, 0x65, 0x78, 0x74,
	0x2e, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x46, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x17, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x92,
	0x01, 0x0a, 0x1a, 0x70, 0x75, 0x72, 0x6c, 0x5f, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x53, 0x2e, 0x73, 0x6d, 0x69, 0x74, 0x68, 0x79, 0x2e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x6d, 0x69, 0x74, 0x68, 0x79, 0x2e, 0x6f, 0x63, 0x73,
	0x66, 0x5f, 0x65, 0x78, 0x74, 0x2e, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x50, 0x75, 0x72, 0x6c, 0x46, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x17, 0x70, 0x75, 0x72, 0x6c,
	0x46, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x7f, 0x0a, 0x03, 0x55, 0x52, 0x49, 0x12, 0x64, 0x0a, 0x0a, 0x75, 0x72,
	0x69, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x45,
	0x2e, 0x73, 0x6d, 0x69, 0x74, 0x68, 0x79, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x73, 0x6d, 0x69, 0x74, 0x68, 0x79, 0x2e, 0x6f, 0x63, 0x73, 0x66, 0x5f, 0x65, 0x78, 0x74,
	0x2e, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x55, 0x52, 0x49, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x09, 0x75, 0x72, 0x69, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x1a, 0x95, 0x01, 0x0a, 0x17, 0x46, 0x69, 0x6c, 0x65, 0x46, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x1a, 0x19, 0x0a, 0x17,
	0x50, 0x75, 0x72, 0x6c, 0x46, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x45, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x01, 0x22, 0x51,
	0x0a, 0x09, 0x55, 0x52, 0x49, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x1a, 0x0a, 0x16, 0x55,
	0x52, 0x49, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x52, 0x49, 0x5f, 0x53,
	0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x55, 0x52, 0x49, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x50, 0x55, 0x52, 0x4c, 0x10,
	0x02, 0x42, 0x0f, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6d, 0x69, 0x74, 0x68, 0x79, 0x2d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2f, 0x73, 0x6d, 0x69, 0x74, 0x68, 0x79, 0x2f, 0x6f, 0x63, 0x73, 0x66, 0x5f, 0x65, 0x78, 0x74,
	0x2f, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ocsf_ext_finding_info_v1_finding_info_proto_rawDescOnce sync.Once
	file_ocsf_ext_finding_info_v1_finding_info_proto_rawDescData = file_ocsf_ext_finding_info_v1_finding_info_proto_rawDesc
)

func file_ocsf_ext_finding_info_v1_finding_info_proto_rawDescGZIP() []byte {
	file_ocsf_ext_finding_info_v1_finding_info_proto_rawDescOnce.Do(func() {
		file_ocsf_ext_finding_info_v1_finding_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_ocsf_ext_finding_info_v1_finding_info_proto_rawDescData)
	})
	return file_ocsf_ext_finding_info_v1_finding_info_proto_rawDescData
}

var file_ocsf_ext_finding_info_v1_finding_info_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ocsf_ext_finding_info_v1_finding_info_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ocsf_ext_finding_info_v1_finding_info_proto_goTypes = []any{
	(DataSource_TargetType)(0),                 // 0: smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource.TargetType
	(DataSource_URISchema)(0),                  // 1: smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource.URISchema
	(*DataSource)(nil),                         // 2: smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource
	(*DataSource_URI)(nil),                     // 3: smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource.URI
	(*DataSource_FileFindingLocationData)(nil), // 4: smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource.FileFindingLocationData
	(*DataSource_PurlFindingLocationData)(nil), // 5: smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource.PurlFindingLocationData
}
var file_ocsf_ext_finding_info_v1_finding_info_proto_depIdxs = []int32{
	0, // 0: smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource.target_type:type_name -> smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource.TargetType
	3, // 1: smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource.uri:type_name -> smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource.URI
	4, // 2: smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource.file_finding_location_data:type_name -> smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource.FileFindingLocationData
	5, // 3: smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource.purl_finding_location_data:type_name -> smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource.PurlFindingLocationData
	1, // 4: smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource.URI.uri_schema:type_name -> smithy.security.smithy.ocsf_ext.finding_info.v1.DataSource.URISchema
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ocsf_ext_finding_info_v1_finding_info_proto_init() }
func file_ocsf_ext_finding_info_v1_finding_info_proto_init() {
	if File_ocsf_ext_finding_info_v1_finding_info_proto != nil {
		return
	}
	file_ocsf_ext_finding_info_v1_finding_info_proto_msgTypes[0].OneofWrappers = []any{
		(*DataSource_FileFindingLocationData_)(nil),
		(*DataSource_PurlFindingLocationData_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ocsf_ext_finding_info_v1_finding_info_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ocsf_ext_finding_info_v1_finding_info_proto_goTypes,
		DependencyIndexes: file_ocsf_ext_finding_info_v1_finding_info_proto_depIdxs,
		EnumInfos:         file_ocsf_ext_finding_info_v1_finding_info_proto_enumTypes,
		MessageInfos:      file_ocsf_ext_finding_info_v1_finding_info_proto_msgTypes,
	}.Build()
	File_ocsf_ext_finding_info_v1_finding_info_proto = out.File
	file_ocsf_ext_finding_info_v1_finding_info_proto_rawDesc = nil
	file_ocsf_ext_finding_info_v1_finding_info_proto_goTypes = nil
	file_ocsf_ext_finding_info_v1_finding_info_proto_depIdxs = nil
}