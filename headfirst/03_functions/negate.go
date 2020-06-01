package main

import "fmt"

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lie := false
	negate(&lie)
	fmt.Println(lie)
}

func negate(myBool *bool) {
	*myBool = !*myBool
}
