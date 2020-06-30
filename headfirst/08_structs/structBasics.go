package main

import "fmt"

type part struct {
	descr string
	count int
}

type car struct {
	name     string
	topSpeed float64
}

func main() {
	var porsche car
	porsche.name = "Porsche Targa 911"
	porsche.topSpeed = 323
	fmt.Println("Name:", porsche.name)
	fmt.Println("Top speed:", porsche.topSpeed)

	var bolts part
	bolts.descr = "Hex bolts"
	bolts.count = 24
	fmt.Println("Description:", bolts.descr)
	fmt.Println("Count:", bolts.count)
}
