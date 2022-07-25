package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

type square struct {
	side float64
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	sq := square{
		side: 5.5,
	}

	printArea(sq)

	tri := triangle{
		base:   6,
		height: 6,
	}

	printArea(tri)
}
