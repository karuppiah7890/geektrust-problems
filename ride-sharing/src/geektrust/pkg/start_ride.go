package pkg

import (
	"fmt"
	"geektrust/pkg/ride"
)

type StartRideInput struct {
	RiderId  string
	DriverId string
	RideId   string
}

func (r *RideSharingApp) StartRide(input *StartRideInput) error {
	rideId := input.RideId
	driverId := input.DriverId
	riderId := input.RiderId

	if _, ok := r.GetRide(rideId); ok {
		return fmt.Errorf("a ride with id %s already exists: %w", rideId, ErrRideIdExist)
	}

	driver, ok := r.GetDriver(driverId)
	if !ok {
		return fmt.Errorf("could not get driver with id %s: %w", driverId, ErrDriverIdNotExist)
	}

	if driver == nil {
		panic(fmt.Sprintf("expected driver details to exist for driver with id %s but none was found", driverId))
	}

	if !driver.IsAvailableForRide() {
		return fmt.Errorf("driver with id %s is not available for a ride: %w", driverId, ErrDriverNotAvailable)
	}

	rider, ok := r.GetRider(riderId)
	if !ok {
		return fmt.Errorf("could not get rider with id %s: %w", riderId, ErrRiderIdNotExist)
	}

	if rider == nil {
		panic(fmt.Sprintf("expected rider details to exist for rider with id %s but none was found", riderId))
	}

	if rider.IsOnRide() {
		return fmt.Errorf("rider with id %s is already on a ride: %w", riderId, ErrRiderOnRide)
	}

	r.rides[rideId] = ride.New(rideId, riderId, driverId, rider.GetLocation())

	driver.MarkAsUnavailableForRide()
	rider.GetOnRide()

	return nil
}
