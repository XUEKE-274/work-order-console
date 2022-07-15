package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type RpcHost struct {
	MkQwAccountSvc string
}

type MongoDb struct {
	MongoUri string
	Username string
	Password string
}

type Mysql struct {
	MysqlUri string
}


type RedisDb struct {
	RHost     string
	RPassword string
	Database  int
}

type LoginConfig struct {
	Rsa string
	TableInit string
}

type Config struct {
	MongoDb
	RpcHost
	RedisDb
	LoginConfig
	Mysql
}


var YmlConfig = fx.Provide(func() *Config{
	return &Config{
		MongoDb{
			MongoUri: viper.GetString("mongo.uri"),
			Username: viper.GetString("mongo.username"),
			Password: viper.GetString("mongo.password"),
		},
		RpcHost {
			MkQwAccountSvc: viper.GetString("rpc.mk-qw-account-svc"),
		},
		RedisDb{
			RHost : viper.GetString("redis.host"),
			RPassword: viper.GetString("redis.password"),
			Database: viper.GetInt("redis.database"),
		},
		LoginConfig{
			Rsa: viper.GetString("application.login.rsa"),
			TableInit: viper.GetString("application.table.init"),
		},
		Mysql{
			MysqlUri: viper.GetString("mysql.uri"),
		},
	}
})



func init() {
	if err := initConfig(); err != nil {
		panic(err)
	}
}

func  initConfig() error {
	viper.AddConfigPath("./")
	viper.SetConfigName("application.yml")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}