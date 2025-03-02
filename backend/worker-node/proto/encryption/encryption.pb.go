// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/encryption.proto

package encryption

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EncryptionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Plaintext     string                 `protobuf:"bytes,1,opt,name=plaintext,proto3" json:"plaintext,omitempty"`
	Key           string                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EncryptionRequest) Reset() {
	*x = EncryptionRequest{}
	mi := &file_proto_encryption_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EncryptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptionRequest) ProtoMessage() {}

func (x *EncryptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_encryption_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptionRequest.ProtoReflect.Descriptor instead.
func (*EncryptionRequest) Descriptor() ([]byte, []int) {
	return file_proto_encryption_proto_rawDescGZIP(), []int{0}
}

func (x *EncryptionRequest) GetPlaintext() string {
	if x != nil {
		return x.Plaintext
	}
	return ""
}

func (x *EncryptionRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type EncryptionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ciphertext    string                 `protobuf:"bytes,1,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EncryptionResponse) Reset() {
	*x = EncryptionResponse{}
	mi := &file_proto_encryption_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EncryptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptionResponse) ProtoMessage() {}

func (x *EncryptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_encryption_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptionResponse.ProtoReflect.Descriptor instead.
func (*EncryptionResponse) Descriptor() ([]byte, []int) {
	return file_proto_encryption_proto_rawDescGZIP(), []int{1}
}

func (x *EncryptionResponse) GetCiphertext() string {
	if x != nil {
		return x.Ciphertext
	}
	return ""
}

var File_proto_encryption_proto protoreflect.FileDescriptor

var file_proto_encryption_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x43, 0x0a, 0x11, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61,
	0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c,
	0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x34, 0x0a, 0x12, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x32,
	0xd5, 0x02, 0x0a, 0x11, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x41, 0x45, 0x53, 0x32, 0x35, 0x36, 0x12, 0x1d, 0x2e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x41, 0x45, 0x53, 0x32, 0x35, 0x36, 0x12, 0x1d, 0x2e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x52, 0x53, 0x41, 0x32, 0x30, 0x34, 0x38, 0x12, 0x1d, 0x2e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x52, 0x53, 0x41, 0x32, 0x30, 0x34, 0x38, 0x12, 0x1d, 0x2e, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x2d, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_encryption_proto_rawDescOnce sync.Once
	file_proto_encryption_proto_rawDescData []byte
)

func file_proto_encryption_proto_rawDescGZIP() []byte {
	file_proto_encryption_proto_rawDescOnce.Do(func() {
		file_proto_encryption_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_encryption_proto_rawDesc), len(file_proto_encryption_proto_rawDesc)))
	})
	return file_proto_encryption_proto_rawDescData
}

var file_proto_encryption_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_encryption_proto_goTypes = []any{
	(*EncryptionRequest)(nil),  // 0: encryption.EncryptionRequest
	(*EncryptionResponse)(nil), // 1: encryption.EncryptionResponse
}
var file_proto_encryption_proto_depIdxs = []int32{
	0, // 0: encryption.EncryptionService.EncryptAES256:input_type -> encryption.EncryptionRequest
	0, // 1: encryption.EncryptionService.DecryptAES256:input_type -> encryption.EncryptionRequest
	0, // 2: encryption.EncryptionService.EncryptRSA2048:input_type -> encryption.EncryptionRequest
	0, // 3: encryption.EncryptionService.DecryptRSA2048:input_type -> encryption.EncryptionRequest
	1, // 4: encryption.EncryptionService.EncryptAES256:output_type -> encryption.EncryptionResponse
	1, // 5: encryption.EncryptionService.DecryptAES256:output_type -> encryption.EncryptionResponse
	1, // 6: encryption.EncryptionService.EncryptRSA2048:output_type -> encryption.EncryptionResponse
	1, // 7: encryption.EncryptionService.DecryptRSA2048:output_type -> encryption.EncryptionResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_encryption_proto_init() }
func file_proto_encryption_proto_init() {
	if File_proto_encryption_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_encryption_proto_rawDesc), len(file_proto_encryption_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_encryption_proto_goTypes,
		DependencyIndexes: file_proto_encryption_proto_depIdxs,
		MessageInfos:      file_proto_encryption_proto_msgTypes,
	}.Build()
	File_proto_encryption_proto = out.File
	file_proto_encryption_proto_goTypes = nil
	file_proto_encryption_proto_depIdxs = nil
}
