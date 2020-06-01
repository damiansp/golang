package main

import "fmt"

func main() {
	amount := 6
	double(&amount)
	fmt.Println(amount)
}

func double(n *int) {
	*n *= 2
}
