package calcArea

import "math"

type Rectangle struct {
	height float64
	length float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.length
}

func (r Circle) Area() float64 {
	return math.Pi * r.radius * r.radius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
