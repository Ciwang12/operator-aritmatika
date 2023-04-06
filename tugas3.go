package main

import "fmt"

func main() {

	// ========login=========

	var nama, sandi string

	fmt.Print("masukan nama : ")
	fmt.Scanln(&nama)
	fmt.Print("masukan sandi : ")
	fmt.Scanln(&sandi)

	if sandi == "1234" && nama == "ciwang" {
		fmt.Println("selamat datang  (=_=)")
	} else if sandi == "1234" {
		fmt.Println("nama anda tidak sesuai !!!")
	} else if nama == "ciwang" {
		fmt.Println("sandi anda salah !!!")
	} else {
		fmt.Println("anda penyusufff !!!")
	}

}
