package main

import "fmt"

func main() {
	// Exercise 1

	// 1. Using the var keyword, declare an array called cities with 2 elements of type string. Don't initialize the array
	// Print out the cities array and notice the value of its elements
	var cities [2]string
	fmt.Printf("%v\n", cities)

	// 2. Declare an array called grades of type [3]float64 and initialize only the first 2 elements using an array literal.
	var grades = [3]float64{45.56, 3.14}
	fmt.Printf("%v\n", grades)

	// 3. Declare an array called taskDone using the ellipsis operator. The elements are of type bool.
	taskDone := [...]bool{true, false, true, true}
	fmt.Printf("%v\n", taskDone)

	// 4. Iterate over the array called cities using the classical for loop syntax and len function and print out element by element.
	for idx := 0; idx < len(taskDone); idx++ {
		fmt.Printf("%v ", taskDone[idx])
	}

	// 5. Iterate over grades using the range keyword and print out element by element.
	for _, value := range taskDone {
		fmt.Printf("%t ", value)
	}

	// Excercise 2

	fmt.Println()

	// Consider the following array declaration: nums := [...]int{30, -1, -6, 90, -6}
	// Write a Go program that counts the number of positive even numbers in the array.
	nums := [...]int{30, -1, -6, 90, -6}
	total := 0
	for _, value := range nums {
		if value%2 == 0 && value > 0 {
			total++
		}
	}

	fmt.Printf("Total: %d\n", total)
}
