package application

import (
	"errors"

	"github.com/adolsalamanca/carries-cars-golang/internal/domain"
)

// UnverifiedDuration should be used when accepting input from untrusted sources (pretty much anywhere) in the model.
// This type models input that has not been verified and is therefore unsafe to use until it has been verified.
// Use Verify() to transform it to trusted input in the form of a duration model.
type UnverifiedDuration struct {
	DurationInMinutes int
}

func (unsafe UnverifiedDuration) Verify() (Duration, error) {
	return DurationInMinutes(unsafe.DurationInMinutes)
}

func CalculatePrice(pricePerMinute domain.Money, duration Duration) domain.Money {
	durationInMinutes := float64(duration.DurationInMinutes())

	return pricePerMinute.MultiplyAndRound(durationInMinutes)
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
