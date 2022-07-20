package dto

import "work-order-console/domain/entity"

type WorkflowDto struct {
	entity.WorkflowEntity
	Nodes []*entity.NodeEntity
}

func (WorkflowDto) TableName() string {
	return "workflow_entity"
}
