package domain

import (
	"errors"
	"math"
	"time"
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
		limit: limit,
	}
}

func (r *RegularReserver) Reserve() {
	r.startTime = time.Now()
}

func (r *RegularReserver) Start() error {
	if time.Since(r.startTime) > r.limit {
		return TimeExceededAfterReservationErr
	}
	return nil
}

type ExtendedReserver struct {
	startTime       time.Time
	limit           time.Duration
	excessInSeconds float64
}

func NewExtendedReserver(limit time.Duration) ExtendedReserver {
	return ExtendedReserver{
		limit: limit,
	}
}

func (r *ExtendedReserver) Reserve() {
	r.startTime = time.Now()
}

func (r *ExtendedReserver) Start() error {
	if startTime := time.Since(r.startTime); startTime > r.limit {
		r.excessInSeconds = math.Round(startTime.Seconds())
	}
	return nil
}
