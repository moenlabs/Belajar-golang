/* Function Return
sebuah function bisa mengembalikan data
untuk memberitahu bahwa function tersebut mengembalikan data,
kita harus menuliskan tipe data kembalian dari function tersebut.
jika function tersebut kita deklarasikan dengan tipe data pengembalian,
maka wajib di dalam functionny kita harus mengembalikan data.
untuk mengembalikan data dari function, kita bisa menggunakan kata kunci
return, diikuti dengan datanya.

artinya :
Jika kita ingin sebuah function mengembalikan sebuah value,
kita harus menentukan tipe data dari value yang dikembalikan
(seperti int, string, dll), dan juga menggunakan kata kunci
return di dalam sebuah function tersebut

Syntax :
func FunctionName(param1 type, param2 type) type {
  // code to be executed
  return output
}
*/

package main

import "fmt"

// contoh return function 1
func penambahan(a int, b int) int { // int ini tipe data dari return
	return a + b // value dari return
}

func sapaan(salam string) string {

	return "Assalamualaikum," + salam
}

func nama(nama string) string {
	return "Nama saya " + nama
}

func main() {

	hasil := penambahan(5, 10) // mengambil return dari function dan inisialisasi data di function
	fmt.Println(hasil)
	fmt.Println(sapaan("Zaid"))

	namaSaya := nama("Adrik")
	fmt.Println(namaSaya)
}
