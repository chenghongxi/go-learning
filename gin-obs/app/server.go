package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"

	"go-learning/gin-obs/router"
	"go-learning/gin-obs/types"
)

func Newserver() error {
	opts, err := types.NewOptions()
	if err != nil {
		klog.Fatalf("unable to initialize command options: %v", err)
	}
	if err = Run(opts); err != nil {
		fmt.Println(err)
	}
	return nil
}

func Run(opt *types.Options) error {
	opt.GinEngine = gin.Default()

	// 初始化 APIs 路由
	router.InstallRouters(opt)

	// 启动 Server
	runServer(opt)
	return nil
}

func runServer(opt *types.Options) {
	srv := &http.Server{
		Addr:    types.Port,
		Handler: opt.GinEngine,
	}

	go func() {
		klog.Infof("starting OBS server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			klog.Fatal("failed to listen OBS server: ", err)
		}
	}()
}
