package main

import "fmt"


func main(){
	// continue 
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("ganjil", i)
	}
}