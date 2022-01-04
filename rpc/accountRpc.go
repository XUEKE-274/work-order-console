package rpc

import (
	"github.com/imroc/req"
	"go.uber.org/fx"
	"time"
	"work-order-console/config"
	"work-order-console/logger"
	"work-order-console/rpc/rpcReq"
	"work-order-console/rpc/rpcResp"
)

type AccountRpcApi interface {
	ListDepartment(request *rpcReq.DepartmentReq) *[]rpcResp.DepartmentResp
}

type accountRpc struct {
	logger logger.NewLogger
	host string
}

var accountRpcReg = fx.Provide(func(logger logger.NewLogger, config *config.Config) AccountRpcApi{
	return &accountRpc{
		logger,
		config.MkQwAccountSvc,
	}
})

func (mine *accountRpc)ListDepartment(request *rpcReq.DepartmentReq) *[]rpcResp.DepartmentResp {
	uri := mine.host + "/internal/qw_account/v1/address-book/department/list_all"
	body := req.BodyJSON(&request)
	r,_ := req.Post(uri, body)
	var entity rpcResp.DepartmentRpcResp
	e := r.ToJSON(&entity)
	if e != nil {}
	time.Sleep(6 * time.Second)
	return entity.Data

}
