package service

import (
	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)


type RedisServiceApi interface {
	Get(key string) string
	Set(key string, value string)
	Del(ket string)
	SetEx(key string, second int32, value string)
}

type redisService struct {
	pool *redis.Pool
	logger *logrus.Logger
}


var redisServiceReg = fx.Provide(func(pool *redis.Pool, logger *logrus.Logger) RedisServiceApi {
	return &redisService{
		pool,
		logger,
	}
})

func (mine *redisService)Get(key string) string {
	conn := mine.pool.Get()
	r, err := redis.String(conn.Do("get", key))
	if err != nil {
		mine.logger.Error("redis get error, key = ", key)
	}
	return r
}

func (mine *redisService)Set(key string, value string)  {
	conn := mine.pool.Get()
	_, err := conn.Do("set", key, value)
	if err != nil {
		mine.logger.Error("redis set error, key,value = ", key, value)
	}
}

func (mine *redisService)Del(key string)  {
	conn := mine.pool.Get()
	_, err := conn.Do("del", key)
	if err != nil {
		mine.logger.Error("redis del error, key ", key)
	}
}

func (mine *redisService)SetEx(key string, second int32, value string) {
	conn := mine.pool.Get()
	_, err := conn.Do("setex", key, second, value)
	if err != nil {
		mine.logger.Error("redis del error, key,value = ", key)
	}
}


