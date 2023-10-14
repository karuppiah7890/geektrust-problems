package main

import (
	"fmt"
	"geektrust/pkg"
)

func bill(rideSharingApp *pkg.RideSharingApp, inputLineNumber int, commandInput []string) {
	numberOfInputs := len(commandInput)
	if numberOfInputs != 1 {
		panic(fmt.Sprintf("expected exactly 1 input for bill command in line %d, but got %d inputs", inputLineNumber, numberOfInputs))
	}

	rideId := commandInput[0]

	_ = rideId
}
