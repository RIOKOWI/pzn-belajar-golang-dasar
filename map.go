package main

import "fmt"

func main(){
	person := map[string]string{
		"name": "rio achyar",
		"address": "puri jaya pasar kemis",
	}


	fmt.Println(person)
	fmt.Println(person["name"]) // name & address = key
	fmt.Println(person["address"])
	
}