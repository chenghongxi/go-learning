/*
Copyright 2021 The Pixiu Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package options

import (
	"go-learning/zaplog/log"
	"go-learning/zaplog/util"
	"strings"

	//pixiuConfig "github.com/caoyingjunz/pixiulib/config"

	"go-learning/zaplog/config"
)

type Options struct {
	ComponentConfig config.Config
}

func (o *Options) register() error {
	if err := o.registerLogger(); err != nil { // 注册日志
		return err
	}
	return nil
}

func (o *Options) registerLogger() error {
	logType := strings.ToLower(o.ComponentConfig.Default.LogType)
	if logType == "file" {
		// 判断文件夹是否存在，不存在则创建
		if err := util.EnsureDirectoryExists(o.ComponentConfig.Default.LogDir); err != nil {
			return err
		}
	}
	// 注册日志
	log.Register(logType, o.ComponentConfig.Default.LogDir, o.ComponentConfig.Default.LogLevel)
	return nil
}
