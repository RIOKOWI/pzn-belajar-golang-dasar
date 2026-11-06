package main

import "fmt"

type Customer struct { // membuat struct
	Name, Address string // pakai PascalCase
	Age           int
}

// STRUCT METHOD
func (customer Customer) greet(name string){ // membuat func struct method (receiver Struct) namaFunc(param tipe data)
	fmt.Println("Halo", name, "Nama Saya", customer.Name); // ucapkan greet + akses field Name dari receiver struct method
}

func main() {
	var rio Customer // bikin variabel untuk memasukkan ke field di dalam struct
	rio.Name = "Rio Achyar" // memasukkan string ke dalam field struct
	rio.Address = "Bandung Resbob"
	rio.Age = 21
	fmt.Println(rio); // akses seluruh isi struct
	fmt.Println(rio.Name); // akses field Name


	// STRUCT METHOD
	var jono = Customer{Name: "Audi"}; // panggil instance Customer
	jono.greet("Rio"); // call the method
}