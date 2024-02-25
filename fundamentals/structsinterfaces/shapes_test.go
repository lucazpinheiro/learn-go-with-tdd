package structsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{Width: 10.0, Height: 10.0}
	got := Perimeter(r)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// TestArea using table driven tests (https://go.dev/wiki/TableDrivenTests)
func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// wrap each test in an t.Run block to have better information when tests fails
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("got %#v want %#v", got, tt.hasArea)
			}
		})
	}
}

// TestArea first version
// func TestArea(t *testing.T) {
// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}
// 	t.Run("rectangle", func(t *testing.T) {
// 		r := Rectangle{Width: 12.0, Height: 6.0}
// 		want := 72.0

// 		checkArea(t, r, want)
// 	})

// 	t.Run("area", func(t *testing.T) {
// 		circle := Circle{10}
// 		want := 314.1592653589793

// 		checkArea(t, circle, want)
// 	})
// }
