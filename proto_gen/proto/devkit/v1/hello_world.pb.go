// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: proto/devkit/v1/hello_world.proto

package devkitv1

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

type HelloWorldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloWorldRequest) Reset() {
	*x = HelloWorldRequest{}
	mi := &file_proto_devkit_v1_hello_world_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloWorldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldRequest) ProtoMessage() {}

func (x *HelloWorldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_devkit_v1_hello_world_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldRequest.ProtoReflect.Descriptor instead.
func (*HelloWorldRequest) Descriptor() ([]byte, []int) {
	return file_proto_devkit_v1_hello_world_proto_rawDescGZIP(), []int{0}
}

func (x *HelloWorldRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloWorldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greet string `protobuf:"bytes,1,opt,name=greet,proto3" json:"greet,omitempty"`
}

func (x *HelloWorldResponse) Reset() {
	*x = HelloWorldResponse{}
	mi := &file_proto_devkit_v1_hello_world_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloWorldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldResponse) ProtoMessage() {}

func (x *HelloWorldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_devkit_v1_hello_world_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldResponse.ProtoReflect.Descriptor instead.
func (*HelloWorldResponse) Descriptor() ([]byte, []int) {
	return file_proto_devkit_v1_hello_world_proto_rawDescGZIP(), []int{1}
}

func (x *HelloWorldResponse) GetGreet() string {
	if x != nil {
		return x.Greet
	}
	return ""
}

var File_proto_devkit_v1_hello_world_proto protoreflect.FileDescriptor

var file_proto_devkit_v1_hello_world_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x6b, 0x69, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x65, 0x76, 0x6b, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x27,
	0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x12, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x42, 0xaa, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x76, 0x6b,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x72, 0x77, 0x69, 0x73, 0x68, 0x64, 0x65, 0x76, 0x2f,
	0x64, 0x65, 0x76, 0x6b, 0x69, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x6b, 0x69,
	0x74, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x65, 0x76, 0x6b, 0x69, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x44, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x44, 0x65, 0x76, 0x6b, 0x69, 0x74, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x09, 0x44, 0x65, 0x76, 0x6b, 0x69, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x44, 0x65,
	0x76, 0x6b, 0x69, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x44, 0x65, 0x76, 0x6b, 0x69, 0x74, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_devkit_v1_hello_world_proto_rawDescOnce sync.Once
	file_proto_devkit_v1_hello_world_proto_rawDescData = file_proto_devkit_v1_hello_world_proto_rawDesc
)

func file_proto_devkit_v1_hello_world_proto_rawDescGZIP() []byte {
	file_proto_devkit_v1_hello_world_proto_rawDescOnce.Do(func() {
		file_proto_devkit_v1_hello_world_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_devkit_v1_hello_world_proto_rawDescData)
	})
	return file_proto_devkit_v1_hello_world_proto_rawDescData
}

var file_proto_devkit_v1_hello_world_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_devkit_v1_hello_world_proto_goTypes = []any{
	(*HelloWorldRequest)(nil),  // 0: devkit.v1.HelloWorldRequest
	(*HelloWorldResponse)(nil), // 1: devkit.v1.HelloWorldResponse
}
var file_proto_devkit_v1_hello_world_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_devkit_v1_hello_world_proto_init() }
func file_proto_devkit_v1_hello_world_proto_init() {
	if File_proto_devkit_v1_hello_world_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_devkit_v1_hello_world_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_devkit_v1_hello_world_proto_goTypes,
		DependencyIndexes: file_proto_devkit_v1_hello_world_proto_depIdxs,
		MessageInfos:      file_proto_devkit_v1_hello_world_proto_msgTypes,
	}.Build()
	File_proto_devkit_v1_hello_world_proto = out.File
	file_proto_devkit_v1_hello_world_proto_rawDesc = nil
	file_proto_devkit_v1_hello_world_proto_goTypes = nil
	file_proto_devkit_v1_hello_world_proto_depIdxs = nil
}
