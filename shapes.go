package geometry

// Shapes is an interface implemented by the shapes in this library
type Shapes interface {
	Area() float64
	Perimiter() float64
}

// Returns total area of any number of shapes passed to it
func totalArea(shapes ...Shapes) float64 {
	var totalArea float64
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}

// Returns total perimeter of any number of shapes passed to it
func totalPerimeter(shapes ...Shapes) float64 {
	var totalPerimeter float64
	for _, shape := range shapes {
		totalPerimeter += shape.Perimiter()
	}
	return totalPerimeter
}
