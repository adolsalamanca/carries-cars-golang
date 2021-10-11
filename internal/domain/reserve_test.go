package domain_test

import (
	"testing"
	"time"

	"github.com/adolsalamanca/carries-cars-golang/internal/domain"
)

func Test_RegularReserver_does_not_allow_exceeding_limit(t *testing.T) {
	reserver := domain.NewRegularReserver(time.Millisecond)

	reserver.Reserve()
	ticker := time.NewTicker(time.Millisecond * 2)
	<-ticker.C

	expected := domain.TimeExceededAfterReservationErr
	err := reserver.Start()
	if err != expected {
		t.Fatalf("expected %s err, obtained %s", expected, err)
	}
}

func Test_ExtendedReserver_does_allow_exceeding_limit(t *testing.T) {
	reserver := domain.NewExtendedReserver(time.Millisecond)

	reserver.Reserve()
	ticker := time.NewTicker(time.Millisecond * 2)

	<-ticker.C
	err := reserver.Start()
	if err != nil {
		t.Fatalf("expected nil err, obtained %s", err)
	}
}
