package main

import (
	"fmt"
	"geektrust/pkg"
)

// TODO: This is kind of like similar/duplicate logic when compared to add driver logic.
// The entity here is rider, and action is add-rider and input is add-rider-input, that's all,
// no other difference, just some small changes in error strings
func addRider(rideSharingApp *pkg.RideSharingApp, inputLineNumber int, commandInput []string) {
	numberOfInputs := len(commandInput)
	if numberOfInputs != 3 {
		panic(fmt.Sprintf("expected exactly 3 inputs for add rider command in line %d, but got %d inputs", inputLineNumber, numberOfInputs))
	}

	riderId := commandInput[0]
	x := parseCoordinate(commandInput[1])
	y := parseCoordinate(commandInput[2])

	location := pkg.NewLocation(x, y)

	input := &pkg.AddRiderInput{
		RiderId:  riderId,
		Location: location,
	}

	err := rideSharingApp.AddRider(input)
	if err != nil {
		panic(fmt.Sprintf("error occurred while adding rider in line %d: %v", inputLineNumber, err))
	}
}
