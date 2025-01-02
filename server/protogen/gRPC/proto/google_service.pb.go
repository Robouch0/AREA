// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.19.6
// source: gRPC/proto/google_service.proto

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

type EmailRequestMe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	// string from = 2; // Will be the user email registered
	Subject     string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	BodyMessage string `protobuf:"bytes,3,opt,name=body_message,json=bodyMessage,proto3" json:"body_message,omitempty"`
}

func (x *EmailRequestMe) Reset() {
	*x = EmailRequestMe{}
	mi := &file_gRPC_proto_google_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailRequestMe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailRequestMe) ProtoMessage() {}

func (x *EmailRequestMe) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_google_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailRequestMe.ProtoReflect.Descriptor instead.
func (*EmailRequestMe) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_google_service_proto_rawDescGZIP(), []int{0}
}

func (x *EmailRequestMe) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *EmailRequestMe) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *EmailRequestMe) GetBodyMessage() string {
	if x != nil {
		return x.BodyMessage
	}
	return ""
}

type DeleteEmailRequestMe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *DeleteEmailRequestMe) Reset() {
	*x = DeleteEmailRequestMe{}
	mi := &file_gRPC_proto_google_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteEmailRequestMe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEmailRequestMe) ProtoMessage() {}

func (x *DeleteEmailRequestMe) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_google_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEmailRequestMe.ProtoReflect.Descriptor instead.
func (*DeleteEmailRequestMe) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_google_service_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteEmailRequestMe) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

type TrashEmailRequestMe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *TrashEmailRequestMe) Reset() {
	*x = TrashEmailRequestMe{}
	mi := &file_gRPC_proto_google_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrashEmailRequestMe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrashEmailRequestMe) ProtoMessage() {}

func (x *TrashEmailRequestMe) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_google_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrashEmailRequestMe.ProtoReflect.Descriptor instead.
func (*TrashEmailRequestMe) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_google_service_proto_rawDescGZIP(), []int{2}
}

func (x *TrashEmailRequestMe) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

type EmailTriggerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionId uint32 `protobuf:"varint,1,opt,name=action_id,json=actionId,proto3" json:"action_id,omitempty"`
}

func (x *EmailTriggerReq) Reset() {
	*x = EmailTriggerReq{}
	mi := &file_gRPC_proto_google_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailTriggerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailTriggerReq) ProtoMessage() {}

func (x *EmailTriggerReq) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_google_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailTriggerReq.ProtoReflect.Descriptor instead.
func (*EmailTriggerReq) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_google_service_proto_rawDescGZIP(), []int{3}
}

func (x *EmailTriggerReq) GetActionId() uint32 {
	if x != nil {
		return x.ActionId
	}
	return 0
}

type GmailTriggerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionId uint32 `protobuf:"varint,1,opt,name=action_id,json=actionId,proto3" json:"action_id,omitempty"`
	Payload  []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *GmailTriggerReq) Reset() {
	*x = GmailTriggerReq{}
	mi := &file_gRPC_proto_google_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GmailTriggerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GmailTriggerReq) ProtoMessage() {}

func (x *GmailTriggerReq) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_google_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GmailTriggerReq.ProtoReflect.Descriptor instead.
func (*GmailTriggerReq) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_google_service_proto_rawDescGZIP(), []int{4}
}

func (x *GmailTriggerReq) GetActionId() uint32 {
	if x != nil {
		return x.ActionId
	}
	return 0
}

func (x *GmailTriggerReq) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type CreateLabelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MessageListVisibility string `protobuf:"bytes,2,opt,name=messageListVisibility,proto3" json:"messageListVisibility,omitempty"`
	LabelListVisibility   string `protobuf:"bytes,3,opt,name=LabelListVisibility,proto3" json:"LabelListVisibility,omitempty"`
	Type                  string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *CreateLabelReq) Reset() {
	*x = CreateLabelReq{}
	mi := &file_gRPC_proto_google_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLabelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLabelReq) ProtoMessage() {}

func (x *CreateLabelReq) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_google_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLabelReq.ProtoReflect.Descriptor instead.
func (*CreateLabelReq) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_google_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateLabelReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateLabelReq) GetMessageListVisibility() string {
	if x != nil {
		return x.MessageListVisibility
	}
	return ""
}

func (x *CreateLabelReq) GetLabelListVisibility() string {
	if x != nil {
		return x.LabelListVisibility
	}
	return ""
}

