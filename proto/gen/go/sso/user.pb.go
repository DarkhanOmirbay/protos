// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0--rc2
// source: sso/user.proto

package ssov1

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fname    string `protobuf:"bytes,1,opt,name=fname,proto3" json:"fname,omitempty"`
	Lname    string `protobuf:"bytes,2,opt,name=lname,proto3" json:"lname,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_sso_user_proto_msgTypes[0]
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
	return file_sso_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetFname() string {
	if x != nil {
		return x.Fname
	}
	return ""
}

func (x *User) GetLname() string {
	if x != nil {
		return x.Lname
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type EditProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	User *User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *EditProfileRequest) Reset() {
	*x = EditProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditProfileRequest) ProtoMessage() {}

func (x *EditProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditProfileRequest.ProtoReflect.Descriptor instead.
func (*EditProfileRequest) Descriptor() ([]byte, []int) {
	return file_sso_user_proto_rawDescGZIP(), []int{1}
}

func (x *EditProfileRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EditProfileRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type EditProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg         string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	UpdatedUser *User  `protobuf:"bytes,2,opt,name=updatedUser,proto3" json:"updatedUser,omitempty"`
}

func (x *EditProfileResponse) Reset() {
	*x = EditProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditProfileResponse) ProtoMessage() {}

func (x *EditProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditProfileResponse.ProtoReflect.Descriptor instead.
func (*EditProfileResponse) Descriptor() ([]byte, []int) {
	return file_sso_user_proto_rawDescGZIP(), []int{2}
}

func (x *EditProfileResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *EditProfileResponse) GetUpdatedUser() *User {
	if x != nil {
		return x.UpdatedUser
	}
	return nil
}

type DeleteAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAccountRequest) Reset() {
	*x = DeleteAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountRequest) ProtoMessage() {}

func (x *DeleteAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountRequest.ProtoReflect.Descriptor instead.
func (*DeleteAccountRequest) Descriptor() ([]byte, []int) {
	return file_sso_user_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteAccountRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DeleteAccountResponse) Reset() {
	*x = DeleteAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountResponse) ProtoMessage() {}

func (x *DeleteAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountResponse.ProtoReflect.Descriptor instead.
func (*DeleteAccountResponse) Descriptor() ([]byte, []int) {
	return file_sso_user_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteAccountResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ShowProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ShowProfileRequest) Reset() {
	*x = ShowProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowProfileRequest) ProtoMessage() {}

func (x *ShowProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowProfileRequest.ProtoReflect.Descriptor instead.
func (*ShowProfileRequest) Descriptor() ([]byte, []int) {
	return file_sso_user_proto_rawDescGZIP(), []int{5}
}

func (x *ShowProfileRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ShowProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ShowProfileResponse) Reset() {
	*x = ShowProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowProfileResponse) ProtoMessage() {}

func (x *ShowProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowProfileResponse.ProtoReflect.Descriptor instead.
func (*ShowProfileResponse) Descriptor() ([]byte, []int) {
	return file_sso_user_proto_rawDescGZIP(), []int{6}
}

func (x *ShowProfileResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_sso_user_proto protoreflect.FileDescriptor

var file_sso_user_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x73, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x73, 0x73, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x43, 0x0a, 0x12, 0x45, 0x64, 0x69,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x73, 0x73, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x54,
	0x0a, 0x13, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x73,
	0x73, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x24, 0x0a, 0x12, 0x53, 0x68, 0x6f, 0x77, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a,
	0x13, 0x53, 0x68, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x32, 0xaf, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x5d, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x73,
	0x6f, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a,
	0x32, 0x10, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x62, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x2a, 0x12, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5d, 0x0a, 0x0b, 0x53, 0x68, 0x6f, 0x77, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x53, 0x68, 0x6f, 0x77,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x12, 0x13, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x12, 0x5a, 0x10, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x73, 0x6f, 0x3b, 0x73, 0x73, 0x6f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_sso_user_proto_rawDescOnce sync.Once
	file_sso_user_proto_rawDescData = file_sso_user_proto_rawDesc
)

func file_sso_user_proto_rawDescGZIP() []byte {
	file_sso_user_proto_rawDescOnce.Do(func() {
		file_sso_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_sso_user_proto_rawDescData)
	})
	return file_sso_user_proto_rawDescData
}

var file_sso_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_sso_user_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: sso.User
	(*EditProfileRequest)(nil),    // 1: sso.EditProfileRequest
	(*EditProfileResponse)(nil),   // 2: sso.EditProfileResponse
	(*DeleteAccountRequest)(nil),  // 3: sso.DeleteAccountRequest
	(*DeleteAccountResponse)(nil), // 4: sso.DeleteAccountResponse
	(*ShowProfileRequest)(nil),    // 5: sso.ShowProfileRequest
	(*ShowProfileResponse)(nil),   // 6: sso.ShowProfileResponse
}
var file_sso_user_proto_depIdxs = []int32{
	0, // 0: sso.EditProfileRequest.user:type_name -> sso.User
	0, // 1: sso.EditProfileResponse.updatedUser:type_name -> sso.User
	0, // 2: sso.ShowProfileResponse.user:type_name -> sso.User
	1, // 3: sso.UserProfile.EditProfile:input_type -> sso.EditProfileRequest
	3, // 4: sso.UserProfile.DeleteAccount:input_type -> sso.DeleteAccountRequest
	5, // 5: sso.UserProfile.ShowProfile:input_type -> sso.ShowProfileRequest
	2, // 6: sso.UserProfile.EditProfile:output_type -> sso.EditProfileResponse
	4, // 7: sso.UserProfile.DeleteAccount:output_type -> sso.DeleteAccountResponse
	6, // 8: sso.UserProfile.ShowProfile:output_type -> sso.ShowProfileResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sso_user_proto_init() }
func file_sso_user_proto_init() {
	if File_sso_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sso_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_sso_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditProfileRequest); i {
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
		file_sso_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditProfileResponse); i {
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
		file_sso_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccountRequest); i {
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
		file_sso_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccountResponse); i {
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
		file_sso_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowProfileRequest); i {
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
		file_sso_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowProfileResponse); i {
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
			RawDescriptor: file_sso_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sso_user_proto_goTypes,
		DependencyIndexes: file_sso_user_proto_depIdxs,
		MessageInfos:      file_sso_user_proto_msgTypes,
	}.Build()
	File_sso_user_proto = out.File
	file_sso_user_proto_rawDesc = nil
	file_sso_user_proto_goTypes = nil
	file_sso_user_proto_depIdxs = nil
}
