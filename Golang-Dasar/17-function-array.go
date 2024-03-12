package main

import "fmt"

func main() {

	nama := [2]string{"aku", "kamu"}

	//untuk mendapatkan panjang array
	panjang := len(nama)

	// mendapatkan data diposisi index
	posisi := nama[1]

	// mengubah data di posisi index
	nama[1] = "tetapKamu"

	fmt.Println(panjang)
	fmt.Println(posisi)
	fmt.Println(nama)
}
