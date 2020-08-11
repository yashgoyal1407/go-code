package main

import "fmt"

func main(){
  for num := 10; num <= 100; num++ {
    fmt.Printf("On modulus if %v by %v, we will get %v\n", num,4,num%4)
  }
}
