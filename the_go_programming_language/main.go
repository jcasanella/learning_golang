package main

import (
	"fmt"
	"learning_golang/the_go_programming_language/tempconv"

)

func main() {
	values := [...]int{-30, -20, -10, 0, 10, 20, 30}
	for _, value := range values {
		fmt.Printf("%d F is %s\n", value, tempconv.FToK(tempconv.Farenheit(value)))
		fmt.Printf("%d F is %s\n", value, tempconv.FToC(tempconv.Farenheit(value)))
		fmt.Printf("%d C is %s\n", value, tempconv.CToK(tempconv.Celsius(value)))
		fmt.Printf("%d C is %s\n", value, tempconv.CToF(tempconv.Celsius(value)))
		fmt.Printf("%d K is %s\n", value, tempconv.KToF(tempconv.Kelvin(value)))
		fmt.Printf("%d K is %s\n", value, tempconv.KToC(tempconv.Kelvin(value)))
	}
}
