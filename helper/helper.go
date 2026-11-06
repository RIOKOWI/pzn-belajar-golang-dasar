package helper // samakan dengan nama direktori

var version = "1.26.1" // gabisa di akses oleh package lain karna huruf depan variable nya tidak kapital
var Application = "Golang" // bisa di akses oleh package lain

func sayGoodbye(name string) string { // gabisa di akses oleh package lain karna huruf depan func nya tidak kapital
	return "Dadah " + name
}

func SayHello(name string) string {
	return "Hello " + name
}