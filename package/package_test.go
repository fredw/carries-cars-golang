package _package_test

import (
	"testing"

	"carries-cars.com/money"
	pkg "carries-cars.com/package"
	"carries-cars.com/pricingEngine"
	"carries-cars.com/rental"
	"github.com/stretchr/testify/assert"
)

func Test_Package_basic_1hr_and_100km(t *testing.T) {
	packageBasic := pkg.Basic()
	trip := rental.NewRental(45, 90)
	finalPrice := pricingEngine.CalculateRentalPrice(packageBasic, trip)

	// trip = 0.24 EUR per minute
	// package basic = 0.20 EUR per minute

	expected := money.EUR(900)

	assert.Equal(t, expected, finalPrice)
}
