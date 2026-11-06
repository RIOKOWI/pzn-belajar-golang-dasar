package main

import "fmt"

type Address struct { // membuat tipe data struct
	City, Province, Country string // buat field
}

func main(){

	address1 := Address{"Tangerang", "Banten", "Indonesia"}
	address2 := &address1 // address2 adalah reference ke &address1 pointer pass by reference

	// address2 = &Address{"Chong", "Qing", "China"} // beda memori tidak reference ke address1
	*address2 = Address{"Chong", "Qing", "China"} // siapapun yag di address1 akan pindah ke address2 karna asterisk operator

	fmt.Println(address1) // data nya tidak berubah
	fmt.Println(address2) 
}