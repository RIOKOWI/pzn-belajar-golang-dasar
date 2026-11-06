package main

import "fmt"


func NewMap(name string) map[string]string {
	if name == "" {
		return nil // nil cuma bisa di pake oleh tipe data interface, function, map, slice, pointer & channel
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main(){
	var data = NewMap("")
	if data == nil { // kalo data nya nil
		fmt.Println("Data Kosong") // print ini
	} else { // kalo data nya ada
		fmt.Println(data) // print data nya
	}
}