package main

import "fmt"


func random() interface{} {
	return true
}


func main(){
	// var hasil = random()
	// var hasilString string = hasil.(string) // ubah sesuai dengan tipe data
	// fmt.Println(hasilString)


	// var hasilInt int = hasil.(int) // kalo konversi tidak sesuai dengan tipe data pasti langsung panic
	// fmt.Println(hasilInt)

	result := random() // menyimpan return value func rando()
	// type assertion switch
	switch value := result.(type){ // kita cek type nya
	case string: // kalo string
		fmt.Println("String", value) // tampilkan String dan value/tipe data dari interface
	case int: // kalo int
		fmt.Println("Int", value) // tampilkan int dan value/tipe data dari interface
	default: // kalo any/kosong
		fmt.Println("Unknown", value) // tampilkan unknown
	}
}