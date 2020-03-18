package geometry

import (
	"fmt"
	"testing"
)

func TestCircleArea(t *testing.T) {
	circle := MakeCircle(3.0)
	want := 28.274333882308138
	got := circle.Area()

	if got != want {
		t.Errorf("Area was %v, wanted %v.", got, want)
	} else {
		fmt.Printf("PASS: Area for %T \n", circle)
	}
}

func TestCircleDiameter(t *testing.T) {
	circle := MakeCircle(3.0)
	want := 6.0
	got := circle.Diameter()

	if got != want {
		t.Errorf("Diameter was %v, wanted %v.", got, want)
	} else {
		fmt.Printf("PASS: Diameter for %T \n", circle)
	}
}

func TestCirclePerimeter(t *testing.T) {
	circle := MakeCircle(3.0)
	want := 18.84955592153876
	got := circle.Perimeter()

	if got != want {
		t.Errorf("Perimeter was %v, wanted %v.", got, want)
	} else {
		fmt.Printf("PASS: Perimeter for %T \n", circle)
	}
}
