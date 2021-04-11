package main

import (
	"fmt"
	"strconv"

)

func main() {
	s := string(99)
	fmt.Println(s) // print ascii code

	// s = string(45.2) this is an error
	var myStr = fmt.Sprintf("%f", 45.2)
	fmt.Printf("%T %s\n", myStr, myStr)

	// Convert strings to numbers
	s1 := "3.123" // type string
	fmt.Printf("%T\n", s1)
	var f1, _ = strconv.ParseFloat(s1, 64) // string to float
	fmt.Printf("%T %f\n", f1, f1)

	i, _ := strconv.Atoi("-50")
	fmt.Printf("%T %d\n", i, i)

	s2 := strconv.Itoa(56)
	fmt.Printf("%T %s\n", s2, s2)
}
