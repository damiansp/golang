package main

import "fmt"

func main() {
	quantity := 4
	len, width := 1.2, 2.4
	customerName := "Philip Damien"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(len*width, "sq. meters")
}
