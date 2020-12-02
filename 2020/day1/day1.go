package day1

import (
	"fmt"
	"strconv"
)

// Day1 ...
type Day1 struct {
	expenseReport []int
}

func find2Numbers(a []int, sumVal int) int {
	returnVal := -1
	for x, y := 0, 0; len(a) > 1; {
		x, a = a[0], a[1:]
		y = sumVal - x
		for i := range a {
			if a[i] == y {
				// Found!
				returnVal = x * y
				break
			}
		}
	}
	return returnVal
}

// Part1 ...
func (e Day1) Part1() {
	fmt.Println("  part 1: " + strconv.Itoa(find2Numbers(e.expenseReport, 2020)))
}

// Part2 ...
func (e Day1) Part2() {
	a := e.expenseReport
	returnVal := 0
	for x, y := 0, 0; len(a) > 1; {
		x, a = a[0], a[1:]
		y = find2Numbers(a, 2020-x)
		if y > 0 {
			returnVal = x * y
			break
		}
	}
	fmt.Println("  part 2: " + strconv.Itoa(returnVal))
}

// Start ...
func Start() {
	e := Day1{
		expenseReport: InputData{}.day1(),
	}
	fmt.Println("Day 1")
	e.Part1()
	e.Part2()
}
