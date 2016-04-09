package util

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

/*
	Communication
*/

//ServerAddr address of the central server
const ServerAddr = "127.0.0.1:7070"

/*
	Struct Point
*/

//Point p(x,y) struct
type Point struct {
	X float64
	Y float64
}

//DistanceTo compute the distance from V1 to V2
func (v1 *Point) DistanceTo(v2 *Point) float64 {
	return math.Sqrt((v1.X-v2.X)*(v1.X-v2.X) + (v1.Y-v2.Y)*(v1.Y-v2.Y))
}

//ParseFloatCoordinates construct a Point using two strings
func ParseFloatCoordinates(strx string, stry string) *Point {
	x, err := strconv.ParseFloat(strx, 64)
	if err != nil {
		fmt.Println("x coordinate is not a valid float value")
		return nil
	}

	y, err := strconv.ParseFloat(stry, 64)
	if err != nil {
		fmt.Println("y coordinate is not a valid float value")
		return nil
	}
	result := Point{X: x, Y: y}
	return &result
}

/*
	Common Utils
*/

//CheckError just check and print error
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
