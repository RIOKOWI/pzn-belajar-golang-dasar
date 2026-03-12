package main

import "fmt"

func main(){
	var name string

	name = "Rio Achyar"
	fmt.Println(name)

	name = "Achyar Rio"
	fmt.Println(name)

	var io = 1 // simplify tidak usah declare tipe data
	fmt.Println(io)

	nama := "Salahudin" // simplify kata kunci var hanya boleh dibuat sekali
	fmt.Println(nama)
	
	nama = "Nuaiman"
	fmt.Println(nama)


	var ( // multiple variable
		firstName = "Rio"
		middleName = "S"
		lastName = "Achyar"
	)
	

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
	fmt.Println(firstName,middleName,lastName)
}