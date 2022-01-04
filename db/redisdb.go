package db

import (
	"github.com/garyburd/redigo/redis"
	"go.uber.org/fx"
	"work-order-console/config"
)

var redisDb = fx.Provide(func(config *config.Config) *redis.Pool{

	setDatabase := redis.DialDatabase(config.Database)
	setPassword := redis.DialPassword(config.RPassword)

	return &redis.Pool{
		MaxActive: 0,
		MaxIdle: 8,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", config.RHost, setDatabase, setPassword)
		},
	}

})

