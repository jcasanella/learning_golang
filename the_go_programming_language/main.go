package main

import (
	"fmt"
	"learning_golang/the_go_programming_language/tempconv"

)

func main() {
	fmt.Printf("%v\n", tempconv.BoilingC)
	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC)

	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", boilingF-tempconv.CToF(tempconv.FreezingC))

	fmt.Printf("%s", tempconv.FreezingC)
}
