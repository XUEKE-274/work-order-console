// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: tempalte.proto

package tRpc

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

type Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok string `protobuf:"bytes,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *Success) Reset() {
	*x = Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tempalte_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_tempalte_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_tempalte_proto_rawDescGZIP(), []int{0}
}

func (x *Success) GetOk() string {
	if x != nil {
		return x.Ok
	}
	return ""
}

type TemplateRpcAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TemplateRpcAdd) Reset() {
	*x = TemplateRpcAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tempalte_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRpcAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRpcAdd) ProtoMessage() {}

func (x *TemplateRpcAdd) ProtoReflect() protoreflect.Message {
	mi := &file_tempalte_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRpcAdd.ProtoReflect.Descriptor instead.
func (*TemplateRpcAdd) Descriptor() ([]byte, []int) {
	return file_tempalte_proto_rawDescGZIP(), []int{1}
}

func (x *TemplateRpcAdd) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PageList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PageList) Reset() {
	*x = PageList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tempalte_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageList) ProtoMessage() {}

func (x *PageList) ProtoReflect() protoreflect.Message {
	mi := &file_tempalte_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageList.ProtoReflect.Descriptor instead.
func (*PageList) Descriptor() ([]byte, []int) {
	return file_tempalte_proto_rawDescGZIP(), []int{2}
}

func (x *PageList) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *PageList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Templates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Templates []*TemplateVo `protobuf:"bytes,1,rep,name=templates,proto3" json:"templates,omitempty"`
}

func (x *Templates) Reset() {
	*x = Templates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tempalte_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Templates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Templates) ProtoMessage() {}

func (x *Templates) ProtoReflect() protoreflect.Message {
	mi := &file_tempalte_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Templates.ProtoReflect.Descriptor instead.
func (*Templates) Descriptor() ([]byte, []int) {
	return file_tempalte_proto_rawDescGZIP(), []int{3}
}

func (x *Templates) GetTemplates() []*TemplateVo {
	if x != nil {
		return x.Templates
	}
	return nil
}

type TemplateVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateBy   string `protobuf:"bytes,2,opt,name=createBy,proto3" json:"createBy,omitempty"`
	ModifyBy   string `protobuf:"bytes,3,opt,name=modifyBy,proto3" json:"modifyBy,omitempty"`
	CreateTime string `protobuf:"bytes,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	ModifyTime string `protobuf:"bytes,5,opt,name=modifyTime,proto3" json:"modifyTime,omitempty"`
	State      string `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	Name       string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TemplateVo) Reset() {
	*x = TemplateVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tempalte_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateVo) ProtoMessage() {}

func (x *TemplateVo) ProtoReflect() protoreflect.Message {
	mi := &file_tempalte_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateVo.ProtoReflect.Descriptor instead.
func (*TemplateVo) Descriptor() ([]byte, []int) {
	return file_tempalte_proto_rawDescGZIP(), []int{4}
}

func (x *TemplateVo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TemplateVo) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *TemplateVo) GetModifyBy() string {
	if x != nil {
		return x.ModifyBy
	}
	return ""
}

func (x *TemplateVo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *TemplateVo) GetModifyTime() string {
	if x != nil {
		return x.ModifyTime
	}
	return ""
}

func (x *TemplateVo) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *TemplateVo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_tempalte_proto protoreflect.FileDescriptor

var file_tempalte_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x61, 0x6c, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x19, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x24, 0x0a, 0x0e, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x70, 0x63, 0x41, 0x64, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x34, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x09, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x56, 0x6f, 0x52, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x22,
	0xbe, 0x01, 0x0a, 0x0a, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x42, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x32, 0x61, 0x0a, 0x12, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x70, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x70, 0x63, 0x41, 0x64, 0x64, 0x1a, 0x08, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x00, 0x12, 0x1f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x09, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x74, 0x52, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tempalte_proto_rawDescOnce sync.Once
	file_tempalte_proto_rawDescData = file_tempalte_proto_rawDesc
)

func file_tempalte_proto_rawDescGZIP() []byte {
	file_tempalte_proto_rawDescOnce.Do(func() {
		file_tempalte_proto_rawDescData = protoimpl.X.CompressGZIP(file_tempalte_proto_rawDescData)
	})
	return file_tempalte_proto_rawDescData
}

var file_tempalte_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tempalte_proto_goTypes = []interface{}{
	(*Success)(nil),        // 0: Success
	(*TemplateRpcAdd)(nil), // 1: TemplateRpcAdd
	(*PageList)(nil),       // 2: PageList
	(*Templates)(nil),      // 3: Templates
	(*TemplateVo)(nil),     // 4: TemplateVo
}
var file_tempalte_proto_depIdxs = []int32{
	4, // 0: Templates.templates:type_name -> TemplateVo
	1, // 1: TemplateRpcService.AddTemplate:input_type -> TemplateRpcAdd
	2, // 2: TemplateRpcService.List:input_type -> PageList
	0, // 3: TemplateRpcService.AddTemplate:output_type -> Success
	3, // 4: TemplateRpcService.List:output_type -> Templates
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tempalte_proto_init() }
func file_tempalte_proto_init() {
	if File_tempalte_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tempalte_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Success); i {
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
		file_tempalte_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRpcAdd); i {
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
		file_tempalte_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageList); i {
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
		file_tempalte_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Templates); i {
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
		file_tempalte_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateVo); i {
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
			RawDescriptor: file_tempalte_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tempalte_proto_goTypes,
		DependencyIndexes: file_tempalte_proto_depIdxs,
		MessageInfos:      file_tempalte_proto_msgTypes,
	}.Build()
	File_tempalte_proto = out.File
	file_tempalte_proto_rawDesc = nil
	file_tempalte_proto_goTypes = nil
	file_tempalte_proto_depIdxs = nil
}
