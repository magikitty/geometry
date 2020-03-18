package geometry

import "math"

// Circle struct
type Circle struct {
	radius float64
}

// Area of circle
func (circle Circle) Area() float64 {
	return math.Pi * circle.radius * circle.radius
}

// Perimeter (circumference) of circle
func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

// Diameter of circle
func (circle Circle) Diameter() float64 {
	return 2 * circle.radius
}

// MakeCircle returns a new circle
func MakeCircle(radius float64) Circle {
	return Circle{radius}
}
