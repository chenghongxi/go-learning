package main

import (
	"flag"
	"k8s.io/klog/v2"
)

var LogPath = "klog/myfile.log"

func main() {

	klog.InitFlags(nil)
	// By default klog writes to stderr. Setting logtostderr to false makes klog
	// write to a log file.
	flag.Set("logtostderr", "false")
	flag.Set("log_file", LogPath)
	flag.Parse()
	klog.Info("nice to meet you")
	klog.Flush()
}
