package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	t := truck{
		vehicle : vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: false,
	}

	fmt.Println("Printing Truck Details", t)
	fmt.Println(t.doors)
	fmt.Println(t.color)
	fmt.Println(t.fourWheel)
	fmt.Println()

	s := sedan{
		vehicle : vehicle{
			doors: 4,
			color: "silver",
		},
		luxury: true,
	}

	fmt.Println("Printing Sedan Details", s)
	fmt.Println(s.doors)
	fmt.Println(s.color)
	fmt.Println(s.luxury)
	fmt.Println()

}
