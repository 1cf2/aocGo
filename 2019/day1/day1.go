package day1

import (
	"fmt"
	"strconv"
)

// Day1 ...
type Day1 struct {
	fuel        int
	moduleArray []int
}

func calculateFuel1(mass int) int {
	return (int(mass / 3.0)) - 2
}

func calculateFuel2(mass int) int {
	fuel, temp := 0, 0
	for {
		temp = calculateFuel1(mass)
		if temp < 0 {
			break
		}
		fuel += temp
		mass = temp
	}
	return fuel
}

// Part1 ...
func (e Day1) Part1() {
	for i := 0; i < len(e.moduleArray); i++ {
		e.fuel += calculateFuel1(e.moduleArray[i])
	}
	fmt.Println("  part 1: " + strconv.Itoa(e.fuel))
}

// Part2 ...
func (e Day1) Part2() {
	for i := 0; i < len(e.moduleArray); i++ {
		e.fuel += calculateFuel2(e.moduleArray[i])
	}
	fmt.Println("  part 2: " + strconv.Itoa(e.fuel))
}

// Start ...
func Start() {
	e := Day1{
		fuel:         0,
		moduleArray: InputData{}.day1(),
	}
	fmt.Println("Day 1")
	e.Part1()
	e.Part2()
}
