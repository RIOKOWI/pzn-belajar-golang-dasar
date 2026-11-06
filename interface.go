package main

import "fmt"


type HasName interface {
	getName() string // contract: tipe yang punya ini berarti mengimplemen interface
}

func Greet(value HasName) {
	fmt.Println("Halo", value.getName()); // implement contract method yang ada di interface 
}

type Person struct {
	Name string	
}


func (person Person) getName() string{ // cara implement interface ke struct/Person memenuh syarat interface
	return person.Name
}


type Animal struct {
	Jenis string
}

func (animal Animal) getName() string { // Animal memenuhi syarat interface karna memanggil contract/method getName() string
	return animal.Jenis
}


func main() {
	var orang = Person{Name: "Rio"} // membuat value bertipe Person
	Greet(orang); // akses func Greet yang sudah implement contract method interface dan value dari Person

	var binatang = Animal{Jenis: "Mamabayu"} // add value bertipe person
	Greet(binatang); // akses func Greet yang sudah implement interface
}