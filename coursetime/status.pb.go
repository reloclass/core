// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/core/coursetime/status.proto

package coursetime

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

// 课节状态
type Status int32

const (
	// 未开始
	Status_UNSTART Status = 0
	// 预课中
	Status_PRIOR Status = 1
	// 进行中
	Status_STARTING Status = 2
	// 结束
	Status_END Status = 3
	// 拖堂
	Status_DELAYED Status = 4
	// 已经中止的
	Status_SUSPENDED Status = 5
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "UNSTART",
		1: "PRIOR",
		2: "STARTING",
		3: "END",
		4: "DELAYED",
		5: "SUSPENDED",
	}
	Status_value = map[string]int32{
		"UNSTART":   0,
		"PRIOR":     1,
		"STARTING":  2,
		"END":       3,
		"DELAYED":   4,
		"SUSPENDED": 5,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_relo_core_coursetime_status_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_relo_core_coursetime_status_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_relo_core_coursetime_status_proto_rawDescGZIP(), []int{0}
}

var File_relo_core_coursetime_status_proto protoreflect.FileDescriptor

var file_relo_core_coursetime_status_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2a, 0x53, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53,
	0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x4e, 0x44,
	0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x41, 0x59, 0x45, 0x44, 0x10, 0x04, 0x12,
	0x0d, 0x0a, 0x09, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x05, 0x42, 0x26,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c,
	0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_core_coursetime_status_proto_rawDescOnce sync.Once
	file_relo_core_coursetime_status_proto_rawDescData = file_relo_core_coursetime_status_proto_rawDesc
)

func file_relo_core_coursetime_status_proto_rawDescGZIP() []byte {
	file_relo_core_coursetime_status_proto_rawDescOnce.Do(func() {
		file_relo_core_coursetime_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_core_coursetime_status_proto_rawDescData)
	})
	return file_relo_core_coursetime_status_proto_rawDescData
}

var file_relo_core_coursetime_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_relo_core_coursetime_status_proto_goTypes = []interface{}{
	(Status)(0), // 0: relo.core.coursetime.Status
}
var file_relo_core_coursetime_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relo_core_coursetime_status_proto_init() }
func file_relo_core_coursetime_status_proto_init() {
	if File_relo_core_coursetime_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relo_core_coursetime_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_core_coursetime_status_proto_goTypes,
		DependencyIndexes: file_relo_core_coursetime_status_proto_depIdxs,
		EnumInfos:         file_relo_core_coursetime_status_proto_enumTypes,
	}.Build()
	File_relo_core_coursetime_status_proto = out.File
	file_relo_core_coursetime_status_proto_rawDesc = nil
	file_relo_core_coursetime_status_proto_goTypes = nil
	file_relo_core_coursetime_status_proto_depIdxs = nil
}