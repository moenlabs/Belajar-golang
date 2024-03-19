/* Variadic Function
Parameter yang berada di Posisi terahir,
memiliki kemampuan dijadikan sebuah varargs (varabel arguments).
Varargs artinya datanya bisa menerima lebih dari satu input,
atau anggap saja semajam array.
apa bedanya dengan parameter biasa dengan tipe data array?
	1. jika parameter array kita wajib membuat arraynya terlebih dahulu
	2. jika parameter menggunakan varargs, kisa bisa langsung mengirim datanya,
	jika lebih dari satu, cukup gunakan tanda koma

Slice Parameter
terkadang ada kasus di mana kita menggunakan variadic function,
namun memiliki variabel berupa slice.
kita bisa menjadikan slice sebagai varargs parameter
*/

package main

import "fmt"

// contoh variadic function array
func penjumlahan(nomor ...int) int {
	total := 0

	for _, number := range nomor {
		total += number

	}
	return total
}

// contoh variadic function slice tapi harus manual slicenya
func perkalian(angka []int) int {
	hasil := 0

	for _, bilangan := range angka {
		hasil += bilangan

	}
	return hasil
}

func main() {
	total := penjumlahan(5, 7, 9)
	hasil := perkalian([]int{2, 4, 1})
	fmt.Println("Total Penjumlahan Variadic dengan Array: ", total)
	fmt.Println("Hasil dari Perkalian Variadic dengan Slice:", hasil)

	// contoh slice jadi parameter
	numbers := []int{1, 3, 5, 7} // ini slice

	// ini menjadikan slice sebagai parameter penjumlahan
	fmt.Println("ini adalah hasil dari menjadikan slice menjadi parameter di penjumlahan: ", penjumlahan(numbers...))
}
