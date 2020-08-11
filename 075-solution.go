package main

import "fmt"

func main() {
	favSport := "Cricket"
	switch favSport {
	case "Rugby":
		fmt.Println("Rugby")
	case "Football":
		fmt.Println("Football")
	case "Tennis":
		fmt.Println("Tennis")
	case "Cricket":
		{
			fmt.Println("Cricket")
		}
	default:
		fmt.Println("Lazy Kid")
	}
}
