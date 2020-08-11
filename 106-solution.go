package main

import (
	"fmt"
)

func main() {
	s := struct {
		name      string
		friends   map[string]int
		favDrinks []string
	}{
		name: "Yash",
		friends: map[string]int{
			"snake": 97879,
			"dog":   3123123,
			"cat":   312313,
		},
		favDrinks: []string{"Pepsi", "Coke"},
	}

	fmt.Println(s)

}
