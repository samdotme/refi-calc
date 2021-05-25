package countdowncalc

import (
	"testing"
)

func TestPaydownIsZeroWhenBalanceIsZero(t *testing.T) {
	months := CalculatePaydownPeriod(0, 0, 0)

	if months != 0 {
		t.Fatalf(`CaculatePaydownPeriod() should return 0 when balance is 0. Instead returned %v`, months)
	}
}

func TestPaydownIsOneWhenBalanceIsSameAsPaydownRateWithNoInterest(t *testing.T) {
	months := CalculatePaydownPeriod(500, 500, 0)

	if months != 1 {
		t.Fatalf(`Failed to calculate 1 month paydown period. Instead returned %v`, months)
	}
}

func TestPaydownIsTwoWhenBalanceIsSameAsPaydownRateWithInterest(t *testing.T) {
	months := CalculatePaydownPeriod(500, 500, .04125)

	if months != 2 {
		t.Fatalf(`Failed to calculate 2 month paydown period due to interest. Instead returned %v`, months)
	}
}

func TestInterestCompounds(t *testing.T) {
	months := CalculatePaydownPeriod(100, 30, .1)

	if months != 5 {
		t.Fatalf(`Failed to take into account compounding of interest. Expected 12, returned %v`, months)
	}
}
