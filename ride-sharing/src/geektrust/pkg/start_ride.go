package pkg

import "fmt"

type StartRideInput struct {
	RiderId  *string
	DriverId *string
	RideId   *string
}

func (r *RideSharingApp) StartRide(input *StartRideInput) error {
	if _, ok := r.rides[*input.RideId]; ok {
		return RideIdAlreadyExists(fmt.Sprintf("a ride with id %s already exists", *input.RideId))
	}

	// TODO: Check if driver with that driver id exists. Return DriverNotFound custom error
	// TODO: Check if driver is available. Return DriverNotAvailable custom error

	r.rides[*input.RideId] = &Ride{
		id:         *input.RideId,
		isComplete: false,
		riderId:    *input.RiderId,
		driverId:   *input.DriverId,
	}

	// TODO: Set ride id of current ride for driver. This is to set that driver is unavailable for new rides
	// TODO: Set ride id of current ride for rider

	return nil
}
