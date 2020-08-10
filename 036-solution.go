package main

import (
  "fmt"
)

var x int
var y string
var z bool

func main() {
  fmt.Println(x)
  fmt.Println(y)
  fmt.Println(z)
  fmt.Println(x,y,z)

  x = 42
  y = "James Bond"
  z = true

  s := fmt.Sprintf( "%v\t%v\t%v\n", x, y, z)
  fmt.Println(s)
}
