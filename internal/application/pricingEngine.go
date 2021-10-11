package application

import (
	"errors"
	"time"

	"github.com/adolsalamanca/carries-cars-golang/internal/domain"
)

var (
	TimeExceededAfterReservationErr = errors.New("exceeded time spent between reservation and car start")
)

type Reserver interface {
	Reserve()
	Start() error
}

type RegularReserver struct {
	startTime time.Time
	limit     time.Duration
}

func NewRegularReserver(limit time.Duration) RegularReserver {
	return RegularReserver{
		startTime: time.Now(),
		limit:     limit,
	}
}

func (r RegularReserver) Reserve() {
	panic("implement me")
}

func (r RegularReserver) Start() error {
	panic("implement me")
}

type ExtendedReserver struct {
	startTime time.Time
	limit     time.Duration
}

func NewExtendedReserver(limit time.Duration) ExtendedReserver {
	return ExtendedReserver{
		startTime: time.Now(),
		limit:     limit,
	}
}

func (e ExtendedReserver) Reserve() {
	panic("implement me")
}

func (e ExtendedReserver) Start() error {
	panic("implement me")
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
