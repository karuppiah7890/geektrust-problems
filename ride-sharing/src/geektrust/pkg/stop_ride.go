package pkg

import (
	"fmt"
	"geektrust/pkg/location"
)

type StopRideInput struct {
	RideId             string
	Destination        *location.Location
	TimeTakenInMinutes int
}

func (r *RideSharingApp) StopRide(input *StopRideInput) error {
	// if ride id does not exist, return error
	ride, ok := r.GetRide(input.RideId)
	if !ok {
		return fmt.Errorf("a ride with id %s does not exist: %w", input.RideId, ErrRideIdNotExist)
	}

	// if ride has already been stopped, return error
	if ride.IsComplete() {
		return fmt.Errorf("ride with id %s is already stopped: %w", input.RideId, ErrRideStopped)
	}

	// complete the ride
	ride.Complete()

	driver, ok := r.GetDriver(ride.GetDriverId())
	if !ok {
		return fmt.Errorf("could not get driver with id %s: %w", ride.GetDriverId(), ErrDriverIdNotExist)
	}

	if driver == nil {
		// TODO: Should we just return an error here? or panic due to system error?
		panic(fmt.Sprintf("expected driver details to exist for driver with id %s but none was found", ride.GetDriverId()))
	}

	// What about a weird case where driver is already available?
	// This could happen due to some glitch in the system
	if driver.IsAvailableForRide() {
		// TODO: Should we just return an error here? or panic due to system error?
		panic(fmt.Sprintf("expected driver with id %s to not be available for a ride but they were already available for a ride", ride.GetDriverId()))
	}

	// make driver available for ride
	driver.MarkAsAvailableForRide()

	rider, ok := r.GetRider(ride.GetRiderId())
	if !ok {
		return fmt.Errorf("could not get rider with id %s: %w", ride.GetRiderId(), ErrRiderIdNotExist)
	}

	if rider == nil {
		// TODO: Should we just return an error here? or panic due to system error?
		panic(fmt.Sprintf("expected rider details to exist for rider with id %s but none was found", ride.GetRiderId()))
	}

	// What about a weird case where rider is already off the ride?
	// This could happen due to some glitch in the system
	if !rider.IsOnRide() {
		// TODO: Should we just return an error here? or panic due to system error?
		panic(fmt.Sprintf("expected rider with id %s to be on the ride but they were already off the ride", ride.GetRiderId()))
	}

	// rider is not on ride anymore
	rider.GetOffRide()

	return nil
}
