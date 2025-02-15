package main

import (
	"errors"
	"fmt"
	"os"
)

func getInput(a string) (float64, error) {
	var input float64
	fmt.Print(a)
	fmt.Scan(&input)
	if input <= 0 {
		return input, errors.New("Invalid input value is less than or equal to 0")
	}
	return input, nil
}

func calculateFunc(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := revenue*(1-taxRate/100) - expenses
	ratio := profit / ebt
	return ebt, profit, ratio
}

func writeInFile(ebt float64, profit float64, ratio float64) {
	text := fmt.Sprint(ebt, ratio, profit)
	os.WriteFile("Profit_sys.txt", []byte(text), 0644)
}

func main() {
	revenue, err1 := getInput("Enter revenue: ")
	expenses, err2 := getInput("Enter expenses:  ")
	taxRate, err3 := getInput("Enter taxRate: ")
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Error!!!")
		fmt.Println(err1)
		fmt.Println(err2)
		fmt.Println(err3)
		panic("Program terminated because of invalid input!!!!")
	}

	ebt, profit, ratio := calculateFunc(revenue, expenses, taxRate)
	writeInFile(ebt, profit, ratio)

	fmt.Printf("Earnings Before Tax: %.2f\n", ebt)
	fmt.Printf("Earnings After Tax: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f", ratio)

}
