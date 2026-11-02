package main

import "fmt"


func main() {
	counter := 0; // set default value

	increment := func () { // membuat anonymous func
		fmt.Println("Increment"); // print text
		counter++ // call variabel and the increment di variable
	}

	increment() // call func
	increment()
	increment()
	fmt.Println(counter); // value == how much we call the func


	// closure membuat AMBIGU karna kita tidak pernah akses variabel counter dan hanya mengakses function increment tetapi value dalam variable counter bertambah jika kita call func increment karna closure bisa memodifikasi variabel dari scope luarnya. maka bijaklah memakai closure
}