package main

import (
  "fmt"
)

func main(){
  arr := [5]int{10, 20, 30, 40, 50}

  for i , v := range arr {
    fmt.Println(i,v)
  }
  fmt.Printf("%T\n", arr)
}
