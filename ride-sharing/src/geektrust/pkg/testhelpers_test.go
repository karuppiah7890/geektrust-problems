package pkg_test

import (
	"geektrust/pkg"
	"testing"
)

func assertLocationEqual(t *testing.T, actual *pkg.Location, expected *pkg.Location) bool {
	if !actual.Equals(expected) {
		t.Errorf("expected the locations to be equal but they were not. Actual: (%v, %v). Expected: (%v, %v)", actual.GetX(), actual.GetY(), expected.GetX(), expected.GetY())
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
	input := &pkg.AddDriverInput{
		DriverId: driver.GetID(),
		Location: driver.GetLocation().Clone(),
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
	input := &pkg.AddRiderInput{
		RiderId:  rider.ID,
		Location: rider.Location.Clone(),
	}

	err := rideSharingApp.AddRider(input)
	if err != nil {
		t.Errorf("expected no error occur while adding rider but got error: %v", err)
	}
}

func driver(driverId string, x float64, y float64) *pkg.Driver {
	location := pkg.NewLocation(x, y)
	return pkg.NewDriver(driverId, location, true)
}

func rider(riderId string, x float64, y float64) *pkg.Rider {
	return &pkg.Rider{
		ID:       riderId,
		Location: pkg.NewLocation(x, y),
	}
}
