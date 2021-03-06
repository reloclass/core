// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/core/school/course.proto

package school

import (
	course "github.com/reloclass/core/course"
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

// 分辨率
type Resolution int32

const (
	// 未知
	Resolution_RESOLUTION_UNSPECIFIED Resolution = 0
	// 360P
	Resolution_P360 Resolution = 1
	// 表情
	Resolution_P480 Resolution = 2
	// 高清
	Resolution_P720 Resolution = 3
	// 超高清
	Resolution_P1080 Resolution = 4
	// 2K
	Resolution_K2 Resolution = 5
	// 4K
	Resolution_K4 Resolution = 6
	// 8K
	Resolution_K8 Resolution = 7
	// 原始分辨率
	Resolution_ORIGINAL Resolution = 15
)

// Enum value maps for Resolution.
var (
	Resolution_name = map[int32]string{
		0:  "RESOLUTION_UNSPECIFIED",
		1:  "P360",
		2:  "P480",
		3:  "P720",
		4:  "P1080",
		5:  "K2",
		6:  "K4",
		7:  "K8",
		15: "ORIGINAL",
	}
	Resolution_value = map[string]int32{
		"RESOLUTION_UNSPECIFIED": 0,
		"P360":                   1,
		"P480":                   2,
		"P720":                   3,
		"P1080":                  4,
		"K2":                     5,
		"K4":                     6,
		"K8":                     7,
		"ORIGINAL":               15,
	}
)

func (x Resolution) Enum() *Resolution {
	p := new(Resolution)
	*p = x
	return p
}

func (x Resolution) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Resolution) Descriptor() protoreflect.EnumDescriptor {
	return file_relo_core_school_course_proto_enumTypes[0].Descriptor()
}

func (Resolution) Type() protoreflect.EnumType {
	return &file_relo_core_school_course_proto_enumTypes[0]
}

func (x Resolution) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Resolution.Descriptor instead.
func (Resolution) EnumDescriptor() ([]byte, []int) {
	return file_relo_core_school_course_proto_rawDescGZIP(), []int{0}
}

// 开设课程
type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 课程类型
	// @gotags: json:"courseType" validate:"required" oneof:"1 3 4"
	Type course.Type `protobuf:"varint,1,opt,name=type,proto3,enum=relo.core.course.Type" json:"courseType" validate:"required" oneof:"1 3 4"`
	//  上课模式
	// @gotags: json:"classMode" validate:"omitempty" oneof:"1 3 4"
	ClassMode course.ClassMode `protobuf:"varint,2,opt,name=class_mode,json=classMode,proto3,enum=relo.core.course.ClassMode" json:"classMode" validate:"omitempty" oneof:"1 3 4"`
	// 最大上课人数
	// @gotags: json:"studentNumber" validate:"omitempty"
	StudentNumber int32 `protobuf:"varint,3,opt,name=student_number,json=studentNumber,proto3" json:"studentNumber" validate:"omitempty"`
	// 课程分辨率
	// @gotags: json:"courseResolution" validate:"omitempty"
	Resolution Resolution `protobuf:"varint,4,opt,name=resolution,proto3,enum=relo.core.school.Resolution" json:"courseResolution" validate:"omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_core_school_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_relo_core_school_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_relo_core_school_course_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetType() course.Type {
	if x != nil {
		return x.Type
	}
	return course.Type_COURSE_TYPE_UNSPECIFIED
}

func (x *Course) GetClassMode() course.ClassMode {
	if x != nil {
		return x.ClassMode
	}
	return course.ClassMode_CLASS_MODE_UNSPECIFIED
}

func (x *Course) GetStudentNumber() int32 {
	if x != nil {
		return x.StudentNumber
	}
	return 0
}

func (x *Course) GetResolution() Resolution {
	if x != nil {
		return x.Resolution
	}
	return Resolution_RESOLUTION_UNSPECIFIED
}

var File_relo_core_school_course_proto protoreflect.FileDescriptor

var file_relo_core_school_course_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f,
	0x6c, 0x1a, 0x21, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x6c,
	0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x72,
	0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x77, 0x0a, 0x0a, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x53, 0x4f, 0x4c,
	0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x33, 0x36, 0x30, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x50, 0x34, 0x38, 0x30, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x37, 0x32, 0x30, 0x10,
	0x03, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x31, 0x30, 0x38, 0x30, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02,
	0x4b, 0x32, 0x10, 0x05, 0x12, 0x06, 0x0a, 0x02, 0x4b, 0x34, 0x10, 0x06, 0x12, 0x06, 0x0a, 0x02,
	0x4b, 0x38, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x41, 0x4c,
	0x10, 0x0f, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_core_school_course_proto_rawDescOnce sync.Once
	file_relo_core_school_course_proto_rawDescData = file_relo_core_school_course_proto_rawDesc
)

func file_relo_core_school_course_proto_rawDescGZIP() []byte {
	file_relo_core_school_course_proto_rawDescOnce.Do(func() {
		file_relo_core_school_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_core_school_course_proto_rawDescData)
	})
	return file_relo_core_school_course_proto_rawDescData
}

var file_relo_core_school_course_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_relo_core_school_course_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_core_school_course_proto_goTypes = []interface{}{
	(Resolution)(0),       // 0: relo.core.school.Resolution
	(*Course)(nil),        // 1: relo.core.school.Course
	(course.Type)(0),      // 2: relo.core.course.Type
	(course.ClassMode)(0), // 3: relo.core.course.ClassMode
}
var file_relo_core_school_course_proto_depIdxs = []int32{
	2, // 0: relo.core.school.Course.type:type_name -> relo.core.course.Type
	3, // 1: relo.core.school.Course.class_mode:type_name -> relo.core.course.ClassMode
	0, // 2: relo.core.school.Course.resolution:type_name -> relo.core.school.Resolution
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_relo_core_school_course_proto_init() }
func file_relo_core_school_course_proto_init() {
	if File_relo_core_school_course_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_core_school_course_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
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
			RawDescriptor: file_relo_core_school_course_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_core_school_course_proto_goTypes,
		DependencyIndexes: file_relo_core_school_course_proto_depIdxs,
		EnumInfos:         file_relo_core_school_course_proto_enumTypes,
		MessageInfos:      file_relo_core_school_course_proto_msgTypes,
	}.Build()
	File_relo_core_school_course_proto = out.File
	file_relo_core_school_course_proto_rawDesc = nil
	file_relo_core_school_course_proto_goTypes = nil
	file_relo_core_school_course_proto_depIdxs = nil
}
