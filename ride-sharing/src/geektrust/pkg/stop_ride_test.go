package pkg_test

import (
	"errors"
	"geektrust/pkg"
	"testing"
)

func TestStopRide(t *testing.T) {
	t.Run("stop a valid ride", func(t *testing.T) {
		rideSharingApp := pkg.NewRideSharingApp()
		addDrivers(t, rideSharingApp, []*pkg.Driver{
			driver("D1", 1, 1),
			driver("D2", 4, 5),
			driver("D3", 2, 2),
		})
		addRiders(t, rideSharingApp, []*pkg.Rider{rider("R1", 0, 0)})

		rideId := "RIDE-001"
		riderId := "R1"
		driverId := "D3"

		input := &pkg.StartRideInput{
			RideId:   &rideId,
			RiderId:  &riderId,
			DriverId: &driverId,
		}

		err := rideSharingApp.StartRide(input)
		if err != nil {
			t.Errorf("expected no error occur while starting ride with given ride id, rider and driver but got error: %v", err)
		}

		stopRideInput := &pkg.StopRideInput{
			RideId: rideId,
			Destination: &pkg.Location{
				X: 4,
				Y: 5,
			},
			TimeTakenInMinutes: 32,
		}

		err = rideSharingApp.StopRide(stopRideInput)
		if err != nil {
			t.Errorf("expected no error occur while stopping ride with given ride id (%v) but got error: %v", rideId, err)
		}

		ride, ok := rideSharingApp.GetRide(rideId)
		if !ok {
			t.Errorf("expected to get ride with given ride id %v, but got none", rideId)
		}

		assertStringEqual(t, ride.GetId(), rideId)
		assertStringEqual(t, ride.GetRiderId(), riderId)
		assertStringEqual(t, ride.GetDriverId(), driverId)
		assertBoolEqual(t, ride.IsComplete(), true)

		driver, ok := rideSharingApp.GetDriver(driverId)
		if !ok {
			t.Errorf("expected to get driver with id %s but got none", driverId)
		}

		assertBoolEqual(t, driver.IsAvailableForRide(), true)

		rider, ok := rideSharingApp.GetRider(riderId)
		if !ok {
			t.Errorf("expected to get rider with id %s but got none", riderId)
		}

		assertBoolEqual(t, rider.IsOnRide(), false)
	})

	t.Run("invalid ride cases", func(t *testing.T) {
		t.Run("ride id does not exist", func(t *testing.T) {
			rideSharingApp := pkg.NewRideSharingApp()

			rideId := "RIDE-999"

			input := &pkg.StopRideInput{
				RideId: rideId,
				Destination: &pkg.Location{
					X: 40,
					Y: 50,
				},
				TimeTakenInMinutes: 25,
			}

			err := rideSharingApp.StopRide(input)

			if err == nil {
				t.Errorf("expected error to occur while stopping a ride with a ride id %v that does not exist, but got none", rideId)
			}

			if !errors.Is(err, pkg.ErrRideIdNotExist) {
				t.Errorf("expected the error to be something but got something else. Actual: %v. Expected: %v", err, pkg.ErrRideIdNotExist)
			}
		})

		t.Run("ride has already been stopped", func(t *testing.T) {
			rideSharingApp := pkg.NewRideSharingApp()
			addDrivers(t, rideSharingApp, []*pkg.Driver{
				driver("D1", 1, 1),
				driver("D2", 4, 5),
				driver("D3", 2, 2),
			})
			addRiders(t, rideSharingApp, []*pkg.Rider{rider("R1", 0, 0)})

			rideId := "RIDE-001"
			riderId := "R1"
			driverId := "D3"

			input := &pkg.StartRideInput{
				RideId:   &rideId,
				RiderId:  &riderId,
				DriverId: &driverId,
			}

			err := rideSharingApp.StartRide(input)
			if err != nil {
				t.Errorf("expected no error occur while starting ride with given ride id, rider and driver but got error: %v", err)
			}

			stopRideInput := &pkg.StopRideInput{
				RideId: rideId,
				Destination: &pkg.Location{
					X: 4,
					Y: 5,
				},
				TimeTakenInMinutes: 32,
			}

			err = rideSharingApp.StopRide(stopRideInput)
			if err != nil {
				t.Errorf("expected no error occur while stopping ride with given ride id (%v) but got error: %v", rideId, err)
			}

			err = rideSharingApp.StopRide(stopRideInput)
			if err == nil {
				t.Errorf("expected error to occur while stopping ride with id (%v) as the ride has already been stopped but got no error", rideId)
			}

			if !errors.Is(err, pkg.ErrRideStopped) {
				t.Errorf("expected the error to be something but got something else. Actual: %v. Expected: %v", err, pkg.ErrRideStopped)
			}
		})
	})
}
