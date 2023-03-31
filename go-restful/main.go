package main

import (
	"go-learning/go-restful/options"
	"go-learning/go-restful/router"
	"go-learning/go-restful/router/cloud"
	"k8s.io/klog/v2"
)

func main() {

	opts := &options.Options{}

	router.InstallRoutes(opts)

	runBootstrap()

}

func runBootstrap() {
	// 处理 kubeconfig
	if err := cloud.Load(); err != nil {
		klog.Fatal("failed to load cloud driver: ", err)
	}

}
