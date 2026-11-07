package main

import "fmt"

func main(){
	type NOKTP string // NOKTP = alias. jadi NOKTP adalah string

	var ktpNyaRio NOKTP = "111122409384329"
	var contoh = "222222"
	var contohKtp NOKTP = NOKTP(contoh)

	fmt.Println(ktpNyaRio)
	fmt.Println(contoh)
	fmt.Println(contohKtp)
}