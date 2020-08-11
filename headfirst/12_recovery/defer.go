package main

import "fmt"

func main() {
	Socialize()
}

func Socialize() {
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	fmt.Println("Nice weather, eh?")
}
