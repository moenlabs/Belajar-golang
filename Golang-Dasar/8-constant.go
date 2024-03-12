/* Materi Constant
constant merupakan variable yang nilainya tidak bisa diubah lagi
setelah pertama kali diberi nilai.
cara untuk membuat constant mirip dengan variable,
yang membedakan hanya kata kunci yang digunakan adalah const, bukan var.
saat pembuatann constant, wajib langsung
menginisialisasi datanya.
perbedaan const dan var adalah, jika 
var harus dipakai, sedangkan const 
todak terjadi error jika tidak dipakai.
*/

// contoh kode 
package main

import "fmt"
func main () {
	const fristName = "misbah"
	const lastName = "munir"

	fmt.Println(fristName,lastName)
}
