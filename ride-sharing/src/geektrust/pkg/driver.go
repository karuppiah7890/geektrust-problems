package pkg

type Driver struct {
	id                 string
	location           *Location
	isAvailableForRide bool
}

func NewDriver(id string, location *Location, isAvailableForRide bool) *Driver {
	return &Driver{
		id: id,
		location: location,
		isAvailableForRide: isAvailableForRide,
	}
}

func (d *Driver) GetLocation() *Location {
	return d.location
}

func (d *Driver) GetID() string {
	return d.id
}

func (d *Driver) IsAvailableForRide() bool {
	return d.isAvailableForRide
}
