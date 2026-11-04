package main

import "fmt"


func logging(){ // membuat function logging untuk digunakan secara defer
	fmt.Println("Selesai Memanggil Logging")
}

func runApp(){
	defer logging(); // call function logging secara defer == di panggil setelah func runApp berhasil/maupun gagal
	fmt.Println("Run Application");
}

func main(){
	runApp(); // call func
}