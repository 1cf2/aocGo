package day3

import (
	"fmt"
	"reflect"
	"strconv"
)

// Coordinate ...
type Coordinate struct {
	x int
	y int
}

// MatrixDot ...
type MatrixDot struct {
	relativeCoord             Coordinate //relative coordinate to starting point
	wire1Pass                 bool
	wire2Pass                 bool
	manhattanDistanceToCenter int
	wire1Steps                int
	wire2Steps                int
}

// MatrixRange ...
type MatrixRange struct {
	maxX int
	minX int
	maxY int
	minY int
}

// Day3 ...
type Day3 struct {
	wire1        []Move
	wire2        []Move
	matrix       [][]MatrixDot
	matrixCenter Coordinate
	matrixRange  MatrixRange
	crossPoints  []MatrixDot
}

//Max ...
func Max(a int, b int) int {
	var c int
	if c = b; a > b {
		c = a
	}
	return c
}

//Min ...
func Min(a int, b int) int {
	var c int
	if c = b; a < b {
		c = a
	}
	return c
}

// MixMatrixRange ...
func MixMatrixRange(mr1 MatrixRange, mr2 MatrixRange) MatrixRange {
	mixedMR := MatrixRange{
		maxX: Max(mr1.maxX, mr2.maxX),
		maxY: Max(mr1.maxY, mr2.maxY),
		minX: Min(mr1.minX, mr2.minX),
		minY: Min(mr1.minY, mr2.minY),
	}
	return mixedMR
}

// GetWireMatrixRange ...
func GetWireMatrixRange(wire []Move) MatrixRange {
	currentPoint := MatrixDot{
		wire1Pass: false,
		wire2Pass: false,
	}
	currentMatrixRange := MatrixRange{0, 0, 0, 0}
	for i := range wire {
		switch wire[i].direction {
		case 1: // up
			currentPoint.relativeCoord.y += wire[i].distance
			currentMatrixRange.maxY = Max(currentMatrixRange.maxY, currentPoint.relativeCoord.y)
		case 2: // Right
			currentPoint.relativeCoord.x += wire[i].distance
			currentMatrixRange.maxX = Max(currentMatrixRange.maxX, currentPoint.relativeCoord.x)
		case 3: // Down
			currentPoint.relativeCoord.y -= wire[i].distance
			currentMatrixRange.minY = Min(currentMatrixRange.minY, currentPoint.relativeCoord.y)
		case 4: // Left
			currentPoint.relativeCoord.x -= wire[i].distance
			currentMatrixRange.minX = Min(currentMatrixRange.minX, currentPoint.relativeCoord.x)
		default:
			panic("invalid direction")
		}
	}
	return currentMatrixRange
}

// CreateMatrix ...
func (e Day3) CreateMatrix() [][]MatrixDot {
	matrix := make([][]MatrixDot, e.matrixRange.maxX-e.matrixRange.minX+1)
	for i := range matrix {
		matrix[i] = make([]MatrixDot, e.matrixRange.maxY-e.matrixRange.minY+1)
	}
	return matrix
}

// MatrixInitial ...
func (e *Day3) MatrixInitial() [][]MatrixDot {
	// create empty map in correct size
	matrix := e.CreateMatrix()
	e.matrixCenter.x = -e.matrixRange.minX
	e.matrixCenter.y = -e.matrixRange.minY
	// initialize each point in the map and calculate the distance to center
	e.matrix = matrix
	return matrix
}

// GetCoordinate ...
func (e *MatrixDot) GetCoordinate(c Coordinate) Coordinate {
	return Coordinate{
		x: e.relativeCoord.x + c.x,
		y: e.relativeCoord.y + c.y,
	}
}

// PassThroughMatrixDot ...
func PassThroughMatrixDot(fromPoint MatrixDot, currentPoint *MatrixDot, wireNumber int, distance int) {
	switch wireNumber {
	case 1:
		currentPoint.wire1Pass = true
		if currentPoint.wire1Steps == 0 {
			currentPoint.wire1Steps = fromPoint.wire1Steps + distance
		}
	case 2:
		currentPoint.wire2Pass = true
		if currentPoint.wire2Steps == 0 {
			currentPoint.wire2Steps = fromPoint.wire2Steps + distance
		}
	default:
		panic("invalid wire")
	}
}

