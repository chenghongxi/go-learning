package main

import (
	"fmt"
	"sync"
)

var waitgroup sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		waitgroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1
		go func(i int) {
			fmt.Println(i)
			waitgroup.Done()
		}(i)
	}
	waitgroup.Wait() //.Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
}
