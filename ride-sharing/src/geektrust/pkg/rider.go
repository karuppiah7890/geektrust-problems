package pkg

import "geektrust/pkg/location"

type Rider struct {
	ID       string
	Location *location.Location
	isOnRide bool
}

func (r *Rider) IsOnRide() bool {
	return r.isOnRide
}
