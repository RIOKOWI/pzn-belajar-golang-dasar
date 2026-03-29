package main

import "fmt"

func main() {
	name := "rio"

	switch name {
	case "Achyar" : // if
		fmt.Println("Hello Achyar")
	case "rio" : // if
		fmt.Println("Hello ", name)
	default : // else
		fmt.Println("hayuuuuuukk")
	}
}