package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"work-order-console/logger"
	"work-order-console/service"
	"work-order-console/web/response"
)

type TicketControllerApi interface {
	CountTicket(c *gin.Context)
	Detail(c *gin.Context)
}

type ticketController struct {
	ticketService service.TicketServiceApi
	logger logger.NewLogger
}

var regTicketController = fx.Provide(func(ticketService service.TicketServiceApi, logger logger.NewLogger) TicketControllerApi {
	return &ticketController{
		ticketService,
		logger,
	}
})

func (mine *ticketController) CountTicket(c *gin.Context) {

	account := c.Query("account")
	if account == "" {
		response.ParamsErrorResponse(c)
		return
	}
	res, e := mine.ticketService.CountTicket(account)
	if e != nil {
		response.ErrorResponse(c, e)
		return
	}
	response.DataResponse(c, res)
}

func (mine *ticketController) Detail(c *gin.Context) () {
	id := c.Param("id")
	if id == "" {
		response.ParamsErrorResponse(c)
		return
	}
	res, e := mine.ticketService.Detail(id, c)

	if e != nil {
		response.ErrorResponse(c, e)
		return
	}

	response.DataResponse(c, res)
}
