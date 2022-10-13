package main

import (
	"context"
	"fmt"
	"go-learning/grpc/grpc-chx/client/auth"
	"go-learning/grpc/grpc-chx/service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	// 加载服务端证书
	creds, err := credentials.NewClientTLSFromFile("D:\\go-work1\\go-learning\\grpc\\cert\\server.pem", "*.xiaobai.com")
	if err != nil {
		log.Fatal("证书错误", err)
	}

	//可以实现jmt，auth等
	token := &auth.Authentication{
		User:     "admin",
		Password: "admin",
	}
	// 创建连接服务端,并读取creds证书和用户名，密码
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(token))
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
