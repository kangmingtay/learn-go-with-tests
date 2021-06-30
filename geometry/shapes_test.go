package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12, 6}
	got := Area(rectangle)
	var want float64 = 72

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
