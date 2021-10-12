package application_test

import (
	"testing"
	"time"

	pricingEngine "github.com/adolsalamanca/carries-cars-golang/internal/application"
	"github.com/adolsalamanca/carries-cars-golang/internal/domain"
)

func TestPricingEngine_With_RegularReserver_CalculatesPrice(t *testing.T) {
	pricePerMinute := domain.EUR(30)
	duration, _ := pricingEngine.DurationInMinutes(1)
	reserver := domain.NewRegularReserver(time.Minute)
	engine := pricingEngine.NewPricingEngine(reserver, domain.EUR(9))

	err := engine.StartAfter(0)
	if err != nil {
		t.Fatalf("expected nil err, obtained %s", err)
	}

	expected := domain.EUR(30)
	if got := engine.CalculatePrice(pricePerMinute, duration); !got.Equals(expected) {
		t.Fatalf("Expected %v, but have %v", expected.Amount(), got.Amount())
	}
}

func TestPricingEngine_With_ExtendedReserver_without_exceed_limit_CalculatesPrice(t *testing.T) {
	pricePerMinute := domain.EUR(30)
	duration, _ := pricingEngine.DurationInMinutes(1)
	reserver := domain.NewExtendedReserver(time.Millisecond)
	engine := pricingEngine.NewPricingEngine(reserver, domain.EUR(9))

	err := engine.StartAfter(0)
	if err != nil {
		t.Fatalf("expected nil err, obtained %s", err)
	}

	expected := domain.EUR(30)
	if got := engine.CalculatePrice(pricePerMinute, duration); !got.Equals(expected) {
		t.Fatalf("Expected %v, but have %v", expected.Amount(), got.Amount())
	}
}

func TestPricingEngine_With_ExtendedReserver_and_lasted_more_than_limit_CalculatesPrice(t *testing.T) {
	pricePerMinute := domain.EUR(30)
	duration, _ := pricingEngine.DurationInMinutes(1)
	reserver := domain.NewExtendedReserver(time.Millisecond)
	engine := pricingEngine.NewPricingEngine(reserver, domain.EUR(9))

	err := engine.StartAfter(time.Minute)
	if err != nil {
		t.Fatalf("expected nil err, obtained %s", err)
	}

	expected := domain.EUR(39)
	if got := engine.CalculatePrice(pricePerMinute, duration); !got.Equals(expected) {
		t.Fatalf("Expected %v, but have %v", expected.Amount(), got.Amount())
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
