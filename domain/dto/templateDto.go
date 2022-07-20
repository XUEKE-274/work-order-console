package dto

import (
	"work-order-console/domain/entity"
)

type TemplateDto struct {
	entity.TemplateEntity
	Workflow WorkflowDto `json:"workflow" gorm:"foreignkey:template_id;association_foreignkey:id"`
	Fields []*entity.FiledEntity   `json:"fields" gorm:"foreignkey:template_id;association_foreignkey:id"`

}


func (TemplateDto) TableName() string {
	return "template_entity"
}
