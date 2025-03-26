package main

import (
	"fmt"
	"math"
)

// Shape interface defines a method that all shapes must implement.
type Shape interface {
	calculate() float32 // Method that calculates area or volume.
}

// Circle struct represents a geometric circle with a radius.
type Circle struct {
	radius float32
}

// calculate() computes the area of a Circle: π * r².
func (c *Circle) calculate() float32 {
	return math.Pi * c.radius * c.radius
}

// Triangle struct represents a geometric triangle with a base and height.
type Triangle struct {
	base, height float32
}

// calculate() computes the area of a Triangle: (1/2) * base * height.
func (t *Triangle) calculate() float32 {
	return 0.5 * t.base * t.height
}

// Rectangle struct represents a geometric rectangle with width and height.
type Rectangle struct {
	width, height float32
}

// calculate() computes the area of a Rectangle: width * height.
func (r *Rectangle) calculate() float32 {
	return r.width * r.height
}

// Square struct represents a square with a side length.
type Square struct {
	side float32
}

// calculate() computes the area of a Square: side².
func (s *Square) calculate() float32 {
	return s.side * s.side
}

// printArea() function takes any Shape and returns its calculated value.
func printArea(s Shape) float32 {
	return s.calculate()
}

func main() {
	// Creating instances of different shapes.
	c := &Circle{radius: 3.15}
	t := &Triangle{base: 4, height: 5}
	r := &Rectangle{width: 6, height: 7}
	s := &Square{side: 4}

	// Printing areas for each shape.
	fmt.Println("Circle Area:", printArea(c))    // Expected: π * 3.15²
	fmt.Println("Triangle Area:", printArea(t))  // Expected: 0.5 * 4 * 5
	fmt.Println("Rectangle Area:", printArea(r)) // Expected: 6 * 7
	fmt.Println("Square Area:", printArea(s))    // Expected: 4²
}
