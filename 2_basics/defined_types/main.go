package main

import "fmt"

type km float64
type mile float64

func main() {
	type speed uint

	var s1 speed = 20
	var s2 speed = 30
	fmt.Println(s2 - s1)
	fmt.Printf("%T %v - %T %v", s1, s1, s2, s2)

	// x = s1 Compilation error
	var x uint = uint(s1)
	fmt.Println(x)

	var s3 speed = speed(x)
	_ = s3

	var parisToLondon km = 465
	var distanceInMiles mile

	distanceInMiles = mile(parisToLondon) / 0.621
	fmt.Println(distanceInMiles)
}
