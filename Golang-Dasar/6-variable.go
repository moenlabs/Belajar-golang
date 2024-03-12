/* Variable
Variable adalah tempat untuk menyimpan data
variable digunakan agar kita bisa mengakses data 
yang sama di manapun kita mau
di Golang, variable hanya bisa menyimpan tipa data yang sama,
jika kita ingin menyimpan data yang berbeda-beda jenis, kita harus membuat
beberapa variable
untuk membuat variable kita bisa menggunakan kata kunci var
lalu diikuti nama variable dan tipe datanya.
*/

// contoh kode 

package main

import "fmt"
/*
func main() {
	var name string

	name = "Misbahul Munir"
	fmt.Println(name)

	name = "moenlabs"
	fmt.Println(name)
}
*/
/* Tipe Data Variable
saat kita membuat variable maka kita wajib menyebutkan tipe data variable tersebut
Namun, jika kita langsung menginisialisasikan data pada variablenya
maka kita tidak wajib menyebutkan tipe datanya.
*/

// Kode Program
/*
func main(){

	var nama = "Misbah"
	fmt.Println(nama)

	nama = "Munir"
	fmt.Println(nama)
}
*/

/* kata kunci Var
di Golang kata kunci variable tidaklah wajib
asalkan saat membuat variable kita langsung menginisialisasi datanya
agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan
kata kunci := (hanya untuk deklarasi pertama, jika ingin mengubah lagi, jangan gunakan := tapi cukup =) saat menginisialisasi data pada variable tersebut
*/

// kode Program

func main() {
	name := "Misbah"
	fmt.Println(name)

	name = "Munir"
	fmt.Println(name)
}