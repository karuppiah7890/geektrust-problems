package pkg

import "fmt"

type AddRiderInput struct {
	RiderId  string
	Location *Location
}

func (r *RideSharingApp) AddRider(input *AddRiderInput) error {
	riderId := input.RiderId

	if _, ok := r.riders[riderId]; ok {
		return fmt.Errorf("a rider with id %s already exists", riderId)
	}

	r.riders[riderId] = &Rider{
		ID:       riderId,
		Location: input.Location.Clone(),
	}

	return nil
}
