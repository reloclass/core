// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/core/homework_sttachment_affiliation_type.proto

package core

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

// 作业附件附属类型
type HomeworkAttachmentAffiliationType int32

const (
	// 归属作业类容
	HomeworkAttachmentAffiliationType_CONTENT HomeworkAttachmentAffiliationType = 0
	// 归属老师评语
	HomeworkAttachmentAffiliationType_COMMENT HomeworkAttachmentAffiliationType = 1
)

// Enum value maps for HomeworkAttachmentAffiliationType.
var (
	HomeworkAttachmentAffiliationType_name = map[int32]string{
		0: "CONTENT",
		1: "COMMENT",
	}
	HomeworkAttachmentAffiliationType_value = map[string]int32{
		"CONTENT": 0,
		"COMMENT": 1,
	}
)

func (x HomeworkAttachmentAffiliationType) Enum() *HomeworkAttachmentAffiliationType {
	p := new(HomeworkAttachmentAffiliationType)
	*p = x
	return p
}

func (x HomeworkAttachmentAffiliationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HomeworkAttachmentAffiliationType) Descriptor() protoreflect.EnumDescriptor {
	return file_relo_core_homework_sttachment_affiliation_type_proto_enumTypes[0].Descriptor()
}

func (HomeworkAttachmentAffiliationType) Type() protoreflect.EnumType {
	return &file_relo_core_homework_sttachment_affiliation_type_proto_enumTypes[0]
}

func (x HomeworkAttachmentAffiliationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HomeworkAttachmentAffiliationType.Descriptor instead.
func (HomeworkAttachmentAffiliationType) EnumDescriptor() ([]byte, []int) {
	return file_relo_core_homework_sttachment_affiliation_type_proto_rawDescGZIP(), []int{0}
}

var File_relo_core_homework_sttachment_affiliation_type_proto protoreflect.FileDescriptor

var file_relo_core_homework_sttachment_affiliation_type_proto_rawDesc = []byte{
	0x0a, 0x34, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2a, 0x3d, 0x0a, 0x21, 0x48, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e,
	0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x01,
	0x42, 0x1b, 0x5a, 0x19, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_core_homework_sttachment_affiliation_type_proto_rawDescOnce sync.Once
	file_relo_core_homework_sttachment_affiliation_type_proto_rawDescData = file_relo_core_homework_sttachment_affiliation_type_proto_rawDesc
)

func file_relo_core_homework_sttachment_affiliation_type_proto_rawDescGZIP() []byte {
	file_relo_core_homework_sttachment_affiliation_type_proto_rawDescOnce.Do(func() {
		file_relo_core_homework_sttachment_affiliation_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_core_homework_sttachment_affiliation_type_proto_rawDescData)
	})
	return file_relo_core_homework_sttachment_affiliation_type_proto_rawDescData
}

var file_relo_core_homework_sttachment_affiliation_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_relo_core_homework_sttachment_affiliation_type_proto_goTypes = []interface{}{
	(HomeworkAttachmentAffiliationType)(0), // 0: relo.core.HomeworkAttachmentAffiliationType
}
var file_relo_core_homework_sttachment_affiliation_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relo_core_homework_sttachment_affiliation_type_proto_init() }
func file_relo_core_homework_sttachment_affiliation_type_proto_init() {
	if File_relo_core_homework_sttachment_affiliation_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relo_core_homework_sttachment_affiliation_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_core_homework_sttachment_affiliation_type_proto_goTypes,
		DependencyIndexes: file_relo_core_homework_sttachment_affiliation_type_proto_depIdxs,
		EnumInfos:         file_relo_core_homework_sttachment_affiliation_type_proto_enumTypes,
	}.Build()
	File_relo_core_homework_sttachment_affiliation_type_proto = out.File
	file_relo_core_homework_sttachment_affiliation_type_proto_rawDesc = nil
	file_relo_core_homework_sttachment_affiliation_type_proto_goTypes = nil
	file_relo_core_homework_sttachment_affiliation_type_proto_depIdxs = nil
}
