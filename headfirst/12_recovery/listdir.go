package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	paths, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, path := range paths {
		if path.IsDir() {
			fmt.Println("Directory:", path.Name())
		} else {
			fmt.Println("File:", path.Name())
		}
	}
}
