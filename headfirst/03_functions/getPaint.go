// getPaint calculates paint needed by area
package main

import (
	"fmt"
	"log"
)

func main() {
	var amount float64
	amount, err := getPaintNeeded(4.2, 3.0)
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed\n", amount)
}

func getPaintNeeded(w, h float64) (float64, error) {
	if w < 0 || h < 0 {
		return 0, fmt.Errorf("Width and height must be > 0")
	}
	area := w * h
	liters := area / 10
	return liters, nil // nil for error
}
