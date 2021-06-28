package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

)

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
	numbers := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("You must enter between  2 and 10 numbers")
	for {
		// var num int
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			fmt.Println("No data entered")
			if len(numbers) > 1 {
				break
			}
		} else {
			num, err := strconv.Atoi(text)
			if err != nil {
				fmt.Println("Error captured")
				fmt.Println(err)
			} else {
				numbers = append(numbers, num)
				fmt.Println("Number entered:", num)
				if len(numbers) == 10 {
					break
				}
			}
		}
	}

	fmt.Println("Calculate sum and average")
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	fmt.Println(numbers)
	average := float32(sum) / float32(len(numbers))
	fmt.Printf("Sum: %d Average: %f\n", sum, average)

	// Exercise 4
	// Consider the following slice declaration: nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	// Using a slice expression and a for loop iterate over the slice ignoring the first and the last two elements.
	// Print those elements and their sum.
	numsA := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	numsB := numsA[1 : len(numsA)-2]
	sumB := 0
	for _, value := range numsB {
		fmt.Print(" ", value)
		sumB += value
	}
	fmt.Println()
	fmt.Println(sumB)

	// Exercise 5
	// Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}
	// Using copy() function create a copy of the slice. Prove that the slices are not connected by modifying one slice and notice
	// that the other slice is not modified.
	friends := []string{"Marry", "John", "Paul", "Diana"}
	friends2 := make([]string, len(friends))
	copy(friends2, friends)
	fmt.Println("Friends:", friends)
	fmt.Println("Friends2:", friends2)

	friends[0] = "Peter"
	fmt.Println("Friends:", friends)
	fmt.Println("Friends2:", friends2)

	friends2[1] = "Joe"
	fmt.Println("Friends:", friends)
	fmt.Println("Friends2:", friends2)

	// Exercise 6
	// Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}
	// Using append() function create a copy of the slice. Prove that the slices are not connected by modifying one
	// slice and notice that the other slice is not modified.
	friendsA := []string{"Marry", "John", "Paul", "Diana"}
	var friendsB []string
	friendsB = append(friendsB, friendsA...)
	fmt.Println(friendsB)
	friendsB[1] = "Adrian"
	fmt.Println(friendsA)
	fmt.Println(friendsB)

	// Exercise 7
	// Consider the following slice declaration:
	// years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	// Using a slice expression and append() function create a new slice called newYears
	// that contains the first 3 and the last 3 elements of the slice.  newYears should be []int{2000, 2001, 2002, 2008, 2009, 2010}
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	var newYears []int
	newYears = append(years[:3], years[len(years)-3:]...)
	fmt.Println(newYears)
}
