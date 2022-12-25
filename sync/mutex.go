package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var a = 0
	var lock sync.Mutex
	for i := 0; i < 1000; i++ {
		go func(idx int) {
			lock.Lock()
			defer lock.Unlock()
			a += 1
			fmt.Printf("goroutine %d, a=%d\n", idx, a)
		}(i)
	}
	time.Sleep(1e9)
}
