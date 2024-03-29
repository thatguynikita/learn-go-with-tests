package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f while expected %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name string
		shape Shape
		want float64
	}{
		{name: "rectangle test", shape: Rectangle{Width: 12.0, Height: 6.0}, want: 72.0},
		{name: "circle test", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "triangle test", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v: got %g while expected %g", tt.shape, got, tt.want)
			}
		})
	}
}