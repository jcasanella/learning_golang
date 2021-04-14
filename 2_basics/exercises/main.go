package main

import (
	"fmt"
	"strconv"

)

var version = "3.1"

func main() {
	var (
		a int
		b float64
		c bool
		d string
	)

	x, y, z := 20, 15.5, "Gopher!"
	score := []int{10, 20, 30}
	fmt.Printf("%d %f %s %#v\n", x, y, z, score)
	fmt.Printf("%q\n", z)
	fmt.Printf("%v %v %v %v\n", x, y, z, score)
	fmt.Printf("%T %T\n", y, score)

	_, _, _, _, _, _, _ = a, b, c, d, x, y, z

	const x2 float64 = 1.422349587101
	fmt.Printf("%.4f\n", x2)

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

	shape := "circle"
	radius := 3.2
	const pi2 float64 = 3.14159
	circumference := 2 * pi2 * radius

	fmt.Printf("Shape: %q\n", shape)
	fmt.Printf("Circle's circumference with radius %f is %f\n", radius, circumference)

	var i int = 3
	var f float64 = 3.2

	fmt.Printf("%f %d\n", float64(i), int(f))

	var i2 = 3
	var f1 = 3.2
	var s1, s2 = "3.14", "5"

	s1 = strconv.Itoa(i2)
	s3, _ := strconv.Atoi(s2)
	s4 := strconv.FormatFloat(f1, 'E', -1, 64)
	s5, _ := strconv.ParseFloat(s1, 64)
	fmt.Printf("%T %s\n", s1, s1)
	fmt.Printf("%T %d\n", s3, s3)
	fmt.Printf("%T %s\n", s4, s4)
	fmt.Printf("%T %f\n", s5, s5)

	x3, y3 := 4, 5.1
	z3 := float64(x3) * y3
	fmt.Println(z3)

	a3 := 5
	b3 := 6.2 * float64(a3)
	fmt.Println(b3)

	type duration int
	var hour duration

	fmt.Printf("%v %T\n", hour, hour)

	hour = 3600
	fmt.Printf("%v %T\n", hour, hour)

	minute := 60
	fmt.Println(int(hour) != minute)
	fmt.Println(hour != duration(minute))

	type mile float64
	type kilometer float64

	const m2km = 1.609
	var (
		mileBerlinToParis mile = 655.3
		kmBerlinToParis   kilometer
	)

	kmBerlinToParis = kilometer(mileBerlinToParis * m2km)
	fmt.Printf("Distance in Km from Berlin to Paris is %f\n", kmBerlinToParis)
}
