package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"work-order-console/web/response"
)

const Healthy = serverName + "/healthy"

var regProbeRouter = fx.Invoke(func(g *gin.Engine) {
	g.GET(Healthy, func(c *gin.Context) {
		response.NilResponse(c)
	})
})

