package pkg

import "fmt"

func (r *RideSharingApp) GetRider(riderId string) (*Rider, error) {
	rider, ok := r.riders[riderId]
	if !ok {
		return nil, fmt.Errorf("rider with id %s does not exist", riderId)
	}

	return rider, nil
}
