package main

import (
	"fmt"
)

type scotch int

var x scotch
var y int

func main() {
	// part a
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	// part b
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