// WireStep ...
func WireStep(fromPoint MatrixDot, move Move, wireNumber int, e *Day3) MatrixDot {
	toPoint := MatrixDot{}
	switch wireNumber {
	case 1:
		if fromPoint.wire1Steps < 0 {
			fromPoint.wire1Steps = 0
		}
	case 2:
		if fromPoint.wire2Steps < 0 {
			fromPoint.wire2Steps = 0
		}
	default:
		panic("invalid wire")
	}
	// get absolute coord for from point
	coord := fromPoint.GetCoordinate(e.matrixCenter)
	currentPoint := &e.matrix[coord.x][coord.y]
	switch move.direction {
	case 1: // Up
		for i := 0; i <= move.distance; i++ {
			currentPoint = &e.matrix[coord.x][coord.y+i]
			currentPoint.relativeCoord.y = coord.y + i - e.matrixCenter.y
			currentPoint.relativeCoord.x = coord.x - e.matrixCenter.x
			currentPoint.GetManhattanDistance()
			PassThroughMatrixDot(fromPoint, currentPoint, wireNumber, i)
		}
		toPoint = e.matrix[coord.x][coord.y+move.distance]
	case 2: // Right
		for i := 0; i <= move.distance; i++ {
			currentPoint = &e.matrix[coord.x+i][coord.y]
			currentPoint.relativeCoord.y = coord.y - e.matrixCenter.y
			currentPoint.relativeCoord.x = coord.x + i - e.matrixCenter.x
			currentPoint.GetManhattanDistance()
			PassThroughMatrixDot(fromPoint, currentPoint, wireNumber, i)
		}
		toPoint = e.matrix[coord.x+move.distance][coord.y]
	case 3: // Down
		for i := 0; i <= move.distance; i++ {
			currentPoint = &e.matrix[coord.x][coord.y-i]
			currentPoint.relativeCoord.y = coord.y - i - e.matrixCenter.y
			currentPoint.relativeCoord.x = coord.x - e.matrixCenter.x
			currentPoint.GetManhattanDistance()
			PassThroughMatrixDot(fromPoint, currentPoint, wireNumber, i)
		}
		toPoint = e.matrix[coord.x][coord.y-move.distance]
	case 4: // Left
		for i := 0; i <= move.distance; i++ {
			currentPoint = &e.matrix[coord.x-i][coord.y]
			currentPoint.relativeCoord.y = coord.y - e.matrixCenter.y
			currentPoint.relativeCoord.x = coord.x - i - e.matrixCenter.x
			currentPoint.GetManhattanDistance()
			PassThroughMatrixDot(fromPoint, currentPoint, wireNumber, i)
		}
		toPoint = e.matrix[coord.x-move.distance][coord.y]
	default:
		panic("invalid direction")
	}
	return toPoint
}

// RunWires ...
func (e *Day3) RunWires() {
	MatrixCenterPoint := e.matrix[e.matrixCenter.x][e.matrixCenter.y]

	//run wire 1 from matrix center
	currentPoint := MatrixCenterPoint
	for i := range e.wire1 {
		currentPoint = WireStep(currentPoint, e.wire1[i], 1, e)
	}
	MatrixCenterPoint.wire1Pass = true
	MatrixCenterPoint.wire1Steps = 0

	//run wire 2 from matrix center
	currentPoint = MatrixCenterPoint
	for i := range e.wire2 {
		currentPoint = WireStep(currentPoint, e.wire2[i], 2, e)
	}
	MatrixCenterPoint.wire2Pass = true
	MatrixCenterPoint.wire2Steps = 0
}

// GetManhattanDistance ...
func (e *MatrixDot) GetManhattanDistance() int {
	e.manhattanDistanceToCenter = (Max(e.relativeCoord.x, 0) - Min(e.relativeCoord.x, 0)) + (Max(e.relativeCoord.y, 0) - Min(e.relativeCoord.y, 0))
	return e.manhattanDistanceToCenter
}

// FindCrossPoints ...
func (e *Day3) FindCrossPoints() {
	//save all cross points in array
	e.crossPoints = []MatrixDot{}
	for m, a := range e.matrix {
		for n := range a {
			if a[n].wire1Pass && a[n].wire2Pass {
				// ignore center point
				if !reflect.DeepEqual(a[n].GetCoordinate(e.matrixCenter), e.matrixCenter) {
					//save all cross point into array
					e.crossPoints = append(e.crossPoints, a[n])
					fmt.Println("    cross point: " + strconv.Itoa(m) + "," + strconv.Itoa(n) +
						" x: " + strconv.Itoa(a[n].relativeCoord.x) +
						" y: " + strconv.Itoa(a[n].relativeCoord.y) +
						" manhattan distance: " + strconv.Itoa(a[n].manhattanDistanceToCenter) +
						" wire1 distance: " + strconv.Itoa(a[n].wire1Steps) +
						" wire2 distance: " + strconv.Itoa(a[n].wire2Steps))
				}
			}
		}
	}
}

// Part1 ...
func (e Day3) Part1() {
	fmt.Println("  Part1: ")
	// get shortest distance
	shortestDistance := e.crossPoints[0].manhattanDistanceToCenter
	for i := range e.crossPoints {
		if shortestDistance > e.crossPoints[i].manhattanDistanceToCenter {
			shortestDistance = e.crossPoints[i].manhattanDistanceToCenter
		}
	}
	fmt.Println("    **********************")
	fmt.Println("    shortest manhattan distance: " + strconv.Itoa(shortestDistance))
	fmt.Println("    **********************")
}

// Part2 ...
func (e Day3) Part2() {
	fmt.Println("  Part2: ")
	// get shortest wire distance
	shortestDistance := e.crossPoints[0].wire1Steps + e.crossPoints[0].wire2Steps
	for i := range e.crossPoints {
		if shortestDistance > (e.crossPoints[i].wire1Steps + e.crossPoints[i].wire2Steps) {
			shortestDistance = e.crossPoints[i].wire1Steps + e.crossPoints[i].wire2Steps
		}
	}

	fmt.Println("    **********************")
	fmt.Println("    shortest combined wire distance: " + strconv.Itoa(shortestDistance))
	fmt.Println("    **********************")
}

// Start ...
func Start() {
	fmt.Println("Day 3")
	fmt.Println("  initializing ...")
	e := Day3{
		wire1:  InputData{}.wire1Moves(),
		wire2:  InputData{}.wire2Moves(),
		matrix: [][]MatrixDot{},
		matrixRange: MixMatrixRange(
			GetWireMatrixRange(InputData{}.wire1Moves()),
			GetWireMatrixRange(InputData{}.wire2Moves())),
	}
	// create the map and initialize it
	e.MatrixInitial()
	// run all wires in the map and mark all point passed
	e.RunWires()
	// find all cross points
	e.FindCrossPoints()
	// do part 1
	e.Part1()
	// do part 2
	e.Part2()
}
