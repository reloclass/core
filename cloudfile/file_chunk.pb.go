// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/core/cloudfile/file_chunk.proto

package cloudfile

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//  分块上传块信息
type ChunkInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  分块序列号
	// @gotags: json:"partNumber" validate:"required"
	PartNumber int64 `protobuf:"varint,2,opt,name=part_number,json=partNumber,proto3" json:"partNumber" validate:"required"`
	//  分块上传返回校验码
	// @gotags: json:"etag" validate:"required"
	ETag string `protobuf:"bytes,3,opt,name=e_tag,json=eTag,proto3" json:"etag" validate:"required"`
}

func (x *ChunkInfo) Reset() {
	*x = ChunkInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_core_cloudfile_file_chunk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkInfo) ProtoMessage() {}

func (x *ChunkInfo) ProtoReflect() protoreflect.Message {
	mi := &file_relo_core_cloudfile_file_chunk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkInfo.ProtoReflect.Descriptor instead.
func (*ChunkInfo) Descriptor() ([]byte, []int) {
	return file_relo_core_cloudfile_file_chunk_proto_rawDescGZIP(), []int{0}
}

func (x *ChunkInfo) GetPartNumber() int64 {
	if x != nil {
		return x.PartNumber
	}
	return 0
}

func (x *ChunkInfo) GetETag() string {
	if x != nil {
		return x.ETag
	}
	return ""
}

//  分块上传数据
type ChunkData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  文件编号
	// @gotags: json:"fileId" validate:"required,len=20"
	FileId string `protobuf:"bytes,2,opt,name=file_id,json=fileId,proto3" json:"fileId" validate:"required,len=20"`
	// Type 文件所属类型
	// @gotags: json:"type" validate:"required,oneof=course public private resource"
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type" validate:"required,oneof=course public private resource"`
}

func (x *ChunkData) Reset() {
	*x = ChunkData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_core_cloudfile_file_chunk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkData) ProtoMessage() {}

func (x *ChunkData) ProtoReflect() protoreflect.Message {
	mi := &file_relo_core_cloudfile_file_chunk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkData.ProtoReflect.Descriptor instead.
func (*ChunkData) Descriptor() ([]byte, []int) {
	return file_relo_core_cloudfile_file_chunk_proto_rawDescGZIP(), []int{1}
}

func (x *ChunkData) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *ChunkData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_relo_core_cloudfile_file_chunk_proto protoreflect.FileDescriptor

var file_relo_core_cloudfile_file_chunk_proto_rawDesc = []byte{
	0x0a, 0x24, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x41, 0x0a, 0x09, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70,
	0x61, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x13, 0x0a, 0x05, 0x65, 0x5f, 0x74,
	0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x54, 0x61, 0x67, 0x22, 0x38,
	0x0a, 0x09, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x69, 0x6c, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_core_cloudfile_file_chunk_proto_rawDescOnce sync.Once
	file_relo_core_cloudfile_file_chunk_proto_rawDescData = file_relo_core_cloudfile_file_chunk_proto_rawDesc
)

func file_relo_core_cloudfile_file_chunk_proto_rawDescGZIP() []byte {
	file_relo_core_cloudfile_file_chunk_proto_rawDescOnce.Do(func() {
		file_relo_core_cloudfile_file_chunk_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_core_cloudfile_file_chunk_proto_rawDescData)
	})
	return file_relo_core_cloudfile_file_chunk_proto_rawDescData
}

var file_relo_core_cloudfile_file_chunk_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_relo_core_cloudfile_file_chunk_proto_goTypes = []interface{}{
	(*ChunkInfo)(nil), // 0: relo.core.cloudfile.ChunkInfo
	(*ChunkData)(nil), // 1: relo.core.cloudfile.ChunkData
}
var file_relo_core_cloudfile_file_chunk_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relo_core_cloudfile_file_chunk_proto_init() }
func file_relo_core_cloudfile_file_chunk_proto_init() {
	if File_relo_core_cloudfile_file_chunk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_core_cloudfile_file_chunk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relo_core_cloudfile_file_chunk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relo_core_cloudfile_file_chunk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_core_cloudfile_file_chunk_proto_goTypes,
		DependencyIndexes: file_relo_core_cloudfile_file_chunk_proto_depIdxs,
		MessageInfos:      file_relo_core_cloudfile_file_chunk_proto_msgTypes,
	}.Build()
	File_relo_core_cloudfile_file_chunk_proto = out.File
	file_relo_core_cloudfile_file_chunk_proto_rawDesc = nil
	file_relo_core_cloudfile_file_chunk_proto_goTypes = nil
	file_relo_core_cloudfile_file_chunk_proto_depIdxs = nil
}