package main

import (
	"errors"
	"fmt"
	"geektrust/cmd/context"
	"geektrust/pkg"
	"strconv"
)

const INVALID_RIDE = "INVALID_RIDE"
const RIDE_STARTED = "RIDE_STARTED"

func startRide(c *context.Context, rideSharingApp *pkg.RideSharingApp, inputLineNumber int, commandInput []string) {
	numberOfInputs := len(commandInput)
	if numberOfInputs != 3 {
		panic(fmt.Sprintf("expected exactly 3 inputs for start ride command in line %d, but got %d inputs", inputLineNumber, numberOfInputs))
	}

	rideId := commandInput[0]
	driverOptionNumber, err := strconv.ParseInt(commandInput[1], 10, 64)
	if err != nil {
		panic("error occurred while parsing %v as a number")
	}
	if driverOptionNumber <= 0 {
		panic(fmt.Sprintf("driver option number (N) should be a number greater than 0 but it is %v", driverOptionNumber))
	}

	riderId := commandInput[2]

	driverId, ok := getDriverIdFromOptions(c, riderId, driverOptionNumber)

	if !ok {
		fmt.Println(INVALID_RIDE)
		return
	}

	input := &pkg.StartRideInput{
		RideId:   rideId,
		RiderId:  riderId,
		DriverId: driverId,
	}

	err = rideSharingApp.StartRide(input)
	if err != nil {
		if isKnownErrorForStartRide(err) {
			fmt.Println(INVALID_RIDE)
			return
		} else {
			panic(fmt.Sprintf("unknown error occurred while starting ride for rider id %v in line %d, driver id %v, ride id %v: %v", riderId, inputLineNumber, driverId, rideId, err))
		}
	}

	c.DeleteDriverOptionsForRider(riderId)

	fmt.Printf("%v %v\n", RIDE_STARTED, rideId)
}

func isKnownErrorForStartRide(err error) bool {
	return errors.Is(err, pkg.ErrRideIdExist) ||
		errors.Is(err, pkg.ErrDriverIdNotExist) ||
		errors.Is(err, pkg.ErrDriverNotAvailable) ||
		errors.Is(err, pkg.ErrRiderIdNotExist) ||
		errors.Is(err, pkg.ErrRiderOnRide)
}

func getDriverIdFromOptions(c *context.Context, riderId string, optionNumber int64) (string, bool) {
	driverOptions, err := c.GetDriverOptionsForRider(riderId)
	if err != nil {
		if errors.Is(err, context.ErrDriverOptionsUnavailable) {
			panic(fmt.Sprintf("driver options unavailable error occurred: %v. maybe MATCH command was not called first", err))
		} else {
			panic(fmt.Sprintf("unknown error occurred while getting driver options for rider: %v", err))
		}
	}

	if driverOptions == nil {
		panic(fmt.Sprintf("expected driver options to not be nil, but driver options is nil for rider id %v", riderId))
	}

	if len(driverOptions) < int(optionNumber) {
		return "", false
	}

	driverId := driverOptions[optionNumber-1]

	return driverId, true
}
