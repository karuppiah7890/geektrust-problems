package ride

type Ride struct {
	id         string
	isComplete bool
	riderId    string
	driverId   string
}

func New(id string, isComplete bool, riderId string, driverId string) *Ride {
	return &Ride{
		id:         id,
		isComplete: isComplete,
		riderId:    riderId,
		driverId:   driverId,
	}
}

// TODO: This is a setter. See if we can get rid of this setter
// and set this value as part of a "start ride" process / method
func (r *Ride) Complete() {
	r.isComplete = true
}

func (r *Ride) GetId() string {
	return r.id
}

func (r *Ride) IsComplete() bool {
	return r.isComplete
}

func (r *Ride) GetRiderId() string {
	return r.riderId
}

func (r *Ride) GetDriverId() string {
	return r.driverId
}
