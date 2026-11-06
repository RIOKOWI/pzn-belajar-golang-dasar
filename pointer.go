package main

import "fmt"

type Address struct { // membuat tipe data struct
	City, Province, Country string // buat field
}

func main(){
	address1 := Address{"Tangerang", "Banten", "Indonesia"}
	address2 := address1 // copy value address1 ke address2 atau istilahnya pass by value

	address2.City = "Kediri" // ubah value dari field City di address2

	fmt.Println(address1) // data nya masih sama tidak berubah
	fmt.Println(address2)
}