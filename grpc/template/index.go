package template

import (
	"context"
	"work-order-console/grpc/template/tRpc"
)

type ServerTemplate struct {
	*tRpc.UnimplementedTemplateRpcServiceServer
}

func (p *ServerTemplate)AddTemplate(context.Context, *tRpc.TemplateRpcAdd) (*tRpc.Success, error) {
	return nil,nil
}

func (p *ServerTemplate) List(context.Context, *tRpc.PageList) (*tRpc.Templates, error) {

	var list []*tRpc.TemplateVo

	item := tRpc.TemplateVo{
		Id: "xx1",
		CreateBy: "System",
		CreateTime: "2022-01-04",
		ModifyBy: "System",
		ModifyTime: "2022-01-04",

		State: "ENABLE",
		Name: "SimpleTemplate",
	}

	list = append(list, &item)

	var res tRpc.Templates
	res.Templates = list

	return &res, nil

}



