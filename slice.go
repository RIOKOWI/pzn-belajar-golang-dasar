package main

import	"fmt"

func main(){
	names := [...]string{
		"udin",
		"dudun",
		"abel",
		"yogo",
		"adly",
		"dono",
		"zee",
		"antonio",
		"faris",
		"coki",
	}

	slice := names[4:6]
	// pointer = index 4
	// length = 2
	// capacity = 6

	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(slice[1])


	slice1 := names[:3]

	fmt.Println(slice1)


	slice2 := names[3:]

	fmt.Println(slice2)

	slice3 := names[:]

	fmt.Println(slice3)
}