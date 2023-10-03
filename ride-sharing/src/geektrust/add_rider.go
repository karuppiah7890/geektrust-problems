package main

import (
	"fmt"
	"geektrust/pkg"
)

// TODO: This is kind of like similar/duplicate logic when compared to add driver logic.
// The entity here is rider, and action is add-rider and input is add-rider-input, that's all,
// no other difference, just some small changes in error strings
func addRider(rideSharingApp *pkg.RideSharingApp, inputLineNumber int, argList []string) {
	numberOfInputs := len(argList)
	if numberOfInputs != 3 {
		panic(fmt.Sprintf("expected exactly 3 inputs for add driver command in line %d, but got %d inputs", inputLineNumber, numberOfInputs))
	}

	riderId := argList[0]
	x := parseCoordinate(argList[1])
	y := parseCoordinate(argList[2])

	location := pkg.Location{
		X: x,
		Y: y,
	}

	input := &pkg.AddRiderInput{
		RiderId:  &riderId,
		Location: &location,
	}

	err := rideSharingApp.AddRider(input)
	if err != nil {
		panic(fmt.Sprintf("error occurred while adding rider in line %d: %v", inputLineNumber, err))
	}
}
