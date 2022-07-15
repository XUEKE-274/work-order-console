package msg

import (
	"context"
	"google.golang.org/grpc"
	"testing"
	"work-order-console/grpc/msg/pRpc"
)

func TestMsgServer(*testing.T)  {



}


func TestMsg(*testing.T)  {



	conn, err := grpc.Dial("127.0.0.1:8083", grpc.WithInsecure())
	if err != nil {
		panic("TestTcpTemplateList error")
	}
	in := &pRpc.PersonRequest{Id: "444"}
	out := new(pRpc.Person)
	ctx := context.Background()


	e := conn.Invoke(ctx, "/UserService/getPerson", in, out)
	if e != nil {
		panic("error >>>> ")
	}


	println(out.Id)
	println(out.Name)
	println(out.Age)

}