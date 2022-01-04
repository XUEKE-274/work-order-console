package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"work-order-console/web/controller"
)

const User = apiPrefix + "/user"
var regUserRouter = fx.Invoke(func(gin *gin.Engine, userController controller.UserControllerApi) {
	ticketGin := gin.Group(User)

	ticketGin.POST("/", userController.AddUser)
	ticketGin.PUT("/:id", userController.ModifyUser)
	ticketGin.DELETE("/:id", userController.DelUser)
	ticketGin.GET("/", userController.ListUser)
	ticketGin.GET("/:id", userController.DetailUser)

})
