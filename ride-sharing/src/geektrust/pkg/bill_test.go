package pkg_test

import (
	"errors"
	"geektrust/pkg"
	"geektrust/pkg/driver"
	"geektrust/pkg/location"
	"geektrust/pkg/rider"
	"testing"
)

func TestCalculateBill(t *testing.T) {
	t.Run("stop a valid ride", func(t *testing.T) {
		rideSharingApp := pkg.NewRideSharingApp()
		addDrivers(t, rideSharingApp, []*driver.Driver{
			newDriver("D1", 1, 1),
			newDriver("D2", 4, 5),
			newDriver("D3", 2, 2),
		})
		addRiders(t, rideSharingApp, []*rider.Rider{newRider("R1", 0, 0)})

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

		destination := location.New(4, 5)
		rideDurationInMinutes := 32

		stopRideInput := &pkg.StopRideInput{
			RideId:             rideId,
			Destination:        destination,
			TimeTakenInMinutes: rideDurationInMinutes,
		}

		err = rideSharingApp.StopRide(stopRideInput)
		if err != nil {
			t.Errorf("expected no error occur while stopping ride with given ride id (%v) but got error: %v", rideId, err)
		}

		calculateBillInput := &pkg.CalculateBillInput{
			RideId: rideId,
		}

		bill, err := rideSharingApp.CalculateBill(calculateBillInput)
		if err != nil {
			t.Errorf("expected no error occur while calculating the bill for ride with given ride id (%v) but got error: %v", rideId, err)
		}

		assertFloatEqual(t, bill.Amount, 186.74)
		assertStringEqual(t, bill.DriverId, driverId)
	})
	t.Run("failure cases", func(t *testing.T) {
		t.Run("ride id does not exist", func(t *testing.T) {
			rideSharingApp := pkg.NewRideSharingApp()

			rideId := "RIDE-999"

			input := &pkg.CalculateBillInput{
				RideId: rideId,
			}

			bill, err := rideSharingApp.CalculateBill(input)

			if err == nil {
				t.Errorf("expected error to occur while calculating the bill of a ride with ride id %v that does not exist, but got none", rideId)
			}

			if !errors.Is(err, pkg.ErrRideIdNotExist) {
				t.Errorf("expected the error to be something but got something else. Actual: %v. Expected: %v", err, pkg.ErrRideIdNotExist)
			}

			if bill != nil {
				t.Errorf("expected no bill for an invalid ride with ride id that does not exist, but got bill: %v", bill)
			}
		})

		t.Run("ride is not completed", func(t *testing.T) {
			rideSharingApp := pkg.NewRideSharingApp()
			addDrivers(t, rideSharingApp, []*driver.Driver{
				newDriver("D1", 1, 1),
				newDriver("D2", 4, 5),
				newDriver("D3", 2, 2),
			})
			addRiders(t, rideSharingApp, []*rider.Rider{newRider("R1", 0, 0)})

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

			calculateBillInput := &pkg.CalculateBillInput{
				RideId: rideId,
			}

			bill, err := rideSharingApp.CalculateBill(calculateBillInput)

			if err == nil {
				t.Errorf("expected error to occur while calculating the bill of a ride with ride id %v that hasn't been completed, but got none", rideId)
			}

			if !errors.Is(err, pkg.ErrRideNotCompleted) {
				t.Errorf("expected the error to be something but got something else. Actual: %v. Expected: %v", err, pkg.ErrRideNotCompleted)
			}

			if bill != nil {
				t.Errorf("expected no bill for a ride with ride id %v that hasn't been completed, but got bill: %v", rideId, bill)
			}
		})

	})
}
