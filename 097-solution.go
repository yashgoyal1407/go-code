package main

import (
  "fmt"
)

func main(){
  m := map[string][]string{
    `bond_james` : []string{`Shaken, not stirred`, `Martinis`, `Women`},
    `moneypenny_miss` : []string{ `James Bond`, `Literature`, `Computer Science`},
    `no_dr` : []string{`Being evil`, `Ice cream`, `Sunsets`},
  }

  fmt.Println(m)
  fmt.Println()

  m["yash"] = []string{"Great", "Cricket", "code"}

  for k, v := range m {
    fmt.Printf("Key = %v, value = %v \n", k, v)
    for i, val := range v{
      fmt.Printf("index = %v, value = %v \n", i, val)
    }
  }

  delete(m, "yash")

  fmt.Println(m)
  fmt.Println()

}
