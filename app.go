package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {

	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	userInput("Investment Amount: ", &investmentAmount)
	userInput("Years: ", &years)
	userInput("Expected Return Rate: ", &expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	_ = futureValue

	fmt.Printf("Future Real Value: %.1f ($)\n", futureRealValue)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

func userInput(str string, val *float64) {
	fmt.Print(str)
	fmt.Scan(val)
}
