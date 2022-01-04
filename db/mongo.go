package db

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"time"
	"work-order-console/config"
)

var (
	MgoClient     *mongo.Client
	MgoSessionCtx mongo.SessionContext
)

type Mongodb struct {
	Client *mongo.Client
	Ctx    mongo.SessionContext
}

var mongodb = fx.Provide(func(logger *logrus.Logger, config *config.Config) *Mongodb{
	uri := config.MongoUri
	logger.Info("uri = ", uri)
	/*
	username := config.Username
	pwd := config.Password
	if ("" == uri) || ("" == username) || ("" == pwd) {
		panic("miss mongo bd config")
	}


	credential := options.Credential{
		AuthMechanism:           "SCRAM-SHA-1",
		AuthMechanismProperties: nil,
		//AuthSource:              "admin",
		//Username:                username,
		//Password:                pwd,
		//PasswordSet:             true,
	}
	*/

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	clientOptions := options.Client().
		ApplyURI(uri)
		//SetAuth(credential)
	//credentialSetMaxPoolSize(20).
	//SetMinPoolSize(10)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Info(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		logger.Info(err)
	}

	session, err := client.StartSession()
	if err != nil {
		logger.Info(err)
	}
	defer session.EndSession(context.Background())
	sessionCtx := mongo.NewSessionContext(ctx, session)

	logger.Println("mongo init")

	return &Mongodb{
		Client: client,
		Ctx:    sessionCtx,
	}
})