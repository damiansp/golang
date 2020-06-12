// average calculates the arithmetic mean of several numbers
package main

import (
	"fmt"
	"log"

	"../datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, n := range numbers {
		sum += n
	}
	mean := sum / float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", mean)
}
