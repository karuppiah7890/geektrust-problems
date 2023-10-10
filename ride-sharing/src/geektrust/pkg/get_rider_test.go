package pkg_test

import (
	"geektrust/pkg"
	"testing"
)

func TestGetRider(t *testing.T) {
	t.Run("fail when getting a rider with rider id that that does not exist", func(t *testing.T) {
		rideSharingApp := pkg.NewRideSharingApp()

		riderId := "R1"

		rider, ok := rideSharingApp.GetRider(riderId)
		if ok {
			t.Errorf("expected to not get rider with rider id %v that does not exist but got one rider: %v", riderId, rider)
		}

		if rider != nil {
			t.Errorf("expected rider to be nil but got %v", rider)
		}
	})
}
