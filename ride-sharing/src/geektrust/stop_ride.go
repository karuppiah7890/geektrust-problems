package main

import (
	"errors"
	"fmt"
	"geektrust/pkg"
	"geektrust/pkg/location"
	"strconv"
)

const RIDE_STOPPED = "RIDE_STOPPED"

func stopRide(c *context, rideSharingApp *pkg.RideSharingApp, inputLineNumber int, commandInput []string) {
	numberOfInputs := len(commandInput)
	if numberOfInputs != 4 {
		panic(fmt.Sprintf("expected exactly 4 inputs for stop ride command in line %d, but got %d inputs", inputLineNumber, numberOfInputs))
	}

	rideId := commandInput[0]
	destinationXStr := commandInput[1]
	destinationYStr := commandInput[2]
	timeTakenInMinutesStr := commandInput[3]

	destinationX := parseCoordinate(destinationXStr)
	destinationY := parseCoordinate(destinationYStr)
	timeTakenInMinutes := stringToInt(timeTakenInMinutesStr, "time taken in minutes")

	input := &pkg.StopRideInput{
		RideId:             rideId,
		Destination:        location.New(destinationX, destinationY),
		TimeTakenInMinutes: timeTakenInMinutes,
	}

	err := rideSharingApp.StopRide(input)
	if err != nil {
		if isKnownErrorForStopRide(err) {
			fmt.Println(INVALID_RIDE)
			return
		} else {
			panic(fmt.Sprintf("unknown error occurred while stopping ride for ride id %v in line %d: %v", rideId, inputLineNumber, err))
		}
	}

	fmt.Printf("%v %v\n", RIDE_STOPPED, rideId)
}

func isKnownErrorForStopRide(err error) bool {
	return errors.Is(err, pkg.ErrRideIdNotExist) ||
		errors.Is(err, pkg.ErrRideStopped)
}

func stringToInt(s string, fieldInfo string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("could not convert string to number for field: %s. Error: %v", fieldInfo, err))
	}
	return number
}
