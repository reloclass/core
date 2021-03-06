// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/core/lesson/user.proto

package lesson

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

// 课节用户信息
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户编号
	// @gotags: json:"id,string"
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,string"`
	// 用户信息
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// 电话
	Phone string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	// 用户状态
	Status UserStatus `protobuf:"varint,5,opt,name=status,proto3,enum=relo.core.lesson.UserStatus" json:"status,omitempty"`
	// 是否禁言
	Forbidden bool `protobuf:"varint,6,opt,name=forbidden,proto3" json:"forbidden,omitempty"`
	// 设备信息
	Equipments []*Equipment `protobuf:"bytes,7,rep,name=equipments,proto3" json:"equipments,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_core_lesson_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_relo_core_lesson_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_relo_core_lesson_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetStatus() UserStatus {
	if x != nil {
		return x.Status
	}
	return UserStatus_NOT_ENTER
}

func (x *User) GetForbidden() bool {
	if x != nil {
		return x.Forbidden
	}
	return false
}

func (x *User) GetEquipments() []*Equipment {
	if x != nil {
		return x.Equipments
	}
	return nil
}

var File_relo_core_lesson_user_proto protoreflect.FileDescriptor

var file_relo_core_lesson_user_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x72,
	0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x1a,
	0x22, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x66, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x3b, 0x0a, 0x0a,
	0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x65,
	0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_core_lesson_user_proto_rawDescOnce sync.Once
	file_relo_core_lesson_user_proto_rawDescData = file_relo_core_lesson_user_proto_rawDesc
)

func file_relo_core_lesson_user_proto_rawDescGZIP() []byte {
	file_relo_core_lesson_user_proto_rawDescOnce.Do(func() {
		file_relo_core_lesson_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_core_lesson_user_proto_rawDescData)
	})
	return file_relo_core_lesson_user_proto_rawDescData
}

var file_relo_core_lesson_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_core_lesson_user_proto_goTypes = []interface{}{
	(*User)(nil),      // 0: relo.core.lesson.User
	(UserStatus)(0),   // 1: relo.core.lesson.UserStatus
	(*Equipment)(nil), // 2: relo.core.lesson.Equipment
}
var file_relo_core_lesson_user_proto_depIdxs = []int32{
	1, // 0: relo.core.lesson.User.status:type_name -> relo.core.lesson.UserStatus
	2, // 1: relo.core.lesson.User.equipments:type_name -> relo.core.lesson.Equipment
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_relo_core_lesson_user_proto_init() }
func file_relo_core_lesson_user_proto_init() {
	if File_relo_core_lesson_user_proto != nil {
		return
	}
	file_relo_core_lesson_user_status_proto_init()
	file_relo_core_lesson_equipment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_relo_core_lesson_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_relo_core_lesson_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_core_lesson_user_proto_goTypes,
		DependencyIndexes: file_relo_core_lesson_user_proto_depIdxs,
		MessageInfos:      file_relo_core_lesson_user_proto_msgTypes,
	}.Build()
	File_relo_core_lesson_user_proto = out.File
	file_relo_core_lesson_user_proto_rawDesc = nil
	file_relo_core_lesson_user_proto_goTypes = nil
	file_relo_core_lesson_user_proto_depIdxs = nil
}
