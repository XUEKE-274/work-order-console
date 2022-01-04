package service

import (
	"context"
	"go.uber.org/fx"
	"time"
	"work-order-console/dao"
	"work-order-console/domain/entity"
	"work-order-console/domain/enum/errorCodeEnum"
	"work-order-console/exception"
	"work-order-console/logger"
	"work-order-console/rpc"
	"work-order-console/rpc/rpcReq"
	"work-order-console/rpc/rpcResp"
	"work-order-console/web/response"
)

type TicketServiceApi interface {
	CountTicket(account string) (*response.TicketCountResponse, error)
	Detail(id string, ctx context.Context) (*response.TicketDetailResponse, error)
}

type ticketService struct {
	logger logger.NewLogger
	ticketDao dao.TicketDaoApi
	accountRpc rpc.AccountRpcApi
	redisService RedisServiceApi
}

var regTicketService = fx.Provide(func(logger logger.NewLogger,
							           ticketDao dao.TicketDaoApi,
									   accountRpc rpc.AccountRpcApi,
									   redisService RedisServiceApi) TicketServiceApi {

	return &ticketService{
		logger,
		ticketDao,
		accountRpc,
		redisService,
	}
})

func (mine *ticketService) CountTicket(account string) (*response.TicketCountResponse, error) {
	r, e := mine.ticketDao.CountByAccount(account)

	if nil != e {
		return nil, e
	}
	return &response.TicketCountResponse{
		Account: account,
		Count:   r,
	}, exception.NewException(errorCodeEnum.TEST_ERROR)
}

func (mine *ticketService) Detail(id string, ctx context.Context) (*response.TicketDetailResponse, error) {

	log := mine.logger.NewInstance(ctx)

	ticket, e := mine.ticketDao.GetById(id)
	if nil != e {
		return nil, e
	}

	// rpc
	var req rpcReq.DepartmentReq
	req.DeptIds = []int64{1}
	req.Account = "W00000000018"
	req.AccountDB = "mk_data_20191209"
	req.ExcludeDelete = true

	start := time.Now().Unix()

	ch := make(chan *[]rpcResp.DepartmentResp)
	go func() {
		ch <- mine.accountRpc.ListDepartment(&req)
	}()

	time.Sleep(5 * time.Second)

	res := <- ch
	end := time.Now().Unix()

	log.Info("cost = ", end -start)
	close(ch)

	for _, item := range *res {
		log.Info("Call Account Rpc, Name = ", item.Name)
	}

	return toDetailVo(ticket), nil

}

func toDetailVo(e *entity.TicketEntity) *response.TicketDetailResponse {



	return &response.TicketDetailResponse{
		TicketNumber:    e.TicketNumber,
		Id:              e.Id,
		CustomFieldList: e.CustomFieldList,
	}
}
