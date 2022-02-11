package service

import (
	"context"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"work-order-console/domain/entity"
	"work-order-console/grpc/template/tRpc"
	"work-order-console/logger"
)

type GrpcServiceApi interface {
	Query() *[]entity.TemplateEntity
	Save(name string)
}
type grpcService struct {
	logger.NewLogger
}

func (p *grpcService) Query() *[]entity.TemplateEntity  {
	// do remote rpc todo
	return nil
}

func (p *grpcService) Save(name string) {
	// do remote rpc  todo
	conn, err := grpc.Dial("1.14.108.113:8082", grpc.WithInsecure())
	//conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := tRpc.NewTemplateRpcServiceClient(conn)

	ok, e := c.AddTemplate(context.Background(), &tRpc.TemplateRpcAdd{Name: "simple"})
	if e != nil {
		log.Fatalf("did not connect: %v", e)
	}

	println("result = ", ok.Ok)

}

var regGrpcService = fx.Provide(func(logger logger.NewLogger) GrpcServiceApi{
	return &grpcService{
		logger,
	}
})



