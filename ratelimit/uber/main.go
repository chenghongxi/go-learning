package main

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
	limiter := ratelimit.New(100)

	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := limiter.Take()
		if i > 0 {
			fmt.Println(i, now.Sub(prev))
		}
		prev = now
	}
}
