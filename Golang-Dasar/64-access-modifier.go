/* Access Modifier
di bahasa pemograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan
access modifier terhadap suatu function atau variabel.
access modifier artinya memberikan akses tertentu dalam kode program
di golang, untuk menentukan access modifier, cukup dengan
nama variabel, function, struct atau lain sebagainya.
jika namanya diawali dengan huruf Besar, maka artinya bisa diakses dari package lain.
jika dimulai dengan huruf kecil, artinya tidak bisa diakses dari package lain.
*/

package main

import "fmt"

// tidak bisa di akses di lain package lain, namu bisa diakses di packe yang sama

func salam(nama string) string {
	return "Assalamu'alaikum... " + nama
}

// bisa di akses di package lain
func Salam(nama string) string {
	return "Assalamu'alaikum... " + nama
}

func main() {

	fmt.Println(salam("Adrik"))
	fmt.Println(Salam("Muhammad"))

}
