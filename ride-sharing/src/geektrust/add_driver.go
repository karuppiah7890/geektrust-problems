package main

import (
	"fmt"
	"geektrust/pkg"
)

func addDriver(rideSharingApp *pkg.RideSharingApp, inputLineNumber int, commandInput []string) {
	numberOfInputs := len(commandInput)
	if numberOfInputs != 3 {
		panic(fmt.Sprintf("expected exactly 3 inputs for add driver command in line %d, but got %d inputs", inputLineNumber, numberOfInputs))
	}

	driverId := commandInput[0]
	x := parseCoordinate(commandInput[1])
	y := parseCoordinate(commandInput[2])

	location := pkg.Location{
		X: x,
		Y: y,
	}

	input := &pkg.AddDriverInput{
		DriverId: &driverId,
		Location: &location,
	}

	err := rideSharingApp.AddDriver(input)
	if err != nil {
		panic(fmt.Sprintf("error occurred while adding driver in line %d: %v", inputLineNumber, err))
	}
}
