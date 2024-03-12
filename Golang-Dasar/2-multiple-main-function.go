/* Multiple Main Function
di Golang, function dalam module/project adalah unik. 
artinya tidak boleh membuat nama function yang sama.
oleh karena itu, jika membuat file baru misal contoh.go, 
lalu membuat nama function yang sama yaitu main
maka kita tidak bisa melakukan build module, karena main function 
tersebut duplikat dengan yang ada di main function sebelumnya.
*/

/* Solusi
karena masih dalam fase belajar, maka solusinya bukan dengan cara build project module terlebih dahulu.
sehingga kita hanya akan menjalankan file Golang satu persatu sehingga todak akan terjadi eror.
tapi, ingat saat membuat project gunakan satu main function saja.
Caranya dengan go run
*/

package main 

import "fmt"

func main () {
	fmt.Println("ini multiple function")
}