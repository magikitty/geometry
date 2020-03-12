package geometry

// Point struct
type Point struct {
	x, y float64
}

// MakePoint returns a new point
func MakePoint(x, y float64) Point {
	return Point{x: x, y: y}
}
