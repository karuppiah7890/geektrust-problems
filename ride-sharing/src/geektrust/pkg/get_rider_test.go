package pkg_test

import (
	"geektrust/pkg"
	"testing"
)

func TestGetRider(t *testing.T) {
	t.Run("fail when getting a rider with rider id that that does not exist", func(t *testing.T) {
		rideSharingApp := pkg.NewRideSharingApp()

		riderId := "R1"

		rider, err := rideSharingApp.GetRider(riderId)
		if err == nil {
			t.Error("expected error to occur while getting rider with rider id that does not exist but got no error")
		}

		if rider != nil {
			t.Errorf("expected rider to be nil but got %v", rider)
		}

		expectedError := "rider with id R1 does not exist"

		if err.Error() != expectedError {
			t.Errorf("expected the error to be equal. Actual: %s. Expected: %s", err.Error(), expectedError)
		}
	})
}
