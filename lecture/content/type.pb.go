// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/core/lecture/content/type.proto

package content

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

// 讲次内容资料类型
type Type int32

const (
	// 视频
	Type_VIDEO Type = 0
	// 资料
	Type_MATERIAL Type = 1
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "VIDEO",
		1: "MATERIAL",
	}
	Type_value = map[string]int32{
		"VIDEO":    0,
		"MATERIAL": 1,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_relo_core_lecture_content_type_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_relo_core_lecture_content_type_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_relo_core_lecture_content_type_proto_rawDescGZIP(), []int{0}
}

var File_relo_core_lecture_content_type_proto protoreflect.FileDescriptor

var file_relo_core_lecture_content_type_proto_rawDesc = []byte{
	0x0a, 0x24, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2a, 0x1f, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x49, 0x44,
	0x45, 0x4f, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x41, 0x54, 0x45, 0x52, 0x49, 0x41, 0x4c,
	0x10, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_core_lecture_content_type_proto_rawDescOnce sync.Once
	file_relo_core_lecture_content_type_proto_rawDescData = file_relo_core_lecture_content_type_proto_rawDesc
)

func file_relo_core_lecture_content_type_proto_rawDescGZIP() []byte {
	file_relo_core_lecture_content_type_proto_rawDescOnce.Do(func() {
		file_relo_core_lecture_content_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_core_lecture_content_type_proto_rawDescData)
	})
	return file_relo_core_lecture_content_type_proto_rawDescData
}

var file_relo_core_lecture_content_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_relo_core_lecture_content_type_proto_goTypes = []interface{}{
	(Type)(0), // 0: relo.core.lecture.content.Type
}
var file_relo_core_lecture_content_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relo_core_lecture_content_type_proto_init() }
func file_relo_core_lecture_content_type_proto_init() {
	if File_relo_core_lecture_content_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relo_core_lecture_content_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_core_lecture_content_type_proto_goTypes,
		DependencyIndexes: file_relo_core_lecture_content_type_proto_depIdxs,
		EnumInfos:         file_relo_core_lecture_content_type_proto_enumTypes,
	}.Build()
	File_relo_core_lecture_content_type_proto = out.File
	file_relo_core_lecture_content_type_proto_rawDesc = nil
	file_relo_core_lecture_content_type_proto_goTypes = nil
	file_relo_core_lecture_content_type_proto_depIdxs = nil
}
