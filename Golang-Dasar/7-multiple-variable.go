/* Deklarasi Multiple Vaiable
di Golang kita bisa membuat variable secara banyak sekaligus
code yang dibuat akan lebih bagus dan mudah di baca
di golang variable harus digunakan, jika tidak digunakan akan terjadi error
*/ 

package main

import "fmt"

func main() {
var (
	fristName = "Misbah"
	lastName = "Munir"
)

fmt.Println(fristName)
fmt.Println(lastName)

}

