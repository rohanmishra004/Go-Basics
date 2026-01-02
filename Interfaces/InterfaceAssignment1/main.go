/**
Write a program - that creates two custom struct called triangle and square

The square type should be a struct with a field caled sideLength of type float64

triangle - has fields height of typefloat64 and base of type float64

Both type have function getArea returns calculated area of square and triangle

Area of triangle 0.5*base*length
area of square sideLength * sideLength

Add a shape interface that defines a function called printArea


**/

package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type shape interface{ getArea() float64 }

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	t := triangle{base: 10, height: 12}
	s := square{sideLength: 5}
	printArea(t)
	printArea(s)
}
