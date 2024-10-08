package structs_test

import (
	"testing"

	"github.com/jlinke1/learn-go-with-tests/structs"
)

func TestPerimeter(t *testing.T) {
	rectangle := structs.Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   structs.Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: structs.Rectangle{Width: 10.0, Height: 10.0}, hasArea: 100.0},
		{name: "Circle", shape: structs.Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: structs.Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}

}
