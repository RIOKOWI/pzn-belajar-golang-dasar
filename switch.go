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


	jeneng := "sumanto"

	switch length := len(jeneng); length <= 7 {
	case true :
		fmt.Println("Hi", jeneng)
	case false :
		fmt.Println("lapo kon?")
	}


	length := len(jeneng)

	switch {
	case length > 10:
		fmt.Println("terlalu panjang")
	case length > 5:
		fmt.Println("oke")
	default:
		fmt.Println("terlalu pendek")
	}
}