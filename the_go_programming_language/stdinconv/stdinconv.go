package stdinconv

import "fmt"

type Celsius float64
type Farenheit float64
type Kelvin float64

const AbsoluteZeroC Celsius = -273.15

func CToF() Farenheit {
	var c Celsius

	fmt.Println("Enter the Celsius to transform to Farenheit")
	fmt.Scanf("%f", &c)
	return Farenheit(c*9/5 + 32)
}

func CToK() Kelvin {
	var c Celsius

	fmt.Println("Enter the Celsius to transform to Kelvin")
	fmt.Scanf("%f", &c)
	return Kelvin(c - AbsoluteZeroC)
}

func (s Farenheit) String() string {
	return fmt.Sprintf("%g F", s)
}

func (s Kelvin) String() string {
	return fmt.Sprintf("%g K", s)
}
