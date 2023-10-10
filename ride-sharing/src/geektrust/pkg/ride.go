package pkg

type Ride struct {
	id         string
	isComplete bool
	riderId    string
	driverId   string
}

func (ride *Ride) GetId() string {
	return ride.id
}

func (ride *Ride) IsComplete() bool {
	return ride.isComplete
}

func (ride *Ride) GetRiderId() string {
	return ride.riderId
}

func (ride *Ride) GetDriverId() string {
	return ride.driverId
}
