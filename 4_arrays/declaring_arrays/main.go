package main

import "fmt"

func main() {
	// define an array and printing values
	var numbers [4]int
	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	var a1 = [4]float64{}
	var a2 = [3]int{5, -3, 5}
	a3 := [4]string{"Dan", "Diana", "Paul", "John"}
	a4 := [4]string{"x", "y"}
	a5 := [...]int{1, 4, 5}
	fmt.Println("The length of a5 is: ", len(a5))

	// define an array multi line
	a6 := [...]int{1,
		2,
		3,
	}
	_, _, _, _, _, _ = a1, a2, a3, a4, a5, a6

	numbers[0] = 7
	fmt.Printf("%v\n", numbers)

	// getting elements
	x := numbers[0]
	fmt.Println("x is ", x)

	for i, v := range numbers {
		fmt.Println("Index:", i, " value:", v)
	}

	// iterating as in java
	for idx := 0; idx < len(numbers); idx++ {
		fmt.Println("index:", idx, "value:", numbers[idx])
	}

	// define multi array
	balances := [2][3]int{
		[3]int{5, 6, 7},
		{8, 9, 10},
	}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}

		fmt.Println()
	}

	m := [3]int{1, 2, 3}
	n := m

	fmt.Println("n is equal to m", n == m)
	m[0] = -1
	fmt.Println("n is equal to m", n == m)
}
