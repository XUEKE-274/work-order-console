package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"work-order-console/web/controller"
)

const (
	Login =  apiPrefix + "/login"
	Logout = apiPrefix + "/logout"
)

var regLoginRouter = fx.Invoke(func(gin *gin.Engine, loginController controller.LoginControllerApi) {
	gin.POST(Login, loginController.Login)
	gin.GET(Logout, loginController.Logout)
})