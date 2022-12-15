package main

import (
	"context"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// 使用 informer 减轻 kubernetes apiserver 并发压力
// use two clients
// 一个 client 负责查资源
// 一个 client 负责资源的创建，更新，删除
func main() {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get all namespace sharedInformers
	sharedInformers := informers.NewSharedInformerFactory(clientset, 0)

	// get Resource
	gvrs := []schema.GroupVersionResource{
		{Group: "apps", Version: "v1", Resource: "deployments"},
		{Group: "", Version: "v1", Resource: "pods"},
	}

	for _, gvr := range gvrs {
		if _, err = sharedInformers.ForResource(gvr); err != nil {
			panic(err)
		}
	}

	// Start all informers
	sharedInformers.Start(ctx.Done())
	// Wait for all caches to sync
	sharedInformers.WaitForCacheSync(ctx.Done())

	log.Printf("all informers has been started")

	// 构造 pod Lister，用于 gin 的查询
	podLister := sharedInformers.Core().V1().Pods().Lister()
	// 启动 gin router
	// 仅作演示， 无封装， 无异常处理
	// 启动之后，curl 127.0.0.1：8080/pods
	r := gin.Default()
	r.GET("/pods", func(c *gin.Context) {
		pods, err := podLister.List(labels.Everything())
		if err != nil {
			panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err, "code": 400})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "pong", "code": 200, "result": pods})
	})
	_ = r.Run(":8080")
}
