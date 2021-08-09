package struct_method_interface

import "math"

type Circle struct {
	Radius float64
}

func (r Circle) Area() float64 {
	return math.Pi * r.Radius * r.Radius
}
