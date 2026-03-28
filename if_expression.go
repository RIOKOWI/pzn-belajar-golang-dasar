package main

import "fmt"

func main(){
	name := "rio"


	if name == "rio" {
		fmt.Print("halo ", name)
	} else if name == "jokowi" {
		fmt.Print("hidup jokowi")
	} else {
		fmt.Print("salah brayyyyy")	
	}
	
	fmt.Print("\n")	
	
	president := "jokowi"
	
	if length := len(president); length >= 6 {
		fmt.Print("hidup ", president)	
	}
}