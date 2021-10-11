package domain_test

import (
	"testing"
	"time"

	"github.com/adolsalamanca/carries-cars-golang/internal/domain"
)

func Test_RegularReserver_excess_is_zero_also_for_exceeded_requests(t *testing.T) {
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

func Test_ExtendedReserver_allows_exceeding_limit(t *testing.T) {
	reserver := domain.NewExtendedReserver(time.Millisecond)

	reserver.Reserve()
	ticker := time.NewTicker(time.Millisecond * 2)

	<-ticker.C
	err := reserver.Start()
	if err != nil {
		t.Fatalf("expected nil err, obtained %s", err)
	}
}

func Test_ExtendedReserver_excess_returned_after_exceeding_limit_is_not_zero(t *testing.T) {
	reserver := domain.NewExtendedReserver(time.Millisecond)

	reserver.Reserve()
	ticker := time.NewTicker(time.Second * 1)

	<-ticker.C
	err := reserver.Start()
	if err != nil {
		t.Fatalf("expected nil err, obtained %s", err)
	}

	expected := 1.0
	excess := reserver.Excess()

	if expected != excess {
		t.Fatalf("expected %f excess, obtained %f", expected, excess)
	}

}
