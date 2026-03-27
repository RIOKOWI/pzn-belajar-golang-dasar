package main

import "fmt"

func main(){
	var nilaiAkhir = 81
	var absensi = 70

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var absensiAkhir bool = absensi > 80

	// var lulus bool = lulusNilaiAkhir && absensiAkhir // tidak lulus
	var lulus bool = lulusNilaiAkhir || absensiAkhir // lulus

	fmt.Println(lulus)
}