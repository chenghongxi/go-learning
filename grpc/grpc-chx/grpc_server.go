package main

import (
	"context"
	"fmt"
	service2 "go-learning/grpc/grpc-chx/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

func main() {
	//添加证书
	creds, err := credentials.NewServerTLSFromFile("D:\\go-work1\\go-learning\\grpc\\cert\\server.pem", "D:\\go-work1\\go-learning\\grpc\\cert\\server.key")
	if err != nil {
		log.Fatal("证书生成错误", err)
	}

	// 实现token认证，需要合法的用户名和密码
	// 实现一个拦截器
	var authInterceptor grpc.UnaryServerInterceptor
	authInterceptor = func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		err = Auth(ctx)
		if err != nil {
			return
		}
		//继续处理请求
		return handler(ctx, req)
	}

	rpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(authInterceptor))

	service2.RegisterProdServiceServer(rpcServer, service2.ProductService)

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

func Auth(ctx context.Context) error {
	//实际上 就是拿到传输的用户名和密码
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}
	var user string
	var password string
	if val, ok := md["user"]; ok {
		user = val[0]
	}
	if val, ok := md["password"]; ok {
		password = val[0]
	}
	if user != "admin" || password != "admin" {
		return status.Error(codes.Unauthenticated, "token不合法")
	}
	return nil
}
