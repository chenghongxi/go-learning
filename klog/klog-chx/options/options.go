package options

import (
	"go-learning/klog/klog-chx/log"
	"os"
	"strings"
)

var (
	// 输出到文件路径
	LogPath = "klog/chx/"
	// true: 标准输出 false: 输出到文件
	LogOptions = "False"
	// 日志类型

)

func RegisterLogger() error {
	logOption := strings.ToLower(LogOptions)
	if logOption == "false" {
		// 判断文件夹是否存在，不存在则创建
		if err := EnsureDirectoryExists(LogPath); err != nil {
			return err
		}
	}
	// 注册日志
	log.Register(logOption, LogPath)

	return nil
}

func EnsureDirectoryExists(path string) (err error) {
	if !IsDirectoryExists(path) {
		err = os.MkdirAll(path, 0755)
	}
	return
}

func IsDirectoryExists(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}
