package main

import "fmt"

func Ups() any { // any untuk interface kosong
	// return 1
	return "Ups"
}


func main(){
	var kosong any = Ups()
	fmt.Println(kosong)


	// interface kosong adalah tipe data paling tinggi
}

