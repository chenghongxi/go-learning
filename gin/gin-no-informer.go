package main

import (
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"net/http"
	"path/filepath"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/pods", func(c *gin.Context) {
		list, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err, "code": 400})
		}
		c.JSON(http.StatusOK, gin.H{"message": 200, "result": list})
	})
	_ = r.Run(":8080")
}
