package template

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"testing"
	"work-order-console/grpc/template/tRpc"
)

func TestServerTemplate_List(t *testing.T) {
	//conn, err := grpc.Dial("1.14.108.113:8082", grpc.WithInsecure())
	conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := tRpc.NewTemplateRpcServiceClient(conn)

	r, err := c.List(context.Background(), &tRpc.PageList{})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("name = ", r.Templates[0].Name)
}