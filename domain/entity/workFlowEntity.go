package entity

type WorkFlowEntity struct {
	BaseEntity
	TemplateId string `json:"templateId"`
	Name string `json:"name"`
	Type string `json:"type"`

}