// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/core/course_type.proto

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

type CourseType int32

const (
	// 未知
	CourseType_COURSE_TYPE_UNSPECIFIED CourseType = 0
	// 互动课
	CourseType_COURSE_TYPE_INTERACTIVE_CLASS CourseType = 1
	// 录播课
	CourseType_COURSE_TYPE_RECORD CourseType = 2
	// 公开课
	CourseType_COURSE_TYPE_PUBLIC CourseType = 3
)

// Enum value maps for CourseType.
var (
	CourseType_name = map[int32]string{
		0: "COURSE_TYPE_UNSPECIFIED",
		1: "COURSE_TYPE_INTERACTIVE_CLASS",
		2: "COURSE_TYPE_RECORD",
		3: "COURSE_TYPE_PUBLIC",
	}
	CourseType_value = map[string]int32{
		"COURSE_TYPE_UNSPECIFIED":       0,
		"COURSE_TYPE_INTERACTIVE_CLASS": 1,
		"COURSE_TYPE_RECORD":            2,
		"COURSE_TYPE_PUBLIC":            3,
	}
)

func (x CourseType) Enum() *CourseType {
	p := new(CourseType)
	*p = x
	return p
}

func (x CourseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CourseType) Descriptor() protoreflect.EnumDescriptor {
	return file_relo_core_course_type_proto_enumTypes[0].Descriptor()
}

func (CourseType) Type() protoreflect.EnumType {
	return &file_relo_core_course_type_proto_enumTypes[0]
}

func (x CourseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CourseType.Descriptor instead.
func (CourseType) EnumDescriptor() ([]byte, []int) {
	return file_relo_core_course_type_proto_rawDescGZIP(), []int{0}
}

var File_relo_core_course_type_proto protoreflect.FileDescriptor

var file_relo_core_course_type_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72,
	0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2a, 0x7c, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x43,
	0x4c, 0x41, 0x53, 0x53, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x10, 0x02, 0x12, 0x16,
	0x0a, 0x12, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x55,
	0x42, 0x4c, 0x49, 0x43, 0x10, 0x03, 0x42, 0x1b, 0x5a, 0x19, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_core_course_type_proto_rawDescOnce sync.Once
	file_relo_core_course_type_proto_rawDescData = file_relo_core_course_type_proto_rawDesc
)

func file_relo_core_course_type_proto_rawDescGZIP() []byte {
	file_relo_core_course_type_proto_rawDescOnce.Do(func() {
		file_relo_core_course_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_core_course_type_proto_rawDescData)
	})
	return file_relo_core_course_type_proto_rawDescData
}

var file_relo_core_course_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_relo_core_course_type_proto_goTypes = []interface{}{
	(CourseType)(0), // 0: relo.core.CourseType
}
var file_relo_core_course_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relo_core_course_type_proto_init() }
func file_relo_core_course_type_proto_init() {
	if File_relo_core_course_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relo_core_course_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_core_course_type_proto_goTypes,
		DependencyIndexes: file_relo_core_course_type_proto_depIdxs,
		EnumInfos:         file_relo_core_course_type_proto_enumTypes,
	}.Build()
	File_relo_core_course_type_proto = out.File
	file_relo_core_course_type_proto_rawDesc = nil
	file_relo_core_course_type_proto_goTypes = nil
	file_relo_core_course_type_proto_depIdxs = nil
}
