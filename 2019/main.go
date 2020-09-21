package main

import (
	"fmt"
	"runtime"

	Day1 "github.com/1cf2/AOC_golang/2019/day1"
	Day2 "github.com/1cf2/AOC_golang/2019/day2"
)

func main() {
	fmt.Println("*********************************")
	fmt.Println("runtime.Version: " + runtime.Version())
	fmt.Println("*********************************")
	Day1.Start()
	Day2.Start()
}
