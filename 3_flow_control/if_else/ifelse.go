package main

import "fmt"

func main() {
	price, inStock := 100, true
	if price > 80 {
		fmt.Println("Too expensive")
	}

	if price <= 100 && inStock {
		fmt.Println("Buy it!!!")
	}

	if price < 100 {
		fmt.Println("it's cheap")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("It's expensive")
	}

	age := 200
	if age >= 0 && age < 18 {
		fmt.Println("You cannot vote")
	} else if age == 18 {
		fmt.Println("This is your first vote")
	} else if age > 18 && age < 100 {
		fmt.Println("Please vote it's important")
	} else {
		fmt.Println("Invalid age")
	}
}
