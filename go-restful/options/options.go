package options

import (
	"github.com/emicklei/go-restful/v3"
)

type Options struct {
	// The default values.
	RestfulEngine *restful.Container

	ConfigFile string
}

func NewOptions() (*Options, error) {
	return &Options{}, nil
}
