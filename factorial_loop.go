package main

import "fmt"

func factorialLoop(value int) int { // membuat function untuk factorila loop yang mengembalikan integer
	result := 1; // set default value

	for i := value; i > 0; i-- { // membuat perulangan untuk factorial loop decrement
		result *= i // setiap decrement di kali 1 sampai 0
	}

	return result // mengembalikan value secara integer yang sudah di set di awal pemuatan func
}


// recursive function
func factorialRecursive(value int) int {
	if value == 1 { // jika value nya = 1
		return 1 // kembalikan / kalikan dengan 1
	} else { // kecuali
		return value * factorialRecursive(value-1) // value nya bukan 1 akan dikalikan dengan value yang ada secara decrement(-) secara terus menerus sampai 1
	}
}

func main() {
	fmt.Println(factorialLoop(5)); // test function dan mengirimkan parameter
	result := 5 * 4 * 3 * 2 * 1; // test manual
	fmt.Println(result);


	fmt.Println(factorialRecursive(5)); // call recursive function
}