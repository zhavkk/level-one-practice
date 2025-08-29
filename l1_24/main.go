package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	pointA := NewPoint(4, 3)
	pointB := NewPoint(2, 1)

	fmt.Println(pointA.Distance(pointB))

	fmt.Println(pointB.Distance(pointA))

}
