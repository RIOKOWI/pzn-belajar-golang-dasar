package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main(){
	byeBye := getGoodBye // variabel function
	fmt.Println(byeBye("Rio"))
}