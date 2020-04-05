// Prints command line args
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Print(i)
		fmt.Println(": " + arg)
	}
}
