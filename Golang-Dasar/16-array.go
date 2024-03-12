/* Array
Array adalah tipe data yang berisikan kumpulan data dengan tipe data yang sama
Saat membuat array, perlu menentukan jumlah data yang bis aditampung oleh array tersebut
Daya tampung array tidak bisa bertambah setelah array dibuat.
Saat membuat array, terdapat yang namanya index.
index adalah nomor urutan dari data yang disimpan dalam aray.
index di array dimulai dari 0
misaln array tersebut berisi 1. aku, 2. kamu, 3. dia
maka berarti 0 untuk nomor aku, 1 untuk nomor kamu, dan 2 untuk nomor dia.
*/

// contoh deklarasi array

package main

import "fmt"

func main() {

	// contoh 1 deklarasi array
	var nama = [3]string{"aku", "kamu", "dia"}

	// contoh 2 deklarasi array tanpa menentukan panjang
	var kelompok = [...]string{"satu", "dua", "tiga"}

	// contoh 3 deklarasi array
	nomor := [5]int8{1, 2, 3, 4, 5}

	// contoh 4 deklarasi array tanpa menentukan panjang array
	urutan := [...]int8{1, 2, 3}

	//contoh 3 mengisi array kosong

	var data [3]string
	data[0] = "zaid"
	data[1] = "umar"
	data[2] = "bakar"

	// contoh error menambahkan array
	// data[3] = "fatimah"
	// urutan[3] = 4

	// contoh merubah array
	data[2] = "siapa"
	urutan[2] = 10

	// contoh array kosong
	kosong := [5]int{}

	// contoh error tanpa menentukan panjang array namun tidak langsung di isi

	/* var mobil [...]string // nvalid use of [...] array (outside a composite literal)
	mobil[0] = "avansa"
	mobil[1] = "innova"
	mobil[2] = "pajero" */

	fmt.Println(nama)
	fmt.Println(kelompok)
	fmt.Println(nomor)
	fmt.Println(urutan)
	fmt.Println(data)
	fmt.Println(data[2]) // mengambil data array
	fmt.Println(kosong)
	// fmt.Println(mobil)
}
