package main

import (
	"errors" // panggil package errors
	"fmt"
)


func Pembagian(value int, bagi int) (int, error) { // (int, error) = mengembalikan data hasilnya, dan kalo terjadi error kembalikan errornya
	if bagi == 0 {
		return 0, errors.New("Dilarang pembagian 0") // return data, dan nil. errors.New tuh masukin pesan error ke interface error
	} else {
		return value / bagi, nil // return data, dan nil
	}
}

func main(){
	result, err := Pembagian(9, 0) // buat dua variabel hasil dan data error untuk function
	if err == nil { // jika error nya nihil atau nil
		fmt.Println("Hasil", result) // tampilkan pesan dan data hasilnya
	} else { // jika error nya ada
		fmt.Println("Error:", err.Error()); // tampilkan pesan error dari method Error() yang dari interface error
	}
}