package pkg

import (
	"fmt"
	"geektrust/pkg/ride"
)

type CalculateBillInput struct {
	RideId string
}

type Bill struct {
	DriverId string
	Amount   float64
}

const BASE_FARE = 50
const COST_PER_KM = 6.5
const COST_PER_MINUTE = 2
const SERVICE_TAX_PERCENTAGE = 0.2

func (r *RideSharingApp) CalculateBill(input *CalculateBillInput) (*Bill, error) {
	rideId := input.RideId

	// if ride id does not exist, return error
	ride, ok := r.GetRide(rideId)
	if !ok {
		return nil, fmt.Errorf("a ride with id %s does not exist: %w", rideId, ErrRideIdNotExist)
	}

	if !ride.IsComplete() {
		return nil, fmt.Errorf("ride with id %s is not completed: %w", rideId, ErrRideNotCompleted)
	}

	amount := calculateBillAmount(ride)

	bill := &Bill{
		Amount:   amount,
		DriverId: ride.GetDriverId(),
	}

	return bill, nil
}

func calculateBillAmount(ride *ride.Ride) float64 {
	source := ride.GetSource()
	destination := ride.GetDestination()

	distanceFee := COST_PER_KM * destination.DistanceBetween(source)

	durationFee := COST_PER_MINUTE * ride.GetRideDurationInMinutes()

	billAmount := BASE_FARE + distanceFee + float64(durationFee)

	tax := SERVICE_TAX_PERCENTAGE * billAmount

	totalBillAmountWithTax := billAmount + tax

	return totalBillAmountWithTax
}
