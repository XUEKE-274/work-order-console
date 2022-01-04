package entity

import "time"

type BaseEntity struct {
	Id       string `json:"_id"`
	CreateTime time.Time `json:"createTime"`
	ModifyTime time.Time `json:"modifyTime"`
	CreateBy   string `json:"createBy"`
	ModifyBy   string `json:"modifyBy"`
}
