package pkg_test

import (
	"errors"
	"geektrust/pkg"
	"geektrust/pkg/driver"
	"testing"
)

// ride started case

// invalid ride case:
// 1. ride id already exists
// 2. driver id does not exist
// 3. driver is not available
// 4. rider id does not exist
// 5. rider is already on ride

func TestStartRide(t *testing.T) {
	t.Run("start a valid ride", func(t *testing.T) {
		rideSharingApp := pkg.NewRideSharingApp()
		addDrivers(t, rideSharingApp, []*driver.Driver{
			newDriver("D1", 1, 1),
			newDriver("D2", 4, 5),
			newDriver("D3", 2, 2),
		})
		addRiders(t, rideSharingApp, []*pkg.Rider{newRider("R1", 0, 0)})

		rideId := "RIDE-001"
		riderId := "R1"
		driverId := "D3"

		input := &pkg.StartRideInput{
			RideId:   rideId,
			RiderId:  riderId,
			DriverId: driverId,
		}
		err := rideSharingApp.StartRide(input)
		if err != nil {
			t.Errorf("expected no error occur while starting ride with given ride id, rider and driver but got error: %v", err)
		}

		ride, ok := rideSharingApp.GetRide(rideId)
		if !ok {
			t.Errorf("expected to get ride with given ride id %v, but got none", rideId)
		}

		assertStringEqual(t, ride.GetId(), rideId)
		assertStringEqual(t, ride.GetRiderId(), riderId)
		assertStringEqual(t, ride.GetDriverId(), driverId)
		assertBoolEqual(t, ride.IsComplete(), false)

		driver, ok := rideSharingApp.GetDriver(driverId)
		if !ok {
			t.Errorf("expected to get driver with id %s but got none", driverId)
		}

		assertBoolEqual(t, driver.IsAvailableForRide(), false)

		rider, ok := rideSharingApp.GetRider(riderId)
		if !ok {
			t.Errorf("expected to get rider with id %s but got none", riderId)
		}

		assertBoolEqual(t, rider.IsOnRide(), true)
	})
	t.Run("invalid ride cases", func(t *testing.T) {
		t.Run("ride id already exists", func(t *testing.T) {
			rideSharingApp := pkg.NewRideSharingApp()
			addDrivers(t, rideSharingApp, []*driver.Driver{
				newDriver("D1", 1, 1),
				newDriver("D2", 4, 5),
				newDriver("D3", 2, 2),
			})
			addRiders(t, rideSharingApp, []*pkg.Rider{
				newRider("R1", 0, 0),
				newRider("R2", 0, 0),
			})

			rideId := "RIDE-001"
			riderId := "R1"
			driverId := "D3"

			input := &pkg.StartRideInput{
				RideId:   rideId,
				RiderId:  riderId,
				DriverId: driverId,
			}

			err := rideSharingApp.StartRide(input)
			if err != nil {
				t.Errorf("expected no error occur while starting ride with given ride id, rider id and driver id (%v) but got error: %v", input, err)
			}

			anotherRiderId := "R2"
			anotherDriverId := "D1"

			anotherStartRideInput := &pkg.StartRideInput{
				RideId:   rideId,
				RiderId:  anotherRiderId,
				DriverId: anotherDriverId,
			}

			err = rideSharingApp.StartRide(anotherStartRideInput)
			if err == nil {
				t.Errorf("expected error to occur while starting ride with already existing ride id %v but got none", rideId)
			}

			if !errors.Is(err, pkg.ErrRideIdExist) {
				t.Errorf("expected the error to be something but got something else. Actual: %v. Expected: %v", err, pkg.ErrRideIdExist)
			}
		})

		t.Run("driver id does not exist", func(t *testing.T) {
			rideSharingApp := pkg.NewRideSharingApp()
			addDrivers(t, rideSharingApp, []*driver.Driver{
				newDriver("D1", 1, 1),
				newDriver("D2", 4, 5),
				newDriver("D3", 2, 2),
			})
			addRiders(t, rideSharingApp, []*pkg.Rider{
				newRider("R1", 0, 0),
				newRider("R2", 0, 0),
			})

			rideId := "RIDE-001"
			riderId := "R1"
			driverId := "D4"

			input := &pkg.StartRideInput{
				RideId:   rideId,
				RiderId:  riderId,
				DriverId: driverId,
			}

			err := rideSharingApp.StartRide(input)
			if err == nil {
				t.Errorf("expected error to occur while starting ride with driver id %v that does not exist but got none", driverId)
			}

			if !errors.Is(err, pkg.ErrDriverIdNotExist) {
				t.Errorf("expected the error to be something but got something else. Actual: %v. Expected: %v", err, pkg.ErrDriverIdNotExist)
			}
		})

		t.Run("driver is not available", func(t *testing.T) {
			rideSharingApp := pkg.NewRideSharingApp()
			addDrivers(t, rideSharingApp, []*driver.Driver{
				newDriver("D1", 1, 1),
				newDriver("D2", 4, 5),
				newDriver("D3", 2, 2),
			})
			addRiders(t, rideSharingApp, []*pkg.Rider{
				newRider("R1", 0, 0),
				newRider("R2", 0, 0),
			})

			rideId := "RIDE-001"
			riderId := "R1"
			driverId := "D3"

			input := &pkg.StartRideInput{
				RideId:   rideId,
				RiderId:  riderId,
				DriverId: driverId,
			}

			err := rideSharingApp.StartRide(input)
			if err != nil {
				t.Errorf("expected no error occur while starting ride with given ride id, rider id and driver id (%v) but got error: %v", input, err)
			}

			anotherRideId := "RIDE-002"
			anotherRiderId := "R2"

			anotherStartRideInput := &pkg.StartRideInput{
				RideId:   anotherRideId,
				RiderId:  anotherRiderId,
				DriverId: driverId,
			}

			err = rideSharingApp.StartRide(anotherStartRideInput)
			if err == nil {
				t.Errorf("expected error to occur while starting ride with driver (id: %v) who is not available but got none", driverId)
			}

			if !errors.Is(err, pkg.ErrDriverNotAvailable) {
				t.Errorf("expected the error to be something but got something else. Actual: %v. Expected: %v", err, pkg.ErrDriverNotAvailable)
			}
		})

		t.Run("rider id does not exist", func(t *testing.T) {
			rideSharingApp := pkg.NewRideSharingApp()
			addDrivers(t, rideSharingApp, []*driver.Driver{
				newDriver("D1", 1, 1),
				newDriver("D2", 4, 5),
				newDriver("D3", 2, 2),
			})
			addRiders(t, rideSharingApp, []*pkg.Rider{
				newRider("R1", 0, 0),
				newRider("R2", 0, 0),
			})

			rideId := "RIDE-001"
			riderId := "R3"
			driverId := "D3"

			input := &pkg.StartRideInput{
				RideId:   rideId,
				RiderId:  riderId,
				DriverId: driverId,
			}

			err := rideSharingApp.StartRide(input)
			if err == nil {
				t.Errorf("expected error to occur while starting ride with driver id %v that does not exist but got none", driverId)
			}

			if !errors.Is(err, pkg.ErrRiderIdNotExist) {
				t.Errorf("expected the error to be something but got something else. Actual: %v. Expected: %v", err, pkg.ErrRiderIdNotExist)
			}
		})

		t.Run("rider is already on ride", func(t *testing.T) {
			rideSharingApp := pkg.NewRideSharingApp()
			addDrivers(t, rideSharingApp, []*driver.Driver{
				newDriver("D1", 1, 1),
				newDriver("D2", 4, 5),
				newDriver("D3", 2, 2),
			})
			addRiders(t, rideSharingApp, []*pkg.Rider{
				newRider("R1", 0, 0),
				newRider("R2", 0, 0),
			})

			rideId := "RIDE-001"
			riderId := "R1"
			driverId := "D3"

			input := &pkg.StartRideInput{
				RideId:   rideId,
				RiderId:  riderId,
				DriverId: driverId,
			}

			err := rideSharingApp.StartRide(input)
			if err != nil {
				t.Errorf("expected no error occur while starting ride with given ride id, rider id and driver id (%v) but got error: %v", input, err)
			}

			anotherRideId := "RIDE-002"
			anotherDriverId := "D1"

			anotherStartRideInput := &pkg.StartRideInput{
				RideId:   anotherRideId,
				RiderId:  riderId,
				DriverId: anotherDriverId,
			}

			err = rideSharingApp.StartRide(anotherStartRideInput)
			if err == nil {
				t.Errorf("expected error to occur while starting ride with rider (id: %v) who is already on a ride, but got none", riderId)
			}

			if !errors.Is(err, pkg.ErrRiderOnRide) {
				t.Errorf("expected the error to be something but got something else. Actual: %v. Expected: %v", err, pkg.ErrRiderOnRide)
			}
		})
	})
}
