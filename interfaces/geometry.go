package interfaces

// Implement an interface for a TwoDShape
// This is really a proof of concept on how to "mark" or "group" a function into an Interface
// now on how to implement a geometry library, so it has a lot of room for improvement

import "math"

type TwoDShape interface {
	Perimeter() float64
	Area() float64
}

func getPerimeterAndArea(shape TwoDShape) (float64, float64) {
	return shape.Perimeter(), shape.Area()
}

// A receiver of type Rectangle would satisfy
// or implement TwoDShape
type Rectangle struct {
	sides []float64
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.sides[0] + 2*r.sides[1]
}
func (r Rectangle) Area() float64 {
	return r.sides[0] * r.sides[1]
}

// A receiver of type Square would also satisfy
// or implement TwoDShape
type Square struct {
	side float64
}

func (s Square) Perimeter() float64 {
	return 4 * s.side
}
func (s Square) Area() float64 {
	return math.Pow(s.side, 2)
}
