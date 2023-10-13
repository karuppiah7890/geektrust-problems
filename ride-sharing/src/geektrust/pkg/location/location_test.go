package location_test

import (
	"geektrust/pkg/location"
	"testing"
)

func TestLocation(t *testing.T) {
	t.Run("Get X and Y", func(t *testing.T) {
		x := float64(1)
		y := float64(2)

		loc := location.New(x, y)

		if loc.GetX() != x {
			t.Errorf("expected getter to work but it didn't. actual: %v. expected: %v", loc.GetX(), x)
		}

		if loc.GetY() != y {
			t.Errorf("expected getter to work but it didn't. actual: %v. expected: %v", loc.GetY(), y)
		}
	})

	t.Run("Clone", func(t *testing.T) {
		loc := location.New(1, 2)
		clonedLocation := loc.Clone()

		if loc.GetX() != clonedLocation.GetX() {
			t.Errorf("expected x coordinates to be equal but they were not. actual: %v. expected: %v", loc.GetX(), clonedLocation.GetX())
		}

		if loc.GetY() != clonedLocation.GetY() {
			t.Errorf("expected y coordinates to be equal but they were not. actual: %v. expected: %v", loc.GetY(), clonedLocation.GetY())
		}

		if loc == clonedLocation {
			t.Errorf("expected cloned location to have different pointer address but got the same address")
		}
	})

	t.Run("Equals", func(t *testing.T) {
		t.Run("Both are equal", func(t *testing.T) {
			x := float64(1)
			y := float64(2)

			loc := location.New(x, y)
			anotherLoc := location.New(x, y)

			if !loc.Equals(anotherLoc) {
				t.Errorf("expected the two locations to be equal but they were not")
			}
		})

		t.Run("Both are unequal", func(t *testing.T) {
			t.Run("x is different", func(t *testing.T) {
				loc := location.New(5, 2)
				anotherLoc := location.New(7, 2)

				if loc.Equals(anotherLoc) {
					t.Errorf("expected the two locations to be unequal but they were equal")
				}
			})

			t.Run("y is different", func(t *testing.T) {
				loc := location.New(1, 10)
				anotherLoc := location.New(1, 30)

				if loc.Equals(anotherLoc) {
					t.Errorf("expected the two locations to be unequal but they were equal")
				}
			})
		})
	})
}
