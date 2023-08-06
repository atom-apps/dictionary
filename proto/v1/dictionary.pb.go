// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/v1/dictionary.proto

package v1

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

type GetDictionariesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   uint64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	TenantID uint64 `protobuf:"varint,2,opt,name=TenantID,proto3" json:"TenantID,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug     string `protobuf:"bytes,4,opt,name=Slug,proto3" json:"Slug,omitempty"`
}

func (x *GetDictionariesRequest) Reset() {
	*x = GetDictionariesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_dictionary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDictionariesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDictionariesRequest) ProtoMessage() {}

func (x *GetDictionariesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_dictionary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDictionariesRequest.ProtoReflect.Descriptor instead.
func (*GetDictionariesRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_dictionary_proto_rawDescGZIP(), []int{0}
}

func (x *GetDictionariesRequest) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *GetDictionariesRequest) GetTenantID() uint64 {
	if x != nil {
		return x.TenantID
	}
	return 0
}

func (x *GetDictionariesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetDictionariesRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type GetDictionariesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*DictionaryItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetDictionariesResponse) Reset() {
	*x = GetDictionariesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_dictionary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDictionariesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDictionariesResponse) ProtoMessage() {}

func (x *GetDictionariesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_dictionary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDictionariesResponse.ProtoReflect.Descriptor instead.
func (*GetDictionariesResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_dictionary_proto_rawDescGZIP(), []int{1}
}

func (x *GetDictionariesResponse) GetItems() []*DictionaryItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type DictionaryItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt   string `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt   string `protobuf:"bytes,3,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	UserID      uint64 `protobuf:"varint,4,opt,name=UserID,proto3" json:"UserID,omitempty"`
	TenantID    uint64 `protobuf:"varint,5,opt,name=TenantID,proto3" json:"TenantID,omitempty"`
	Name        string `protobuf:"bytes,6,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug        string `protobuf:"bytes,7,opt,name=Slug,proto3" json:"Slug,omitempty"`
	Description string `protobuf:"bytes,8,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *DictionaryItem) Reset() {
	*x = DictionaryItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_dictionary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictionaryItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictionaryItem) ProtoMessage() {}

func (x *DictionaryItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_dictionary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictionaryItem.ProtoReflect.Descriptor instead.
func (*DictionaryItem) Descriptor() ([]byte, []int) {
	return file_proto_v1_dictionary_proto_rawDescGZIP(), []int{2}
}

func (x *DictionaryItem) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DictionaryItem) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *DictionaryItem) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *DictionaryItem) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *DictionaryItem) GetTenantID() uint64 {
	if x != nil {
		return x.TenantID
	}
	return 0
}

func (x *DictionaryItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DictionaryItem) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *DictionaryItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_proto_v1_dictionary_proto protoreflect.FileDescriptor

var file_proto_v1_dictionary_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x74, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6c, 0x75, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x6c, 0x75, 0x67, 0x22, 0x49, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xda, 0x01, 0x0a, 0x0e, 0x44, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x53, 0x6c, 0x75, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x6c, 0x75,
	0x67, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x32, 0x6d, 0x0a, 0x11, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x74, 0x6f, 0x6d, 0x2d, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_dictionary_proto_rawDescOnce sync.Once
	file_proto_v1_dictionary_proto_rawDescData = file_proto_v1_dictionary_proto_rawDesc
)

func file_proto_v1_dictionary_proto_rawDescGZIP() []byte {
	file_proto_v1_dictionary_proto_rawDescOnce.Do(func() {
		file_proto_v1_dictionary_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_dictionary_proto_rawDescData)
	})
	return file_proto_v1_dictionary_proto_rawDescData
}

var file_proto_v1_dictionary_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_v1_dictionary_proto_goTypes = []interface{}{
	(*GetDictionariesRequest)(nil),  // 0: proto.v1.GetDictionariesRequest
	(*GetDictionariesResponse)(nil), // 1: proto.v1.GetDictionariesResponse
	(*DictionaryItem)(nil),          // 2: proto.v1.DictionaryItem
}
var file_proto_v1_dictionary_proto_depIdxs = []int32{
	2, // 0: proto.v1.GetDictionariesResponse.items:type_name -> proto.v1.DictionaryItem
	0, // 1: proto.v1.DictionaryService.GetDictionaries:input_type -> proto.v1.GetDictionariesRequest
	1, // 2: proto.v1.DictionaryService.GetDictionaries:output_type -> proto.v1.GetDictionariesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_v1_dictionary_proto_init() }
func file_proto_v1_dictionary_proto_init() {
	if File_proto_v1_dictionary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_dictionary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDictionariesRequest); i {
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
		file_proto_v1_dictionary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDictionariesResponse); i {
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
		file_proto_v1_dictionary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictionaryItem); i {
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
			RawDescriptor: file_proto_v1_dictionary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_dictionary_proto_goTypes,
		DependencyIndexes: file_proto_v1_dictionary_proto_depIdxs,
		MessageInfos:      file_proto_v1_dictionary_proto_msgTypes,
	}.Build()
	File_proto_v1_dictionary_proto = out.File
	file_proto_v1_dictionary_proto_rawDesc = nil
	file_proto_v1_dictionary_proto_goTypes = nil
	file_proto_v1_dictionary_proto_depIdxs = nil
}
