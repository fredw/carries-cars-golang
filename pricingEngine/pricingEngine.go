package pricingEngine

import (
	"errors"

	"carries-cars.com/money"
)

type Duration interface {
	DurationInMinutes() int
}

// UnverifiedDuration should be used when accepting input from untrusted sources (pretty much anywhere) in the model.
// This type models input that has not been verified and is therefore unsafe to use until it has been verified.
// Use Verify() to transform it to trusted input in the form of a duration model.
type UnverifiedDuration struct {
	DurationInMinutes int
}

func (unsafe UnverifiedDuration) Verify() (Duration, error) {
	return DurationInMinutes(unsafe.DurationInMinutes)
}

func DurationInMinutes(durationInMinutes int) (Duration, error) {
	if durationInMinutes <= 0 {
		defaultDuration := duration{durationInMinutes: 1}

		return defaultDuration, errors.New("duration should be a positive number in minutes")
	}

	return duration{durationInMinutes: durationInMinutes}, nil
}

type duration struct {
	durationInMinutes int
}

func (duration duration) DurationInMinutes() int {
	return duration.durationInMinutes
}

func CalculatePrice(pricePerMinute money.Money, duration Duration, distanceInKM int) money.Money {
	durationInMinutes := float64(duration.DurationInMinutes())

	return pricePerMinute.MultiplyAndRound(durationInMinutes)
}

const reservationTimeLimitInMinutes = 20

func CalculateReservationPrice(duration Duration) money.Money {
	durationInMinutes := float64(duration.DurationInMinutes())
	extraTimeSurcharge := money.EUR(0)
	if durationInMinutes > reservationTimeLimitInMinutes {
		extraTimeSurcharge = money.EUR(9)
		durationInMinutes -= reservationTimeLimitInMinutes
	}

	return extraTimeSurcharge.MultiplyAndRound(durationInMinutes)
}
