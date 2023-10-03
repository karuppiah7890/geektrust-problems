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
