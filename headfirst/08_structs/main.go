package main

import (
	"fmt"

	"./magazine"
)

func main() {
	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)

	subscriber := magazine.Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	address := magazine.Address{
		Street: "123 Elm St.", City: "Elgin", State: "IL", PostalCode: "61407"}
	subscriber.Address = address
	fmt.Println("Name:", subscriber.Name)
	fmt.Println("Rate:", subscriber.Rate)
	fmt.Println("Active:", subscriber.Active)
	fmt.Println(subscriber.Address)
}
