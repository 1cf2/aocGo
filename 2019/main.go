package main

import (
	"fmt"
	"runtime"
	"github.com/1cf2/AOC_golang/tree/master/2019/day1"
)

func main() {
	fmt.Println("*********************************")
	fmt.Println("runtime.Version: " + runtime.Version())
	fmt.Println("*********************************")
	Day1.Start()
}
