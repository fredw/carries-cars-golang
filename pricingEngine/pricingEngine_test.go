package pricingEngine_test

import (
	"testing"

	"carries-cars.com/money"
	pricingEngine "carries-cars.com/pricingEngine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CalculatePrice_charged_per_minute(t *testing.T) {
	pricePerMinute := money.EUR(30)
	duration, _ := pricingEngine.DurationInMinutes(1)
	expected := money.EUR(30)

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

func Test_CalculatePrice_reservation(t *testing.T) {
	duration, err := pricingEngine.DurationInMinutes(15)
	require.NoError(t, err, "failed to calculate the duration")

	got := pricingEngine.CalculateReservationPrice(duration)
	expected := money.EUR(0)

	assert.Equal(t, expected, got)
}
