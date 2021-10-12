package domain

import (
	"errors"
	"time"
)

var (
	TimeExceededAfterReservationErr = errors.New("exceeded time spent between reservation and car start")
)

type Reserver interface {
	StartAfter(time.Duration) error
	ExcessInMinutes() time.Duration
}

type RegularReserver struct {
	startTime    time.Time
	limitMinutes time.Duration
}

func NewRegularReserver(limit time.Duration) *RegularReserver {
	return &RegularReserver{
		limitMinutes: limit,
	}
}

func (r *RegularReserver) ExcessInMinutes() time.Duration {
	return 0
}

func (r *RegularReserver) StartAfter(timeSinceReservation time.Duration) error {
	if timeSinceReservation.Minutes() > r.limitMinutes.Minutes() {
		return TimeExceededAfterReservationErr
	}
	return nil
}

type ExtendedReserver struct {
	startTime       time.Time
	limitInMinutes  time.Duration
	excessInMinutes time.Duration
}

func (r *ExtendedReserver) ExcessInMinutes() time.Duration {
	return r.excessInMinutes - r.limitInMinutes
}

func NewExtendedReserver(limit time.Duration) *ExtendedReserver {
	return &ExtendedReserver{
		limitInMinutes: limit,
	}
}

func (r *ExtendedReserver) StartAfter(timeSinceReservation time.Duration) error {
	r.excessInMinutes = timeSinceReservation

	return nil
}
