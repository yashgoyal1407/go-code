package main

import (
  "fmt"
)

func main(){
  slc := [][]string{}
  fmt.Println(slc)
  fmt.Println(len(slc))
  a := []string{"James", "Bond", "Shaken, not stirred"}
  b := []string{"Miss", "Moneypenny", "Helloooooo, James."}
  slc = append(slc, a,b)
  fmt.Println(slc)
  fmt.Println(len(slc))

  for _, v := range slc {
    fmt.Println(v)
    for _,k := range v{
      fmt.Println(k)
    }
  }
}
