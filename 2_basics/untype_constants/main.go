package main

import "fmt"

func main() {
	const a float64 = 5.64      // type constant
	const b = 7.14              // untype constant - behind scenes are different
	const c float64 = a * b     // remember constants are evaluated during compile time
	const str = "hello " + "go" // untype constant
	const d = 5 > 10

	fmt.Printf("%t\n", d)

	const y = 5 * 4.5
	// const err int = 5 * 4.5 - fail to compile because this type constant is type int and can not multiply int * float
	fmt.Printf("%T %f\n", y, y)

	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Printf("%d %d %d\n", c1, c2, c3)
	fmt.Println(c1, c2, c3)

	// other way to assign iota constant
	const (
		c11 = iota
		c12
		c13
	)
	fmt.Println(c11, c12, c13)

	// iota with steps of 2
	const (
		c21 = iota * 2
		c22
		c23
		c24
	)
	fmt.Println(c21, c22, c23, c24)

	const (
		c31 = (iota * 2) + 1
		c32
		c33
		c34
	)
	fmt.Println(c31, c32, c33, c34)

	// skip iota element
	const (
		d51 = -(iota + 2)
		_   // skip -3
		d53
		d54
		d55
	)
	fmt.Println(d51, d53, d54, d55)

}
