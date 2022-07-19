package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"work-order-console/domain/entity"
	"work-order-console/logger"
	"work-order-console/service"
	"work-order-console/web/request"
	"work-order-console/web/response"
)

type TemplateControllerApi interface {
	QueryAll(c *gin.Context)
	GrpcQuery(c *gin.Context)
	GrpcSave(c *gin.Context)
	CurrentSave(ctx *gin.Context)
}


type templateController struct {
	service.TemplateServiceApi
	logger.NewLogger
	service.GrpcServiceApi
}

var regTemplateController = fx.Provide(func(templateService service.TemplateServiceApi,
	logger logger.NewLogger, grpcService service.GrpcServiceApi) TemplateControllerApi {
	return &templateController{
		templateService,
		logger,
		grpcService,
	}
})

func (p *templateController)QueryAll(c *gin.Context)  {
	res := p.TemplateServiceApi.QueryAll()
	response.DataResponse(c, res)
}

func (p *templateController)GrpcQuery(c *gin.Context)  {
	res := p.Query()
	response.DataResponse(c, res)
}

func (p *templateController)GrpcSave(c *gin.Context)  {
	log := p.NewInstance(c)
	// request grpc
	params := &request.TemplateAddRequest{}
	e := c.BindJSON(&params)
	if e != nil {
		log.Error("json parse error")
		response.ParamsErrorResponse(c)
		return
	}
	p.Save(params.Name)
	response.NilResponse(c)
}
func (p *templateController)CurrentSave(c *gin.Context)  {
	log := p.NewInstance(c)
	params := request.NewTemplateAddRequest{}
	log.Info("params = ", params)
	e := c.BindJSON(&params)
	if e != nil {
		log.Error("json parse error")
		response.ParamsErrorResponse(c)
		return
	}
	session := c.Value("session")
	 u :=  session.(*entity.UserEntity)
	p.FullSave(params.TemplateName, params.WorkflowName, params.Fields, u.Username)
	response.NilResponse(c)
}