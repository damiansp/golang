package main

import (
	"fmt"
	"math"
)

func main() {
	root, err := squareRoot(-9.3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.3f", root)
	}
}

func squareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("Cannot get the square root of a negative")
	}
	return math.Sqrt(n), nil
}
