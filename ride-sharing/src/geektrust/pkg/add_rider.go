package pkg

import (
	"fmt"
	"geektrust/pkg/location"
	"geektrust/pkg/rider"
)

type AddRiderInput struct {
	RiderId  string
	Location *location.Location
}

func (r *RideSharingApp) AddRider(input *AddRiderInput) error {
	riderId := input.RiderId

	if _, ok := r.riders[riderId]; ok {
		return fmt.Errorf("a rider with id %s already exists", riderId)
	}

	r.riders[riderId] = rider.New(riderId, input.Location, false)

	return nil
}
