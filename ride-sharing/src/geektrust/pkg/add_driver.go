package pkg

import "fmt"

type AddDriverInput struct {
	DriverId string
	Location *Location
}

func (r *RideSharingApp) AddDriver(input *AddDriverInput) error {
	driverId := input.DriverId

	if _, ok := r.drivers[driverId]; ok {
		return fmt.Errorf("a driver with id %s already exists", driverId)
	}

	location := input.Location.Clone()

	r.drivers[driverId] = NewDriver(driverId, location, true)

	return nil
}
