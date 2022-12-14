package main

import "github.com/gin-gonic/gin"

func main() {

	// gin 指定代理，apis 原始请求转发到 k8s apiserver
	route := gin.Default()

	route.Any("/*proxy", proxyHandler)

	_ = route.Run(":8090")

}

func proxyHandler(c *gin.Context) {

}
