package main

import (
	"pzn-belajar-golang-dasar/database" // import package "nama module yang ada di go.mod/direktori"
	_ "pzn-belajar-golang-dasar/internal" // _ = adalah blank indentifier untuk menjalankan init di package lain tanpa harus call function di package tersebut
	"fmt"
)


func main(){
	fmt.Println(database.GetDatabase()); // connect ke database
}