package geometry

import (
	"math"
)

// Rectangle struct
type Rectangle struct {
	point1 Point
	point2 Point
}

// Length of rectangle
func (rectangle Rectangle) Length() float64 {
	length := math.Abs(rectangle.point2.x) - math.Abs(rectangle.point1.x)
	return math.Abs(length)
}

// Width of rectangle
func (rectangle Rectangle) Width() float64 {
	width := math.Abs(rectangle.point2.y) - math.Abs(rectangle.point1.y)
	return math.Abs(width)
}

// Area of rectangle
func (rectangle Rectangle) Area() float64 {
	return rectangle.Length() * rectangle.Width()
}

// Perimeter of rectangle
func (rectangle Rectangle) Perimeter() float64 {
	return (rectangle.Length() * 2) + (rectangle.Width() * 2)
}

// Points returns rectangle's point fields
func (rectangle Rectangle) Points() (Point, Point) {
	return rectangle.point1, rectangle.point2
}

// MakeRectangle returns a new rectangle
func MakeRectangle(point1, point2 Point) Rectangle {
	return Rectangle{point1, point2}
}
