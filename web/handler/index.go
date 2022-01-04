package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"work-order-console/logger"
	"work-order-console/service"
)



type result struct {
	fx.Out
	LogHandler gin.HandlerFunc `name:"logHandler"`
	AuthHandler gin.HandlerFunc `name:"authHandler"`
	AuthorizationHandler gin.HandlerFunc `name:"authorizationHandler"`
}



var Model = fx.Provide(func(logger logger.NewLogger, sessionService service.SessionServiceApi) result{
	return result{
		LogHandler: logHandler(logger),
		AuthHandler: authHandler(logger, sessionService),
		AuthorizationHandler: authorizationHandler(logger, sessionService),
	}
})

