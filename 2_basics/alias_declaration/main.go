package main

import "fmt"

func main() {
	var a uint8 = 10
	var b byte

	b = a

	fmt.Println(b)

	type second = uint
	var hour second = 3600
	fmt.Printf("%d\n", hour/60)

	type duration second
	var d duration
	fmt.Printf("%T\n", d)

	type second2 uint
	type duration2 second2
	type minute2 = uint

	var t1 duration2 = 10
	// var x uint = t1 compilation error
	_ = t1

	var t2 minute2 = 10
	var x uint = t2

	_ = x
}
