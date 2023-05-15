package main

import (
	"github.com/gin-gonic/gin"
	"go-learning/gin-obs/app"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	err := app.Newserver()
	if err != nil {
		os.Exit(1)
	}
}
