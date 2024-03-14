/* Else If
kadang dalam if, kita butuh membuat beberapa kondisi
kasus seperti ini, kita bisa menggunakan else if

pernyataan else if untuk menentukan kondisi baru
jika kondisi pertama false (salah)
*/

package main

import "fmt"

func main() {

	// kode variabel
	angka := 25

	// kode if
	if angka < 11 {
		fmt.Println("Angka 11")
	} else if angka < 30 { // kode else if
		fmt.Println("Angka 30")
	} else { // kode else
		fmt.Println("Angka salah")
	}

	// contoh 2
	nama := "Andi"

	if nama == "Zaid" {
		fmt.Println("Namaku Zaid")
	} else if nama == "Budi" {
		fmt.Println("Namaku bukan budi")
	} else if nama == "Budi" {
		fmt.Println("Namaku adalah Budi")
	} else if nama == "Andi" {
		fmt.Println("Namaku Andi")
	} else {
		fmt.Println("Namaku Tidak Tahu")
	}

}
