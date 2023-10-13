package pkg

type Location struct {
	x float64
	y float64
}

func NewLocation(x float64, y float64) *Location {
	return &Location{
		x: x,
		y: y,
	}
}

func (l *Location) GetX() float64 {
	return l.x
}

func (l *Location) GetY() float64 {
	return l.y
}

func (l *Location) Clone() *Location {
	return &Location{
		x: l.x,
		y: l.y,
	}
}

func (l *Location) Equals(anotherLocation *Location) bool {
	return l.x == anotherLocation.x && l.y == anotherLocation.y
}