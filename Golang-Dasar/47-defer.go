/* Defer
defer adalah function yang bisa kita jadwalkan untuk dieksekusi setelah
sebuah function selesai di eksekusi
defer function akan selalu dieksekusi walaupun terjadi
error di function yang dieksekusi.
*/

package main

import "fmt"

func login() {
	fmt.Println("mencoba defer")
}

func percobaan() {

	defer login()
	fmt.Println("Sedang menjalankan Percobaan")

}

func main() {
	percobaan()
}
