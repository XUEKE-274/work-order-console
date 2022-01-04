package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"work-order-console/web/controller"
)

const System = apiPrefix + "/system"
var regSystemRouter = fx.Invoke(func(gin *gin.Engine, systemController controller.SystemControllerApi) {
	ticketGin := gin.Group(System)

	ticketGin.GET("/pubKey", systemController.FetchPubKey)


})