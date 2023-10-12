package pkg

type Driver struct {
	ID                 string
	Location           *Location
	isAvailableForRide bool
}

func (d *Driver) IsAvailableForRide() bool {
	return d.isAvailableForRide
}
