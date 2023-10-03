package pkg

import "fmt"

type AddRiderInput struct {
	RiderId  *string
	Location *Location
}

func (r *RideSharingApp) AddRider(input *AddRiderInput) error {
	if _, ok := r.riders[*input.RiderId]; ok {
		return fmt.Errorf("a rider with id %s already exists", *input.RiderId)
	}

	r.riders[*input.RiderId] = &Rider{
		ID: *input.RiderId,
		Location: &Location{
			X: input.Location.X,
			Y: input.Location.Y,
		},
	}

	return nil
}
