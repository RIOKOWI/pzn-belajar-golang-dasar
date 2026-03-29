package main

import "fmt"

func sayID(fisrtName string, age int) {
	fmt.Print("halo ", fisrtName, "berumur ", age)
}

func main(){
	sayID("rio ", 19)
}