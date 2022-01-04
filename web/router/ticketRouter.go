package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"work-order-console/web/controller"
)

const Ticket = apiPrefix + "/ticket"

var regTicketRouter = fx.Invoke(func(gin *gin.Engine, ticketController controller.TicketControllerApi) {
	ticketGin := gin.Group(Ticket)

	ticketGin.GET("/count", ticketController.CountTicket)
	ticketGin.GET("/detail/:id", ticketController.Detail)

})