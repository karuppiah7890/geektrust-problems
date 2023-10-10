package pkg_test

import (
	"geektrust/pkg"
	"testing"
)

// ride started case

// invalid ride case:
// 1. ride id already exists
// 2. drive id does not exist
// 3. driver is not available
// 4. rider id does not exist
// 5. rider is already on ride

func TestStartRide(t *testing.T) {
	t.Run("start a valid ride", func(t *testing.T) {
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

		ride, ok := rideSharingApp.GetRide(rideId)
		if !ok {
			t.Errorf("expected to get ride with given ride id %v, but got none", rideId)
		}

		assertStringEqual(t, ride.GetId(), rideId)
		assertStringEqual(t, ride.GetRiderId(), riderId)
		assertStringEqual(t, ride.GetDriverId(), driverId)
		assertStringEqual(t, ride.GetDriverId(), driverId)
		assertBoolEqual(t, ride.IsComplete(), false)
	})
	t.Run("invalid ride cases", func(t *testing.T) {
		t.Run("ride id already exists", func(t *testing.T) {})
	})
}
