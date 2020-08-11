package main

import (
  "fmt"
)

func main(){
  slc := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

  for i , v := range slc {
    fmt.Println(i,v)
  }
  fmt.Printf("%T\n", slc)
}
