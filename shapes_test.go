package geometry

import (
	"fmt"
	"testing"
)

func TestTotalArea(t *testing.T) {
	// test rectangle
	wantRectangle := 24.830000000000002
	rectangle1 := MakeRectangle(MakePoint(0, 0), MakePoint(2.1, 2.2))
	rectangle2 := MakeRectangle(MakePoint(0, 0), MakePoint(4.7, 4.3))
	gotRectangle := TotalArea(rectangle1, rectangle2)

	//test circle
	wantCircle := 15.707963267948966
	circle1 := MakeCircle(1)
	circle2 := MakeCircle(2)
	gotCircle := TotalArea(circle1, circle2)

	if gotRectangle != wantRectangle {
		t.Errorf("Area was %v, wanted %v.", gotRectangle, wantRectangle)
	} else {
		fmt.Printf("PASS: TotalArea for %T \n", rectangle1)
	}

	if gotCircle != wantCircle {
		t.Errorf("Area was %v, wanted %v.", gotCircle, wantCircle)
	} else {
		fmt.Printf("PASS: TotalArea for %T \n", circle1)
	}
}
