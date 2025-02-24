package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"Carpet"}
	prices := [4]float64{10.5, 20.2, 40.7, 5.3}
	fmt.Println(prices)

	productNames[2] = "Toothpaste"
	fmt.Println(productNames[2])

	featuredPrices := prices[1:]
	featuredPrices[0] = 199.99
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(prices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

}
