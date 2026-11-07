package main

import "fmt"

func main(){
	var names [3]string // tidak bisa lebih dari 3 array
	names[0] = "rio"
	names[1] = "achyar"
	names[2] = "dudung"
	

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// cara buat array secara langsung

	var values = [3]int{
		1,
		2,
		// data ketiga defaultnya kosong
	}

	fmt.Println(values)


	// function array
	var aruy = [...]string{ // ... adalah jumlah tidak ditentukan
		"io",
		"rio",
		"iyok",
	}

	aruy[2] = "bedul" // ubah isi array

	fmt.Println(len(aruy))
	fmt.Println(aruy)

}