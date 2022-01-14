// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/core/homework/attachment_type.proto

package homework

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

// 附件类型
type AttachmentType int32

const (
	// 图片类型
	AttachmentType_PICTURE AttachmentType = 0
	// 录音类型
	AttachmentType_RECORDING AttachmentType = 1
	// 云盘资源
	AttachmentType_CLOUD_RESOURCE AttachmentType = 2
	// 本地文件资源
	AttachmentType_LOCAL_FILE AttachmentType = 3
	// 超链接资源
	AttachmentType_HYPER_LINK AttachmentType = 4
)

// Enum value maps for AttachmentType.
var (
	AttachmentType_name = map[int32]string{
		0: "PICTURE",
		1: "RECORDING",
		2: "CLOUD_RESOURCE",
		3: "LOCAL_FILE",
		4: "HYPER_LINK",
	}
	AttachmentType_value = map[string]int32{
		"PICTURE":        0,
		"RECORDING":      1,
		"CLOUD_RESOURCE": 2,
		"LOCAL_FILE":     3,
		"HYPER_LINK":     4,
	}
)

func (x AttachmentType) Enum() *AttachmentType {
	p := new(AttachmentType)
	*p = x
	return p
}

func (x AttachmentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttachmentType) Descriptor() protoreflect.EnumDescriptor {
	return file_relo_core_homework_attachment_type_proto_enumTypes[0].Descriptor()
}

func (AttachmentType) Type() protoreflect.EnumType {
	return &file_relo_core_homework_attachment_type_proto_enumTypes[0]
}

func (x AttachmentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttachmentType.Descriptor instead.
func (AttachmentType) EnumDescriptor() ([]byte, []int) {
	return file_relo_core_homework_attachment_type_proto_rawDescGZIP(), []int{0}
}

var File_relo_core_homework_attachment_type_proto protoreflect.FileDescriptor

var file_relo_core_homework_attachment_type_proto_rawDesc = []byte{
	0x0a, 0x28, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72, 0x65, 0x6c, 0x6f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2a, 0x60,
	0x0a, 0x0e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x49, 0x43, 0x54, 0x55, 0x52, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x02,
	0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x03,
	0x12, 0x0e, 0x0a, 0x0a, 0x48, 0x59, 0x50, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x04,
	0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x6f,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_core_homework_attachment_type_proto_rawDescOnce sync.Once
	file_relo_core_homework_attachment_type_proto_rawDescData = file_relo_core_homework_attachment_type_proto_rawDesc
)

func file_relo_core_homework_attachment_type_proto_rawDescGZIP() []byte {
	file_relo_core_homework_attachment_type_proto_rawDescOnce.Do(func() {
		file_relo_core_homework_attachment_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_core_homework_attachment_type_proto_rawDescData)
	})
	return file_relo_core_homework_attachment_type_proto_rawDescData
}

var file_relo_core_homework_attachment_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_relo_core_homework_attachment_type_proto_goTypes = []interface{}{
	(AttachmentType)(0), // 0: relo.core.homework.AttachmentType
}
var file_relo_core_homework_attachment_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relo_core_homework_attachment_type_proto_init() }
func file_relo_core_homework_attachment_type_proto_init() {
	if File_relo_core_homework_attachment_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relo_core_homework_attachment_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_core_homework_attachment_type_proto_goTypes,
		DependencyIndexes: file_relo_core_homework_attachment_type_proto_depIdxs,
		EnumInfos:         file_relo_core_homework_attachment_type_proto_enumTypes,
	}.Build()
	File_relo_core_homework_attachment_type_proto = out.File
	file_relo_core_homework_attachment_type_proto_rawDesc = nil
	file_relo_core_homework_attachment_type_proto_goTypes = nil
	file_relo_core_homework_attachment_type_proto_depIdxs = nil
}