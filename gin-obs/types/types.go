package types

import "github.com/gin-gonic/gin"

const (
	AccessKey = "XXXXXXXXXXX"
	SecretKey = "XXXXXXXXXXXXXXXXXXXXXX"
	Endpoint  = "https://XXXXXX"
	Port      = "8090"
)

type Options struct {
	// The default values.
	GinEngine *gin.Engine
}

func NewOptions() (*Options, error) {
	return &Options{}, nil
}
