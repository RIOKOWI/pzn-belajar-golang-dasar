package main

import "fmt"

func identity()(firstName, middleName, lastName string){
	firstName = "Rio"
	middleName = "Achyar"
	lastName = "nichol"

	return firstName, middleName, lastName 
}


func main(){
	a, b, _ := identity() // a = firstname, b - middlename, _ = abaikan
	fmt.Println(a, b)
}