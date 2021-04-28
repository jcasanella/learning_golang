package main

import (
	"fmt"
	"os"

)

func main() {
	fmt.Println("os.Args", os.Args)
	fmt.Println("Path", os.Args[0])
	fmt.Println("First arg", os.Args[1])
	fmt.Println("Second arg", os.Args[2])
	fmt.Println("Num args", len(os.Args))
}
