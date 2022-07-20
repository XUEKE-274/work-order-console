package dto

import "work-order-console/domain/entity"

type WorkflowDto struct {
	entity.WorkflowEntity
	Nodes []*entity.NodeEntity `json:"nodes" gorm:"foreignkey:workflow_id;association_foreignkey:id"`
}

func (WorkflowDto) TableName() string {
	return "workflow_entity"
}
