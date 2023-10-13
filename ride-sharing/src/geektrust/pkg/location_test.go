package pkg_test

import (
	"geektrust/pkg"
	"testing"
)

func LocationTest(t *testing.T) {
	t.Run("Get X and Y", func(t *testing.T) {
		x := float64(1)
		y := float64(2)
		
		location := pkg.NewLocation(x, y)

		if location.GetX() != x {
			t.Errorf("expected getter to work but it didn't. actual: %v. expected: %v", location.GetX(), x)
		}

		if location.GetY() != y {
			t.Errorf("expected getter to work but it didn't. actual: %v. expected: %v", location.GetY(), y)
		}
	})
}
