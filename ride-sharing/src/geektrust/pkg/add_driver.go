package pkg

import (
	"fmt"
	"geektrust/pkg/driver"
	"geektrust/pkg/location"
)

type AddDriverInput struct {
	DriverId string
	Location *location.Location
}

func (r *RideSharingApp) AddDriver(input *AddDriverInput) error {
	driverId := input.DriverId

	if _, ok := r.drivers[driverId]; ok {
		return fmt.Errorf("a driver with id %s already exists", driverId)
	}

	loc := input.Location.Clone()

	r.drivers[driverId] = driver.New(driverId, loc, true)

	return nil
}
