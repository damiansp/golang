package main

import "fmt"

type Appliance interface {
	TurnOn()
}

type Fan string

func (f Fan) TurnOn() {
	fmt.Println("Blowing")
}

type CoffeePot string

func (c CoffeePot) TurnOn() {
	fmt.Println("Warming up")
}

func (c CoffeePot) Brew() {
	fmt.Println("Brewing")
}

func main() {
	var device Appliance
	device = Fan("Mr. Breezy")
	device.TurnOn()
	device = CoffeePot("Dale E. Grind")
	device.TurnOn()
}
