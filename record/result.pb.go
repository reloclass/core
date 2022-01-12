// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/core/record/result.proto

package record

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

// RecordResult 录制任务结果
type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 错误类型，只有请求出错的时候才会返回此字段
	// @gotags: json:"ErrorCode"
	ErrorCode string `protobuf:"bytes,2,opt,name=error_code,json=errorCode,proto3" json:"ErrorCode"`
	// 错误说明，只有请求出错的时候才会返回此字段
	// @gotags: json:"ErrorMessage"
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"ErrorMessage"`
	// 录制视频⽂件列表
	// @gotags: json:"Videos"
	Videos []*Video `protobuf:"bytes,4,rep,name=videos,proto3" json:"Videos"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_core_record_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_relo_core_record_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_relo_core_record_result_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

func (x *Result) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *Result) GetVideos() []*Video {
	if x != nil {
		return x.Videos
	}
	return nil
}

// 录制任务视频
type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 视频文件名称
	// @gotags: json:"Filename"
	FileName string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"Filename"`
	// 视频大小
	// @gotags: json:"FileSize"
	FileSize int64 `protobuf:"varint,3,opt,name=file_size,json=fileSize,proto3" json:"FileSize"`
	// 视频文件链接
	// @gotags: json:"FileURL"
	FileUrl string `protobuf:"bytes,4,opt,name=file_url,json=fileUrl,proto3" json:"FileURL"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_core_record_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_relo_core_record_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_relo_core_record_result_proto_rawDescGZIP(), []int{1}
}

func (x *Video) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *Video) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *Video) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

var File_relo_core_record_result_proto protoreflect.FileDescriptor

var file_relo_core_record_result_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x22, 0x7d, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x2f, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73,
	0x22, 0x5c, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x42, 0x22,
	0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c,
	0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_core_record_result_proto_rawDescOnce sync.Once
	file_relo_core_record_result_proto_rawDescData = file_relo_core_record_result_proto_rawDesc
)

func file_relo_core_record_result_proto_rawDescGZIP() []byte {
	file_relo_core_record_result_proto_rawDescOnce.Do(func() {
		file_relo_core_record_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_core_record_result_proto_rawDescData)
	})
	return file_relo_core_record_result_proto_rawDescData
}

var file_relo_core_record_result_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_relo_core_record_result_proto_goTypes = []interface{}{
	(*Result)(nil), // 0: relo.core.record.Result
	(*Video)(nil),  // 1: relo.core.record.Video
}
var file_relo_core_record_result_proto_depIdxs = []int32{
	1, // 0: relo.core.record.Result.videos:type_name -> relo.core.record.Video
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_relo_core_record_result_proto_init() }
func file_relo_core_record_result_proto_init() {
	if File_relo_core_record_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_core_record_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_relo_core_record_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
			RawDescriptor: file_relo_core_record_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_core_record_result_proto_goTypes,
		DependencyIndexes: file_relo_core_record_result_proto_depIdxs,
		MessageInfos:      file_relo_core_record_result_proto_msgTypes,
	}.Build()
	File_relo_core_record_result_proto = out.File
	file_relo_core_record_result_proto_rawDesc = nil
	file_relo_core_record_result_proto_goTypes = nil
	file_relo_core_record_result_proto_depIdxs = nil
}
