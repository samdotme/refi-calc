package countdowncalc

// Calculates a total paydown period in months.
func CalculatePaydownPeriod(balance float32, monthlyPayment float32, interest float32) int32 {
	var paydownPeriod int32
	remainingBalance := balance

	for remainingBalance > 0 {
		balanceWithInterest := (remainingBalance + remainingBalance*interest)

		remainingBalance = balanceWithInterest
		remainingBalance -= monthlyPayment
		paydownPeriod++
	}

	return paydownPeriod
}
