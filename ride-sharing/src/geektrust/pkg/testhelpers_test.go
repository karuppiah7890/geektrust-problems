package pkg_test

import (
	"geektrust/pkg"
	"geektrust/pkg/driver"
	"geektrust/pkg/location"
	"geektrust/pkg/rider"
	"testing"
)

func assertLocationEqual(t *testing.T, actual *location.Location, expected *location.Location) bool {
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

func addDrivers(t *testing.T, rideSharingApp *pkg.RideSharingApp, drivers []*driver.Driver) {
	for _, driver := range drivers {
		addDriver(t, rideSharingApp, driver)
	}
}

func addDriver(t *testing.T, rideSharingApp *pkg.RideSharingApp, driver *driver.Driver) {
	input := &pkg.AddDriverInput{
		DriverId: driver.GetID(),
		Location: driver.GetLocation(),
	}

	err := rideSharingApp.AddDriver(input)
	if err != nil {
		t.Errorf("expected no error occur while adding driver but got error: %v", err)
	}
}

func addRiders(t *testing.T, rideSharingApp *pkg.RideSharingApp, riders []*rider.Rider) {
	for _, rider := range riders {
		addRider(t, rideSharingApp, rider)
	}
}

func addRider(t *testing.T, rideSharingApp *pkg.RideSharingApp, rider *rider.Rider) {
	input := &pkg.AddRiderInput{
		RiderId:  rider.GetID(),
		Location: rider.GetLocation(),
	}

	err := rideSharingApp.AddRider(input)
	if err != nil {
		t.Errorf("expected no error occur while adding rider but got error: %v", err)
	}
}

func newDriver(driverId string, x float64, y float64) *driver.Driver {
	loc := location.New(x, y)
	return driver.New(driverId, loc, true)
}

func newRider(riderId string, x float64, y float64) *rider.Rider {
	loc := location.New(x, y)
	return rider.New(riderId, loc, false)
}
