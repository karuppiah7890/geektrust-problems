package ride

import "geektrust/pkg/location"

type Ride struct {
	id                    string
	isComplete            bool
	riderId               string
	driverId              string
	source                *location.Location
	destination           *location.Location
	rideDurationInMinutes int
}

func New(id string, riderId string, driverId string, source *location.Location) *Ride {
	return &Ride{
		id:                    id,
		isComplete:            false,
		riderId:               riderId,
		driverId:              driverId,
		source:                source,
		destination:           nil,
		rideDurationInMinutes: 0,
	}
}

// TODO: This is a setter. See if we can get rid of this setter
// and set this value as part of a "stop ride" process / method
func (r *Ride) Complete(destination *location.Location, rideDurationInMinutes int) {
	r.isComplete = true
	r.destination = destination
	r.rideDurationInMinutes = rideDurationInMinutes
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

func (r *Ride) GetSource() *location.Location {
	return r.source
}

func (r *Ride) GetDestination() *location.Location {
	return r.destination
}

func (r *Ride) GetRideDurationInMinutes() int {
	return r.rideDurationInMinutes
}
