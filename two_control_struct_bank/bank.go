package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const balanceFile = "Balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(balanceFile)
	if err != nil {
		fmt.Println("Error!!!")
		fmt.Println(err)
		fmt.Println("--------------------")
	}
	var choice int
	for choice != 4 {
		fmt.Println("Welcome to Gods bank")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")

		fmt.Scan(&choice)
		if choice == 1 {

			fmt.Printf("Your balance is %.2f\n", accountBalance)

		} else if choice == 2 {

			fmt.Print("How much do you want to deposit? ")
			var deposit float64
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Println("Invalid Amount")
			} else {
				accountBalance += deposit
				fmt.Printf("Your balance is now %.2f\n", accountBalance)
				fileops.WriteFloatToFile(accountBalance, balanceFile)
			}

		} else if choice == 3 {

			fmt.Print("How much do you want to withdraw? ")
			var withdraw float64
			fmt.Scan(&withdraw)
			if withdraw > accountBalance {
				fmt.Println("Insufficient funds")
			} else {
				accountBalance -= withdraw
				fmt.Printf("Your balance is now %.2f\n", accountBalance)
				fileops.WriteFloatToFile(accountBalance, balanceFile)
			}
		} else {
			fmt.Println("Thank you godding with us Nye!!!!!")
		}

	}

}
