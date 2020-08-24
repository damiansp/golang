package main

import "fmt"

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func slow() {
	a()
	b()
	fmt.Println("\nend slow()")
}

func fast() {
	go a()
	go b()
	fmt.Println("\nend fast()")
}

func main() {
	slow()
	fast()
	fmt.Println("\nend main()")
}
