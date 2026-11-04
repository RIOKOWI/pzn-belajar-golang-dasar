package main

import "fmt"

func end(){
	fmt.Print("program kelar ")
}

func run(error bool){
	defer end() // masih jalan mau panic/nggak karna defer
	if error {
		panic("BREAK PROGRAM") // panic == benar benar menghentikan program
	}
}

func main(){
	run(true) // true == panic program bakal berhenti defer func masih jalan, false defer func tetap jalan tetapi panic tidak
}
