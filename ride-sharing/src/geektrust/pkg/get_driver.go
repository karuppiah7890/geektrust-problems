package pkg

import "fmt"

func (r *RideSharingApp) GetDriver(driverId string) (*Driver, error) {
	driver, ok := r.drivers[driverId]
	if !ok {
		return nil, fmt.Errorf("driver with driver id %s does not exist", driverId)
	}

	return driver, nil
}
