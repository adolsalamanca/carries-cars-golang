package application_test

import (
	"testing"
	"time"

	pricingEngine "github.com/adolsalamanca/carries-cars-golang/internal/application"
	"github.com/adolsalamanca/carries-cars-golang/internal/domain"
)

func Test_RegularReserver_does_not_allow_exceeding_limit(t *testing.T) {
	reserver := pricingEngine.NewRegularReserver(time.Millisecond)

	reserver.Reserve()
	ticker := time.NewTicker(time.Millisecond * 2)
	<-ticker.C

	expected := pricingEngine.TimeExceededAfterReservationErr
	err := reserver.Start()
	if err != expected {
		t.Fatalf("expected %s err, obtained %s", expected, err)
	}
}

func Test_ExtendedReserver_does_allow_exceeding_limit(t *testing.T) {
	reserver := pricingEngine.NewExtendedReserver(time.Millisecond)

	reserver.Reserve()
	ticker := time.NewTicker(time.Millisecond * 2)

	<-ticker.C
	err := reserver.Start()
	if err != nil {
		t.Fatalf("expected nil err, obtained %s", err)
	}
}

func Test_CalculatePrice_charged_per_minute(t *testing.T) {
	pricePerMinute := domain.EUR(30)
	duration, _ := pricingEngine.DurationInMinutes(1)
	expected := domain.EUR(30)

	if !pricingEngine.CalculatePrice(pricePerMinute, duration).Equals(expected) {
		t.Fatalf("Price EUR(30) x 1min, want = EUR(30), have = EUR(%v)", expected.Amount())
	}
}

func Test_Duration_guards_against_zero_or_negative_duration(t *testing.T) {
	_, err := pricingEngine.DurationInMinutes(0)
	expected := "duration should be a positive number in minutes"

	if nil == err {
		t.Fatalf("DurationInMinutes(0), want = error(%q), have = nil", expected)
	}

	actual := err.Error()

	if expected != actual {
		t.Fatalf("DurationInMinutes(0), want = error(%q), have = error(%q)", expected, actual)
	}
}

func Test_UnverifiedDuration_Valid_Input(t *testing.T) {
	inMinutes := 1
	unverifiedInput := pricingEngine.UnverifiedDuration{DurationInMinutes: inMinutes}

	actual, _ := unverifiedInput.Verify()
	expected, _ := pricingEngine.DurationInMinutes(inMinutes)

	if expected != actual {
		t.Fatalf("UnverifiedDuration({DurationInMinutes: %v}), want = DurationInMinutes(%v), have = %T(%v)", inMinutes, expected, actual, actual)
	}
}

func Test_UnverifiedDuration_Invalid_Input(t *testing.T) {
	inMinutes := 0
	unverifiedInput := pricingEngine.UnverifiedDuration{DurationInMinutes: inMinutes}

	_, actual := unverifiedInput.Verify()
	expected := "duration should be a positive number in minutes"

	if nil == actual {
		t.Fatalf("UnverifiedDuration{DurationInMinutes: 0}.Verify(), want = error, have = nil")
	}

	if expected != actual.Error() {
		t.Fatalf("UnverifiedDuration{DurationInMinutes: 0}.Verify(), want = error(%q), have = error(%q)", expected, actual.Error())
	}
}
