package pkg

import "geektrust/pkg/ride"

func (app *RideSharingApp) GetRide(rideId string) (*ride.Ride, bool) {
	r, ok := app.rides[rideId]
	return r, ok
}
