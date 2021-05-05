package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// Exercise 1 - Using a for loop and an if statement print out all the numbers between 1 and 50
	// that divisible by 7
	for idx := 1; idx <= 50; idx++ {
		if idx%7 == 0 {
			fmt.Printf("%d is divisible by 7\n", idx)
		}
	}

	// Exercise 2 - Change the code from the previous exercise and use the continue statement
	// to print out all the numbers divisible by 7 between 1 and 50
	for idx := 1; idx <= 50; idx++ {
		if idx%7 != 0 {
			continue
		}

		fmt.Printf("%d is divisible by 7\n", idx)
	}

	// Exercise 3 - Change the code from the preveious exercise and use the break statement to print out
	// only the first 3 numbers divisible by 7 between 1 and 50
	count := 0
	for idx := 1; idx <= 50; idx++ {
		if idx%7 != 0 {
			continue
		}

		fmt.Printf("%d is divisible by 7\n", idx)
		count++
		if count == 3 {
			break
		}
	}

	// Exercise 4 - Using a for loop, an if statement and the logical and operator print out all the numbers
	// between 1 and 500 that divisible both by 7 and 5
	for idx := 1; idx <= 500; idx++ {
		if idx%7 == 0 && idx%5 == 0 {
			fmt.Printf("%d is divisible by 7 and 5\n", idx)
		}
	}

	// Exercise 5 - Using a for loop print out all the years from your birthday to the current year. Use a variant
	// of for loop where the post statement is moved inside the for block code
	t := time.Now()
	year := t.Year()
	dob, _ := strconv.Atoi(os.Args[0])
	for dob <= year {
		fmt.Printf("Year: %d\n", dob)
		dob++
	}

	// Exercise 6 - Change the following code to use a swith instead of a if else
	// age := -9
	// if age < 0 || age > 100 {
	// 	fmt.Println("Invalid Age")
	// } else if age < 18 {
	// 	fmt.Println("You are minor!")
	// } else if age == 18 {
	// 	fmt.Println("Congratulations! You've just become major!")
	//  } else {
	//  	fmt.Println("You are major!")
	// }
	age := -9
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid age")
	case age < 18:
		fmt.Println("You're a minor!!!")
	case age == 18:
		fmt.Println("Congrats!!! U've just become major!!!")
	default:
		fmt.Println("You're major")

	}
}
