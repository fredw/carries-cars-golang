package _package

import (
	"carries-cars.com/money"
)

type Package interface {
	Price() money.Money
}

type rentalPackage struct {
	minutes        int
	maximumMileage int
	price          money.Money
}

func (rp rentalPackage) Price() money.Money {
	return rp.price
}

func Basic() Package {
	return rentalPackage{
		minutes:        60,
		maximumMileage: 100,
		price:          money.EUR(20),
	}
}
