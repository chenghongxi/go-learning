package main

import (
	"fmt"
	"go-learning/klog/options"
)

func main() {
	if err := register(); err != nil {
		fmt.Println(err)
	}

}

func register() error {
	// 注册日志
	if err := options.RegisterLogger(); err != nil {
		fmt.Println(err)
	}
	return nil
}
