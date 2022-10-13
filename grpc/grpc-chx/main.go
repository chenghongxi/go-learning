package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-learning/grpc/grpc-chx/service"
)

func main() {
	user := &service.User{
		Username: "chengxiaobai",
		Age:      18,
	}
	// 序列化的过程
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	// 反序列化
	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(newUser.Username, newUser.Age)
}
