package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"work-order-console/web/controller"
)

const Template = apiPrefix + "/template"

var regTemplateRouter = fx.Invoke(func(gin *gin.Engine, template controller.TemplateControllerApi) {
	templateGin := gin.Group(Template)
	templateGin.GET("/list", template.List)
	templateGin.GET("/grpc/list", template.GrpcQuery)
	templateGin.POST("/grpc/save", template.GrpcSave)
	templateGin.POST("/save", template.CurrentSave)


})