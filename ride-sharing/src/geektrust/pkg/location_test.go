package pkg_test

import (
	"geektrust/pkg"
	"testing"
)

func TestLocation(t *testing.T) {
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

	t.Run("Clone", func(t *testing.T) {
		location := pkg.NewLocation(1, 2)
		clonedLocation := location.Clone()

		if location.GetX() != clonedLocation.GetX() {
			t.Errorf("expected x coordinates to be equal but they were not. actual: %v. expected: %v", location.GetX(), clonedLocation.GetX())
		}

		if location.GetY() != clonedLocation.GetY() {
			t.Errorf("expected y coordinates to be equal but they were not. actual: %v. expected: %v", location.GetY(), clonedLocation.GetY())
		}

		if location == clonedLocation {
			t.Errorf("expected cloned location to have different pointer address but got the same address")
		}
	})

	t.Run("Equals", func(t *testing.T) {
		t.Run("Both are equal", func(t *testing.T) {
			x := float64(1)
			y := float64(2)

			location := pkg.NewLocation(x, y)
			anotherLocation := pkg.NewLocation(x, y)

			if !location.Equals(anotherLocation) {
				t.Errorf("expected the two locations to be equal but they were not")
			}
		})

		t.Run("Both are unequal", func(t *testing.T) {
			t.Run("x is different", func(t *testing.T) {
				location := pkg.NewLocation(5, 2)
				anotherLocation := pkg.NewLocation(7, 2)

				if location.Equals(anotherLocation) {
					t.Errorf("expected the two locations to be unequal but they were equal")
				}
			})

			t.Run("y is different", func(t *testing.T) {
				location := pkg.NewLocation(1, 10)
				anotherLocation := pkg.NewLocation(1, 30)

				if location.Equals(anotherLocation) {
					t.Errorf("expected the two locations to be unequal but they were equal")
				}
			})
		})
	})
}
