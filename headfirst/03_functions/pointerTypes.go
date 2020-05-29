package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	var myFloat float64
	var myBool bool
	fmt.Println(reflect.TypeOf(&myInt))   // *int
	fmt.Println(reflect.TypeOf(&myFloat)) // *float64
	fmt.Println(reflect.TypeOf(&myBool))  // *bool
}
