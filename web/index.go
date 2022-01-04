package web

import (
	"go.uber.org/fx"
	"work-order-console/web/controller"
	"work-order-console/web/handler"
	"work-order-console/web/router"
	"work-order-console/web/server"
)

var Model =fx.Options(
	server.InvokeGinStart,
	handler.Model,
	controller.Model,
	router.Model,
	server.InvokeGrpc,
)
