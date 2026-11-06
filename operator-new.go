package main

import "fmt"

type Address struct { // membuat tipe data struct
	City, Province, Country string // buat field
}

func main(){

	address1 := new(Address) // untuk yang data awal nya kosong saja
	address2 := address1 

	address2.Province = "Tidore"

	fmt.Println(address1) 
	fmt.Println(address2) 
}