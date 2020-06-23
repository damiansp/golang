package main

import (
	"fmt"
	"math"
)

func maximum(ns ...float64) float64 {
	max := math.Inf(-1)
	for _, n := range ns {
		if n > max {
			max = n
		}
	}
	return max
}

func inRange(min float64, max float64, ns ...float64) []float64 {
	var result []float64
	for _, n := range ns {
		if n >= min && n <= max {
			result = append(result, n)
		}
	}
	return result
}

func main() {
	fmt.Println(maximum(77, 55, 99, 2))
	fmt.Println(inRange(-10, 10, 1.1, 8.8, 12.2, -9, -14.4))
}
