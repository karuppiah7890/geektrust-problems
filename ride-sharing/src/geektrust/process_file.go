package main

import (
	"bufio"
	"fmt"
	"geektrust/pkg"
	"os"
	"strings"
)

func processFile(file *os.File) {
	rideSharingApp := pkg.NewRideSharingApp()

	scanner := bufio.NewScanner(file)

	lineNumber := 1

	// Read input file line by line
	for scanner.Scan() {

		args := scanner.Text()
		argList := strings.Fields(args)

		if len(argList) == 0 {
			panic(fmt.Sprintf("expected every line to contain some input but found none in line %d", lineNumber))
		}

		command := argList[0]

		switch command {
		case "ADD_DRIVER":
			addDriver(rideSharingApp, lineNumber, argList[1:])

		case "ADD_RIDER":
		case "MATCH":
		case "START_RIDE":
		case "STOP_RIDE":
		case "BILL":
		default:
			panic(fmt.Sprintf("expected every line to contain a known command but found command %s in line %d. Known commands are ADD_DRIVER, ADD_RIDER, MATCH, START_RIDE, STOP_RIDE, BILL", command, lineNumber))
		}

		lineNumber++
	}
}
