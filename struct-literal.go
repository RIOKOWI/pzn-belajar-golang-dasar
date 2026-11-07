package main

import "fmt"

type Customer struct { // membuat struct
	Name, Address string // pakai PascalCase
	Age           int
}

func main() {
	var rio Customer // bikin variabel untuk memasukkan ke field di dalam struct
	rio.Name = "Rio Achyar" // memasukkan string ke dalam field struct
	rio.Address = "Bandung Resbob"
	rio.Age = 21
	fmt.Println(rio); // akses seluruh isi struct
	fmt.Println(rio.Name); // akses field Name


	// STRUCT LITERALS
	jokowi := Customer{ // masukin data ke field bisa juga kayak gini
		Name: "Pangerang Nipunegoro",
		Address: "Solo",
		Age: 69,
	}
	fmt.Println(jokowi)
	
	fufufafa := Customer{"Gibran", "Solo", 10} // masukin data ke field bisa juga kayak gini
	fmt.Println(fufufafa)
}