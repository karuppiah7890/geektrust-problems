package pkg

import "geektrust/pkg/driver"

func (r *RideSharingApp) GetDriver(driverId string) (*driver.Driver, bool) {
	d, ok := r.drivers[driverId]

	return d, ok
}
