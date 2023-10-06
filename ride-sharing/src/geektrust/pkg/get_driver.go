package pkg

func (r *RideSharingApp) GetDriver(driverId string) (*Driver, bool) {
	driver, ok := r.drivers[driverId]

	return driver, ok
}
