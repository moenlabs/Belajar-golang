/*Interface Kosong

Golang bukanlah bahasa pemograman berorientasi object
Biasanya dalam bahasa pemograman berorientasi object, ada satu data parent di puncak
yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemograman tersebut.
di Java ada java.lang.Object

Untuk menangani kasus ini, di Golang kita bisa menggunakan interface kosong
interface kosong adalah interface yang
tidak memiliki deklarasi method satupu,
hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya
interface kosong, juga memiliki type alias bernama any.

penggunaan interface kosong di golang seperti
fmt.Println(a ...interface[])
panic(v interface[])
recover()interface[]
dan lain-lain
*/

package main

import "fmt"

func Ups() any {

	return "Ups"
}

func main() {
	interfaceKosong := Ups()
	fmt.Println(interfaceKosong)

}
