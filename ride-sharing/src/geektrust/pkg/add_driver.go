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

	r.drivers[driverId] = &Driver{
		ID: driverId,
		Location: &Location{
			X: input.Location.X,
			Y: input.Location.Y,
		},
		isAvailableForRide: true,
	}

	return nil
}
