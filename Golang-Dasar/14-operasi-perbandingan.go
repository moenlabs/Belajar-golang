/* Operasi Perbandingan
operasi perbandingan adalah operasi untuk membandingkan
dua buah data
operasi perbandingan adalah operasi yang menghasilkan nilai boolean (benar atau salah)
jika hasil operasinya adalah salah, maka nilainya false
berikut ini adalah operasi perbandingan :

> : lebih dari
< : kurang dari
>= : lebih dari sama dengan
<= : kurang dari sama dengan
== : sama dengan
!= : tidak sama dengan

Penting : tidak semua data bisa dibandingkan, untuk data int bisa menggunakan
operator perbandingan. namun untuk tipe data string bisa menggunakan
operator == atau !=
*/

// Contoh kode

package main

import "fmt"

func main() {

	namaKu := "zaid"
	namaMu := "umar"

	namaKita := namaKu == namaMu
	// membandingkan string tidak error namun tidak kompatibel karena hasilnya berdasarkan urutan huruf alpabhet,
	namaSalah := namaKu > namaMu

	// urutan alphabet bahwa nilai dari itu ter kecil z itu terbesar
	n := "c"
	m := "d"

	nm := n < m

	nomorMu := 100
	nomorKu := 50

	nomorKita := nomorMu <= nomorKu

	fmt.Println(namaKita)
	fmt.Println(namaSalah)
	fmt.Println(nm)
	fmt.Println(nomorKita)
}
