# Getting Started

### Activity 1. 

In a directory called **exercise2** under **$GOPATH/src/master_go_programming/1_getting_started** create a file named **[main.go](exercise2/main.go)** with the typical structure of any Go Application
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

### Activity 2. 

Compile the program from the previous exercise using the **go** command. The executable should appear in the current working directory. 
```
cd exercise2

# This creates an app with the name main
go build main.go

# This creates an app with the name of the folder
go build
```

### Activity 3. 
Compile the program from **Exercise 1** manually using the *go* command. Produce an executable called:

* my_first_go_app.exe if your are on Windows
* my_first_go_app if you are on Linux/Mac in the current working directory

```
go build -o my_first_go_app.exe main.go
```

### Activity 4. 

Do the following depending your OS:

* If you are on Windows compile the program from Exercise #1 for Linux using the go command. 
* If you are on Mac compile the program from Exercise #1 for Windows using the go command. 
* If you are on Linux compile the program from Exercise #1 for Mac using the go command.

https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures

https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04

https://binx.io/blog/2018/11/25/go-cross-compilation/

https://golangcookbook.com/chapters/running/cross-compiling/

```
# Print all valids platforms
go tool dist list

# Print current platform and architecture
go env GOOS GOARCH

GOOS=linux GOARCH=amd64 go build -o  my_go_program
```

### Activity 5. 

Using the **go** command and the appropiate option, format the source code file in the idiomatic Go style assuming this file: [main.go](main.go)

```
package main
import "fmt"
func main(){	SayHello()
}
func SayHello() {
	fmt.Println("Hello package v1.0.0!")

	var x int=10
	_=x}
```

```
gofmt -l -w main.go

# or simply run
go fmt
```


