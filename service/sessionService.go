package service

import (
	"context"
	"encoding/json"
	"go.uber.org/fx"
	"work-order-console/domain/entity"
	"work-order-console/utils"
)

type SessionServiceApi interface {
	GenToken(po *entity.UserEntity) string
	FetchSessionByToken(token string) (*entity.UserEntity, bool)
	ContextSession(ctx context.Context) *entity.UserEntity

}

type sessionService struct {
	redisService RedisServiceApi
}



var regSessionService = fx.Provide(func(redisService RedisServiceApi) SessionServiceApi {
	return &sessionService{
		redisService,
	}
})

func (mine *sessionService) GenToken(po *entity.UserEntity) string {
	userStr, _ := json.Marshal(po)
	token := utils.NewUuid()
	mine.redisService.SetEx(token, 60 * 60 * 2, string(userStr))
	return token
}

func (mine *sessionService) FetchSessionByToken(token string) (*entity.UserEntity, bool) {
	var res = entity.UserEntity{}
	sessionStr := mine.redisService.Get(token)
	if sessionStr == "" {
		return &res, false
	}

	e := json.Unmarshal([]byte(sessionStr), &res)
	if e!= nil {
		panic("redis error")
	}
	return &res, true
}

func (mine *sessionService) ContextSession(ctx context.Context) *entity.UserEntity {
	return ctx.Value("session").(*entity.UserEntity)
}



