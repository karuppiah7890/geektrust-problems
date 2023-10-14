package main

import (
	"bufio"
	"fmt"
	"geektrust/pkg"
	"os"
	"strings"
)

func processFile(file *os.File) {
	c := newContext()
	rideSharingApp := pkg.NewRideSharingApp()

	scanner := bufio.NewScanner(file)

	inputLineNumber := 1

	// Read input file line by line
	for scanner.Scan() {

		args := scanner.Text()
		argList := strings.Fields(args)

		if len(argList) == 0 {
			panic(fmt.Sprintf("expected every line to contain some input but found none in line %d", inputLineNumber))
		}

		command := argList[0]
		commandInput := argList[1:]

		switch command {
		case "ADD_DRIVER":
			addDriver(rideSharingApp, inputLineNumber, commandInput)

		case "ADD_RIDER":
			addRider(rideSharingApp, inputLineNumber, commandInput)

		case "MATCH":
			match(c, rideSharingApp, inputLineNumber, commandInput)

		case "START_RIDE":
			startRide(c, rideSharingApp, inputLineNumber, commandInput)

		case "STOP_RIDE":
			stopRide(c, rideSharingApp, inputLineNumber, commandInput)

		case "BILL":
			bill(rideSharingApp, inputLineNumber, commandInput)

		default:
			panic(fmt.Sprintf("expected every line to contain a known command but found command %s in line %d. Known commands are ADD_DRIVER, ADD_RIDER, MATCH, START_RIDE, STOP_RIDE, BILL", command, inputLineNumber))
		}

		inputLineNumber++
	}
}
