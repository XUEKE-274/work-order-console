package response

import (
	"work-order-console/domain/dto"
	"work-order-console/domain/entity"
)

type TemplateResponse struct {
	List *[]*TemplateVo
}
type TemplateVo struct {
	Id string
	Name string
	State string
	Workflow *WorkflowVo
	Fields *[]*FiledVo
}

type WorkflowVo struct {
	Id string
	Name string
	Type string
	Nodes *[]*NodeVo
}
type NodeVo struct {
	Id string
	Name string
}
type FiledVo struct {
	Id string
	Name string
	Type string
}

func CoverTemplateResponse(in *[]*dto.TemplateDto)  *TemplateResponse{
	if in == nil || len(*in) == 0 {
		return &TemplateResponse{}
	}
	res := TemplateResponse{}
	templates := make([]*TemplateVo, 0,  len(*in))
	for _, item := range *in {
		t := &TemplateVo{}
		t.Id = item.Id
		t.Name = item.Name
		t.State = item.State
		t.Fields = newFiledVo(&item.Fields)
		t.Workflow = newWorkflow(&item.Workflow)
		templates = append(templates, t)
	}
	res.List = &templates
	return &res
}

func newWorkflow(in *dto.WorkflowDto) *WorkflowVo {
	if in == nil  {
		return &WorkflowVo{}
	}

	r := WorkflowVo{}
	r.Id = in.Id
	r.Name = in.Name
	r.Type = in.Type
	r.Nodes = newNodes(&in.Nodes)
	return &r
}

func newNodes(in *[]*entity.NodeEntity) *[]*NodeVo {
	if in == nil || len(*in) == 0 {
		r :=  make([]*NodeVo,0)
		return &r
	}

	res :=  make([]*NodeVo,0, len(*in))
	for _, item := range *in {
		f := NodeVo{}
		f.Id = item.Id
		f.Name = item.Name
		res = append(res, &f)
	}
	return &res
}

func newFiledVo(in *[]*entity.FiledEntity) *[]*FiledVo {
	if in == nil || len(*in) == 0 {
		r :=  make([]*FiledVo,0)
		return &r
	}
	res := make([]*FiledVo,0, len(*in))

	for _, item := range *in {
		f := FiledVo{}
		f.Id = item.Id
		f.Name = item.Name
		f.Type = item.Type
		res = append(res, &f)
	}
	return &res

}
