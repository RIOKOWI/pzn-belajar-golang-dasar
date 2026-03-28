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
	fmt.Println(len(person))
	
	
	person["name"] = "Rio Nichol"
	fmt.Println(person)
	
	president := make(map[string]string) // buat map baru
	president["name"] = "Obama"
	president["country"] = "Amerika"
	president["food"] = "Bakso"
	fmt.Println(president)
	
	delete(president, "food") // hapus key map
	fmt.Println(president)



	
}