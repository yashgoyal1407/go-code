package main

import (
	"fmt"
)

type person struct {
	first_name string
	last_name  string
	ice_cream  []string
}

func main() {
	p1 := person{
		first_name: "yash",
		last_name:  "goyal",
		ice_cream:  []string{"vanila", "strawberry"},
	}

	p2 := person{
		first_name: "todd",
		last_name:  "mcload",
		ice_cream:  []string{"vanila", "cranberry"},
	}

	fmt.Println(p1)
	fmt.Println(p1.first_name)
	fmt.Println(p1.last_name)
	print_icecream(p1.ice_cream)
	fmt.Println()

	fmt.Println(p2)
	fmt.Println(p2.first_name)
	fmt.Println(p2.last_name)
	print_icecream(p2.ice_cream)
	fmt.Println()

}

func print_icecream(icecream []string) {
	for i, v := range icecream {
		fmt.Println(i, v)
	}
}
