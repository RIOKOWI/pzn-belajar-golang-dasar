package database // membuat packaage database

var connection string // membuat variable tipe data string

func init(){ // membuat func init
	connection = "MySQL" // isi variabel logic connect ke database nantinya
}

func GetDatabase() string{ // func access modifier untuk di panggil di package lain
	return connection // mengembalikan connection
}