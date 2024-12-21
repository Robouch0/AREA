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
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x50, 0x0a, 0x0d, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x53, 0x65, 0x6e,
	0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x65, 0x61, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_gRPC_proto_google_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gRPC_proto_google_service_proto_goTypes = []any{
	(*EmailRequestMe)(nil), // 0: google.EmailRequestMe
}
var file_gRPC_proto_google_service_proto_depIdxs = []int32{
	0, // 0: google.GoogleService.SendEmailMe:input_type -> google.EmailRequestMe
	0, // 1: google.GoogleService.SendEmailMe:output_type -> google.EmailRequestMe
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
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
			NumMessages:   1,
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
