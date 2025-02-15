package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(firstName, lastName, birthdate)

	if err != nil {
		fmt.Print(err)
		return

	}

	// ... do something awesome with that gathered data!

	appUser.DelNames()
	appUser.SploutputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
