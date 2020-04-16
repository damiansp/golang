// Prints the boiling point of water
package main

import "fmt"

const boilingF = 212.

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Boiling point: %gC (%g°F)\n", c, f)
}
