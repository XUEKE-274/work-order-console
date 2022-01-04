package entity



type TemplateEntity struct {
	BaseEntity
	Name  string `json:"name"`
	State string `json:"state"`
}