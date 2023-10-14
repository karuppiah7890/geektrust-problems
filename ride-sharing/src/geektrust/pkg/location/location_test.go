package location_test

import (
	"geektrust/pkg/location"
	"math"
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

	t.Run("Distance Between", func(t *testing.T) {
		loc := location.New(1, 1)
		anotherLoc := location.New(2, 2)

		expected := math.Sqrt(2)
		actual := loc.DistanceBetween(anotherLoc)

		if expected != actual {
			t.Errorf("actual distance and expected distance is different. actual: %v. expected: %v", actual, expected)
		}
	})
}
