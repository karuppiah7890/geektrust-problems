package rider

import "geektrust/pkg/location"

type Rider struct {
	id       string
	location *location.Location
	isOnRide bool
}

func NewRider(id string, loc *location.Location, isOnRide bool) *Rider {
	return &Rider{
		id:       id,
		location: loc,
		isOnRide: isOnRide,
	}
}

func (r *Rider) GetID() string {
	return r.id
}

func (r *Rider) GetLocation() *location.Location {
	return r.location
}

func (r *Rider) IsOnRide() bool {
	return r.isOnRide
}

// TODO: This is a setter. See if we can get rid of this setter
// and set this value as part of a "start ride" process / method
func (r *Rider) GetOnRide() {
	r.isOnRide = true
}

// TODO: This is a setter. See if we can get rid of this setter
// and set this value as part of a "stop ride" process / method
func (r *Rider) GetOffRide() {
	r.isOnRide = false
}
