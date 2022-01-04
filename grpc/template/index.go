package template

import (
	"context"
	"go.uber.org/fx"
	"work-order-console/grpc/template/tRpc"
	"work-order-console/service"
)

var RegServerTemplate = fx.Provide(func(templateService service.TemplateServiceApi) *ServerTemplate {
	return &ServerTemplate{
		templateService,
		&tRpc.UnimplementedTemplateRpcServiceServer{},
	}
})

type ServerTemplate struct {
	service.TemplateServiceApi
	*tRpc.UnimplementedTemplateRpcServiceServer

}

func (p *ServerTemplate)AddTemplate(c context.Context, request *tRpc.TemplateRpcAdd) (*tRpc.Success, error) {
	p.SaveTemplate(request.Name)
	return &tRpc.Success{Ok: "SUCCESS"}, nil
}

func (p *ServerTemplate) List(context.Context, *tRpc.PageList) (*tRpc.Templates, error) {
	arr := p.QueryAll()
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



