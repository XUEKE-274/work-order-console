package entity

type VersionEntity struct {
	BaseEntity
	Version  int32 `json:"version"`

}