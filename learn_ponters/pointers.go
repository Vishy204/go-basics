package main

import (
	"fmt"
)

func main() {
	age := 18
	agePointer := &age

	fmt.Println(adultYears(agePointer))
	fmt.Print(age)
}

func adultYears(age *int) int {
	*age = *age - 18
	return *age
}
