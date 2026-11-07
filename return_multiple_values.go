package main

import "fmt"

func getFullName() (string, string) {
	return "Rio", "Achyar"
}

func main(){
	firstName, _ := getFullName() // _ = hiraukan value
	fmt.Println(firstName)
}