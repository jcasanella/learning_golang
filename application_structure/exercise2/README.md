# Exercises Section 3 - Getting Started

## 1. In a directory called *exercise2* under *$GOPATH/src/master_go_programming/2_getting_started* create a file named *main.go* with the typical structure of any Go Application
```
package main

import "fmt"

var x int = 100

const y = 3600

func main() {
	var a int = 7
	var b float64 = 3.5
	const c int = 10

	fmt.Println("Hello test")
	fmt.Println(a, b, c)
}
```

## 2. Compile the program from the previous exercise using the *go* command. The executable should appear in the current working directory. Notice its name and run it!

## 3. Compile the program from *Exercise 1* manually using the *go* command. Produce an executable called:

* my_first_go_app.exe if your are on Windows
* my_first_go_app if you are on Linux/Mac in the current working directory

## 4. Do the following depending your OS:

* If you are on Windows compile the program from Exercise #1 for Linux using the go command. 
* If you are on Mac compile the program from Exercise #1 for Windows using the go command. 
* If you are on Linux compile the program from Exercise #1 for Mac using the go command.

## 5. Using the go command and the appropiate option, format the source code file in the idiomatic Go style.


