/* If Statment
If merupakan salah satu kata kunci yang digunakan
untuk percabangan.
percabangan artinya kita bisa mengeksekusi kode program
tertentu ketika suatu kondisi terpenuhi
Hampir di semua pemrograman mendukung if statement

if statement digunakan untuk menentukan blok kode Go yang di mulai
dengan tanda kurung kurawal {} akan dieksekusi jika suatu kondisi
bernilai true.
*/

package main

import "fmt"

func main() {

	// contoh if statement
	if 100 > 76 {
		fmt.Println("100 Lebih besar dari pada 76")
	}

	nilai := 30

	if nilai > 50 {
		fmt.Println("Nilai Lebih Besar Dari 50")
	} else {
		fmt.Println("Nilai Lebih Kecil Dari 50")
	}

}
