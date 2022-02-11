package service

import (
	"go.uber.org/fx"
	"work-order-console/domain/entity"
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

}

var regGrpcService = fx.Provide(func(logger logger.NewLogger) GrpcServiceApi{
	return &grpcService{
		logger,
	}
})



