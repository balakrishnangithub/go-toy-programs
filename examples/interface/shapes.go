package main

import "fmt"

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

type square struct {
	side float64
}

func (s square) getArea() float64 {
	return s.side * s.side
}

type triangle struct {
	base   float64
	height float64
}

func (t *triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func mainShapes() {
	shape1 := square{side: 10}
	shape2 := triangle{base: 10, height: 10}
	printArea(shape1)
	printArea(&shape2) // Note: only the pointer of triangle implements getArea
}
