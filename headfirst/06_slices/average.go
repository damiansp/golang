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
	var sum float64 = 0
	for _, n := range numbers {
		sum += n
	}
	mean := sum / float64(len(numbers))
	fmt.Printf("Mean is: %0.4f\n", mean)
}
