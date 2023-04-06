package main

import "fmt"

func main() {

	// ======menentukan ganjil genap=========

	var angka int8

	fmt.Print("masukan angka : ")
	fmt.Scanln(&angka)

	if angka%2 == 0 {
		fmt.Println(angka, "adalah angka genap")
	} else {
		fmt.Println(angka, "adalah angka ganjil")
	}
}
