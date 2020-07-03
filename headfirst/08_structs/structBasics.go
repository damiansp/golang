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

	bolts := minimumOrder("Hex Bolts")
	showInfo(bolts)
}

func minimumOrder(descr string) part {
	var p part
	p.descr = descr
	p.count = 24
	return p
}

func showInfo(p part) {
	fmt.Println("Description:", p.descr)
	fmt.Println("Count:", p.count)
}
