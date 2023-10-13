package pkg

type Driver struct {
	id                 string
	Location           *Location
	isAvailableForRide bool
}

func NewDriver(id string, location *Location, isAvailableForRide bool) *Driver {
	return &Driver{
		id: id,
		Location: location,
		isAvailableForRide: isAvailableForRide,
	}
}

func (d *Driver) GetID() string {
	return d.id
}

func (d *Driver) IsAvailableForRide() bool {
	return d.isAvailableForRide
}
