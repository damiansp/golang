package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep.  Boop.")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
}

func main() {
	var toy NoiseMaker
	toy = Whistle("Tin Lizzy")
	toy.MakeSound()
	play(toy)
	toy = Horn("Klaxomatic")
	toy.MakeSound()
	play(toy)
	var thingy NoiseMaker = Robot("Robbie")
	thingy.MakeSound()
	var robit Robot = thingy.(Robot) // convert to concrete Robot type
	robit.Walk()
}
