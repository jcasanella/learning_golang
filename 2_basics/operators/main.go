package main

import "fmt"

func main() {
	a, b := 4, 2
	r := (a + b) / (a - b) * 2
	fmt.Printf("%T %d\n", r, r)
	fmt.Printf("%t\n", a == b)
	fmt.Printf("%t\n", a != b)
	fmt.Println(a > b, a <= b)
	fmt.Println(a > 1 && b < 5)
	fmt.Println(a == 4 || b == 1)

	r = 9 % a
	fmt.Printf("%T %d\n", r, r)

	c, d := 2, 3
	c += d
	fmt.Printf("%T %d\n", c, c)

	d -= 2
	fmt.Printf("%T %d\n", d, d)

	d *= 2
	fmt.Printf("%T %d\n", d, d)

	d /= 2
	fmt.Printf("%T %d\n", d, d)

	x := 1
	x++
	fmt.Printf("%d\n", x)

	x--
	fmt.Printf("%d\n", x)
}
