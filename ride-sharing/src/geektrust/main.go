package main

import (
	"fmt"
	"os"
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
