package domain_test

import (
	"testing"
	"time"

	"github.com/adolsalamanca/carries-cars-golang/internal/domain"
)

func Test_RegularReserver_start_fails_for_exceeded_requests(t *testing.T) {
	reserver := domain.NewRegularReserver(time.Minute * 20)
	expected := domain.TimeExceededAfterReservationErr

	err := reserver.StartAfter(time.Minute * 30)

	if err != expected {
		t.Fatalf("expected %s err, obtained %s", expected, err)
	}
}

func Test_ExtendedReserver_allows_exceeding_limit(t *testing.T) {
	reserver := domain.NewExtendedReserver(time.Minute * 20)

	err := reserver.StartAfter(time.Minute * 25)
	if err != nil {
		t.Fatalf("expected nil err, obtained %s", err)
	}
}

func Test_ExtendedReserver_excess_returned_after_exceeding_limit_is_not_zero(t *testing.T) {
	reserver := domain.NewExtendedReserver(time.Minute * 20)

	err := reserver.StartAfter(time.Minute * 25)
	if err != nil {
		t.Fatalf("expected nil err, obtained %s", err)
	}

	expected := time.Minute * 5
	excess := reserver.ExcessInMinutes()

	if excess.Minutes() != expected.Minutes() {
		t.Fatalf("expected %f excess, obtained %f", expected.Minutes(), excess.Minutes())
	}

}
