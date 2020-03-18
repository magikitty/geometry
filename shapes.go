package geometry

// Shape is an interface implemented by the shapes in this library
type Shape interface {
	Area() float64
	Perimeter() float64
}

// TotalArea returns total area of any number of shapes passed to it
func TotalArea(shapes ...Shape) float64 {
	var totalArea float64
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}

// TotalPerimeter returns total perimeter of any number of shapes passed to it
func TotalPerimeter(shapes ...Shape) float64 {
	var totalPerimeter float64
	for _, shape := range shapes {
		totalPerimeter += shape.Perimeter()
	}
	return totalPerimeter
}
