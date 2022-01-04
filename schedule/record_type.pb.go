// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/core/schedule/record_type.proto

package schedule

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

// 录制操作类型
type RecordType int32

const (
	// 未知
	RecordType_RECORD_TYPE_UNSPECIFIED RecordType = 0
	// 不进行录制
	RecordType_RECORD_NO RecordType = 1
	// 仅仅进行录制
	RecordType_ONLY RecordType = 2
	// 录制并可以进行回放
	RecordType_PLAY_BACK RecordType = 3
)

// Enum value maps for RecordType.
var (
	RecordType_name = map[int32]string{
		0: "RECORD_TYPE_UNSPECIFIED",
		1: "RECORD_NO",
		2: "ONLY",
		3: "PLAY_BACK",
	}
	RecordType_value = map[string]int32{
		"RECORD_TYPE_UNSPECIFIED": 0,
		"RECORD_NO":               1,
		"ONLY":                    2,
		"PLAY_BACK":               3,
	}
)

func (x RecordType) Enum() *RecordType {
	p := new(RecordType)
	*p = x
	return p
}

func (x RecordType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecordType) Descriptor() protoreflect.EnumDescriptor {
	return file_relo_core_schedule_record_type_proto_enumTypes[0].Descriptor()
}

func (RecordType) Type() protoreflect.EnumType {
	return &file_relo_core_schedule_record_type_proto_enumTypes[0]
}

func (x RecordType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecordType.Descriptor instead.
func (RecordType) EnumDescriptor() ([]byte, []int) {
	return file_relo_core_schedule_record_type_proto_rawDescGZIP(), []int{0}
}

var File_relo_core_schedule_record_type_proto protoreflect.FileDescriptor

var file_relo_core_schedule_record_type_proto_rawDesc = []byte{
	0x0a, 0x24, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2a, 0x51, 0x0a, 0x0a, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x43, 0x4f,
	0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f,
	0x4e, 0x4f, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x03, 0x42, 0x24, 0x5a,
	0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_core_schedule_record_type_proto_rawDescOnce sync.Once
	file_relo_core_schedule_record_type_proto_rawDescData = file_relo_core_schedule_record_type_proto_rawDesc
)

func file_relo_core_schedule_record_type_proto_rawDescGZIP() []byte {
	file_relo_core_schedule_record_type_proto_rawDescOnce.Do(func() {
		file_relo_core_schedule_record_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_core_schedule_record_type_proto_rawDescData)
	})
	return file_relo_core_schedule_record_type_proto_rawDescData
}

var file_relo_core_schedule_record_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_relo_core_schedule_record_type_proto_goTypes = []interface{}{
	(RecordType)(0), // 0: relo.core.schedule.RecordType
}
var file_relo_core_schedule_record_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relo_core_schedule_record_type_proto_init() }
func file_relo_core_schedule_record_type_proto_init() {
	if File_relo_core_schedule_record_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relo_core_schedule_record_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_core_schedule_record_type_proto_goTypes,
		DependencyIndexes: file_relo_core_schedule_record_type_proto_depIdxs,
		EnumInfos:         file_relo_core_schedule_record_type_proto_enumTypes,
	}.Build()
	File_relo_core_schedule_record_type_proto = out.File
	file_relo_core_schedule_record_type_proto_rawDesc = nil
	file_relo_core_schedule_record_type_proto_goTypes = nil
	file_relo_core_schedule_record_type_proto_depIdxs = nil
}
