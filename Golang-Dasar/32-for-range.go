/* For Range
for bisa digunakan untuk melakukan iterasi terhadap semua data collection
data collection contohnya Array, Slice dan Map
dengan for kita bisa melakukan perulangan secara manual
untuk array, slice dan map.
dan bisa dengan cara range
Kata kunci `range` digunakan untuk mengulang
array, slice, atau map dengan lebih mudah.
Kata kunci ini mengembalikan index dan valuenya.

Syntax :
for index, value := array|slice|map {
   // kode yang akan dieksekusi untuk setiap iterasi
}

Tip: Untuk hanya menampilkan value atau index,
kita bisa menghilangkan for lainnya
dengan menggunakan underscore (_).
*/

package main

import "fmt"

func main() {

	// contoh perulangan for manual dari array
	nama := [...]string{"zaid", "umar", "bakar"}

	for i := 0; i < len(nama); i++ {
		fmt.Println(nama[i])
	}
	// contoh perulangan for dengan range dari slice
	komputer := []string{"Macbook", "Hp", "Lenovo", "Asus"}

	for index, jenis := range komputer {
		fmt.Println("Index ke: ", index, "komputer", jenis)
	}

	// contoh jika tidak butuh index di ganti underscore
	komputer1 := []string{"Macbook", "Hp", "Lenovo", "Asus"}

	for _, jenis := range komputer1 {
		fmt.Println("komputer", jenis)
	}

	// contoh jika tidak butuh valuenya
	komputer2 := []string{"Macbook", "Hp", "Lenovo", "Asus"}

	for index, _ := range komputer2 {
		fmt.Println("index ke: ", index)
	}

	// contoh hanya akses key nya di map dengan kata kunci tetep index
	inventaris := map[string]string{"Komputer": "Macbook", "Handphone": "Apple", "Kamera": "Sonny"}

	for index, _ := range inventaris {
		fmt.Println("Nama Barang: ", index)
	}

	alat := []string{"Pacul", "Sekop", "Palu"}

	for _, nama := range alat {
		fmt.Println("Nama Alat: ", nama)
	}
}
