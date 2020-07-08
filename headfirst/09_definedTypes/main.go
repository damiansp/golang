package main

import "fmt"

type Liters float64
type Gallons float64

func main() {
	var usFuel Gallons
	usFuel = Gallons(10.)
	euroFuel := Liters(240.)
	fmt.Println(usFuel, euroFuel)
}
