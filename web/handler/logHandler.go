package handler

import (
	"github.com/gin-gonic/gin"
	"work-order-console/logger"
	"work-order-console/utils"
)

var logHandler = func(logger logger.NewLogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := utils.NewUuid()
		ctx.Set("requestId", uuid)
		log := logger.NewInstance(ctx)
		log.Info("do logHandler  ")

		ctx.Next()
	}
}