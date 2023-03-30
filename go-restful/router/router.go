package router

import (
	"go-learning/go-restful/options"
	"go-learning/go-restful/router/cloud"
	"go-learning/go-restful/router/menu"
)

func InstallRoutes(opt *options.Options) {
	cloud.NewRouter(opt.RestfulEngine) //  注册 cloud 路由

	menu.NewRouter(opt.RestfulEngine) //  注册 menu 路由
}
