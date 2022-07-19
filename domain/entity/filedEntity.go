package entity


type FiledEntity struct {
	BaseEntity
	TemplateId string `json:"templateId"`
	Name string `json:"name"`
	Type string `json:"type"`
}
