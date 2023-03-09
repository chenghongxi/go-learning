package main

import (
	"flag"
)

var strFlag = flag.String("s", "", "Description")

func main() {
	flag.Parse()
	println(*strFlag)
}
