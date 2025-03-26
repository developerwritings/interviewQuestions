package main

import "fmt"

// Circle struct represents a geometric circle with a radius.
type Circle struct {
	radius float32 // Field to store the radius of the circle.
}

// calculate() modifies the radius inside the method but does NOT change the original struct.
func (c Circle) calculate() float32 {
	c.radius = 3.11 // This modifies a COPY of the struct, NOT the original.
	return c.radius // Returns the modified value (3.11).
}

func main() {
	// Creating an instance of Circle with radius 3.15.
	c := Circle{radius: 3.15}

	// Calling the calculate() method.
	fmt.Println(c.calculate()) // Prints 3.11 because the function modifies a copy.

	// Printing the original radius in 'main()'.
	fmt.Println(c.radius) // Still prints 3.15 because the original struct remains unchanged.
}
