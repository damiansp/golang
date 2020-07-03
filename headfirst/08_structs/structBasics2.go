package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	subscriber1 := makeDefaultSubscriber("Aman Singh")
	subscriber1.rate = 4.99
	printInfo(subscriber1)
	subscriber2 := makeDefaultSubscriber("Beth Ryan")
	applyDiscount(&subscriber2)
	printInfo(subscriber2)
}

func printInfo(s subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active:", s.active)
}

func makeDefaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}
