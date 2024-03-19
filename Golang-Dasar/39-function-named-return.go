/* Named Return Values

Biasanya saat kita memberi tahu bahwa sebuah function
mengembalikan value, maka kita hanya mendeklarasikan tipe data return value di function.
namun, kita juga bisa membuat variable secara langsung di tipe data return functionnya.
*/

package main

import "fmt"

func inventaris() (komputer, laptop, hp int) {
	komputer = 5
	laptop = 2
	hp = 10

	return komputer, laptop, hp
}

func daftarStock() (makanan string, jumlah int) {
	makanan = "jagung"
	jumlah = 10

	return makanan, jumlah
}

func rakaatShalat() (shalat string, rakaat int8) {
	shalat = "Dhuhur"
	rakaat = 4

	return shalat, rakaat
}

func penjumlahan(x int, y int) (hasil int) {
	hasil = x + y
	return
}

func perkalian(l string, m int, n int) (ucapan string, hasil int) {
	ucapan = l
	hasil = m * n
	return
}

func main() {
	a, b, _ := inventaris()
	fmt.Println(a, b)
	fmt.Println(inventaris())
	fmt.Println(daftarStock())
	fmt.Println(rakaatShalat())
	fmt.Println(penjumlahan(20, 30))
	fmt.Println(perkalian("Hasil Perkalian: ", 11, 7))
}
