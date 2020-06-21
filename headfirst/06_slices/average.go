// Average calcultes averages of numbers read in from a file
package main

import (
	"fmt"
	"log"

	"./datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	avg := mean(numbers...)
	fmt.Printf("Mean is: %0.4f\n", avg)
}

func mean(ns ...float64) float64 {
	var sum float64 = 0
	for _, n := range ns {
		sum += n
	}
	return sum / float64(len(ns))
}
