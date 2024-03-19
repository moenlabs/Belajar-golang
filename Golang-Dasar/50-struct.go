/* Struct
struct adalah template data yang digunakan untuk menggabungkan
nol atau lebih type data lainnya dalam satu kesatuan.
struct biasanya direpresentasi data dalam program aplikasi yang dibuat
data di struct di simpan dalam field atau atribut
maka dari itu struct adalah kumpulan dari field atau atribut

Struct (singkatan dari structure) digunakan untuk membuat kumpulan anggota
dengan tipe data yang berbeda, ke dalam satu variabel.

struct berbeda dengan array, karena array digunakan untuk menyimpan beberapa nilai
dengan tipe data yang sama ke dalam satu variabel,
sedangkan struct digunakan untuk menyimpan beberapa nilai dengan tipe data yang berbeda ke dalam satu variabel.

Struct dapat berguna untuk mengelompokkan data bersama untuk membuat catatan

Untuk mendeklarasikan sebuah struct di Go,
gunakan kata kunci `type` dan `struct`.

Syntax :
type struct_name struct {
  anggota1 datatype
  anggota2 datatype
  anggota3 datatype
}
*/

package main

import "fmt"

type Santri struct {
	Nama, Alamat string
	Usia         int
}

func main() {
	var santri1 Santri
	var santri2 Santri

	// data santri1
	santri1.Nama = "Ahmad"
	santri1.Alamat = "Jogja"
	santri1.Usia = 20

	fmt.Println(santri1)

	// data santri2
	santri2.Nama = "Nur"
	santri2.Alamat = "Magelang"
	santri2.Usia = 17

	fmt.Println(santri2)

	// cuma akses bagian di struc santri 1

	nama := santri1.Nama
	alamat := santri1.Alamat

	fmt.Println(nama, alamat)
	fmt.Println(santri1.Usia)

	// mengganti nama santri 1
	santri1.Nama = "Adrik"

	fmt.Println(santri1)

}
