package main

import (
	"fmt"
	"math"
)

const InflationRate = 2.5

func main() {
	const InflationRate = 2.5
	var InvestmentAmount float64
	ExpectedReturnRate := 5.5
	var years float64

	fmt.Print("Enter the Investment Amount: ")
	fmt.Scan(&InvestmentAmount)

	fmt.Print("Enter the Expected Return Rate: ")
	fmt.Scan(&ExpectedReturnRate)

	fmt.Print("Enter the years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := caculationInvestment(InvestmentAmount, ExpectedReturnRate, years)
	printCalulation(futureValue, futureRealValue)
}
func caculationInvestment(InvestmentAmount, ExpectedReturnRate, years float64) (fv float64, frv float64) {
	fv = InvestmentAmount * math.Pow(1+ExpectedReturnRate/100, years)
	frv = fv / math.Pow(1+InflationRate/100, years)
	//	return fv, frv
	return
}
func printCalulation(futureValue, futureRealValue float64) {
	fmt.Printf("The future values is %.1f", futureValue)
	fmt.Printf("\nThe future value after inflation is %.1f", futureRealValue)
}
