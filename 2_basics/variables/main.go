package main

import "fmt"

func main() {
	var name string = "Peter"
	fmt.Printf("%s\n", name)

	// := its only for new variables
	name = name + " Parker"
	fmt.Printf("%s\n", name)

	// infer data type
	var age = 40
	fmt.Printf("Age %d\n", age)

	// Create new variables
	var (
		address string
		city    string
		zipcode string
	)

	address = "24 Boulevard"
	city = "London"
	zipcode = "Ld1"
	fmt.Println("Address: " + address + " City: " + city + " ZipCode: " + zipcode)

	// Create new variables
	price, units := 50, 10
	fmt.Printf("Total Price: %d, unit price: %d, total units: %d\n", price*units, price, units)

	// to use := at least one of the variables must be a new one
	age, age2 := 99, 24
	fmt.Printf("old age: %d new age: %d\n", age, age2)

	// when a var is not used, needs to be either dropped or _ assigned to blank operator
	var notUsed = "dummy"
	_ = notUsed
}
