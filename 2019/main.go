package main

import (
	"fmt"
	"runtime"

	Day1 "./day1"
)

func main() {
	fmt.Println("*********************************")
	fmt.Println("runtime.Version: " + runtime.Version())
	fmt.Println("*********************************")
	Day1.Start()
}
