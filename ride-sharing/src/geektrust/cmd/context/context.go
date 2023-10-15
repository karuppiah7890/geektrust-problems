package context

import "fmt"

type Context struct {
	riderDetails *riderDetails
}

type riderDetails struct {
	driverOptionsForRider map[string][]string
}

func NewContext() *Context {
	return &Context{
		riderDetails: &riderDetails{
			driverOptionsForRider: make(map[string][]string),
		},
	}
}

func (c *Context) GetDriverOptionsForRider(riderId string) ([]string, error) {
	checkContext(c)
	driverOptions, ok := c.riderDetails.driverOptionsForRider[riderId]
	if !ok {
		return nil, fmt.Errorf("driver options are unavailable for rider id %v: %w", riderId, ErrDriverOptionsUnavailable)
	}

	return driverOptions, nil
}

func (c *Context) StoreDriverOptionsForRider(riderId string, driverOptions []string) {
	checkContext(c)
	c.riderDetails.driverOptionsForRider[riderId] = driverOptions
}

func (c *Context) DeleteDriverOptionsForRider(riderId string) {
	checkContext(c)
	delete(c.riderDetails.driverOptionsForRider, riderId)
}

func checkContext(c *Context) {
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
