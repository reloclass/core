// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/core/cloudfile/operation_type.proto

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

// 文件操作类型
type OperationType int32

const (
	// 保留
	OperationType_RESERVE OperationType = 0
	// 直接下载
	OperationType_DOWNLOAD OperationType = 1
	// 打开
	OperationType_OPEN OperationType = 2
	// 编辑
	OperationType_EDIT OperationType = 3
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "RESERVE",
		1: "DOWNLOAD",
		2: "OPEN",
		3: "EDIT",
	}
	OperationType_value = map[string]int32{
		"RESERVE":  0,
		"DOWNLOAD": 1,
		"OPEN":     2,
		"EDIT":     3,
	}
)

func (x OperationType) Enum() *OperationType {
	p := new(OperationType)
	*p = x
	return p
}

func (x OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_relo_core_cloudfile_operation_type_proto_enumTypes[0].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_relo_core_cloudfile_operation_type_proto_enumTypes[0]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_relo_core_cloudfile_operation_type_proto_rawDescGZIP(), []int{0}
}

var File_relo_core_cloudfile_operation_type_proto protoreflect.FileDescriptor

var file_relo_core_cloudfile_operation_type_proto_rawDesc = []byte{
	0x0a, 0x28, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x72, 0x65, 0x6c, 0x6f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x69, 0x6c, 0x65, 0x2a,
	0x3e, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4f,
	0x50, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x44, 0x49, 0x54, 0x10, 0x03, 0x42,
	0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65,
	0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_core_cloudfile_operation_type_proto_rawDescOnce sync.Once
	file_relo_core_cloudfile_operation_type_proto_rawDescData = file_relo_core_cloudfile_operation_type_proto_rawDesc
)

func file_relo_core_cloudfile_operation_type_proto_rawDescGZIP() []byte {
	file_relo_core_cloudfile_operation_type_proto_rawDescOnce.Do(func() {
		file_relo_core_cloudfile_operation_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_core_cloudfile_operation_type_proto_rawDescData)
	})
	return file_relo_core_cloudfile_operation_type_proto_rawDescData
}

var file_relo_core_cloudfile_operation_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_relo_core_cloudfile_operation_type_proto_goTypes = []interface{}{
	(OperationType)(0), // 0: relo.core.cloudfile.OperationType
}
var file_relo_core_cloudfile_operation_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relo_core_cloudfile_operation_type_proto_init() }
func file_relo_core_cloudfile_operation_type_proto_init() {
	if File_relo_core_cloudfile_operation_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relo_core_cloudfile_operation_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_core_cloudfile_operation_type_proto_goTypes,
		DependencyIndexes: file_relo_core_cloudfile_operation_type_proto_depIdxs,
		EnumInfos:         file_relo_core_cloudfile_operation_type_proto_enumTypes,
	}.Build()
	File_relo_core_cloudfile_operation_type_proto = out.File
	file_relo_core_cloudfile_operation_type_proto_rawDesc = nil
	file_relo_core_cloudfile_operation_type_proto_goTypes = nil
	file_relo_core_cloudfile_operation_type_proto_depIdxs = nil
}
