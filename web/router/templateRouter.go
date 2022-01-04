package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"work-order-console/web/controller"
)

const Template = apiPrefix + "/template"

var regTemplateRouter = fx.Invoke(func(gin *gin.Engine, template controller.TemplateControllerApi) {
	ticketGin := gin.Group(Template)
	ticketGin.GET("/list", template.QueryAll)
})