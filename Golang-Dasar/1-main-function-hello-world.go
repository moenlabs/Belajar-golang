/* Main Function

Go-lang mirip seperti bahasa pemograman C/C++
di mana perlu ada yang namanya main function.
main function adalah sebuah fungsi yang dijalankan 
ketika program  berjalan.
untuk membuat function bisa menggunakan kata kunci func
Main function harus terdapat di dalam main package
Titik koma digolang tidak wajib.
huruf besar kecil berpengaruh (seperti pada javascript).
*/

/* Println
untuk menulis tulisan, perlu melakukan import "fmt" (singkatan dari format) terlebih dahulu
materi tentang umport akan dibahas di bagian tersendiri
*/

// Contoh Program
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("Halo Dunia")
}

// untuk menjalankan program dengan mengkompilasi dengan kode go build kemudian jalankan dengan ./ nama program
// atau menjalankan tanpa mengkompilasi dengan cara go run kemudian nama file, namun cocoknya ketika pada fase pengembangan (belum produk jadi)
