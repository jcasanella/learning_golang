package main

import "fmt"

func main() {
	language := "golang"

	switch language {
	case "java":
		fmt.Println("Java is the language")

	case "go", "golang":
		fmt.Println("Golang is the language")

	default:
		fmt.Println("Any other programming language")
	}

	n := 5
	switch true {
	case n%2 == 0:
		fmt.Println("Even number")

	case n%2 != 0:
		fmt.Println("Odd number")

	default:
		fmt.Println("Never here")
	}

	switch n := 10; true {
	case n > 0:
		fmt.Println("Positive number")

	case n < 0:
		fmt.Println("Negative number")

	default:
		fmt.Println("Zero number")
	}
}
