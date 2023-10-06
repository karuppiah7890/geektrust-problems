package pkg

import (
	"fmt"
)

type StartRideInput struct {
	RiderId  *string
	DriverId *string
	RideId   *string
}

func (r *RideSharingApp) StartRide(input *StartRideInput) error {
	if _, ok := r.GetRide(*input.RideId); ok {
		return fmt.Errorf("a ride with id %s already exists: %w", *input.RideId, ErrRideIdExist)
	}

	driver, ok := r.GetDriver(*input.DriverId)
	if !ok {
		return fmt.Errorf("could not get driver with id %s: %w", *input.DriverId, ErrDriverIdNotExist)
	}

	if driver == nil {
		panic(fmt.Sprintf("expected driver details to exist for driver with id %s but none was found", *input.DriverId))
	}

	if !driver.isAvailableForRide {
		return fmt.Errorf("driver with id %s is not available for a ride: %w", *input.DriverId, ErrDriverNotAvailable)
	}

	rider, ok := r.GetRider(*input.RiderId)
	if !ok {
		return fmt.Errorf("could not get rider with id %s: %w", *input.RiderId, ErrRiderIdNotExist)
	}

	if rider == nil {
		panic(fmt.Sprintf("expected rider details to exist for rider with id %s but none was found", *input.RiderId))
	}

	if rider.isOnRide {
		return fmt.Errorf("rider with id %s is already on a ride: %w", *input.RiderId, ErrRiderOnRide)
	}

	r.rides[*input.RideId] = &Ride{
		id:         *input.RideId,
		isComplete: false,
		riderId:    *input.RiderId,
		driverId:   *input.DriverId,
	}

	driver.isAvailableForRide = false
	rider.isOnRide = true

	return nil
}
