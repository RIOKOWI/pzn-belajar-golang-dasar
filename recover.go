package main

import "fmt"

func end(){
	fmt.Print("program kelar ")
	// contoh kode program recover yang benar yang nantinya func dipanggil dengan defer sehinggan setelah panic kode recover masih bisa di jalankan
	message := recover()
	fmt.Println("Terjadi Error", message);
}

// func run(error bool){
// 	defer end() // masih jalan mau panic/nggak karna defer
// 	if error {
// 		panic("BREAK PROGRAM") // panic == benar benar menghentikan program
// 	}
// 	// dibawah adalah contoh kode recover yang salah, soalnya setelah kode panic kode program akan berhenti sehingga recover tidak dijalankan kecuali defer function
// 	message := recover()
// 	fmt.Println("Terjadi Erro", message);
// }

func run(error bool){
	defer end() // masih jalan mau panic/nggak karna defer
	if error {
		panic("BREAK PROGRAM") // panic == benar benar menghentikan program
	}
}

func main(){
	run(true) // true == panic program bakal berhenti defer func masih jalan, false defer func tetap jalan tetapi panic tidak
}
