package main

import "fmt"

var version = "3.1"

func main() {
	var (
		a int
		b float64
		c bool
		d string
	)

	x, y, z := 20, 15.5, "Gopher!"

	_, _, _, _, _, _, _ = a, b, c, d, x, y, z

	var a1 float64 = 7.1
	x1, y1 := true, 3.7
	a1, x1 = 5.5, false
	_, _, _ = a1, x1, y1

	name := "Golang"
	_ = name

	const (
		daysWeek   = 7
		lightSpeed = 299792458
		pi         = 3.14159
	)

	const (
		secPerDay = 60 * 60 * 24
		daysYear  = 365
	)
	fmt.Printf("%d\n", daysYear*secPerDay)

	const (
		Jun = iota + 6
		Jul
		Aug
	)

	fmt.Println(Jun, Jul, Aug)
}
