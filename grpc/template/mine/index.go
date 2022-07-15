package mine



type Templates struct {
	Templates []*TemplateVo
}


type TemplateVo struct {
	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateBy   string `protobuf:"bytes,2,opt,name=createBy,proto3" json:"createBy,omitempty"`
	ModifyBy   string `protobuf:"bytes,3,opt,name=modifyBy,proto3" json:"modifyBy,omitempty"`
	CreateTime string `protobuf:"bytes,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	ModifyTime string `protobuf:"bytes,5,opt,name=modifyTime,proto3" json:"modifyTime,omitempty"`
	State      string `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	Name       string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
}