package main

import "fmt"

func main() {
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
}
