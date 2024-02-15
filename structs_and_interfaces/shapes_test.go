package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

//func TestArea(t *testing.T) {

//	checkArea := func(t testing.TB, shape Shape, want float64) {
//		t.Helper()
//		got := shape.Area() //access meth via interface
		//use of %g will give more precise decimal in                 //the error message
//		if got != want {
//			t.Errorf("got %g want %g", got, want)
//		}
//	}

//	t.Run("rectangles", func(t *testing.T) {
//		rectangle := Rectangle{10.0, 10.0}
//		checkArea(t, rectangle, 100.0)
//	})

//	t.Run("circle", func(t *testing.T) {
//		circle := Circle{10}
//		checkArea(t, circle, 314.1592653589793)
//	})
//}

func TestArea(t *testing.T) {
	//table driven tests
	areaTests := []struct {
		name string
		shape Shape //pass in interface
		hasArea float64 //the return type
	}{
		{"Rectangle", Rectangle{10, 10}, 100.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	//(_ is used as index & tt assigned to curr val)
	for _, tt := range areaTests {
		//using tt.name as t.Run test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
