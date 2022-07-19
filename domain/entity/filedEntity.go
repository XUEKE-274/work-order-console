package entity


type FiledEntity struct {
	BaseEntity
	Name string `json:"name"`
	Type string `json:"type"`
}
