package domain

import (
	"errors"
	"math"
)

type CurrencyIsoCode string

type Money interface {
	Add(moneyToBeAdded Money) (Money, error)

	// Amount is denoted in the lowest denominator of the corresponding currency.
	// E.g. amount is in whole cents for the Euro or UnitedStatesDollar
	Amount() int

	CurrencyIsoCode() CurrencyIsoCode

	MultiplyAndRound(multiplier float64) Money

	Equals(other Money) bool
}

// EUR acts as a named constructor function to create trustedMoney for the Euro currency.
// Provide the amount in cents.
func EUR(amount int) Money {
	return trustedMoney{
		amount:          amount,
		currencyIsoCode: Euro,
	}
}

// USD acts as a named constructor function to create trustedMoney for the UnitedStatesDollar currency.
// Provide the amount in cents.
func USD(amount int) Money {
	return trustedMoney{
		amount:          amount,
		currencyIsoCode: UnitedStatesDollar,
	}
}

func (money trustedMoney) Equals(other Money) bool {
	return money.amount == other.Amount() && money.currencyIsoCode == other.CurrencyIsoCode()
}

func (money trustedMoney) MultiplyAndRound(multiplier float64) Money {
	multipliedAmount := float64(money.amount) * multiplier
	multipliedAmountRounded := int(math.Round(multipliedAmount))

	return trustedMoney{
		amount:          multipliedAmountRounded,
		currencyIsoCode: money.currencyIsoCode,
	}
}

const Euro = "EUR"
const UnitedStatesDollar = "USD"

// Internals and boring (machine-generated) code below

// trustedMoney is hidden from the API surface to ensure that this type is trustworthy because it can only be created
// through one of the named constructors (EUR() or USD()).
type trustedMoney struct {
	// amount is denoted in the lowest denominator of the corresponding currency.
	// E.g. amount is in whole cents for the Euro or UnitedStatesDollar
	amount          int
	currencyIsoCode CurrencyIsoCode
}

func (money trustedMoney) Add(moneyToBeAdded Money) (Money, error) {
	if money.currencyIsoCode != moneyToBeAdded.CurrencyIsoCode() {
		return nil, errors.New("could not add different currency money")
	}

	return trustedMoney{
		amount:          money.amount + moneyToBeAdded.Amount(),
		currencyIsoCode: money.currencyIsoCode,
	}, nil
}

// Boring (machine-generated) code below
func (money trustedMoney) CurrencyIsoCode() CurrencyIsoCode {
	return money.currencyIsoCode
}

func (money trustedMoney) Amount() int {
	return money.amount
}
