package Day2

import (
	"fmt"
	"strconv"
)

type Day2 struct {
	noun   int
	verb   int
	OPCode []int
}

func Start() {
	fmt.Println("Day 2")
	e := Day2{
		noun:   12,
		verb:   2,
		OPCode: InputData{}.day2(),
	}
	e.Part1()
	e.Part2()
}

func (e Day2) InitCode() {
	ResetCode(e.OPCode, e.noun, e.verb)
}

func (e Day2) Part1() {
	e.InitCode()
	fmt.Println("  part 1: " + strconv.Itoa(RunCode(e.OPCode)))
}

func ResetCode(a []int, noun int, verb int) []int {
	a[1] = noun
	a[2] = verb
	return a
}

func RunCode(code []int) int {
	currentOpCode := 0
	for op := 0; currentOpCode != 99; op++ {
		currentOpCode = code[op]
		switch currentOpCode {
		case 1:
			code[code[op+3]] = code[code[op+1]] + code[code[op+2]]
			op += 3
		case 2:
			code[code[op+3]] = code[code[op+1]] * code[code[op+2]]
			op += 3
		case 99:
			return code[0]
		default:
			panic("invalid op code")
		}
	}
	return code[0]
}

func (e Day2) Part2() {
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			e.OPCode = InputData{}.day2()
			if RunCode(ResetCode(e.OPCode, i, j)) == 19690720 {
				fmt.Println("  part 2: " + strconv.Itoa(i) + "," + strconv.Itoa(j))
			}
		}
	}
}
