/* For Loop


Perulangan for mengulang blok kode beberapa kali.
Perulangan for adalah satu-satunya perulangan yang tersedia di Go.

Perulangan sangat berguna jika Anda ingin menjalankan kode
yang sama berulang kali, dengan nilai yang berbeda.

dalam for, kita bisa menambahkan statement (pernyataan)
di mana terdapat 2 statement yang bisa di tambahkan di for

init statement : statement sebelum for di eksekusi
post statement: statement yang akan selalu di eksekusi
di akhir setiap perulangan

Setiap eksekusi dari sebuah perulangan disebut iterasi.
Perulangan for dapat terdiri dari tiga pernyataan

Syntax :
for pernyataan1; pernyataan2; pernyataan3 {
   // kode yang akan dieksekusi untuk setiap iterasi
}

pernyataan1: Menginisialisasi nilai penghitung perulangan.
pernyataan2 : mengevaluasi setiap iterasi perulangan.
Jika bernilai TRUE, perulangan berlanjut.
Jika bernilai FALSE, perulangan berakhir.
pernyataan3 : Meningkatkan nilai penghitung perulangan.
*/

package main

import "fmt"

func main() {

	// Contoh Sederhata For
	perulangan1 := 1

	for perulangan1 <= 10 {
		fmt.Println("Perulangan Ke :", perulangan1)
		perulangan1++
	}

	fmt.Println("Finnish")

	// contoh init dan post statement dalam for
	for perulangan2 := 5; perulangan2 <= 7; perulangan2++ {
		fmt.Println("Perulangan ke: ", perulangan2)
	}

	fmt.Println("Selesai")

	// latihan ulang for loop
	for counter := 1; counter <= 5; counter++ {
		fmt.Println("Counter ke: ", counter)
	}
}
