package main

import "fmt"


func random() interface{} {
	return "Ok"
}


func main(){
	var hasil = random()
	var hasilString string = hasil.(string) // ubah sesuai dengan tipe data
	fmt.Println(hasilString)


	var hasilInt int = hasil.(int) // kalo konversi tidak sesuai dengan tipe data pasti langsung panic
	fmt.Println(hasilInt)
}