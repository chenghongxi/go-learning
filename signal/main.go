package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"k8s.io/klog/v2"
)

func main() {
	// 创建监听chan 监听源：系统信号
	c := make(chan os.Signal)
	signal.Notify(c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT:
				klog.Error("shutdown server......")
				exit()

			default:
				klog.Error("other signal shutdown server......")
			}
		}
	}()

	klog.Info("starting server......")
	sum := 0
	for {
		sum++
		klog.Infoln("sum :", sum)
		time.Sleep(time.Second)
	}
}

func exit() {
	klog.Infoln("Execute Clean...")
	klog.Infoln("End Exit...")
	os.Exit(0)
}
