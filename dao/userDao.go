package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"work-order-console/db"
	"work-order-console/domain/entity"
	"work-order-console/domain/enum"
	"work-order-console/web/request"
)

type UserDaoApi interface {
	QueryByUsername(username string) (*entity.UserEntity, error)
	QueryById(id string) (*entity.UserEntity, error)
	RemoveById(id string) error
	UpdateById(id string, role enum.RoleEnum) error
	Save(po *entity.UserEntity) error
	PageQuery(params *request.UserListRequest) (*[]entity.UserEntity, int64, error)
}


type userDao struct {
	mongo *db.Mongodb
}

func (mine *userDao) PageQuery(params *request.UserListRequest) (*[]entity.UserEntity, int64, error) {
	c := userCollect(mine.mongo)
	f := bson.M{}
	if params.Username != "" {
		f["username"] = bson.M{"$regex" : params.Username}
	}
	if params.Role != "" {
		f["role"] = params.Role
	}

	var data []entity.UserEntity

	var findOption options.FindOptions
	if params.PageSize >0 {
		findOption.SetLimit(params.PageSize)
		findOption.SetSkip(params.PageIndex * params.PageSize)
	}
	total, err := c.CountDocuments(context.Background(), f)
	if err != nil {
		return &data, 0, err
	}

	cur, e := c.Find(context.Background(), f, &findOption)
	if e != nil {
		return &data, 0, e
	}

	err = cur.All(context.Background(), &data)
	if err != nil {
		return &data, 0, e
	}

	return &data, total, e

}

func (mine *userDao) UpdateById(id string, role enum.RoleEnum) error {
	c := userCollect(mine.mongo)
	hexId, _ := primitive.ObjectIDFromHex(id)
	f := bson.M{"_id": hexId}

	upd := bson.M{"$set" : bson.M{"role": role}}
	_, e := c.UpdateOne(context.Background(), f, upd)
	return e

}

func (mine *userDao) RemoveById(id string) error {
	c := userCollect(mine.mongo)
	hexId, _ := primitive.ObjectIDFromHex(id)
	f := bson.M{"_id": hexId}

	_, e := c.DeleteOne(context.Background(), f)
	return e
}

func (mine *userDao) QueryById(id string) (*entity.UserEntity, error) {
	c := userCollect(mine.mongo)
	hexId, _ := primitive.ObjectIDFromHex(id)
	f := bson.M{"_id": hexId}
	one := c.FindOne(context.Background(), f)
	ticket := entity.UserEntity{}
	e := one.Decode(&ticket)
	if e != nil {
		return &ticket, e
	}
	return &ticket, nil
}

var regUserDao = fx.Provide(func(mongo *db.Mongodb) UserDaoApi{
	return &userDao{
		mongo,
	}
})

func (mine *userDao) QueryByUsername(username string) (*entity.UserEntity, error)  {
	c := userCollect(mine.mongo)

	f := bson.M{"username": username}
	one := c.FindOne(context.Background(), f)
	ticket := entity.UserEntity{}
	e := one.Decode(&ticket)
	if e != nil {
		return &ticket, e
	}
	return &ticket, nil
}

func (mine *userDao) Save(po *entity.UserEntity) error {
	c := userCollect(mine.mongo)
	_, e := c.InsertOne(context.Background(), po)
	return e
}








func userCollect(mongo *db.Mongodb) *mongo.Collection {
	return Collection(mongo, "user")
}