// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: cloud/storage/core/protos/tablet.proto

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

type TTabletChannelHistoryEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromGeneration uint32 `protobuf:"varint,1,opt,name=FromGeneration,proto3" json:"FromGeneration,omitempty"`
	GroupId        uint32 `protobuf:"varint,2,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
}

func (x *TTabletChannelHistoryEntry) Reset() {
	*x = TTabletChannelHistoryEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_storage_core_protos_tablet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TTabletChannelHistoryEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TTabletChannelHistoryEntry) ProtoMessage() {}

func (x *TTabletChannelHistoryEntry) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_storage_core_protos_tablet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TTabletChannelHistoryEntry.ProtoReflect.Descriptor instead.
func (*TTabletChannelHistoryEntry) Descriptor() ([]byte, []int) {
	return file_cloud_storage_core_protos_tablet_proto_rawDescGZIP(), []int{0}
}

func (x *TTabletChannelHistoryEntry) GetFromGeneration() uint32 {
	if x != nil {
		return x.FromGeneration
	}
	return 0
}

func (x *TTabletChannelHistoryEntry) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type TTabletChannelInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoragePool string                        `protobuf:"bytes,1,opt,name=StoragePool,proto3" json:"StoragePool,omitempty"`
	History     []*TTabletChannelHistoryEntry `protobuf:"bytes,2,rep,name=History,proto3" json:"History,omitempty"`
}

func (x *TTabletChannelInfo) Reset() {
	*x = TTabletChannelInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_storage_core_protos_tablet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TTabletChannelInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TTabletChannelInfo) ProtoMessage() {}

func (x *TTabletChannelInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_storage_core_protos_tablet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TTabletChannelInfo.ProtoReflect.Descriptor instead.
func (*TTabletChannelInfo) Descriptor() ([]byte, []int) {
	return file_cloud_storage_core_protos_tablet_proto_rawDescGZIP(), []int{1}
}

func (x *TTabletChannelInfo) GetStoragePool() string {
	if x != nil {
		return x.StoragePool
	}
	return ""
}

func (x *TTabletChannelInfo) GetHistory() []*TTabletChannelHistoryEntry {
	if x != nil {
		return x.History
	}
	return nil
}

type TTabletStorageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TabletId uint64                `protobuf:"varint,1,opt,name=TabletId,proto3" json:"TabletId,omitempty"`
	Version  uint64                `protobuf:"varint,2,opt,name=Version,proto3" json:"Version,omitempty"`
	Channels []*TTabletChannelInfo `protobuf:"bytes,3,rep,name=Channels,proto3" json:"Channels,omitempty"`
}

func (x *TTabletStorageInfo) Reset() {
	*x = TTabletStorageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_storage_core_protos_tablet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TTabletStorageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TTabletStorageInfo) ProtoMessage() {}

func (x *TTabletStorageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_storage_core_protos_tablet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TTabletStorageInfo.ProtoReflect.Descriptor instead.
func (*TTabletStorageInfo) Descriptor() ([]byte, []int) {
	return file_cloud_storage_core_protos_tablet_proto_rawDescGZIP(), []int{2}
}

func (x *TTabletStorageInfo) GetTabletId() uint64 {
	if x != nil {
		return x.TabletId
	}
	return 0
}

func (x *TTabletStorageInfo) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *TTabletStorageInfo) GetChannels() []*TTabletChannelInfo {
	if x != nil {
		return x.Channels
	}
	return nil
}

var File_cloud_storage_core_protos_tablet_proto protoreflect.FileDescriptor

var file_cloud_storage_core_protos_tablet_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x1a, 0x54, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x46,
	0x72, 0x6f, 0x6d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x7b, 0x0a, 0x12, 0x54, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a,
	0x0b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12,
	0x43, 0x0a, 0x07, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x22, 0x89, 0x01, 0x0a, 0x12, 0x54, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x74,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x3d, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73,
	0x42, 0x2c, 0x5a, 0x2a, 0x61, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65, 0x61,
	0x6d, 0x2e, 0x72, 0x75, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_storage_core_protos_tablet_proto_rawDescOnce sync.Once
	file_cloud_storage_core_protos_tablet_proto_rawDescData = file_cloud_storage_core_protos_tablet_proto_rawDesc
)

func file_cloud_storage_core_protos_tablet_proto_rawDescGZIP() []byte {
	file_cloud_storage_core_protos_tablet_proto_rawDescOnce.Do(func() {
		file_cloud_storage_core_protos_tablet_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_storage_core_protos_tablet_proto_rawDescData)
	})
	return file_cloud_storage_core_protos_tablet_proto_rawDescData
}

var file_cloud_storage_core_protos_tablet_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cloud_storage_core_protos_tablet_proto_goTypes = []interface{}{
	(*TTabletChannelHistoryEntry)(nil), // 0: NCloud.NProto.TTabletChannelHistoryEntry
	(*TTabletChannelInfo)(nil),         // 1: NCloud.NProto.TTabletChannelInfo
	(*TTabletStorageInfo)(nil),         // 2: NCloud.NProto.TTabletStorageInfo
}
var file_cloud_storage_core_protos_tablet_proto_depIdxs = []int32{
	0, // 0: NCloud.NProto.TTabletChannelInfo.History:type_name -> NCloud.NProto.TTabletChannelHistoryEntry
	1, // 1: NCloud.NProto.TTabletStorageInfo.Channels:type_name -> NCloud.NProto.TTabletChannelInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cloud_storage_core_protos_tablet_proto_init() }
func file_cloud_storage_core_protos_tablet_proto_init() {
	if File_cloud_storage_core_protos_tablet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_storage_core_protos_tablet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TTabletChannelHistoryEntry); i {
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
		file_cloud_storage_core_protos_tablet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TTabletChannelInfo); i {
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
		file_cloud_storage_core_protos_tablet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TTabletStorageInfo); i {
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
			RawDescriptor: file_cloud_storage_core_protos_tablet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_storage_core_protos_tablet_proto_goTypes,
		DependencyIndexes: file_cloud_storage_core_protos_tablet_proto_depIdxs,
		MessageInfos:      file_cloud_storage_core_protos_tablet_proto_msgTypes,
	}.Build()
	File_cloud_storage_core_protos_tablet_proto = out.File
	file_cloud_storage_core_protos_tablet_proto_rawDesc = nil
	file_cloud_storage_core_protos_tablet_proto_goTypes = nil
	file_cloud_storage_core_protos_tablet_proto_depIdxs = nil
}
