package day4

import (
	"fmt"
	// "math/rand"
	"strconv"
	// "time"

	DV "github.com/1cf2/aocGo/2019/debugVisualizer"
)

func is6Digit(i int) bool {
	r := false
	if i > 99999 && i < 1000000 {
		r = true
	}
	return r
}

func isTwoDigitsSame(i int) bool {
	r := false
	s := strconv.Itoa(i)
	for j := 0; j < (len(s) - 1); j++ {
		if s[j] == s[j+1] {
			r = true
			break
		}
	}
	return r
}

func isThreeDigitsSame(i int) bool {
	r := false
	s := strconv.Itoa(i)
	for j := 0; j < (len(s) - 2); j++ {
		if (s[j] == s[j+1]) && (s[j] == s[j+2]) {
			r = true
			break
		}
	}
	return r
}

func isOnlyTwoDigitsSameExist(i int) bool {
	r := false
	s := strconv.Itoa(i)
	j := 0
	// c := "."[0]
	for j = 0; j < (len(s) - 1); j++ {
		// if (s[j] != c) && (s[j] == s[j+1]) {
		// 	if s[j] != s[j+2] {
		// 		r = true
		// 		break
		// 	} else {
		// 		c = s[j]
		// 		j += 2
		// 	}
		// }
		if s[j] == s[j+1] {
			if j > 0 && j < (len(s)-2) {
				if s[j-1] != s[j] && s[j+1] != s[j+2] {
					r = true
					return r
				}
			}
			if j == 0 {
				if s[j+1] != s[j+2] {
					r = true
					return r
				}
			}
			if j == (len(s) - 2) {
				if s[j-1] != s[j] {
					r = true
					return r
				}
			}
		}
	}
	return r
}

func isNeverDecreasing(i int) bool {
	r := true
	s := strconv.Itoa(i)
	for j := 0; j < (len(s) - 1); j++ {
		if s[j] > s[j+1] {
			r = false
			break
		}
	}
	return r
}

// Start ...
func Start() {
	fmt.Println("Day 4")

	// Part 1
	s := ""
	d4DataTable := DV.NewD4DataTable()

	m := ""
	mesh := DV.NewMesh()
	mesh.Data = []DV.MeshData{DV.MeshData{
		Type: "scatter3d",
		X:    []int{},
		Y:    []int{},
		Z:    []int{},
	}}
	md := &(mesh.Data[0])

	rangeFrom := 197487
	rangeTo := 673251
	for i := rangeFrom; i < rangeTo; i++ {
		if is6Digit(i) && isTwoDigitsSame(i) && isNeverDecreasing(i) {
			d4DataTable.Rows = append(d4DataTable.Rows, DV.D4Row{
				SixDigit:        strconv.Itoa(i),
				TwoDigitsSame:   "True",
				NeverDecreasing: "True",
			})
			md.X = append(md.X, i/10000)
			md.Y = append(md.Y, (i%10000)/100)
			md.Z = append(md.Z, i%100)
			s = d4DataTable.ToString()
			_ = s
			m = mesh.ToString()
			_ = m
			// debug visualizer break point here
		}
	}

	fmt.Println("  part 1: " + strconv.Itoa(len(d4DataTable.Rows)))

	// Part 2
	d4p2DataTable := DV.NewD4DataTable()
	for i := range d4DataTable.Rows {
		n, _ := strconv.ParseInt(d4DataTable.Rows[i].SixDigit, 10, 0)
		if isOnlyTwoDigitsSameExist(int(n)) {
			d4p2DataTable.Rows = append(d4p2DataTable.Rows, d4DataTable.Rows[i])
			s = d4p2DataTable.ToString()
			_ = s
		}
	}
	fmt.Println("  part 1: " + strconv.Itoa(len(d4p2DataTable.Rows)))
	_ = s
}
