package main

import "fmt"

func main() {
	switch  {
	case false:
		fmt.Println("false")
  case true:
    fmt.Println(true)
	default:
		fmt.Println("true")
	}
}
