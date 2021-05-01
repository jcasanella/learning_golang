package main

import "fmt"

func main() {
	// declaring a variable with the same name as a label
	outer := 19

	people := [5]string{"paul", "alex", "sofi", "ann", "john"}
	friends := [2]string{"alex", "dave"}

outer:
	for index, name := range people {
		for _, name2 := range friends {
			if name == name2 {
				fmt.Printf("They're friends %q at index %d\n", name, index)
				break outer
			}
		}
	}

	_ = outer
}
