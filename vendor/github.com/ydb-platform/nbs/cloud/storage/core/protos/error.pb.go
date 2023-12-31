// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: cloud/storage/core/protos/error.proto

package protos

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

type EErrorFlag int32

const (
	EErrorFlag_EF_NONE EErrorFlag = 0
	// Be careful using EF_SILENT with retriable errors. This flag changes error
	// kind to EErrorKind::ErrorSilent making the error non-retriable.
	EErrorFlag_EF_SILENT EErrorFlag = 1
	// Currently set if a problem with nonreplicated disk/agent was detected.
	EErrorFlag_EF_HW_PROBLEMS_DETECTED EErrorFlag = 2
)

// Enum value maps for EErrorFlag.
var (
	EErrorFlag_name = map[int32]string{
		0: "EF_NONE",
		1: "EF_SILENT",
		2: "EF_HW_PROBLEMS_DETECTED",
	}
	EErrorFlag_value = map[string]int32{
		"EF_NONE":                 0,
		"EF_SILENT":               1,
		"EF_HW_PROBLEMS_DETECTED": 2,
	}
)

func (x EErrorFlag) Enum() *EErrorFlag {
	p := new(EErrorFlag)
	*p = x
	return p
}

func (x EErrorFlag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EErrorFlag) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_storage_core_protos_error_proto_enumTypes[0].Descriptor()
}

func (EErrorFlag) Type() protoreflect.EnumType {
	return &file_cloud_storage_core_protos_error_proto_enumTypes[0]
}

func (x EErrorFlag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EErrorFlag.Descriptor instead.
func (EErrorFlag) EnumDescriptor() ([]byte, []int) {
	return file_cloud_storage_core_protos_error_proto_rawDescGZIP(), []int{0}
}

type TError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Error code.
	Code uint32 `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	// Error message.
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	// Flags.
	Flags uint32 `protobuf:"varint,3,opt,name=Flags,proto3" json:"Flags,omitempty"`
}

func (x *TError) Reset() {
	*x = TError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_storage_core_protos_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TError) ProtoMessage() {}

func (x *TError) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_storage_core_protos_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TError.ProtoReflect.Descriptor instead.
func (*TError) Descriptor() ([]byte, []int) {
	return file_cloud_storage_core_protos_error_proto_rawDescGZIP(), []int{0}
}

func (x *TError) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *TError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TError) GetFlags() uint32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

var File_cloud_storage_core_protos_error_proto protoreflect.FileDescriptor

var file_cloud_storage_core_protos_error_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x06, 0x54, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x46,
	0x6c, 0x61, 0x67, 0x73, 0x2a, 0x45, 0x0a, 0x0a, 0x45, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x46, 0x6c,
	0x61, 0x67, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x46, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x45, 0x46, 0x5f, 0x53, 0x49, 0x4c, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x1b,
	0x0a, 0x17, 0x45, 0x46, 0x5f, 0x48, 0x57, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x53,
	0x5f, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x42, 0x2c, 0x5a, 0x2a, 0x61,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x72, 0x75, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cloud_storage_core_protos_error_proto_rawDescOnce sync.Once
	file_cloud_storage_core_protos_error_proto_rawDescData = file_cloud_storage_core_protos_error_proto_rawDesc
)

func file_cloud_storage_core_protos_error_proto_rawDescGZIP() []byte {
	file_cloud_storage_core_protos_error_proto_rawDescOnce.Do(func() {
		file_cloud_storage_core_protos_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_storage_core_protos_error_proto_rawDescData)
	})
	return file_cloud_storage_core_protos_error_proto_rawDescData
}

var file_cloud_storage_core_protos_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_storage_core_protos_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cloud_storage_core_protos_error_proto_goTypes = []interface{}{
	(EErrorFlag)(0), // 0: NCloud.NProto.EErrorFlag
	(*TError)(nil),  // 1: NCloud.NProto.TError
}
var file_cloud_storage_core_protos_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_storage_core_protos_error_proto_init() }
func file_cloud_storage_core_protos_error_proto_init() {
	if File_cloud_storage_core_protos_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_storage_core_protos_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TError); i {
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
			RawDescriptor: file_cloud_storage_core_protos_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_storage_core_protos_error_proto_goTypes,
		DependencyIndexes: file_cloud_storage_core_protos_error_proto_depIdxs,
		EnumInfos:         file_cloud_storage_core_protos_error_proto_enumTypes,
		MessageInfos:      file_cloud_storage_core_protos_error_proto_msgTypes,
	}.Build()
	File_cloud_storage_core_protos_error_proto = out.File
	file_cloud_storage_core_protos_error_proto_rawDesc = nil
	file_cloud_storage_core_protos_error_proto_goTypes = nil
	file_cloud_storage_core_protos_error_proto_depIdxs = nil
}
