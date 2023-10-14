package pkg

import "geektrust/pkg/rider"

func (r *RideSharingApp) GetRider(riderId string) (*rider.Rider, bool) {
	rider, ok := r.riders[riderId]
	return rider, ok
}
