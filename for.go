package main

import "fmt"

func main(){
	counter := 1


	for counter <= 10 {
		fmt.Println("Perulangan ke :", counter)
		counter++
	}


	// for statement
	for tung := 0; tung < 5; tung++ {
		fmt.Println("ke", tung)
	}
}