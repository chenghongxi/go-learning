package router

import (
	"go-learning/go-restful/options"
	"go-learning/go-restful/router/cloud"
)

func InstallRoutes(opt *options.Options) {
	cloud.NewRouter(opt.RestfulEngine) //  注册 cloud 路由
}
