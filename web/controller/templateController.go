package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"work-order-console/logger"
	"work-order-console/service"
	"work-order-console/web/response"
)

type TemplateControllerApi interface {
	QueryAll(c *gin.Context)
}


type templateController struct {
	service.TemplateServiceApi
	logger.NewLogger
}

var regTemplateController = fx.Provide(func(templateService service.TemplateServiceApi, logger logger.NewLogger) TemplateControllerApi {
	return &templateController{
		templateService,
		logger,
	}
})

func (p *templateController)QueryAll(c *gin.Context)  {
	res := p.TemplateServiceApi.QueryAll()
	response.DataResponse(c, res)
}