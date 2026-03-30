package main
import "fmt"

type Filter func(string) string // membuat alias agar gak kepanjagan di parameter

func sayHelloWithFilter(name string, filter Filter){
	filteredName := filter(name)
	fmt.Println("Hello " + filteredName)
}

func spamFilter(name string) string { // ini dipaakai untuk parameter func sayHelloWithFilter
	if name == "anjing" {
		return "****"
	} else {
		return name
	}
}


func main(){
	sayHelloWithFilter("Rio", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("anjing", filter)
}