// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: api/bye/v1/bye.proto

package bye

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The request message containing the user's name.
type ByeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ByeRequest) Reset() {
	*x = ByeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bye_v1_bye_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByeRequest) ProtoMessage() {}

func (x *ByeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_bye_v1_bye_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByeRequest.ProtoReflect.Descriptor instead.
func (*ByeRequest) Descriptor() ([]byte, []int) {
	return file_api_bye_v1_bye_proto_rawDescGZIP(), []int{0}
}

func (x *ByeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type ByeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ByeReply) Reset() {
	*x = ByeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bye_v1_bye_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByeReply) ProtoMessage() {}

func (x *ByeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_bye_v1_bye_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByeReply.ProtoReflect.Descriptor instead.
func (*ByeReply) Descriptor() ([]byte, []int) {
	return file_api_bye_v1_bye_proto_rawDescGZIP(), []int{1}
}

func (x *ByeReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_bye_v1_bye_proto protoreflect.FileDescriptor

var file_api_bye_v1_bye_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x79, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x79, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x70, 0x69, 0x2e, 0x62, 0x79, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x20, 0x0a, 0x0a, 0x42, 0x79, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x42, 0x79, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x24, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x79, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x14, 0x73, 0x70, 0x6f, 0x72, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x62, 0x79, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x79, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_bye_v1_bye_proto_rawDescOnce sync.Once
	file_api_bye_v1_bye_proto_rawDescData = file_api_bye_v1_bye_proto_rawDesc
)

func file_api_bye_v1_bye_proto_rawDescGZIP() []byte {
	file_api_bye_v1_bye_proto_rawDescOnce.Do(func() {
		file_api_bye_v1_bye_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_bye_v1_bye_proto_rawDescData)
	})
	return file_api_bye_v1_bye_proto_rawDescData
}

var file_api_bye_v1_bye_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_bye_v1_bye_proto_goTypes = []interface{}{
	(*ByeRequest)(nil), // 0: Api.bye.v1.ByeRequest
	(*ByeReply)(nil),   // 1: Api.bye.v1.ByeReply
}
var file_api_bye_v1_bye_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_bye_v1_bye_proto_init() }
func file_api_bye_v1_bye_proto_init() {
	if File_api_bye_v1_bye_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_bye_v1_bye_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByeRequest); i {
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
		file_api_bye_v1_bye_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByeReply); i {
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
			RawDescriptor: file_api_bye_v1_bye_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_bye_v1_bye_proto_goTypes,
		DependencyIndexes: file_api_bye_v1_bye_proto_depIdxs,
		MessageInfos:      file_api_bye_v1_bye_proto_msgTypes,
	}.Build()
	File_api_bye_v1_bye_proto = out.File
	file_api_bye_v1_bye_proto_rawDesc = nil
	file_api_bye_v1_bye_proto_goTypes = nil
	file_api_bye_v1_bye_proto_depIdxs = nil
}
