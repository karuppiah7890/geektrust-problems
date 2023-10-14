package location

import "math"

type Location struct {
	x float64
	y float64
}

func New(x float64, y float64) *Location {
	return &Location{
		x: x,
		y: y,
	}
}

func (loc *Location) GetX() float64 {
	return loc.x
}

func (loc *Location) GetY() float64 {
	return loc.y
}

func (loc *Location) Equals(anotherLoc *Location) bool {
	return loc.x == anotherLoc.x && loc.y == anotherLoc.y
}

func (loc *Location) DistanceBetween(anotherLoc *Location) float64 {
	// Euclidean distance formula: SquareRoot( (x2 - x1)^2 + (y2 - y1)^2 )
	return squareRoot(square(anotherLoc.GetX()-loc.GetX()) + square(anotherLoc.GetY()-loc.GetY()))
}

func square(x float64) float64 {
	return math.Pow(x, 2)
}

func squareRoot(x float64) float64 {
	return math.Sqrt(x)
}
