package main

import (
	"pzn-belajar-golang-dasar/helper" // import package dari dir helper
	"fmt"
)

func main (){
	hasil := helper.SayHello("Rio") // call package.func
	fmt.Println(hasil)
}