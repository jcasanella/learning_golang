package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

)

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

	// runes
	// characters or rune literals are expressed in Go by enclosing them in single quotes
	// declaring a variable of type rune (alias to int32)
	var1 := 'a'
	fmt.Printf("Type: %T, Value %d\n", var1, var1)

	// declaring a string value that contains non-ascii characters
	str := "ţară" // ţară means country in Romanian
	// 't', 'a' ,'r' and 'a' are runes and each rune occupies beetween 1 and 4 bytes.

	//The len() built-in function returns the no. of bytes not runes or chars.
	fmt.Println(len(str)) // -> 6,  4 runes in the string but the length is 6

	// returning the number of runes in the string
	m := utf8.RuneCountInString(str)
	fmt.Println(m) // => 4

	// by using indexes we get the byte at that positioin, not rune.
	fmt.Println("Byte (not rune) at position 1:", str[1]) // -> 163

	// Decoding a string byte by byte
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	fmt.Println("\n" + strings.Repeat("#", 10))

	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c", r)
		i += size
	}

	fmt.Println("\n" + strings.Repeat("#", 10))

	// decoding a string rune by rune automatically:
	for i, r := range str { //the first value returned by range is the index of the byte in string where rune starts
		fmt.Printf("%d -> %c", i, r) // => ţară
	}

	s1A := "abcdefghijkl"
	fmt.Println(s1A[2:5]) // -> cde, bytes from 2 (included) to 5 (e

	s2A := "中文维基是世界上"
	fmt.Println(s2A[0:2]) // -> � - the unicode representation of bytes fro

	// returning a slice of runes
	// 1st step: converting string to rune slice
	rs := []rune(s2A)
	fmt.Printf("%T\n", rs) // => []int32

	// 2st step: slicing the rune slice
	fmt.Println(string(rs[0:3])) // => 中文维
}
