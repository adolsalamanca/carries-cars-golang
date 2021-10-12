package application

import (
	"errors"
	"log"

	"github.com/adolsalamanca/carries-cars-golang/internal/domain"
)

type Calculator interface {
	CalculatePrice(pricePerMinute domain.Money, duration Duration) domain.Money
}

type PricingEngine struct {
	domain.Reserver
	pricePerExcessMinute domain.Money
}

func (p PricingEngine) CalculatePrice(pricePerMinute domain.Money, duration Duration) domain.Money {
	tripDurationInMinutes := float64(duration.DurationInMinutes())

	exceededReservationMinutes := p.ExcessInMinutes()
	if exceededReservationMinutes != 0 {
		excessPrice := p.pricePerExcessMinute.MultiplyAndRound(exceededReservationMinutes.Minutes())

		price, err := pricePerMinute.MultiplyAndRound(tripDurationInMinutes).Add(excessPrice)
		if err != nil {
			log.Fatalf("could not calculate price, %v", err)
		}
		return price
	}

	return pricePerMinute.MultiplyAndRound(tripDurationInMinutes)
}

func NewPricingEngine(reserver domain.Reserver, pricePerExcessMinute domain.Money) PricingEngine {
	return PricingEngine{
		Reserver:             reserver,
		pricePerExcessMinute: pricePerExcessMinute,
	}
}

// UnverifiedDuration should be used when accepting input from untrusted sources (pretty much anywhere) in the model.
// This type models input that has not been verified and is therefore unsafe to use until it has been verified.
// Use Verify() to transform it to trusted input in the form of a duration model.
type UnverifiedDuration struct {
	DurationInMinutes int
}

func (unsafe UnverifiedDuration) Verify() (Duration, error) {
	return DurationInMinutes(unsafe.DurationInMinutes)
}

type Duration interface {
	DurationInMinutes() int
}

func DurationInMinutes(durationInMinutes int) (Duration, error) {
	if durationInMinutes <= 0 {
		defaultDuration := duration{durationInMinutes: 1}

		return defaultDuration, errors.New("duration should be a positive number in minutes")
	}

	return duration{durationInMinutes: durationInMinutes}, nil
}

func (duration duration) DurationInMinutes() int {
	return duration.durationInMinutes
}

type duration struct {
	durationInMinutes int
}
