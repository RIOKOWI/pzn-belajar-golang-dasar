package main

import "fmt"

func main(){
	const firstName = "Rio"
	const lastName = "Achyar"
	
	// error tidak bisa diubah
	// firstName = "lutpi"
	// lastName = "fandi"

	const ( // multiple constant
		first = "Rio"
		last = "Kowi"
	)

	fmt.Println(first)
	fmt.Println(last)


}