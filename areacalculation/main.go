package main

import (
	"fmt"
)

type shape interface {
	printArea()
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	sq := square{sideLength: 2}

	fmt.Println(sq.getArea())

	tri := triangle{base: 3, height: 4}

	fmt.Println(tri.getArea())

	sq.printArea()

	tri.printArea()

}

func (t triangle) getArea() float64 {

	return 0.5 * t.base * t.height

}

func (s square) getArea() float64 {

	return s.sideLength * s.sideLength

}

func (t triangle) printArea() {

	fmt.Println(0.5 * t.base * t.height)

}

func (s square) printArea() {

	fmt.Println(s.sideLength * s.sideLength)

}
