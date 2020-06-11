// average calculates the arithmetic mean of several numbers
package main

import "fmt"

func main() {
	numbers := [3]float64{12.3, 24.6, 36.9}
	var sum float64 = 0
	for _, n := range numbers {
		sum += n
	}
	mean := sum / float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", mean)
}
