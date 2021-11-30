// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/core/question_status.proto

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

// 问题状态
type QuestionStatus int32

const (
	// 未阅读
	QuestionStatus_UNREAD QuestionStatus = 0
	// 已经阅读
	QuestionStatus_READ QuestionStatus = 1
	// 已经回答
	QuestionStatus_ANSWERED QuestionStatus = 2
)

// Enum value maps for QuestionStatus.
var (
	QuestionStatus_name = map[int32]string{
		0: "UNREAD",
		1: "READ",
		2: "ANSWERED",
	}
	QuestionStatus_value = map[string]int32{
		"UNREAD":   0,
		"READ":     1,
		"ANSWERED": 2,
	}
)

func (x QuestionStatus) Enum() *QuestionStatus {
	p := new(QuestionStatus)
	*p = x
	return p
}

func (x QuestionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QuestionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_relo_core_question_status_proto_enumTypes[0].Descriptor()
}

func (QuestionStatus) Type() protoreflect.EnumType {
	return &file_relo_core_question_status_proto_enumTypes[0]
}

func (x QuestionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QuestionStatus.Descriptor instead.
func (QuestionStatus) EnumDescriptor() ([]byte, []int) {
	return file_relo_core_question_status_proto_rawDescGZIP(), []int{0}
}

var File_relo_core_question_status_proto protoreflect.FileDescriptor

var file_relo_core_question_status_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2a, 0x34, 0x0a, 0x0e,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a,
	0x0a, 0x06, 0x55, 0x4e, 0x52, 0x45, 0x41, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45,
	0x41, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x45, 0x44,
	0x10, 0x02, 0x42, 0x1b, 0x5a, 0x19, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_core_question_status_proto_rawDescOnce sync.Once
	file_relo_core_question_status_proto_rawDescData = file_relo_core_question_status_proto_rawDesc
)

func file_relo_core_question_status_proto_rawDescGZIP() []byte {
	file_relo_core_question_status_proto_rawDescOnce.Do(func() {
		file_relo_core_question_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_core_question_status_proto_rawDescData)
	})
	return file_relo_core_question_status_proto_rawDescData
}

var file_relo_core_question_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_relo_core_question_status_proto_goTypes = []interface{}{
	(QuestionStatus)(0), // 0: relo.core.QuestionStatus
}
var file_relo_core_question_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relo_core_question_status_proto_init() }
func file_relo_core_question_status_proto_init() {
	if File_relo_core_question_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relo_core_question_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_core_question_status_proto_goTypes,
		DependencyIndexes: file_relo_core_question_status_proto_depIdxs,
		EnumInfos:         file_relo_core_question_status_proto_enumTypes,
	}.Build()
	File_relo_core_question_status_proto = out.File
	file_relo_core_question_status_proto_rawDesc = nil
	file_relo_core_question_status_proto_goTypes = nil
	file_relo_core_question_status_proto_depIdxs = nil
}
