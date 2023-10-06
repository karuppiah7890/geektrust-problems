package pkg

func (r *RideSharingApp) GetRider(riderId string) (*Rider, bool) {
	rider, ok := r.riders[riderId]
	return rider, ok
}
