package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"work-order-console/web/controller"
)

const System = apiPrefix + "/system"
var regSystemRouter = fx.Invoke(func(gin *gin.Engine, systemController controller.SystemControllerApi) {
	router := gin.Group(System)

	router.GET("/pubKey", systemController.FetchPubKey)
	router.GET("/test", systemController.Test)



})