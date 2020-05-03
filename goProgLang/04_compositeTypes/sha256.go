package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	diffBits := 0
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	for i := range c1 {
		if c1[i] != c2[i] {
			diffBits++
		}
	}
	fmt.Printf("N bits differing: %d\n", diffBits)
}
