package geometry

import (
	"fmt"
	"testing"
)

func TestRectangleLength(t *testing.T) {

	rectangle := MakeRectangle(MakePoint(0, 0), MakePoint(2, 2))
	want := 2.0
	got := rectangle.Length()

	if got != want {
		t.Errorf("Length was %v, wanted %v.", got, want)
	} else {
		fmt.Printf("PASS: Length for %T \n", rectangle)
	}
}

func TestRectangleWidth(t *testing.T) {
	rectangle := MakeRectangle(MakePoint(0, 0), MakePoint(2, 2))
	want := 2.0
	got := rectangle.Width()

	if got != want {
		t.Errorf("Width was %v, wanted %v.", got, want)
	} else {
		fmt.Printf("PASS: Width for %T \n", rectangle)
	}
}

func TestRectangleArea(t *testing.T) {
	rectangle := MakeRectangle(MakePoint(0, 0), MakePoint(2, 2))
	want := 4.0
	got := rectangle.Area()

	if got != want {
		t.Errorf("Area was %v, wanted %v.", got, want)
	} else {
		fmt.Printf("PASS: Area for %T \n", rectangle)
	}
}

func TestRectanglePerimeter(t *testing.T) {

	rectangle := MakeRectangle(MakePoint(0, 0), MakePoint(2, 2))
	want := 8.0
	got := rectangle.Perimeter()

	if got != want {
		t.Errorf("Perimeter was %v, wanted %v.", got, want)
	} else {
		fmt.Printf("PASS: Perimeter for %T \n", rectangle)
	}
}
