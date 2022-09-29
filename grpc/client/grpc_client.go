package main

import (
	"context"
	"fmt"
	"go-learning/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	//创建连接服务端
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("fail connect server ", err)
	}

	defer conn.Close()
	prodClient := service.NewProdServiceClient(conn)
	request := &service.ProductRequest{
		ProdId: 123,
	}
	stockResponse, err := prodClient.GetProductStock(context.Background(), request)
	if err != nil {
		log.Fatal("查询出错 ", err)
	}
	fmt.Println("查询成功", stockResponse.ProdStock)
}
