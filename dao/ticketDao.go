package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"work-order-console/db"
	"work-order-console/domain/entity"
)

type TicketDaoApi interface {
	GetById(id string) (*entity.TicketEntity, error)
	CountByAccount(account string) (int64, error)
}

type ticketDao struct {
	mongo *db.Mongodb
}

var ticketDaoReg = fx.Provide(func(mongo *db.Mongodb) TicketDaoApi {
	return &ticketDao{
		mongo,
	}
})

func (mine *ticketDao) GetById(id string) (*entity.TicketEntity, error) {
	c := ticketCollect(mine.mongo)
	hexId, _ := primitive.ObjectIDFromHex(id)
	f := bson.M{"_id": hexId}
	one := c.FindOne(context.Background(), f)
	ticket := entity.TicketEntity{}
	e := one.Decode(&ticket)
	if e != nil {
		return &ticket, e
	}
	return &ticket, nil
}

func (mine *ticketDao)CountByAccount(account string) (int64, error) {
	c := ticketCollect(mine.mongo)
	f := bson.M{"account": account}
	r, e := c.CountDocuments(context.Background(), f)
	if e != nil {
		return 0, e
	}
	return r, nil
}

func ticketCollect(mongo *db.Mongodb) *mongo.Collection {
	return Collection(mongo, "ticket")
}




