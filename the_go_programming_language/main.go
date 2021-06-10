package main

import (
	"fmt"
	"learning_golang/the_go_programming_language/tempconv"

)

func main() {
	// Fah to Kel
	fmt.Println("Fahrenheit to Kelvin")
	fmt.Printf("%g\n", tempconv.FToK(-30))
	fmt.Printf("%g\n", tempconv.FToK(-20))
	fmt.Printf("%g\n", tempconv.FToK(-10))
	fmt.Printf("%g\n", tempconv.FToK(0))
	fmt.Printf("%g\n", tempconv.FToK(10))
	fmt.Printf("%g\n", tempconv.FToK(20))
	fmt.Printf("%g\n", tempconv.FToK(30))

	// Fah to Cel
	fmt.Println("Fahrenheit to Celsius")
	fmt.Printf("%g\n", tempconv.FToC(-30))
	fmt.Printf("%g\n", tempconv.FToC(-20))
	fmt.Printf("%g\n", tempconv.FToC(-10))
	fmt.Printf("%g\n", tempconv.FToC(0))
	fmt.Printf("%g\n", tempconv.FToC(10))
	fmt.Printf("%g\n", tempconv.FToC(20))
	fmt.Printf("%g\n", tempconv.FToC(30))

	// Cel to Kel
	fmt.Println("Celcius to Kelvin")
	fmt.Printf("%g\n", tempconv.CToK(-30))
	fmt.Printf("%g\n", tempconv.CToK(-20))
	fmt.Printf("%g\n", tempconv.CToK(-10))
	fmt.Printf("%g\n", tempconv.CToK(0))
	fmt.Printf("%g\n", tempconv.CToK(10))
	fmt.Printf("%g\n", tempconv.CToK(20))
	fmt.Printf("%g\n", tempconv.CToK(30))

	// Celc to Fah
	fmt.Println("Celcius to Fahrenheit")
	fmt.Printf("%g\n", tempconv.CToF(-30))
	fmt.Printf("%g\n", tempconv.CToF(-20))
	fmt.Printf("%g\n", tempconv.CToF(-10))
	fmt.Printf("%g\n", tempconv.CToF(0))
	fmt.Printf("%g\n", tempconv.CToF(10))
	fmt.Printf("%g\n", tempconv.CToF(20))
	fmt.Printf("%g\n", tempconv.CToF(30))

	// Kel to Fah
	fmt.Println("Kelvin to Fahrenheit")
	fmt.Printf("%g\n", tempconv.KToF(-30))
	fmt.Printf("%g\n", tempconv.KToF(-20))
	fmt.Printf("%g\n", tempconv.KToF(-10))
	fmt.Printf("%g\n", tempconv.KToF(0))
	fmt.Printf("%g\n", tempconv.KToF(10))
	fmt.Printf("%g\n", tempconv.KToF(20))
	fmt.Printf("%g\n", tempconv.KToF(30))

	// Kel to Cel
	fmt.Println("Kelvin to Celsius")
	fmt.Printf("%g\n", tempconv.KToC(-30))
	fmt.Printf("%g\n", tempconv.KToC(-20))
	fmt.Printf("%g\n", tempconv.KToC(-10))
	fmt.Printf("%g\n", tempconv.KToC(0))
	fmt.Printf("%g\n", tempconv.KToC(10))
	fmt.Printf("%g\n", tempconv.KToC(20))
	fmt.Printf("%g\n", tempconv.KToC(30))
}
