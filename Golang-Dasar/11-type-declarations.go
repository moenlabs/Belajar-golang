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

	// type declaration banya

	type (
		Tipe1 string
		Tipe2 int
	)

	type (
		Variasi1 = string
		Variasi2 = int
	)

	var tipe1 Tipe1 = "zaid"
	var tipe1a Tipe1 = "umar"
	gabungkan := Tipe1(tipe1)

	var tipe2 Tipe2 = 20
	gabunganTipe2 := (tipe2)

	var variasi1 Variasi1 = "Bakar"
	var variasi2 Variasi2 = 20

	fmt.Println(nomor, no)
	fmt.Println(jujur)
	fmt.Println(salah, nilai)
	fmt.Println(rt, rw1)
	fmt.Println(bagi, hasil)

	// cetak banyak type declaration

	fmt.Println(tipe1a, gabungkan)
	fmt.Println(gabunganTipe2)

	fmt.Println(variasi1, variasi2)

}
