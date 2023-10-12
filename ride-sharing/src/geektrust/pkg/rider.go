package pkg

type Rider struct {
	ID       string
	Location *Location
	isOnRide bool
}

func (r *Rider) IsOnRide() bool {
	return r.isOnRide
}
