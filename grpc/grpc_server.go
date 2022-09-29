package main

import (
	"fmt"
	"go-learning/grpc/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()

	service.RegisterProdServiceServer(rpcServer, service.ProductService)

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("启动监听出错", err)
	}
	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatal("启动服务出错", err)
	}
	fmt.Println("启动grpc服务端成功")
}
