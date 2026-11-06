package main

import "fmt"

type Address struct { // membuat tipe data struct
	City, Province, Country string // buat field
}

// func changeCountryToIndonesia(address Address) { // tidak akan berubah karna Address copy ke memori lain
// 	address.Country = "Indonesia"
// }

func changeCountryToIndonesia(address *Address) { //  set parameter menjadi pointer
	address.Country = "Indonesia"
}

func main(){
	change := Address{"Tangerang", "Banten", ""};
	changeCountryToIndonesia(&change) // kirim ke pointer yang sama atau pass by reference ke Address

	fmt.Println(change)
}