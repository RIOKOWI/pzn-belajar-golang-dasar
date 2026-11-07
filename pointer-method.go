package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() { // membuat method dengan pointer. sangat di rekomendasikan
	man.Name = "Mr. " + man.Name
}

func main(){
	rio := Man{"Rio"}
	rio.Married()

	fmt.Println(rio.Name)
}