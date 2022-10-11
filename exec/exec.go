package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func upper(data string) string {
	return strings.ToUpper(data)
}

func main() {
	cmd := exec.Command("echo 111111111")
	//标准输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// Start函数执行； 它不会等待它完成
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// 从管道中读取数据
	data, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}

	// Wait 等待命令退出并等待任何复制到stdin或从stdout或stderr复制完成。看到命令退出
	// 关闭通道
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", upper(string(data)))
}
