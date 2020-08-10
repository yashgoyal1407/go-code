package main
import (
  "fmt"
)
type whisky int
var x whisky
func main(){
  fmt.Println(x)
  fmt.Printf("%T\n", x)
  x = 42
  fmt.Println(x)
}
