package main

import "fmt"

func main() {
	var price int = 100
	var taxRate float64 = 0.08
	var tax float64 = taxRate * float64(price)
	var total float64 = float64(price) + tax
	var availableFunds int = 120
	fmt.Println("Price is $", price)
	fmt.Println("Tax is $", tax)
	fmt.Println("Total cost is $", total)
	fmt.Println("Available: $", availableFunds)
	fmt.Println("Within budget:", float64(availableFunds) >= total)
}
