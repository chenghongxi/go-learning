package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	"go-learning/loadconfig/config"
)

const file = "loadconfig/test.yaml"

func loadConfigFromFile(file string) (*config.PixiuConfiguration, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return loadConfig(data)
}

func loadConfig(data []byte) (*config.PixiuConfiguration, error) {
	var pc config.PixiuConfiguration
	if err := yaml.Unmarshal(data, &pc); err != nil {
		return nil, err
	}

	return &pc, nil
}

func main() {
	pixiuConfiguration, err := loadConfigFromFile(file)
	if err != nil {
		panic(err)
	}
	// 向结构体中重新写入数据
	pixiuConfiguration.Mysql.Host = "9.9.9.9"
	out, err := yaml.Marshal(pixiuConfiguration)
	// 通过ioutil.writeFile写入文件
	err = ioutil.WriteFile(file, out, 0777)
	fmt.Println(pixiuConfiguration)

}
