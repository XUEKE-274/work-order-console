package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"work-order-console/domain/enum"
)

type UserEntity struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	CreateTime time.Time `json:"createTime" bson:"createTime"`
	ModifyTime time.Time `json:"modifyTime" bson:"modifyTime"`
	CreateBy   string `json:"createBy" bson:"createBy"`
	ModifyBy   string `json:"modifyBy" bson:"modifyBy"`
	Username string        `json:"username" bson:"username"`
	EnPwd    string        `json:"enPwd" bson:"enPwd"`
	Role     enum.RoleEnum `json:"role" bson:"role"`
}