package money

import "math"

type CurrencyIsoCode string

type Money interface {
	// Amount is denoted in the lowest denominator of the corresponding currency.
	// E.g. amount is in whole cents for the Euro or UnitedStatesDollar
	Amount() int
	CurrencyIsoCode() CurrencyIsoCode
	MultiplyAndRound(multiplier float64) Money
	Add(amount int) Money
	Equals(other Money) bool
}

// trustedMoney is hidden from the API surface to ensure that this type is trustworthy because it can only be created
// through one of the named constructors (EUR() or USD()).
type trustedMoney struct {
	// amount is denoted in the lowest denominator of the corresponding currency.
	// E.g. amount is in whole cents for the Euro or UnitedStatesDollar
	amount          int
	currencyIsoCode CurrencyIsoCode
}

func (money trustedMoney) Amount() int {
	return money.amount
}

func (money trustedMoney) CurrencyIsoCode() CurrencyIsoCode {
	return money.currencyIsoCode
}

func (money trustedMoney) MultiplyAndRound(multiplier float64) Money {
	multipliedAmount := float64(money.amount) * multiplier
	multipliedAmountRounded := int(math.Round(multipliedAmount))

	return trustedMoney{
		amount:          multipliedAmountRounded,
		currencyIsoCode: money.currencyIsoCode,
	}
}

func (money trustedMoney) Add(amount int) Money {
	return trustedMoney{
		amount:          money.amount + amount,
		currencyIsoCode: money.currencyIsoCode,
	}
}

func (money trustedMoney) Equals(other Money) bool {
	return money.amount == other.Amount() && money.currencyIsoCode == other.CurrencyIsoCode()
}

const Euro = "EUR"
const UnitedStatesDollar = "USD"

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
