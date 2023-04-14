package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ConfigMap struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
	Data struct {
		MicroConfigProdYaml string `yaml:"micro.config.prod.yaml"`
	} `yaml:"data"`
}

type Service struct {
	Name    string `yaml:"name"`
	Package string `yaml:"package"`
	Version int    `yaml:"version"`
	IP      string `yaml:"ip"`
	Port    int    `yaml:"port"`
}

type MicroConfigProd struct {
	Service struct {
		MessageMicroservice Service `yaml:"message_microservice"`
	} `yaml:"service"`
}

func main() {
	// 读取 YAML 文件
	data, err := ioutil.ReadFile("yaml/config.yaml")
	if err != nil {
		panic(err)
	}

	// 解析 YAML 文件
	var configMap ConfigMap
	err = yaml.Unmarshal(data, &configMap)
	if err != nil {
		panic(err)
	}

	// 获取 micro.config.prod.yaml 字段的值
	microConfigProdYaml := configMap.Data.MicroConfigProdYaml

	// 解析 micro.config.prod.yaml 字段的值
	var microConfigProd MicroConfigProd
	err = yaml.Unmarshal([]byte(microConfigProdYaml), &microConfigProd)
	if err != nil {
		panic(err)
	}

	// 打印解析结果
	fmt.Printf("%+v\n", configMap.Data.MicroConfigProdYaml)
}
