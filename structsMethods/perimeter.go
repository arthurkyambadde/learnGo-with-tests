package perimeter

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func perimeter(r Rectangle) float64 {

	return 2 * (r.width + r.height)

}

func Area(r Rectangle) float64 {
	return (r.width * r.height)
}

func AreaCircle(r Circle) float64 {
	return math.Pi * r.radius * r.radius
}
