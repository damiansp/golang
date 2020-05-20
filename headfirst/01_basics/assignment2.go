package main

import (
	"fmt"
	"reflect"
)

func main() {
	var quantity = 4                   // int inferred
	var len, width = 1.2, 2.4          // float64 inferred
	var customerName = "Philip Damien" // string inferred

	fmt.Println(reflect.TypeOf(quantity))
	fmt.Println(reflect.TypeOf(len))
	fmt.Println(reflect.TypeOf(width))
	fmt.Println(reflect.TypeOf(customerName))
}
