package main

import "fmt"

func main(){
  i := 1997
  for i <= 2020 {
    if i > 2020{
      break
    }

    fmt.Println(i)
    i++
  }
}
