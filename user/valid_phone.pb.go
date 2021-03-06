// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/core/user/valid_phone.proto

package user

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

// 有效手机号
type ValidPhone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 手机号
	// @gotags: json:"phone" validate:"omitempty,mobile"
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone" validate:"omitempty,mobile"`
}

func (x *ValidPhone) Reset() {
	*x = ValidPhone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_core_user_valid_phone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidPhone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidPhone) ProtoMessage() {}

func (x *ValidPhone) ProtoReflect() protoreflect.Message {
	mi := &file_relo_core_user_valid_phone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidPhone.ProtoReflect.Descriptor instead.
func (*ValidPhone) Descriptor() ([]byte, []int) {
	return file_relo_core_user_valid_phone_proto_rawDescGZIP(), []int{0}
}

func (x *ValidPhone) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_relo_core_user_valid_phone_proto protoreflect.FileDescriptor

var file_relo_core_user_valid_phone_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x22, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_core_user_valid_phone_proto_rawDescOnce sync.Once
	file_relo_core_user_valid_phone_proto_rawDescData = file_relo_core_user_valid_phone_proto_rawDesc
)

func file_relo_core_user_valid_phone_proto_rawDescGZIP() []byte {
	file_relo_core_user_valid_phone_proto_rawDescOnce.Do(func() {
		file_relo_core_user_valid_phone_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_core_user_valid_phone_proto_rawDescData)
	})
	return file_relo_core_user_valid_phone_proto_rawDescData
}

var file_relo_core_user_valid_phone_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_core_user_valid_phone_proto_goTypes = []interface{}{
	(*ValidPhone)(nil), // 0: relo.core.user.ValidPhone
}
var file_relo_core_user_valid_phone_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relo_core_user_valid_phone_proto_init() }
func file_relo_core_user_valid_phone_proto_init() {
	if File_relo_core_user_valid_phone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_core_user_valid_phone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidPhone); i {
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
			RawDescriptor: file_relo_core_user_valid_phone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_core_user_valid_phone_proto_goTypes,
		DependencyIndexes: file_relo_core_user_valid_phone_proto_depIdxs,
		MessageInfos:      file_relo_core_user_valid_phone_proto_msgTypes,
	}.Build()
	File_relo_core_user_valid_phone_proto = out.File
	file_relo_core_user_valid_phone_proto_rawDesc = nil
	file_relo_core_user_valid_phone_proto_goTypes = nil
	file_relo_core_user_valid_phone_proto_depIdxs = nil
}
