package main

import "fmt"

type Liters float64
type Gallons float64

func main() {
	var usFuel Gallons
	usFuel = Gallons(1.2)
	euroFuel := Liters(4.5)
	usFuel += liters2Gallons(Liters(40.))
	euroFuel += gallons2Liters(Gallons(30.))
	fmt.Printf("US Fuel: %0.1f gallons\n", usFuel)
	fmt.Printf("Euro Fuel: %0.1f liters\n", euroFuel)
}

func liters2Gallons(liters Liters) Gallons {
	return Gallons(liters * 0.264)
}

func gallons2Liters(gallons Gallons) Liters {
	return Liters(gallons * 3.785)
}
