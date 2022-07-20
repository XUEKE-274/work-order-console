package service

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/fx"
	"time"
	"work-order-console/domain/dto"
	"work-order-console/domain/entity"
	"work-order-console/logger"
	"work-order-console/utils"
)

type TemplateServiceApi interface {
	QueryAll() *[]entity.TemplateEntity
	SaveTemplate(name string)
	List() *[]*dto.TemplateDto
	SqlList() *[]*dto.TemplateDto
}

type templateService struct {
	logger.NewLogger
	*gorm.DB
}

func (p *templateService)QueryAll() *[]entity.TemplateEntity {
	var res []entity.TemplateEntity
	p.Find(&res)
	return &res
}

func (p *templateService) List() *[]*dto.TemplateDto  {
	//var w = dto.WorkflowDto{}
	var res []*dto.TemplateDto
	p.Preload("Fields").
		Preload("Workflow.Nodes").
		Find(&res)
	return &res
}
func (p *templateService) SqlList() *[]*dto.TemplateDto  {
	r := struct {
		TemplateId string
		TemplateName string
		TemplateState string
		WorkflowId string
		WorkflowName string
		WorkflowType string
		FiledId string
		FiledName string
		FiledType string
		NodeId string
		NodeName string
	}{}
	var res []*dto.TemplateDto
	p.Raw(createSql()).Scan(&r)

	return &res
}
func createSql() string {
	sql := "select \nt.id   TemplateId , \nt.name TemplateName, \nt.state TemplateState,\nw.id WorkflowId , " +
		"w.name WorkflowName ,\nw.type WorkflowType,\nf.id FiledId,\nf.name FiledName,\nf.type FiledType, " +
		" n.id NodeId,\nn.name NodeName\n from template_entity  t\nleft join workflow_entity w \non t.id = w.template_id\nleft join filed_entity f \non t.id = f.template_id\nleft join node_entity n\non w.id = n.workflow_id\n"


	return sql
}


func (p *templateService)SaveTemplate(name string)  {
	var po entity.TemplateEntity
	po.Id = utils.NewUuid()
	po.Name = name
	po.State = "ENABLE"
	po.CreateBy = "SYSTEM"
	po.ModifyBy = "SYSTEM"
	po.CreateTime = time.Now()
	po.ModifyTime = time.Now()
	p.Save(po)
}

var regTemplateService = fx.Provide(func(logger logger.NewLogger, db *gorm.DB) TemplateServiceApi {
	return &templateService{
		logger,
		db,
	}
})
