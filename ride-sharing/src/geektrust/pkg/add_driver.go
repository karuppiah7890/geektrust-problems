package pkg

import "fmt"

type AddDriverInput struct {
	DriverId *string
	Location *Location
}

func (r *RideSharingApp) AddDriver(input *AddDriverInput) error {
	if _, ok := r.drivers[*input.DriverId]; ok {
		return fmt.Errorf("a driver with id %s already exists", *input.DriverId)
	}

	r.drivers[*input.DriverId] = &Driver{
		ID: *input.DriverId,
		Location: &Location{
			X: input.Location.X,
			Y: input.Location.Y,
		},
	}

	return nil
}
