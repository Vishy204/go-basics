package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	var interestRate float64
	var years float64
	const inflationRate = 7.5

	outputText("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText.Print("Enter Interest Rate: ")
	fmt.Scan(&interestRate)

	outputText.Print("Enter number of years:")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, interestRate, years)

	//futureValue := investmentAmount * math.Pow(1+interestRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future value (adjusted after inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)

	//fmt.Println("Future Value: ",futureValue)

	//fmt.Printf("Future Value: %.2f\n",futureValue)
	//fmt.Println("Future Value (adjusted after inflation): "futureRealValue)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, interestRate, years float64) float64.float64 {
	fv := investmentAmount * math.Pow(1+interestRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)

	return fv, frv
}
