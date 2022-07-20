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
		Preload("Workflow").
		Find(&res)
	return &res
}
func (p *templateService) SqlList() *[]*dto.TemplateDto  {
	//var w = dto.WorkflowDto{}
	var res []*dto.TemplateDto
	//p.Raw(createSql()).Scan(&r)

	return &res
}
func createSql() string {
	sql := "select " +
		"t.id   templateId , " +
		"t.name templateName, " +
		"w.id workflowId , " +
		"w.name workflowName ," +
		"f.name filedName" +
		" from template_entity  t"  +
		"left join workflow_entity w " +
		" on t.id = w.template_id " +
		"left join filed_entity f  " +
		"on t.id = f.template_id"
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
