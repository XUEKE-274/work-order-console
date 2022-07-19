package entity

type WorkFlowEntity struct {
	BaseEntity
	Name string `json:"name"`
	Type string `json:"type"`

}