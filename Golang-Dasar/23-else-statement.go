/* Else statement
jika if statement akan mengeksekusi block kode jika bernilai true
maka jika ingin mengeksekusi program yang bernilai false
menggunakan statement else
artinya : else digunakan untuk menentukan blok kode
yang akan dieksekusi jika kondisinya bernilai false (salah)
*/

package main

import "fmt"

func main() {

	// variabelnya
	namamu := "kamu"
	namaku := "aku"

	// if statement
	if namamu == namaku {
		fmt.Println("Nama kita sama")
	} else { // kode else
		fmt.Println("Nama kita berbeda")
	}

}
