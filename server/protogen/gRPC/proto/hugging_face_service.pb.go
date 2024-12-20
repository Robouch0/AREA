// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
<<<<<<< Updated upstream
// 	protoc-gen-go v1.36.0
=======
// 	protoc-gen-go v1.32.0
>>>>>>> Stashed changes
// 	protoc        v3.21.12
// source: gRPC/proto/hugging_face_service.proto

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

type TextGenerationReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Model         string                 `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	Inputs        string                 `protobuf:"bytes,2,opt,name=inputs,proto3" json:"inputs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TextGenerationReq) Reset() {
	*x = TextGenerationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_hugging_face_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextGenerationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextGenerationReq) ProtoMessage() {}

func (x *TextGenerationReq) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_hugging_face_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextGenerationReq.ProtoReflect.Descriptor instead.
func (*TextGenerationReq) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_hugging_face_service_proto_rawDescGZIP(), []int{0}
}

func (x *TextGenerationReq) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *TextGenerationReq) GetInputs() string {
	if x != nil {
		return x.Inputs
	}
	return ""
}

type TextGenerationRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GeneratedText string                 `protobuf:"bytes,1,opt,name=generated_text,json=generatedText,proto3" json:"generated_text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TextGenerationRes) Reset() {
	*x = TextGenerationRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_hugging_face_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextGenerationRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextGenerationRes) ProtoMessage() {}

func (x *TextGenerationRes) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_hugging_face_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextGenerationRes.ProtoReflect.Descriptor instead.
func (*TextGenerationRes) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_hugging_face_service_proto_rawDescGZIP(), []int{1}
}

func (x *TextGenerationRes) GetGeneratedText() string {
	if x != nil {
		return x.GeneratedText
	}
	return ""
}

var File_gRPC_proto_hugging_face_service_proto protoreflect.FileDescriptor

var file_gRPC_proto_hugging_face_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x75, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x75, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x66, 0x61, 0x63, 0x65, 0x22, 0x41, 0x0a, 0x11, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x22, 0x3a, 0x0a, 0x11, 0x54, 0x65, 0x78, 0x74, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x54,
	0x65, 0x78, 0x74, 0x32, 0x6e, 0x0a, 0x12, 0x48, 0x75, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x46, 0x61,
	0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x14, 0x4c, 0x61, 0x75,
	0x6e, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x2e, 0x68, 0x75, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x66, 0x61, 0x63, 0x65, 0x2e,
	0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x1a, 0x1e, 0x2e, 0x68, 0x75, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x66, 0x61, 0x63, 0x65, 0x2e,
	0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x72, 0x65, 0x61, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gRPC_proto_hugging_face_service_proto_rawDescOnce sync.Once
	file_gRPC_proto_hugging_face_service_proto_rawDescData = file_gRPC_proto_hugging_face_service_proto_rawDesc
)

func file_gRPC_proto_hugging_face_service_proto_rawDescGZIP() []byte {
	file_gRPC_proto_hugging_face_service_proto_rawDescOnce.Do(func() {
		file_gRPC_proto_hugging_face_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_gRPC_proto_hugging_face_service_proto_rawDescData)
	})
	return file_gRPC_proto_hugging_face_service_proto_rawDescData
}

var file_gRPC_proto_hugging_face_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gRPC_proto_hugging_face_service_proto_goTypes = []interface{}{
	(*TextGenerationReq)(nil), // 0: huggingface.TextGenerationReq
	(*TextGenerationRes)(nil), // 1: huggingface.TextGenerationRes
}
var file_gRPC_proto_hugging_face_service_proto_depIdxs = []int32{
	0, // 0: huggingface.HuggingFaceService.LaunchTextGeneration:input_type -> huggingface.TextGenerationReq
	1, // 1: huggingface.HuggingFaceService.LaunchTextGeneration:output_type -> huggingface.TextGenerationRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gRPC_proto_hugging_face_service_proto_init() }
func file_gRPC_proto_hugging_face_service_proto_init() {
	if File_gRPC_proto_hugging_face_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gRPC_proto_hugging_face_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextGenerationReq); i {
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
		file_gRPC_proto_hugging_face_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextGenerationRes); i {
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
			RawDescriptor: file_gRPC_proto_hugging_face_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gRPC_proto_hugging_face_service_proto_goTypes,
		DependencyIndexes: file_gRPC_proto_hugging_face_service_proto_depIdxs,
		MessageInfos:      file_gRPC_proto_hugging_face_service_proto_msgTypes,
	}.Build()
	File_gRPC_proto_hugging_face_service_proto = out.File
	file_gRPC_proto_hugging_face_service_proto_rawDesc = nil
	file_gRPC_proto_hugging_face_service_proto_goTypes = nil
	file_gRPC_proto_hugging_face_service_proto_depIdxs = nil
}
