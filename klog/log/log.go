package log

import (
	"flag"

	"k8s.io/klog/v2"
)

func Register(LogOption string, LogPath string) {
	klog.InitFlags(nil)
	// By default klog writes to stderr. Setting logtostderr to false makes klog
	flag.Set("logtostderr", LogOption)
	flag.Set("log_file", LogPath)
	flag.Parse()
	klog.Info("nice to meet you")
	klog.Flush()
}
