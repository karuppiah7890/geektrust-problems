package pkg_test

import (
	"geektrust/pkg"
	"testing"
)

func TestGetDriver(t *testing.T) {
	t.Run("fail when getting a driver with driver id that that does not exist", func(t *testing.T) {
		rideSharingApp := pkg.NewRideSharingApp()

		driverId := "D1"

		driver, err := rideSharingApp.GetDriver(driverId)
		if err == nil {
			t.Error("expected error to occur while getting driver with driver id that does not exist but got no error")
		}

		if driver != nil {
			t.Errorf("expected driver to be nil but got %v", driver)
		}

		expectedError := "driver with id D1 does not exist"

		if err.Error() != expectedError {
			t.Errorf("expected the error to be equal. Actual: %s. Expected: %s", err.Error(), expectedError)
		}
	})
}
