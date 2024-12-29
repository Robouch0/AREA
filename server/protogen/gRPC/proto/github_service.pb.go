// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.21.12
// source: gRPC/proto/github_service.proto

package service

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

type GitWebHookInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Id of the action that will be sent while creating the webhook
	ActionId int32 `protobuf:"varint,1,opt,name=action_id,json=actionId,proto3" json:"action_id,omitempty"`
	// Name of the owner of the repository
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// Name of the repository
	Repo          string `protobuf:"bytes,3,opt,name=repo,proto3" json:"repo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GitWebHookInfo) Reset() {
	*x = GitWebHookInfo{}
	mi := &file_gRPC_proto_github_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GitWebHookInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitWebHookInfo) ProtoMessage() {}

func (x *GitWebHookInfo) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_github_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitWebHookInfo.ProtoReflect.Descriptor instead.
func (*GitWebHookInfo) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_github_service_proto_rawDescGZIP(), []int{0}
}

func (x *GitWebHookInfo) GetActionId() int32 {
	if x != nil {
		return x.ActionId
	}
	return 0
}

func (x *GitWebHookInfo) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GitWebHookInfo) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

type UpdateRepoInfos struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Owner         string                 `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"` // Path parameter
	Repo          string                 `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`   // Path parameter
	Name          string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"` // There is more to add if needed
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRepoInfos) Reset() {
	*x = UpdateRepoInfos{}
	mi := &file_gRPC_proto_github_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRepoInfos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRepoInfos) ProtoMessage() {}

func (x *UpdateRepoInfos) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_github_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRepoInfos.ProtoReflect.Descriptor instead.
func (*UpdateRepoInfos) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_github_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRepoInfos) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *UpdateRepoInfos) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *UpdateRepoInfos) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRepoInfos) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UpdateRepoFile struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Owner         string                 `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"` // Path parameter
	Repo          string                 `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`   // Path parameter
	Path          string                 `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`   // Path parameter
	Message       string                 `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Content       string                 `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"` // Base64
	Sha           string                 `protobuf:"bytes,6,opt,name=sha,proto3" json:"sha,omitempty"`         // There is more to add if needed
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRepoFile) Reset() {
	*x = UpdateRepoFile{}
	mi := &file_gRPC_proto_github_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRepoFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRepoFile) ProtoMessage() {}

func (x *UpdateRepoFile) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_github_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRepoFile.ProtoReflect.Descriptor instead.
func (*UpdateRepoFile) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_github_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRepoFile) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *UpdateRepoFile) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *UpdateRepoFile) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UpdateRepoFile) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpdateRepoFile) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdateRepoFile) GetSha() string {
	if x != nil {
		return x.Sha
	}
	return ""
}

type DeleteRepoFile struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Owner         string                 `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Repo          string                 `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
	Path          string                 `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Message       string                 `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Sha           string                 `protobuf:"bytes,5,opt,name=sha,proto3" json:"sha,omitempty"` // There is more to add if needed
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRepoFile) Reset() {
	*x = DeleteRepoFile{}
	mi := &file_gRPC_proto_github_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRepoFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRepoFile) ProtoMessage() {}

func (x *DeleteRepoFile) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_github_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRepoFile.ProtoReflect.Descriptor instead.
func (*DeleteRepoFile) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_github_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRepoFile) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *DeleteRepoFile) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *DeleteRepoFile) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *DeleteRepoFile) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeleteRepoFile) GetSha() string {
	if x != nil {
		return x.Sha
	}
	return ""
}

type GithubWebHookTriggerReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ActionId      uint32                 `protobuf:"varint,1,opt,name=action_id,json=actionId,proto3" json:"action_id,omitempty"`
	Payload       []byte                 `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GithubWebHookTriggerReq) Reset() {
	*x = GithubWebHookTriggerReq{}
	mi := &file_gRPC_proto_github_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GithubWebHookTriggerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GithubWebHookTriggerReq) ProtoMessage() {}

func (x *GithubWebHookTriggerReq) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_github_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GithubWebHookTriggerReq.ProtoReflect.Descriptor instead.
func (*GithubWebHookTriggerReq) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_github_service_proto_rawDescGZIP(), []int{4}
}

func (x *GithubWebHookTriggerReq) GetActionId() uint32 {
	if x != nil {
		return x.ActionId
	}
	return 0
}

func (x *GithubWebHookTriggerReq) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_gRPC_proto_github_service_proto protoreflect.FileDescriptor

var file_gRPC_proto_github_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x22, 0x57, 0x0a, 0x0e, 0x47, 0x69, 0x74,
	0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65,
	0x70, 0x6f, 0x22, 0x71, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x65, 0x70, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x94, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65,
	0x70, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x68,
	0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x68, 0x61, 0x22, 0x7a, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x68, 0x61, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x68, 0x61, 0x22, 0x50, 0x0a, 0x17, 0x47, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0xf4, 0x02, 0x0a, 0x0d, 0x47,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x17, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x1a, 0x17, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x46, 0x69,
	0x6c, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x46, 0x69,
	0x6c, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75,
	0x73, 0x68, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x47, 0x69, 0x74, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x47, 0x69, 0x74, 0x57, 0x65,
	0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0e, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x1f, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x57, 0x65, 0x62,
	0x48, 0x6f, 0x6f, 0x6b, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1f,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x57, 0x65,
	0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x22,
	0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x72, 0x65, 0x61, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gRPC_proto_github_service_proto_rawDescOnce sync.Once
	file_gRPC_proto_github_service_proto_rawDescData = file_gRPC_proto_github_service_proto_rawDesc
)

func file_gRPC_proto_github_service_proto_rawDescGZIP() []byte {
	file_gRPC_proto_github_service_proto_rawDescOnce.Do(func() {
		file_gRPC_proto_github_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_gRPC_proto_github_service_proto_rawDescData)
	})
	return file_gRPC_proto_github_service_proto_rawDescData
}

var file_gRPC_proto_github_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gRPC_proto_github_service_proto_goTypes = []any{
	(*GitWebHookInfo)(nil),          // 0: github.GitWebHookInfo
	(*UpdateRepoInfos)(nil),         // 1: github.UpdateRepoInfos
	(*UpdateRepoFile)(nil),          // 2: github.UpdateRepoFile
	(*DeleteRepoFile)(nil),          // 3: github.DeleteRepoFile
	(*GithubWebHookTriggerReq)(nil), // 4: github.GithubWebHookTriggerReq
}
var file_gRPC_proto_github_service_proto_depIdxs = []int32{
	1, // 0: github.GithubService.UpdateRepository:input_type -> github.UpdateRepoInfos
	2, // 1: github.GithubService.UpdateFile:input_type -> github.UpdateRepoFile
	3, // 2: github.GithubService.DeleteFile:input_type -> github.DeleteRepoFile
	0, // 3: github.GithubService.CreatePushWebhook:input_type -> github.GitWebHookInfo
	4, // 4: github.GithubService.TriggerWebHook:input_type -> github.GithubWebHookTriggerReq
	1, // 5: github.GithubService.UpdateRepository:output_type -> github.UpdateRepoInfos
	2, // 6: github.GithubService.UpdateFile:output_type -> github.UpdateRepoFile
	3, // 7: github.GithubService.DeleteFile:output_type -> github.DeleteRepoFile
	0, // 8: github.GithubService.CreatePushWebhook:output_type -> github.GitWebHookInfo
	4, // 9: github.GithubService.TriggerWebHook:output_type -> github.GithubWebHookTriggerReq
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gRPC_proto_github_service_proto_init() }
func file_gRPC_proto_github_service_proto_init() {
	if File_gRPC_proto_github_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gRPC_proto_github_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gRPC_proto_github_service_proto_goTypes,
		DependencyIndexes: file_gRPC_proto_github_service_proto_depIdxs,
		MessageInfos:      file_gRPC_proto_github_service_proto_msgTypes,
	}.Build()
	File_gRPC_proto_github_service_proto = out.File
	file_gRPC_proto_github_service_proto_rawDesc = nil
	file_gRPC_proto_github_service_proto_goTypes = nil
	file_gRPC_proto_github_service_proto_depIdxs = nil
}
