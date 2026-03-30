package main

import "fmt"

func sumAll(numbers ...int) int { // varargs = variabel argumen
	total := 0
	for _, number := range numbers{
		total += number
	}

	return total
}

func main(){
	bedugul := sumAll(10, 10, 10, 20, 30, 40)
	fmt.Println(bedugul)
}