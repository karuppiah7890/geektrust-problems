package main

type DriverOptionsUnavailableForRider string

func (d DriverOptionsUnavailableForRider) Error() string {
	return string(d)
}
