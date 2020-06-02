// fToC converts temperature in degrees Fahrenheit to Celsius
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter the temperature (°F): ")
	f, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	c := fToC(f)
	fmt.Printf("%0.2fC\n", c)
}

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func fToC(f float64) float64 {
	c := (f - 32) * 5 / 9
	return c
}
