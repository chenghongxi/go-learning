package main

import (
	"fmt"

	"go-learning/klog/klog-chx/options"
)

func main() {

	if err := register(); err != nil {
		return
	}

}

func register() error {
	// 注册日志
	if err := options.RegisterLogger(); err != nil {
		fmt.Println(err)
	}
	return nil
}
