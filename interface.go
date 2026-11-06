package main

import "fmt"


type HasName interface {
	getName() string 
}

func Greet(value HasName) {
	fmt.Println("Halo", value.getName()); // implement contract method yang ada di interface 
}