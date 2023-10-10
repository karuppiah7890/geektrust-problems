package pkg_test

import (
	"geektrust/pkg"
	"testing"
)

func TestAddDriver(t *testing.T) {
	t.Run("add a valid driver", func(t *testing.T) {
		rideSharingApp := pkg.NewRideSharingApp()

		driverId := "D1"

		location := &pkg.Location{
			X: 1,
			Y: 1,
		}

		input := &pkg.AddDriverInput{
			DriverId: &driverId,
			Location: location,
		}

		err := rideSharingApp.AddDriver(input)
		if err != nil {
			t.Errorf("expected no error occur while adding driver but got error: %v", err)
		}

		driver, ok := rideSharingApp.GetDriver(driverId)
		if !ok {
			t.Errorf("could not get driver with driver id %v", driverId)
		}

		assertLocationEqual(t, driver.Location, location)
		assertStringEqual(t, driver.ID, driverId)
	})

	t.Run("fail when adding a driver with driver id that already exists", func(t *testing.T) {
		rideSharingApp := pkg.NewRideSharingApp()

		driverId := "D1"

		location := &pkg.Location{
			X: 1,
			Y: 1,
		}

		input := &pkg.AddDriverInput{
			DriverId: &driverId,
			Location: location,
		}

		err := rideSharingApp.AddDriver(input)
		if err != nil {
			t.Errorf("expected no error occur while adding driver but got error: %v", err)
		}

		err = rideSharingApp.AddDriver(input)
		if err == nil {
			t.Error("expected error to occur while adding driver with driver id that already exists but got no error")
			return
		}

		expectedError := "a driver with id D1 already exists"

		if err.Error() != expectedError {
			t.Errorf("expected the error to be equal. Actual: %s. Expected: %s", err.Error(), expectedError)
		}
	})
}
