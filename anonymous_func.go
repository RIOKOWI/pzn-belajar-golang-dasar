package main
import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name){
		fmt.Println("You are Blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}


func main(){
	blackList := func(name string) bool { // bikin anonymous function kayak gini bisa
		return name == "anjing"
	}

	registerUser("Rio", blackList)

	registerUser("Guguk", func(name string) bool { // bikin anonymous function kayak gini juga bisa
		return name == "Guguk"
	})

}