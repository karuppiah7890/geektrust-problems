package pkg

import "errors"

var ErrRideIdExist = errors.New("ride id already exists")

var ErrDriverIdNotExist = errors.New("driver id does not exist")

var ErrDriverNotAvailable = errors.New("driver is not available for a ride")

var ErrRiderIdNotExist = errors.New("rider id does not exist")

var ErrRiderOnRide = errors.New("rider is already on a ride")
