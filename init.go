package main

import (
	"pzn-belajar-golang-dasar/database" // import package "nama module yang ada di go.mod/direktori"
	"fmt"
)


func main(){
	fmt.Println(database.GetDatabase()); // connect ke database
}