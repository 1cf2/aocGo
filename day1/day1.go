package Day1

import (
	"fmt"
	"strconv"
)

type Day1 struct {
	fuel         int
	module_array []int
}

func calculate_fuel_1(mass int) int {
	return (int(mass / 3.0)) - 2
}

func calculate_fuel_2(mass int) int {
	fuel, temp := 0, 0
	for {
		temp = calculate_fuel_1(mass)
		if temp < 0 {
			break
		}
		fuel += temp
		mass = temp
	}
	return fuel
}

func (e Day1) Part1() {
	for i := 0; i < len(e.module_array); i++ {
		e.fuel += calculate_fuel_1(e.module_array[i])
	}
	fmt.Println("  part 1: " + strconv.Itoa(e.fuel))
}

func (e Day1) Part2() {
	for i := 0; i < len(e.module_array); i++ {
		e.fuel += calculate_fuel_2(e.module_array[i])
	}
	fmt.Println("  part 2: " + strconv.Itoa(e.fuel))
}

func Start() {
	e := Day1{
		fuel:         0,
		module_array: InputData{}.day1(),
	}
	fmt.Println("Day 1")
	e.Part1()
	e.Part2()
}
