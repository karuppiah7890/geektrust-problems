package pkg

import "geektrust/pkg/ride"

type RideSharingApp struct {
	drivers map[string]*Driver
	riders  map[string]*Rider
	rides   map[string]*ride.Ride
}

func NewRideSharingApp() *RideSharingApp {
	return &RideSharingApp{
		drivers: make(map[string]*Driver),
		riders:  make(map[string]*Rider),
		rides:   make(map[string]*ride.Ride),
	}
}
