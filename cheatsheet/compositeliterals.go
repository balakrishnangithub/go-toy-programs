// Composite literals are literals constructed combining different types
// ref: https://golang.org/ref/spec#Composite_literals
package main

import (
	"fmt"
)

func mainCompositeLiterals() {
	type Point3D struct{ x, y, z float64 }
	type Line struct{ p, q Point3D }
	origin := Point3D{}
	fmt.Printf("%+v\n", origin)                                //    "{x:0 y:0 z:0}"
	fmt.Printf("%+v\n", Line{origin, Point3D{y: -4, z: 12.3}}) // "{p:{x:0 y:0 z:0} q:{x:0 y:-4 z:12.3}}"

	p1 := &[]int{}                      // p1 points to an initialized, empty slice with value []int{} and length 0
	p2 := new([]int)                    // p2 points to an uninitialized slice with value nil and length 0
	fmt.Println(*p1 == nil, *p2 == nil) // "false true"

	fmt.Println(len([10]string{}))              // "10"
	fmt.Println(len([6]int{1, 2, 3, 5}))        // "6"
	fmt.Println(len([...]string{"Sat", "Sun"})) // "2"

	type Point struct{ x, y float64 }
	fmt.Printf("%+v\n", [...]Point{{1.5, -3.5}, {0, 0}})  // same as [...]Point{Point{1.5, -3.5}, Point{0, 0}}
	fmt.Printf("%+v\n", [][]int{{1, 2, 3}, {4, 5}})       // same as [][]int{[]int{1, 2, 3}, []int{4, 5}}
	fmt.Printf("%+v\n", [][]Point{{{0, 1}, {1, 2}}})      // same as [][]Point{[]Point{Point{0, 1}, Point{1, 2}}}
	fmt.Printf("%+v\n", map[string]Point{"orig": {0, 0}}) // same as map[string]Point{"orig": Point{0, 0}}
	fmt.Printf("%+v\n", map[Point]string{{0, 0}: "orig"}) // same as map[Point]string{Point{0, 0}: "orig"}

	type PPoint *Point
	fmt.Printf("%+v\n", [2]*Point{{1.5, -3.5}, {}}) // same as [2]*Point{&Point{1.5, -3.5}, &Point{}}
	fmt.Printf("%+v\n", [2]PPoint{{1.5, -3.5}, {}}) // same as [2]PPoint{PPoint(&Point{1.5, -3.5}), PPoint(&Point{})}

	// list of prime numbers
	primes := []int{2, 3, 5, 7, 9, 2147483647}
	fmt.Printf("%+v\n", primes) // "[2 3 5 7 9 2147483647]"

	// vowels[ch] is true if ch is a vowel
	vowels := [128]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true}
	fmt.Printf("%+v\n", vowels) // "[false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false false true false false false true false false false true false false false false false true false false false false false true false false false true false false false false false false]"

	// the array [10]float32{-1, 0, 0, 0, -0.1, -0.1, 0, 0, 0, -1}
	filter := [10]float32{-1, 4: -0.1, -0.1, 9: -1}
	fmt.Printf("%+v\n", filter) // "[-1 0 0 0 -0.1 -0.1 0 0 0 -1]"

	// frequencies in Hz for equal-tempered scale (A4 = 440Hz)
	noteFrequency := map[string]float32{
		"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
		"G0": 24.50, "A0": 27.50, "B0": 30.87,
	}
	fmt.Printf("%+v\n", noteFrequency) // "map[A0:27.5 B0:30.87 C0:16.35 D0:18.35 E0:20.6 F0:21.83 G0:24.5]"
}
