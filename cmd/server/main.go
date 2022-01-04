package main

import (
	"go.uber.org/fx"
	"work-order-console/config"
	"work-order-console/dao"
	"work-order-console/db"
	"work-order-console/grpc"
	"work-order-console/logger"
	"work-order-console/rpc"
	"work-order-console/service"
	"work-order-console/web"
)

var Model = fx.Options(
	logger.Model,
	config.Model,
	db.Model,
	web.Model,
	service.Model,
	dao.Model,
	rpc.Model,
	grpc.Model,
)
// @Title 工单控制台
// @version 1.0.0
// @description 开发人员使用的后台运维系统
// @BasePath /work-order-console/api/v1
func main() {
	// check
	err := fx.ValidateApp(Model)
	if err != nil {
		panic(err)
	}
	// run
	fx.New(Model).Run()
}
