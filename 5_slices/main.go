package main

import "fmt"

func main() {
	// declaring a string slices
	var cities []string
	fmt.Println("Cities is equal to nil: ", cities == nil)
	fmt.Printf("Cities: %#v\n", cities)

	// declaring a slice using literal
	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers)

	// Creating a slice of length 2
	nums := make([]int, 2)
	fmt.Println(nums)

	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)

	x := numbers[0]
	fmt.Println("x is ", x)

	// Modifying an element of the slice
	numbers[1] = 200
	fmt.Printf("%#v\n", numbers)

	// iterate over a slice
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("index: %v, value: %v\n", i, numbers[i])
	}

	// slices with the same element type can be assigned each other
	var n []int
	n = numbers
	_ = n

	// Comparing slices, a go slice can only be compared to nil

	// uninitialized slice, equal to nil
	var nn []int
	fmt.Println(nn == nil)

	// empty slice but initialized, not equal to nil
	mm := []int{}
	fmt.Println(mm == nil)

	// we can not compare slices using the equal (=) operator
	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	var eq bool = true
	for i, valueA := range a {
		if valueA != b[i] {
			eq = false
			break
		}
	}

	// also we need to check the length og slices
	if len(a) != len(b) {
		eq = false
	}

	if eq {
		fmt.Println("Slices A anb B are equal")
	} else {
		fmt.Println("Slices are not ")
	}
}
