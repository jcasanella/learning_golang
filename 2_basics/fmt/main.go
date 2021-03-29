package main

import "fmt"

func main() {
	fmt.Println("Hello Golang!!!")

	var name, age = "Peter", 35
	fmt.Println(name, "is", age, "years old")

	// Printf
	// %d -> decimal
	// %f -> float
	// %s -> string
	// %q -> quoted string
	// %v -> any value
	// %T -> value type
	// %t -> boolean
	// %p -> pointer
	// %c -> char represented by the corresponding unicode
	a, b, c := 10, 15.5, "Goephers"
	grades := []int{10, 20, 30}

	fmt.Printf("a is %d, b is %f, c is %s\n", a, b, c)
	fmt.Printf("%q\n", c)
	fmt.Printf("\"%s\"\n", c)

	fmt.Printf("%v\n", grades)
	fmt.Printf("%#v\n", grades)

	fmt.Printf("b is of type %T and grades is of type %T", b, grades)

	fmt.Printf("The address of a: %d\n", &a)
	fmt.Printf("%c and %c\n", 100, 51011)

	const pi float64 = 3.1415926
	fmt.Printf("Pi is %.4f\n", pi)

	fmt.Printf("255 in base2 is %b\n", 255)
	fmt.Printf("101 in base16 is %x\n", 101)

	s := fmt.Sprintf("a is %d, b is %f, c is %s\n", a, b, c)
	fmt.Println(s)
}
