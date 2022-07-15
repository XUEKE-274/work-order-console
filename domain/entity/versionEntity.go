package entity

type VersionEntity struct {
	BaseEntity
	Resource string `json:"resource"`
	Version  int32 `json:"version"`

}