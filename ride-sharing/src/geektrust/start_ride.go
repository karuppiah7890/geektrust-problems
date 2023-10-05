package main

import (
	"fmt"
	"geektrust/pkg"
	"strconv"
)

const INVALID_RIDE = "INVALID_RIDE"
const RIDE_STARTED = "RIDE_STARTED"

func startRide(c *context, rideSharingApp *pkg.RideSharingApp, inputLineNumber int, commandInput []string) {
	numberOfInputs := len(commandInput)
	if numberOfInputs != 3 {
		panic(fmt.Sprintf("expected exactly 3 inputs for match command in line %d, but got %d inputs", inputLineNumber, numberOfInputs))
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

	driverId := getDriverIdFromOptions(c, riderId, driverOptionNumber)

	if driverId == nil {
		fmt.Println(INVALID_RIDE)
		return
	}

	input := &pkg.StartRideInput{
		RideId:   &rideId,
		RiderId:  &riderId,
		DriverId: driverId,
	}

	err = rideSharingApp.StartRide(input)
	if err != nil {
		_, ok := err.(pkg.RideIdAlreadyExists)
		if ok {
			fmt.Println(INVALID_RIDE)
		} else {
			panic(fmt.Sprintf("unknown error occurred while starting ride for rider id %v in line %d, driver id %v, ride id %v: %v", riderId, inputLineNumber, driverId, rideId, err))
		}
	}

	c.deleteDriverOptionsForRider(riderId)

	fmt.Printf("%v %v\n", RIDE_STARTED, rideId)
}

func getDriverIdFromOptions(c *context, riderId string, optionNumber int64) *string {
	driverOptions, err := c.getDriverOptionsForRider(riderId)
	if err != nil {
		_, ok := err.(DriverOptionsUnavailableForRider)
		if ok {
			panic(fmt.Sprintf("driver options unavailable error occurred: %v. maybe MATCH command was not called first", err))
		} else {
			panic(fmt.Sprintf("unknown error occurred while getting driver options for rider: %v", err))
		}
	}

	if driverOptions == nil {
		panic(fmt.Sprintf("expected driver options to not be nil, but driver options is nil for rider id %v", riderId))
	}

	if len(driverOptions) < int(optionNumber) {
		return nil
	}

	driverId := driverOptions[optionNumber-1]

	return &driverId
}
