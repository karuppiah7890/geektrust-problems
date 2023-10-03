package main

import (
	"bufio"
	"fmt"
	"geektrust/pkg"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := getInputFilePath()

	file := openFile(filePath)

	defer file.Close()

	processFile(file)
}

func getInputFilePath() string {
	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")
		os.Exit(1)
	}

	return cliArgs[0]
}

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening the input file")
		os.Exit(1)
	}

	return file
}

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
		}

		lineNumber++
	}
}

func addDriver(rideSharingApp *pkg.RideSharingApp, lineNumber int, argList []string) {
	numberOfInputs := len(argList)
	if numberOfInputs != 3 {
		panic(fmt.Sprintf("expected exactly 3 inputs for add driver command in line %d, but got %d inputs", lineNumber, numberOfInputs))
	}

	driverId := argList[0]
	x := parseCoordinate(argList[1])
	y := parseCoordinate(argList[2])

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
		panic(fmt.Sprintf("error occurred while adding driver in line %d: %v", lineNumber, err))
	}
}

func parseCoordinate(x string) float64 {
	floatX, err := strconv.ParseFloat(x, 64)

	if err != nil {
		panic(fmt.Sprintf("expected coordinate to be a valid number with or without decimal points. %s is not a valid coordinate", x))
	}

	return floatX
}
