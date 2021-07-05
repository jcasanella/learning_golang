package main

import "fmt"

func main() {
	s1 := "Here we go"
	fmt.Printf("%s\n", s1) // print the string
	fmt.Printf("%q\n", s1) // print the string quoted

	fmt.Println("He say \"Hello\"")
	fmt.Println(`He say "Hello"`)

	s2 := "Here we go2"
	fmt.Println(s2)

	fmt.Println("Price 100\nBrand: Ekin")

	fmt.Println(`
	Price: 100
	Brand: Ekin
	`)

	fmt.Println(`c:\users\jcasanella`)
	fmt.Println("c:\\users\\jcasanella")

	var s3 = "This is a test " + "concatenating strings " + "tra tra"
	fmt.Println(s3 + "!!!")

	fmt.Println("Element at index 0: ", s3[0])
}
