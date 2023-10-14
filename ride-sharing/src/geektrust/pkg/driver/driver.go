package driver

import "geektrust/pkg/location"

type Driver struct {
	id                 string
	location           *location.Location
	isAvailableForRide bool
}

func New(id string, loc *location.Location, isAvailableForRide bool) *Driver {
	return &Driver{
		id:                 id,
		location:           loc,
		isAvailableForRide: isAvailableForRide,
	}
}

// TODO: This is a setter. See if we can get rid of this setter
// and set this value as part of a "start ride" process / method
func (d *Driver) MarkAsUnavailableForRide() {
	d.isAvailableForRide = false
}

// TODO: This is a setter. See if we can get rid of this setter
// and set this value as part of a "stop ride" process / method
func (d *Driver) MarkAsAvailableForRide() {
	d.isAvailableForRide = true
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
