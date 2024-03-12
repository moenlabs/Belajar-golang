/* Type Deklarasi
type declaration merupakan kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
type daclaration biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada
dengan tujuan agar mudah di mengerti
*/

// contoh code

package main

import "fmt"

func main() {

	type noHp string

	var nomor noHp = "085643182929"
	var nomorHp = "08237829236"
	var no noHp = noHp(nomorHp)

	type benar bool

	var jujur benar = true
	var nilai = false
	var salah benar = benar(nilai)

	type angka int
	var rt angka = 30
	var rw = 40
	var rw1 angka = angka(rw)

	type koma float32
	var bagi koma = 3.800
	var kali = 4.00
	var hasil koma = koma(kali)

	fmt.Println(nomor, no)
	fmt.Println(jujur)
	fmt.Println(salah, nilai)
	fmt.Println(rt, rw1)
	fmt.Println(bagi, hasil)

}
