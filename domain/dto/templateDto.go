package dto

import "work-order-console/domain/entity"

type TemplateDto struct {
	entity.TemplateEntity
	Workflow entity.WorkFlowEntity `json:"workflow" gorm:"foreignkey:template_id;association_foreignkey:id"`
	Fields []*entity.FiledEntity `json:"fields" gorm:"foreignkey:template_id;association_foreignkey:id"`

}
