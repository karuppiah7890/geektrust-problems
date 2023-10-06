package pkg

func (r *RideSharingApp) GetRide(rideId string) (*Ride, bool) {
	ride, ok := r.rides[rideId]
	return ride, ok
}