func (x *CreateLabelReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type UpdateLabelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldName               string `protobuf:"bytes,1,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	NewName               string `protobuf:"bytes,2,opt,name=new_name,json=newName,proto3" json:"new_name,omitempty"`
	MessageListVisibility string `protobuf:"bytes,3,opt,name=messageListVisibility,proto3" json:"messageListVisibility,omitempty"`
	LabelListVisibility   string `protobuf:"bytes,4,opt,name=LabelListVisibility,proto3" json:"LabelListVisibility,omitempty"`
	Type                  string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *UpdateLabelReq) Reset() {
	*x = UpdateLabelReq{}
	mi := &file_gRPC_proto_google_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateLabelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLabelReq) ProtoMessage() {}

func (x *UpdateLabelReq) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_google_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLabelReq.ProtoReflect.Descriptor instead.
func (*UpdateLabelReq) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_google_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateLabelReq) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *UpdateLabelReq) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

func (x *UpdateLabelReq) GetMessageListVisibility() string {
	if x != nil {
		return x.MessageListVisibility
	}
	return ""
}

func (x *UpdateLabelReq) GetLabelListVisibility() string {
	if x != nil {
		return x.LabelListVisibility
	}
	return ""
}

func (x *UpdateLabelReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type DeleteLabelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteLabelReq) Reset() {
	*x = DeleteLabelReq{}
	mi := &file_gRPC_proto_google_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteLabelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLabelReq) ProtoMessage() {}

func (x *DeleteLabelReq) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_google_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLabelReq.ProtoReflect.Descriptor instead.
func (*DeleteLabelReq) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_google_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteLabelReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_gRPC_proto_google_service_proto protoreflect.FileDescriptor

var file_gRPC_proto_google_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x22, 0x5d, 0x0a, 0x0e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6f, 0x64,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2f, 0x0a, 0x13, 0x54, 0x72,
	0x61, 0x73, 0x68, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2e, 0x0a, 0x0f, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x0f, 0x47,
	0x6d, 0x61, 0x69, 0x6c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xa0, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x15,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x73, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x30, 0x0a, 0x13, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x56,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xc2, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x34, 0x0a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x73,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x30, 0x0a, 0x13, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x56,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x24, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x32, 0x87, 0x05, 0x0a, 0x0d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x4d, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x54,
	0x72, 0x61, 0x73, 0x68, 0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x54, 0x72,
	0x61, 0x73, 0x68, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d,
	0x65, 0x1a, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x73, 0x68,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x0d, 0x4d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x72, 0x61, 0x73,
	0x68, 0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x73, 0x68,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x1a, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x73, 0x68, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x22, 0x00,
	0x12, 0x45, 0x0a, 0x0f, 0x57, 0x61, 0x74, 0x63, 0x68, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0e, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x4d, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x47, 0x6d, 0x61, 0x69,
	0x6c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x22, 0x00, 0x42, 0x21, 0x5a,
	0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x65, 0x61,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gRPC_proto_google_service_proto_rawDescOnce sync.Once
	file_gRPC_proto_google_service_proto_rawDescData = file_gRPC_proto_google_service_proto_rawDesc
)

func file_gRPC_proto_google_service_proto_rawDescGZIP() []byte {
	file_gRPC_proto_google_service_proto_rawDescOnce.Do(func() {
		file_gRPC_proto_google_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_gRPC_proto_google_service_proto_rawDescData)
	})
	return file_gRPC_proto_google_service_proto_rawDescData
}

var file_gRPC_proto_google_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_gRPC_proto_google_service_proto_goTypes = []any{
	(*EmailRequestMe)(nil),       // 0: google.EmailRequestMe
	(*DeleteEmailRequestMe)(nil), // 1: google.DeleteEmailRequestMe
	(*TrashEmailRequestMe)(nil),  // 2: google.TrashEmailRequestMe
	(*EmailTriggerReq)(nil),      // 3: google.EmailTriggerReq
	(*GmailTriggerReq)(nil),      // 4: google.GmailTriggerReq
	(*CreateLabelReq)(nil),       // 5: google.CreateLabelReq
	(*UpdateLabelReq)(nil),       // 6: google.UpdateLabelReq
	(*DeleteLabelReq)(nil),       // 7: google.DeleteLabelReq
}
var file_gRPC_proto_google_service_proto_depIdxs = []int32{
	0, // 0: google.GoogleService.SendEmailMe:input_type -> google.EmailRequestMe
	1, // 1: google.GoogleService.DeleteEmailMe:input_type -> google.DeleteEmailRequestMe
	2, // 2: google.GoogleService.MoveToTrash:input_type -> google.TrashEmailRequestMe
	2, // 3: google.GoogleService.MoveFromTrash:input_type -> google.TrashEmailRequestMe
	5, // 4: google.GoogleService.CreateLabel:input_type -> google.CreateLabelReq
	7, // 5: google.GoogleService.DeleteLabel:input_type -> google.DeleteLabelReq
	6, // 6: google.GoogleService.UpdateLabel:input_type -> google.UpdateLabelReq
	3, // 7: google.GoogleService.WatchGmailEmail:input_type -> google.EmailTriggerReq
	4, // 8: google.GoogleService.WatchMeTrigger:input_type -> google.GmailTriggerReq
	0, // 9: google.GoogleService.SendEmailMe:output_type -> google.EmailRequestMe
	1, // 10: google.GoogleService.DeleteEmailMe:output_type -> google.DeleteEmailRequestMe
	2, // 11: google.GoogleService.MoveToTrash:output_type -> google.TrashEmailRequestMe
	2, // 12: google.GoogleService.MoveFromTrash:output_type -> google.TrashEmailRequestMe
	5, // 13: google.GoogleService.CreateLabel:output_type -> google.CreateLabelReq
	7, // 14: google.GoogleService.DeleteLabel:output_type -> google.DeleteLabelReq
	6, // 15: google.GoogleService.UpdateLabel:output_type -> google.UpdateLabelReq
	3, // 16: google.GoogleService.WatchGmailEmail:output_type -> google.EmailTriggerReq
	4, // 17: google.GoogleService.WatchMeTrigger:output_type -> google.GmailTriggerReq
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gRPC_proto_google_service_proto_init() }
func file_gRPC_proto_google_service_proto_init() {
	if File_gRPC_proto_google_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gRPC_proto_google_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gRPC_proto_google_service_proto_goTypes,
		DependencyIndexes: file_gRPC_proto_google_service_proto_depIdxs,
		MessageInfos:      file_gRPC_proto_google_service_proto_msgTypes,
	}.Build()
	File_gRPC_proto_google_service_proto = out.File
	file_gRPC_proto_google_service_proto_rawDesc = nil
	file_gRPC_proto_google_service_proto_goTypes = nil
	file_gRPC_proto_google_service_proto_depIdxs = nil
}
