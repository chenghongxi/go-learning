package log

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"os"
)

func Register(LogOption string, LogPath string) {

	write, _ := os.OpenFile("error.log", os.O_WRONLY|os.O_CREATE, 0755)
	klog.SetOutput(write)
	klog.Info("nice to meet you")

}
