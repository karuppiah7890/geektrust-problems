package pkg_test

import (
	"geektrust/pkg"
	"testing"
)

func TestGetDriver(t *testing.T) {
	t.Run("fail when getting a driver with driver id that that does not exist", func(t *testing.T) {
		rideSharingApp := pkg.NewRideSharingApp()

		driverId := "D1"

		driver, ok := rideSharingApp.GetDriver(driverId)
		if ok {
			t.Errorf("expected to not get a driver with driver id %v as driver id does not exist, but got a driver", driverId)
		}

		if driver != nil {
			t.Errorf("expected driver to be nil but got %v", driver)
		}
	})
}
