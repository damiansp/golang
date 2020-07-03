package main

import (
	"fmt"

	"./magazine"
)

func main() {
	var s magazine.Subscriber
	s.rate = 4.99
	fmt.Print(s.rate)
}
