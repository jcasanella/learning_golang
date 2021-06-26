package main

import "fmt"

func main() {
	// Exercise 1.
	// Using a composite literal declare and initialize a slice of type string with 3 elements.
	// Iterate over the slice and print each element in the slice and its index.
	sports := []string{"roller", "wake", "skate"}
	for index, value := range sports {
		fmt.Printf("Index %d Value: %s\n", index, value)
	}

	// Exercsise 2.
	// 1. Declare a slice called nums with three float64 numbers.
	// 2. Append the value 10.1 to the slice
	// 3. In one statement append to the slice the values: 4.1,  5.5 and 6.6
	// 4. Print out the slice
	// 5. Declare a slice called n with two float64 values
	// 6. Append n to nums
	//7. Print out the nums slice
	nums := []float64{1.5, 2.5, 3.5}
	nums = append(nums, 10.1)
	nums = append(nums, 4.1, 5.5, 6.6)
	fmt.Println(nums)

	n := []float64{1.1, 2.2}
	nums = append(nums, n...)
	fmt.Println(nums)

	// Exercise 3
	// Create a Go program that reads some numbers from the command line and then calculates the sum and the product of all the numbers given at command line.
	// The user should give between 2 and 10 numbers.
}
