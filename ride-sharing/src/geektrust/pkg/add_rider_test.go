package pkg_test

import (
	"geektrust/pkg"
	"geektrust/pkg/location"
	"testing"
)

func TestAddRider(t *testing.T) {
	t.Run("add a valid rider", func(t *testing.T) {
		rideSharingApp := pkg.NewRideSharingApp()

		riderId := "R1"

		loc := location.New(1, 1)

		input := &pkg.AddRiderInput{
			RiderId:  riderId,
			Location: loc,
		}

		err := rideSharingApp.AddRider(input)
		if err != nil {
			t.Errorf("expected no error occur while adding rider but got error: %v", err)
		}

		rider, ok := rideSharingApp.GetRider(riderId)
		if !ok {
			t.Errorf("expected to get rider for rider id %v but got none", riderId)
		}

		assertLocationEqual(t, rider.GetLocation(), loc)
		assertStringEqual(t, rider.GetID(), riderId)
	})

	t.Run("fail when adding a rider with rider id that already exists", func(t *testing.T) {
		rideSharingApp := pkg.NewRideSharingApp()

		riderId := "R1"

		loc := location.New(1, 1)

		input := &pkg.AddRiderInput{
			RiderId:  riderId,
			Location: loc,
		}

		err := rideSharingApp.AddRider(input)
		if err != nil {
			t.Errorf("expected no error occur while adding rider but got error: %v", err)
		}

		err = rideSharingApp.AddRider(input)
		if err == nil {
			t.Error("expected error to occur while adding rider with rider id that already exists but got no error")
			return
		}

		expectedError := "a rider with id R1 already exists"

		if err.Error() != expectedError {
			t.Errorf("expected the error to be equal. Actual: %s. Expected: %s", err.Error(), expectedError)
		}
	})
}
