package countdowncalc

// Calculates a total paydown period in months.
func CalculatePaydownPeriod(balance float64, monthlyPayment float64, interest float64) int32 {
	var paydownPeriod int32
	remainingBalance := balance
	monthlyInterest := interest / 12

	for remainingBalance > 0 {
		balanceWithInterest := (remainingBalance + remainingBalance*monthlyInterest)

		remainingBalance = balanceWithInterest
		remainingBalance -= monthlyPayment
		paydownPeriod++
	}

	return paydownPeriod
}
