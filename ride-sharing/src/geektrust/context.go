package main

import "fmt"

type context struct {
	riderDetails *riderDetails
}

type riderDetails struct {
	driverOptionsForRider map[string][]string
}

func newContext() *context {
	return &context{
		riderDetails: &riderDetails{
			driverOptionsForRider: make(map[string][]string),
		},
	}
}

func (c *context) getDriverOptionsForRider(riderId string) ([]string, error) {
	checkContext(c)
	driverOptions, ok := c.riderDetails.driverOptionsForRider[riderId]
	if !ok {
		return nil, DriverOptionsUnavailableForRider(fmt.Sprintf("driver options are unavailable for rider id %v", riderId))
	}

	return driverOptions, nil
}

func (c *context) storeDriverOptionsForRider(riderId string, driverOptions []string) {
	checkContext(c)
	c.riderDetails.driverOptionsForRider[riderId] = driverOptions
}

func (c *context) deleteDriverOptionsForRider(riderId string) {
	checkContext(c)
	delete(c.riderDetails.driverOptionsForRider, riderId)
}

func checkContext(c *context) {
	if c == nil {
		panic("expected context to not be empty but it is empty")
	}

	if c.riderDetails == nil {
		panic("expected rider details in context to not be empty but it is empty")
	}

	if c.riderDetails.driverOptionsForRider == nil {
		panic("expected driver options in rider details in context to not be empty but it is empty")
	}
}
