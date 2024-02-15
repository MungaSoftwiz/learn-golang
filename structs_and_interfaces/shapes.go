package main

import "math"

//struct just like in C
type Rectangle struct {
	Width float64
	Height float64
}
type Circle struct {
	Radius float64
}
type Triangle struct {
	Base float64
	Height float64
}

//interface declaration.Specifies a set of method signatures
//provides a way to achieve polymorphism
type Shape interface {
	Area() float64
}


func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

//func Area(rectangle Rectangle) float64 {
//	return rectangle.Width * rectangle.Height
//}


//methods follow below:implementing area method for rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//convention in go to have receiver name "c" name as first
//letter type "(receiverName ReceiverType) MethodName(args)"
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.Base * t.Height)
}
