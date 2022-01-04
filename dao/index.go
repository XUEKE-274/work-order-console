package dao

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"work-order-console/db"
)

var Model = fx.Options(
	regUserDao,
	ticketDaoReg,
)


func Collection(mongo *db.Mongodb, tableName string) *mongo.Collection{
	return mongo.Client.Database("kms").Collection(tableName)
}