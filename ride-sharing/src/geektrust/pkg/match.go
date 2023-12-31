package pkg

import (
	"container/heap"
	"fmt"
	"geektrust/pkg/driver"
)

type MatchRiderWithDriverInput struct {
	RiderId    string
	RadiusInKm float64
	MaxDrivers int
}

type MatchedDriver struct {
	*driver.Driver
	distanceFromRider float64
}

type MatchedDrivers []*MatchedDriver

// MatchRiderWithDriver returns Driver IDs of Drivers in ascending order of
// their distance from the rider. In the event of multiple drivers being
// equidistant, it will return them in lexicographical order.
func (r *RideSharingApp) MatchRiderWithDriver(input *MatchRiderWithDriverInput) ([]string, error) {
	riderId := input.RiderId
	rider, ok := r.GetRider(riderId)
	if !ok {
		return nil, fmt.Errorf("could not get rider with id %s: %w", riderId, ErrRiderIdNotExist)
	}

	matchedDrivers := make(MatchedDrivers, 0)

	// find the distance between the rider and each of the drivers in the drivers list.
	// TODO: Get only available drivers and not all drivers. Available means - drivers
	// who are not on a ride
	for _, driver := range r.drivers {
		distance := rider.GetLocation().DistanceBetween(driver.GetLocation())
		// if distance between them is less than radius in KM then insert it into matched drivers, or else leave it.
		if distance <= input.RadiusInKm {
			matchedDrivers = append(matchedDrivers, &MatchedDriver{
				Driver:            driver,
				distanceFromRider: distance,
			})
		}
	}

	if len(matchedDrivers) == 0 {
		return []string{}, nil
	}

	idsOfMatchedDrivers := make([]string, 0, min(input.MaxDrivers, matchedDrivers.Len()))

	// use min heap to get drivers in ascending order,
	// sorted by distance of driver from rider
	heap.Init(&matchedDrivers)

	// return MaxDrivers amount of drivers in the output.
	// it's also possible that the numer of matched drivers
	// is less than MaxDrivers.
	for count := 0; count < input.MaxDrivers && matchedDrivers.Len() > 0; count++ {
		value := heap.Pop(&matchedDrivers)
		d, ok := value.(*MatchedDriver)
		if !ok {
			panic(fmt.Sprintf("unexpected error occurred: not able to convert a value popped from matched drivers heap to a driver. value: %v", value))
		}
		idsOfMatchedDrivers = append(idsOfMatchedDrivers, d.GetID())
	}

	return idsOfMatchedDrivers, nil
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Ensuring that MatchedDrivers implements heap.Interface
var _ heap.Interface = &MatchedDrivers{}

func (m MatchedDrivers) Len() int {
	return len(m)
}

func (m MatchedDrivers) Less(i, j int) bool {
	driver := m[i]
	anotherDriver := m[j]

	if driver.distanceFromRider == anotherDriver.distanceFromRider {
		return isLexicographicallyOrdered(driver.GetID(), anotherDriver.GetID())
	}

	return driver.distanceFromRider < anotherDriver.distanceFromRider
}

func isLexicographicallyOrdered(s1 string, s2 string) bool {
	return s1 < s2
}

func (m MatchedDrivers) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MatchedDrivers) Push(driver interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*m = append(*m, driver.(*MatchedDriver))
}

func (m *MatchedDrivers) Pop() interface{} {
	old := *m
	n := len(old)
	driver := old[n-1]
	old[n-1] = nil // avoid memory leak
	*m = old[0 : n-1]
	return driver
}
