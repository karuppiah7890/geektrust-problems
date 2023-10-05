package main

import (
	"fmt"
	"geektrust/pkg"
)

const RADIUS_IN_KM = 5
const MAX_MATCHED_DRIVERS = 5

func match(c *context, rideSharingApp *pkg.RideSharingApp, inputLineNumber int, commandInput []string) {
	numberOfInputs := len(commandInput)
	if numberOfInputs != 1 {
		panic(fmt.Sprintf("expected exactly 1 inputs for match command in line %d, but got %d inputs", inputLineNumber, numberOfInputs))
	}

	riderId := commandInput[0]

	input := &pkg.MatchRiderWithDriverInput{
		RiderId:    riderId,
		RadiusInKm: RADIUS_IN_KM,
		MaxDrivers: MAX_MATCHED_DRIVERS,
	}

	idsOfMatchedDrivers, err := rideSharingApp.MatchRiderWithDriver(input)
	if err != nil {
		panic(fmt.Sprintf("error occurred while trying to match rider with nearest available drivers in line %d: %v", inputLineNumber, err))
	}

	if len(idsOfMatchedDrivers) == 0 {
		c.storeDriverOptionsForRider(riderId, idsOfMatchedDrivers)
		fmt.Println("NO_DRIVERS_AVAILABLE")
	} else {
		c.storeDriverOptionsForRider(riderId, idsOfMatchedDrivers)
		fmt.Print("DRIVERS_MATCHED")
		for _, driverId := range idsOfMatchedDrivers {
			fmt.Printf(" %s", driverId)
		}
		fmt.Println()
	}
}
