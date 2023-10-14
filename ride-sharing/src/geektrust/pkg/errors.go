package pkg

import "errors"

var ErrRideIdExist = errors.New("ride id already exists")

var ErrRideIdNotExist = errors.New("ride id does not exist")

var ErrRideStopped = errors.New("ride is already stopped")

var ErrRideNotCompleted = errors.New("ride is not completed")

var ErrDriverIdNotExist = errors.New("driver id does not exist")

var ErrDriverNotAvailable = errors.New("driver is not available for a ride")

var ErrRiderIdNotExist = errors.New("rider id does not exist")

var ErrRiderOnRide = errors.New("rider is already on a ride")
