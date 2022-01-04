package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)


const (
	SwaggerPath = "/swagger"
	swaggerAll = "/swagger/*any"
)

var regSwaggerRouter = fx.Invoke(func(g *gin.Engine) {
	g.GET(swaggerAll, ginSwagger.WrapHandler(swaggerFiles.Handler))
})