package main

import (
	"fmt"
	"strconv"
)

func parseCoordinate(x string) float64 {
	floatX, err := strconv.ParseFloat(x, 64)

	if err != nil {
		panic(fmt.Sprintf("expected coordinate to be a valid number with or without decimal points. %s is not a valid coordinate", x))
	}

	return floatX
}
