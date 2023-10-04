package pkg

type MatchRiderWithDriverInput struct {
	RiderId    string
	RadiusInKm float64
	MaxDrivers int
}

// MatchRiderWithDriver returns Driver IDs of Drivers in ascending order of
// their distance from the rider. In the event of multiple drivers being
// equidistant, it will return them in lexicographical order.
func (r *RideSharingApp) MatchRiderWithDriver(input *MatchRiderWithDriverInput) ([]string, error) {
	return nil, nil
}
