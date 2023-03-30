package router

import (
	"github.com/emicklei/go-restful/v3"

	"go-learning/go-restful/router/cloud"
)

type Options struct {
	restfulEngine *restful.Container
}

func InstallRoutes(opt Options) {
	cloud.NewRouter(opt.restfulEngine) //  注册 cloud 路由
}
