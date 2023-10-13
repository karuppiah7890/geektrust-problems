package pkg

import "geektrust/pkg/location"

type Driver struct {
	id                 string
	location           *location.Location
	isAvailableForRide bool
}

func NewDriver(id string, location *location.Location, isAvailableForRide bool) *Driver {
	return &Driver{
		id:                 id,
		location:           location,
		isAvailableForRide: isAvailableForRide,
	}
}

func (d *Driver) GetLocation() *location.Location {
	return d.location
}

func (d *Driver) GetID() string {
	return d.id
}

func (d *Driver) IsAvailableForRide() bool {
	return d.isAvailableForRide
}
