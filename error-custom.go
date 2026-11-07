package main

import (
	"fmt"
)


type validationError struct { // bikin struct 
	Message string // bikin field untuk message error
}

type notFoundError struct { // bikin struct 
	Message string // bikin field untuk message error
}

func (v *validationError) Error() string { // bikin func receiver pointer pass by reference yang implement interface error
	return v.Message // kembalikan pesan error nya
}

func (v *notFoundError) Error() string { // bikin func receiver pointer pass by reference yang implement interface error
	return v.Message // kembalikan pesan error nya
}

func saveData(id string, data any) error { // buat function (value, isi data) dengan implement interface error
	if id == "" { // jika id nya kosong
		return &validationError{Message: "validaton error:"} // kembalikan pointer pass by value custom error dengan message validation error
	}

	if id != "rio" { // jika id nya bukan rio
		return &notFoundError{Message: "data not found:"} // kembalikan pointer pass by value custom error dengan message data not found
	}

	return nil // ok
}

func main(){
	err := saveData("rioo", nil) // call function
	if err != nil { // jika terjadi error
		// Periksa apakah err bertipe pointer validationError, lalu simpan hasilnya ke validationErr dan status kecocokannya ke ok(boolean).
		if validationErr, ok := err.(*validationError); ok { // jika error nya validationError
			fmt.Println(validationErr.Error(), "Data lu Kosong Brayyyyy") // tampilkan pesan error dari interface error custom

			// Periksa apakah err bertipe pointer notFoundError, lalu simpan hasilnya ke notFoundError dan status kecocokannya ke ok(boolean).
		} else if notFoundErr, ok := err.(*notFoundError); ok { // jika error nya notfoundError
			fmt.Println(notFoundErr.Error(), "Data lu Gak Ketemu")	 // tampilkan pesan error dari interface error custom
		} else { // jika error tidak keduanya
			fmt.Println(err.Error(), "Gatau Error Apaan"); // tampilkan pesan error dari interface error custom
		}
	} else {
		fmt.Print("Success") // jika tidak error
	}
}
