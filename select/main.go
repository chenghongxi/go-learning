package main

import (
	"fmt"
	"time"
)

// 通过 select 设置超时设置
func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				// do someting
				fmt.Println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
	fmt.Println("程序结束")
}
