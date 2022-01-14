// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/core/meeting/mqtt_formatter.proto

package meeting

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

// MQTT格式化
type MqttFormatter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 面板
	// @gotags: json:"board,omitempty"
	Board string `protobuf:"bytes,2,opt,name=board,proto3" json:"board,omitempty"`
	// 用户
	// @gotags: json:"user,omitempty"
	User string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *MqttFormatter) Reset() {
	*x = MqttFormatter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_core_meeting_mqtt_formatter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MqttFormatter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MqttFormatter) ProtoMessage() {}

func (x *MqttFormatter) ProtoReflect() protoreflect.Message {
	mi := &file_relo_core_meeting_mqtt_formatter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MqttFormatter.ProtoReflect.Descriptor instead.
func (*MqttFormatter) Descriptor() ([]byte, []int) {
	return file_relo_core_meeting_mqtt_formatter_proto_rawDescGZIP(), []int{0}
}

func (x *MqttFormatter) GetBoard() string {
	if x != nil {
		return x.Board
	}
	return ""
}

func (x *MqttFormatter) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

var File_relo_core_meeting_mqtt_formatter_proto protoreflect.FileDescriptor

var file_relo_core_meeting_mqtt_formatter_proto_rawDesc = []byte{
	0x0a, 0x26, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x71, 0x74, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x39, 0x0a, 0x0d, 0x4d,
	0x71, 0x74, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_relo_core_meeting_mqtt_formatter_proto_rawDescOnce sync.Once
	file_relo_core_meeting_mqtt_formatter_proto_rawDescData = file_relo_core_meeting_mqtt_formatter_proto_rawDesc
)

func file_relo_core_meeting_mqtt_formatter_proto_rawDescGZIP() []byte {
	file_relo_core_meeting_mqtt_formatter_proto_rawDescOnce.Do(func() {
		file_relo_core_meeting_mqtt_formatter_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_core_meeting_mqtt_formatter_proto_rawDescData)
	})
	return file_relo_core_meeting_mqtt_formatter_proto_rawDescData
}

var file_relo_core_meeting_mqtt_formatter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_core_meeting_mqtt_formatter_proto_goTypes = []interface{}{
	(*MqttFormatter)(nil), // 0: relo.core.meeting.MqttFormatter
}
var file_relo_core_meeting_mqtt_formatter_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relo_core_meeting_mqtt_formatter_proto_init() }
func file_relo_core_meeting_mqtt_formatter_proto_init() {
	if File_relo_core_meeting_mqtt_formatter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_core_meeting_mqtt_formatter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MqttFormatter); i {
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
			RawDescriptor: file_relo_core_meeting_mqtt_formatter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_core_meeting_mqtt_formatter_proto_goTypes,
		DependencyIndexes: file_relo_core_meeting_mqtt_formatter_proto_depIdxs,
		MessageInfos:      file_relo_core_meeting_mqtt_formatter_proto_msgTypes,
	}.Build()
	File_relo_core_meeting_mqtt_formatter_proto = out.File
	file_relo_core_meeting_mqtt_formatter_proto_rawDesc = nil
	file_relo_core_meeting_mqtt_formatter_proto_goTypes = nil
	file_relo_core_meeting_mqtt_formatter_proto_depIdxs = nil
}