package server

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"net"
	"work-order-console/grpc/template"
	"work-order-console/grpc/template/tRpc"
)
import "github.com/gin-gonic/gin"




type targets struct {
	fx.In
	LogHandler gin.HandlerFunc `name:"logHandler"`
	AuthHandler gin.HandlerFunc `name:"authHandler"`
	AuthorizationHandler gin.HandlerFunc `name:"authorizationHandler"`
}

var InvokeGinStart = fx.Provide(func(logger *logrus.Logger, t targets) *gin.Engine{
	gin.DefaultWriter = logger.Writer()
	g := gin.Default()
	g.Use(t.LogHandler)
	g.Use(t.AuthHandler)
	g.Use(t.AuthorizationHandler)
	return g
})


var InvokeGrpc = fx.Invoke(func(logger *logrus.Logger, templateServer *template.ServerTemplate) {
	go func() {
		logger.Info("Grpc start")
		lis, err := net.Listen("tcp", ":8082")

		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()

		tRpc.RegisterTemplateRpcServiceServer(s, templateServer)

		e := s.Serve(lis)
		if e != nil {
			panic("init grpc error ")
		}
	}()
})