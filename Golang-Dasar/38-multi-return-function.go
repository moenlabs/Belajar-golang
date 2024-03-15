/* Returning Multiple Value
function tidak hanya mengembalikan satu value,
tetapi juga bisa multiple value.
Untuk memberitahu jika function mengembalikan multiple value,
kita harus menulis semua tipe data returnnya di function.

multiple return value wajib ditangkap semua valuenya,
jika kita ingin tidak menangkap return valuenya,
kita bisa menggunakan tanda underscore (_)
*/

package main

import "fmt"

func getNamaLengkap() (string, string) {
	return "Muhammad", "Nur"
}

func uangSaku() (string, int) {
	return "Muhammad", 2000
}

func main() {

	namaDepan, namaBelakang := getNamaLengkap()
	fmt.Println(namaDepan, namaBelakang)

	// contoh tidak menangkap return value
	_, uang := uangSaku()
	fmt.Println(uang)
}
