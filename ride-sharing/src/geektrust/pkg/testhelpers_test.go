package pkg_test

import (
	"geektrust/pkg"
	"testing"
)

func assertLocationEqual(t *testing.T, actual *pkg.Location, expected *pkg.Location) bool {
	if actual.X != expected.X || actual.Y != expected.Y {
		t.Errorf("expected the locations to be equal but they were not. Actual: (%v, %v). Expected: (%v, %v)", actual.X, actual.Y, expected.X, expected.Y)
		return false
	}

	return true
}

func assertStringEqual(t *testing.T, actual string, expected string) bool {
	if actual != expected {
		t.Errorf("expected the strings to be equal but they were not. Actual: %v. Expected: %v", actual, expected)
		return false
	}

	return true
}

func assertBoolEqual(t *testing.T, actual bool, expected bool) bool {
	if actual != expected {
		t.Errorf("expected the bools to be equal but they were not. Actual: %v. Expected: %v", actual, expected)
		return false
	}

	return true
}

func assertStringArrayEqual(t *testing.T, actual []string, expected []string) bool {
	if len(actual) != len(expected) {
		t.Errorf("expcted the length of the string arrays to be equal but they were not. Actual: %v. Expected: %v", len(actual), len(expected))
		return false
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("expcted the string arrays to be equal but they were not. Actual: %v. Expected: %v", actual, expected)
			return false
		}
	}

	return true
}

func addDrivers(t *testing.T, rideSharingApp *pkg.RideSharingApp, drivers []*pkg.Driver) {
	for _, driver := range drivers {
		addDriver(t, rideSharingApp, driver)
	}
}

func addDriver(t *testing.T, rideSharingApp *pkg.RideSharingApp, driver *pkg.Driver) {
	location := &pkg.Location{
		X: driver.Location.X,
		Y: driver.Location.Y,
	}

	input := &pkg.AddDriverInput{
		DriverId: driver.ID,
		Location: location,
	}

	err := rideSharingApp.AddDriver(input)
	if err != nil {
		t.Errorf("expected no error occur while adding driver but got error: %v", err)
	}
}

func addRiders(t *testing.T, rideSharingApp *pkg.RideSharingApp, riders []*pkg.Rider) {
	for _, rider := range riders {
		addRider(t, rideSharingApp, rider)
	}
}

func addRider(t *testing.T, rideSharingApp *pkg.RideSharingApp, rider *pkg.Rider) {
	location := &pkg.Location{
		X: rider.Location.X,
		Y: rider.Location.Y,
	}

	input := &pkg.AddRiderInput{
		RiderId:  &rider.ID,
		Location: location,
	}

	err := rideSharingApp.AddRider(input)
	if err != nil {
		t.Errorf("expected no error occur while adding rider but got error: %v", err)
	}
}

func driver(driverId string, x float64, y float64) *pkg.Driver {
	return &pkg.Driver{
		ID: driverId,
		Location: &pkg.Location{
			X: x,
			Y: y,
		},
	}
}

func rider(riderId string, x float64, y float64) *pkg.Rider {
	return &pkg.Rider{
		ID: riderId,
		Location: &pkg.Location{
			X: x,
			Y: y,
		},
	}
}
