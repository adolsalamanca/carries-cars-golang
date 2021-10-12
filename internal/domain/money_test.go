package domain_test

import (
	"github.com/adolsalamanca/carries-cars-golang/internal/domain"

	"testing"
)

func Test_Money_Equals_detects_equal_values(t *testing.T) {
	actual := domain.EUR(99).Equals(domain.EUR(99))
	expected := true

	if actual != expected {
		t.Fatalf("EUR(99).Equals(EUR(99)) want = %t, have = %t", expected, actual)
	}
}

func Test_Money_Equals_detects_currency_differences(t *testing.T) {
	actual := domain.EUR(10).Equals(domain.USD(10))
	expected := false

	if actual != expected {
		t.Fatalf("EUR(10).Equals(USD(10)) want = %t, have = %t", expected, actual)
	}
}

func Test_Money_Equals_detects_amount_differences(t *testing.T) {
	actual := domain.EUR(1).Equals(domain.EUR(2))
	expected := false

	if actual != expected {
		t.Fatalf("EUR(1).Equals(EUR(2)) want = %t, have = %t", expected, actual)
	}
}

func Test_Money_Multiply_multiplies(t *testing.T) {
	actual := domain.EUR(200).MultiplyAndRound(2.00)
	expected := domain.EUR(400)

	if actual != expected {
		t.Fatalf("EUR(200).MultiplyAndRound(2.00) want = EUR(%v), have = EUR(%v)", expected.Amount(), actual.Amount())
	}
}

func Test_Money_Add_Same_Currency(t *testing.T) {
	actual := domain.EUR(200)
	got, err := actual.Add(domain.EUR(200))
	if err != nil {
		t.Fatalf("expected nil err, obtained %s", err)
	}

	expected := domain.EUR(400)

	if got != expected {
		t.Fatalf("EUR(200).MultiplyAndRound(2.00) want = EUR(%v), have = EUR(%v)", expected.Amount(), got.Amount())
	}
}

func Test_Money_Add_Different_Currency(t *testing.T) {
	actual := domain.EUR(200)
	_, err := actual.Add(domain.USD(200))
	if err == nil {
		t.Fatal("expected not nil err, but no error was obtained")
	}
}

func Test_Money_Multiply_rounds_upward_correctly(t *testing.T) {
	actual := domain.EUR(100).MultiplyAndRound(1.999)
	expected := domain.EUR(200)

	if actual != expected {
		t.Fatalf("EUR(100).MultiplyAndRound(1.999) want = EUR(%v), have = EUR(%v)", expected.Amount(), actual.Amount())
	}
}

func Test_Money_Multiply_rounds_downward_correctly(t *testing.T) {
	actual := domain.EUR(100).MultiplyAndRound(1.994)
	expected := domain.EUR(199)

	if actual != expected {
		t.Fatalf("EUR(100).MultiplyAndRound(1.994) want = EUR(%v), have = EUR(%v)", expected.Amount(), actual.Amount())
	}
}

func Test_Money_Amount_exposes_value(t *testing.T) {
	t.Skip("Todo")
}

func Test_Money_CurrencyIsoCode_exposes_value(t *testing.T) {
	t.Skip("Todo")
}
