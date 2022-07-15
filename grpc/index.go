package grpc

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"work-order-console/grpc/template"
)

var Model = fx.Options(
	template.RegServerTemplate,
	regGrpcPool,

)




type Pool struct {
	cs []*grpc.ClientConn
}

var regGrpcPool = fx.Provide(func() *Pool{
	// 初始化 2个
	conn1, err1 := grpc.Dial("1.14.108.113:8082", grpc.WithInsecure())
	if err1 != nil {
		panic("grpc conn init fail")
	}

	conn2, err2 := grpc.Dial("1.14.108.113:8082", grpc.WithInsecure())
	if err2 != nil {
		panic("grpc conn init fail")
	}

	cs := make([]*grpc.ClientConn, 0)
	cs = append(cs, conn1)
	cs = append(cs, conn2)

	return &Pool{cs}
})

func (p *Pool) GetConn() *grpc.ClientConn {
	for index, item := range p.cs {
		if item == nil {
			println("GetConn = nil, index = ", index)
		} else {
			println("GetConn, index = ", index)
			temp := item
			item = nil
			return temp
		}
	}
	// result


	println("not has conn")
	return nil
}

func (p *Pool) ReturnConn(c *grpc.ClientConn)  {

}
