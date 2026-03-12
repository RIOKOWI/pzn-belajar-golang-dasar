package main

import "fmt"

func main(){
	a := 5
	b := 10
	d := 7
	e := 20
	c := a + b - d * e
	fmt.Println(c)


	// augmented assigments
	i := 10
	i += 30
	i += 100
	fmt.Println(i)


	// unary operator
	j := 1
	j++
	j++
	fmt.Println(j)
	j--
	fmt.Println(j)
}