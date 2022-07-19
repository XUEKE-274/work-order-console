package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"time"
	"work-order-console/domain/entity"
	"work-order-console/grpc/template/tRpc"
	"work-order-console/logger"
	"work-order-console/utils"
	"work-order-console/web/request"
)

type GrpcServiceApi interface {
	Query() *[]entity.TemplateEntity
	Save(name string)
	FullSave(string, string, []*request.FieldRequest, string)
}
type grpcService struct {
	logger.NewLogger
	db *gorm.DB
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


func (p *grpcService)FullSave(templateName string, workflowName string, fields []*request.FieldRequest, operator string){
	client := p.db
	// save template
	templateId := utils.NewUuid()
	t := entity.TemplateEntity{}
	t.Id = templateId
	t.Name = templateName
	t.State = "ON"
	t.ModifyTime = time.Now()
	t.CreateTime = time.Now()
	t.CreateBy = operator
	t.ModifyBy = operator
	client.Save(&t)
	// save workflow
	w := entity.WorkFlowEntity{}
	w.Id = utils.NewUuid()
	w.TemplateId = templateId
	w.Name = workflowName
	w.ModifyTime = time.Now()
	w.CreateTime = time.Now()
	w.CreateBy = operator
	w.ModifyBy = operator
	client.Save(&w)
	// save field
	if len(fields) != 0{
		for _, item := range fields {
			f := entity.FiledEntity{}
			f.Id = utils.NewUuid()
			f.TemplateId = templateId
			f.ModifyTime = time.Now()
			f.CreateTime = time.Now()
			f.CreateBy = operator
			f.ModifyBy = operator
			f.Name = item.Name
			f.Type = item.Type
			client.Save(&f)
		}
	}




}

var regGrpcService = fx.Provide(func(logger logger.NewLogger, db *gorm.DB) GrpcServiceApi{
	return &grpcService{
		logger,
		db,
	}
})



