package template

import (
	"context"
	"go.uber.org/fx"
	"work-order-console/grpc/template/tRpc"
	"work-order-console/service"
)

type ServerTemplate struct {
	*tRpc.UnimplementedTemplateRpcServiceServer
}

func (p *ServerTemplate)AddTemplate(context.Context, *tRpc.TemplateRpcAdd) (*tRpc.Success, error) {
	return nil,nil
}

func (p *ServerTemplate) List(context.Context, *tRpc.PageList) (*tRpc.Templates, error) {
	var templateService service.TemplateServiceApi
	fx.Populate(&templateService)
	// todo 以来注入为什么失败
	arr := templateService.QueryAll()


	var list []*tRpc.TemplateVo
	for _, item := range *arr{
		var vo tRpc.TemplateVo
		vo.Id = item.Id
		vo.Name = item.Name
		vo.State = item.State

		vo.CreateBy = item.CreateBy
		vo.CreateTime = item.CreateTime.String()
		vo.ModifyBy = item.ModifyBy
		vo.ModifyTime = item.ModifyTime.String()
		list = append(list, &vo)
	}

	var res tRpc.Templates
	res.Templates = list

	return &res, nil

}



