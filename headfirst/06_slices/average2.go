// average2 calculates the mean of several numbers
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	var sum float64 = 0
	for _, arg := range args {
		n, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += n
	}
	mean := sum / float64(len(args))
	fmt.Printf("Mean: %0.2f\n", mean)
}
