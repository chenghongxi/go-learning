package router

import (
	"go-learning/gin-obs/types"
)

func InstallRouters(ginEngine *types.Options) {

	uploadgroup := ginEngine.GinEngine.Group("/upload")
	{
		uploadgroup.POST("", Upload)
	}

	downloadRoute := ginEngine.GinEngine.Group("/download")
	{
		downloadRoute.GET("", Download)
	}
}
