package main

import "fmt"

func main() {

	// ========mencari nilai terbesar=========
	var (
		angka1 int
		angka2 int
		angka3 int
	)

	fmt.Print("masukan angka pertama : ")
	fmt.Scanln(&angka1)

	fmt.Print("masukan angka kedua : ")
	fmt.Scanln(&angka2)

	fmt.Print("masukan angka ketiga : ")
	fmt.Scanln(&angka3)

	if angka1 > angka2 && angka1 > angka3 {
		fmt.Print("angka petrama terbesar = ", angka1)
	} else if angka2 > angka1 && angka2 > angka3 {
		fmt.Print("angka kedua terbesar = ", angka2)
	} else if angka3 > angka1 && angka3 > angka2 {
		fmt.Print("angka ketiga terbesar = ", angka3)
	}
}
