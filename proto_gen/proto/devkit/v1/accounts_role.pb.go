// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: proto/devkit/v1/accounts_role.proto

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

type RoleCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleName        string `protobuf:"bytes,1,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	RoleDescription string `protobuf:"bytes,2,opt,name=role_description,json=roleDescription,proto3" json:"role_description,omitempty"`
}

func (x *RoleCreateRequest) Reset() {
	*x = RoleCreateRequest{}
	mi := &file_proto_devkit_v1_accounts_role_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleCreateRequest) ProtoMessage() {}

func (x *RoleCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_devkit_v1_accounts_role_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleCreateRequest.ProtoReflect.Descriptor instead.
func (*RoleCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_devkit_v1_accounts_role_proto_rawDescGZIP(), []int{0}
}

func (x *RoleCreateRequest) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *RoleCreateRequest) GetRoleDescription() string {
	if x != nil {
		return x.RoleDescription
	}
	return ""
}

type AccountsSchemaRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId          int32  `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	RoleName        string `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"` // Unique
	RoleDescription string `protobuf:"bytes,3,opt,name=role_description,json=roleDescription,proto3" json:"role_description,omitempty"`
	CreatedAt       string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt       string `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *AccountsSchemaRole) Reset() {
	*x = AccountsSchemaRole{}
	mi := &file_proto_devkit_v1_accounts_role_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountsSchemaRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountsSchemaRole) ProtoMessage() {}

func (x *AccountsSchemaRole) ProtoReflect() protoreflect.Message {
	mi := &file_proto_devkit_v1_accounts_role_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountsSchemaRole.ProtoReflect.Descriptor instead.
func (*AccountsSchemaRole) Descriptor() ([]byte, []int) {
	return file_proto_devkit_v1_accounts_role_proto_rawDescGZIP(), []int{1}
}

func (x *AccountsSchemaRole) GetRoleId() int32 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *AccountsSchemaRole) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *AccountsSchemaRole) GetRoleDescription() string {
	if x != nil {
		return x.RoleDescription
	}
	return ""
}

func (x *AccountsSchemaRole) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AccountsSchemaRole) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *AccountsSchemaRole) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type RoleCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role *AccountsSchemaRole `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *RoleCreateResponse) Reset() {
	*x = RoleCreateResponse{}
	mi := &file_proto_devkit_v1_accounts_role_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleCreateResponse) ProtoMessage() {}

func (x *RoleCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_devkit_v1_accounts_role_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleCreateResponse.ProtoReflect.Descriptor instead.
func (*RoleCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_devkit_v1_accounts_role_proto_rawDescGZIP(), []int{2}
}

func (x *RoleCreateResponse) GetRole() *AccountsSchemaRole {
	if x != nil {
		return x.Role
	}
	return nil
}

var File_proto_devkit_v1_accounts_role_proto protoreflect.FileDescriptor

var file_proto_devkit_v1_accounts_role_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x6b, 0x69, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x65, 0x76, 0x6b, 0x69, 0x74, 0x2e, 0x76, 0x31,
	0x22, 0x5b, 0x0a, 0x11, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x6f,
	0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd2, 0x01,
	0x0a, 0x12, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x6f, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x47, 0x0a, 0x12, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x65, 0x76, 0x6b, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0xac, 0x01, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x76, 0x6b, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x61, 0x72, 0x77, 0x69, 0x73, 0x68, 0x64, 0x65, 0x76, 0x2f, 0x64, 0x65, 0x76, 0x6b, 0x69, 0x74,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x6b, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x64,
	0x65, 0x76, 0x6b, 0x69, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x09,
	0x44, 0x65, 0x76, 0x6b, 0x69, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x44, 0x65, 0x76, 0x6b,
	0x69, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x44, 0x65, 0x76, 0x6b, 0x69, 0x74, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a,
	0x44, 0x65, 0x76, 0x6b, 0x69, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_devkit_v1_accounts_role_proto_rawDescOnce sync.Once
	file_proto_devkit_v1_accounts_role_proto_rawDescData = file_proto_devkit_v1_accounts_role_proto_rawDesc
)

func file_proto_devkit_v1_accounts_role_proto_rawDescGZIP() []byte {
	file_proto_devkit_v1_accounts_role_proto_rawDescOnce.Do(func() {
		file_proto_devkit_v1_accounts_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_devkit_v1_accounts_role_proto_rawDescData)
	})
	return file_proto_devkit_v1_accounts_role_proto_rawDescData
}

var file_proto_devkit_v1_accounts_role_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_devkit_v1_accounts_role_proto_goTypes = []any{
	(*RoleCreateRequest)(nil),  // 0: devkit.v1.RoleCreateRequest
	(*AccountsSchemaRole)(nil), // 1: devkit.v1.AccountsSchemaRole
	(*RoleCreateResponse)(nil), // 2: devkit.v1.RoleCreateResponse
}
var file_proto_devkit_v1_accounts_role_proto_depIdxs = []int32{
	1, // 0: devkit.v1.RoleCreateResponse.role:type_name -> devkit.v1.AccountsSchemaRole
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_devkit_v1_accounts_role_proto_init() }
func file_proto_devkit_v1_accounts_role_proto_init() {
	if File_proto_devkit_v1_accounts_role_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_devkit_v1_accounts_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_devkit_v1_accounts_role_proto_goTypes,
		DependencyIndexes: file_proto_devkit_v1_accounts_role_proto_depIdxs,
		MessageInfos:      file_proto_devkit_v1_accounts_role_proto_msgTypes,
	}.Build()
	File_proto_devkit_v1_accounts_role_proto = out.File
	file_proto_devkit_v1_accounts_role_proto_rawDesc = nil
	file_proto_devkit_v1_accounts_role_proto_goTypes = nil
	file_proto_devkit_v1_accounts_role_proto_depIdxs = nil
}
