package main

import (
	"fmt"
)

// Vertex ...
type Vertex struct{ X, Y float64 }

// Mutation is possible when receiver is a pointer (*Vertex).
func (v *Vertex) reset() {
	v.X = 0
	v.Y = 0
}

// Everything in Go is passed by value, slices too.
// But slice value only contains a pointer to the array.
func getSliceAddr(s []int) string {
	return fmt.Sprintf("%p", s)
}

// Arrays are passed by value.
func getArrayAddr(a [7]int) string {
	return fmt.Sprintf("%p", &a)
}

func main() {
	var x, y int
	fmt.Println(x == 0, x == y)  // "true true"
	fmt.Println(&x == &y)        // "false"
	fmt.Printf("%T %T\n", x, &x) // "int *int"
	fmt.Printf("%p\n", x)        // "%!p(int=0)" // golint: format %p has arg x of wrong type int
	fmt.Printf("%p\n", &x)       // "0xc0000b6010"

	var p *int
	fmt.Println(p == nil) // "true"

	p = &x
	fmt.Println(p == &x, x == *p) // "true true"
	fmt.Printf("%T %T\n", p, &p)  // "*int **int"

	v := Vertex{X: 1, Y: 2}
	vptr := &v
	fmt.Println(vptr.X == (*vptr).X)       // "true"
	fmt.Println(&(vptr.X) == &((*vptr).X)) // "true"

	fmt.Println(v.X == 0, v.X == v.Y) // "false false"
	v.reset()
	fmt.Println(v.X == 0, v.X == v.Y) // "true true"

	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(fmt.Sprintf("%p", s) == getSliceAddr(s)) // "true"

	a := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(fmt.Sprintf("%p", &a) == getArrayAddr(a)) // "false"
}
