package main

import (
	"fmt"
	"os"
	"strconv"

)

func main() {
	i, err := strconv.Atoi("55a")
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(i)
	}

	if i, err := strconv.Atoi("20"); err == nil {
		fmt.Println("No error", i)
	} else {
		fmt.Println(err)
	}

	if args := os.Args; len(os.Args) != 2 {
		fmt.Println("One argument is required")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("The argument must be an integer!!! Error", err)
	} else {
		fmt.Printf("%d km in miles is %v\n", km, float64(km)*1.609)
	}
}
