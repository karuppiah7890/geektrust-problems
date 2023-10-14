package pkg

import (
	"geektrust/pkg/driver"
	"geektrust/pkg/ride"
)

type RideSharingApp struct {
	drivers map[string]*driver.Driver
	riders  map[string]*Rider
	rides   map[string]*ride.Ride
}

func NewRideSharingApp() *RideSharingApp {
	return &RideSharingApp{
		drivers: make(map[string]*driver.Driver),
		riders:  make(map[string]*Rider),
		rides:   make(map[string]*ride.Ride),
	}
}
