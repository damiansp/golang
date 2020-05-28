// getPaint calculates paint needed by area
package main

import "fmt"

func main() {
	getPaintNeeded(4.2, 3.0)
	getPaintNeeded(5.2, 3.5)
	getPaintNeeded(5.5, 3.3)
}

func getPaintNeeded(w, h float64) {
	area := w * h
	liters := area / 10
	fmt.Printf("%.2f liters needed\n", liters)
}
